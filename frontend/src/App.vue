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
          {{ currentPageTitle }}
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
        <v-btn icon @click="requestNotificationPermission" :color="notificationsEnabled ? 'success' : 'warning'">
          <v-icon>{{ notificationsEnabled ? 'mdi-bell' : 'mdi-bell-off' }}</v-icon>
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
          <router-view :key="$route.fullPath" v-slot="{ Component, route }">
            <component 
              :is="Component" 
              v-bind="route.meta || {}" 
              :salas="salas"
              :salas-cargadas="!loading && salas.length > 0"
            />
          </router-view>
        </transition>
      </v-container>
    </v-main>

    <!-- Snackbar para notificaciones - Siempre visible -->
    <v-snackbar v-model="snackbar.show" :color="snackbar.color" :timeout="snackbar.timeout" top right>
      {{ snackbar.text }}
      <template v-slot:action="{ attrs }">
        <v-btn icon v-bind="attrs" @click="snackbar.show = false">
          <v-icon>mdi-close</v-icon>
        </v-btn>
      </template>
    </v-snackbar>

    <!-- Sistema de notificaciones Toast -->
    <div class="notifications-container">
      <transition-group name="notification" tag="div">
        <div 
          v-for="notification in notificacionesToast" 
          :key="notification.id"
          class="notification-toast"
          :class="notification.type"
        >
          <div class="notification-icon">
            <v-icon :color="getNotificationIconColor(notification.type)" size="24">
              {{ getNotificationIcon(notification.type) }}
            </v-icon>
          </div>
          <div class="notification-content">
            <div class="notification-title">{{ notification.title }}</div>
            <div class="notification-message">{{ notification.message }}</div>
            <div class="notification-time">{{ formatNotificationTime(notification.timestamp) }}</div>
          </div>
          <div class="notification-actions">
            <v-btn 
              icon 
              small 
              @click="dismissNotification(notification.id)"
              class="close-btn"
            >
              <v-icon size="18">mdi-close</v-icon>
            </v-btn>
          </div>
        </div>
      </transition-group>
    </div>
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
      pollingInterval: null,
      // Nuevo sistema de notificaciones tipo toast
      notificacionesToast: [],
      nextNotificationId: 1,
      // Sistema de deduplicación de alertas
      alertasNotificadas: new Set() // Para evitar duplicados
    }
  },
  computed: {
    currentPageTitle() {
      const routeTitles = {
        'Home': 'Sistema de Monitoreo Ambiental',
        'SalaDetail': this.getSalaTitleById(this.$route.params.id),
        'SalaAlertas': `Alertas - ${this.getSalaTitleById(this.$route.params.id)}`,
        'TodasAlertas': 'Todas las Alertas'
      }
      return routeTitles[this.$route.name] || 'Sistema de Monitoreo Ambiental'
    }
  },
  watch: {
    // Los watchers ya no son necesarios porque la computed property es reactiva automáticamente
    // Pero mantenemos uno simple para debug
    salas: {
      handler() {
        console.log('Salas actualizadas, títulos se actualizarán automáticamente');
      }
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

          // Debug: verificar datos de cada sala
          console.log('Sala procesada:', salaData);

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

        console.log('Salas finales cargadas:', this.salas);
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

    getSalaTitleById(salaId) {
      if (!salaId) return 'Sala';
      
      // Si las salas aún no se han cargado, mostrar el ID temporalmente
      if (!this.salas || this.salas.length === 0) {
        return `Sala ${salaId}`;
      }
      
      // Buscar la sala en el array de salas cargadas
      const sala = this.salas.find(s => s.id == salaId);
      if (sala && sala.nombre) {
        return `Sala ${sala.nombre}`;
      }
      
      // Fallback al ID si no se encuentra la sala
      return `Sala ${salaId}`;
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

    showSnackbar(text, color = 'info') {
      this.snackbar = {
        show: true,
        text,
        color,
        timeout: color === 'error' ? 6000 : 4000 // Más tiempo para alertas críticas
      }
    },

    showNotification(title, body) {
      console.log('Intentando mostrar notificación:', { title, body, notificationsEnabled: this.notificationsEnabled });
      
      // Notificación nativa del navegador
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

          console.log('Notificación nativa creada exitosamente');
        } catch (error) {
          console.error('Error al mostrar notificación nativa:', error);
        }
      }

      // Mostrar notificación toast en la página
      this.showToastNotification(title, body, 'alert');
    },

    showToastNotification(title, message, type = 'info', alertId = null) {
      // Verificar duplicados si se proporciona alertId
      if (alertId) {
        const alertKey = `${alertId}-${type}`;
        if (this.alertasNotificadas.has(alertKey)) {
          console.log('Notificación duplicada evitada:', alertKey);
          return; // No mostrar la notificación duplicada
        }
        this.alertasNotificadas.add(alertKey);
        
        // Limpiar alertas antiguas del set (mantener solo las últimas 50)
        if (this.alertasNotificadas.size > 50) {
          const alertasArray = Array.from(this.alertasNotificadas);
          const nuevasAlertas = alertasArray.slice(-30); // Mantener solo las últimas 30
          this.alertasNotificadas = new Set(nuevasAlertas);
        }
      }

      const notification = {
        id: this.nextNotificationId++,
        title,
        message,
        type,
        timestamp: new Date(),
        alertId // Guardamos el ID para referencia
      };

      this.notificacionesToast.unshift(notification);

      // Auto-dismiss después de 8 segundos para alertas críticas, 5 para otras
      const timeout = type === 'critical' ? 8000 : 5000;
      setTimeout(() => {
        this.dismissNotification(notification.id);
      }, timeout);

      // Limitar a máximo 3 notificaciones visibles
      if (this.notificacionesToast.length > 3) {
        this.notificacionesToast = this.notificacionesToast.slice(0, 3);
      }
    },

    dismissNotification(id) {
      const index = this.notificacionesToast.findIndex(n => n.id === id);
      if (index > -1) {
        const notification = this.notificacionesToast[index];
        this.notificacionesToast.splice(index, 1);
        
        // Opcional: remover del set de notificadas después de un tiempo
        // para permitir re-notificaciones de la misma alerta más tarde
        if (notification.alertId) {
          setTimeout(() => {
            const alertKey = `${notification.alertId}-${notification.type}`;
            this.alertasNotificadas.delete(alertKey);
          }, 30000); // Permitir re-notificación después de 30 segundos
        }
      }
    },

    getNotificationIcon(type) {
      const icons = {
        'alert': 'mdi-alert-circle',
        'critical': 'mdi-alert-octagon',
        'warning': 'mdi-alert',
        'info': 'mdi-information',
        'success': 'mdi-check-circle'
      };
      return icons[type] || icons.info;
    },

    getNotificationIconColor(type) {
      const colors = {
        'alert': '#ff9800',
        'critical': '#f44336',
        'warning': '#ff9800',
        'info': '#2196f3',
        'success': '#4caf50'
      };
      return colors[type] || colors.info;
    },

    formatNotificationTime(timestamp) {
      const now = new Date();
      const diff = now - timestamp;
      const seconds = Math.floor(diff / 1000);
      
      if (seconds < 60) {
        return 'Ahora';
      } else if (seconds < 3600) {
        const minutes = Math.floor(seconds / 60);
        return `Hace ${minutes} min`;
      } else {
        return timestamp.toLocaleTimeString('es-CL', { 
          hour: '2-digit', 
          minute: '2-digit' 
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

.v-snackbar--top {
  margin-top: 24px !important;
}

/* Sistema de notificaciones Toast */
.notifications-container {
  position: fixed;
  top: 80px;
  right: 20px;
  z-index: 9999;
  width: 400px;
  max-width: calc(100vw - 40px);
}

.notification-toast {
  background: white;
  border-radius: 12px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.12);
  border: 1px solid rgba(0, 0, 0, 0.08);
  padding: 16px;
  margin-bottom: 12px;
  display: flex;
  align-items: flex-start;
  gap: 12px;
  backdrop-filter: blur(8px);
  position: relative;
  overflow: hidden;
}

.notification-toast::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: #2196f3;
}

.notification-toast.alert::before {
  background: #ff9800;
}

.notification-toast.critical::before {
  background: #f44336;
}

.notification-toast.warning::before {
  background: #ff9800;
}

.notification-toast.success::before {
  background: #4caf50;
}

.notification-icon {
  flex-shrink: 0;
  margin-top: 2px;
}

.notification-content {
  flex: 1;
  min-width: 0;
}

.notification-title {
  font-weight: 600;
  font-size: 14px;
  color: #333;
  margin-bottom: 4px;
  line-height: 1.3;
}

.notification-message {
  font-size: 13px;
  color: #666;
  line-height: 1.4;
  margin-bottom: 6px;
  word-wrap: break-word;
}

.notification-time {
  font-size: 11px;
  color: #999;
  font-weight: 500;
}

.notification-actions {
  flex-shrink: 0;
}

.close-btn {
  opacity: 0.6;
  transition: opacity 0.2s ease;
}

.close-btn:hover {
  opacity: 1;
}

/* Animaciones para las notificaciones */
.notification-enter-active {
  transition: all 0.4s cubic-bezier(0.34, 1.56, 0.64, 1);
}

.notification-leave-active {
  transition: all 0.3s ease-out;
}

.notification-enter-from {
  opacity: 0;
  transform: translateX(100%) scale(0.8);
}

.notification-leave-to {
  opacity: 0;
  transform: translateX(100%) scale(0.9);
}

.notification-move {
  transition: all 0.3s ease;
}

/* Responsive para móviles */
@media (max-width: 768px) {
  .notifications-container {
    top: 70px;
    right: 10px;
    left: 10px;
    width: auto;
    max-width: none;
  }
  
  .notification-toast {
    padding: 12px;
    margin-bottom: 8px;
  }
  
  .notification-title {
    font-size: 13px;
  }
  
  .notification-message {
    font-size: 12px;
  }
}

/* Efectos hover para las notificaciones */
.notification-toast:hover {
  transform: translateY(-2px);
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.15);
  transition: all 0.2s ease;
}
</style>