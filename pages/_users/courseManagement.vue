<template>
  <div class="wrapper">
    <Form class="operation">
      <FormItem label="学期">
        <Row type="flex" justify="space-between">
          <Select
            v-model="selectedClassId"
            style="width:200px"
            placeholder="请选择学期"
          >
            <Option v-for="term of terms" :value="term.name" :key="term.name">{{term.name}}</Option>
          </Select>
          <ButtonGroup>
            <Button type="primary" icon="md-add" size="large">
              新增
            </Button>
          </ButtonGroup>
        </Row>
      </FormItem>
      <FormItem>
        <Table
          stripe
          border
          :columns="columns"
          :data="data1"
        ></Table>
      </FormItem>
    </Form>
  </div>
</template>

<script>
  import mock from '~/assets/js/scoreManagementMock'

  export default {
    name: 'teacherManagement',
    mixins: [mock],
    data() {
      const that = this
      return {
        columns: [{
          'title': '课程号',
          'key': 'kh',
          'align': 'center',
          'sortable': true
        }, {
          'title': '课程名',
          'key': 'km',
          'align': 'center'
        }, {
          'title': '学分',
          'key': 'xf',
          'align': 'center'
        }, {
          'title': '学时',
          'key': 'xs',
          'align': 'center'
        }, {
          'title': '开课院系',
          'key': 'dep',
          'align': 'center'
        }, {
          'title': '开课数',
          'key': 'kks',
          'align': 'center',
          'sortable': true
        }, {
          'title': '操作',
          'key': 'action',
          'fixed': 'right',
          'width': 300,
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
                  type: 'info',
                  icon: 'md-apps'
                },
                on: {
                  click: () => {
                    that.showModal = true
                    that.thisRow = params.row
                  }
                }
              }, '开课安排'),
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
          kh: '08305001',
          km: '离散数学',
          xf: '4',
          xs: '40',
          dep: '计算机工程与科学学院',
          kks: 3
        }],
        selectedClassId: ''
      }
    }
  }
</script>

<style scoped>

</style>
