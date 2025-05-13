<template>
    <el-card class="ProcessPanel" shadow="never">
        <div class="CardHeader">
            <span> 处理机(CPU)状态</span>
        </div>
        <div class="mainContainer">
            <p>
                <span> CPU0:{{ CPU0State }}</span>
                <el-button type="primary" size="small" :icon="View" @click="showCPU(0)"></el-button>
            </p>
            <p>
                <span> CPU1:{{ CPU1State }}</span>
                <el-button type="primary" size="small" :icon="View" @click="showCPU(1)"></el-button>
            </p>
        </div>
    </el-card>

    <el-dialog v-model="dialogVisible" title="当前CPU上进程：" width="500">
        <template v-if="currentCPUProcess">
            <div v-for="(value, key) in currentCPUProcess" :key="key" class="process-info-item">
                <span class="info-label">{{ key }}:</span>
                <span class="info-value">{{ value ?? 'null' }}</span>
            </div>
        </template>
    </el-dialog>
</template>

<script lang="ts" setup>
import useSystemStatusStore from '@/store/modules/SystemStatus';
import { computed, ref } from 'vue';
import { View } from '@element-plus/icons-vue';
import { ElMessage } from 'element-plus';

const systemStatusStore = useSystemStatusStore();
const dialogVisible = ref(false);
const currentCPUIndex = ref<number>(0);

const CPU0 = computed(() => systemStatusStore.ProcessorsStatus[0]);
const CPU1 = computed(() => systemStatusStore.ProcessorsStatus[1]);

const CPU0State = computed(() => CPU0.value === null ? '空闲' : '占用');
const CPU1State = computed(() => CPU1.value === null ? '空闲' : '占用');

const showCPU = (cpuIndex: number) => {
    console.log('showCPU', cpuIndex);
    const cpu = cpuIndex === 0 ? CPU0 : CPU1;
    if (cpu.value !== null) {
        currentCPUIndex.value = cpuIndex;
        dialogVisible.value = true;
    } else {
        ElMessage.warning('当前CPU空闲');
    }
};

const currentCPUProcess = computed(() => {
    const cpu = currentCPUIndex.value === 0 ? CPU0.value : CPU1.value;
    return cpu;
});
</script>

<style lang="scss" scoped>
.CardHeader {
    font-size: 14px;
    opacity: 0.8;
    margin-bottom: 10px;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.ProcessPanel {
    width: 25vw;
    height: 10vh;
}

.mainContainer {
    p {
        margin: 10px auto;

        span {
            font-size: 18px;
            margin-right: 20px;
        }
    }
}

.process-info-item {
    margin: 10px 0;
    display: flex;
    align-items: center;

    .info-label {
        font-weight: bold;
        width: 120px;
        margin-right: 10px;
    }

    .info-value {
        flex: 1;
    }
}
</style>
