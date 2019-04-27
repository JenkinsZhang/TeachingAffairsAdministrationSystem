<template>
  <div class="wrapper">
    <no-ssr placeholder="Loading...">
      <ve-line :data="chartData" :settings="chartSettings" class="operation"></ve-line>
    </no-ssr>
  </div>
</template>

<script>
  import mock from './scoreManagementMock'

  export default {
    name: 'scoreTrend',
    mixins: [mock],
    data() {
      this.chartSettings = {
        labelMap: {
          'val': '平均分数'
        },
      }
      return {
        chartData: {
          columns: ['date', 'val'],
          rows: [
            { 'date': '2017-2018冬', 'val': 87 },
            { 'date': '2017-2018春', 'val': 89 },
            { 'date': '2017-2018夏', 'val': 94 },
            { 'date': '2018-2019秋', 'val': 92 },
            { 'date': '2018-2019冬', 'val': 85 },
            { 'date': '2018-2019春', 'val': 90 },
            { 'date': '2018-2019夏', 'val': 91 }
          ]
        },
        selectedClassId: ''
      }
    }
  }
</script>

<style scoped>

</style>
