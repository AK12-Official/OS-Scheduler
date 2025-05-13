<template>
    <el-card class="QueuePanel">
        <div class="header">
            <el-select v-model="currentQueue" class="queue-select" @change="updateProcessArray">
                <el-option label="就绪队列" value="ready" />
                <el-option label="运行队列" value="running" />
                <el-option label="等待队列" value="waiting" />
                <el-option label="后备队列" value="backup" />
                <el-option label="挂起队列" value="suspended" />
            </el-select>
        </div>
        <el-table :data="sortedProcessArray" stripe style="width: 100%" border>
            <el-table-column prop="pid" label="pid" width="50px" />
            <el-table-column prop="name" label="进程名" />
            <el-table-column prop="priority" label="优先级" />
            <el-table-column prop="requiredTime" label="剩余时间" />
            <el-table-column prop="state" label="状态" />
            <el-table-column prop="memoryStart" label="内存位置" />
            <el-table-column prop="memorySize" label="内存大小" />
            <el-table-column label="类型">
                <template #default="{ row }">
                    {{ getProcessType(row) }}
                </template>
            </el-table-column>
        </el-table>
    </el-card>
</template>

<script lang="ts" setup>
import { ref, computed, watch } from 'vue';
import useSystemStatusStore from '@/store/modules/SystemStatus';
import type { Process } from '@/types';

const systemStatusStore = useSystemStatusStore();
const currentQueue = ref('ready');
const ProcessArray = ref(systemStatusStore.SystemStatus.queue.ready);

// 监视 SystemStatus 的变化
watch(
    () => systemStatusStore.SystemStatus,
    () => {
        updateProcessArray();
    },
    { deep: true } // 启用深度监听以检测嵌套对象的变化
);

// 添加计算属性，按pid排序
const sortedProcessArray = computed(() => {
    return [...ProcessArray.value].sort((a, b) => a.pid - b.pid);
});

const updateProcessArray = () => {
    switch (currentQueue.value) {
        case 'ready':
            ProcessArray.value = systemStatusStore.SystemStatus.queue.ready;
            break;
        case 'running':
            ProcessArray.value = systemStatusStore.SystemStatus.queue.running;
            break;
        case 'waiting':
            ProcessArray.value = systemStatusStore.SystemStatus.queue.waiting;
            break;
        case 'backup':
            ProcessArray.value = systemStatusStore.SystemStatus.queue.backup;
            break;
        case 'suspended':
            ProcessArray.value = systemStatusStore.SystemStatus.queue.suspended;
            break;
    }
};

// 判断进程类型的函数
const getProcessType = (process: Process) => {
    return (!process.predecessors && !process.successors) ? '独立进程' : '同步进程';
};
</script>

<style lang="scss" scoped>
.QueuePanel {
    height: 80vh;
    width: 40vw;

    .header {
        margin-bottom: 20px;
        padding: 10px;

        .queue-select {
            width: 200px;
        }
    }
}
</style>
