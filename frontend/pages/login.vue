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
  export default {
    name: 'login',
    data: () => ({
      host: '',
      datalist: {},
      loading: false,
      username: '',
      password: ''
    }),
    asyncData({ req }) {
      if (process.server) {
        return { host: req.headers.host, datalist: { 'ok': true } }
      }
    },
    head() {
      return {
        title: '登录'
      }
    },
    methods: {
      login() {

      }
    }
  }
</script>

<style lang="scss" scoped>
  @import "./login";
</style>
