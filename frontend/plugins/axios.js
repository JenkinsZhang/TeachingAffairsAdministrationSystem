import Cookies from 'js-cookie'

export default function({ $axios, redirect, store }) {
  $axios.onRequest(config => {
    if (store.state.token) {
      config.headers.Authorization = store.state.token
    } else {
      console.warn('无法在headers添加Authorization', config.url)
    }
  })

  $axios.onError(error => {
    // const code = parseInt(error.response && error.response.status)
    console.error(error)
  })
}
