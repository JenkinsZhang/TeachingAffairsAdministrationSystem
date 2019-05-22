<template>
  <div class="wrapper">
    <no-ssr placeholder="Loading...">
      <Select
        v-model="selected"
        style="width:400px"
        placeholder="请选择课程"
        class="operation"
        @on-change="onSelectChange"
      >
        <OptionGroup :label="term.name" v-for="term of terms" :key="term.name">
          <Option v-for="cla of term.classes" :value="cla.kh+'|'+term.name" :key="cla.kh">{{ cla.km }} | {{cla.kh}} |
            {{cla.sksj}}
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
  //cid,term
  function requestData(requestBody, axios) {
    return axios({
      url: '/teacher/scoreManagement',
      method: 'post',
      data: {
        ...requestBody,
        op: 'query'
      }
    }).then((res) => {
      const data = res.data
      const arr = []
      if (!data.id) {
        return arr
      }
      for (let i = 0; i < data.id.length; i++) {
        arr.push({
          id: data.id[i],
          dname: data.dname[i],
          name: data.name[i],
          score: data.score[i]
        })
      }
      return arr
    })
  }

  let that

  export default {
    name: 'scoreManagement',
    async asyncData({ app }) {
      const terms = []
      let selected = null
      let selectedCid = null
      let selectedTerm = null
      await app.$axios({
        url: '/teacher/scoreManagement'
      }).then((res) => {
        const m = {}
        let cnt = 0
        const data = res.data
        if (!data.cid) {
          return
        }
        for (let i = 0; i < data.cid.length; i++) {
          if (!selected) {
            selected = data.cid[i] + '|' + data.term[i]
            selectedCid = data.cid[i]
            selectedTerm = data.term[i]
          }
          if (m[data.term[i]] == null) {
            m[data.term[i]] = cnt++
            terms.push({
              name: data.term[i],
              classes: [{
                km: data.cname[i],
                kh: data.cid[i],
                sksj: data.classTime[i]
              }]
            })
          } else {
            terms[m[data.term[i]]].classes.push({
              km: data.cname[i],
              kh: data.cid[i],
              sksj: data.classTime[i]
            })
          }
        }
      })
      let data1 = await requestData({ cid: selectedCid, term: selectedTerm }, app.$axios)
      console.log(data1)
      return { terms, selected, data1 }
    },
    data: () => ({
      columns: [
        {
          'title': '学号',
          'key': 'id',
          'align': 'center',
          'sortable': true
        },
        {
          'title': '姓名',
          'key': 'name',
          'align': 'center'
        },
        {
          'title': '院系',
          'key': 'dname',
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
          'width': 150,
          'align': 'center',
          render(h, params) {
            return h('ButtonGroup', {
              props: {
                size: 'small'
              }
            }, [
              h('Button', {
                props: {
                  type: 'info',
                  icon: 'md-checkmark',
                  disabled: params.row.grade
                },
                on: {
                  click: () => {
                    that.showModal = true
                    that.thisRow = params.row
                  }
                }
              }, '录入成绩')
            ])
          }
        }
      ],
      loading: false,
      data1: [],
      selected: '',
      scoreValue: '',
      showModal: false,
      thisRow: null,
      terms: []
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
      editScore() {
        const score = this.scoreValue
        if (parseInt(score) < 0 || parseInt(score) > 100) {
          this.$Message.warning('成绩必须是0~100的整数')
          return
        }
        const arr = this.selected.split('|')
        this.$axios({
          url: '/teacher/scoreManagement',
          method: 'post',
          data: {
            'id': this.thisRow.id,
            'cid': arr[0],
            'term': arr[1],
            'score': this.scoreValue,
            'op': 'change'
          }
        }).then((res) => {
          if (res.data.message === 'ok') {
            this.thisRow.score = score
          } else {
            this.$Message.warning(res.data.message)
          }
        })
        this.showModal = false
        this.scoreValue = ''
      },
      async onSelectChange(val) {
        const arr = val.split('|')
        requestData({ cid: arr[0], term: arr[1]}, this.$axios).then((val) => {
          this.data1 = val
        })
      }
    }
  }
</script>
