import { createRouter, createWebHistory } from 'vue-router'
import Dashboard from '../views/Dashboard.vue'
import Sensors from '../views/Sensors.vue'
import Alerts from '../views/Alerts.vue'
import Settings from '../views/Settings.vue'
import Login from '../views/Login.vue'

const routes = [
  {
    path: '/',
    name: 'Login',
    component: Login
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: Dashboard
  },
  {
    path: '/sensors',
    name: 'Sensors',
    component: Sensors
  },
  {
    path: '/alerts',
    name: 'Alerts',
    component: Alerts
  },
  {
    path: '/settings',
    name: 'Settings',
    component: Settings
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
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
