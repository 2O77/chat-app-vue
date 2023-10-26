<template>
    <div id="home">
        <sidebar :username="user.Username" :id="user.ID"></sidebar>
        <mainbar></mainbar>
    </div>
</template>

<script>
import mainbar from '../components/mainbar.vue';
import sidebar from '../components/sidebar.vue';
import VueCookies from "vue-cookies"
import {ref, onMounted} from "vue"
import axios from "axios"

    export default {
        name: "Home", 
        setup() {
            const user = ref('');
            const userToken = VueCookies.get('userToken');

            onMounted(async () => {
                try {
                    const response = await axios.get('http://localhost:2000/user', {
                        headers: {
                            Token: userToken
                        }
                    });
                    user.value = response.data;
                } catch (error) {
                    console.error(error);
                }
            });

            return {
                user
            };
        },
        components: {
           "sidebar" : sidebar,
           "mainbar" : mainbar
        },
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