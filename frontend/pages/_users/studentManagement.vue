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
      title="新增学生"
      @on-ok="handleAddStudent"
      :loading="modalLoading"
      :mask-closable="false"
    >
      <!--
          id,
          name,
          gender,
          birthday,
          birthplace,
          phone
          dname,
          grade,
      -->
      <Form
        :model="form"
        label-position="right"
        :rules="ruleValidate"
        :label-width="80"
        style="padding:20px 30px 20px 15px"
        ref="form"
      >
        <FormItem label="学号" prop="id">
          <Input v-model="form.id"></Input>
        </FormItem>
        <FormItem label="姓名" prop="name">
          <Input v-model="form.name"></Input>
        </FormItem>
        <FormItem label="学院" prop="dname">
          <Select v-model="form.dname" placeholder="选择所属学院">
            <Option :value="dep" v-for="dep in deps" :key="dep">{{dep}}</Option>
          </Select>
        </FormItem>
        <FormItem label="出生日期" prop="birthday">
          <DatePicker type="date" placeholder="选择出生日期" v-model="form.birthday"></DatePicker>
        </FormItem>
        <FormItem label="性别" prop="gender">
          <RadioGroup v-model="form.gender">
            <Radio label="男">男</Radio>
            <Radio label="女">女</Radio>
          </RadioGroup>
        </FormItem>
        <FormItem label="籍贯" prop="birthplace">
          <Input v-model="form.birthplace"></Input>
        </FormItem>
        <FormItem label="手机号码" prop="phone">
          <Input v-model="form.phone" number></Input>
        </FormItem>
      </Form>
    </Modal>
  </div>
</template>

<script>
  export default {
    name: 'studentManagement',
    async asyncData({ app }) {
      const data1 = []
      await app.$axios({
        url: '/admin/studentManagement'
      }).then((res) => {
        const {
          birthday,
          birthplace,
          did,
          dname,
          gender,
          id,
          name,
          phone
        } = res.data
        if (!id) {
          return
        }
        for (let i = 0; i < id.length; i++) {
          data1.push({
            birthday: birthday[i],
            birthplace: birthplace[i],
            did: did[i],
            dname: dname[i],
            gender: gender[i],
            id: id[i],
            name: name[i],
            phone: phone[i]
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
        modalLoading: true,
        deps: ['计算机学院'],
        ruleValidate: {
          // id,
          // name,
          // gender,
          // birthday,
          // birthplace,
          // phone,
          // dname
          id: [
            { required: true, message: '学号不能为空', trigger: 'blur' }
          ],
          name: [
            { required: true, message: '姓名不能为空', trigger: 'blur' }
          ],
          birthday: [
            { required: true, type: 'date', message: '出生日期不能为空，且必须符合格式规范', trigger: 'change' }
          ],
          gender: [
            { required: true, message: '性别必选', trigger: 'change' }
          ],
          phone: [
            { required: true, type: 'integer', min: 1, message: '电话不能为空，且必须为数字', trigger: 'change' }
          ],
          birthplace: [
            { required: true, message: '籍贯不能为空', trigger: 'change' }
          ],
          dname: [
            { required: true, message: '学院不能为空', trigger: 'change' }
          ]
        },
        form: {
          birthday: '',
          birthplace: '',
          did: '',
          dname: '',
          gender: '',
          grade: '',
          id: '',
          name: '',
          phone: ''
        },
        showModal: false,
        columns: [{
          'title': '学号',
          'key': 'id',
          'align': 'center',
          'sortable': true
        }, {
          'title': '姓名',
          'key': 'name',
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
          'title': '籍贯',
          'key': 'birthplace',
          'align': 'center'
        }, {
          'title': '手机号码',
          'key': 'phone',
          'align': 'center'
        }, {
          'title': '院系',
          'key': 'dname',
          'align': 'center'
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
                    that.thisRow = params.row
                    that.isAdding = false
                    that.form = Object.assign({}, params.row)
                    that.form.phone = parseInt(that.form.phone)
                    that.form.birthday = that.$dayjs(that.form.birthday, 'YYYYMMDD').toDate()
                    that.showModal = true
                    console.log(that.form)
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
                      content: `确定要删除学生 ${params.row.name} 吗？`,
                      loading: true,
                      onOk: () => {
                        that.$axios({
                          url: '/admin/studentManagement',
                          method: 'post',
                          data: {
                            id: params.row.id,
                            op: 'delete'
                          }
                        }).then(() => {
                          that.$Modal.remove()
                          that.$Message.info('删除成功')
                          that.data1.splice(params.index, 1)
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
        isAdding: false,
        thisRow: null
      }
    },
    methods: {
      handleClickAddButton() {
        if (!this.isAdding) {
          this.$refs.form.resetFields()
          this.isAdding = true
        }
        this.showModal = true
        this.$axios({
          url: '/getDepartment'
        }).then((res) => {
          this.deps = res.data.dname
        })
      },
      async handleAddStudent() {
        const that = this
        if (!await this.$refs.form.validate()) {
          this.modalLoading = false
          this.$nextTick(() => {
            this.modalLoading = true
          })
          return
        }
        let {
          id,
          name,
          gender,
          birthday,
          birthplace,
          phone,
          dname
        } = this.form
        birthday = this.$dayjs(birthday).format('YYYYMMDD')
        this.$axios({
          url: '/admin/studentManagement',
          method: 'post',
          data: {
            id,
            name,
            gender,
            birthday,
            birthplace,
            phone: phone.toString(),
            dname,
            'op': this.isAdding ? 'add' : 'modify'
          }
        }).then((res) => {
          if (res.data.message === 'ok') {
            this.$Message.info(`${this.isAdding ? '新增' : '修改'}成功`)
            this.showModal = false
            this.form.birthday = birthday
            if (!this.isAdding) {
              for (let k in this.thisRow) {
                if (k in this.form) {
                  this.thisRow[k] = this.form[k]
                }
              }
            } else {
              this.data1.push(Object.assign({}, this.form))
            }
            this.$refs.form.resetFields()
          } else {
            this.$Message.warning(res.data.message)
            this.modalLoading = false
            this.$nextTick(() => {
              that.modalLoading = true
            })
          }
        })
      }
    }
  }
</script>

<style scoped>

</style>
