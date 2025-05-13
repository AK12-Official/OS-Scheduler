<template>
    <el-card class="ControlPanel" shadow="never">
        <div class="CardHeader">
            <span> 控制面板</span>
            <span>当前时间片：{{ systemStatusStore.time }}</span>

        </div>
        <div class="button-container">
            <el-button type="primary" size="large">新建进程</el-button>
            <el-button type="primary" size="large">新建随机进程</el-button>
            <el-button type="primary" size="large">单步执行</el-button>
            <el-button type="primary" size="large">自动执行</el-button>
        </div>
    </el-card>
</template>

<script lang="ts" setup>
import useSystemStatusStore from '@/store/modules/SystemStatus';
import { onMounted } from 'vue';

const systemStatusStore = useSystemStatusStore();

onMounted(() => {
    // 组件挂载后执行的逻辑
    systemStatusStore.getSystemStatus();
});
</script>

<style lang="scss" scoped>
.CardHeader {
    font-size: 15px;
    opacity: 0.5;
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
</style>