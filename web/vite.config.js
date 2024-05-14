import { resolve } from "path"
import { defineConfig } from "vite"
import vue from "@vitejs/plugin-vue"
import VueSetupExtend from "vite-plugin-vue-setup-extend"

export default defineConfig({
  build: {
    chunkSizeWarningLimit: 4096,
  },
  plugins: [vue(), VueSetupExtend()],
  resolve: {
    alias: {
      "@": resolve(__dirname, "src"),
    },
  },
})
