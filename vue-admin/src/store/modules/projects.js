import axios from '../../axios.js'

/**
 * Route
 */

const route = '/projects'
var struct = [
  { text: 'Id', align: 'left', sortable: true, value: 'id' },
  { text: 'Nombre', value: 'name' },
  { text: 'Cliente', value: 'client' },
  { text: 'Estado', value: 'closed' },
  { text: 'Publico', value: 'public' },
  { text: 'Servicio', value: 'service.name' },
  {text: 'Acciones', align: 'center', value: ''}
]
const options = { headers: { 'Content-Type': 'multipart/form-data' } }

var getForm = (data) => {
  let formData = new FormData()
  let images
  if (data.files) {
    images = data.files
  } else if (data.file) {
    images = data.file
  }
  // Set Images
  if (images) {
    let length = images.length
    for (let i = 0; i < length; i++) {
      const element = images[i]
      formData.append('images', element)
    }
  }

  formData.append('name', data.name)
  formData.append('description', data.description)
  formData.append('client', data.client)
  formData.append('priority', data.priority)
  formData.append('location[id]', data.location.id)
  formData.append('service[id]', data.service.id)
  formData.append('activity[id]', data.activity.id)

  return formData
}

var getFormUploadImage = (data) => {
  let formData = new FormData()
  let images
  if (data.files) {
    images = data.files
  } else if (data.file) {
    images = data.file
  }
  // Set Images
  if (images) {
    let length = images.length
    for (let i = 0; i < length; i++) {
      const element = images[i]
      formData.append('images', element)
    }
  }

  formData.append('priority', data.priorityImage)
  formData.append('portfolio[id]', data.id)

  return formData
}

export default {
  namespaced: true,
  state: { all: [], defaultItem: {}, editedItem: {}, trashed: [], struct: struct },
  mutations: {
    GET_ALL (state, data) {
      state.all = data
    },
    CREATE (state, data) {
      state.all.push(data.res)
    },
    UPDATE_ONE (state, data) {
      state.all[data.item.in] = data.item
    },
    UPLOAD_IMAGE (state, data) {
      if (state.all[data.item.in].images !== undefined) {
        state.all[data.item.in].images.push(data.res)
      } else {
        state.all[data.item.in].images = []
        state.all[data.item.in].images.push(data.res)
      }
    },
    DELETE_IMAGE (state, data) {
      state.all[data.item.in].images.splice(data.indexImage, 1)
    }
  },
  actions: {
    async getAll ({commit}) {
      await axios.get(route).then((res) => {
        commit('GET_ALL', res.data)
      }).catch(() => {})
    },
    async create ({ commit }, params) {
      let formData = getForm(params.item)
      let item = params.item

      await axios.post('/' + params.state, formData, options).then((res) => {
        params.res = res.data
        params.res.location = item.location
        params.res.service = item.service
        params.res.activity = item.activity
        commit('CREATE', params)
      }).catch(() => {})
    },
    async updateOne ({ commit }, params) {
      if (!params.edited) {
        let formData = getForm(params.item)

        await axios.put('/' + params.state + '/' + params.item.id, formData, options).then((res) => {
          params.res = res.data
          commit('UPDATE_ONE', params)
        }).catch(() => {})
      } else {
        await axios.put('/' + params.state + '/' + params.item.id, params.item).then((res) => {
          params.res = res.data
          commit('UPDATE_ONE', params)
        }).catch(() => {})
      }
    },
    async uploadImage ({ commit }, params) {
      let formData = getFormUploadImage(params.item)

      await axios.post('/images', formData, options).then((res) => {
        params.res = res.data
        commit('UPLOAD_IMAGE', params)
      }).catch(() => {})
    },
    async imageDelete ({ commit }, params) {
      let id = params.item.images[params.indexImage].id
      await axios.delete('/images/' + id + '?trash=true').then((res) => {
        params.res = res.data
        commit('DELETE_IMAGE', params)
      }).catch(() => {})
    },
    async imagePriority ({ commit }, params) {
      await axios.put('/images/' + params.item.id, params.item)
    }
  }
}
