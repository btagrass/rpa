<template>
  <el-form>
    <el-form-item :label="$t('l.Resources')">
      <el-tree ref="tree" :data="resources" :props="{ label: 'name' }" default-expand-all node-key="id"
        show-checkbox></el-tree>
    </el-form-item>
    <el-form-item>
      <div class="row-center">
        <el-button type="primary" @click="save">{{ $t("l.Save") }}</el-button>
      </div>
    </el-form-item>
  </el-form>
</template>

<script setup>
import { inject, reactive, toRefs, onMounted } from "vue"
import { useGet, usePost } from "@/http"

const props = defineProps({
  values: Object,
})
const emits = defineEmits(["close"])
const api = inject("api")
const state = reactive({
  ...props.values,
  tree: null,
  resources: [],
})
const save = async () => {
  await usePost(`${api}/${state.id}/resources`, state.tree.getCheckedNodes())
  emits("close")
}
onMounted(async () => {
  state.resources = await useGet("/mgt/sys/resources")
  state.tree.setCheckedKeys(await useGet(`${api}/${state.id}/resources`))
})

const { tree, resources } = toRefs(state)
</script>
