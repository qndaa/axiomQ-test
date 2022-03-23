import {createApp} from 'vue'
import {createRouter, createWebHistory} from "vue-router";
import Home from "@/components/Home";
import Login from "@/components/Login";
import NewBook from "@/components/NewBook";
import store from "@/store";
import App from "@/App";
import Book from "@/components/Book";
import Forbidden from "@/components/Forbidden";
import NotFoundedPage from "@/components/NotFoundedPage";

const router = createRouter({
    mode: 'history',
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            component: Home
        }, {
            path: '/login',
            component: Login
        }, {
            path: '/new-book',
            component: NewBook,
        }, {
            path: '/home',
            component: Home
        }, {
            path: '/book/:id',
            component: Book,
            props: true
        }, {
            path: '/forbidden',
            component: Forbidden
        }, {
            path: '/*',
            component: NotFoundedPage
        }
    ]
})

const app = createApp(App);
app.use(router);
app.use(store);
app.mount("#app")

