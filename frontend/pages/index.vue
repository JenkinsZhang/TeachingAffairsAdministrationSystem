<template>
</template>

<script>
  import { getUserInfoFromToken } from '~/assets/js/tokenTools'

  export default {
    fetch({ redirect, req, app }) {
      let token = null
      if (process.server) {
        let cookies = {}
        req.headers.cookie && req.headers.cookie.split(';').forEach(function(Cookie) {
          var parts = Cookie.split('=')
          cookies[parts[0].trim()] = (parts[1] || '').trim()
        })
        token = cookies.token
      } else {
        token = app.$cookies.get('token')
      }
      if (token) {
        const info = getUserInfoFromToken(token)
        redirect('/' + info.identity)
      } else {
        redirect('/login')
      }
    }
  }
</script>

<style>
</style>
