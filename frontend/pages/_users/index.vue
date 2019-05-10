<template>
</template>

<script>
  export default {
    name: 'index',
    validate({ params, redirect }) {
      redirect(`/${params.users}/profile`)
      return true
    }
  }
</script>

<style scoped>

</style>
