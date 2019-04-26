<template>
  <div class="layout">
    <Layout style="height:100%">
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
          >{{sidebarSelectedTitle}}
          </div>
          <div
            class="userBox"
            @click="showUserPopTip = !showUserPopTip;$event.stopPropagation()"
          >
            <div
              class="userBox-avatar"
              :style="{backgroundImage: 'url(https://cdn-daoweb-prod.daocloud.io/objects/avatar/04148988-d6c8-4e8b-9081-ebe1f9e78a2f.jpg)'}"
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
            <li @click="$router.push('/organization')">
              <Icon
                type="md-checkmark"
                size="20"
              ></Icon>
              {{profile.name}}
              <span class="right-align" style="margin-right:10px">{{profile.role}}</span>
            </li>
            <li class="divider"></li>
            <li @click="$router.push('/organization')">
              <Icon
                type="ios-person"
                size="20"
              ></Icon>
              用户中心
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
        <Content :style="{margin: '20px', background: '#fff', minHeight: 'calc(100% - 64px - 40px)', display:'table'}">
          <div style="position: relative">
            <nuxt-child></nuxt-child>
            <!--            <Spin-->
            <!--              size="large"-->
            <!--              fix-->
            <!--              v-if="showLoading"-->
            <!--              style="min-height: calc(100vh - 100px)"-->
            <!--            >加载中...</Spin>-->
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

  const identityMap = {
    'student': '学生',
    'teacher': '教师',
    'admin': '管理员',
    'superAdmin': '系统管理员'
  }

  export default {
    name: 'users',
    components: { SideMenu },
    validate({ params }) {
      console.log(params.users)
      return !!identityMap[params.users]
    },
    data() {
      return {
        isCollapsed: false,
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
      if (process.server) {
        let menu = menuJson[params.users]
        console.log(menu)
        menu.forEach((g) => {
          g.items.forEach((i) => {
            i.name = `${params.users}/${i.name}`
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
      }
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
        localStorage.removeItem('area')
      },
      //获取用户信息
      initInfoProfile() {
        this.$axios({
          url: apiRoot + '/profile'
        }).then(res => {
          return res.data.data
        }).catch(() => {
          return null
        }).then(data => {
          if (data && data instanceof Array) {
            this.profile = data[0]
          } else if (data) {
            this.profile = data
          }
        })
      }
    }
  }
</script>

<style lang="scss" scoped>
  @import "./users";
</style>

<style lang="scss">
  .ivu-layout-content {
    width: calc(100% - 40px);
  }

  .layout-header-title > * {
    vertical-align: middle;
  }

  .layout-header-title > i {
    font-size: 30px;
  }
</style>
