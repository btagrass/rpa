<template>
  <el-form>
    <el-form-item :label="$t('l.Roles')">
      <el-checkbox-group v-model="data">
        <el-checkbox v-for="r in roles" :key="r.id" :label="r.id">{{ r.name }}</el-checkbox>
      </el-checkbox-group>
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
  data: {},
  roles: [],
})
const save = async () => {
  await usePost(`${api}/${state.id}/roles`, state.data)
  emits("close")
}
onMounted(async () => {
  state.roles = await useGet("/mgt/sys/roles")
  state.data = await useGet(`${api}/${state.id}/roles`)
})

const { data, roles } = toRefs(state)
</script>
