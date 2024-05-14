<template>
  <div class="background center">
    <el-form ref="form" :model="data" :rules="rules">
      <el-form-item prop="userName">
        <el-input v-model="data.userName" maxlength="60" prefix-icon="user" @keyup.enter="login"></el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input v-model="data.password" maxlength="60" show-password prefix-icon="lock" @keyup.enter="login"></el-input>
      </el-form-item>
      <el-form-item>
        <div class="row-center">
          <el-button type="primary" @click="login">{{ $t("l.Login") }}</el-button>
        </div>
      </el-form-item>
    </el-form>
  </div>
</template>

<script setup name="login">
import { reactive, toRefs } from "vue"
import { useI18n } from "vue-i18n"
import { useRouter } from "vue-router"
import { useGet, usePost } from "@/http"
import { useStore } from "@/store"

const { t } = useI18n()
const router = useRouter()
const { saveResources, saveUser } = useStore()
const state = reactive({
  form: null,
  data: {
    userName: "admin",
    password: "",
  },
  rules: {
    userName: {
      required: true,
      message: t("r.Input").replace("$v", t("l.UserName").toLowerCase()),
      trigger: "blur",
    },
    password: {
      required: true,
      message: t("r.Input").replace("$v", t("l.Password").toLowerCase()),
      trigger: "blur",
    },
  },
})
const login = () => {
  state.form.validate(async (valid) => {
    if (valid) {
      const user = await usePost("/mgt/login", state.data)
      if (user) {
        saveUser({
          userName: user.userName,
          token: user.token,
        })
        saveResources(await useGet("/mgt/sys/resources/menu"))
        router.push("/")
      }
    }
  })
}

const { form, data, rules } = toRefs(state)
</script>
