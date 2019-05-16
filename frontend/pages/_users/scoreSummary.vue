<template>
  <div class="wrapper">
    <div class="operation">
      <Row
        type="flex"
        justify="space-around"
        class="header"
      >
        <Col>
          <p>院系：{{profile.dname}}</p>
          <p>姓名：{{profile.name}}</p>
        </Col>
        <Col>
          <p>学号：{{profile.id}}</p>
          <p>性别：{{profile.gender}}</p>
        </Col>
      </Row>
    </div>
    <Table class="operation" stripe border :columns="columns" :data="data1" size="small"></Table>
    <p class="tail">获得学分：{{sumxf}}，平均成绩：{{avecj}}</p>
  </div>
</template>

<script>
  export default {
    name: 'scoreSummary',
    async asyncData({ app, params }) {
      let data1 = [], sumxf = 0, avecj = 0, profile = null
      const p1 = app.$axios({
        url: `/${params.users}/profile`
      }).then((res) => {
        profile = res.data
      })
      const p2 = app.$axios({
        url: '/student/scoreSummary'
      }).then((res) => {
        const { cid, cname, credit, score, tid, tname, term } = res.data
        for (let i = 0; i < cid.length; i++) {
          data1.push({
            kh: cid[i],
            km: cname[i],
            xf: credit[i],
            cj: score[i],
            xq: term[i],
            bz: ''
          })
          sumxf += Number(credit[i])
          avecj += score[i] * credit[i]
        }
        avecj = (avecj / sumxf).toFixed(2)
      })
      await Promise.all([p1, p2])
      return { data1, sumxf, avecj, profile }
    },
    data() {
      return {
        profile: {},
        columns: [{
          'title': '课程号',
          'key': 'kh',
          'align': 'center'
        }, {
          'title': '课程名',
          'key': 'km',
          'align': 'center'
        }, {
          'title': '学分',
          'key': 'xf',
          'align': 'center'
        }, {
          'title': '成绩',
          'key': 'cj',
          'align': 'center'
        }, {
          'title': '学期',
          'key': 'xq',
          'align': 'center'
        }, {
          'title': '备注',
          'key': 'bz',
          'align': 'center'
        }],
        data1: [],
        avecj: 0,
        sumxf: 0
      }
    }
  }
</script>

<style scoped>
  .header {
    padding: 0 20%;
    font-size: 14px;
  }

  .tail {
    margin: 3px 0 0 32.5px;
    font-size: 13px
  }
</style>
