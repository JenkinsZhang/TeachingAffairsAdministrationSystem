<template>
  <div class="wrapper">
    <Form class="operation">
      <FormItem label="学期">
        <Select
          v-model="selected"
          style="width:200px"
          placeholder="请选择学期"
          @on-change="handleSelectChange"
        >
          <Option v-for="term of terms" :value="term" :key="term">{{term}}</Option>
        </Select>
      </FormItem>
    </Form>
    <Row class="operation">
      <Col span="12">
        <Table
          class="calendarRaw"
          stripe
          border
          :columns="columns"
          :data="data1"
          size="default"
          ref="calendarRaw"
        ></Table>
      </Col>
      <Col span="12">
        <Table
          class="calendar"
          border
          disabled-hover
          :columns="calendarColumns"
          :data="calendar"
          ref="calendar"
          :max-height="calendarHeight"
        ></Table>
      </Col>
    </Row>
  </div>
</template>

<script>
  import calendarColumns from 'assets/js/calendarColumnsMock'

  let that
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
  const colorArray = ['#d50000', '#E91E63', '#9C27B0', '#673AB7', '#3F51B5', '#03A9F4', '#00BCD4', '#009688', '#4CAF50', '#FF5722', '#795548', '#607D8B']
  export default {
    name: 'courseCalender',
    mixins: [calendarColumns],
    async asyncData({ params, app }) {
      let isStudent = false
      if (params.users === 'student') {
        isStudent = true
      }
      const data1 = []
      let terms = null
      let selected = ''
      await app.$axios({
        url: `/${params.users}/courseCalendar`
      }).then(async (res) => {
        terms = res.data.term.term
        if (!terms) {
          return
        }
        res.data.term.isCurrent.some((x, index) => {
          if (x === 'yes') {
            selected = terms[index]
            return true
          }
          return false
        })
        await app.$axios({
          url: `/${params.users}/courseCalendar`,
          method: 'post',
          data: {
            term: selected
          }
        }).then((res) => {
          const { cid, cname, credit, classTime, tid, tname } = res.data
          if (!cid) {
            return
          }
          for (let i = 0; i < cid.length; i++) {
            data1.push({
              kh: cid[i],
              km: cname[i],
              gh: tid[i],
              xm: tname[i],
              xf: credit[i],
              sksj: classTime[i]
            })
          }
        })
      })
      return {
        isStudent,
        data1,
        terms,
        selected: selected
      }
    },
    data() {
      return {
        columns: [
          {
            'title': '课程号',
            'key': 'kh',
            'align': 'center'
          },
          {
            'title': '课程名',
            'key': 'km',
            'align': 'center'
          },
          {
            'title': '教师号',
            'key': 'gh',
            'align': 'center'
          },
          {
            'title': '教师名',
            'key': 'xm',
            'align': 'center'
          },
          {
            'title': '学分',
            'key': 'xf',
            'align': 'center'
          },
          {
            'title': '上课时间',
            'key': 'sksj',
            'width': 150,
            'align': 'center'
          }
        ],
        data1: [],
        selected: '',
        calendar: [{ time: '1' }, { time: '2' }, { time: '3' }, { time: '4' }, { time: '5' }, { time: '6' }, { time: '7' }, { time: '8' }, { time: '9' }, { time: '10' }, { time: '11' }, { time: '12' }, { time: '13' }],
        nodeMatrix: [],
        isStudent: false,
        calendarHeight: null,
        terms: []
      }
    },
    methods: {
      handleSelectChange(term) {
        this.data1 = []
        this.renderCalendar()
        this.$axios({
          url: `/${this.isStudent ? 'student' : 'teacher'}/courseCalendar`,
          method: 'post',
          data: {
            term
          }
        }).then((res) => {
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
        }).finally(() => {
          this.renderCalendar()
        })
      },
      renderCalendar({
                       raw = this.data1,
                       hover: {
                         kh, sksj
                       } = {}
                     } = {}) {
        //raw是课程数组
        //clear
        let matrix = new Array(13)
        if (this.nodeMatrix.length) {
          for (let i = 0; i < this.nodeMatrix.length; i++) {
            for (let j = 0; j < this.nodeMatrix[j].length; j++) {
              this.nodeMatrix[i][j].firstChild.innerHTML = ''
              this.nodeMatrix[i][j].style = ''
            }
          }
        }
        for (let i = 0; i < 13; i++) {
          matrix[i] = new Array(7)
        }
        //计算，并赋予颜色
        let colorArrayBak = Object.assign([], colorArray)

        function swap(arr, i, j) {
          [arr[i], arr[j]] = [arr[j], arr[i]]
        }

        let cnt = 0
        raw.forEach((obj, index) => {
          const x = Object.assign({}, obj)
          let ct = x.sksj.split(' ')
          //每个不同的时间段
          if (!ct) {
            ct = [x.sksj]
          }
          for (let j = 0; j < ct.length; j++) {
            if (ct[j].indexOf('星期') === 0) {
              ct[j] = ct[j].substring(2)
            }
            const day = tiptopMap[ct[j][0]]
            const arr = ct[j].substring(1).split('-')
            let begin = parseInt(arr[0]),
              end = parseInt(arr[1])
            //给课程初始化一个颜色
            let color
            if ((kh === undefined || sksj === undefined) || (obj.kh === kh && obj.sksj === sksj)) {//没有hover的元素，或hover的就是这个课程
              if (!raw[index]._color) {
                swap(colorArrayBak, cnt, Math.floor(Math.random() * (colorArrayBak.length - cnt) + cnt))
                raw[index]._color = colorArrayBak[cnt++]
              }
              color = raw[index]._color
            } else {
              color = 'rgba(0,0,0,.2)'//灰色
            }
            matrix[begin - 1][day - 1] = {
              ...obj,
              color,
              _length: end - begin + 1
            }
          }
        })
        //开始渲染
        let tr = document.querySelectorAll('.calendar .ivu-table-tbody tr')
        for (let i = 0; i < tr.length; i++) {
          const tds = tr[i].querySelectorAll('td:not(:first-child)')
          this.nodeMatrix.push(tds)
        }
        const rawTbody = document.querySelector('.calendarRaw .ivu-table-tbody')
        //在.calendarRaw上挂特效函数
        rawTbody.onmouseleave = this.mouseLeaveTbody
        tr = rawTbody.querySelectorAll('tr')
        for (let i = 0; i < tr.length; i++) {
          tr[i].onmouseenter = () => (this.mouseEnterTr(i))
        }

        for (let i = 0; i < matrix.length; i++) {
          for (let j = 0; j < matrix[i].length; j++) {
            if (!matrix[i][j]) {
              continue
            }
            this.nodeMatrix[i][j].style.position = 'relative'
            this.nodeMatrix[i][j].firstChild.innerHTML = `
              <div class="node" style="
                height: ${matrix[i][j]._length}00%;
                background-color: ${matrix[i][j].color};
              ">
                <div>
                  <p style="font-weight:bold">${matrix[i][j].kh}</p>
                  <p style="font-weight:bold">${matrix[i][j].km}</p>
                </div>
              </div>
              `
            this.nodeMatrix[i][j].style.color = 'white'
          }
        }
      },
      mouseEnterTr(line) {
        // console.log(line)
        this.renderCalendar({
          hover: {
            kh: this.$refs.calendarRaw.data[line].kh,
            sksj: this.$refs.calendarRaw.data[line].sksj
          }
        })
      },
      mouseLeaveTbody() {
        // console.log('out')
        this.renderCalendar()
      }
    },
    mounted() {
      that = this
      this.calendarHeight = document.body.clientHeight - 66
      if (this.isStudent) {
        this.columns.push({
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
                  type: 'error',
                  icon: 'md-close'
                },
                on: {
                  click: () => {
                    that.$Modal.confirm({
                      title: '确认',
                      content: `确定要退课《${params.row.km}》吗？`,
                      loading: true,
                      onOk: () => {
                        that.$axios({
                          url: '/student/courseCalendar',
                          method: 'post',
                          data: {
                            cid: params.row.kh,
                            term: that.selected,
                            op: 'delete'
                          }
                        }).then(() => {
                          that.$Modal.remove()
                          that.$Message.info('退课成功')
                          that.data1.splice(params.index, 1)
                          that.renderCalendar()
                        })
                      },
                      onCancel: () => {
                      }
                    })
                  }
                }
              }, '退课')
            ])
          }
        })
      }
      this.renderCalendar()
    }
  }
</script>

<style lang="scss">
  .calendar {
    /*margin: 25px 32.5px;*/

    td {
      text-align: center;
      height: 45px;
    }
  }

  .calendarRaw {
    .ivu-table-cell {
      padding: 0 5px !important;
    }
  }
</style>
