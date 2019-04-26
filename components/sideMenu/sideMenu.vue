<template>
  <div class="menu-list" id="menu-list" ref="menuList">
    <CustomMenu :active-name="MenuActiveName" theme="custom" width="auto" :class="menuitemClasses"
                @on-select="handleSelectItem">
      <MenuGroup
        v-for="group of menu"
        :key="group.groupName"
        :title="group.groupName"
      >
        <MenuItem
          v-for="item of group.items"
          :key="item.name"
          :name="item.name"
        >
          <i :class="['icon','iconfont',item.icon]"></i>
          <span>{{item.title}}</span>
        </MenuItem>
      </MenuGroup>
    </CustomMenu>
    <div id="menu-more">
      <Icon type="ios-arrow-down" size="30" color="#d5d5d5" v-show="showMenuMore"></Icon>
    </div>
  </div>
</template>

<script>
  import CustomMenu from './customMenu/customMenu'
  import stickybits from 'stickybits'


  export default {
    name: 'SideMenu',
    components: { CustomMenu },
    props: {
      collapse: Boolean,
      menu: Array
    },
    data: () => ({
      isCollapsed: false,
      showMenuMore: false,
      MenuActiveName: null
    }),
    computed: {
      menuitemClasses() {
        return [
          'menu-item',
          this.isCollapsed ? 'collapsed-menu' : ''
        ]
      }
    },
    watch: {
      collapse(val) {
        this.isCollapsed = val
      },
      '$route'() {
        this.$nextTick(() => {
          this.initMenuActive()
        })
      }
    },
    mounted() {
      //将menu-more设为sticky
      stickybits('#menu-more', {
        verticalPosition: 'bottom',
        useStickyClasses: true,
        scrollEl: '#menu-list'
      })
      //监听滚动事件，用以控制menu-more的显示
      this.$refs.menuList.addEventListener('scroll', this.menuListScrollHandler)
      this.$nextTick(() => {
        //初始化menu-more的显示
        this.menuListScrollHandler(this.$refs.menuList)
        //初始化menu现选项
        this.initMenuActive()
      })
    },
    methods: {
      menuListScrollHandler(event) {
        const el = event.target || event
        this.showMenuMore = el.scrollTop + el.clientHeight <= el.scrollHeight - 30
        //  30是menu-more的高度
      },
      initMenuActive(activeName) {
        this.MenuActiveName = activeName || this.$route.matched[this.$route.matched.length - 1].components.default.name
        setTimeout(() => {
          const e = document.querySelector('.ivu-menu-item-selected')
          e && this.$emit('input', e.innerHTML)
        }, 20)
      },
      handleSelectItem(name) {
        const that = this
        this.$nextTick(() => {
          that.$router.push(name)
        })
      }
    }
  }
</script>

<style lang="scss" scoped>
  @import "./sideMenu";
</style>

<style lang="scss">
  .ivu-menu-item-group-title {
    $--height: 25px;
    color: #595f69 !important;
    height: $--height !important;
    line-height: $--height !important;
    margin-top: 10px;
    padding: 0 0 0 15px !important;
    font-size: 13px !important;
  }

  .collapsed-menu .ivu-menu-item-group-title {
    text-align: center;
    padding: 0 !important;
  }

  .ivu-layout-sider-children {
    /*height: fit-content !important;*/
  }
</style>
