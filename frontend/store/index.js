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
    if (state.xhrCount === 0 && process.client) {
      clearTimeout(state.spinShowDelayTimer)
      commit('setSpinShowDelayTimer', setTimeout(() => {
        commit('setSpinShowDelayTimer', null)
        app.$Spin.show()
      }, 300))
      commit('xhrInc')
    }
  },
  xhrDecWithSpin({ commit, state }, { app }) {
    if (state.xhrCount === 1 && process.client) {
      clearTimeout(state.spinHideDelayTimer)
      clearTimeout(state.spinShowDelayTimer)
      commit('setSpinShowDelayTimer', null)
      commit('setSpinHideDelayTimer', setTimeout(() => {
        commit('setSpinHideDelayTimer', null)
        app.$Spin.hide()
      }, 100))
      commit('xhrDec')
    }
  }
}
