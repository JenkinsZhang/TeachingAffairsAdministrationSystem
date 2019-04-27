<template>
  <div class="wrapper">
    <div class="operation">
      <Row
        type="flex"
        justify="space-around"
        class="header"
      >
        <Col>
          <p>院系：{{profile.dep}}</p>
          <p>姓名：{{profile.name}}</p>
        </Col>
        <Col>
          <p>学号：{{profile.id}}</p>
          <p>性别：{{profile.xb}}</p>
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
    data() {
      return {
        profile: {
          name: '莫之章',
          id: '16121663',
          xb: '男',
          dep: '计算机工程与科学学院',
          type: '本科生'
        },
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
        avecj: 92,
        sumxf: 117
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
