<template>
  <el-form ref="form" :model="data" :rules="rules">
    <el-form-item :label="$t('l.Id')" prop="id">
      <el-input v-model="data.id" disabled></el-input>
    </el-form-item>
    <el-form-item :label="$t('l.ParentId')" prop="parentId">
      <el-input-number v-model="data.parentId" :min="0"></el-input-number>
    </el-form-item>
    <el-form-item :label="$t('l.Name')" prop="name">
      <el-input v-model="data.name" clearable maxlength="100" show-word-limit></el-input>
    </el-form-item>
    <el-form-item :label="$t('l.Type')" prop="type">
      <el-select v-model="data.type" filterable>
        <el-option v-for="t in types" :key="t.code" :label="t.name" :value="t.code"></el-option>
      </el-select>
    </el-form-item>
    <el-form-item :label="$t('l.Icon')" prop="icon">
      <el-input v-model="data.icon" clearable maxlength="100" placeholder="ElementPlus 图标" show-word-limit></el-input>
    </el-form-item>
    <el-form-item :label="$t('l.Uri')" prop="uri">
      <el-input v-model="data.uri" clearable maxlength="100" show-word-limit></el-input>
    </el-form-item>
    <el-form-item :label="$t('l.Act')" prop="act">
      <el-select v-model="data.act" clearable filterable>
        <el-option v-for="a in acts" :key="a.code" :label="a.name" :value="a.code"></el-option>
      </el-select>
    </el-form-item>
    <el-form-item :label="$t('l.Sequence')" prop="sequence">
      <el-input-number v-model="data.sequence" :min="0"></el-input-number>
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
import { useGet } from "@/http"

const { t } = useI18n()
const props = defineProps({
  values: Object,
})
const emits = defineEmits(["close"])
const { form, data, types, acts, rules, save } = useEdit({
  ...props.values,
  types: [],
  acts: [],
  rules: {
    name: {
      required: true,
      message: t("r.Input").replace("$v", t("l.Name").toLowerCase()),
      trigger: "blur",
    },
    type: {
      required: true,
      message: t("r.Select").replace("$v", t("l.Type").toLowerCase()),
      trigger: "blur",
    },
    uri: {
      required: true,
      message: t("r.Input").replace("$v", t("l.Uri").toLowerCase()),
      trigger: "blur",
    },
  },
}, emits, async () => {
  types.value = await useGet("/mgt/sys/dicts?type=ResourceType")
  acts.value = await useGet("/mgt/sys/dicts?type=ResourceAct")
})
</script>
