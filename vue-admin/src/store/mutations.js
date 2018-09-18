export const STORAGE_KEY = 'bazam'

// for testing
if (navigator.userAgent.indexOf('PhantomJS') > -1) {
  window.localStorage.clear()
}

export const mutations = {
  action (state, todo) {
    state.todos.push(todo)
  }
}
