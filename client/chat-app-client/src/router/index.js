import { createRouter, createWebHistory } from 'vue-router'
import Home from "../components/home"
import Register from "../components/register"
import Login from "../components/login"

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
    }
  ]
})

export default router
