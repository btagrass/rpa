import { provide, inject, reactive, ref, toRefs, watch, onMounted, onActivated, onDeactivated } from "vue"
import { ElMessage, ElMessageBox } from "element-plus"
import { useGet, usePost, useDelete } from "@/http"

export function useComp() {
  const state = reactive({
    name: "",
    values: {},
    visible: false,
  })
  const open = (name, values) => {
    state.name = name
    state.values = values
    state.visible = true
  }
  const close = () => {
    state.name = ""
    state.values = {}
    state.visible = false
  }
  return {
    ...toRefs(state),
    open,
    close,
  }
}

export function useList(api, state, mounted) {
  provide("api", api)
  const s = reactive({
    table: null,
    ids: [],
    query: {
      current: 1,
      size: 10,
    },
    data: {
      records: [],
      total: 0,
    },
    ...state,
  })
  const list = async () => {
    s.data = await useGet(api, {
      params: s.query,
    })
  }
  const remove = (selection) => {
    if (!(selection instanceof PointerEvent)) {
      select(selection)
    }
    if (s.ids.length > 0) {
      ElMessageBox.confirm("确定要删除吗？")
        .then(async () => {
          await useDelete(`${api}/${s.ids}`)
          s.query.current = 1
          list()
        })
        .catch(() => {
          s.table.clearSelection()
          s.ids = []
        })
    } else {
      ElMessage.error("请选择记录")
    }
  }
  const select = (selection) => {
    if (selection instanceof Array) {
      s.ids = selection.map((s) => s.id)
    } else if (selection instanceof Object) {
      if (selection.children) {
        selection.children.forEach((s) => {
          select(s)
        })
      }
      s.ids.push(selection.id)
    }
  }
  watch(s.query, list)
  onMounted(() => {
    list()
    if (mounted) {
      mounted()
    }
  })
  return {
    ...toRefs(s),
    list,
    remove,
    select,
  }
}

export function useEdit(state, emits, mounted) {
  const api = inject("api")
  const s = reactive({
    form: null,
    id: 0,
    data: {},
    rules: {},
    ...state,
  })
  const edit = async (id) => {
    id = id ?? s.id
    if (id) {
      s.data = await useGet(`${api}/${id}`)
      if (!s.data) {
        ElMessage.error("该记录不存在")
      }
    } else {
      s.data = { ...state, id: 0 }
    }
  }
  const save = (id) => {
    s.form.validate(async (valid) => {
      if (valid) {
        await usePost(api, s.data)
        if (id instanceof MouseEvent) {
          if (emits) {
            emits("close")
          }
        } else {
          edit(id)
        }
      }
    })
  }
  onMounted(() => {
    edit()
    if (mounted) {
      mounted()
    }
  })
  return {
    api,
    ...toRefs(s),
    save,
  }
}

export function useRefresh(func) {
  const id = ref(0)
  const interval = ref(0)
  watch(interval, () => {
    clearInterval(id.value)
    if (interval.value > 0) {
      id.value = setInterval(func, interval.value * 1000)
    }
  })
  onActivated(() => {
    if (interval.value > 0) {
      id.value = setInterval(func, interval.value * 1000)
    }
  })
  onDeactivated(() => {
    clearInterval(id.value)
  })
  return {
    id,
    interval,
  }
}
