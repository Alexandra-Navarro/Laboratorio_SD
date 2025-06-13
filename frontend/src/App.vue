<template>
  <v-app>
    <!-- Navigation Drawer y App Bar se muestran en todas las rutas excepto login -->
    <template v-if="!isLoginPage">
      <!-- Navigation Drawer -->
      <v-navigation-drawer v-model="drawer" app :width="280" class="elevation-3">
        <!-- Header del drawer -->
        <v-list-item class="px-2 py-4" to="/home" link>
          <v-list-item-avatar>
            <v-icon size="40" color="primary">mdi-home-analytics</v-icon>
          </v-list-item-avatar>
          <v-list-item-content>
            <v-list-item-title class="text-h6 font-weight-bold">
              Monitoreo
            </v-list-item-title>
            <v-list-item-subtitle>Ambiental</v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>

        <v-divider></v-divider>

        <v-list nav dense>
          <!-- Dashboard General 
          <v-list-item to="/dashboard" link class="mb-2">
            <v-list-item-icon>
              <v-icon>mdi-view-dashboard</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title>Dashboard General</v-list-item-title>
            </v-list-item-content>
          </v-list-item>-->

          <!-- Todas las Alertas -->
          <v-list-item to="/alertas" link class="mb-2">
            <v-list-item-icon>
              <v-icon>mdi-bell-ring</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title>Todas las Alertas</v-list-item-title>
            </v-list-item-content>
            <!-- Badge de alertas críticas globales -->
            <v-list-item-action v-if="totalAlertasCriticas > 0">
              <v-chip color="error" text-color="white" small class="ml-2">
                {{ totalAlertasCriticas }}
              </v-chip>
            </v-list-item-action>
          </v-list-item>

          <v-divider class="my-2"></v-divider>

          <!-- Grupo de Salas -->
          <v-list-group :value="true" prepend-icon="mdi-home-city" no-action>
            <template v-slot:activator="{ props }">
              <v-list-item v-bind="props">
                <v-list-item-content>
                  <v-list-item-title class="font-weight-medium">
                    Salas ({{ salas.length }})
                  </v-list-item-title>
                </v-list-item-content>
              </v-list-item>
            </template>

            <!-- Loading state para salas -->
            <v-list-item v-if="loading" class="ml-4">
              <v-list-item-icon>
                <v-progress-circular indeterminate size="20" color="primary"></v-progress-circular>
              </v-list-item-icon>
              <v-list-item-content>
                <v-list-item-title class="text-caption">
                  Cargando salas...
                </v-list-item-title>
              </v-list-item-content>
            </v-list-item>

            <!-- Error state para salas -->
            <v-list-item v-else-if="error" class="ml-4">
              <v-list-item-icon>
                <v-icon color="error" size="20">mdi-alert-circle</v-icon>
              </v-list-item-icon>
              <v-list-item-content>
                <v-list-item-title class="text-error text-caption">
                  {{ error }}
                </v-list-item-title>
              </v-list-item-content>
              <v-list-item-action>
                <v-btn icon small @click="obtenerSalas" :loading="loading">
                  <v-icon small>mdi-refresh</v-icon>
                </v-btn>
              </v-list-item-action>
            </v-list-item>

            <!-- Lista de salas -->
            <template v-else>
              <div v-for="sala in salas" :key="sala.id">
                <!-- Subgrupo para cada sala -->
                <v-list-group :value="isCurrentSala(sala.id)" sub-group no-action class="ml-4">
                  <template v-slot:activator="{ props }">
                    <v-list-item v-bind="props">
                      <v-list-item-icon>
                        <v-icon size="20">{{ sala.icon }}</v-icon>
                      </v-list-item-icon>
                      <v-list-item-content>
                        <v-list-item-title class="text-sm">
                          {{ sala.nombre }}
                        </v-list-item-title>
                      </v-list-item-content>
                      <!-- Badge de alertas por sala -->
                      <v-list-item-action v-if="sala.alertasCriticas > 0">
                        <v-chip color="error" text-color="white" x-small>
                          {{ sala.alertasCriticas }}
                        </v-chip>
                      </v-list-item-action>
                    </v-list-item>
                  </template>

                  <!-- Opciones de cada sala -->
                  <v-list-item :to="sala.path" link class="ml-8" exact>
                    <v-list-item-icon>
                      <v-icon size="18">mdi-monitor-dashboard</v-icon>
                    </v-list-item-icon>
                    <v-list-item-content>
                      <v-list-item-title class="text-sm">Dashboard</v-list-item-title>
                    </v-list-item-content>
                  </v-list-item>

                  <v-list-item :to="`/salas/${sala.id}/alertas`" link class="ml-8" exact>
                    <v-list-item-icon>
                      <v-icon size="18">mdi-bell-alert</v-icon>
                    </v-list-item-icon>
                    <v-list-item-content>
                      <v-list-item-title class="text-sm">Alertas</v-list-item-title>
                    </v-list-item-content>
                    <v-list-item-action v-if="sala.alertasCriticas > 0">
                      <v-chip color="error" text-color="white" x-small>
                        {{ sala.alertasCriticas }}
                      </v-chip>
                    </v-list-item-action>
                  </v-list-item>
                </v-list-group>
              </div>
            </template>
          </v-list-group>
        </v-list>

        <!-- Footer del drawer -->
        <template v-slot:append>
          <div class="pa-2">
            <v-btn block outlined small @click="obtenerSalas" :loading="loading">
              <v-icon left small>mdi-refresh</v-icon>
              Actualizar
            </v-btn>
          </div>
        </template>
      </v-navigation-drawer>

      <!-- App Bar -->
      <v-app-bar app color="primary" dark elevation="4" :class="{ 'blur-background': drawer }">
        <v-app-bar-nav-icon @click="drawer = !drawer" class="mr-2"></v-app-bar-nav-icon>

        <v-toolbar-title class="font-weight-bold">
          {{ getPageTitle() }}
        </v-toolbar-title>

        <v-spacer></v-spacer>

        <!-- Indicador de conexión 
        <v-chip :color="isOnline ? 'success' : 'error'" small class="mr-2">
          <v-icon left small>
            {{ isOnline ? 'mdi-wifi' : 'mdi-wifi-off' }}
          </v-icon>
          {{ isOnline ? 'Conectado' : 'Sin conexión' }}
        </v-chip>-->

        <!-- Botón de notificaciones -->
        <v-btn icon @click="requestNotificationPermission" v-if="!notificationsEnabled">
          <v-icon>mdi-bell-off</v-icon>
        </v-btn>
        <v-btn @click="logout" color="red" variant="elevated" class="logout-btn" rounded>
          <v-icon left>mdi-logout</v-icon>
          Cerrar sesión
        </v-btn>
      </v-app-bar>
    </template>

    <!-- Main Content - Siempre visible -->
    <v-main>
      <v-container fluid class="pa-0">
        <transition name="fade" mode="out-in">
          <router-view :key="$route.fullPath"></router-view>
        </transition>
      </v-container>
    </v-main>

    <!-- Snackbar para notificaciones - Siempre visible -->
    <v-snackbar v-model="snackbar.show" :color="snackbar.color" :timeout="snackbar.timeout" bottom right>
      {{ snackbar.text }}
      <template v-slot:action="{ attrs }">
        <v-btn text v-bind="attrs" @click="snackbar.show = false">
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

