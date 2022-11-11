import { createRouter, createWebHistory, RouteLocationNormalized, RouteRecordRaw } from "vue-router"

const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        name: 'Home',
        component: () => import('../views/Home.vue')
    },
    {
        path: "/login",
        name: "Login",
        component: () => import("../views/Login.vue")
    },
    {
        path: "/:id",
        name: "Tournament",
        component: () => import("../views/Tournament.vue")
    },
    {
        path: "/:id/suggestions",
        name: "Suggestions",
        component: () => import("../views/Suggestions.vue")
    },
    {
        path: "/:id/replays",
        name: "Replays",
        component: () => import("../views/Replays.vue")
    },
    {
        path: "/:id/mappool",
        name: "Mappool",
        component: () => import("../views/Mappool.vue")
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

const DEFAULT_NAME = import.meta.env.VITE_PAGE_NAME;

router.beforeResolve((to: RouteLocationNormalized) => {
    // set the page title to "osu! Mappool Collabo | `${route name}`"
    document.title = `${DEFAULT_NAME} | ${String(to.name)}`
})

export default router