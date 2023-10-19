import { createRouter, createWebHistory } from 'vue-router'
import Login from "../views/login.vue"
import Register from "../views/register.vue"
import Home from "../views/home.vue"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: "/register",
      name: "register",
      component: Register
    },
    {
      
    }
  ]
})

export default router
