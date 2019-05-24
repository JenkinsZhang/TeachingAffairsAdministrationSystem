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
      @on-ok="handleAddTerm"
      :loading="modalLoading"
      :mask-closable="false"
    >
      <Form
        :model="form"
        label-position="right"
        :rules="ruleValidate"
        :label-width="80"
        style="padding:20px 30px 20px 15px"
        ref="form"
      >
        <FormItem label="学期名" prop="term">
          <Input v-model="form.term"></Input>
        </FormItem>
      </Form>
    </Modal>
  </div>
</template>

<script>
  export default {
    name: 'term',
    async asyncData({ app }) {
      const data1 = []
      await app.$axios({
        url: '/admin/termManagement'
      }).then((res) => {
        const { isCurrent, term } = res.data
        if (!term) {
          return
        }
        for (let i = 0; i < term.length; i++) {
          data1.push({
            state: isCurrent[i] === 'yes' ? '当前学期' : '',
            isCurrent: isCurrent[i],
            term: term[i]
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
          'title': '学期名',
          'key': 'term',
          'align': 'center'
        }, {
          'title': '状态',
          'key': 'state',
          'align': 'center'
        }, {
          'title': '操作',
          'key': 'action',
          'fixed': 'right',
          'width': 270,
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
                  type: params.row.isCurrent === 'yes' ? 'warning' : 'primary',
                  icon: 'md-create'
                },
                on: {
                  click: () => {
                    that.$Modal.confirm({
                      title: '确认',
                      content: params.row.isCurrent === 'yes' ? `确定要结束学期 ${params.row.term} 吗？` : `确定要设置学期 ${params.row.term} 为当前学期吗？`,
                      loading: true,
                      onOk: () => {
                        //结束学期
                        const op = params.row.isCurrent === 'yes' ? 'end' : 'set'
                        that.$axios({
                          url: '/admin/termManagement',
                          method: 'post',
                          data: {
                            op,
                            term: params.row.term
                          }
                        }).then((res) => {
                          that.$Modal.remove()
                          if (res.data.message === 'ok') {
                            that.$Message.info('操作成功')
                            if (op === 'end') {
                              params.row.state = ''
                              params.row.isCurrent = 'no'
                            } else {
                              params.row.state = '当前学期'
                              params.row.isCurrent = 'yes'
                            }
                          } else {
                            that.$Message.warning(res.data.message)
                          }
                        })
                      },
                      onCancel: () => {
                      }
                    })
                  }
                }
              }, params.row.isCurrent === 'yes' ? '结束当前学期' : '设为当前学期'),
              h('Button', {
                props: {
                  type: 'error',
                  icon: 'md-trash'
                },
                on: {
                  click: () => {
                    that.$Modal.confirm({
                      title: '确认',
                      content: `确定要删除学期 ${params.row.term} 吗？`,
                      loading: true,
                      onOk: () => {
                        that.$axios({
                          url: '/admin/termManagement',
                          method: 'post',
                          data: {
                            term: params.row.term,
                            op: 'delete'
                          }
                        }).then((res) => {
                          that.$Modal.remove()
                          if (res.data.message === 'ok') {
                            that.$Message.info('删除成功')
                            that.data1.splice(params.index, 1)
                          } else {
                            that.$Message.warning(res.data.message)
                          }
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
          term: ''
        },
        ruleValidate: {
          term: [
            { required: true, message: '学期名不能为空', trigger: 'blur' }
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
          url: '/getDepartment'
        }).then((res) => {
          this.deps = res.data.dname
        })
      },
      async handleAddTerm() {
        const that = this
        if (!await this.$refs.form.validate()) {
          this.modalLoading = false
          this.$nextTick(() => {
            this.modalLoading = true
          })
          return
        }
        let {
          term
        } = this.form
        this.$axios({
          url: '/admin/termManagement',
          method: 'post',
          data: {
            term,
            'op': 'add'
          }
        }).then((res) => {
          console.log('handleAddTerm', res.data)
          if (res.data.message === 'ok') {
            this.$Message.info('新增成功')
            this.showModal = false
            this.data1.push(Object.assign({ state: '', isCurrent: 'no' }, this.form))
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
