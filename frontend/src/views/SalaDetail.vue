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
            <v-chip small outlined>
              <v-icon left small>mdi-map-marker</v-icon>
              {{ escuelaInfo?.direccion }}
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
    <!--<div class="tabs-section">
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
            <v-tab>
              <v-icon left>mdi-tune-variant</v-icon>
              Umbrales
            </v-tab>
          </v-tabs>
        </v-container>
      </div>-->

    <!-- Contenido de las tabs -->
    <v-container class="sala-content">
      <v-tabs-items v-model="tabActivo">
        <!-- Tab Dashboard -->
        <v-tab-item>
          <div class="dashboard-tab">
            <!-- Métricas rápidas -->
            <v-row class="mb-6" justify="center" align="center" style="max-width:1100px;margin:0 auto;">
              <v-col cols="12" sm="6" md="2">
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
              <v-col cols="12" sm="6" md="2">
                <v-card class="metric-card">
                  <v-card-text class="text-center">
                    <div class="metric-icon">
                      <v-icon size="40" color="success">mdi-molecule-co2</v-icon>
                    </div>
                    <div class="metric-value" style="font-size:1.5rem;">{{ metricas.co2 }} ppm</div>
                    <div class="metric-label">CO2</div>
                  </v-card-text>
                </v-card>
              </v-col>
              <v-col cols="12" sm="6" md="2">
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
              <v-col cols="12" sm="6" md="2">
                <v-card class="metric-card">
                  <v-card-text class="text-center">
                    <div class="metric-icon">
                      <v-icon size="40" color="warning">mdi-volume-high</v-icon>
                    </div>
                    <div class="metric-value">{{ metricas.ruido }} dB</div>
                    <div class="metric-label">Ruido</div>
                  </v-card-text>
                </v-card>
              </v-col>
              <v-col cols="12" sm="6" md="2">
                <v-card class="metric-card">
                  <v-card-text class="text-center">
                    <div class="metric-icon">
                      <v-icon size="40" color="amber">mdi-white-balance-sunny</v-icon>
                    </div>
                    <div class="metric-value">{{ metricas.iluminacion }} lux</div>
                    <div class="metric-label">Iluminación</div>
                  </v-card-text>
                </v-card>
              </v-col>
            </v-row>

            <!-- Gráficos y datos adicionales 
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
              </v-row>-->
          </div>
        </v-tab-item>

        <!-- Tab Alertas -->
        <v-tab-item>
          <v-row class="alertas-umbrales-row" align="start" justify="center">
            <v-col cols="12" md="7">
              <div class="alertas-tab">
                <v-card>
                  <v-card-title>
                    <v-icon left>mdi-bell-alert</v-icon>
                    Alertas
                  </v-card-title>
                  <v-card-text class="alertas-container">
                    <AlertsComponent :sala-id="salaId" @nueva-alerta="manejarNuevaAlerta"
                      @alertas-cargadas="actualizarConteoAlertas" />
                  </v-card-text>
                </v-card>
              </div>
            </v-col>
            <v-col cols="12" md="5">
              <div class="umbrales-tab">
                <v-card>
                  <v-card-title>
                    <v-icon left>mdi-tune-variant</v-icon>
                    Gestión de Umbrales Personalizados
                  </v-card-title>
                  <v-card-text>
                    <v-alert type="info" outlined class="mb-4">
                      Aquí puedes configurar umbrales personalizados para cada variable ambiental de esta sala. Si no
                      existe un umbral personalizado, se mostrará como "No definido".
                    </v-alert>
                    <div class="tabla-umbrales-wrapper">
                      <v-simple-table class="tabla-umbrales">
                        <thead>
                          <tr>
                            <th>Variable</th>
                            <th>Unidad</th>
                            <th>Umbral Bajo</th>
                            <th>Umbral Alto</th>
                            <th>Acciones</th>
                          </tr>
                        </thead>
                        <tbody>
                          <tr v-for="variable in variablesAmbientales" :key="variable.id"
                            :class="{ 'fila-edicion': editandoId === variable.id }">
                            <td>{{ variable.nombre }}</td>
                            <td>{{ variable.unidad_medida }}</td>
                            <td>
                              <template v-if="editandoId === variable.id">
                                <v-text-field v-model.number="formUmbral.umbral_bajo" dense hide-details type="number"
                                  step="any" style="max-width:100px" class="campo-edicion" />
                              </template>
                              <template v-else>
                                {{ obtenerUmbral(variable.id, 'bajo') }}
                              </template>
                            </td>
                            <td>
                              <template v-if="editandoId === variable.id">
                                <v-text-field v-model.number="formUmbral.umbral_alto" dense hide-details type="number"
                                  step="any" style="max-width:100px" class="campo-edicion" />
                              </template>
                              <template v-else>
                                {{ obtenerUmbral(variable.id, 'alto') }}
                              </template>
                            </td>
                            <td>
                              <template v-if="editandoId === variable.id">
                                <v-btn color="success" icon small
                                  @click="guardarUmbral(variable)"><v-icon>mdi-content-save</v-icon></v-btn>
                                <v-btn color="grey" icon small
                                  @click="cancelarEdicion"><v-icon>mdi-close</v-icon></v-btn>
                              </template>
                              <template v-else>
                                <v-btn color="primary" icon small
                                  @click="editarUmbral(variable)"><v-icon>mdi-pencil</v-icon></v-btn>
                              </template>
                            </td>
                          </tr>
                        </tbody>
                      </v-simple-table>
                    </div>
                  </v-card-text>
                </v-card>
              </div>
            </v-col>
          </v-row>
        </v-tab-item>

        <!-- Tab Sensores 
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
          </v-tab-item>-->

        <!-- Tab Historial 
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
          </v-tab-item>-->

        <!-- Tab Umbrales 
        <v-tab-item>
          <div class="umbrales-tab">
            <v-card>
              <v-card-title>
                <v-icon left>mdi-tune-variant</v-icon>
                Gestión de Umbrales Personalizados
              </v-card-title>
              <v-card-text>
                <v-alert type="info" outlined class="mb-4">
                  Aquí puedes configurar umbrales personalizados para cada variable ambiental de esta sala. Si no existe
                  un
                  umbral personalizado, se mostrará como "No definido".
                </v-alert>
                <div class="tabla-umbrales-wrapper">
                  <v-simple-table class="tabla-umbrales">
                    <thead>
                      <tr>
                        <th>Variable</th>
                        <th>Unidad</th>
                        <th>Umbral Bajo</th>
                        <th>Umbral Alto</th>
                        <th>Acciones</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr v-for="variable in variablesAmbientales" :key="variable.id"
                        :class="{ 'fila-edicion': editandoId === variable.id }">
                        <td>{{ variable.nombre }}</td>
                        <td>{{ variable.unidad_medida }}</td>
                        <td>
                          <template v-if="editandoId === variable.id">
                            <v-text-field v-model.number="formUmbral.umbral_bajo" dense hide-details type="number"
                              step="any" style="max-width:100px" class="campo-edicion" />
                          </template>
                          <template v-else>
                            {{ obtenerUmbral(variable.id, 'bajo') }}
                          </template>
                        </td>
                        <td>
                          <template v-if="editandoId === variable.id">
                            <v-text-field v-model.number="formUmbral.umbral_alto" dense hide-details type="number"
                              step="any" style="max-width:100px" class="campo-edicion" />
                          </template>
                          <template v-else>
                            {{ obtenerUmbral(variable.id, 'alto') }}
                          </template>
                        </td>
                        <td>
                          <template v-if="editandoId === variable.id">
                            <v-btn color="success" icon small
                              @click="guardarUmbral(variable)"><v-icon>mdi-content-save</v-icon></v-btn>
                            <v-btn color="grey" icon small @click="cancelarEdicion"><v-icon>mdi-close</v-icon></v-btn>
                          </template>
                          <template v-else>
                            <v-btn color="primary" icon small
                              @click="editarUmbral(variable)"><v-icon>mdi-pencil</v-icon></v-btn>
                          </template>
                        </td>
                      </tr>
                    </tbody>
                  </v-simple-table>
                </div>
              </v-card-text>
            </v-card>
          </div>
        </v-tab-item>-->
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
        co2: '--',
        humedad: '--',
        ruido: '--',
        iluminacion: '--'
      },
      variablesAmbientales: [],
      umbralesPersonalizados: [],
      editandoId: null,
      formUmbral: {
        umbral_bajo: '',
        umbral_alto: ''
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
        this.cargarMetricas()
        this.cargarVariablesYUmbrales()
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

        this.cargarMetricas()

      } catch (error) {
        console.error('Error al cargar datos de la sala:', error)
      } finally {
        this.loading = false
      }
    },

    async cargarMetricas() {
      try {
        // 1. Obtener sensores de la sala
        const sensoresRes = await axios.get(`/api/sensores/sala/${this.salaId}`);
        let sensores = sensoresRes.data;
        if (!Array.isArray(sensores)) {
          if (Array.isArray(sensores.sensores)) {
            sensores = sensores.sensores;
          } else {
            console.warn('La respuesta de sensores no es un array:', sensores);
            sensores = [];
          }
        }
        const sensorToVariable = {};
        sensores.forEach(sensor => {
          sensorToVariable[sensor.id] = sensor.variable_id;
        });

        // 2. Obtener mediciones de la sala
        const medicionesRes = await axios.get(`/api/mediciones/sala/${this.salaId}`);
        let mediciones = medicionesRes.data;
        if (!Array.isArray(mediciones)) {
          if (Array.isArray(mediciones.mediciones)) {
            mediciones = mediciones.mediciones;
          } else {
            console.warn('La respuesta de mediciones no es un array:', mediciones);
            mediciones = [];
          }
        }

        // 3. Obtener la medición más reciente para cada variable
        const latestByVariable = {};
        mediciones.forEach(med => {
          const variableId = sensorToVariable[med.sensor_id];
          if (!variableId) return;
          if (!latestByVariable[variableId] || new Date(med.fecha) > new Date(latestByVariable[variableId].fecha)) {
            latestByVariable[variableId] = med;
          }
        });

        // 4. Mapear variable_id a nombre de variable
        this.metricas = {
          temperatura: latestByVariable[1]?.valor?.toFixed(1) ?? '--',
          humedad: latestByVariable[2]?.valor?.toFixed(1) ?? '--',
          co2: latestByVariable[3]?.valor?.toFixed(0) ?? '--',
          ruido: latestByVariable[4]?.valor?.toFixed(1) ?? '--',
          iluminacion: latestByVariable[5]?.valor?.toFixed(0) ?? '--'
        };
      } catch (error) {
        console.error('Error al cargar métricas:', error);
      }
    },

    manejarNuevaAlerta(alerta) {
      if (alerta.tipo === 'critico') {
        this.alertasCriticas++
      }

      this.$root.showNotification(
        `Nueva Alerta en ${this.salaInfo.nombre} - ${alerta.tipo}`,
        `${alerta.descripcion}\nValor detectado: ${alerta.valor_detectado}`
      )
    },

    actualizarConteoAlertas(alertas) {
      this.alertasCriticas = alertas.filter(a => a.tipo === 'critico').length
    },

    formatDate(date) {
      return new Date(date).toLocaleString('es-CL')
    },

    async cargarVariablesYUmbrales() {
      try {
        // Obtener variables ambientales
        const variablesRes = await axios.get('/api/variables')
        this.variablesAmbientales = variablesRes.data
        // Obtener umbrales personalizados de la sala
        const umbralesRes = await axios.get(`/api/umbrales?sala_id=${this.salaId}`)
        this.umbralesPersonalizados = umbralesRes.data
      } catch (error) {
        console.error('Error al cargar variables o umbrales:', error)
      }
    },
    obtenerUmbral(variableId, tipo) {
      const umbral = this.umbralesPersonalizados.find(u => u.variable_id === variableId)
      if (umbral) {
        return tipo === 'bajo' ? umbral.umbral_bajo : umbral.umbral_alto
      }
      return 'No definido'
    },
    tieneUmbralPersonalizado(variableId) {
      return this.umbralesPersonalizados.some(u => u.variable_id === variableId)
    },
    editarUmbral(variable) {
      const umbral = this.umbralesPersonalizados.find(u => u.variable_id === variable.id)
      if (umbral) {
        this.formUmbral = { ...umbral }
      } else {
        this.formUmbral = { variable_id: variable.id, sala_id: this.salaId, umbral_bajo: '', umbral_alto: '' }
      }
      this.editandoId = variable.id
    },
    cancelarEdicion() {
      this.editandoId = null
      this.formUmbral = { umbral_bajo: '', umbral_alto: '' }
    },
    async guardarUmbral(variable) {
      try {
        // Validar que el umbral bajo no sea mayor que el umbral alto
        if (this.formUmbral.umbral_bajo >= this.formUmbral.umbral_alto) {
          this.$root.showNotification('Error', 'El umbral bajo no puede ser mayor que el umbral alto.', 'error');
          return;
        }

        if (this.tieneUmbralPersonalizado(variable.id)) {
          // Editar umbral existente
          const umbral = this.umbralesPersonalizados.find(u => u.variable_id === variable.id);
          await axios.put(`/api/umbrales/${umbral.id}`, {
            ...umbral,
            umbral_bajo: this.formUmbral.umbral_bajo,
            umbral_alto: this.formUmbral.umbral_alto
          });
        } else {
          // Crear nuevo umbral personalizado
          await axios.post('/api/umbrales', {
            sala_id: this.salaId,
            variable_id: variable.id,
            umbral_bajo: this.formUmbral.umbral_bajo,
            umbral_alto: this.formUmbral.umbral_alto
          });
        }
        this.editandoId = null;
        this.formUmbral = { umbral_bajo: '', umbral_alto: '' };
        await this.cargarVariablesYUmbrales();
        this.$root.showNotification('Umbral guardado', 'El umbral se guardó correctamente.');
      } catch (error) {
        console.error('Error al guardar umbral:', error);
        this.$root.showNotification('Error', 'No se pudo guardar el umbral.', 'error');
      }
    },
    async eliminarUmbral(variable) {
      try {
        const umbral = this.umbralesPersonalizados.find(u => u.variable_id === variable.id)
        if (umbral) {
          await axios.delete(`/api/umbrales/${umbral.id}`)
          await this.cargarVariablesYUmbrales()
          this.$root.showNotification('Umbral eliminado', 'El umbral personalizado fue eliminado.')
        }
      } catch (error) {
        console.error('Error al eliminar umbral:', error)
        this.$root.showNotification('Error', 'No se pudo eliminar el umbral.', 'error')
      }
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
  background: linear-gradient(135deg, #00b4db 0%, #0083b0 100%);
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

.tabla-umbrales-wrapper {
  background: #fff;
  border-radius: 16px;
  box-shadow: 0 4px 16px rgba(0,0,0,0.06);
  padding: 8px 0 8px 0;
  margin-bottom: 16px;
  overflow-x: auto;
}

.tabla-umbrales {
  width: 100%;
  border-radius: 12px;
  overflow: hidden;
  border-collapse: separate;
  border-spacing: 0;
}

.tabla-umbrales th,
.tabla-umbrales td {
  text-align: center;
  vertical-align: middle;
  padding: 6px 10px;
  border-bottom: 1.5px solid #f0f0f0;
}

.tabla-umbrales th:not(:last-child),
.tabla-umbrales td:not(:last-child) {
  border-right: 1px solid #e0e4ea;
}

.tabla-umbrales th {
  background: #f5f7fa;
  color: #00b4db;
  font-weight: 600;
  font-size: 1rem;
  border-bottom: 2px solid #e0e4ea;
}

.tabla-umbrales td {
  font-size: 0.98rem;
}

.tabla-umbrales td:last-child {
  text-align: center;
  vertical-align: middle;
}

.tabla-umbrales .v-btn {
  margin: 2px 0;
}

.fila-edicion {
  background: #e3f2fd !important;
  transition: background 0.2s;
}

.campo-edicion {
  margin: 0 auto;
  min-width: 70px;
  max-width: 100px;
}

.alertas-container {
  max-height: 600px;
  overflow-y: auto;
  padding: 0;
}

.alertas-container::-webkit-scrollbar {
  width: 8px;
}

.alertas-container::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 4px;
}

.alertas-container::-webkit-scrollbar-thumb {
  background: #888;
  border-radius: 4px;
}

.alertas-container::-webkit-scrollbar-thumb:hover {
  background: #555;
}
</style>