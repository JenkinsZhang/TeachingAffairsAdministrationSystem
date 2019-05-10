import { getUserInfoFromToken } from 'assets/js/tokenTools'

export default function({ store, redirect }) {
  // 只有游客才能访问，如果用户通过身份验证，则重定向到主页
  if (store.state.token) {
    const info = getUserInfoFromToken(store.state.token)
    return redirect('/' + info.identity)
  }
}
