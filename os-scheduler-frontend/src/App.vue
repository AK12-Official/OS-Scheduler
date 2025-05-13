<template>
  <controlPanel></controlPanel>
  <processorPanel></processorPanel>
  <queuePanel></queuePanel>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import controlPanel from './components/control-panel.vue';
import processorPanel from './components/processor-panel.vue';
import queuePanel from './components/queue-panel.vue';
import useSystemStatusStore from '@/store/modules/SystemStatus';



const systemStatusStore = useSystemStatusStore();

onMounted(async () => {
  // 组件挂载后执行的逻辑
  await systemStatusStore.getSystemStatus();  // 获取系统状态
  await systemStatusStore.getProcessor();     // 获取处理器状态
});
</script>

<style scoped></style>
