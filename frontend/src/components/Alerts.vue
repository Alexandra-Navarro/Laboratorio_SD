<template>
  <div class="alerts-component">
    <!-- Loading state -->
    <div v-if="loading" class="loading-container">
      <div class="loading-spinner"></div>
      <p>Cargando alertas...</p>
    </div>
    
    <!-- Error state -->
    <div v-else-if="error" class="error-container">
      <p class="error-message">{{ error }}</p>
      <button @click="cargarAlertas" class="retry-button">Reintentar</button>
    </div>
    
    <!-- Empty state -->
    <div v-else-if="alertas.length === 0" class="empty-container">
      <div class="empty-icon">🔔</div>
      <h3>No hay alertas</h3>
      <p>No se han registrado alertas para esta sala.</p>
      <!-- Debug info -->
      <div style="font-size: 12px; color: #999; margin-top: 20px;">
        Debug: {{ alertasOriginales.length }} alertas originales, {{ alertas.length }} alertas filtradas
      </div>
    </div>
    
    <!-- Alerts list -->
    <div v-else class="alerts-container">
      <div v-for="alerta in alertas" :key="alerta.id" class="alerta" :class="alerta.tipo">
        <div class="alerta-header">
          <span class="tipo" :class="`tipo-${alerta.tipo}`">{{ formatTipo(alerta.tipo) }}</span>
          <span class="fecha">{{ formatDate(alerta.fecha) }}</span>
        </div>
        <div class="alerta-body">
          <p class="descripcion">{{ alerta.descripcion }}</p>
          <div class="alerta-details">
            <span class="valor">Valor detectado: <strong>{{ alerta.valor_detectado }}</strong></span>
            <span class="umbral">Umbral: <strong>{{ formatUmbral(alerta.umbral) }}</strong></span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'AlertsComponent',
  props: {
    salaId: {
      type: [String, Number],
      default: null
    },
    filtroTipo: {
      type: String,
      default: null
    },
    filtroFecha: {
      type: String,
      default: null
    },
    limit: {
      type: Number,
      default: null
    },
    salaNombre: {
      type: String,
      default: ''
    },
    salas: {
      type: Array,
      default: () => []
    },
    salasCargadas: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      alertas: [],
      alertasOriginales: [],
      loading: false,
      error: null,
      pollingInterval: null,
      ultimaAlertaNotificada: null,
      salasInternas: [], // Para guardar las salas si las tenemos que cargar internamente
      componentId: Math.random().toString(36).substr(2, 9) // ID único para debug
    };
  },
  computed: {
    salasDisponibles() {
      // Priorizar las salas que vienen como prop
      if (this.salas && this.salas.length > 0) {
        return this.salas;
      }
      // Usar salas internas si las hemos cargado
      return this.salasInternas;
    },
    salasNombres() {
      const salasAUsar = this.salasDisponibles;
      
      if (!salasAUsar || salasAUsar.length === 0) {
        console.log('Alerts: no hay salas disponibles aún');
        return {};
      }

      const dic = {};
      salasAUsar.forEach(s => {
        if (s && s.id && s.nombre) {
          dic[s.id] = s.nombre;
        }
      });

      console.log('Alerts: salasNombres generado:', {
        salasAUsar,
        diccionario: dic
      });

      return dic;
    },
    alertasFiltradas() {
      console.log(`Alerts[${this.componentId}]: Calculando alertas filtradas`, {
        alertasOriginales: this.alertasOriginales.length,
        filtroTipo: this.filtroTipo,
        filtroFecha: this.filtroFecha,
        limit: this.limit
      });
      
      let filtradas = [...this.alertasOriginales];
      
      // Filtrar por tipo
      if (this.filtroTipo) {
        filtradas = filtradas.filter(alerta => alerta.tipo === this.filtroTipo);
        console.log(`Alerts[${this.componentId}]: Después de filtrar por tipo: ${filtradas.length}`);
      }
      
      // Filtrar por fecha
      if (this.filtroFecha) {
        const fechaFiltro = new Date(this.filtroFecha).toDateString();
        filtradas = filtradas.filter(alerta => {
          const fechaAlerta = new Date(alerta.fecha).toDateString();
          return fechaAlerta === fechaFiltro;
        });
        console.log(`Alerts[${this.componentId}]: Después de filtrar por fecha: ${filtradas.length}`);
      }
      
      // Aplicar límite si se especifica
      if (this.limit && this.limit > 0) {
        filtradas = filtradas.slice(0, this.limit);
        console.log(`Alerts[${this.componentId}]: Después de aplicar límite: ${filtradas.length}`);
      }
      
      console.log(`Alerts[${this.componentId}]: Alertas filtradas finales: ${filtradas.length}`);
      return filtradas;
    }
  },
  watch: {
    // Recargar alertas cuando cambie el ID de la sala
    salaId: {
      immediate: true,
      handler(newId, oldId) {
        if (newId !== oldId) {
          this.cargarAlertas();
        }
      }
    },
    // Actualizar alertas mostradas cuando cambien los filtros
    alertasFiltradas: {
      immediate: true,
      handler(nuevasAlertas) {
        console.log(`Alerts[${this.componentId}]: Watcher alertasFiltradas activado`, {
          nuevasAlertas: nuevasAlertas.length,
          alertasActuales: this.alertas.length
        });
        this.alertas = nuevasAlertas;
        this.$emit('alertas-cargadas', nuevasAlertas);
      }
    },
    // Nuevo watcher: cuando cambien las salas, verificar notificaciones pendientes
    salas: {
      deep: true,
      handler(nuevasSalas) {
        console.log('Salas cambiaron en Alerts:', nuevasSalas);
        // Las salas han cambiado, el diccionario se actualizará automáticamente
      }
    }
  },
  mounted() {
    console.log(`Alerts[${this.componentId}] mounted con props:`, {
      salaId: this.salaId,
      salaNombre: this.salaNombre,
      salas: this.salas,
      salasLength: this.salas ? this.salas.length : 'undefined',
      salasCargadas: this.salasCargadas
    });
    
    // Debug adicional: verificar estado inicial
    console.log(`Alerts[${this.componentId}] estado inicial:`, {
      alertas: this.alertas.length,
      alertasOriginales: this.alertasOriginales.length,
      loading: this.loading,
      error: this.error
    });
    
    this.iniciarPolling();
  },
  beforeUnmount() {
    this.detenerPolling();
  },
  methods: {
    // Método para cargar salas si no las recibimos como prop
    async cargarSalasInternas() {
      if (this.salas && this.salas.length > 0) {
        console.log('Alerts: usando salas del prop');
        return; // Ya tenemos salas del prop
      }

      try {
        console.log('Alerts: cargando salas internamente');
        const response = await axios.get('/api/salas');
        
        if (Array.isArray(response.data)) {
          this.salasInternas = response.data.map(sala => ({
            id: sala.id,
            nombre: sala.nombre
          }));
          console.log('Alerts: salas internas cargadas:', this.salasInternas);
        }
      } catch (error) {
        console.error('Alerts: error al cargar salas internas:', error);
      }
    },

    iniciarPolling() {
      this.cargarSalasInternas(); // Cargar salas si es necesario
      this.cargarAlertas();
      this.pollingInterval = setInterval(this.cargarAlertas, 3000); // 3 segundos = 3000 ms
    },
    
    detenerPolling() {
      if (this.pollingInterval) {
        clearInterval(this.pollingInterval);
        this.pollingInterval = null;
      }
    },

    // Método para obtener el nombre correcto de la sala
    obtenerNombreSala(alerta) {
      // Debug: logs para ver qué datos tenemos
      console.log('Debug obtenerNombreSala:', {
        salaNombre: this.salaNombre,
        salasDisponibles: this.salasDisponibles.length,
        salasInternas: this.salasInternas.length,
        salasProp: this.salas ? this.salas.length : 0,
        salasNombres: this.salasNombres,
        salasNombresKeys: Object.keys(this.salasNombres || {}),
        alertaSalaId: alerta.sala_id,
        nombreEncontrado: this.salasNombres[alerta.sala_id]
      });
      
      // Priorizar el nombre directo si estamos en una sala específica
      if (this.salaNombre && this.salaNombre.trim() !== '') {
        console.log('Usando salaNombre:', this.salaNombre);
        return this.salaNombre;
      }
      
      // Usar el diccionario de salas si está disponible y no está vacío
      if (this.salasNombres && 
          Object.keys(this.salasNombres).length > 0 && 
          this.salasNombres[alerta.sala_id]) {
        const nombreSala = this.salasNombres[alerta.sala_id];
        console.log('Usando nombre del diccionario:', nombreSala);
        return nombreSala;
      }
      
      // Fallback al ID de la sala
      console.log('Usando fallback para sala:', alerta.sala_id);
      return `Sala ${alerta.sala_id}`;
    },
    
    async cargarAlertas() {
      try {
        this.loading = true;
        this.error = null;
        
        let url;
        if (this.salaId) {
          url = `/api/alertas/sala/${this.salaId}`;
        } else {
          url = '/api/alertas';
        }
        
        console.log(`Alerts[${this.componentId}]: Cargando alertas desde ${url}`);
        const response = await axios.get(url);
        const nuevasAlertas = response.data;
        console.log(`Alerts[${this.componentId}]: Recibidas ${nuevasAlertas.length} alertas`);
        
        // Ordenar por fecha descendente (más nueva primero)
        nuevasAlertas.sort((a, b) => new Date(b.fecha) - new Date(a.fecha));

        // Verificar nuevas alertas para notificaciones
        if (nuevasAlertas.length > 0) {
          const alertaMasReciente = nuevasAlertas[0];

          if (this.ultimaAlertaNotificada === null) {
            this.ultimaAlertaNotificada = alertaMasReciente.id;
            console.log(`Alerts[${this.componentId}]: Primera carga, estableciendo última alerta: ${alertaMasReciente.id}`);
          } else if (alertaMasReciente.id !== this.ultimaAlertaNotificada) {
            const nuevas = [];
            for (const alerta of nuevasAlertas) {
              if (alerta.id === this.ultimaAlertaNotificada) break;
              nuevas.push(alerta);
            }
            
            if (nuevas.length > 0) {
              console.log(`Alerts[${this.componentId}]: ${nuevas.length} nuevas alertas detectadas`);
              
              // Emitir evento para nuevas alertas y mostrar notificación
              nuevas.reverse().forEach(alerta => {
                console.log(`Alerts[${this.componentId}]: procesando alerta ${alerta.id}`);
                this.$emit('nueva-alerta', alerta);
                
                // Obtener el nombre de la sala de forma consistente
                const nombreSala = this.obtenerNombreSala(alerta);
                
                const titulo = `${this.formatTipo(alerta.tipo)} en ${nombreSala}`;
                
                // Solo mostrar notificación toast (evitar duplicados)
                if (this.$root && typeof this.$root.showToastNotification === 'function') {
                  const tipoNotificacion = alerta.tipo === 'critico' ? 'critical' : 'alert';
                  const mensajeToast = `${alerta.descripcion}. Valor detectado: ${alerta.valor_detectado}`;
                  console.log(`Alerts[${this.componentId}]: mostrando notificación para alerta ${alerta.id}`);
                  this.$root.showToastNotification(titulo, mensajeToast, tipoNotificacion, alerta.id);
                }
              });
              
              this.ultimaAlertaNotificada = alertaMasReciente.id;
            }
          }
        }

        // IMPORTANTE: Siempre actualizar las alertas mostradas
        this.alertasOriginales = nuevasAlertas;
        console.log(`Alerts[${this.componentId}]: Alertas originales actualizadas: ${this.alertasOriginales.length} alertas`);
        this.$emit('alertas-todas-cargadas', this.alertasOriginales);
      } catch (error) {
        console.error(`Alerts[${this.componentId}]: Error al cargar alertas:`, error);
        this.error = 'Error al cargar las alertas. Verifique su conexión.';
      } finally {
        this.loading = false;
      }
    },
    
    formatDate(date) {
      return new Date(date).toLocaleString('es-CL', {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit',
        hour: '2-digit',
        minute: '2-digit'
      });
    },
    
    formatTipo(tipo) {
      const tipos = {
        //'informativo': 'Informativo',
        'preventivo': 'Preventivo',
        'critico': 'Crítico'
      };
      return tipos[tipo] || tipo;
    },
    
    formatUmbral(umbral) {
      const umbrales = {
        'bajo': 'Bajo',
        'alto': 'Alto',
        'variable': 'Variable'
      };
      return umbrales[umbral] || umbral;
    }
  }
};
</script>

