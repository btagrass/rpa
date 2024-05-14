import { createRouter, createWebHashHistory } from "vue-router"
import { useStore } from "@/store"

var refreshed = true
const pages = import.meta.glob("/src/pages/**/*.vue")
const router = createRouter({
  history: createWebHashHistory(),
  routes: [{
    name: "login",
    path: "/login",
    meta: {
      title: "Login",
    },
    component: () => import("@/pages/home/Login.vue"),
  }, {
    name: "index",
    path: "/",
    redirect: "/",
    component: () => import("@/pages/home/Index.vue"),
    children: [
      {
        name: "dashboard",
        path: "/",
        meta: {
          title: "Home",
        },
        component: () => import("@/pages/home/Dashboard.vue"),
      },
    ],
  }],
})
router.beforeEach((to, _, next) => {
  const { resources, user, web } = useStore()
  document.title = web.title
  if (user.token == "") {
    if (to.fullPath.startsWith("/login")) {
      next()
    } else {
      next("/login")
    }
  } else if (refreshed) {
    resources.forEach((r) => {
      r.children.forEach((c) => {
        router.addRoute("index", {
          name: c.uri.substring(c.uri.lastIndexOf("/") + 1),
          path: c.uri,
          meta: {
            title: c.name,
          },
          component: pages[`/src/pages${c.uri}/Index.vue`],
        })
      })
    })
    refreshed = false
    next({
      ...to,
      replace: true,
    })
  } else {
    if (to.fullPath.startsWith("/http")) {
      window.open(to.fullPath.substring(1), "_blank")
    } else {
      next()
    }
  }
})
router.afterEach((to, _) => {
  const { savePage } = useStore()
  if (!to.fullPath.startsWith("/login")) {
    savePage({
      name: to.name,
      path: to.fullPath,
      title: to.meta.title,
    })
  }
})

export default router
