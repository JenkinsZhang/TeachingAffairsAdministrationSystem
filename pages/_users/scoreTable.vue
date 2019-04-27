<template>
  <div class="wrapper">
    <Form class="operation">
      <FormItem label="学期">
        <Select
          v-model="selectedClassId"
          style="width:200px"
          placeholder="请选择学期"
        >
          <Option v-for="term of terms" :value="term.name" :key="term.name">{{term.name}}</Option>
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
  import mock from '~/assets/js/scoreManagementMock'

  export default {
    name: 'scoreQuery',
    mixins: [mock],
    data() {
      return {
        selectedClassId: '',
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
        data1: []
      }
    }
  }
</script>

<style scoped>

</style>
