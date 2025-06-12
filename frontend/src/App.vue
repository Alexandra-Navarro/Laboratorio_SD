<template>
  <v-app>
    <!-- Solo mostrar navegación si el usuario está logueado y no está en login -->
    <template v-if="isLoggedIn && !isLoginPage">
      <v-navigation-drawer v-model="drawer" app>
        <v-list>
          <v-list-item
            v-for="item in menuItems"
            :key="item.title"
            :to="item.path"
            link
          >
            <v-list-item-icon>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title>{{ item.title }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-navigation-drawer>

      <v-app-bar app color="primary" dark>
        <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
        <v-toolbar-title>Sistema de Monitoreo Ambiental</v-toolbar-title>
        <v-btn @click="logout" color="red" variant="tonal">Cerrar sesión</v-btn>
      </v-app-bar>
    </template>

    <!-- Main Content -->
    <v-main>
      <v-container fluid class="pa-0">
        <transition name="fade" mode="out-in">
          <router-view :key="$route.fullPath"></router-view>
        </transition>
      </v-container>
    </v-main>

    <!-- Snackbar para notificaciones -->
    <v-snackbar
      v-model="snackbar.show"
      :color="snackbar.color"
      :timeout="snackbar.timeout"
      bottom
      right
    >
      {{ snackbar.text }}
      <template v-slot:action="{ attrs }">
        <v-btn
          text
          v-bind="attrs"
          @click="snackbar.show = false"
        >
          Cerrar
        </v-btn>
      </template>
    </v-snackbar>
  </v-app>
</template>

<script>
import axios from 'axios'
import { ref, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()
const drawer = ref(true)

const isLoggedIn = computed(() => !!localStorage.getItem('usuario'))
const isLoginPage = computed(() => route.path === '/')

const menuItems = [
  {
    title: 'Dashboard',
    icon: 'mdi-view-dashboard',
    path: '/dashboard'
  },
  {
    title: 'Sensores',
    icon: 'mdi-thermometer',
    path: '/sensors'
  },
  {
    title: 'Alertas',
    icon: 'mdi-alert',
    path: '/alerts'
  },
  {
    title: 'Configuración',
    icon: 'mdi-cog',
    path: '/settings'
  }
]

const logout = () => {
  localStorage.removeItem('usuario')
  router.push('/')

export default {
  name: 'App',
  data() {
    return {
      drawer: true,
      salas: [],
      loading: false,
      error: null,
      totalAlertasCriticas: 0,
      isOnline: navigator.onLine,
      notificationsEnabled: false,
      snackbar: {
        show: false,
        text: '',
        color: 'info',
        timeout: 4000
      },
      pollingInterval: null
    }
  },
  async created() {
    await this.obtenerSalas()
    this.checkNotificationPermission()
    this.setupConnectionMonitoring()
    this.startPolling()
  },
  beforeUnmount() {
    this.stopPolling()
  },
  methods: {
    async obtenerSalas() {
      try {
        this.loading = true
        this.error = null
        
        const response = await axios.get('/api/salas')
        
        if (!Array.isArray(response.data)) {
          throw new Error('La respuesta de la API no es un array válido')
        }
        
        // Mapear las salas y obtener alertas críticas
        this.salas = await Promise.all(response.data.map(async (sala) => {
          const salaData = {
            id: sala.id,
            nombre: sala.nombre,
            icon: this.obtenerIconoPorEscuela(sala.escuela_id),
            path: `/salas/${sala.id}`,
            alertasCriticas: 0
          }
          
          // Obtener alertas críticas para cada sala
          try {
            const alertasResponse = await axios.get(`/api/alertas/sala/${sala.id}?tipo=critico`)
            salaData.alertasCriticas = alertasResponse.data.length || 0
          } catch (error) {
            console.warn(`No se pudieron cargar alertas para sala ${sala.id}:`, error)
          }
          
          return salaData
        }))

        // Calcular total de alertas críticas
        this.totalAlertasCriticas = this.salas.reduce((total, sala) => total + sala.alertasCriticas, 0)
        
        this.showSnackbar('Salas actualizadas correctamente', 'success')
      } catch (error) {
        console.error('Error al obtener salas:', error)
        this.error = 'No se pudieron cargar las salas'
        
        // Salas de respaldo en caso de error
        this.salas = [
          {
            id: 1,
            nombre: 'Sala por defecto',
            icon: 'mdi-home-city',
            path: '/salas/1',
            alertasCriticas: 0
          }
        ]
        
        this.showSnackbar('Error al cargar salas', 'error')
      } finally {
        this.loading = false
      }
    },

    obtenerIconoPorEscuela(escuela_id) {
      const iconos = {
        1: 'mdi-school',
        2: 'mdi-book-open-variant',
        3: 'mdi-account-group',
        4: 'mdi-city',
        5: 'mdi-office-building',
        'default': 'mdi-door'
      }
      return iconos[escuela_id] || iconos.default
    },

    getPageTitle() {
      const routeTitles = {
        'Home': 'Sistema de Monitoreo Ambiental',
        'Dashboard': 'Dashboard General',
        'SalaDetail': `Sala ${this.$route.params.id || ''}`,
        'SalaAlertas': `Alertas - Sala ${this.$route.params.id || ''}`,
        'TodasAlertas': 'Todas las Alertas'
      }
      return routeTitles[this.$route.name] || 'Sistema de Monitoreo Ambiental'
    },

    isCurrentSala(salaId) {
      return this.$route.params.id == salaId
    },

    checkNotificationPermission() {
      if ('Notification' in window) {
        console.log('Estado actual de notificaciones:', Notification.permission);
        if (Notification.permission === 'default') {
          Notification.requestPermission().then(permission => {
            console.log('Permiso de notificación:', permission);
            this.notificationsEnabled = permission === 'granted';
            if (permission === 'granted') {
              this.showSnackbar('Notificaciones activadas', 'success');
              // Mostrar una notificación de prueba
              this.showNotification('Notificaciones activadas', 'El sistema de notificaciones está funcionando correctamente');
            }
          });
        } else {
          this.notificationsEnabled = Notification.permission === 'granted';
          console.log('Notificaciones ya configuradas:', this.notificationsEnabled);
        }
      } else {
        console.log('Las notificaciones no están soportadas en este navegador');
      }
    },

    async requestNotificationPermission() {
      if ('Notification' in window) {
        const permission = await Notification.requestPermission()
        this.notificationsEnabled = permission === 'granted'
        
        if (this.notificationsEnabled) {
          this.showSnackbar('Notificaciones habilitadas', 'success')
        } else {
          this.showSnackbar('Notificaciones denegadas', 'warning')
        }
      }
    },

    setupConnectionMonitoring() {
      window.addEventListener('online', () => {
        this.isOnline = true
        this.showSnackbar('Conexión restaurada', 'success')
      })
      
      window.addEventListener('offline', () => {
        this.isOnline = false
        this.showSnackbar('Sin conexión a internet', 'error')
      })
    },

    startPolling() {
      // Actualizar salas cada 30 segundos
      this.pollingInterval = setInterval(() => {
        if (this.isOnline) {
          this.obtenerSalas()
        }
      }, 30000)
    },

    stopPolling() {
      if (this.pollingInterval) {
        clearInterval(this.pollingInterval)
        this.pollingInterval = null
      }
    },

    showSnackbar(text, color = 'info', timeout = 4000) {
      this.snackbar = {
        show: true,
        text,
        color,
        timeout
      }
    },

    showNotification(title, body) {
      console.log('Intentando mostrar notificación:', { title, body, notificationsEnabled: this.notificationsEnabled });
      if (this.notificationsEnabled && 'Notification' in window) {
        try {
          const notification = new Notification(title, {
            body: body,
            icon: '/favicon.ico',
            badge: '/favicon.ico',
            tag: 'alerta',
            requireInteraction: true
          });
          
          notification.onclick = function() {
            window.focus();
            this.close();
          };
          
          console.log('Notificación creada exitosamente');
        } catch (error) {
          console.error('Error al mostrar notificación:', error);
        }
      } else {
        console.log('No se puede mostrar la notificación:', {
          notificationsEnabled: this.notificationsEnabled,
          notificationsSupported: 'Notification' in window
        });
      }
    }
  }
}
</script>

<style>
@import '@mdi/font/css/materialdesignicons.css';

/* Transiciones suaves */
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from, .fade-leave-to {
  opacity: 0;
}

/* Efecto blur en app bar cuando drawer está abierto */
.blur-background {
  backdrop-filter: blur(2px);
}

/* Mejoras visuales para el drawer */
.v-navigation-drawer {
  border-right: 1px solid rgba(0, 0, 0, 0.12) !important;
}

/* Estilos para chips pequeños */
.v-chip.v-size--x-small {
  height: 18px !important;
  font-size: 10px !important;
  min-width: 18px !important;
}

/* Hover effects para list items */
.v-list-item:hover {
  background-color: rgba(0, 0, 0, 0.04) !important;
}

/* Estilo activo para rutas */
.v-list-item--active {
  color: var(--v-primary-base) !important;
}

.v-list-item--active .v-list-item__icon {
  color: var(--v-primary-base) !important;
}
</style>