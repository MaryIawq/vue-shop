import './assets/main.css'
import { createApp } from 'vue'
import {createRouter, createWebHistory} from 'vue-router';
import { autoAnimatePlugin } from '@formkit/auto-animate/vue'
import App from './App.vue'
const app = createApp(App).use(store)

import Home from './pages/home.vue'
import Favorites from './pages/favorites.vue'
import Profile from "./pages/Profile.vue";
import Auth from "./pages/Auth.vue";
import store from './store'

const routes = [
    { path: '/', name: 'Home', component: Home },
    { path: '/favorites', name: 'Favorites', component: Favorites },
    { path: '/profile', name: 'Profile', component: Profile },
    { path: '/auth', name: 'Auth', component: Auth },
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

app.use(router)


app.use(autoAnimatePlugin)
app.mount('#app')
