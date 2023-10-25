<template>
  <header id="index">
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

<style>

   * {
    margin: 0;
    box-sizing: border-box;
   }

   html {
    width: 100%;
    height: 100%;
   }

   body {
    width: 100%;
    height: 100%;
   }

   #app {
    width: 100%;
    height: 100%;
   }

  #index {
    width: 100%;
    height: 100%;
  }
  
</style>