<template>
  <div class="wrapper">
    <Row>
      <Col span="17">
        <div class="operation" style="width:600px">
          <span class="label">课程号：</span><Input enter-button style="width: 160px" v-model="cid" @on-enter="handleSearch"/>
          <span class="label">课程名：</span><Input enter-button style="width: 160px" v-model="cname"
                                                @on-enter="handleSearch"/>
        </div>
        <div class="operation" style="width:600px">
          <span class="label">教师名：</span><Input enter-button style="width: 160px" v-model="tname"
                                                @on-enter="handleSearch"/>
          <span class="label">教师号：</span><Input enter-button style="width: 160px" v-model="tid" @on-enter="handleSearch"/>
        </div>
      </Col>
      <Col span="7">
        <Form class="operation">
          <FormItem label="学期">
            <Select
              v-model="selectedClassId"
              style="width:200px"
              placeholder="请选择学期"
              @on-change="handleSelectChange"
            >
              <Option v-for="term of terms" :value="term" :key="term">{{term}}</Option>
            </Select>
          </FormItem>
        </Form>
      </Col>
    </Row>
    <Table class="operation" stripe border :columns="columns" :data="data1" size="large"></Table>
  </div>
</template>

<script>
  var that = null
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
                    that.$Modal.confirm({
                      title: '确认',
                      content: `确定要选课《${params.row.km}》吗？`,
                      loading: true,
                      onOk: () => {
                        that.$axios({
                          url: apiRoot + '/student/courseQuery',
                          method: 'post',
                          data: {
                            cid: params.row.kh,
                            classTime: params.row.sksj,
                            tid: params.row.gh,
                            term: that.selectedClassId,
                            op: 'select'
                          }
                        }).then((res) => {
                          if (res.data.message === 'success') {
                            that.$Message.info('选课成功')
                          } else {
                            that.$Message.warning(res.data.message)
                          }
                          that.$Modal.remove()
                        })
                      },
                      onCancel: () => {
                      }
                    })
                  }
                }
              }, '选课')
            ])
          }
        }],
        data1: [],
        cid: '',
        cname: '',
        tid: '',
        tname: ''
      }
    },
    mounted() {
      that = this
    },
    methods: {
      handleSearch() {
        this.data1.length = 0
        this.$axios({
          url: apiRoot + '/student/courseQuery',
          method: 'post',
          data: {
            cid: this.cid.trim(),
            cname: this.cname.trim(),
            tid: this.tid.trim(),
            tname: this.tname.trim(),
            op: 'query'
          }
        }).then((res) => {
          console.log(res.data)
          const { cid, cname, credit, classTime, tid, tname } = res.data
          if (!cid) {
            return
          }
          for (let i = 0; i < cid.length; i++) {
            this.data1.push({
              kh: cid[i],
              km: cname[i],
              gh: tid[i],
              xm: tname[i],
              xf: credit[i],
              sksj: classTime[i]
            })
          }
        })
      }
    }
  }
</script>

<style scoped>

</style>
