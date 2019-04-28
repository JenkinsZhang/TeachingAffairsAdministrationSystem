<template>
  <div class="wrapper">
    <no-ssr placeholder="Loading...">
      <Select
        v-model="selectedClassId"
        style="width:400px"
        placeholder="请选择课程"
        class="operation"
      >
        <OptionGroup :label="term.name" v-for="term of terms" :key="term.name">
          <Option v-for="cla of term.classes" :value="cla.kh" :key="cla.kh">{{ cla.km }} | {{cla.kh}} | {{cla.sksj}}
          </Option>
        </OptionGroup>
      </Select>
    </no-ssr>
    <div class="operation">
      <Table stripe border :columns="columns" :data="data1" size="large"></Table>
    </div>
    <Modal
      title="录入成绩"
      v-model="showModal"
      @on-ok="editScore"
      @on-cancel="scoreValue=''"
      width="400px"
    >
      <Input
        v-model="scoreValue"
        autofocus
        placeholder="键入成绩"
        type="number"
        :maxlength=3
        @on-enter="editScore"
      />
    </Modal>
  </div>
</template>

<script>
  // import columns from '~/assets/json/classColumns.json'
  import mock from '~/assets/js/scoreManagementMock'

  let that

  export default {
    name: 'courseTable',
    mixins: [mock],
    data: () => ({
      columns: [
        {
          'title': '学号',
          'key': 'id',
          'align': 'center',
          'sortable': true
        },
        {
          'title': '类型',
          'key': 'type',
          'align': 'center'
        },
        {
          'title': '姓名',
          'key': 'name',
          'align': 'center'
        },
        {
          'title': '性别',
          'key': 'gender',
          'align': 'center'
        },
        {
          'title': '籍贯',
          'key': 'jg',
          'align': 'center'
        },
        {
          'title': '手机号码',
          'key': 'sjhm',
          'align': 'center'
        },
        {
          'title': '院系',
          'key': 'dep',
          'align': 'center',
          'sortable': true
        }
      ],
      loading: false,
      data1: [{
        id: '16121663',
        type: '本科生',
        name: '莫之章',
        gender: '男',
        jg: '四川',
        sjhm: '189xxxx9181',
        dep: '计算机工程与科学学院'
      }, {
        id: '16121670',
        type: '本科生',
        name: 'ybmj',
        gender: '男',
        jg: '河北',
        sjhm: '166xxxx0694',
        dep: '计算机工程与科学学院'
      }],
      selectedClassId: '',
      scoreValue: '',
      showModal: false,
      thisRow: null
    }),
    mounted() {
      that = this
    },
    watch: {
      showModal(val) {
        if (val) {
          setTimeout(() => {
            const e = document.querySelector('.ivu-input[autofocus=autofocus]')
            e && e.focus()
          }, 20)
        }
      }
    },
    methods: {
      editScore(score) {
        this.thisRow.score = typeof score == 'number' ? score : parseInt(this.scoreValue)
        this.showModal = false
        this.scoreValue = ''
      }
    }
  }
</script>
