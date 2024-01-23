import './assets/main.css'

import { createApp } from 'vue'
import {createRouter, createWebHistory} from 'vue-router';
import { autoAnimatePlugin } from '@formkit/auto-animate/vue'
import App from './App.vue'
const app = createApp(App)

import Home from './pages/home.vue'
import Favorites from './pages/favorites.vue'
import Profile from "@/pages/Profile.vue";

const routes = [
    { path: '/', name: 'Home', component: Home },
    { path: '/favorites', name: 'Favorites', component: Favorites },
    { path: '/profile', name: 'Profile', component: Profile },
]



const router = createRouter({
    history: createWebHistory(),
    routes
})

app.use(router)


app.use(autoAnimatePlugin)
app.mount('#app')
