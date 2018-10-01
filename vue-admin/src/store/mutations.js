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
    state[data.state].editedItem = data.res
  },
  CREATE (state, data) {
    state[data.state].all.unshift(data.res)
  },
  UPDATE_ONE (state, data) {
    state[data.state].all[data.item.in] = data.item
  },
  DELETE_ONE (state, data) {
    state[data.state].all.splice(data.item.index, 1)
  },
  GET_ALL_TRASHED (state, data) {
    state[data.state].trashed = data.res
  },
  RESTORE (state, data) {
    state[data.state].trashed.splice(data.item.index, 1)
  },
  TRASH (state, data) {
    state[data.state].trashed.splice(data.item.index, 1)
  }
}
