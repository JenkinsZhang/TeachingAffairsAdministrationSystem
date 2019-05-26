<template>
  <div class="wrapper">
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
    <Table class="operation" stripe border :columns="columns" :data="data1" size="large" placeholder="请按条件搜索"></Table>
  </div>
</template>

<script>
  const tiptopMap = {
    '一': 1,
    '二': 2,
    '三': 3,
    '四': 4,
    '五': 5,
    '六': 6,
    '七': 7,
    '天': 7,
    '日': 7
  }
  var that = null
  export default {
    name: 'courseQuery',
    asyncData() {

    },
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
                    //判选课开放
                    that.$axios({
                      url: '/admin/courseManagement',
                      method: 'post',
                      data: {
                        op: 'queryOpen'
                      }
                    }).then((res) => {
                      if (res.data.open === 'open') {
                        //判cid重复, 判时间冲突
                        that.$axios({
                          url: '/student/courseCalendar',
                          method: 'post',
                          data: {
                            op: 'query'
                          }
                        }).then((res) => {
                          const data = res.data
                          if (!data.cid) {
                            data.cid = []
                            data.classTime = []
                          }
                          //判cid重复
                          for (let i = 0; i < data.cid.length; i++) {
                            if (params.row.kh === data.cid[i]) {
                              return that.$Notice.warning({
                                title: '提示',
                                desc: `您本学期已选[${data.cid[i]}]《${data.cname[i]}》(${data.tname[i]})，不能重复选择相同课程号的课程`
                              })
                            }
                          }

                          function forEachClassUnit({ classTime }, func) {
                            //每门课
                            for (let i = 0; i < classTime.length; i++) {
                              let ct = classTime[i].split(' ')
                              //每个不同的时间段
                              if (!ct) {
                                ct = [classTime[i]]
                              }
                              for (let j = 0; j < ct.length; j++) {
                                const g = ct[j].match(/星期[一二三四五六七日天](\d+)-(\d+)/)
                                console.log(g)
                                if (!(g && g.length > 0 && g[1] >= 1 && g[1] <= 13 && g[2] >= 1 && g[2] <= 13 && parseInt(g[1]) <= g[2])) {
                                  throw new Error('上课时间格式不规范，示例：星期一11-13 星期五3-4')
                                }
                                const column = tiptopMap[ct[j].substring(2)[0]]
                                const arr = ct[j].substring(3).split('-')
                                let rowBegin = parseInt(arr[0]),
                                  rowEnd = parseInt(arr[1])
                                for (let k = rowBegin; k <= rowEnd; k++) {
                                  const stop = func(k, column, i, j)//课表中的行号（1开始），课表中的列号（2开始），课程下标，时间段下标
                                  if (stop) {
                                    return
                                  }
                                }
                              }
                            }
                          }

                          //初始化课表，把0空出来
                          const matrix = new Array(14)
                          for (let i = 0; i < 14; i++) {
                            matrix[i] = new Array(8)
                          }
                          try {
                            forEachClassUnit({ classTime: data.classTime }, (i, j) => {
                              matrix[i][j] = 1
                            })
                            //查看上课时间是否重复
                            let invalid = false
                            let cIndex = -1
                            forEachClassUnit({ classTime: [params.row.sksj] }, (i, j, index) => {
                              if (matrix[i][j]) {
                                invalid = true
                                cIndex = index
                                return true
                              }
                              return false
                            })
                            if (invalid) {
                              return that.$Notice.warning({
                                title: '提示',
                                desc: `上课时间与所选的《${data.cname[cIndex]}》[${data.classTime[cIndex]}]冲突`
                              })
                            }
                            //发送提示
                            that.$Modal.confirm({
                              title: '确认',
                              content: `确定要选课《${params.row.km}》吗？`,
                              loading: true,
                              onOk: () => {
                                that.$axios({
                                  url: '/student/courseQuery',
                                  method: 'post',
                                  data: {
                                    cid: params.row.kh,
                                    classTime: params.row.sksj,
                                    tid: params.row.gh,
                                    term: that.selected,
                                    op: 'select'
                                  }
                                }).then((res) => {
                                  if (res.data.message === 'ok') {
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
                          } catch (e) {
                            console.log(e)
                            that.$Notice.warning({
                              title: '提示',
                              desc: e.toString()
                            })
                          }
                        }).catch((err) => {
                          console.log(err)
                          that.$Notice.warning({
                            title: '提示',
                            desc: err.toString()
                          })
                        })
                      } else {
                        that.$Notice.warning({
                          title: '提示',
                          desc: '当前时间选课尚未开放'
                        })
                      }
                    }).catch((err) => {
                      console.log(err)
                      that.$Notice.warning({
                        title: '提示',
                        desc: err.toString()
                      })
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
          url: '/student/courseQuery',
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
