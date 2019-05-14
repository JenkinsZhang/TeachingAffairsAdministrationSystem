<template>
  <div class="wrapper">
    <div class="operation" style="width:600px">
      <span class="label">课程号：</span><Input enter-button  style="width: 160px"/>
      <span class="label">课程名：</span><Input enter-button  style="width: 160px"/>
    </div>
    <div class="operation" style="width:600px">
      <span class="label">教师名：</span><Input enter-button  style="width: 160px"/>
      <span class="label">教师号：</span><Input enter-button  style="width: 160px"/>
    </div>
    <div class="operation" style="width:600px">
      <span class="label">上课时间：</span><Input enter-button  style="width: 160px"/>
      <span class="label">学分数：</span><Input enter-button  style="width: 160px"/>
    </div>
    <Table class="operation" stripe border :columns="columns" :data="data1" size="large"></Table>
  </div>
</template>

<script>
  export default {
    name: 'courseQuery',
    data() {
      return {
        columns: [{
          'title': '课程号',
          'key': 'kh',
          'align': 'center'
        }, {
          'title': '课程名',
          'key': 'km',
          'align': 'center'
        }, {
          'title': '教师号',
          'key': 'gh',
          'align': 'center'
        }, {
          'title': '教师名',
          'key': 'xm',
          'align': 'center'
        }, {
          'title': '学分',
          'key': 'xf',
          'align': 'center'
        }, {
          'title': '上课时间',
          'key': 'sksj',
          'align': 'center'
        }, {
          'title': '操作',
          'key': 'action',
          'fixed': 'right',
          'width': 100,
          'align': 'center',
          render(h, params) {
            return h('ButtonGroup', {
              props: {
                size: 'small'
              }
            }, [
              h('Button', {
                props: {
                  type: 'success',
                  icon: 'md-close'
                },
                on: {
                  click: () => {
                    console.log('点了退课')
                  }
                }
              }, '选课')
            ])
          }
        }],
        data1: []
      }
    },
    methods: {
      handleSearch(keyword) {
        keyword = keyword.trim()
        console.log(keyword)
      }
    }
  }
</script>

<style scoped>

</style>
