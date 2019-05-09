<template>
</template>

<script>
  export default {
    name: 'index',
    validate({ params, redirect }) {
      redirect(`/${params.users}/profile`)
    }
  }
</script>

<style scoped>

</style>
