import { createRouter, createWebHistory, RouteLocationNormalized, RouteRecordRaw } from "vue-router"
import { getSelf } from "../compositions/useUser";

const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        name: 'Home',
        component: () => import('../views/Home.vue')
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

const DEFAULT_NAME = import.meta.env.VITE_PAGE_NAME ?? "Mappool Collabo";

router.beforeResolve((to: RouteLocationNormalized) => {
    // set the page title to "osu! Mappool Collabo | `${route name}`"
    document.title = `${DEFAULT_NAME} | ${String(to.name)}`
})

router.beforeEach(async (to: RouteLocationNormalized) => {
  // if we get redirected to /login?token= grab the token and put it in localstorage
    if (to.query.token) {
        localStorage.setItem("auth_token", String(to.query.token))
        window.location.href = "/"
    }
})

export default router