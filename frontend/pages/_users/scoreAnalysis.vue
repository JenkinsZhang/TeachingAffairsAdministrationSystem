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
          <Option v-for="cla of term.classes" :value="cla.kh+'|'+term.name" :key="cla.kh">
            {{ cla.km }} | {{cla.kh}} | {{cla.sksj}}
          </Option>
        </OptionGroup>
      </Select>
      <Row type="flex" class="operation">
        <Col span="12">
          <ve-histogram :data="chartData" :settings="chartSettings" class="operation"></ve-histogram>
        </Col>
        <Col span="12">
          <ve-pie :data="chartData" style="position: relative; top:25px;"></ve-pie>
        </Col>
      </Row>

    </no-ssr>
  </div>
</template>

<script>
  function requestData(requestBody, axios) {
    return axios({
      url: '/teacher/scoreAnalysis',
      method: 'post',
      data: requestBody
    }).then((res) => {
      let dis = [0, 0, 0, 0, 0]
      res.data.score.forEach((score) => {
        if (score >= 90) {
          dis[4]++
        } else if (score >= 80) {
          dis[3]++
        } else if (score >= 70) {
          dis[2]++
        } else if (score >= 60) {
          dis[1]++
        } else {
          dis[0]++
        }
      })
      const ret = []
      if (dis[0]) {
        ret.push({ 'date': '0~59', 'val': dis[0] })
      }
      if (dis[1]) {
        ret.push({ 'date': '60~69', 'val': dis[1] })
      }
      if (dis[2]) {
        ret.push({ 'date': '70~79', 'val': dis[2] })
      }
      if (dis[3]) {
        ret.push({ 'date': '80~89', 'val': dis[3] })
      }
      if (dis[4]) {
        ret.push({ 'date': '90~100', 'val': dis[4] })
      }
      return ret
    })
  }

  export default {
    name: 'scoreSummary',
    async asyncData({ app }) {
      const terms = []
      let selected = null
      let selectedCid = null
      let selectedTerm = null
      await app.$axios({
        url: '/teacher/scoreAnalysis'
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
      const rows = await requestData({ cid: selectedCid, term: selectedTerm }, app.$axios)
      return {
        terms,
        selected,
        chartData: {
          columns: ['date', 'val'],
          rows
        }
      }
    },
    data() {
      this.chartSettings = {
        labelMap: {
          'val': '人数'
        }
      }
      return {
        chartData: null,
        selected: '',
        terms: null
      }
    },
    methods: {
      async handleSwitchTerm(val) {
        const arr = val.split('|')
        requestData({ cid: arr[0], term: arr[1] }, this.$axios).then((val) => {
          this.chartData.rows = val
        })
      }
    },
    mounted() {
      console.log(this.chartData)
    }
  }
</script>

<style scoped>

</style>
