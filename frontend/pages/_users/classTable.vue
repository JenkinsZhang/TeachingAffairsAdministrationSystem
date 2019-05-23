<template>
  <div class="wrapper">
    <no-ssr placeholder="Loading...">
      <Select
        v-model="selected"
        style="width:400px"
        placeholder="请选择课程"
        class="operation"
        @on-change="handleSwitchTerm"
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
  </div>
</template>

<script>
  //cid,term
  function requestData(requestBody, axios) {
    return axios({
      url: '/teacher/classTable',
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
          name: data.name[i],
          gender: data.gender[i],
          dname: data.dname[i],
          birthplace: data.birthplace[i],
          phone: data.phone[i],
        })
      }
      return arr
    })
  }

  let that

  export default {
    name: 'courseTable',
    async asyncData({ app }) {
      const terms = []
      let selected = null
      let selectedCid = null
      let selectedTerm = null
      await app.$axios({
        url: '/teacher/classTable'
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
      const data1 = await requestData({ cid: selectedCid, term: selectedTerm }, app.$axios)
      return {
        terms,
        selected,
        data1
      }
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
          'title': '性别',
          'key': 'gender',
          'align': 'center'
        },
        {
          'title': '籍贯',
          'key': 'birthplace',
          'align': 'center'
        },
        {
          'title': '手机号码',
          'key': 'phone',
          'align': 'center'
        },
        {
          'title': '院系',
          'key': 'dname',
          'align': 'center',
          'sortable': true
        }
      ],
      loading: false,
      data1: [],
      selected: '',
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
      async handleSwitchTerm(val) {
        const arr = val.split('|')
        requestData({ cid: arr[0], term: arr[1]}, this.$axios).then((val) => {
          this.data1 = val
        })
      }
    }
  }
</script>
