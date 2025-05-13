<template>
    <el-card class="ControlPanel" shadow="never">
        <div class="CardHeader">
            <span> 控制面板</span>
            <span>当前时间片：{{ systemStatusStore.time }}</span>
        </div>
        <div class="button-container">
            <el-button type="primary" size="large" @click="showCreateDialog()">新建进程</el-button>
            <el-button type="primary" size="large">新建随机进程</el-button>
            <el-button type="primary" size="large">单步执行</el-button>
            <el-button type="primary" size="large">自动执行</el-button>
        </div>
    </el-card>

    <el-dialog v-model="dialogVisible" title="新建进程" width="500">
        <el-form :model="newProcess" :rules="rules" ref="formRef" label-width="120px" class="process-form">
            <el-form-item prop="name" label="进程名称">
                <el-input v-model="newProcess.name" placeholder="进程名称"></el-input>
            </el-form-item>
            <el-form-item prop="requiredTime" label="所需时间片">
                <el-input v-model.number="newProcess.requiredTime" placeholder="所需时间片"></el-input>
            </el-form-item>
            <el-form-item prop="priority" label="优先级">
                <el-input v-model.number="newProcess.priority" placeholder="优先级"></el-input>
            </el-form-item>
            <el-form-item prop="memorySize" label="内存大小">
                <el-input v-model.number="newProcess.memorySize" placeholder="内存大小"></el-input>
            </el-form-item>
        </el-form>
        <template #footer>
            <el-button type="success" size="large" @click="CreatNew()">新建进程</el-button>
        </template>
    </el-dialog>
</template>

<script lang="ts" setup>
import useSystemStatusStore from '@/store/modules/SystemStatus';
import { ref } from 'vue';
import type { FormInstance, FormRules } from 'element-plus';

const systemStatusStore = useSystemStatusStore();
const dialogVisible = ref(false);
const formRef = ref<FormInstance>();

const newProcess = ref({
    name: '',
    requiredTime: 0,
    priority: 0,
    memorySize: 0,
});

const rules = ref<FormRules>({
    name: [
        { required: true, message: '请输入进程名称', trigger: 'blur' },
    ],
    requiredTime: [
        { required: true, message: '请输入所需时间片', trigger: 'blur' },
        { type: 'number', min: 1, message: '时间片必须大于0', trigger: 'blur' }
    ],
    priority: [
        { required: true, message: '请输入优先级', trigger: 'blur' },
        { type: 'number', min: 1, message: '优先级必须大于0', trigger: 'blur' }
    ],
    memorySize: [
        { required: true, message: '请输入内存大小', trigger: 'blur' },
        { type: 'number', min: 1, message: '内存大小必须大于0', trigger: 'blur' }
    ],
});

const showCreateDialog = () => {
    dialogVisible.value = true;
};

const CreatNew = async () => {
    const isValid = await formRef.value?.validate();
    if (isValid) {
        await systemStatusStore.createNewProcess(newProcess.value);
        dialogVisible.value = false;
        newProcess.value = {
            name: '',
            requiredTime: 0,
            priority: 0,
            memorySize: 0,
        };
    }
};
</script>

<style lang="scss" scoped>
.CardHeader {
    font-size: 15px;
    opacity: 0.7;
    margin-bottom: 10px;
    display: flex;
    justify-content: space-between; // 两端对齐
    align-items: center; // 垂直居中
}

.ControlPanel {
    width: 25vw;
    height: 10vh;
}

.button-container {
    height: 100%;
    display: flex;
    align-items: center;
    /* 垂直居中 */
    gap: 10px;
    /* 按钮之间的间距 */
}

.process-form {
    .el-form-item {
        margin-bottom: 20px;

        :deep(.el-form-item__content) {
            width: calc(100% - 120px);

            .el-input {
                width: 100%;
            }
        }
    }
}

// 调整对话框内边距
:deep(.el-dialog__body) {
    padding: 20px 30px;
}
</style>