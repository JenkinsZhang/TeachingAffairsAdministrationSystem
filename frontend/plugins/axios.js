export default function({ $axios, app, store, nuxtState }) {
  $axios.onRequest(config => {
    if (store.state.token) {
      config.headers.Authorization = store.state.token
    } else {
      console.warn('无法在headers添加Authorization', config.url)
    }
    const { state, commit } = store
    if (state.xhrCount === 0 && process.client) {
      clearTimeout(state.spinShowDelayTimer)
      commit('setSpinShowDelayTimer', setTimeout(() => {
        commit('setSpinShowDelayTimer', null)
        app.$Spin.show()
      }, 300))
      commit('xhrInc')
    }
  })
  $axios.onResponse(config => {
    const { state, commit } = store
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
  })

  $axios.onError(error => {
    // const code = parseInt(error.response && error.response.status)
    console.error(error)
  })
}
