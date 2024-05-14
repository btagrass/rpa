<template>
  <el-form ref="form" :model="data" :rules="rules">
    <el-form-item :label="$t('l.Id')" prop="id">
      <el-input v-model="data.id" disabled></el-input>
    </el-form-item>
    <el-form-item :label="$t('l.Dept')" prop="deptId">
      <el-tree-select v-model="data.deptId" :data="depts" :props="{ label: 'name' }" check-strictly clearable
        node-key="id">
      </el-tree-select>
    </el-form-item>
    <el-form-item :label="$t('l.UserName')" prop="userName">
      <el-input v-model="data.userName" clearable maxlength="60" show-word-limit></el-input>
    </el-form-item>
    <el-form-item :label="$t('l.FullName')" prop="fullName">
      <el-input v-model="data.fullName" clearable maxlength="100" show-word-limit></el-input>
    </el-form-item>
    <el-form-item :label="$t('l.Mobile')" prop="mobile">
      <el-input v-model="data.mobile" clearable maxlength="100" show-word-limit></el-input>
    </el-form-item>
    <el-form-item :label="$t('l.Password')" prop="password">
      <el-input v-model="data.password" clearable maxlength="60" show-password></el-input>
    </el-form-item>
    <el-form-item :label="$t('l.Frozen')" prop="frozen">
      <el-switch v-model="data.frozen"></el-switch>
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
const { form, data, depts, rules, save } = useEdit({
  ...props.values,
  depts: [],
  rules: {
    deptId: {
      required: true,
      message: t("r.Select").replace("$v", t("l.Dept").toLowerCase()),
      trigger: "blur",
    },
    userName: {
      required: true,
      message: t("r.Input").replace("$v", t("l.UserName").toLowerCase()),
      trigger: "blur",
    },
    mobile: {
      required: true,
      pattern: "^[1][3-9][0-9]{9}$",
      message: t("r.InputCorrect").replace("$v", t("l.Mobile").toLowerCase()),
      trigger: "blur",
    },
    password: {
      required: true,
      message: t("r.Input").replace("$v", t("l.Password").toLowerCase()),
      trigger: "blur",
    },
  },
}, emits, async () => {
  depts.value = await useGet("/mgt/sys/depts")
})
</script>
