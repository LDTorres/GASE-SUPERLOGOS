import axios from '../axios.js'

export default {
  async login ({ commit }, params) {
    await axios.post('/users/login', params.item).then((res) => {
      params.res = res.data
      commit('LOGIN', params)
    }).catch(() => {})
  },
  async register ({ commit }, params) {
    await axios.post('/users', params.item).then((res) => {
      params.res = res.data
      commit('LOGIN', params)
    }).catch(() => {})
  },
  async getAll ({ commit }, params) {
    await axios.get('/' + params.state).then((res) => {
      params.res = res.data
      commit('GET_ALL', params)
    }).catch(() => {})
  },
  async getOne ({ commit }, params) {
    await axios.get('/' + params.state + '/' + params.item.id).then((res) => {
      params.res = res.data
      commit('GET_ONE', params)
    }).catch(() => {})
  },
  async create ({ commit }, params) {
    await axios.post('/' + params.state, params.item).then((res) => {
      params.res = res.data
      commit('CREATE', params)
    }).catch(() => {})
  },
  async updateOne ({ commit }, params) {
    await axios.put('/' + params.state + '/' + params.item.id, params.item).then((res) => {
      params.res = res.data
      alert('El elemento fue actualizado')
    }).catch(() => {})
  },
  async deleteOne ({ commit }, params) {
    await axios.delete('/' + params.state + '/' + params.item.id).then((res) => {
      params.res = res.data
      commit('DELETE_ONE', params)
    }).catch(() => {})
  },
  async getAllTrashed ({ commit }, params) {
    await axios.get('/' + params.state + '/trashed').then((res) => {
      params.res = res.data
      commit('GET_ALL_TRASHED', params)
    }).catch(() => {})
  },
  async restore ({ commit }, params) {
    await axios.put('/' + params.state + '/' + params.item.id + '/restore').then((res) => {
      params.res = res.data
      commit('RESTORE', params)
    }).catch(() => {})
  },
  async trash ({ commit }, params) {
    await axios.delete('/' + params.state + '/' + params.item.id + '?trash=true').then((res) => {
      params.res = res.data
      commit('TRASH', params)
    }).catch(() => {})
  }
}
