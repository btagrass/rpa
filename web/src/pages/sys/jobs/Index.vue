<template>
  <div>
    <div class="row">
      <el-button type="primary" icon="Search" @click="list">{{ $t("l.Search") }}</el-button>
    </div>
    <el-table ref="table" :data="data.records" border>
      <el-table-column :label="$t('l.Id')" prop="id" width="120"></el-table-column>
      <el-table-column :label="$t('l.Name')" prop="name"></el-table-column>
      <el-table-column :label="$t('l.Desc')" prop="desc"></el-table-column>
      <el-table-column :label="$t('l.Cron')" prop="cron"></el-table-column>
      <el-table-column :label="$t('l.Params')" prop="params"></el-table-column>
      <el-table-column :label="$t('l.ParamsDesc')" prop="paramsDesc"></el-table-column>
      <el-table-column :label="$t('l.Instance')" prop="instance" align="center"></el-table-column>
      <el-table-column :label="$t('l.State')" prop="state" align="center" :formatter="format"></el-table-column>
      <el-table-column :label="$t('l.UpdatedAt')" width="130">
        <template #default="scope">{{ useDateFormat(scope.row.updatedAt, "YYYY-MM-DD HH:mm:ss").value }}</template>
      </el-table-column>
      <el-table-column :label="$t('l.Operation')" width="150">
        <template #default="scope">
          <el-button-group>
            <el-button type="primary" icon="Edit" :title="$t('l.Edit')"
              @click="open(Edit, { id: scope.row.id })"></el-button>
            <el-button v-if="scope.row.state == 0" type="warning" icon="VideoPlay" :title="$t('l.Start')"
              @click="start(scope.row.id)"></el-button>
            <el-button v-else type="danger" icon="VideoPause" :title="$t('l.Stop')"
              @click="stop(scope.row.id)"></el-button>
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

<script setup name="jobs">
import { useDateFormat } from "@vueuse/core"
import { useComp, useList } from "@/crud"
import { useGet, usePost } from "@/http"
import Edit from "./Edit.vue"

const { name, values, visible, open, close } = useComp()
const { table, query, data, states, list } = useList("/mgt/sys/jobs", {
  states: [],
}, async () => {
  states.value = await useGet("/mgt/sys/dicts?type=JobState")
})
const format = (row) => {
  const state = states.value.find((i) => i.code == row.state)
  return state ? state.name : row.state
}
const start = async (id) => {
  await usePost(`/mgt/sys/jobs/${id}/start`)
  list()
}
const stop = async (id) => {
  await usePost(`/mgt/sys/jobs/${id}/stop`)
  list()
}
</script>
