<template>
  <div class="layout">
    <Layout style="height: 100vh;">
      <Sider
        ref="side1"
        hide-trigger
        collapsible
        :collapsed-width="78"
        v-model="isCollapsed"
      >
        <div class="layout-logo-left">
          <span>教学</span><span>事务管理系统</span>
        </div>
        <SideMenu
          :collapse="isCollapsed"
          v-model="sidebarSelectedTitle"
          :menu="menu"
        ></SideMenu>
      </Sider>
      <Layout>
        <Header
          :style="{padding: 0}"
          class="layout-header-bar"
          id="layout-header-bar"
        >
          <Icon
            @click.native="collapsedSider"
            :class="rotateIcon"
            :style="{margin: '0 20px'}"
            type="md-menu"
            size="24"
          ></Icon>
          <div
            class="layout-header-title"
            id="layout-header-title"
            v-html="sidebarSelectedTitle"
          ></div>
          <div
            class="userBox"
            @click="showUserPopTip = !showUserPopTip;$event.stopPropagation()"
          >
            <div
              class="userBox-avatar"
              :style="{backgroundImage: 'url(http://mzzeast.shumsg.cn/FjSAoF6FqSXeDpu0aRI7eIOhGbM-)'}"
            ></div>
            <div class="userBox-content">
              <p class="userBox-content-name">{{profile.name}}</p>
              <!-- <p class="userBox-content-type">个人</p> -->
            </div>
            <Icon
              type="ios-arrow-down"
              size="15"
              color="#000"
            ></Icon>
          </div>
        </Header>
        <div
          class="user-PopTip"
          v-show="showUserPopTip"
        >
          <ul>
            <li @click="$router.push('profile')">
              <Icon
                type="md-checkmark"
                size="20"
              ></Icon>
              {{profile.name}}
              <span class="right-align" style="margin-right:10px">{{profile.identity|identityMap}}</span>
            </li>
            <li class="divider"></li>
            <li @click="$router.push('profile')">
              <Icon
                type="ios-person"
                size="20"
              ></Icon>
              个人信息
            </li>
            <li @click="logout">
              <Icon
                type="md-exit"
                size="20"
              ></Icon>
              登出
            </li>
          </ul>
        </div>
        <Content :style="{margin: '20px', background: '#fff'}">
          <div style="position: relative">
            <nuxt-child></nuxt-child>
          </div>
        </Content>
      </Layout>
    </Layout>
  </div>
</template>

<script>
  import SideMenu from '~/components/sideMenu/sideMenu'
  import stickybits from 'stickybits'
  import menuJson from '~/assets/json/menu.json'
  import identityMap from '~/assets/json/identity.json'
  import { getUserInfoFromToken } from '~/assets/js/tokenTools'

  export default {
    name: 'users',
    components: { SideMenu },
    middleware: 'authenticated',
    validate({ params, req, error, app }) {
      const info = getUserInfoFromToken(null, req)
      console.log('validate', !!info, !!identityMap[params.users], !!menuJson[params.users], info.identity === params.users)
      return !!info && !!identityMap[params.users] && !!menuJson[params.users] && info.identity === params.users
    },
    head() {
      return {
        title: identityMap[this.profile.identity] + '端'
      }
    },
    filters: {
      identityMap(k) {
        return identityMap[k]
      }
    },
    data() {
      return {
        isCollapsed: true,
        showUserPopTip: false,
        profile: {
          birthplace: '河北',
          dname: '计算机学院',
          gender: '男',
          grade: '大三',
          id: '16121674',
          name: '苗伟华',
          phone: '16601700694'
        },
        sidebarSelectedTitle: '',
        menu: []
      }
    },
    asyncData({ params, app, route, redirect, req }) {//tip: 在客户端可能会多次触发，例如从404页面回退时
      //特判
      if (route.path.substr(-params.users.length) === params.users) {
        return
      }
      let menu = menuJson[params.users]
      console.log('asyncData', process.server)
      menu.forEach((g) => {
        g.items.forEach((i) => {
          if (i.done) {
            return
          }
          i.name = `/${params.users}/${i.name}`
          i.done = true
        })
      })
      let profile = {}
      profile = getUserInfoFromToken(null, req)
      console.log(profile)
      return { menu, profile }
    },
    mounted() {
      document.addEventListener('click', () => {
        this.showUserPopTip = false
      })
      stickybits('#layout-header-bar')
      this.initInfoProfile()
    },
    computed: {
      rotateIcon() {
        return ['menu-icon', this.isCollapsed ? 'rotate-icon' : '']
      }
    },
    methods: {
      collapsedSider() {
        this.$refs.side1.toggleCollapse()
      },
      logout() {
        this.$cookies.remove('token')
        this.$store.commit('setToken', null)
        this.$router.push({
          path: '/login'
        })
      },
      //获取用户信息
      initInfoProfile() {
        // this.$axios({
        //   url: apiRoot + '/user/info/16122131',
        //   headers: {
        //     Authorization: 'eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJzdWIiOiIxNjEyMjEzMSIsImV4cCI6MTg1NjM3MDkzMSwidXNlck5hbWUiOiLpg63lrZ_nhLYiLCJpYXQiOjE1NDUzMzA5MzF9.xmeHnjdFMj6sTDl9qJoJnRwUu-I1iUX2VXznQal9DL6kAw8CyGWoKsNDgIAejqPriOksy9Faee96tZkCeZ5W5w'
        //   }
        // }).then(res => {
        //   return res.data.data
        // }).catch(() => {
        //   return null
        // }).then(data => {
        //   if (data && data instanceof Array) {
        //     this.profile = data[0]
        //   } else if (data) {
        //     this.profile = data
        //   }
        // })
      }
    }
  }
</script>

<style lang="scss" scoped>
  @import "./users";
</style>

<style lang="scss">
  .layout-header-title > * {
    vertical-align: middle;
  }

  .layout-header-title > i {
    font-size: 30px;
  }

  .ivu-layout-sider-collapsed .layout-logo-left {
    font-size: 9px;
  }
</style>
