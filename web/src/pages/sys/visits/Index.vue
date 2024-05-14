<template>
  <div>
    <div class="row">
      <el-button type="danger" icon="Delete" @click="remove">{{ $t("l.Delete") }}</el-button>
      <el-button type="primary" icon="Search" @click="list">{{ $t("l.Search") }}</el-button>
      <el-button type="danger" icon="Search" @click="clear">{{ $t("l.Clear") }}</el-button>
    </div>
    <el-table ref="table" :data="data.records" border @selection-change="select">
      <el-table-column type="selection"></el-table-column>
      <el-table-column :label="$t('l.Id')" prop="id" width="120"></el-table-column>
      <el-table-column :label="$t('l.UserId')" prop="userId"></el-table-column>
      <el-table-column :label="$t('l.UserName')" prop="userName"></el-table-column>
      <el-table-column :label="$t('l.Ip')" prop="ip"></el-table-column>
      <el-table-column :label="$t('l.Method')" prop="method" align="center"></el-table-column>
      <el-table-column :label="$t('l.Url')" prop="url"></el-table-column>
      <el-table-column :label="$t('l.UserAgent')" prop="userAgent" show-overflow-tooltip></el-table-column>
      <el-table-column :label="$t('l.Req')" prop="req" show-overflow-tooltip></el-table-column>
      <el-table-column :label="$t('l.Resp')" prop="resp" show-overflow-tooltip></el-table-column>
      <el-table-column :label="$t('l.CreatedAt')" prop="createdAt"></el-table-column>
      <el-table-column :label="$t('l.Operation')" width="150">
        <template #default="scope">
          <el-button-group>
            <el-button type="danger" icon="Delete" :title="$t('l.Delete')" @click="remove(scope.row)"></el-button>
          </el-button-group>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination v-model:current-page="query.current" v-model:page-size="query.size" :total="data.total" background
      layout="total,prev,pager,next,sizes"></el-pagination>
  </div>
</template>

<script setup name="visits">
import { useList } from "@/crud"
import { useDelete } from "@/http"

const { table, query, data, list, remove, select } = useList("/mgt/sys/visits")
const clear = async () => {
  await useDelete("/mgt/sys/visits/all")
  list()
}
</script>
