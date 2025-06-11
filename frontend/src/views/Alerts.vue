<template>
  <div class="view-alerts">
    <h1>Alertas</h1>
    <div class="alerts-container">
      <div v-for="alerta in alertas" :key="alerta.id" class="alerta" :class="alerta.tipo">
        <div class="alerta-header">
          <span class="tipo">{{ alerta.tipo }}</span>
          <span class="fecha">{{ formatDate(alerta.fecha) }}</span>
        </div>
        <div class="alerta-body">
          <p class="descripcion">{{ alerta.descripcion }}</p>
          <p class="valor">Valor detectado: {{ alerta.valor_detectado }}</p>
          <p class="umbral">Umbral: {{ alerta.umbral }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Alerts',
  data() {
    return {
      alertas: [],
      pollingInterval: null,
      ultimaAlertaNotificada: null
    };
  },
  created() {
    this.cargarAlertas();
    this.pollingInterval = setInterval(this.cargarAlertas, 5000);
  },
  beforeUnmount() {
    if (this.pollingInterval) {
      clearInterval(this.pollingInterval);
    }
  },
  methods: {
    async cargarAlertas() {
      try {
        const response = await axios.get('/api/alertas');
        const nuevasAlertas = response.data;

        if (nuevasAlertas.length > 0) {
          const alertaMasReciente = nuevasAlertas[0];

          if (this.ultimaAlertaNotificada === null) {
            this.ultimaAlertaNotificada = alertaMasReciente.id;
          } else if (alertaMasReciente.id !== this.ultimaAlertaNotificada) {
            const nuevas = [];
            for (const alerta of nuevasAlertas) {
              if (alerta.id === this.ultimaAlertaNotificada) break;
              nuevas.push(alerta);
            }
            nuevas.reverse().forEach(alerta => {
              alert(`Nueva Alerta: ${alerta.descripcion}`);
            });
            this.ultimaAlertaNotificada = alertaMasReciente.id;
          }
        }

        this.alertas = nuevasAlertas;
      } catch (error) {
        console.error('Error al cargar alertas:', error);
      }
    },
    formatDate(date) {
      return new Date(date).toLocaleString();
    }
  }
};
</script>

<style scoped>
.view-alerts {
  padding: 16px;
}

.alerts-container {
  display: flex;
  flex-direction: column;
  gap: 16px;
  max-width: 800px;
  margin: 0 auto;
}

.alerta {
  padding: 16px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.alerta.preventivo {
  background-color: #fff3cd;
  border-left: 4px solid #ffc107;
}

.alerta.critico {
  background-color: #f8d7da;
  border-left: 4px solid #dc3545;
}

.alerta-header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 8px;
}

.tipo {
  font-weight: bold;
  text-transform: capitalize;
}

.fecha {
  color: #666;
  font-size: 0.9em;
}

.alerta-body {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.descripcion {
  margin: 0;
  font-weight: 500;
}

.valor, .umbral {
  margin: 0;
  font-size: 0.9em;
  color: #666;
}
</style>
