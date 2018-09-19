export const STORAGE_KEY = 'bazam'

// for testing
if (navigator.userAgent.indexOf('PhantomJS') > -1) {
  window.localStorage.clear()
}

export const mutations = {
  GET_ALL (state, data) {
    state[data.state].all = data.res
  },
  GET_ONE (state, data) {
    state[data.state].all.push(data.res)
  },
  UPDATE_ONE (state, data) {
    state[data.state].all[data.index] = data.item
  },
  DELETE_ONE (state, data) {
    state[data.state].all.splice(data.index, 1)
  }
}
