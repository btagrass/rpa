import { createI18n } from "vue-i18n"
import en from "@/i18n/en"
import zh from "@/i18n/zh"

const i18n = createI18n({
    legacy: false,
    messages: {
        en: en,
        zh: zh,
    }
})

export default i18n
