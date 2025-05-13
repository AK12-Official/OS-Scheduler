package models

type MemoryBlock struct {
    Start    int  `json:"start"`
    Length   int  `json:"length"`
    IsUsed   bool `json:"isUsed"`
}

type MemoryManager struct {
    TotalSize    int           `json:"totalSize"`
    OSSize       int           `json:"osSize"`
    Blocks       []MemoryBlock `json:"blocks"`
}