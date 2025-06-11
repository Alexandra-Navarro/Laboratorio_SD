<template>
    <div class="sala-detail">
      <!-- Header de la sala -->
      <div class="sala-header">
        <div class="header-content">
          <div v-if="loading" class="header-loading">
            <v-skeleton-loader type="heading, subtitle"></v-skeleton-loader>
          </div>
          <div v-else-if="salaInfo" class="header-info">
            <h1 class="sala-title">{{ salaInfo.nombre }}</h1>
            <p class="sala-subtitle">
              <v-icon left>mdi-school</v-icon>
              {{ escuelaInfo?.nombre || 'Cargando escuela...' }}
            </p>
            <div class="sala-meta">
              <v-chip class="mr-2" small outlined>
                <v-icon left small>mdi-identifier</v-icon>
                ID: {{ salaInfo.id }}
              </v-chip>
              <v-chip small outlined>
                <v-icon left small>mdi-map-marker</v-icon>
                {{ escuelaInfo?.comuna || 'Ubicación' }}
              </v-chip>
            </div>
          </div>
          <div v-else class="header-error">
            <h1>Error al cargar la sala</h1>
            <p>No se pudo obtener la información de la sala.</p>
          </div>
        </div>
      </div>
  
      <!-- Navegación de tabs -->
      <div class="tabs-section">
        <v-container>
          <v-tabs v-model="tabActivo" class="sala-tabs">
            <v-tab>
              <v-icon left>mdi-view-dashboard</v-icon>
              Dashboard
            </v-tab>
            <v-tab>
              <v-icon left>mdi-bell-alert</v-icon>
              Alertas
              <v-chip 
                v-if="alertasCriticas > 0" 
                color="error" 
                text-color="white" 
                small 
                class="ml-2"
              >
                {{ alertasCriticas }}
              </v-chip>
            </v-tab>
            <v-tab>
              <v-icon left>mdi-chart-line</v-icon>
              Sensores
            </v-tab>
            <v-tab>
              <v-icon left>mdi-history</v-icon>
              Historial
            </v-tab>
          </v-tabs>
        </v-container>
      </div>
  
      <!-- Contenido de las tabs -->
      <v-container class="sala-content">
        <v-tabs-items v-model="tabActivo">
          <!-- Tab Dashboard -->
          <v-tab-item>
            <div class="dashboard-tab">
              <!-- Métricas rápidas -->
              <v-row class="mb-6">
                <v-col cols="12" sm="6" md="3">
                  <v-card class="metric-card">
                    <v-card-text class="text-center">
                      <div class="metric-icon">
                        <v-icon size="40" color="primary">mdi-thermometer</v-icon>
                      </div>
                      <div class="metric-value">{{ metricas.temperatura }}°C</div>
                      <div class="metric-label">Temperatura</div>
                    </v-card-text>
                  </v-card>
                </v-col>
                <v-col cols="12" sm="6" md="3">
                  <v-card class="metric-card">
                    <v-card-text class="text-center">
                      <div class="metric-icon">
                        <v-icon size="40" color="info">mdi-water-percent</v-icon>
                      </div>
                      <div class="metric-value">{{ metricas.humedad }}%</div>
                      <div class="metric-label">Humedad</div>
                    </v-card-text>
                  </v-card>
                </v-col>
                <v-col cols="12" sm="6" md="3">
                  <v-card class="metric-card">
                    <v-card-text class="text-center">
                      <div class="metric-icon">
                        <v-icon size="40" color="success">mdi-air-filter</v-icon>
                      </div>
                      <div class="metric-value">{{ metricas.calidadAire }}</div>
                      <div class="metric-label">Calidad del Aire</div>
                    </v-card-text>
                  </v-card>
                </v-col>
                <v-col cols="12" sm="6" md="3">
                  <v-card class="metric-card">
                    <v-card-text class="text-center">
                      <div class="metric-icon">
                        <v-icon size="40" color="warning">mdi-volume-high</v-icon>
                      </div>
                      <div class="metric-value">{{ metricas.ruido }} dB</div>
                      <div class="metric-label">Nivel de Ruido</div>
                    </v-card-text>
                  </v-card>
                </v-col>
              </v-row>
  
              <!-- Gráficos y datos adicionales -->
              <v-row>
                <v-col cols="12" md="8">
                  <v-card>
                    <v-card-title>
                      <v-icon left>mdi-chart-line</v-icon>
                      Tendencias de Variables Ambientales
                    </v-card-title>
                    <v-card-text>
                      <div class="chart-placeholder">
                        <p class="text-center text-h6 grey--text">
                          Gráfico de tendencias
                        </p>
                        <p class="text-center grey--text">
                          Aquí se mostraría un gráfico con las mediciones de las últimas 24 horas
                        </p>
                      </div>
                    </v-card-text>
                  </v-card>
                </v-col>
                <v-col cols="12" md="4">
                  <v-card class="mb-4">
                    <v-card-title>
                      <v-icon left>mdi-router-wireless</v-icon>
                      Estado de Sensores
                    </v-card-title>
                    <v-card-text>
                      <div class="sensor-status">
                        <div class="sensor-item">
                          <v-icon color="success" small>mdi-check-circle</v-icon>
                          <span class="ml-2">Sensor Temperatura</span>
                        </div>
                        <div class="sensor-item">
                          <v-icon color="success" small>mdi-check-circle</v-icon>
                          <span class="ml-2">Sensor Humedad</span>
                        </div>
                        <div class="sensor-item">
                          <v-icon color="warning" small>mdi-alert-circle</v-icon>
                          <span class="ml-2">Sensor CO2</span>
                        </div>
                        <div class="sensor-item">
                          <v-icon color="success" small>mdi-check-circle</v-icon>
                          <span class="ml-2">Sensor Ruido</span>
                        </div>
                      </div>
                    </v-card-text>
                  </v-card>
  
                  <v-card>
                    <v-card-title>
                      <v-icon left>mdi-clock-outline</v-icon>
                      Última Actualización
                    </v-card-title>
                    <v-card-text>
                      <p class="mb-2">
                        <strong>{{ formatDate(new Date()) }}</strong>
                      </p>
                      <p class="text-caption grey--text mb-0">
                        Los datos se actualizan cada 5 segundos
                      </p>
                    </v-card-text>
                  </v-card>
                </v-col>
              </v-row>
            </div>
          </v-tab-item>
  
          <!-- Tab Alertas -->
          <v-tab-item>
            <div class="alertas-tab">
              <AlertsComponent 
                :sala-id="salaId" 
                @nueva-alerta="manejarNuevaAlerta"
                @alertas-cargadas="actualizarConteoAlertas"
              />
            </div>
          </v-tab-item>
  
          <!-- Tab Sensores -->
          <v-tab-item>
            <div class="sensores-tab">
              <v-card>
                <v-card-title>
                  <v-icon left>mdi-router-wireless</v-icon>
                  Sensores Instalados
                </v-card-title>
                <v-card-text>
                  <p class="text-center text-h6 grey--text">
                    Lista de sensores
                  </p>
                  <p class="text-center grey--text">
                    Aquí se mostraría la lista detallada de todos los sensores instalados en la sala
                  </p>
                </v-card-text>
              </v-card>
            </div>
          </v-tab-item>
  
          <!-- Tab Historial -->
          <v-tab-item>
            <div class="historial-tab">
              <v-card>
                <v-card-title>
                  <v-icon left>mdi-history</v-icon>
                  Historial de Mediciones
                </v-card-title>
                <v-card-text>
                  <p class="text-center text-h6 grey--text">
                    Historial de datos
                  </p>
                  <p class="text-center grey--text">
                    Aquí se mostraría el historial completo de mediciones y eventos
                  </p>
                </v-card-text>
              </v-card>
            </div>
          </v-tab-item>
        </v-tabs-items>
      </v-container>
    </div>
  </template>
  
  <script>
  import AlertsComponent from '@/components/Alerts.vue'
  import axios from 'axios'
  
  export default {
    name: 'SalaDetail',
    components: {
      AlertsComponent
    },
    props: {
      id: {
        type: [String, Number],
        required: true
      }
    },
    data() {
      return {
        tabActivo: 0,
        salaInfo: null,
        escuelaInfo: null,
        loading: false,
        alertasCriticas: 0,
        metricas: {
          temperatura: '--',
          humedad: '--',
          calidadAire: '--',
          ruido: '--'
        }
      }
    },
    computed: {
      salaId() {
        return this.id || this.$route.params.id
      }
    },
    watch: {
      salaId: {
        immediate: true,
        handler() {
          this.cargarDatosSala()
        }
      }
    },
    methods: {
      async cargarDatosSala() {
        if (!this.salaId) return
        
        try {
          this.loading = true
          
          // Cargar información de la sala
          const salaResponse = await axios.get(`/api/salas/${this.salaId}`)
          this.salaInfo = salaResponse.data
          
          // Cargar información de la escuela
          if (this.salaInfo.escuela_id) {
            try {
              const escuelaResponse = await axios.get(`/api/escuelas/${this.salaInfo.escuela_id}`)
              this.escuelaInfo = escuelaResponse.data
            } catch (error) {
              console.warn('No se pudo cargar información de la escuela:', error)
            }
          }
          
          // Cargar métricas actuales (simuladas por ahora)
          this.cargarMetricas()
          
        } catch (error) {
          console.error('Error al cargar datos de la sala:', error)
        } finally {
          this.loading = false
        }
      },
      
      async cargarMetricas() {
        // Aquí deberías hacer llamadas a la API para obtener las métricas reales
        // Por ahora, simulamos algunos datos
        try {
          // Simular llamada a API de métricas
          this.metricas = {
            temperatura: (20 + Math.random() * 10).toFixed(1),
            humedad: (40 + Math.random() * 30).toFixed(0),
            calidadAire: ['Buena', 'Regular', 'Mala'][Math.floor(Math.random() * 3)],
            ruido: (30 + Math.random() * 20).toFixed(0)
          }
        } catch (error) {
          console.error('Error al cargar métricas:', error)
        }
      },
      
      manejarNuevaAlerta(alerta) {
        if (alerta.tipo === 'critico') {
          this.alertasCriticas++
        }
        
        // Mostrar notificación
        if ('Notification' in window && Notification.permission === 'granted') {
          new Notification(`Nueva Alerta - ${alerta.tipo}`, {
            body: alerta.descripcion,
            icon: '/favicon.ico'
          })
        }
      },
      
      actualizarConteoAlertas(alertas) {
        this.alertasCriticas = alertas.filter(a => a.tipo === 'critico').length
      },
      
      formatDate(date) {
        return new Date(date).toLocaleString('es-CL')
      }
    }
  }
  </script>
  
  <style scoped>
  .sala-detail {
    min-height: 100vh;
    background-color: #f5f5f5;
  }
  
  .sala-header {
    background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%);
    color: white;
    padding: 32px 0;
  }
  
  .header-content {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 24px;
  }
  
  .sala-title {
    font-size: 2.5rem;
    font-weight: 300;
    margin: 0 0 8px 0;
  }
  
  .sala-subtitle {
    font-size: 1.2rem;
    opacity: 0.9;
    margin: 0 0 16px 0;
    display: flex;
    align-items: center;
  }
  
  .sala-meta {
    display: flex;
    gap: 8px;
    flex-wrap: wrap;
  }
  
  .tabs-section {
    background: white;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    position: sticky;
    top: 64px;
    z-index: 10;
  }
  
  .sala-tabs {
    max-width: 1200px;
    margin: 0 auto;
  }
  
  .sala-content {
    max-width: 1200px;
    margin: 24px auto 0;
    padding-bottom: 24px;
  }
  
  .metric-card {
    height: 140px;
    transition: transform 0.2s ease;
  }
  
  .metric-card:hover {
    transform: translateY(-2px);
  }
  
  .metric-icon {
    margin-bottom: 8px;
  }
  
  .metric-value {
    font-size: 1.8rem;
    font-weight: bold;
    margin-bottom: 4px;
  }
  
  .metric-label {
    font-size: 0.875rem;
    opacity: 0.7;
    text-transform: uppercase;
    letter-spacing: 0.5px;
  }
  
  .chart-placeholder {
    height: 300px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    background: #f9f9f9;
    border-radius: 8px;
  }
  
  .sensor-status {
    display: flex;
    flex-direction: column;
    gap: 12px;
  }
  
  .sensor-item {
    display: flex;
    align-items: center;
  }
  
  /* Responsive */
  @media (max-width: 768px) {
    .sala-title {
      font-size: 2rem;
    }
    
    .header-content {
      padding: 0 16px;
    }
    
    .sala-content {
      padding: 0 16px 24px;
    }
  }
  </style>