export default {
  name: 'App',
  setup() {
    const router = useRouter()
    const route = useRoute()
    const drawer = ref(true)

    const isLoggedIn = computed(() => !!localStorage.getItem('usuario'))
    const isLoginPage = computed(() => route.path === '/')

    return {
      router,
      drawer,
      isLoggedIn,
      isLoginPage
    }
  },
  data() {
    return {
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
    logout() {
      localStorage.removeItem('usuario')
      this.$router.push('/')
    },
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

        //this.showSnackbar('Salas actualizadas correctamente', 'success')
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
        //'Dashboard': 'Dashboard General',
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

    /*startPolling() {
      // Actualizar salas cada 30 segundos
      this.pollingInterval = setInterval(() => {
        if (this.isOnline) {
          this.obtenerSalas()
        }
      }, 30000)
    },*/

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

          notification.onclick = function () {
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
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
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

.logout-btn {
  font-weight: bold;
  letter-spacing: 1px;
  box-shadow: 0 2px 8px rgba(255, 0, 0, 0.10);
  transition: box-shadow 0.2s, background 0.2s;
  background: linear-gradient(90deg, #ff5252 0%, #ff1744 100%);
  color: white !important;
}
.logout-btn:hover {
  box-shadow: 0 4px 16px rgba(255, 23, 68, 0.18);
  background: linear-gradient(90deg, #ff1744 0%, #ff5252 100%);
}
</style>