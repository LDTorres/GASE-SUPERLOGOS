import axios from '../../axios.js'

/**
 * Route
 */

const route = '/portfolios'
var struct = [
  { text: 'Id', align: 'left', sortable: true, value: 'id' },
  { text: 'Nombre', value: 'name' },
  { text: 'Descripción', value: 'description' },
  { text: 'Cliente', value: 'client' },
  { text: 'Locación', value: 'location.name' },
  { text: 'Servicio', value: 'service.name' },
  { text: 'Actividad', value: 'activity.name' },
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
  let length = images.length

  formData.append('name', data.name)
  formData.append('description', data.description)
  formData.append('client', data.client)
  formData.append('priority', data.priority)
  formData.append('location[id]', data.location.id)
  formData.append('service[id]', data.service.id)
  formData.append('activity[id]', data.activity.id)

  // Set Images
  for (let i = 0; i < length; i++) {
    const element = images[i]
    formData.append('images', element)
  }

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
  let length = images.length

  formData.append('priority', data.priorityImage)
  formData.append('portfolio[id]', data.id)

  // Set Images
  for (let i = 0; i < length; i++) {
    const element = images[i]
    formData.append('images', element)
  }

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
      alert('El elemento fue creado')
    },
    UPDATE_ONE (state, data) {
      state.all[data.item.in] = data.item
      alert('El elemento fue actualizado')
    },
    UPLOAD_IMAGE (state, data) {
      if (state.all[data.item.in].images !== undefined) {
        state.all[data.item.in].images.push(data.res)
      } else {
        state.all[data.item.in].images = []
        state.all[data.item.in].images.push(data.res)
      }
      alert('El elemento fue actualizado')
    },
    DELETE_IMAGE (state, data) {
      state.all[data.item.in].images.splice(data.indexImage, 1)
      alert('El elemento fue eliminado')
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
      let formData = getForm(params.item)

      await axios.put('/' + params.state + '/' + params.item.id, formData, options).then((res) => {
        params.res = res.data
        commit('UPDATE_ONE', params)
      }).catch(() => {})
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
      await axios.put('/images/' + params.item.id, params.item).then((res) => {
        alert('Se ha colocado la prioridad')
      }).catch(() => {})
    }
  }
}
