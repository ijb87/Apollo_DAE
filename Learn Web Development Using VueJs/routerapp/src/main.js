// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import VueRouter from 'vue-router'
import App from './App'
import Home from './components/Home'
import About from './components/About'
import Services from './components/Services'

Vue.use(VueRouter)

const router = new VueRouter({
  mode: 'history',
  base: __dirname,
  routes: [
    { path: '/', component: Home },
    { path: '/about', component: About },
    { path: '/services', component: Services }
  ]
})

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  router,
  template: `<div id="app">
    <ul class="nav">
      <li><router-link to="/">Home</router-link></li>
      <li><router-link to="/about">About</router-link></li>
      <li><router-link to="/services">Services</router-link></li>
    </ul>
    <router-view class="view"></router-view>
  </div>
  `
}).$mount('#app')
