export default {
  namespaced: true,
  state: () => ({
    token: ''
  }),
  actions: {
    setToken({ commit }, token) {
      commit('setTonken', token)
    }
  },
    mutations: {
      setToken(state, token) {
        state.token = token
      }
    },
    getters: {},
  }

