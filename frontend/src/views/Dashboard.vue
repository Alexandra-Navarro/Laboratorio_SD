<template>
  <div class="dashboard">
    <!-- Header del Dashboard -->
    <div class="dashboard-header">
      <div class="header-content">
        <h1 class="dashboard-title">Dashboard General</h1>
        <p class="dashboard-subtitle">
          Vista general del sistema de monitoreo ambiental
        </p>
        <div class="header-actions">
          <v-btn
            color="white"
            outlined
            @click="actualizarDatos"
            :loading="loading"
            class="mr-2"
          >
            <v-icon left>mdi-refresh</v-icon>
            Actualizar
          </v-btn>
          <v-chip
            :color="isOnline ? 'success' : 'error'"
            dark
          >
            <v-icon left small>
              {{ isOnline ? 'mdi-wifi' : 'mdi-wifi-off' }}
            </v-icon>
            {{ isOnline ? 'Conectado' : 'Sin conexión' }}
          </v-chip>
        </div>
      </div>
    </div>

    <v-container class="dashboard-content" fluid>
      <!-- Resumen Estadísticas -->
      <v-row class="mb-6">
        <v-col cols="12" sm="6" md="3">
          <v-card class="stat-card">
            <v-card-text class="text-center">
              <v-icon size="48" color="primary" class="mb-2">mdi-home-city</v-icon>
              <div class="stat-number">{{ estadisticas.totalSalas }}</div>
              <div class="stat-label">Salas Activas</div>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3">
          <v-card class="stat-card">
            <v-card-text class="text-center">
              <v-icon size="48" color="error" class="mb-2">mdi-alert-circle</v-icon>
              <div class="stat-number">{{ estadisticas.alertasCriticas }}</div>
              <div class="stat-label">Alertas Críticas</div>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3">
          <v-card class="stat-card">
            <v-card-text class="text-center">
              <v-icon size="48" color="success" class="mb-2">mdi-check-circle</v-icon>
              <div class="stat-number">{{ estadisticas.sensoresOnline }}</div>
              <div class="stat-label">Sensores Online</div>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3">
          <v-card class="stat-card">
            <v-card-text class="text-center">
              <v-icon size="48" color="info" class="mb-2">mdi-chart-line</v-icon>
              <div class="stat-number">{{ estadisticas.medicionesHoy }}</div>
              <div class="stat-label">Mediciones Hoy</div>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>

      <!-- Grid de Salas -->
      <v-row>
        <v-col cols="12" md="8">
          <v-card>
            <v-card-title>
              <v-icon left>mdi-home-city</v-icon>
              Estado de las Salas
              <v-spacer></v-spacer>
              <v-text-field
                v-model="busquedaSala"
                append-icon="mdi-magnify"
                label="Buscar sala"
                single-line
                hide-details
                dense
                outlined
                style="max-width: 250px"
              ></v-text-field>
            </v-card-title>
            <v-card-text>
              <v-row v-if="loading">
                <v-col v-for="i in 6" :key="i" cols="12" sm="6" md="4">
                  <v-skeleton-loader type="card"></v-skeleton-loader>
                </v-col>
              </v-row>
              <v-row v-else>
                <v-col 
                  v-for="sala in salasFiltradas" 
                  :key="sala.id" 
                  cols="12" 
                  sm="6" 
                  md="4"
                  class="mb-3"
                >
                  <v-card 
                    class="sala-card" 
                    :class="getSalaCardClass(sala)"
                    hover
                    @click="irASala(sala.id)"
                  >
                    <v-card-text class="pa-3">
                      <div class="d-flex align-center mb-2">
                        <v-icon 
                          :color="getSalaStatusColor(sala)" 
                          class="mr-2"
                        >
                          {{ sala.icon }}
                        </v-icon>
                        <div class="flex-grow-1">
                          <div class="sala-name">{{ sala.nombre }}</div>
                          <div class="sala-status text-caption">
                            {{ getSalaStatusText(sala) }}
                          </div>
                        </div>
                        <v-menu offset-y>
                          <template v-slot:activator="{ on, attrs }">
                            <v-btn
                              icon
                              small
                              v-bind="attrs"
                              v-on="on"
                              @click.stop
                            >
                              <v-icon>mdi-dots-vertical</v-icon>
                            </v-btn>
                          </template>
                          <v-list dense>
                            <v-list-item @click="irASala(sala.id)">
                              <v-list-item-icon>
                                <v-icon>mdi-eye</v-icon>
                              </v-list-item-icon>
                              <v-list-item-title>Ver Detalle</v-list-item-title>
                            </v-list-item>
                            <v-list-item @click="irAAlertas(sala.id)">
                              <v-list-item-icon>
                                <v-icon>mdi-bell</v-icon>
                              </v-list-item-icon>
                              <v-list-item-title>Ver Alertas</v-list-item-title>
                            </v-list-item>
                          </v-list>
                        </v-menu>
                      </div>
                      
                      <!-- Métricas de la sala -->
                      <div class="sala-metrics">
                        <div class="metric-row">
                          <v-icon small class="mr-1">mdi-thermometer</v-icon>
                          <span class="text-caption">{{ getRandomTemp() }}°C</span>
                        </div>
                        <div class="metric-row">
                          <v-icon small class="mr-1">mdi-water-percent</v-icon>
                          <span class="text-caption">{{ getRandomHumidity() }}%</span>
                        </div>
                      </div>
                      
                      <!-- Alertas -->
                      <div v-if="sala.alertasCriticas > 0" class="mt-2">
                        <v-chip
                          color="error"
                          text-color="white"
                          small
                          block
                        >
                          <v-icon left small>mdi-alert</v-icon>
                          {{ sala.alertasCriticas }} alerta(s) crítica(s)
                        </v-chip>
                      </div>
                    </v-card-text>
                  </v-card>
                </v-col>
              </v-row>
              
              <div v-if="!loading && salasFiltradas.length === 0" class="text-center py-8">
                <v-icon size="64" color="grey">mdi-home-search</v-icon>
                <h3 class="grey--text mt-2">No se encontraron salas</h3>
                <p class="grey--text">Intenta con un término de búsqueda diferente</p>
              </div>
            </v-card-text>
          </v-card>
        </v-col>
        
        <!-- Panel lateral -->
        <v-col cols="12" md="4">
          <!-- Alertas Recientes -->
          <v-card class="mb-4">
            <v-card-title>
              <v-icon left>mdi-bell-alert</v-icon>
              Alertas Recientes
            </v-card-title>
            <v-card-text class="pa-0">
              <AlertsComponent 
                :limit="3"
                @alertas-cargadas="actualizarEstadisticasAlertas"
              />
            </v-card-text>
            <v-card-actions>
              <v-btn
                text
                color="primary"
                to="/alertas"
                block
              >
                Ver Todas las Alertas
                <v-icon right>mdi-arrow-right</v-icon>
              </v-btn>
            </v-card-actions>
          </v-card>
          
          <!-- Estado del Sistema -->
          <v-card>
            <v-card-title>
              <v-icon left>mdi-cog</v-icon>
              Estado del Sistema
            </v-card-title>
            <v-card-text>
              <div class="system-status">
                <div class="status-item">
                  <v-icon color="success" class="mr-2">mdi-check-circle</v-icon>
                  <span>API Funcionando</span>
                </div>
                <div class="status-item">
                  <v-icon color="success" class="mr-2">mdi-check-circle</v-icon>
                  <span>Base de Datos OK</span>
                </div>
                <div class="status-item">
                  <v-icon color="warning" class="mr-2">mdi-alert-circle</v-icon>
                  <span>2 Sensores Offline</span>
                </div>
                <div class="status-item">
                  <v-icon color="success" class="mr-2">mdi-check-circle</v-icon>
                  <span>Notificaciones Activas</span>
                </div>
              </div>
              
              <v-divider class="my-3"></v-divider>
              
              <div class="text-caption grey--text">
                Última actualización: {{ formatTime(new Date()) }}
              </div>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import AlertsComponent from '@/components/Alerts.vue'

export default {
  name: 'Dashboard',
  components: {
    AlertsComponent
  },
  data() {
    return {
      loading: false,
      isOnline: navigator.onLine,
      busquedaSala: '',
      salas: [],
      estadisticas: {
        totalSalas: 0,
        alertasCriticas: 0,
        sensoresOnline: 0,
        medicionesHoy: 0,
        escuelas: 0
      }
    }
  }
}
</script>