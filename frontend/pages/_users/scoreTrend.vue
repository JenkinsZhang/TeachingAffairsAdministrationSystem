<template>
  <div class="wrapper">
    <no-ssr placeholder="Loading...">
      <ve-line :data="chartData" :settings="chartSettings" class="operation"></ve-line>
    </no-ssr>
  </div>
</template>

<script>
  import mock from '~/assets/js/scoreManagementMock'

  export default {
    name: 'scoreTrend',
    mixins: [mock],
    async asyncData({ app }) {
      const rows = []
      await app.$axios({
        url: '/student/scoreTrend'
      }).then((res) => {
        const { score, term } = res.data
        if (!term) {
          return
        }
        for (let i = 0; i < term.length; i++) {
          rows.push({
            date: term[i],
            val: score[i]
          })
        }
      })
      return {
        chartData: {
          columns: ['date', 'val'],
          rows
        }
      }
    },
    data() {
      this.chartSettings = {
        labelMap: {
          'val': '平均分数'
        }
      }
      return {
        chartData: {
          columns: ['date', 'val'],
          rows: []
        },
      }
    }
  }
</script>

<style scoped>

</style>
