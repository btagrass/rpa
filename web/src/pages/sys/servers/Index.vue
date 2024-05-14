<template>
  <div class="row-side">
    <div class="row">
      <el-icon>
        <Timer />
      </el-icon>
      <el-input-number v-model="interval" :min="0" :max="60" :step="3"></el-input-number>
      <el-button type="primary" icon="Search" @click="list">{{ $t("l.Search") }}</el-button>
    </div>
    <div class="row">
      <el-link type="primary" :href="`${env.VITE_MGT_URL}/debug/pprof`" target="_blank">{{ $t("l.debug") }}</el-link>
    </div>
  </div>
  <el-row :gutter="5">
    <el-col :span="12">
      <el-card>
        <template #header>
          <span>{{ $t("l.Hostname") }} ({{ data.host.hostname }})</span>
        </template>
        <el-row>
          <el-col :span="12">{{ $t("l.Uptime") }}</el-col>
          <el-col :span="12">{{ (data.host.uptime / 60 / 60 / 60).toFixed(2) }} {{ $t("l.Days").toLowerCase()
          }}</el-col>
        </el-row>
        <el-row>
          <el-col :span="12">{{ $t("l.Platform") }}</el-col>
          <el-col :span="12">{{ data.host.platform }} {{ data.host.platformVersion }}</el-col>
        </el-row>
        <el-row>
          <el-col :span="12">{{ $t("l.Arch") }}</el-col>
          <el-col :span="12">{{ data.host.kernelArch }}</el-col>
        </el-row>
      </el-card>
    </el-col>
    <el-col :span="12">
      <el-card>
        <template #header>
          <span>{{ $t("l.Cpu") }} ({{ data.cpu.modelName }})</span>
        </template>
        <el-row>
          <el-col :span="12">{{ $t("l.Cores") }}</el-col>
          <el-col :span="12">{{ data.cpu.cores }}</el-col>
        </el-row>
        <el-row>
          <el-col :span="12">{{ $t("l.Usage") }}</el-col>
          <el-col :span="12">{{ data.cpu.usage.toFixed(0) }} %</el-col>
        </el-row>
        <el-row>
          <el-col :span="12">{{ $t("l.Load") }}</el-col>
          <el-col :span="12">{{ (data.cpu.load * 100).toFixed(0) }} %</el-col>
        </el-row>
      </el-card>
    </el-col>
  </el-row>
  <el-row :gutter="5">
    <el-col :span="12">
      <el-card>
        <template #header>
          <span>{{ $t("l.Memory") }} ({{ data.memory.usedPercent.toFixed(0) }}%)</span>
        </template>
        <el-row>
          <el-col :span="12">{{ $t("l.Total") }}</el-col>
          <el-col :span="12">{{ (data.memory.total / 1024 / 1024 / 1024).toFixed(0) }} GB</el-col>
        </el-row>
        <el-row>
          <el-col :span="12">{{ $t("l.Available") }}</el-col>
          <el-col :span="12">{{ (data.memory.available / 1024 / 1024 / 1024).toFixed(0) }} GB</el-col>
        </el-row>
        <el-row>
          <el-col :span="12">{{ $t("l.Cache") }}</el-col>
          <el-col :span="12">{{ (data.memory.cached / 1024 / 1024 / 1024).toFixed(0) }} GB</el-col>
        </el-row>
      </el-card>
    </el-col>
    <el-col :span="12">
      <el-card>
        <template #header>
          <span>{{ $t("l.Network") }}</span>
        </template>
        <el-row>
          <el-col :span="12">{{ $t("l.BytesRecv") }}</el-col>
          <el-col :span="12">{{ (data.net.bytesRecv / 1024 / 1024).toFixed(1) }} MB</el-col>
        </el-row>
        <el-row>
          <el-col :span="12">{{ $t("l.BytesSent") }}</el-col>
          <el-col :span="12">{{ (data.net.bytesSent / 1024 / 1024).toFixed(1) }} MB</el-col>
        </el-row>
        <el-row>
          <el-col :span="12">{{ $t("l.PacketsRecv") }}</el-col>
          <el-col :span="12">{{ data.net.packetsRecv }}</el-col>
        </el-row>
        <el-row>
          <el-col :span="12">{{ $t("l.PacketsSent") }}</el-col>
          <el-col :span="12">{{ data.net.packetsSent }}</el-col>
        </el-row>
      </el-card>
    </el-col>
  </el-row>
  <el-row :gutter="5">
    <el-col :span="24">
      <el-card>
        <template #header>
          <span>{{ $t("l.Disk") }}</span>
        </template>
        <el-row>
          <el-col :span="4"></el-col>
          <el-col :span="4">{{ $t("l.Name") }}</el-col>
          <el-col :span="4">{{ $t("l.Mount") }}</el-col>
          <el-col :span="4">{{ $t("l.Total") }}</el-col>
          <el-col :span="4">{{ $t("l.Available") }}</el-col>
          <el-col :span="4">{{ $t("l.Usage") }}</el-col>
        </el-row>
        <el-divider></el-divider>
        <el-row v-for="d, i in data.disks" :key="d">
          <el-col :span="4">{{ i + 1 }}</el-col>
          <el-col :span="4">{{ d.device }}</el-col>
          <el-col :span="4">{{ d.mountpoint }}</el-col>
          <el-col :span="4">{{ (d.total / 1000 / 1000 / 1000).toFixed(0) }} GB</el-col>
          <el-col :span="4">{{ (d.free / 1000 / 1000 / 1000).toFixed(0) }} GB</el-col>
          <el-col :span="4">{{ d.usedPercent.toFixed(2) }} %</el-col>
        </el-row>
      </el-card>
    </el-col>
  </el-row>
</template>

<script setup name="servers">
import { reactive, toRefs, onMounted } from "vue"
import { useRefresh } from "@/crud"
import { useGet } from "@/http"

const env = import.meta.env
const state = reactive({
  data: {
    host: {
      hostname: "",
      uptime: 0,
      platform: "",
      kernelArch: "",
    },
    cpu: {
      model: "",
      cores: 0,
      usage: 0,
      load: 0,
    },
    memory: {
      usedPercent: 0,
      total: 0,
      available: 0,
      cached: 0,
    },
    net: {
      bytesSent: 0,
      bytesRecv: 0,
      packetsSent: 0,
      packetsRecv: 0,
    },
    disks: [],
  },
})
const list = async () => {
  state.data = await useGet("/mgt/sys/servers")
}
onMounted(list)
const { interval } = useRefresh(list)

const { data } = toRefs(state)
</script>
