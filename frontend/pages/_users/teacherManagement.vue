<template>
  <div class="wrapper">
    <ButtonGroup class="operation">
      <Button type="primary" icon="md-add" size="large">
        新增
      </Button>
    </ButtonGroup>
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
    name: 'teacherManagement',
    data() {
      const that = this
      return {
        columns: [{
          'title': '工号',
          'key': 'gh',
          'align': 'center',
          'sortable': true
        }, {
          'title': '姓名',
          'key': 'xm',
          'align': 'center'
        }, {
          'title': '性别',
          'key': 'xb',
          'align': 'center'
        }, {
          'title': '出生日期',
          'key': 'csrq',
          'align': 'center'
        }, {
          'title': '学历',
          'key': 'xl',
          'align': 'center'
        }, {
          'title': '基本工资',
          'key': 'jbgz',
          'align': 'center'
        }, {
          'title': '院系',
          'key': 'dep',
          'align': 'center',
          'sortable': true
        }, {
          'title': '操作',
          'key': 'action',
          'fixed': 'right',
          'width': 200,
          'align': 'center',
          render(h, params) {
            return h('ButtonGroup', {
              props: {
                size: 'default',
                shape: 'circle'
              }
            }, [
              h('Button', {
                props: {
                  type: 'primary',
                  icon: 'md-create'
                },
                on: {
                  click: () => {
                    that.showModal = true
                    that.thisRow = params.row
                  }
                }
              }, '修改'),
              h('Button', {
                props: {
                  type: 'warning',
                  icon: 'md-trash'
                },
                on: {
                  click: () => {
                    that.showModal = true
                    that.thisRow = params.row
                  }
                }
              }, '删除')
            ])
          }
        }],
        data1: [{
          gh: '10000663',
          xb: '男',
          xm: '陈迪茂',
          csrq: '1973-03-06',
          xl: '副教授',
          jbgz: '3567.00',
          dep: '计算机工程与科学学院'
        }]
      }
    }
  }
</script>

<style scoped>

</style>
