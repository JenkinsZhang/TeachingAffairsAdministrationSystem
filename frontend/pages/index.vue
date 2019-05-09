<template>
</template>

<script>
  import { getUserInfoFromToken } from '~/assets/js/tokenTools'

  export default {
    fetch({ redirect, req }) {
      let token = null
      if (process.server) {
        token = req.headers.Authorization
      } else {
        token = localStorage.getItem('token')
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
