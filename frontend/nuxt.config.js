const webpack = require('webpack')
const pkg = require('./package')


module.exports = {
  server: {
    port: 3000,
    host: '0.0.0.0'
  },
  mode: 'universal',
  /*
  ** Headers of the page
  */
  head: {
    titleTemplate: chunk => chunk ? `${chunk} | 教学事务管理系统` : '教学事务管理系统',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: pkg.description }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
      { rel: 'stylesheet', href: '//at.alicdn.com/t/font_1164860_8vr2y50xuyd.css' }
    ]
  },

  /*
  ** Customize the progress-bar color
  */
  loading: {
    color: '#359bff',
    height: '3px'
  },

  /*
  ** Global CSS
  */
  css: [
    'iview/dist/styles/iview.css',
    'normalize.css',
    { src: '~assets/style/reset.scss', lang: 'scss' },
    { src: '~assets/style/main.scss', lang: 'scss' }
  ],

  /*
  ** Plugins to load before mounting the App
  */
  plugins: [
    '@/plugins/iview',
    '@/plugins/vcharts.client',
    '@/plugins/cookies.client'
  ],

  /*
  ** Nuxt.js modules
  */
  modules: [
    // Doc: https://axios.nuxtjs.org/usage
    '@nuxtjs/axios',
    '@nuxtjs/pwa'
  ],
  /*
  ** Axios module configuration
  */
  axios: {
    // See https://github.com/nuxt-community/axios-module#options
    baseURL: undefined
  },

  /*
  ** Build configuration
  */
  build: {
    /*
    ** You can extend webpack config here
    */
    extend(config, ctx) {
      if (ctx.isDev) {
        config.plugins.push(new webpack.DefinePlugin({
          'apiRoot': '\'/api\''
        }))
      } else {
        config.plugins.push(new webpack.DefinePlugin({
          'apiRoot': '\'/api\''
        }))
      }
    },
    // babel: {
    // babelrc:true,
    // 'plugins': [['import', {
    //   'libraryName': 'iview',
    //   'libraryDirectory': 'src/components'
    // }]]
    // },
    postcss: {
      // 添加插件名称作为键，参数作为值
      // 使用npm或yarn安装它们
      plugins: {
        // 通过传递 false 来禁用插件
        'postcss-url': false,
        'postcss-nested': {},
        'postcss-responsive-type': {},
        'postcss-hexrgba': {}
      },
      preset: {
        // 更改postcss-preset-env 设置
        autoprefixer: {
          grid: true
        }
      }
    }
  }
}
