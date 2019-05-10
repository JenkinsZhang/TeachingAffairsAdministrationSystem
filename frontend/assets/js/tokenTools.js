import Cookies from 'js-cookie'

var Base64 = require('js-base64').Base64

export function getUserInfoFromToken(token = null, req) {
  if (!token) {
    if (process.server) {
      let cookies = {}
      req.headers.cookie && req.headers.cookie.split(';').forEach(function(Cookie) {
        var parts = Cookie.split('=')
        cookies[parts[0].trim()] = (parts[1] || '').trim()
      })
      token = cookies.token
    } else {
      token = Cookies.get('token')
    }
    if (!token) {
      return null
    }
  }

  return JSON.parse(decodeURIComponent(Base64.decode(token.split(/\./)[1])))
}
