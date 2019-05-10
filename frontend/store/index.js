import Vuex from 'vuex'

const cookieparser = process.server ? require('cookieparser') : undefined

export const state = () => ({
  token: null
})

export const mutations = {
  setToken(state, auth) {
    state.token = auth
  }
}

export const actions = {
  nuxtServerInit({ commit }, { req }) {
    let token = null
    if (req.headers.cookie) {
      const parsed = cookieparser.parse(req.headers.cookie)
      token = parsed.token
    }
    commit('setToken', token)
  }
}
