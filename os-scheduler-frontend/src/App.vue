<template>
  <div class="container">
    <controlPanel />
    <processorPanel />
    <memoryPanel />
    <queuePanel />
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import controlPanel from './components/control-panel.vue';
import processorPanel from './components/processor-panel.vue';
import memoryPanel from './components/memory-panel.vue';
import queuePanel from './components/queue-panel.vue';
import useSystemStatusStore from '@/store/modules/SystemStatus';

const systemStatusStore = useSystemStatusStore();

onMounted(async () => {
  // 组件挂载后执行的逻辑
  await systemStatusStore.getSystemStatus();  // 获取系统状态
  await systemStatusStore.getProcessor();     // 获取处理器状态
});
</script>

<style lang="scss" scoped>
.container {
  padding: 10px;
  display: flex;
  flex-direction: column;
  gap: 5px;

  .control-panel {
    height: 15vh;
  }

  .processor-panel {
    height: 15vh;
  }

  .memory-panel {
    height: 15vh;
  }

  .queue-panel {
    height: 50vh;
    flex: 1; // 剩余空间都给队列面板
    min-height: 0; // 防止内容溢出
  }
}
</style>
