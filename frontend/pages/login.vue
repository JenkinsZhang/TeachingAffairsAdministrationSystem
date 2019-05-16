<template>
  <div class="login">
    <div class="filter">
      <div class="SysTitle">
        教学事务管理系统
        <div class="login-box">
          <Input v-model="username" style="margin-bottom: 10px" size="large" prefix='md-person' placeholder="用户名"/>
          <Input v-model="password" type="password" style="margin-bottom: 10px" prefix="md-lock" size="large"
                 placeholder="密码"/>
          <ButtonGroup size="large" style="margin-top:10px">
            <div>
              <Button size="large" type="warning">
                忘记密码？
              </Button>
              <Button size="large" type="primary" :loading="loading" @click="login">
                登录
              </Button>
            </div>
          </ButtonGroup>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import { getUserInfoFromToken } from '~/assets/js/tokenTools'

  export default {
    name: 'login',
    middleware: 'anonymous',
    data: () => ({
      host: '',
      datalist: {},
      loading: false,
      username: '',
      password: ''
    }),
    head() {
      return {
        title: '登录'
      }
    },
    methods: {
      login() {
        this.$axios({
          url: '/login',
          method: 'post',
          data: {
            id: this.username,
            password: this.password
          }
        }).then((res) => {
          if (res.data.message === 'ok') {
            this.$store.commit('setToken', res.data.token) // 存储在vuex中用来进行客户端渲染
            this.$cookies.set('token', res.data.token, { expires: 1 }) // 在cookie中保存token用来进行服务器端渲染
            this.$router.push('/')
          } else {
            this.$Notice.warning({
              title: '登录失败',
              desc: res.data.message
            })
          }
        })
      }
    }
  }
</script>

<style lang="scss" scoped>
  @import "./login";
</style>