<style scoped>
.alerts-component {
  width: 100%;
}

.loading-container, .error-container, .empty-container {
  text-align: center;
  padding: 40px 20px;
}

.loading-spinner {
  width: 40px;
  height: 40px;
  border: 4px solid #f3f3f3;
  border-top: 4px solid #2196f3;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin: 0 auto 16px;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.error-container {
  color: #d32f2f;
}

.error-message {
  margin-bottom: 16px;
}

.retry-button {
  background: #2196f3;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
}

.retry-button:hover {
  background: #00b4db;
}

.empty-container {
  color: #666;
}

.empty-icon {
  font-size: 48px;
  margin-bottom: 16px;
}

.alerts-container {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.alerta {
  padding: 20px;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  border-left: 6px solid;
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.alerta:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

/*.alerta.informativo {
  background-color: #e3f2fd;
  border-left-color: #2196f3;
}*/

.alerta.preventivo {
  background-color: #fff3e0;
  border-left-color: #ff9800;
}

.alerta.critico {
  background-color: #ffebee;
  border-left-color: #f44336;
}

.alerta-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.tipo {
  padding: 4px 12px;
  border-radius: 16px;
  font-weight: 600;
  font-size: 0.85em;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.tipo-informativo {
  background: #2196f3;
  color: white;
}

.tipo-preventivo {
  background: #ff9800;
  color: white;
}

.tipo-critico {
  background: #f44336;
  color: white;
}

.fecha {
  color: #666;
  font-size: 0.9em;
  font-weight: 500;
}

.alerta-body {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.descripcion {
  margin: 0;
  font-weight: 500;
  font-size: 1.1em;
  color: #333;
  line-height: 1.4;
}

.alerta-details {
  display: flex;
  gap: 24px;
  flex-wrap: wrap;
}

.valor, .umbral {
  margin: 0;
  font-size: 0.95em;
  color: #555;
}

.valor strong, .umbral strong {
  color: #333;
}

/* Responsive design */
@media (max-width: 768px) {
  .alerta {
    padding: 16px;
  }
  
  .alerta-details {
    flex-direction: column;
    gap: 8px;
  }
}

.alert-info {
    background: #00b4db;
    color: white;
}
</style>