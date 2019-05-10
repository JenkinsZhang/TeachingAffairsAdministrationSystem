import Cookies from 'js-cookie'
import Vue from 'vue'

export default ({ app }, inject) => {
  inject('cookies', Cookies)
}
