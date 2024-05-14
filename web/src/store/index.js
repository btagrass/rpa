import { computed, reactive, toRefs } from "vue"
import { defineStore } from "pinia"

export const useStore = defineStore("store", () => {
    const state = reactive({
        pages: [],
        resources: [],
        user: {
            token: "",
        },
        web: {
            collapsed: false,
            locale: navigator.language,
            title: "Grass RPA",
        },
    })
    const names = computed(() => {
        return state.pages.map((i) => {
            return i.name
        }).join(",")
    })
    const clearPages = () => {
        state.pages.splice(0, state.pages.length)
    }
    const removePage = (path) => {
        const index = state.pages.findIndex((i) => {
            return i.path == path
        })
        state.pages.splice(index, 1)
    }
    const savePage = (page) => {
        const exist = state.pages.some((i) => {
            return i.path == page.path
        })
        if (!exist) {
            state.pages.push(page)
        }
    }
    const clearResources = () => {
        state.resources = []
    }
    const saveResources = (resources) => {
        state.resources = resources
    }
    const clearUser = () => {
        state.user = {
            token: "",
        }
    }
    const saveUser = (user) => {
        state.user = user
    }
    const toggleCollapsed = () => {
        state.web.collapsed = !state.web.collapsed
    }
    const saveLocale = (locale) => {
        state.web.locale = locale
    }
    return {
        ...toRefs(state),
        names,
        clearPages,
        removePage,
        savePage,
        clearResources,
        saveResources,
        clearUser,
        saveUser,
        toggleCollapsed,
        saveLocale,
    }
}, {
    persist: true,
})
