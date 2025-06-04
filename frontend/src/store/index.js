import { createStore } from 'vuex'
import axios from 'axios'

const API_URL = process.env.VUE_APP_API_URL || 'http://localhost:8080/api'

export default createStore({
  state: {
    sensors: [],
    alerts: [],
    thresholds: [],
    loading: false,
    error: null
  },
  mutations: {
    SET_SENSORS(state, sensors) {
      state.sensors = sensors
    },
    SET_ALERTS(state, alerts) {
      state.alerts = alerts
    },
    SET_THRESHOLDS(state, thresholds) {
      state.thresholds = thresholds
    },
    SET_LOADING(state, loading) {
      state.loading = loading
    },
    SET_ERROR(state, error) {
      state.error = error
    }
  },
  actions: {
    async fetchSensors({ commit }) {
      commit('SET_LOADING', true)
      try {
        const response = await axios.get(`${API_URL}/sensors`)
        commit('SET_SENSORS', response.data)
      } catch (error) {
        commit('SET_ERROR', error.message)
      } finally {
        commit('SET_LOADING', false)
      }
    },
    async fetchAlerts({ commit }) {
      commit('SET_LOADING', true)
      try {
        const response = await axios.get(`${API_URL}/alerts`)
        commit('SET_ALERTS', response.data)
      } catch (error) {
        commit('SET_ERROR', error.message)
      } finally {
        commit('SET_LOADING', false)
      }
    },
    async fetchThresholds({ commit }) {
      commit('SET_LOADING', true)
      try {
        const response = await axios.get(`${API_URL}/thresholds`)
        commit('SET_THRESHOLDS', response.data)
      } catch (error) {
        commit('SET_ERROR', error.message)
      } finally {
        commit('SET_LOADING', false)
      }
    },
    async updateThreshold({ commit }, threshold) {
      commit('SET_LOADING', true)
      try {
        const response = await axios.put(`${API_URL}/thresholds/${threshold.id}`, threshold)
        commit('SET_THRESHOLDS', response.data)
      } catch (error) {
        commit('SET_ERROR', error.message)
      } finally {
        commit('SET_LOADING', false)
      }
    }
  },
  getters: {
    getSensorsByRoom: (state) => (roomId) => {
      return state.sensors.filter(sensor => sensor.room_id === roomId)
    },
    getAlertsByRoom: (state) => (roomId) => {
      return state.alerts.filter(alert => alert.room_id === roomId)
    },
    getThresholdsByRoom: (state) => (roomId) => {
      return state.thresholds.filter(threshold => threshold.room_id === roomId)
    }
  }
}) 