import { createRouter, createWebHistory } from 'vue-router'

// Importar las vistas
import Home from '@/views/Home.vue'
import SalaDetail from '@/views/SalaDetail.vue'
import AlertasView from '@/views/AlertasView.vue'
//import Dashboard from '@/views/Dashboard.vue'
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
  /*{
    path: '/dashboard',
    name: 'Dashboard',
    component: Dashboard,
    meta: {
      title: 'Dashboard General'
    }
  },*/
  {
    path: '/salas/:id',
    name: 'SalaDetail',
    component: SalaDetail,
    props: true,
    meta: {
      title: 'Detalle de Sala'
    }
  },
  {
    path: '/salas/:id/alertas',
    name: 'SalaAlertas',
    component: AlertasView,
    props: true,
    meta: {
      title: 'Alertas de Sala'
    }
  },
  {
    path: '/alertas',
    name: 'TodasAlertas',
    component: AlertasView,
    meta: {
      title: 'Todas las Alertas'
    }
  },
  
  {
    path: '/:pathMatch(.*)*',
    redirect: '/home'
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

// Navigation guard para manejar títulos y autenticación
router.beforeEach((to, from, next) => {
  // Actualizar título de la página
  if (to.meta.title) {
    document.title = `${to.meta.title} - Sistema de Monitoreo`
  }

  // Lógica de autenticación
  const publicPages = ['/']
  const authRequired = !publicPages.includes(to.path)
  const usuario = localStorage.getItem('usuario')

  if (authRequired && !usuario) {
    return next('/')
  }

  // Si el usuario está logueado y va a la página de login, redirigir al dashboard
  if (to.path === '/' && usuario) {
    return next('/home')
  }

  next()
})

export default router