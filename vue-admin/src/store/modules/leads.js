import axios from '../../axios.js'

/**
 * Route
 */

const route = '/leads'
var struct = [
  { text: 'Id', align: 'left', sortable: true, value: 'id' },
  { text: 'Correo', value: 'email' },
  { text: 'Teléfono', value: 'phone' },
  { text: 'Horario', value: 'schedule' },
  { text: 'Medio', value: 'medium' },
  { text: 'Pais', value: 'countries.name' },
  {text: 'Eliminar', align: 'center', value: ''}
]

export default {
  namespaced: true,
  state: { all: [], defaultItem: {}, editedItem: {}, trashed: [], struct: struct },
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
