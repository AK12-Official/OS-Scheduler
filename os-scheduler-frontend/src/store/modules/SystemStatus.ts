import { defineStore } from 'pinia'
import { getStatus, createProcess, singleSchedule, suspendProcess, resumeProcess, getProcessorStatus, resetSystem } from '@/api/index';
import { ref } from 'vue';
import { ElMessage, ElNotification } from 'element-plus';
import type { Process, ProcessInfo } from '@/types';

const useSystemStatusStore = defineStore('SystemStatus', () => {
    const SystemStatus = ref({
        queue: {
            ready: [] as Process[],
            running: [] as Process[],
            waiting: [] as Process[],
            backup: [] as Process[],
            suspended: [] as Process[]
        },
        memory: {
            totalSize: 0,
            osSize: 0,
            blocks: [] as {
                start: number,
                length: number,
                isUsed: boolean
            }[]
        }
    });
    const time = ref(0);
    const ProcessorsStatus = ref<(Process | null)[]>([null, null]);

    const getSystemStatus = async () => {
        const res = await getStatus();
        if (res.code === 0) {
            SystemStatus.value = res.data;
        } else {
            ElNotification({
                type: 'error',
                message: res.message || '获取系统状态失败'
            })
        }
    };

    const createNewProcess = async (data: ProcessInfo) => {
        const res = await createProcess(data);
        if (res.code === 0) {
            ElMessage.success(res.message)
            await getSystemStatus();
        } else {
            ElNotification({
                type: 'error',
                message: res.message || '创建进程失败'
            })
        }


    };

    const Schedule = async () => {
        const res = await singleSchedule();
        if (res.code === 0) {
            SystemStatus.value = {
                ...SystemStatus.value,
                ...res.data
            };
            ElMessage.success(res.message);
            time.value++;
            await getSystemStatus();
        } else {
            ElNotification({
                type: 'error',
                message: res.message || '调度失败'
            })
        }
    }

    const Suspend = async (pid: number) => {
        const res = await suspendProcess(pid);
        if (res.code === 0) {
            ElMessage.success(res.message)
            await getSystemStatus();
        } else {
            ElNotification({
                type: 'error',
                message: res.message || '挂起失败'
            })
        }
    }

    const Resume = async (pid: number) => {
        const res = await resumeProcess(pid);
        if (res.code === 0) {
            ElMessage.success(res.message)
            await getSystemStatus();
        } else {
            ElNotification({
                type: 'error',
                message: res.message || '恢复失败'
            })
        }
    }

    const getProcessor = async () => {
        const res = await getProcessorStatus();

        if (res.code == 0) {
            ProcessorsStatus.value = res.data.processors;
        } else {
            ElNotification({
                type: 'error',
                message: res.message || '获取处理机状态失败'
            })
        }
    }

    const reset = async () => {
        const res = await resetSystem();
        if (res.code === 0) {
            ElMessage.success(res.message)
            await getSystemStatus();
            time.value = 0;
        } else {
            ElNotification({
                type: 'error',
                message: res.message || '重置系统失败'
            })
        }
    }

    return {
        SystemStatus,
        time,
        ProcessorsStatus,
        getSystemStatus,
        createNewProcess,
        Schedule,
        Suspend,
        Resume,
        getProcessor,
        reset
    }
})

export default useSystemStatusStore;

