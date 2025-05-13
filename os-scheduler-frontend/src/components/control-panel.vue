<template>
    <el-card class="ControlPanel" shadow="never">
        <div class="CardHeader">
            <span> 控制面板</span>
            <span>当前时间片：{{ systemStatusStore.time }}</span>
        </div>
        <div class="button-container">
            <el-button type="primary" size="large" @click="showCreateDialog()">新建进程</el-button>
            <el-button type="primary" size="large">新建随机进程</el-button>
            <el-button type="primary" size="large" @click="SingleMove()">单步执行</el-button>
            <el-button type="primary" size="large">自动执行</el-button>
            <el-button type="danger" size="large" @click="reset()">重置系统</el-button>
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
            <el-form-item prop="predecessors" label="前驱进程">
                <el-input v-model="predecessorsInput" placeholder="请输入进程ID，多个用逗号分隔"
                    @change="handlePredecessorsChange"></el-input>
            </el-form-item>
            <el-form-item prop="successors" label="后继进程">
                <el-input v-model="successorsInput" placeholder="请输入进程ID，多个用逗号分隔"
                    @change="handleSuccessorsChange"></el-input>
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
import { ElMessage, type FormInstance, type FormRules } from 'element-plus';

const systemStatusStore = useSystemStatusStore();
const dialogVisible = ref(false);
const formRef = ref<FormInstance>();

const newProcess = ref<{
    name: string;
    requiredTime: number;
    priority: number;
    memorySize: number;
    predecessors: number[] | null;
    successors: number[] | null;
}>({
    name: '',
    requiredTime: 0,
    priority: 0,
    memorySize: 0,
    predecessors: null,
    successors: null,
});

const predecessorsInput = ref('');
const successorsInput = ref('');

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
    predecessors: [
        {
            validator: (rule, value, callback) => {
                if (predecessorsInput.value.trim() === '') {
                    callback();
                } else {
                    const numbers = predecessorsInput.value.split(',')
                        .map(n => parseInt(n.trim()));
                    if (numbers.some(n => isNaN(n))) {
                        callback(new Error('请输入有效的进程ID'));
                    } else {
                        callback();
                    }
                }
            },
            trigger: 'blur'
        }
    ],
    successors: [
        {
            validator: (rule, value, callback) => {
                if (successorsInput.value.trim() === '') {
                    callback();
                } else {
                    const numbers = successorsInput.value.split(',')
                        .map(n => parseInt(n.trim()));
                    if (numbers.some(n => isNaN(n))) {
                        callback(new Error('请输入有效的进程ID'));
                    } else {
                        callback();
                    }
                }
            },
            trigger: 'blur'
        }
    ]
});

const showCreateDialog = () => {
    dialogVisible.value = true;
};

const handlePredecessorsChange = (value: string) => {
    if (!value.trim()) {
        newProcess.value.predecessors = null;
        return;
    }
    const numbers = value.split(',')
        .map(n => parseInt(n.trim()))
        .filter(n => !isNaN(n));
    newProcess.value.predecessors = numbers.length ? numbers : null;
};

const handleSuccessorsChange = (value: string) => {
    if (!value.trim()) {
        newProcess.value.successors = null;
        return;
    }
    const numbers = value.split(',')
        .map(n => parseInt(n.trim()))
        .filter(n => !isNaN(n));
    newProcess.value.successors = numbers.length ? numbers : null;
};

const resetForm = () => {
    newProcess.value = {
        name: '',
        requiredTime: 0,
        priority: 0,
        memorySize: 0,
        predecessors: null,
        successors: null,
    };
    predecessorsInput.value = '';
    successorsInput.value = '';
};

const CreatNew = async () => {
    try {
        const isValid = await formRef.value?.validate();
        if (isValid) {
            // 等待创建进程操作完成
            await systemStatusStore.createNewProcess(newProcess.value);
            // 重置表单
            resetForm();
            // 关闭对话框
            dialogVisible.value = false;
        }
    } catch (error) {
        console.error('创建进程失败:', error);
        ElMessage.error('创建进程失败');
    }
};

const reset = async () => {
    console.log('重置系统');
    await systemStatusStore.reset();
};

const SingleMove = async () => {
    await systemStatusStore.Schedule();
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
    width: 100vw;
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