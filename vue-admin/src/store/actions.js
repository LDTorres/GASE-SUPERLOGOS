import axios from '../axios.js'

export default {
  async getAll ({ commit }, params) {
    try {
      let res = await axios.get(params.route)
      params.res = res.data
      commit('GET_ALL', params)
    } catch (error) {
      console.log(error)
    }
  },
  async getOne ({ commit }, params) {
    try {
      let res = await axios.get(params.route + '/' + params.item.id)
      params.res = res.data
      commit('GET_ONE', params)
    } catch (error) {
      console.log(error)
    }
  },
  async create ({ commit }, params) {
    try {
      let res = await axios.post(params.route, params.item)
      params.res = res.data
      commit('CREATE', params)
    } catch (error) {
      console.log(error)
    }
  },
  async updateOne ({ commit }, params) {
    try {
      let res = await axios.put(params.route + '/' + params.item.id, params.item)
      params.res = res.data
      commit('UPDATE_ONE', params)
    } catch (error) {
      console.log(error)
    }
  },
  async deleteOne ({ commit }, params) {
    try {
      let res = await axios.delete(params.route + '/' + params.item.id + '?trash=false')
      params.res = res.data
      commit('DELETE_ONE', params)
    } catch (error) {
      console.log(error)
    }
  }
}
