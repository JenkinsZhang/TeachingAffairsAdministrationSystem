import Cookies from 'js-cookie'

export default function({ $axios, redirect }) {
  $axios.onRequest(config => {
    let token = null
    if (process.server) {
      let cookies = {}
      config.headers.common.cookie && config.headers.common.cookie.split(';').forEach(function(Cookie) {
        var parts = Cookie.split('=')
        cookies[parts[0].trim()] = (parts[1] || '').trim()
      })
      token = cookies.token
    } else {
      token = Cookies.get('token')
    }
    if (token) {
      config.headers.Authorization = token
    } else {
      console.error('无效的token')
    }
  })

  $axios.onError(error => {
    // const code = parseInt(error.response && error.response.status)
    console.error(error)
  })
}
