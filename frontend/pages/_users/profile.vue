<template>
  <Collapse v-model="whichCollapse" accordion>
    <Panel name="1">
      我的信息
      <p slot="content">
        <Tag color="blue">{{idLabel}}</Tag>
        {{profile.id}}
      </p>
      <p slot="content">
        <Tag color="green">类　型</Tag>
        {{profile.type}}
      </p>
      <p slot="content">
        <Tag color="orange">姓　名</Tag>
        {{profile.name}}
      </p>
      <p slot="content">
        <Tag color="cyan">院　系</Tag>
        {{profile.dep}}
      </p>
    </Panel>
  </Collapse>
</template>

<script>

  export default {
    name: 'profile',
    asyncData({ params }) {
      if (params.users === 'student') {
        return { idLabel: '学　号' }
      } else {
        return { idLabel: '工　号' }
      }
    },
    data: () => ({
      whichCollapse: '1',
      idLabel: '学　号',
      profile: {
        name: '莫之章',
        id: '16121663',
        dep: '计算机工程与科学学院',
        type: '本科生'
      }
    })
  }
</script>

<style scoped>

</style>
