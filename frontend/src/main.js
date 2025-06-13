import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

// Vuetify
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import '@mdi/font/css/materialdesignicons.css'

const vuetify = createVuetify({
  components,
  directives,
  theme: {
    defaultTheme: 'light',
    themes: {
      light: {
        colors: {
          primary: '#00b4db',
          secondary: '#0083b0',
          accent: '#00b4db',
          error: '#f44336',
          warning: '#ff9800',
          info: '#00b4db',
          success: '#4caf50'
        }
      }
    }
  }
})

const app = createApp(App)

app.use(router)
app.use(store)
app.use(vuetify)

app.mount('#app') 