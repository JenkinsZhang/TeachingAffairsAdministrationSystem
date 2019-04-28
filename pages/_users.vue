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
              :style="{backgroundImage: 'url(http://ww1.sinaimg.cn/thumbnail/006P5HMAly1g2icle6yjlj30f50ecguy.jpg)'}"
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
              <span class="right-align" style="margin-right:10px">{{profile.role}}</span>
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

  export default {
    name: 'users',
    components: { SideMenu },
    validate({ params }) {
      console.log('validate', process.server, params.users)
      return !!identityMap[params.users] && !!menuJson[params.users]
    },
    head() {
      return {
        title: this.profile.role + '端'
      }
    },
    data() {
      return {
        isCollapsed: true,
        showUserPopTip: false,
        profile: {
          id: '',
          name: '',
          pass: '',
          group_id: '',
          role: '',
          email: ''
        },
        sidebarSelectedTitle: '',
        menu: []
      }
    },
    asyncData({ params }) {
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
      let profile = {
        id: '10001',
        name: '莫之章',
        pass: '',
        group_id: '',
        role: identityMap[params.users],
        email: ''
      }
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
        this.$router.push({
          path: '/login'
        })
        localStorage.removeItem('token')
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
