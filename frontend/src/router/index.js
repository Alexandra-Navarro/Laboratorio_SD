import { createRouter, createWebHistory } from 'vue-router'

// Importar las vistas
import Home from '@/views/Home.vue'
import SalaDetail from '@/views/SalaDetail.vue'
import AlertasView from '@/views/AlertasView.vue'
import Dashboard from '@/views/Dashboard.vue'
import Login from '../views/Login.vue'

const routes = [
  {
    path: '/',
    name: 'Login',
    component: Login
  },
  {
    path: '/home',
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

// 🔐 Protección de rutas
router.beforeEach((to, from, next) => {
  const publicPages = ['/']
  const authRequired = !publicPages.includes(to.path)
  const usuario = localStorage.getItem('usuario')

  if (authRequired && !usuario) {
    return next('/')
  }

  // Si el usuario está logueado y va a la página de login, redirigir al dashboard
  if (to.path === '/' && usuario) {
    return next('/dashboard')
  }

  next()
})

export default router