// import axios from '../../axios.js'

/**
 * Route
 */

// const route = '/activities'

var struct = [
  { text: 'Id', align: 'left', sortable: true, value: 'id' },
  { text: 'Nombre', value: 'name' },
  { text: 'Descripción', value: 'description' },
  { text: 'Sector', value: 'sector.name' },
  {text: 'Acciones', align: 'center', value: ''}
]

export default {
  namespaced: true,
  state: { all: [], defaultItem: {}, editedItem: {}, trash: [], struct: struct },
  mutations: {
    GET_ALL (state, data) {
      state.all = data
    }
  },
  actions: {

  },
  getters: {}
}
