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
    <Table
      class="operation"
      stripe
      border
      :columns="columns"
      :data="data1"
    ></Table>
  </div>
</template>

<script>

  export default {
    name: 'scoreQuery',
    async asyncData({ app }) {
      const data1 = []
      let terms = null
      let selected = ''
      await app.$axios({
        url: '/student/scoreTable'
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
          url: '/student/scoreTable',
          method: 'post',
          data: {
            term: terms[0]
          }
        }).then((res) => {
          const { cid, cname, credit, score, tid, tname } = res.data
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
              cj: score[i]
            })
          }
        })
      })
      return {
        data1,
        terms,
        selected: selected
      }
    },
    data() {
      return {
        selected: '',
        columns: [{
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
            'title': '成绩',
            'key': 'cj',
            'align': 'center'
          }],
        data1: [],
        terms: []
      }
    },
    methods:{
      handleSelectChange(term) {
        this.data1.length = 0
        this.$axios({
          url: '/student/scoreTable',
          method: 'post',
          data: {
            term
          }
        }).then((res) => {
          const { cid, cname, credit, score, tid, tname } = res.data
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
              cj: score[i]
            })
          }
        })
      },
    }
  }
</script>

<style scoped>

</style>
