<template>
  <v-app>
    <!-- Solo mostrar navegación si el usuario está logueado y no está en login -->
    <template v-if="isLoggedIn && !isLoginPage">
      <v-navigation-drawer v-model="drawer" app>
        <v-list>
          <v-list-item
            v-for="item in menuItems"
            :key="item.title"
            :to="item.path"
            link
          >
            <v-list-item-icon>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title>{{ item.title }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-navigation-drawer>

      <v-app-bar app color="primary" dark>
        <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
        <v-toolbar-title>Sistema de Monitoreo Ambiental</v-toolbar-title>
        <v-btn @click="logout" color="red" variant="tonal">Cerrar sesión</v-btn>
      </v-app-bar>
    </template>

    <v-main>
      <v-container fluid>
        <router-view></router-view>
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()
const drawer = ref(true)

const isLoggedIn = computed(() => !!localStorage.getItem('usuario'))
const isLoginPage = computed(() => route.path === '/')

const menuItems = [
  {
    title: 'Dashboard',
    icon: 'mdi-view-dashboard',
    path: '/dashboard'
  },
  {
    title: 'Sensores',
    icon: 'mdi-thermometer',
    path: '/sensors'
  },
  {
    title: 'Alertas',
    icon: 'mdi-alert',
    path: '/alerts'
  },
  {
    title: 'Configuración',
    icon: 'mdi-cog',
    path: '/settings'
  }
]

const logout = () => {
  localStorage.removeItem('usuario')
  router.push('/')
}
</script>

<style>
@import '@mdi/font/css/materialdesignicons.css';
</style> 