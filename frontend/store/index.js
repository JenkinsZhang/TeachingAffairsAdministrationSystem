const cookieparser = process.server ? require('cookieparser') : undefined

export const state = () => ({
  token: null,
  xhrCount: 0,
  spinHideDelayTimer: null,
  spinShowDelayTimer: null
})

export const mutations = {
  setToken(state, auth) {
    state.token = auth
  },
  xhrInc(state) {
    state.xhrCount++
  },
  xhrDec(state) {
    state.xhrCount--
  },
  setSpinHideDelayTimer(state, val) {
    state.spinHideDelayTimer = val
  },
  setSpinShowDelayTimer(state, val) {
    state.spinShowDelayTimer = val
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
  },
  xhrIncWithSpin({ commit, state }, { app }) {

  },
  xhrDecWithSpin({ commit, state }, { app }) {

  }
}
