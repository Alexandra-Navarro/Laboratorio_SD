<template>
  <v-container class="fill-height login-container" fluid>
    <v-row justify="center" align="center" class="login-row">
      <v-col cols="12" sm="8" md="5" lg="4">
        <div class="login-wrapper">
          <v-card class="login-card" elevation="2">
            <v-card-text class="pa-8">
              <div class="text-center mb-8">
                <div class="login-icon mb-4">
                  <v-icon size="56" color="#00b4db">mdi-waves</v-icon>
                </div>
                <h2 class="login-title">Bienvenido</h2>
                <p class="login-subtitle">Ingresa a tu cuenta</p>
              </div>

              <v-form @submit.prevent="handleLogin">
                <v-text-field
                  v-model="email"
                  label="Email"
                  type="email"
                  prepend-inner-icon="mdi-email-outline"
                  variant="outlined"
                  color="#00b4db"
                  class="mb-4 custom-input"
                  density="comfortable"
                  bg-color="white"
                />

                <v-text-field
                  v-model="password"
                  label="Contraseña"
                  type="password"
                  prepend-inner-icon="mdi-lock-outline"
                  variant="outlined"
                  color="#00b4db"
                  class="mb-6 custom-input"
                  density="comfortable"
                  bg-color="white"
                />

                <v-btn
                  type="submit"
                  color="#00b4db"
                  size="large"
                  block
                  :loading="loading"
                  class="login-btn mb-4"
                  variant="elevated"
                  elevation="4"
                >
                  <v-icon left class="mr-2">mdi-login-variant</v-icon>
                  Iniciar Sesión
                </v-btn>
              </v-form>

              <v-alert
                v-if="error"
                type="error"
                variant="tonal"
                class="mt-4 error-alert"
                icon="mdi-alert-circle-outline"
              >
                {{ error }}
              </v-alert>
            </v-card-text>
          </v-card>
        </div>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

const email = ref('')
const password = ref('')
const error = ref('')
const loading = ref(false)
const router = useRouter()

const handleLogin = async () => {
  error.value = ''
  loading.value = true

  try {
    const res = await axios.post('http://localhost:8080/api/login', {
      email: email.value,
      password: password.value,
    })
    localStorage.setItem('usuario', JSON.stringify(res.data.usuario))
    router.push('/dashboard')
  } catch (err) {
    error.value = 'Credenciales inválidas. Verifica tu email y contraseña.'
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-container {
  background: #fff;
  min-height: 100vh;
  position: fixed;
  width: 100vw;
  height: 100vh;
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
}

.login-row {
  min-height: 100vh;
  align-items: center !important;
  justify-content: center !important;
}

.login-wrapper {
  position: relative;
  z-index: 1;
}

.login-card {
  background: rgba(255, 255, 255, 0.98);
  border-radius: 24px;
  border: 1px solid #e0e0e0;
  box-shadow: 0 8px 32px rgba(0, 180, 219, 0.08);
}

.login-icon {
  width: 80px;
  height: 80px;
  background: #e0f7fa;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto;
  border: 2px solid #00b4db33;
}

.login-title {
  color: #00b4db;
  font-size: 2rem;
  font-weight: 400;
  margin-bottom: 8px;
}

.login-subtitle {
  color: #666;
  font-size: 1rem;
  font-weight: 300;
  margin: 0;
}

.custom-input {
  background: white !important;
  border-radius: 12px !important;
  box-shadow: 0 2px 8px rgba(0,180,219,0.04);
}

:deep(.custom-input .v-field) {
  border-radius: 12px !important;
  background: white !important;
  border: 1.5px solid #e0e0e0 !important;
}

:deep(.custom-input .v-field--focused) {
  border-color: #00b4db !important;
  box-shadow: 0 0 0 2px #00b4db22;
}

:deep(.custom-input .v-field__input) {
  color: #222;
  font-weight: 400;
}

:deep(.custom-input .v-field__input::placeholder) {
  color: #aaa;
}

:deep(.custom-input .v-label) {
  color: #00b4db !important;
  background: white !important;
  z-index: 2;
}

:deep(.custom-input .v-icon) {
  color: #00b4db;
}

.login-btn {
  border-radius: 16px;
  color: #fff !important;
  font-weight: 600;
  font-size: 1rem;
  text-transform: none;
  height: 56px;
  background: linear-gradient(90deg, #00b4db 0%, #0083b0 100%) !important;
  box-shadow: 0 4px 16px rgba(0,180,219,0.10);
  transition: all 0.3s ease;
}

.login-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(0, 180, 219, 0.15);
}

.error-alert {
  background: rgba(244, 67, 54, 0.1);
  color: #ff5252;
  border: 1px solid rgba(244, 67, 54, 0.3);
  border-radius: 12px;
  backdrop-filter: blur(10px);
}

:deep(.error-alert .v-icon) {
  color: #ff5252;
}

@media (max-width: 600px) {
  .login-card {
    margin: 16px;
    border-radius: 20px;
  }
  .login-title {
    font-size: 1.5rem;
  }
  .pa-8 {
    padding: 24px !important;
  }
}
</style>