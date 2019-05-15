<template>
  <div class="wrapper">
    <ButtonGroup class="operation">
      <Button type="primary" icon="md-add" size="large" @click="handleClickAddButton">
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
    <Modal
      v-model="showModal"
      title="新增教师"
    >
      <!--
                tid: '',
                tname: '',
                gender: '',
                birthday: '',
                wage: '',
                education: '',
                dname: ''
      -->
      <Form
        :model="form"
        label-position="right"
        :label-width="60"
        style="padding:20px 30px 20px 15px"
        ref="form"
      >
        <FormItem label="工号" prop="tid">
          <Input v-model="form.tid"></Input>
        </FormItem>
        <FormItem label="姓名" prop="tname">
          <Input v-model="form.tname"></Input>
        </FormItem>
        <FormItem label="学院" prop="dname">
          <Select v-model="form.dname" placeholder="选择所属学院">
            <Option value="计算机学院">计算机学院</Option>
            <Option value="通信学院">通信学院</Option>
            <Option value="机自学院">机自学院</Option>
            <Option value="数据是mock的">数据是mock的</Option>
          </Select>
        </FormItem>
        <FormItem label="学历" prop="education">
          <Select v-model="form.education" placeholder="选择学历">
            <Option value="学士">学士</Option>
            <Option value="硕士">硕士</Option>
            <Option value="博士">博士</Option>
          </Select>
        </FormItem>
        <FormItem label="出生日期" prop="birthday">
          <DatePicker type="date" placeholder="选择出生日期" v-model="form.date"></DatePicker>
        </FormItem>
        <FormItem label="性别" prop="gender">
          <RadioGroup v-model="form.gender">
            <Radio label="male">男</Radio>
            <Radio label="female">女</Radio>
          </RadioGroup>
        </FormItem>
        <FormItem label="薪资" prop="wage">
          <Input v-model="form.wage"></Input>
        </FormItem>
      </Form>
    </Modal>
  </div>
</template>

<script>
  export default {
    name: 'teacherManagement',
    async asyncData({ app }) {
      const data1 = []
      await app.$axios({
        url: apiRoot + '/admin/teacherManagement'
      }).then((res) => {
        const {
          tid,
          tname,
          gender,
          birthday,
          wage,
          education,
          dname
        } = res.data
        if (!tid) {
          return
        }
        for (let i = 0; i < tid.length; i++) {
          data1.push({
            tid: tid[i],
            gender: gender[i],
            birthday: birthday[i],
            tname: tname[i],
            wage: wage[i],
            education: education[i],
            dname: dname[i]
          })
        }
      })
      return {
        data1
      }
    },
    data() {
      const that = this
      return {
        columns: [{
          'title': '工号',
          'key': 'tid',
          'align': 'center',
          'sortable': true
        }, {
          'title': '姓名',
          'key': 'tname',
          'align': 'center'
        }, {
          'title': '性别',
          'key': 'gender',
          'align': 'center'
        }, {
          'title': '出生日期',
          'key': 'birthday',
          'align': 'center'
        }, {
          'title': '学历',
          'key': 'education',
          'align': 'center'
        }, {
          'title': '基本工资',
          'key': 'wage',
          'align': 'center'
        }, {
          'title': '院系',
          'key': 'dname',
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
                    that.$Modal.confirm({
                      title: '确认',
                      content: `确定要删除教师 ${params.row.tname} 吗？`,
                      loading: true,
                      onOk: () => {
                        that.$axios({
                          url: apiRoot + '/admin/teacherManagement',
                          method: 'post',
                          data: {
                            tid: params.row.tid,
                            op: 'delete'
                          }
                        }).then(() => {
                          that.$Modal.remove()
                          that.$Message.info('删除成功')
                          that.data1.splice(params.index, 1)
                          that.$refs.form.resetFields()
                        })
                      },
                      onCancel: () => {
                      }
                    })
                  }
                }
              }, '删除')
            ])
          }
        }],
        data1: [],
        form: {
          tid: '',
          tname: '',
          gender: '',
          birthday: '',
          wage: '',
          education: '',
          dname: ''
        },
        showModal: false
      }
    },
    methods: {
      handleClickAddButton() {
        this.showModal = true
      },
      handleAddTeacher() {
        let {
          tid,
          tname,
          gender,
          birthday,
          wage,
          education,
          dname
        } = this.form
        this.$axios({
          url: apiRoot + '/admin/teacherManagement',
          method: 'post',
          data: {
            tid,
            tname,
            gender,
            birthday,
            wage,
            education,
            dname,
            'op': 'add'
          }
        }).then((res) => {

        })
      }
    }
  }
</script>

<style scoped>

</style>
