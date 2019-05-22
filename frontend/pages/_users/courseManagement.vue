<template>
  <div class="wrapper">
    <Form class="operation">
      <FormItem label="学期">
        <Row type="flex" justify="space-between">
          <Select
            v-model="selected"
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

    <Drawer
      class="courseManagementDrawer"
      v-model="showModal"
      width="100"
      :mask-closable="false"
    >
      <div slot="header" style="overflow: unset;" class="ivu-drawer-header-inner">
        开课安排
        <Button size="small" type="dashed" icon="md-add"></Button>
      </div>
      <Row>
        <Col span="12">
          <Table
            class="calendarRaw"
            stripe
            border
            :columns="columns2"
            :data="data2"
            ref="calendarRaw"
          ></Table>
        </Col>
        <Col span="12">
          <Table
            class="calendar"
            disabled-hover
            border
            :columns="calendarColumns"
            :data="calendar"
            ref="calendar"
            :height="drawerContentHeight"
          ></Table>
        </Col>
      </Row>
    </Drawer>
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
  export default {
    name: 'teacherManagement',
    mixins: [mock, calendarColumnsMock],
    data() {
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
                    that.$nextTick(() => {
                      that.drawerContentHeight = document.querySelector('.courseManagementDrawer .ivu-drawer-body').offsetHeight
                      console.log(that.drawerContentHeight)
                    })
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
        columns2: [
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
          },
          {
            'title': '操作',
            'key': 'action',
            'fixed': 'right',
            'width': 160,
            'align': 'center',
            render(h, params) {
              return h('ButtonGroup', {
                props: {
                  size: 'small',
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
                      that.showModal = true
                      that.thisRow = params.row
                    }
                  }
                }, '删除')
              ])
            }
          }
        ],
        data1: [{
          kh: '08305001',
          km: '离散数学',
          xf: '4',
          xs: '40',
          dep: '计算机工程与科学学院',
          kks: 3
        }],
        data2: [{
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
        calendar: [{ time: '1' }, { time: '2' }, { time: '3' }, { time: '4' }, { time: '5' }, { time: '6' }, { time: '7' }, { time: '8' }, { time: '9' }, { time: '10' }, { time: '11' }, { time: '12' }, { time: '13' }],
        nodeMatrix: [],
        selected: '',
        showModal: false,
        drawerContentHeight: ''
      }
    },
    methods: {
      renderCalendar({
                       raw = this.data2,
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
            raw[index]._color = color = raw[index]._color || `rgb(${(Math.random() * 195 + 30).toFixed(0)},${(Math.random() * 195 + 30).toFixed(0)},${(Math.random() * 195 + 30).toFixed(0)})`
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
                  <p style="font-weight:bold">${matrix[i][j].km}</p>
                  <p style="font-weight:bold">${matrix[i][j].xm}</p>
                  <p style="font-weight:bold">${matrix[i][j].sksj}</p>
                </div>
              </div>
              `
            this.nodeMatrix[i][j].style.color = 'white'
          }
        }
      },
      mouseEnterTr(line) {
        this.renderCalendar({ raw: [this.$refs.calendarRaw.data[line]], hoverKh: this.$refs.calendarRaw.data[line].kh })
      },
      mouseLeaveTbody() {
        this.renderCalendar({ raw: [] })
      }
    },
    mounted() {
      that = this
      this.renderCalendar({ raw: [] })
    }
  }
</script>

<style lang="scss">
  .courseManagementDrawer .ivu-drawer-body {
    padding: 0 !important;
  }
</style>
<style lang="scss" scoped>

</style>
