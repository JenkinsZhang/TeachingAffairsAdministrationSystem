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
  })

  $axios.onError(error => {
    store.dispatch('xhrDecWithSpin', { app })
    // const code = parseInt(error.response && error.response.status)
    console.error(error)
  })
}
