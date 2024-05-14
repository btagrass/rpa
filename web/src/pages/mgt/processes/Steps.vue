<template>
  <div>
    <div class="row">
      <el-button type="primary" icon="DocumentAdd" @click="add">{{ $t("l.Add") }}</el-button>
      <el-button type="primary" icon="Search" @click="list">{{ $t("l.Search") }}</el-button>
    </div>
    <el-tree ref="table" :data="data" default-expand-all draggable node-key="id" @node-click="edit" @node-drop="dragDrop">
      <template #default="{ data }">
        <div class="row-side">
          <span>{{ data.type }}</span>
          <el-button-group>
            <el-button type="danger" icon="Delete" :title="$t('l.Delete')" @click="remove(data)"></el-button>
            <el-button type="primary" icon="DocumentAdd" :title="$t('l.AddSibling')"
              @click="open(Edit, { parentId: data.parentId })"></el-button>
            <el-button type="warning" icon="DocumentAdd" :title="$t('l.AddChild')"
              @click="open(Edit, { parentId: data.id })"></el-button>
          </el-button-group>
        </div>
      </template>
    </el-tree>
    <el-form ref="form" :model="datum" :rules="rules">
      <el-form-item :label="$t('l.Id')" prop="id">
        <el-input v-model="datum.id" disabled></el-input>
      </el-form-item>
      <el-form-item :label="$t('l.ParentId')" prop="parentId">
        <el-input v-model="datum.parentId" disabled></el-input>
      </el-form-item>
      <el-form-item :label="$t('l.Type')" prop="type">
        <el-select v-model="datum.type" filterable>
          <el-option v-for="t in state.types" :key="t.code" :label="t.name" :value="t.name"></el-option>
        </el-select>
      </el-form-item>
      <template v-if="datum.type == 'Wait'">
        <el-form-item :label="$t('l.Duration')" prop="value">
          <el-input v-model="datum.value" clearable maxlength="100" show-word-limit></el-input>
        </el-form-item>
      </template>
      <template v-if="datum.type == '[Web]Goto'">
        <el-form-item :label="$t('l.Url')" prop="name">
          <el-input v-model="datum.name" clearable maxlength="100" show-word-limit></el-input>
        </el-form-item>
        <el-form-item :label="$t('l.Visible')" prop="visible">
          <el-checkbox v-model="datum.visible"></el-checkbox>
        </el-form-item>
        <el-form-item :label="$t('l.Timeout')" prop="duration">
          <el-input-number v-model="datum.duration" :min="0"></el-input-number>
        </el-form-item>
      </template>
      <template v-if="datum.type == '[Web]Attr'">
        <el-form-item :label="$t('l.Name')" prop="name">
          <el-input v-model="datum.name" clearable maxlength="100" show-word-limit></el-input>
        </el-form-item>
        <el-form-item :label="$t('l.Attr')" prop="value">
          <el-input v-model="datum.value" clearable maxlength="100" show-word-limit></el-input>
        </el-form-item>
      </template>
      <template v-if="datum.type == '[Web]Click'">
        <el-form-item :label="$t('l.Name')" prop="name">
          <el-input v-model="datum.name" clearable maxlength="100" show-word-limit></el-input>
        </el-form-item>
      </template>
      <template v-if="datum.type == '[Web]Input'">
        <el-form-item :label="$t('l.Name')" prop="name">
          <el-input v-model="datum.name" clearable maxlength="100" show-word-limit></el-input>
        </el-form-item>
        <el-form-item :label="$t('l.Value')" prop="value">
          <el-input v-model="datum.value" clearable maxlength="100" show-word-limit></el-input>
        </el-form-item>
      </template>
      <template v-if="datum.type == '[Web]Text'">
        <el-form-item :label="$t('l.Name')" prop="name">
          <el-input v-model="datum.name" clearable maxlength="100" show-word-limit></el-input>
        </el-form-item>
      </template>
      <el-form-item :label="$t('l.Variable')" prop="variable">
        <el-input v-model="datum.variable" clearable maxlength="100" show-word-limit></el-input>
      </el-form-item>
      <el-form-item>
        <div class="row-center">
          <el-button type="primary" @click="save">{{ $t("l.Save") }}</el-button>
          <el-button type="primary" @click="save(0)">{{ $t("l.SaveAdd") }}</el-button>
        </div>
      </el-form-item>
    </el-form>
  </div>
</template>

<script setup>
import { reactive, onMounted } from "vue"
import { useList } from "@/crud"
import { useGet, usePost } from "@/http"
import { toRefs } from "vue";

const props = defineProps({
  values: Object,
})
const { data, list, remove } = useList(`/mgt/processes/${props.values.id}/steps`)
const dragDrop = async (dragNode, dropNode) => {
  dragNode.data.parentId = dropNode.data.id
  await usePost(`/mgt/processes/${props.values.id}/steps`, dragNode.data)
}
const state = reactive({
  form: null,
  id: 0,
  datum: {},
  rules: {},
  types: [],
})
const add = () => {
  state.datum = {
    processId: props.values.id,
  }
}
const edit = (data) => {
  state.datum = data
}
const save = (id) => {
  state.form.validate(async (valid) => {
    if (valid) {
      await usePost(`/mgt/processes/${props.values.id}/steps`, state.datum)
      if (id instanceof MouseEvent) {
      } else {
        edit(id)
      }
    }
  })
}
onMounted(async () => {
  state.types = await useGet("/mgt/sys/dicts?type=ProcessType")
})

const { form, datum } = toRefs(state)
</script>
