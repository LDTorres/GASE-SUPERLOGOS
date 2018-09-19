import axios from '../../axios.js'

/**
 * Route
 */

const route = '/payments-methods'

export default {
  namespaced: true,
  state: {all: []},
  mutations: {
    GET_ALL (state, data) {
      state.all = data
    }
  },
  actions: {
    async getAll ({commit}) {
      try {
        let res = await axios.get(route)
        commit('GET_ALL', res.data)
      } catch (error) {
        console.log(error)
      }
    }
  },
  getters: { }
}
