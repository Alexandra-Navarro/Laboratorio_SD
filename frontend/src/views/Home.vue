<template>
    <div class="home">
        <!-- Hero Section -->
        <div class="hero-section">
            <v-container>
                <v-row align="center" justify="center">
                    <v-col cols="12" md="8" class="text-center">
                        <h1 class="hero-title">
                            Sistema de Monitoreo Ambiental
                        </h1>
                        <p class="hero-subtitle">
                            Monitoreo en tiempo real de las condiciones ambientales en salas educativas
                        </p>
                        <div class="hero-actions">
                            <v-btn color="primary" large to="/dashboard" class="mr-4">
                                <v-icon left>mdi-view-dashboard</v-icon>
                                Ver Dashboard
                            </v-btn>
                            <v-btn outlined large to="/alertas">
                                <v-icon left>mdi-bell-ring</v-icon>
                                Ver Alertas
                            </v-btn>
                        </div>
                    </v-col>
                </v-row>
            </v-container>
            <!-- Features Section -->
            <div class="features-section">
                <v-container>
                    <v-row justify="center" class="mb-8">
                        <v-col cols="12" md="8" class="text-center">
                            <h2 class="section-title">Funcionalidades Principales</h2>
                            <p class="section-subtitle">
                                Nuestro sistema ofrece monitoreo completo y alertas en tiempo real
                            </p>
                        </v-col>
                    </v-row>

                    <v-row>
                        <v-col cols="12" md="4" class="mb-4">
                            <v-card class="feature-card h-100">
                                <v-card-text class="text-center pa-6">
                                    <div class="feature-icon mb-4">
                                        <v-icon size="64" color="primary">mdi-thermometer</v-icon>
                                    </div>
                                    <h3 class="feature-title mb-3">Monitoreo Ambiental</h3>
                                    <p class="feature-description">
                                        Seguimiento continuo de temperatura, humedad, calidad del aire y niveles de
                                        ruido
                                    </p>
                                </v-card-text>
                            </v-card>
                        </v-col>

                        <v-col cols="12" md="4" class="mb-4">
                            <v-card class="feature-card h-100">
                                <v-card-text class="text-center pa-6">
                                    <div class="feature-icon mb-4">
                                        <v-icon size="64" color="warning">mdi-bell-alert</v-icon>
                                    </div>
                                    <h3 class="feature-title mb-3">Alertas Inteligentes</h3>
                                    <p class="feature-description">
                                        Sistema de alertas automático que notifica cuando las condiciones salen de los
                                        parámetros normales
                                    </p>
                                </v-card-text>
                            </v-card>
                        </v-col>

                        <v-col cols="12" md="4" class="mb-4">
                            <v-card class="feature-card h-100">
                                <v-card-text class="text-center pa-6">
                                    <div class="feature-icon mb-4">
                                        <v-icon size="64" color="success">mdi-chart-line</v-icon>
                                    </div>
                                    <h3 class="feature-title mb-3">Análisis de Datos</h3>
                                    <p class="feature-description">
                                        Visualización de tendencias y patrones para la toma de decisiones informadas
                                    </p>
                                </v-card-text>
                            </v-card>
                        </v-col>
                    </v-row>
                </v-container>
            </div>

            <!-- Quick Access Section -->
            <div class="quick-access-section">
                <v-container>
                    <v-row justify="center" class="mb-6">
                        <v-col cols="12" md="8" class="text-center">
                            <h2 class="section-title">Acceso Rápido a Salas</h2>
                            <p class="section-subtitle">
                                Selecciona una sala para ver su estado actual
                            </p>
                        </v-col>
                    </v-row>

                    <v-row v-if="loading">
                        <v-col v-for="i in 6" :key="i" cols="12" sm="6" md="4">
                            <v-skeleton-loader type="card"></v-skeleton-loader>
                        </v-col>
                    </v-row>

                    <v-row v-else>
                        <v-col v-for="sala in salasDestacadas" :key="sala.id" cols="12" sm="6" md="4" class="mb-4">
                            <v-card class="sala-card" hover :to="sala.path">
                                <v-card-text class="pa-4">
                                    <div class="d-flex align-center mb-3">
                                        <v-icon :color="getSalaStatusColor(sala)" size="24" class="mr-2">
                                            {{ sala.icon }}
                                        </v-icon>
                                        <div>
                                            <h4 class="sala-name">{{ sala.nombre }}</h4>
                                            <p class="sala-status text-caption mb-0">
                                                {{ getSalaStatus(sala) }}
                                            </p>
                                        </div>
                                        <v-spacer></v-spacer>
                                        <v-chip v-if="sala.alertasCriticas > 0" color="error" text-color="white" small>
                                            {{ sala.alertasCriticas }}
                                        </v-chip>
                                    </div>

                                    <div class="sala-metrics">
                                        <div class="metric-item">
                                            <v-icon small class="mr-1">mdi-thermometer</v-icon>
                                            <span class="text-caption">{{ getRandomMetric(18, 25) }}°C</span>
                                        </div>
                                        <div class="metric-item">
                                            <v-icon small class="mr-1">mdi-water-percent</v-icon>
                                            <span class="text-caption">{{ getRandomMetric(40, 60) }}%</span>
                                        </div>
                                    </div>
                                </v-card-text>
                            </v-card>
                        </v-col>
                    </v-row>

                    <v-row justify="center" class="mt-4">
                        <v-col cols="auto">
                            <v-btn color="primary" outlined to="/dashboard" large>
                                Ver Todas las Salas
                                <v-icon right>mdi-arrow-right</v-icon>
                            </v-btn>
                        </v-col>
                    </v-row>
                </v-container>
            </div>

            <!-- Recent Alerts Section -->
            <div class="recent-alerts-section">
                <v-container>
                    <v-row justify="center" class="mb-6">
                        <v-col cols="12" md="8" class="text-center">
                            <h2 class="section-title">Alertas Recientes</h2>
                            <p class="section-subtitle">
                                Últimas alertas registradas en el sistema
                            </p>
                        </v-col>
                    </v-row>

                    <v-row justify="center">
                        <v-col cols="12" md="10">
                            <AlertsComponent :limit="5" @alertas-cargadas="actualizarEstadisticasHome" />
                            <div class="text-center mt-4">
                                <v-btn color="primary" outlined to="/alertas">
                                    Ver Todas las Alertas
                                    <v-icon right>mdi-arrow-right</v-icon>
                                </v-btn>
                            </div>
                        </v-col>
                    </v-row>
                </v-container>
            </div>
        </div>
    </div>
