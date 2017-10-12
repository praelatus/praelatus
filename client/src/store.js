import Vue from 'vue'
import Vuex from 'vuex'
import createPersistedState from 'vuex-persistedState'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    sidebarWidth: '250px',
    currentUser: null,
    token: '',
    showSidebar: false
  },

  getters: {
    showSidebar: function (state) {
      return state.showSidebar
    },

    sidebarWidth: function (state) {
      return state.sidebarWidth
    },

    currentUser: function (state) {
      return state.currentUser
    },

    token: function (state) {
      return state.token
    }
  },

  mutations: {
    setSidebarWidth: function (state, width) {
      if (typeof width === 'number') {
        state.sidebarWidth = width.toString() + 'px'
      } else if (width) {
        state.sidebarWidth = width
      } else {
        state.sidebarWidth = '250px'
      }
    },

    login: function (state, { token, user }) {
      state.currentUser = user
      state.token = token
    },

    logout: function (state) {
      state.currentUser = null
      state.token = null
    },

    sidebarShown: function (state, show) {
      state.showSidebar = show
    }
  },

  plugins: [createPersistedState()]
})
