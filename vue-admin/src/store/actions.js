import axios from '../axios.js'

export default {
  async login ({ state }, params) {
    await axios.post('/users/login', params.item).then((res) => {
      localStorage.setItem('bazam-token', res.data.token)
      delete res.data.token
      state.user = res.data
      localStorage.setItem('bazam-user', JSON.stringify(res.data))
      location.reload()
    }).catch(() => {})
  },
  async register ({ state }, params) {
    await axios.post('/users', params.item).then((res) => {
      localStorage.setItem('bazam-token', res.data.token)
      delete res.data.token
      state.user = res.data
      localStorage.setItem('bazam-user', JSON.stringify(res.data))
      location.reload()
    }).catch(() => {})
  },
  async getAll ({ commit }, params) {
    await axios.get('/' + params.state + '?limit=1000').then((res) => {
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
      commit('UPDATE_ONE', params)
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
