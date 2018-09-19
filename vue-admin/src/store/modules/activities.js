// import axios from '../../axios.js'

/**
 * Route
 */

// const route = '/activities'

export default {
  namespaced: true,
  state: { all: [], defaultItem: {}, editedItem: {} },
  mutations: {
    GET_ALL (state, data) {
      state.all = data
    }
  },
  actions: {

  },
  getters: {}
}
