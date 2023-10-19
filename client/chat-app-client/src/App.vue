<template>
  <header>
    <RouterView/>
  </header>
</template>

<script>
import { RouterLink, RouterView } from 'vue-router';

  export default {
    name: "App",
    components: { RouterLink, RouterView }
}

</script>

<style scoped>

</style>