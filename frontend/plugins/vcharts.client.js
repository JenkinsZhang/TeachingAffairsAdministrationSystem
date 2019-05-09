import Vue from 'vue'
import VeLine from 'v-charts/lib/line.common'
import VePie from 'v-charts/lib/pie.common'
import VeHistogram from 'v-charts/lib/histogram.common'


export default () => {
  Vue.component(VeLine.name, VeLine)
  Vue.component(VePie.name, VePie)
  Vue.component(VeHistogram.name, VeHistogram)
}
