import Vue from 'vue'
import VueRouter, { RouteConfig } from 'vue-router'
import LoginView from "@/pages/auth/views/LoginView.vue";
import RegisterView from "@/pages/auth/views/RegisterView.vue";

Vue.use(VueRouter)

const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'login',
    component: LoginView
  },
  {
    path: '/register',
    name: 'register',
    component: RegisterView
  }
]

const router = new VueRouter({
  routes
})

export default router
