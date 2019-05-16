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
      @on-ok="handleAddTeacher"
      :loading="modalLoading"
      :mask-closable="false"
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
        :rules="ruleValidate"
        :label-width="80"
        style="padding:20px 30px 20px 15px"
        ref="form"
      >
        <FormItem label="工号" prop="tid">
          <Input v-model="form.tid" :disabled="!isAdding"></Input>
        </FormItem>
        <FormItem label="姓名" prop="tname">
          <Input v-model="form.tname"></Input>
        </FormItem>
        <FormItem label="学院" prop="dname">
          <Select v-model="form.dname" placeholder="选择所属学院">
            <Option :value="dep" v-for="dep in deps" :key="dep">{{dep}}</Option>
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
          <DatePicker type="date" placeholder="选择出生日期" v-model="form.birthday"></DatePicker>
        </FormItem>
        <FormItem label="性别" prop="gender">
          <RadioGroup v-model="form.gender">
            <Radio label="男">男</Radio>
            <Radio label="女">女</Radio>
          </RadioGroup>
        </FormItem>
        <FormItem label="薪资" prop="wage">
          <Input v-model="form.wage" number></Input>
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
        modalLoading: true,
        deps: ['计算机学院'],
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
                    that.thisRow = params.row
                    that.isAdding = false
                    that.form = Object.assign({}, params.row)
                    that.form.wage = parseInt(that.form.wage)
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
        ruleValidate: {
          tid: [
            { required: true, message: '工号不能为空', trigger: 'blur' }
          ],
          tname: [
            { required: true, message: '姓名不能为空', trigger: 'blur' }
          ],
          birthday: [
            { required: true, type: 'date', message: '出生日期不能为空，且必须符合格式规范', trigger: 'change' }
          ],
          gender: [
            { required: true, message: '性别必选', trigger: 'change' }
          ],
          wage: [
            { required: true, type: 'integer', min: 1, message: '薪资不能为空，且必须为正数', trigger: 'change' }
          ],
          education: [
            { required: true, message: '学历不能为空', trigger: 'change' }
          ],
          dname: [
            { required: true, message: '学院不能为空', trigger: 'change' }
          ]
        },
        showModal: false,
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
          url: apiRoot + '/getDepartment'
        }).then((res) => {
          this.deps = res.data.dname
        })
      },
      async handleAddTeacher() {
        const that = this
        if (!await this.$refs.form.validate()) {
          this.modalLoading = false
          this.$nextTick(() => {
            this.modalLoading = true
          })
          return
        }
        let {
          tid,
          tname,
          gender,
          birthday,
          wage,
          education,
          dname
        } = this.form
        birthday = this.$dayjs(birthday).format('YYYYMMDD')
        this.$axios({
          url: apiRoot + '/admin/teacherManagement',
          method: 'post',
          data: {
            tid,
            tname,
            gender,
            birthday,
            wage: wage.toString(),
            education,
            dname,
            'op': this.isAdding ? 'add' : 'modify'
          }
        }).then((res) => {
          console.log('handleAddTeacher', res.data)
          if (res.data.message === 'ok') {
            this.$Message.info(`${this.isAdding ? '新增' : '修改'}成功`)
            this.showModal = false
            this.form.birthday = birthday
            if (!this.isAdding) {
              for (let k in this.thisRow) {
                if (k in this.form) {
                  this.thisRow[k] = this.form[k]
                  console.log(k, this.form[k])
                }
              }
            } else {
              this.data1.push(Object.assign({}, this.form))
            }
            this.$refs.form.resetFields()
          } else {
            this.$Message.info(res.data.message)
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
