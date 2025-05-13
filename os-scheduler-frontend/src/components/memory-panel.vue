<template>
    <el-card class="MemoryPanel" shadow="never">
        <div class="CardHeader">
            <span>内存使用情况</span>
            <div class="memory-info">
                <span>总大小: {{ systemStatusStore.SystemStatus.memory.totalSize }}KB</span>
                <span>系统占用: {{ systemStatusStore.SystemStatus.memory.osSize }}KB</span>
            </div>
        </div>
        <div class="memory-blocks">
            <!-- 显示系统占用的内存块 -->
            <div class="memory-block system-block"
                :style="{ width: `${(systemStatusStore.SystemStatus.memory.osSize / systemStatusStore.SystemStatus.memory.totalSize) * 100}%` }">
                <el-tooltip content="系统占用" placement="top">
                    <div class="block-content">
                        系统占用: {{ systemStatusStore.SystemStatus.memory.osSize }}KB
                    </div>
                </el-tooltip>
            </div>
            <!-- 显示其他内存块 -->
            <div v-for="(block, index) in systemStatusStore.SystemStatus.memory.blocks" :key="index"
                class="memory-block" :class="{ used: block.isUsed }" :style="{
                    width: `${calculateBlockWidth(block.length, index)}%`,
                    flexGrow: index === systemStatusStore.SystemStatus.memory.blocks.length - 1 ? 1 : 0
                }">
                <el-tooltip :content="`起始位置: ${block.start}KB\n大小: ${block.length}KB`" placement="top">
                    <div class="block-content">
                        {{ block.length }}KB
                    </div>
                </el-tooltip>
            </div>
        </div>
    </el-card>
</template>

<script lang="ts" setup>
import useSystemStatusStore from '@/store/modules/SystemStatus';

const systemStatusStore = useSystemStatusStore();

// 计算内存块宽度的函数
const calculateBlockWidth = (length: number, index: number) => {
    const totalSize = systemStatusStore.SystemStatus.memory.totalSize;
    // 如果是最后一个块，确保填充剩余空间
    if (index === systemStatusStore.SystemStatus.memory.blocks.length - 1) {
        return (length / totalSize) * 100;
    }
    return (length / totalSize) * 100;
};
</script>

<style lang="scss" scoped>
.MemoryPanel {
    width: 100%;
    margin: 20px 0;

    .CardHeader {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 20px;

        .memory-info {
            display: flex;
            gap: 20px;
        }
    }

    .memory-blocks {
        display: flex;
        width: 100%;
        height: 50px;
        border: 1px solid #dcdfe6;
        border-radius: 4px;
        overflow: hidden;

        .memory-block {
            height: 100%;
            transition: all 0.3s;
            position: relative;

            &.system-block {
                background-color: #67c23a; // 使用绿色表示系统占用
                color: white;
                border-right: 1px solid #dcdfe6;
            }

            &.used {
                background-color: #409eff;
                color: white;
            }

            &:not(.used):not(.system-block) {
                background-color: #f5f7fa;
                color: #606266;
            }

            .block-content {
                position: absolute;
                width: 100%;
                height: 100%;
                display: flex;
                align-items: center;
                justify-content: center;
                font-size: 12px;
                overflow: hidden;
                white-space: nowrap;
            }

            &:not(:last-child) {
                border-right: 1px solid #dcdfe6;
            }

            &:last-child {
                flex-grow: 1; // 让最后一个块自动填充剩余空间
            }
        }
    }
}
</style>