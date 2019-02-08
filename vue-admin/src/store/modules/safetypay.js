import axios from '../../axios.js'

/**
 * Route
 */

const route = '/gateways/safetypay'
var struct = [
  { text: 'Fecha', sortable: true, value: 'CreationDateTime' },
  { text: 'Id de la operación', value: 'OperationID' },
  { text: 'Id de la orden', value: 'MerchantOrderID' },
  { text: 'Estado de la operación', value: 'OperationStatus' },
  { text: 'Confirmar Pago', align: 'center', value: '' }
]

export default {
  namespaced: true,
  state: { all: [], defaultItem: { priority: 0 }, editedItem: {}, trashed: [], struct: struct },
  getters: {
    getAll: (state) => (stateName) => {
      return state.all
    }
  },
  mutations: {
    GET_ALL (state, data) {
      state.all = data
    }
  },
  actions: {
    async getAll ({commit}) {
      await axios.get(route + '/notifications').then((res) => {
        commit('GET_ALL', res.data)
      }).catch(() => {})
    },
    async confirm ({ actions }, params) {
      await axios.put(route + '/notifications/confirm', params).then((res) => {
        location.reload()
      }).catch(() => {})
    }
  }
}
