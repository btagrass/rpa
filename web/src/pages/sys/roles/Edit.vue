<template>
  <el-form ref="form" :model="data" :rules="rules">
    <el-form-item :label="$t('l.Id')" prop="id">
      <el-input v-model="data.id" disabled></el-input>
    </el-form-item>
    <el-form-item :label="$t('l.Name')" prop="name">
      <el-input v-model="data.name" clearable maxlength="100" show-word-limit></el-input>
    </el-form-item>
    <el-form-item>
      <div class="row-center">
        <el-button type="primary" @click="save">{{ $t("l.Save") }}</el-button>
        <el-button type="primary" @click="save(0)">{{ $t("l.SaveAdd") }}</el-button>
      </div>
    </el-form-item>
  </el-form>
</template>

<script setup>
import { useI18n } from "vue-i18n"
import { useEdit } from "@/crud"

const { t } = useI18n()
const props = defineProps({
  values: Object,
})
const emits = defineEmits(["close"])
const { form, data, rules, save } = useEdit({
  ...props.values,
  rules: {
    name: {
      required: true,
      message: t("r.Input").replace("$v", t("l.Name").toLowerCase()),
      trigger: "blur",
    },
  },
}, emits)
</script>
