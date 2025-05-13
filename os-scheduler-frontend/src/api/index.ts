import request from "@/utils/request";
import type { createProcessResponse, ProcessorResponse, ProcessInfo, ScheduleResponse, SuspendAndResumeResponse, SystemStatusResponse } from "@/types";

enum API {
    GET_STATUS = "/status",
    CREATE_PROCESS = "/process",
    SINGELE_SCHEDULE = "/schedule",
    SUSPEND = "/suspend/",
    RESUME = "/resume/",
    GET_PROCESSER_STATUS = "/processor-status",
}

export const getStatus = () => request.get<void, SystemStatusResponse>(API.GET_STATUS);

export const createProcess = (data: ProcessInfo) => request.post<unknown, createProcessResponse>(API.CREATE_PROCESS, data);

export const singleSchedule = () => request.post<void, ScheduleResponse>(API.SINGELE_SCHEDULE);

export const suspendProcess = (pid: number) => request.post<void, SuspendAndResumeResponse>(API.SUSPEND + pid);

export const resumeProcess = (pid: number) => request.post<void, SuspendAndResumeResponse>(API.RESUME + pid);

export const getProcessorStatus = () => request.get<void, ProcessorResponse>(API.GET_PROCESSER_STATUS);