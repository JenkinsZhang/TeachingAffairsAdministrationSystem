<template>
  <div class="wrapper">
    <Form class="operation">
      <FormItem label="学期">
        <Select
          v-model="selectedClassId"
          style="width:200px"
          placeholder="请选择学期"
        >
          <Option v-for="term of terms" :value="term.name" :key="term.name">{{term.name}}</Option>
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
  import mock from '~/assets/js/scoreManagementMock'
  import calendarColumnsMock from '~/assets/js/calendarColumnsMock'

  let that
  const tiptopMap = {
    '一': 1,
    '二': 2,
    '三': 3,
    '四': 4,
    '五': 5,
    '六': 6,
    '七': 7,
    '日': 7
  }
  const colorArray = ['#d50000', '#E91E63', '#9C27B0', '#673AB7', '#3F51B5', '#2196F3', '#03A9F4', '#00BCD4', '#009688', '#4CAF50', '#FF5722', '#795548', '#607D8B']
  export default {
    name: 'courseCalender',
    mixins: [mock, calendarColumnsMock],
    asyncData({ params }) {
      let isStudent = false
      if (params.users === 'student') {
        isStudent = true
      }
      return {
        isStudent
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
            'align': 'center'
          }
        ],
        data1: [{
          kh: '01015045',
          km: '数据结构与算法',
          gh: '1000',
          xm: '朱明',
          xf: '4',
          sksj: '星期三11-13'
        }, {
          kh: '08305013',
          km: '编译原理',
          gh: '1004',
          xm: '滕中梅',
          xf: '5',
          sksj: '星期二7-9'
        }, {
          kh: '08305015',
          km: '数据库原理(2)',
          gh: '1005',
          xm: '叶飞跃',
          xf: '4',
          sksj: '星期三1-2'
        }, {
          kh: '08305124',
          km: '计算机系统结构',
          gh: '1004',
          xm: '刘方方',
          xf: '4',
          sksj: '星期三5-6'
        }],
        selectedClassId: '',
        calendar: [{ time: '1' }, { time: '2' }, { time: '3' }, { time: '4' }, { time: '5' }, { time: '6' }, { time: '7' }, { time: '8' }, { time: '9' }, { time: '10' }, { time: '11' }, { time: '12' }, { time: '13' }],
        nodeMatrix: [],
        isStudent: false,
        calendarHeight: null
      }
    },
    methods: {
      renderCalendar({
                       raw = this.data1,
                       hoverKh
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
          if (x.sksj.indexOf('星期') === 0) {
            x.sksj = x.sksj.substring(2)
          }
          const day = tiptopMap[x.sksj[0]]
          const arr = x.sksj.substring(1).split('-')
          let begin = parseInt(arr[0]),
            end = parseInt(arr[1])
          //给课程初始化一个颜色
          let color
          if (hoverKh === undefined || obj.kh === hoverKh) {//没有hover的元素，或hover的就是这个课程
            swap(colorArrayBak, cnt, Math.floor(Math.random() * (colorArrayBak.length - cnt) + cnt))
            raw[index]._color = color = raw[index]._color || colorArrayBak[cnt++]
          } else {
            color = 'rgba(0,0,0,.2)'//灰色
          }
          matrix[begin - 1][day - 1] = {
            ...obj,
            color,
            _length: end - begin + 1
          }
        })
        //开始渲染
        if (!this.nodeMatrix.length) {
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
        this.renderCalendar({ hoverKh: this.$refs.calendarRaw.data[line].kh })
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
                    console.log('点了退课')
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
