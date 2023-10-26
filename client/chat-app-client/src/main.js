import { VueElement, createApp } from 'vue'
import App from './App.vue'
import router from './router'
import axios from "axios"
import vueCookies from "vue-cookies"

const app = createApp(App)
app.use(router)
app.config.globalProperties.$axios = axios;

app.use(vueCookies)

app.mount('#app')
