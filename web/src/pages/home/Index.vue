<template>
  <el-config-provider :locale="eLocale">
    <el-container>
      <el-header>
        <div class="row-side">
          <div class="row-center">
            <div class="hover" @click="toggleCollapsed">
              <el-icon v-if="web.collapsed">
                <expand></expand>
              </el-icon>
              <el-icon v-else>
                <fold></fold>
              </el-icon>
            </div>
            <span>{{ web.title }}</span>
          </div>
          <div class="row-center">
            <div class="hover" @click="clearTabs">
              <el-tooltip :content="$t('l.Clear')">
                <el-icon>
                  <FolderDelete />
                </el-icon>
              </el-tooltip>
            </div>
            <div class="hover" @click="toggle">
              <el-tooltip :content="$t('l.FullScreen')">
                <el-icon>
                  <FullScreen />
                </el-icon>
              </el-tooltip>
            </div>
            <div class="hover" @click="toggleLang">
              <el-tooltip :content="$t('l.Lang')">
                <el-icon>
                  <Eleme />
                </el-icon>
              </el-tooltip>
            </div>
            <el-dropdown class="hover" @command="command">
              <div>
                <el-icon>
                  <User />
                </el-icon>
                <span>{{ user.userName }}</span>
              </div>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item command="logout">{{ $t("l.Logout") }}</el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </div>
        </div>
      </el-header>
      <el-container>
        <el-aside :width="web.collapsed ? '72px' : '302px'">
          <el-menu :collapse="web.collapsed" :default-active="route.fullPath" unique-opened router>
            <template v-for="r in resources" :key="r.id">
              <el-sub-menu v-if="r.children.length" :index="r.uri">
                <template #title>
                  <el-icon>
                    <component :is="r.icon"></component>
                  </el-icon>
                  <span>{{ $t("m.$v".replace("$v", r.name)) }}</span>
                </template>
                <el-menu-item v-for="c in r.children" :key="c.id" :index="c.uri">
                  <el-icon>
                    <component :is="c.icon"></component>
                  </el-icon>
                  <span>{{ $t("m.$v".replace("$v", c.name)) }}</span>
                </el-menu-item>
              </el-sub-menu>
              <el-menu-item v-else :index="r.uri">
                <el-icon>
                  <component :is="r.icon"></component>
                </el-icon>
                <span>{{ $t("m.$v".replace("$v", r.name)) }}</span>
              </el-menu-item>
            </template>
          </el-menu>
        </el-aside>
        <el-main>
          <el-tabs v-model="route.fullPath" type="card" @tab-click="clickTab" @tab-remove="removeTab">
            <el-tab-pane v-for="p in pages" :key="p.path" :label="$t('m.$v'.replace('$v', p.title))" :name="p.path"
              :closable="p.path != '/'"></el-tab-pane>
          </el-tabs>
          <div ref="fullscreen">
            <router-view v-slot="{ Component }">
              <keep-alive :include="names">
                <component :is="Component"></component>
              </keep-alive>
            </router-view>
          </div>
        </el-main>
      </el-container>
      <el-footer>https://github.com/btagrass</el-footer>
    </el-container>
  </el-config-provider>
</template>

<script setup name="index">
import { ref, onMounted } from "vue"
import { useI18n } from "vue-i18n"
import { useRoute, useRouter } from "vue-router"
import en from "element-plus/es/locale/lang/en"
import zh from "element-plus/es/locale/lang/zh-cn"
import { useFullscreen } from "@vueuse/core"
import { useStore } from "@/store"
import { computed } from "vue"

const { locale } = useI18n()
const route = useRoute()
const router = useRouter()
const { pages, resources, user, web, names, clearPages, removePage, clearResources, clearUser, toggleCollapsed, saveLocale } = useStore()
const clearTabs = () => {
  clearPages()
  router.push("/")
}
const fullscreen = ref(null)
const { toggle } = useFullscreen(fullscreen)
const toggleLang = () => {
  if (locale.value == "en") {
    locale.value = "zh"
  } else {
    locale.value = "en"
  }
  saveLocale(locale)
}
const eLocale = computed(() => (locale.value == "en" ? en : zh))
const command = (cmd) => {
  if (cmd == "logout") {
    clearUser()
    clearPages()
    clearResources()
    router.push("/login")
  }
}
const clickTab = (tab) => router.push(tab.paneName)
const removeTab = (name) => {
  removePage(name)
  router.push("/")
}
onMounted(() => {
  locale.value = web.locale
})
</script>
