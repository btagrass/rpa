<template>
  <el-form ref="form" :model="data" :rules="rules">
    <el-form-item :label="$t('l.Id')" prop="id">
      <el-input v-model="data.id" disabled></el-input>
    </el-form-item>
    <el-form-item :label="$t('l.Name')" prop="name">
      <el-input v-model="data.name" disabled></el-input>
    </el-form-item>
    <el-form-item :label="$t('l.Desc')" prop="desc">
      <el-input v-model="data.desc" disabled></el-input>
    </el-form-item>
    <el-form-item :label="$t('l.Cron')" prop="cron">
      <el-input v-model="data.cron" clearable maxlength="100" show-word-limit></el-input>
    </el-form-item>
    <el-form-item :label="$t('l.Params')" prop="params">
      <el-input v-model="data.params" clearable maxlength="100" show-word-limit></el-input>
    </el-form-item>
    <el-form-item :label="$t('l.ParamsDesc')" prop="paramsDesc">
      <el-input v-model="data.paramsDesc" disabled></el-input>
    </el-form-item>
    <el-form-item :label="$t('l.Instance')" prop="instance">
      <el-input v-model="data.instance" disabled></el-input>
    </el-form-item>
    <el-form-item :label="$t('l.State')" prop="state">
      <el-select v-model="data.state" disabled>
        <el-option v-for="s in states" :key="s.code" :label="s.name" :value="s.code"></el-option>
      </el-select>
    </el-form-item>
    <el-form-item>
      <div class="row-center">
        <el-button type="primary" @click="save">{{ $t("l.Save") }}</el-button>
      </div>
    </el-form-item>
  </el-form>
</template>

<script setup>
import { useI18n } from "vue-i18n"
import { useEdit } from "@/crud"
import { useGet } from "@/http"

const { t } = useI18n()
const props = defineProps({
  values: Object,
})
const emits = defineEmits(["close"])
const { form, data, states, rules, save } = useEdit({
  ...props.values,
  states: [],
  rules: {
    cron: {
      required: true,
      message: t("r.Input").replace("$v", t("l.Cron").toLowerCase()),
      trigger: "blur",
    },
  },
}, emits, async () => {
  states.value = await useGet("/mgt/sys/dicts?type=JobState")
})
</script>
