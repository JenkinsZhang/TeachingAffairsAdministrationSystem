export default function ({ store, redirect }) {
  // 只有用户才能访问，如果用户未经过身份验证，则重定向到登录
  if (!store.state.token) {
    return redirect('/login')
  }
}
