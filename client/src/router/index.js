import Vue from 'vue'
import VueRouter from 'vue-router'
import Hello from '../views/Home'
import Login from '../views/Login'
import VerifikasiUser from '../views/VerifikasiUser'
import auth from '../auth'


import VueResource from 'vue-resource'
Vue.use(VueResource)

Vue.use(VueRouter)

function requireAuth(to, from, next) {
    if (!auth.isAuthenticated()) {
        router.push({
            path: '/login'
        })
    } else {
        next()
    }
}
const routes = [{
        path: '/login',
        name: 'login',
        component: Login,
    },
    {
        path: '/home',
        name: 'home',
        component: Hello,
        beforeEnter: requireAuth
    },
    {
        path: '/verifikasi/:token',
        name: 'verifikasi',
        component: VerifikasiUser,
    },

]

const router = new VueRouter({
    mode: 'history',
    routes
})


export default router