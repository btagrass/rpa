<template>
  <div>
    <div class="row">
      <el-button type="danger" icon="Delete" @click="remove">{{ $t("l.Delete") }}</el-button>
      <el-select v-model="query.type" filterable :placeholder="$t('l.Type')">
        <el-option v-for="t in types" :key="t.code" :label="t.name" :value="t.code"></el-option>
      </el-select>
      <el-input v-model="query.keyword" clearable :placeholder="$t('l.Key')"></el-input>
      <el-button type="primary" icon="Search" @click="list">{{ $t("l.Search") }}</el-button>
    </div>
    <el-table ref="table" :data="data.records" border @selection-change="select">
      <el-table-column type="selection"></el-table-column>
      <el-table-column :label="$t('l.Id')" prop="id" width="60"></el-table-column>
      <el-table-column :label="$t('l.Key')" prop="key"></el-table-column>
      <el-table-column :label="$t('l.Value')" prop="value"></el-table-column>
      <el-table-column :label="$t('l.Expiration')" prop="expiration" width="60"></el-table-column>
      <el-table-column :label="$t('l.Operation')" width="150">
        <template #default="scope">
          <el-button-group>
            <el-button type="danger" icon="Delete" :title="$t('l.Delete')" @click="remove(scope.row)"></el-button>
          </el-button-group>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script setup name="caches">
import { useList } from "@/crud"
import { useGet } from "@/http"

const { table, query, data, types, list, remove, select } = useList("/mgt/sys/caches", {
  query: {
    type: 1,
    keyword: "",
  },
  types: [],
}, async () => {
  types.value = await useGet("/mgt/sys/dicts?type=CacheType")
})
</script>
