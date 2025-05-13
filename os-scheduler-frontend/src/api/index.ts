import request from "@/utils/request";
import type { createProcessResponse, ProcessInfo, SystemStatusResponse } from "@/types";

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