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
    <Table
      class="operation"
      stripe
      border
      :columns="columns"
      :data="data1"
      size="small"
    ></Table>
    <Table
      class="calendar"
      border
      disabled-hover
      :columns="calendarColumns"
      :data="calendar"
      ref="calendar"
    ></Table>
  </div>
</template>

<script>
  import mock from './scoreManagementMock'
  import calendarColumnsMock from './calendarColumnsMock.js'

  let that

  export default {
    name: 'courseCalender',
    mixins: [mock, calendarColumnsMock],
    asyncData({ params }) {
      let isStudent = false
      if (params.users === 'student') {
        isStudent = true
      }
      //transfer raw to data1 and matrix
      const raw = [{
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
      }]

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
      const matrix = new Array(13)
      for (let i = 0; i < 13; i++) {
        matrix[i] = new Array(7)
      }
      raw.forEach((obj) => {
        const x = Object.assign({}, obj)
        if (x.sksj.indexOf('星期') === 0) {
          x.sksj = x.sksj.substring(2)
        }
        const day = tiptopMap[x.sksj[0]]
        const arr = x.sksj.substring(1).split('-')
        let begin = parseInt(arr[0]),
          end = parseInt(arr[1])
        //给课程初始化一个颜色
        const color = `rgb(${Math.random() * 195 + 30},${Math.random() * 195 + 30},${Math.random() * 195 + 30})`
        for (let i = begin; i <= end; i++) {
          matrix[i - 1][day - 1] = {
            ...x,
            color
          }
        }
      })
      return {
        data1: raw,
        isStudent,
        matrix
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
        data1: [],
        selectedClassId: '',
        calendar: [{
          time: '1'
        }, { time: '2' }, { time: '3' }, { time: '4' }, { time: '5' }, { time: '6' }, { time: '7' }, { time: '8' }, { time: '9' }, { time: '10' }, { time: '11' }, { time: '12' }, { time: '13' }],
        matrix: [],
        nodeMatrix: [],
        isStudent: false
      }
    },
    methods: {
      renderCalendar() {
        // for (let i = 0; i < this.matrix.length; i++) {
        //   for (let j = 0; j < this.matrix[i].length; j++) {
        //     if (!this.matrix[i][j]) {
        //       continue
        //     }
        //     this.calendar[i][j + 1] = `<div>${this.matrix[i][j].kh}</div>`
        //   }
        // }
        // this.updateCalendar++
        setTimeout(() => {
          const tr = document.querySelectorAll('.calendar .ivu-table-tbody tr')
          for (let i = 0; i < tr.length; i++) {
            const tds = tr[i].querySelectorAll('td:not(:first-child)')
            this.nodeMatrix.push(tds)
          }
          for (let i = 0; i < this.matrix.length; i++) {
            for (let j = 0; j < this.matrix[i].length; j++) {
              if (!this.matrix[i][j]) {
                continue
              }
              this.nodeMatrix[i][j].style.background = this.matrix[i][j].color
              this.nodeMatrix[i][j].firstChild.innerHTML = `
              <p style="font-weight:bold">${this.matrix[i][j].kh}</p>
              <p style="font-weight:bold">${this.matrix[i][j].km}</p>
              `
              this.nodeMatrix[i][j].style.color = 'white'
            }
          }
        }, 20)
      }
    },
    mounted() {
      that = this
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
    margin: 25px 32.5px;

    td {
      text-align: center;
      height: 60px;
    }
  }
</style>
