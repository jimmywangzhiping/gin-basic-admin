<template>
  <div>
    <div class="gva-form-box">
      <el-form
        ref="elFormRef"
        :model="formData"
        label-position="right"
        :rules="rule"
        label-width="80px"
      >
        <el-form-item
          label="远程服务器Id:"
          prop="remoteId"
        >
          <el-input
            v-model.number="formData.remoteId"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item
          label="命令名称:"
          prop="command"
        >
          <el-input
            v-model="formData.command"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item
          label="执行命令结果:"
          prop="status"
        >
          <el-switch
            v-model="formData.status"
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="是"
            inactive-text="否"
            clearable
          />
        </el-form-item>
        <el-form-item
          label="消息:"
          prop="message"
        >
          <el-input
            v-model="formData.message"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            @click="save"
          >保存</el-button>
          <el-button
            type="primary"
            @click="back"
          >返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createSysRemoteRecord,
  updateSysRemoteRecord,
  findSysRemoteRecord
} from '@/api/sysRemoteRecord'

defineOptions({
  name: 'SysRemoteRecordForm'
})

// 自动获取字典
// import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
  remoteId: 0,
  command: '',
  status: false,
  message: '',
})
// 验证规则
const rule = reactive({
  remoteId: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  command: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
})

const elFormRef = ref()

// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findSysRemoteRecord({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.reremoteRecord
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
}

init()
// 保存按钮
const save = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createSysRemoteRecord(formData.value)
        break
      case 'update':
        res = await updateSysRemoteRecord(formData.value)
        break
      default:
        res = await createSysRemoteRecord(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
    }
  })
}

// 返回按钮
const back = () => {
  router.go(-1)
}

</script>

<style>
</style>
