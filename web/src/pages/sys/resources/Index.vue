<template>
  <div>
    <div class="row">
      <el-button type="primary" icon="Search" @click="list">{{ $t("l.Search") }}</el-button>
    </div>
    <el-table ref="table" :data="data.records" :tree-props="{ children: 'children' }" border default-expand-all
      row-key="id">
      <el-table-column :label="$t('l.Id')" prop="id" width="170"></el-table-column>
      <el-table-column :label="$t('l.Code')" prop="code"></el-table-column>
      <el-table-column :label="$t('l.Name')" prop="name"></el-table-column>
      <el-table-column :label="$t('l.Type')" align="center" :formatter="format"></el-table-column>
      <el-table-column :label="$t('l.Icon')" prop="icon"></el-table-column>
      <el-table-column :label="$t('l.Uri')" prop="uri"></el-table-column>
      <el-table-column :label="$t('l.Act')" prop="act" align="center"></el-table-column>
      <el-table-column :label="$t('l.Sequence')" prop="sequence" align="center"></el-table-column>
      <el-table-column :label="$t('l.Operation')" width="170">
        <template #default="scope">
          <el-button-group>
            <el-button type="primary" icon="Edit" :title="$t('l.Edit')"
              @click="open(Edit, { id: scope.row.id, parentId: scope.row.parentId })"></el-button>
            <el-button type="danger" icon="Delete" :title="$t('l.Delete')" @click="remove(scope.row)"></el-button>
            <el-button type="primary" icon="DocumentAdd" :title="$t('l.AddSibling')"
              @click="open(Edit, { parentId: scope.row.parentId })"></el-button>
            <el-button type="warning" icon="DocumentAdd" :title="$t('l.AddChild')"
              @click="open(Edit, { parentId: scope.row.id })"></el-button>
          </el-button-group>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination v-model:current-page="query.current" v-model:page-size="query.size" :total="data.total" background
      layout="total,prev,pager,next,sizes"></el-pagination>
    <el-drawer v-model="visible" destroy-on-close @close="list">
      <component :is="name" :values="values" @close="close"></component>
    </el-drawer>
  </div>
</template>

<script setup name="resources">
import { useComp, useList } from "@/crud"
import { useGet } from "@/http"
import Edit from "./Edit.vue"

const { name, values, visible, open, close } = useComp()
const { table, query, data, types, list, remove } = useList("/mgt/sys/resources", {
  types: [],
}, async () => {
  types.value = await useGet("/mgt/sys/dicts?type=ResourceType")
})
const format = (row) => {
  const type = types.value.find((i) => i.code == row.type)
  return type ? type.name : row.type
}
</script>
