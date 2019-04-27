<template>
  <div class="wrapper">
    <no-ssr placeholder="Loading...">
      <p class="operation">
        <Select
          v-model="selectedClassId"
          style="width:400px"
          placeholder="请选择课程"
        >
          <OptionGroup :label="term.name" v-for="term of terms" :key="term.name">
            <Option v-for="cla of term.classes" :value="cla.kh" :key="cla.kh">{{ cla.km }} | {{cla.kh}} | {{cla.sksj}}
            </Option>
          </OptionGroup>
        </Select>
      </p>
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
  import mock from './scoreManagementMock'

  let that

  export default {
    name: 'scoreManagement',
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
          'title': '院系',
          'key': 'dep',
          'align': 'center',
          'sortable': true
        },
        {
          'title': '成绩',
          'key': 'score',
          'align': 'center',
          'sortable': true,
          filters: [{
            label: '未登分',
            value: 1
          }],
          filterMethod(value, row) {
            if (value === 1) {
              return !row.score
            }
          }
        },
        {
          'title': '操作',
          'key': 'action',
          'fixed': 'right',
          'width': 200,
          'align': 'center',
          render: (h, params) => {
            return (
              <ButtonGroup size="default">
                <Button type="info" icon="md-checkmark" onClick={() => {
                  that.showModal = true
                  that.thisRow = params.row
                }}>修改成绩</Button>
              </ButtonGroup>
            )
          }
        }
      ],
      loading: false,
      data1: [{
        id: '16121663',
        type: '本科生',
        name: '莫之章',
        dep: '计算机工程与科学学院',
        score: 100
      }, {
        id: '16121670',
        type: '本科生',
        name: 'ybmj',
        dep: '计算机工程与科学学院',
        score: null
      }],
      selectedClassId: '',
      scoreValue: '',
      showModal: false,
      thisRow: null
    }),
    mounted() {
      that = this
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