</template>

<script>
import AlertsComponent from '@/components/Alerts.vue'
import axios from 'axios'

export default {
    name: 'Home',
    components: {
        AlertsComponent
    },
    data() {
        return {
            loading: false,
            estadisticas: {
                totalSalas: 0,
                sensoresActivos: 0,
                alertasHoy: 0,
                escuelas: 0
            },
            salasDestacadas: []
        }
    },
    async created() {
        await this.cargarDatos()
    },
    methods: {
        async cargarDatos() {
            try {
                this.loading = true

                // Cargar salas
                const salasResponse = await axios.get('/api/salas')
                if (Array.isArray(salasResponse.data)) {
                    this.salasDestacadas = await Promise.all(
                        salasResponse.data.slice(0, 6).map(async (sala) => {
                            const salaData = {
                                id: sala.id,
                                nombre: sala.nombre,
                                icon: this.obtenerIconoPorEscuela(sala.escuela_id),
                                path: `/salas/${sala.id}`,
                                alertasCriticas: 0,
                                status: 'normal'
                            }

                            // Obtener alertas críticas
                            try {
                                const alertasResponse = await axios.get(`/api/alertas/sala/${sala.id}?tipo=critico`)
                                salaData.alertasCriticas = alertasResponse.data.length || 0
                                salaData.status = salaData.alertasCriticas > 0 ? 'critical' : 'normal'
                            } catch (error) {
                                console.warn(`No se pudieron cargar alertas para sala ${sala.id}`)
                            }

                            return salaData
                        })
                    )

                    this.estadisticas.totalSalas = salasResponse.data.length
                }

                // Cargar estadísticas adicionales
                await this.cargarEstadisticas()

            } catch (error) {
                console.error('Error al cargar datos:', error)
            } finally {
                this.loading = false
            }
        },

        async cargarEstadisticas() {
            try {
                // Simular datos de estadísticas (deberías obtenerlos de tu API)
                this.estadisticas = {
                    ...this.estadisticas,
                    sensoresActivos: Math.floor(Math.random() * 50) + 20,
                    alertasHoy: Math.floor(Math.random() * 10) + 2,
                    escuelas: Math.floor(Math.random() * 5) + 3
                }
            } catch (error) {
                console.error('Error al cargar estadísticas:', error)
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

        getSalaStatus(sala) {
            if (sala.alertasCriticas > 0) return 'Requiere atención'
            return 'Estado normal'
        },

        getSalaStatusColor(sala) {
            if (sala.alertasCriticas > 0) return 'error'
            return 'success'
        },

        getRandomMetric(min, max) {
            return (Math.random() * (max - min) + min).toFixed(1)
        },

        actualizarEstadisticasHome(alertas) {
            // Actualizar estadísticas basadas en las alertas cargadas
            const hoy = new Date().toDateString()
            this.estadisticas.alertasHoy = alertas.filter(alerta => {
                const fechaAlerta = new Date(alerta.fecha).toDateString()
                return fechaAlerta === hoy
            }).length
        }
    }
}
</script>

<style scoped>
.home {
    min-height: 100vh;
}

.hero-section {
    background: linear-gradient(135deg, #1976d2 0%, #1565c0 100%);
    color: white;
    padding: 80px 0;
    text-align: center;
}

.hero-title {
    font-size: 3.5rem;
    font-weight: 300;
    margin-bottom: 24px;
    line-height: 1.2;
}

.hero-subtitle {
    font-size: 1.4rem;
    opacity: 0.9;
    margin-bottom: 40px;
    max-width: 600px;
    margin-left: auto;
    margin-right: auto;
}

.hero-actions {
    margin-top: 32px;
}

.stats-section {
    padding: 60px 0;
    background: #f8f9fa;
}

.stat-card {
    height: 160px;
    transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.stat-card:hover {
    transform: translateY(-4px);
    box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15);
}

.stat-number {
    font-size: 2.5rem;
    font-weight: bold;
    margin-bottom: 8px;
    color: #1976d2;
}

.stat-label {
    font-size: 1rem;
    opacity: 0.8;
    text-transform: uppercase;
    letter-spacing: 0.5px;
}

.features-section,
.quick-access-section,
.recent-alerts-section {
    padding: 60px 0;
}

.features-section {
    background: white;
}

.quick-access-section {
    background: #f8f9fa;
}

.recent-alerts-section {
    background: white;
}

.section-title {
    font-size: 2.5rem;
    font-weight: 300;
    margin-bottom: 16px;
    color: #333;
}

.section-subtitle {
    font-size: 1.2rem;
    opacity: 0.7;
    margin-bottom: 0;
}

.feature-card {
    transition: transform 0.3s ease, box-shadow 0.3s ease;
    border: 2px solid transparent;
}

.feature-card:hover {
    transform: translateY(-4px);
    box-shadow: 0 12px 30px rgba(0, 0, 0, 0.1);
    border-color: #1976d2;
}

.feature-icon {
    margin-bottom: 24px;
}

.feature-title {
    font-size: 1.5rem;
    font-weight: 500;
    margin-bottom: 16px;
    color: #333;
}

.feature-description {
    font-size: 1rem;
    line-height: 1.6;
    opacity: 0.8;
    margin-bottom: 0;
}

.sala-card {
    transition: transform 0.2s ease, box-shadow 0.2s ease;
    border: 1px solid #e0e0e0;
}

.sala-card:hover {
    transform: translateY(-2px);
    box-shadow: 0 6px 20px rgba(0, 0, 0, 0.1);
    border-color: #1976d2;
}

.sala-name {
    font-size: 1.1rem;
    font-weight: 500;
    margin-bottom: 4px;
    color: #333;
}

.sala-status {
    color: #666;
}

.sala-metrics {
    display: flex;
    gap: 16px;
}

.metric-item {
    display: flex;
    align-items: center;
    color: #666;
}

.h-100 {
    height: 100%;
}

/* Responsive */
@media (max-width: 768px) {
    .hero-title {
        font-size: 2.5rem;
    }

    .hero-subtitle {
        font-size: 1.2rem;
    }

    .section-title {
        font-size: 2rem;
    }

    .hero-actions .v-btn {
        display: block;
        margin: 8px auto;
        width: 200px;
    }

    .sala-metrics {
        flex-direction: column;
        gap: 8px;
    }
}
</style>