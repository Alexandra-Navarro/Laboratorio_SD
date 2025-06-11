import { createRouter, createWebHistory } from 'vue-router'

// Importar las vistas
import Home from '@/views/Home.vue'
import SalaDetail from '@/views/SalaDetail.vue'
import AlertasView from '@/views/AlertasView.vue'
import Dashboard from '@/views/Dashboard.vue'

const routes = [
  // Ruta principal
  {
    path: '/',
    name: 'Home',
    component: Home,
    meta: {
      title: 'Sistema de Monitoreo Ambiental'
    }
  },
  
  // Dashboard general
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: Dashboard,
    meta: {
      title: 'Dashboard General'
    }
  },
  
  // Detalle de sala específica
  {
    path: '/salas/:id',
    name: 'SalaDetail',
    component: SalaDetail,
    props: true,
    meta: {
      title: 'Detalle de Sala'
    }
  },
  
  // Alertas de una sala específica
  {
    path: '/salas/:id/alertas',
    name: 'SalaAlertas',
    component: AlertasView,
    props: true,
    meta: {
      title: 'Alertas de Sala'
    }
  },
  
  // Todas las alertas del sistema
  {
    path: '/alertas',
    name: 'TodasAlertas',
    component: AlertasView,
    meta: {
      title: 'Todas las Alertas'
    }
  },
  
  // Redirigir rutas no encontradas al home
  {
    path: '/:pathMatch(.*)*',
    redirect: '/'
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
  // Scroll behavior para mejorar UX
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    }
    return { top: 0 }
  }
})

// Navigation guard para actualizar el título de la página
router.beforeEach((to, from, next) => {
  // Actualizar título de la página
  if (to.meta.title) {
    document.title = `${to.meta.title} - Sistema de Monitoreo`
  }
  
  next()
})

export default router