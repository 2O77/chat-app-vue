<template>
    <div id="home">
        <sidebar></sidebar>
        <mainbar></mainbar>
    </div>
</template>

<script>
import mainbar from '../components/mainbar.vue';
import sidebar from '../components/sidebar.vue';

    export default {
        name: "Home", 
        components: {
           "sidebar" : sidebar,
           "mainbar" : mainbar
        }
    }

</script>

<style>

   #home {
        width: 100%;
        height: 100%;
        padding: 0;
        display: flex;
        flex-direction: row;
   } 

</style>