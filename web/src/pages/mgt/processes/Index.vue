<template>
  <div>
    <div class="row">
      <el-button type="danger" icon="Delete" @click="remove">{{ $t("l.Delete") }}</el-button>
      <el-button type="primary" icon="DocumentAdd" @click="open(Edit)">{{ $t("l.Add") }}</el-button>
      <el-button type="primary" icon="Search" @click="list">{{ $t("l.Search") }}</el-button>
    </div>
    <el-table ref="table" :data="data.records" border @selection-change="select">
      <el-table-column type="selection"></el-table-column>
      <el-table-column :label="$t('l.Id')" prop="id" width="120"></el-table-column>
      <el-table-column :label="$t('l.Name')" prop="name"></el-table-column>
      <el-table-column :label="$t('l.Operation')" width="150">
        <template #default="scope">
          <el-button-group>
            <el-button type="primary" icon="Edit" :title="$t('l.Edit')"
              @click="open(Edit, { id: scope.row.id })"></el-button>
            <el-button type="danger" icon="Delete" :title="$t('l.Delete')" @click="remove(scope.row)"></el-button>
            <el-button type="warning" icon="Setting" :title="$t('l.Steps')"
              @click="open(Steps, { id: scope.row.id })"></el-button>
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

<script setup name="processes">
import { useComp, useList } from "@/crud"
import Edit from "./Edit.vue"
import Steps from "./Steps.vue"

const { name, values, visible, open, close } = useComp()
const { table, query, data, list, remove, select } = useList("/mgt/processes")
</script>
