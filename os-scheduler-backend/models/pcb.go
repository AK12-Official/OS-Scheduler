package models

type ProcessState string

const (
	Ready     ProcessState = "ready"
	Running   ProcessState = "running"
	Waiting   ProcessState = "waiting"
	Finished  ProcessState = "finished"
	Suspended ProcessState = "suspended"
)

type PCB struct {
	Name              string       `json:"name"`
	PID               int          `json:"pid"`
	RequiredTime      int         `json:"requiredTime"` // 剩余运行时间
	TotalRequiredTime int         `json:"totalTime"`    // 总运行时间
	Priority          int         `json:"priority"`
	State             ProcessState `json:"state"`
	MemorySize        int          `json:"memorySize"`
	MemoryStart       int          `json:"memoryStart"`
	ProcessorID       int          `json:"processorId"`  // -1表示未分配处理机
	Predecessors      []int        `json:"predecessors"` // 前驱进程PID列表
	Successors        []int        `json:"successors"`   // 后继进程PID列表
}
