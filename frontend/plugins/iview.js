import Vue from 'vue'
import iView, { Spin } from 'iview'

export default ({}, inject) => {
  Vue.use(iView)
  inject('Spin', Spin)
}
