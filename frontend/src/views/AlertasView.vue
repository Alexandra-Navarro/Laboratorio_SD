<template>
  <div class="alertas-view">
    <!-- Header de la vista -->
    <div class="view-header">
      <div class="header-content">
        <div class="header-text">
          <h1 class="view-title">
            {{ pageTitle }}
          </h1>
          <p class="view-subtitle" v-if="salaInfo">
            {{ salaInfo.nombre }}
          </p>
          <p class="view-subtitle" v-else-if="!salaId">
            Monitoreo de todas las salas del sistema
          </p>
        </div>
        
        <!-- Filtros y acciones -->
        <div class="header-actions">
          <v-card class="pa-4" outlined>
            <div class="d-flex align-center gap-4 flex-wrap">
              <!-- Filtro por tipo -->
              <v-select
                v-model="filtroTipo"
                :items="tiposAlerta"
                label="Filtrar por tipo"
                clearable
                dense
                outlined
                hide-details
                style="min-width: 150px"
              ></v-select>
              
              <!-- Filtro por fecha -->
              <v-menu
                v-model="menuFecha"
                :close-on-content-click="false"
                transition="scale-transition"
                offset-y
                min-width="auto"
              >
                <template v-slot:activator="{ on, attrs }">
                  <v-text-field
                    v-model="filtroFecha"
                    label="Fecha"
                    prepend-icon="mdi-calendar"
                    readonly
                    clearable
                    dense
                    outlined
                    hide-details
                    v-bind="attrs"
                    v-on="on"
                  ></v-text-field>
                </template>
                <v-date-picker
                  v-model="filtroFecha"
                  @input="menuFecha = false"
                ></v-date-picker>
              </v-menu>
              
              <!-- Botón de actualizar -->
              <v-btn
                color="primary"
                @click="refrescarAlertas"
                :loading="loading"
                icon
              >
                <v-icon>mdi-refresh</v-icon>
              </v-btn>
            </div>
          </v-card>
        </div>
      </div>
    </div>

    <!-- Estadísticas rápidas -->
    <div class="stats-section" v-if="estadisticas">
      <v-row>
        <v-col cols="12" sm="6" md="3">
          <v-card class="stat-card stat-total">
            <v-card-text class="text-center">
              <div class="stat-number">{{ estadisticas.total }}</div>
              <div class="stat-label">Total</div>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3">
          <v-card class="stat-card stat-criticas">
            <v-card-text class="text-center">
              <div class="stat-number">{{ estadisticas.criticas }}</div>
              <div class="stat-label">Críticas</div>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col cols="12" sm="6" md="3">
          <v-card class="stat-card stat-preventivas">
            <v-card-text class="text-center">
              <div class="stat-number">{{ estadisticas.preventivas }}</div>
              <div class="stat-label">Preventivas</div>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </div>

    <!-- Componente de alertas -->
    <div class="alerts-content">
      <AlertsComponent 
        :sala-id="salaId" 
        :filtro-tipo="filtroTipo"
        :filtro-fecha="filtroFecha"
        @nueva-alerta="manejarNuevaAlerta"
        @alertas-cargadas="actualizarEstadisticas"
      />
    </div>
  </div>
</template>

<script>
import AlertsComponent from '@/components/Alerts.vue'
import axios from 'axios'

export default {
  name: 'AlertasView',
  components: {
    AlertsComponent
  },
  props: {
    id: {
      type: [String, Number],
      default: null
    }
  },
  data() {
    return {
      salaInfo: null,
      loading: false,
      filtroTipo: null,
      filtroFecha: null,
      menuFecha: false,
      estadisticas: {
        total: 0,
        criticas: 0,
        preventivas: 0,
      },
      tiposAlerta: [
        { text: 'Crítico', value: 'critico' },
        { text: 'Preventivo', value: 'preventivo' },
      ]
    }
  },
  computed: {
    salaId() {
      return this.id || this.$route.params.id || null
    },
    pageTitle() {
      if (this.salaId) {
        return 'Alertas de Sala'
      }
      return 'Todas las Alertas'
    }
  },
  watch: {
    salaId: {
      immediate: true,
      handler() {
        if (this.salaId) {
          this.cargarInfoSala()
        }
      }
    }
  },
  methods: {
    async cargarInfoSala() {
      if (!this.salaId) return
      
      try {
        const response = await axios.get(`/api/salas/${this.salaId}`)
        this.salaInfo = response.data
      } catch (error) {
        console.error('Error al cargar información de la sala:', error)
      }
    },
    
    manejarNuevaAlerta(alerta) {
      // Mostrar notificación de nueva alerta
      if ('Notification' in window && Notification.permission === 'granted') {
        new Notification(`Nueva Alerta - ${alerta.tipo}`, {
          body: alerta.descripcion,
          icon: '/favicon.ico'
        })
      }
      
      // Actualizar estadísticas
      this.estadisticas.total++
      this.estadisticas[alerta.tipo]++
    },
    
    actualizarEstadisticas(alertas) {
      this.estadisticas = {
        total: alertas.length,
        criticas: alertas.filter(a => a.tipo === 'critico').length,
        preventivas: alertas.filter(a => a.tipo === 'preventivo').length,
      }
    },
    
    refrescarAlertas() {
      // El componente AlertsComponent maneja su propia actualización
      // Este método puede ser usado para forzar una actualización
      this.$refs.alertsComponent?.cargarAlertas()
    }
  }
}
</script>

<style scoped>
.alertas-view {
  min-height: 100vh;
  background-color: #f5f5f5;
}

.view-header {
  background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%);
  color: white;
  padding: 24px 0;
  margin-bottom: 24px;
}

.header-content {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 24px;
}

.header-text {
  margin-bottom: 24px;
}

.view-title {
  font-size: 2.5rem;
  font-weight: 300;
  margin: 0 0 8px 0;
}

.view-subtitle {
  font-size: 1.1rem;
  opacity: 0.9;
  margin: 0;
}

.header-actions {
  margin-top: 16px;
}

.stats-section {
  max-width: 1200px;
  margin: 0 auto 24px auto;
  padding: 0 24px;
}

.stat-card {
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.stat-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.stat-number {
  font-size: 2.5rem;
  font-weight: bold;
  line-height: 1;
  margin-bottom: 8px;
}

.stat-label {
  font-size: 0.875rem;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  opacity: 0.8;
}

.stat-total .stat-number {
  color: #1976d2;
}

.stat-criticas .stat-number {
  color: #f44336;
}

.stat-preventivas .stat-number {
  color: #ff9800;
}



.alerts-content {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 24px 24px 24px;
}

.gap-4 {
  gap: 16px;
}

/* Responsive */
@media (max-width: 768px) {
  .view-title {
    font-size: 2rem;
  }
  
  .header-content {
    padding: 0 16px;
  }
}
</style>