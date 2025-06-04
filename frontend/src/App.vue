<template>
  <v-app>
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
    </v-app-bar>

    <v-main>
      <v-container fluid>
        <router-view></router-view>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
export default {
  name: 'App',
  data: () => ({
    drawer: true,
    menuItems: [
      {
        title: 'Dashboard',
        icon: 'mdi-view-dashboard',
        path: '/'
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
  })
}
</script>

<style>
@import '@mdi/font/css/materialdesignicons.css';
</style> 