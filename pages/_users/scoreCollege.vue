<template>
  <div class="wrapper">
    <Table
      class="operation"
      stripe
      border
      :columns="columns"
      :data="data1"
    ></Table>
  </div>
</template>

<script>
  export default {
    name: 'scoreCollege',
    data() {
      return {
        selectedClassId: '',
        columns: [{
          'title': '学号',
          'key': 'xh',
          'align': 'center'
        },
          {
            'title': '姓名',
            'key': 'xm',
            'align': 'center'
          },
          {
            'title': '平均成绩',
            'key': 'pjcj',
            'align': 'center'
          },
          {
            'title': '院系',
            'key': 'yx',
            'align': 'center'
          },
          {
            'title': '年级总人数',
            'key': 'zs',
            'align': 'center'
          },
          {
            'title': '排名',
            'key': 'xf',
            'align': 'center'
          },
          {
            'title': '百分比',
            'key': 'bfb',
            'align': 'center'
          }],
        data1: []
      }
    }
  }
</script>

<style scoped>

</style>
