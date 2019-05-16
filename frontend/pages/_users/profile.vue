<template>
  <Collapse v-model="whichCollapse" accordion>
    <Panel name="1">
      我的信息
      <p slot="content">
        <Tag color="blue">{{idLabel}}</Tag>
        {{profile.id||profile.tid}}
      </p>
      <p slot="content">
        <Tag color="orange">姓　名</Tag>
        {{profile.name||profile.tname}}
      </p>
      <p slot="content">
        <Tag color="green">性　别</Tag>
        {{profile.gender}}
      </p>
      <p slot="content">
        <Tag color="cyan">院　系</Tag>
        {{profile.dname}}
      </p>
      <p slot="content" v-if="profile.grade">
        <Tag color="geekblue">年　级</Tag>
        {{profile.grade}}
      </p>
      <p slot="content" v-if="profile.birthplace">
        <Tag color="magenta">籍　贯</Tag>
        {{profile.birthplace}}
      </p>
      <p slot="content" v-if="profile.phone">
        <Tag color="purple">手机号</Tag>
        {{profile.phone}}
      </p>
    </Panel>
  </Collapse>
</template>

<script>
  export default {
    name: 'profile',
    async asyncData({ params, app }) {
      let idLabel = ''
      let ret = {}
      if (params.users === 'student') {
        idLabel = '学　号'
      } else {
        idLabel = '工　号'
      }
      ret.idLabel = idLabel
      await app.$axios({
        url: apiRoot + `/${params.users}/profile`
      }).then((res) => {
        console.log(res.data)
        ret.profile = { ...res.data }
      })
      return ret
    },
    data: () => ({
      whichCollapse: '1',
      idLabel: '学　号',
      profile: {
        name: '',
        id: '',
        dep: '',
        type: ''
      }
    })
  }
</script>

<style scoped>

</style>
