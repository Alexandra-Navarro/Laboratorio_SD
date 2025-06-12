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
    }
  },
  data() {
    return {
      alertas: [],
      alertasOriginales: [],
      loading: false,
      error: null,
      pollingInterval: null,
      ultimaAlertaNotificada: null
    };
  },
  computed: {
    alertasFiltradas() {
      let filtradas = [...this.alertasOriginales];
      
      // Filtrar por tipo
      if (this.filtroTipo) {
        filtradas = filtradas.filter(alerta => alerta.tipo === this.filtroTipo);
      }
      
      // Filtrar por fecha
      if (this.filtroFecha) {
        const fechaFiltro = new Date(this.filtroFecha).toDateString();
        filtradas = filtradas.filter(alerta => {
          const fechaAlerta = new Date(alerta.fecha).toDateString();
          return fechaAlerta === fechaFiltro;
        });
      }
      
      // Aplicar límite si se especifica
      if (this.limit && this.limit > 0) {
        filtradas = filtradas.slice(0, this.limit);
      }
      
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
        this.alertas = nuevasAlertas;
        this.$emit('alertas-cargadas', nuevasAlertas);
      }
    }
  },
  mounted() {
    this.iniciarPolling();
  },
  beforeUnmount() {
    this.detenerPolling();
  },
  methods: {
    iniciarPolling() {
      this.cargarAlertas();
      this.pollingInterval = setInterval(this.cargarAlertas, 5000);
    },
    
    detenerPolling() {
      if (this.pollingInterval) {
        clearInterval(this.pollingInterval);
        this.pollingInterval = null;
      }
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
        
        const response = await axios.get(url);
        const nuevasAlertas = response.data;
        console.log('Alertas cargadas:', nuevasAlertas);

        // Verificar nuevas alertas para notificaciones
        if (nuevasAlertas.length > 0) {
          const alertaMasReciente = nuevasAlertas[0];
          console.log('Última alerta notificada:', this.ultimaAlertaNotificada);
          console.log('Alerta más reciente:', alertaMasReciente.id);

          if (this.ultimaAlertaNotificada === null) {
            this.ultimaAlertaNotificada = alertaMasReciente.id;
            console.log('Primera carga de alertas, no se muestran notificaciones');
          } else if (alertaMasReciente.id !== this.ultimaAlertaNotificada) {
            const nuevas = [];
            for (const alerta of nuevasAlertas) {
              if (alerta.id === this.ultimaAlertaNotificada) break;
              nuevas.push(alerta);
            }
            
            console.log('Nuevas alertas detectadas:', nuevas);
            
            // Emitir evento para nuevas alertas y mostrar notificación
            nuevas.reverse().forEach(alerta => {
              console.log('Procesando nueva alerta:', alerta);
              this.$emit('nueva-alerta', alerta);
              
              // Asegurarse de que el método showNotification esté disponible
              if (this.$root && typeof this.$root.showNotification === 'function') {
                this.$root.showNotification(
                  `Nueva Alerta - ${this.formatTipo(alerta.tipo)}`,
                  `${alerta.descripcion}\nValor detectado: ${alerta.valor_detectado}`
                );
              } else {
                console.error('Método showNotification no disponible en $root');
              }
            });
            
            this.ultimaAlertaNotificada = alertaMasReciente.id;
          }
        }

        this.alertasOriginales = nuevasAlertas;
      } catch (error) {
        console.error('Error al cargar alertas:', error);
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
        'informativo': 'Informativo',
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
  background: #1976d2;
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

.alerta.informativo {
  background-color: #e3f2fd;
  border-left-color: #2196f3;
}

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
</style>