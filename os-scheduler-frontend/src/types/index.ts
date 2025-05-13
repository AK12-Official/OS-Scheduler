export interface ProcessInfo {
    name: string,
    requiredTime: number,
    priority: number,
    memorySize: number
}

interface Response {
    code: number,
    msg: string,
}

interface Process {
    name: string,
    pid: number,
    requiredTime: number,
    totalTime: number,
    priority: number,
    state: string,
    memorySize: number,
    memoryStart: number,
    processorId: number,
    predecessors: null | number[],
    successors: null | number[]
}

export interface SystemStatusResponse extends Response {
    data: {
        queue: {
            ready: Process[],
            running: Process[],
            waiting: Process[],
            backup: Process[],
            suspended: Process[]
        },
        memory: {
            totalSize: number,
            osSize: number,
            blocks: [
                {
                    start: number,
                    length: number,
                    isUsed: boolean
                }
            ]
        }
    }
}

export interface createProcessResponse extends Response {
    data: Process
}