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
  import mock from '~/assets/js/scoreManagementMock'

  export default {
    name: 'scoreSummary',
    mixins: [mock],
    data() {
      this.chartSettings = {
        labelMap: {
          'val': '人数'
        }
      }
      return {
        chartData: {
          columns: ['date', 'val'],
          rows: [
            { 'date': '0~59', 'val': 2 },
            { 'date': '60~69', 'val': 15 },
            { 'date': '70~79', 'val': 21 },
            { 'date': '80~89', 'val': 12 },
            { 'date': '90~100', 'val': 6 }
          ]
        },
        selectedClassId: ''
      }
    }
  }
</script>

<style scoped>

</style>
