import Cookies from 'js-cookie'

export default function({ $axios, app, store, nuxtState }) {
  $axios.onRequest(config => {
    if (store.state.token) {
      config.headers.Authorization = store.state.token
    } else {
      console.warn('无法在headers添加Authorization', config.url)
    }
    store.dispatch('xhrIncWithSpin', { app })
  })
  $axios.onResponse(config => {
    store.dispatch('xhrDecWithSpin', { app })
    if (config.data.message && config.data.message === 'invalid token') {
      if (process.client) {
        app.$cookies.remove('token')
      }
      store.commit('setToken', null)
    }
  })

  $axios.onError(error => {
    store.dispatch('xhrDecWithSpin', { app })
    // const code = parseInt(error.response && error.response.status)
    console.error(error)
  })
}
