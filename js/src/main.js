// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import Vuex from 'vuex'

import App from './App'
import router from './router'

Vue.config.productionTip = false

Vue.use(Vuex)

const store = new Vuex.Store({
  state: { todos: [] },
  actions: {
    initialize ({ commit }) {
      fetch('http://localhost:9000/api/todos')
        .then(response => response.json())
        .then(data => commit('setTodos', data))
    }
  },
  mutations: {
    setTodos (state, todos) {
      state.todos = todos
    }
  }
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
