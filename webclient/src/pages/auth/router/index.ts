import Vue from 'vue'
import VueRouter, { RouteConfig } from 'vue-router'
import LoginView from "@/pages/auth/views/LoginView.vue";

Vue.use(VueRouter)

const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'login',
    component: LoginView
  }
]

const router = new VueRouter({
  routes
})

export default router
