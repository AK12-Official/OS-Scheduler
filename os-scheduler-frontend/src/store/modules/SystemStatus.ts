import { defineStore } from 'pinia'
import { getStatus, createProcess } from '@/api/index';
import { ref } from 'vue';
import { ElNotification } from 'element-plus';
import type { ProcessInfo } from '@/types';

const useSystemStatusStore = defineStore('SystemStatus', () => {
    const SystemStatus = ref({})
    const time: number = 0;
    const getSystemStatus = async () => {
        const res = await getStatus();
        if (res.code === 0) {
            SystemStatus.value = res.data;
        } else {
            ElNotification({
                type: 'error',
                message: res.msg || '获取系统状态失败'
            })
        }
    };

    const createNewProcess = async (data: ProcessInfo) => {
        const res = await createProcess(data);
        if (res.code === 0) {
            ElNotification({
                type: 'success',
                message: res.msg || '创建进程成功'
            })
            await getSystemStatus();
        } else {
            ElNotification({
                type: 'error',
                message: res.msg || '创建进程失败'
            })
        }
    };

    return {
        SystemStatus,
        time,
        getSystemStatus,
        createNewProcess
    }
})

export default useSystemStatusStore;