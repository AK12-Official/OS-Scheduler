package services

import (
    "os-scheduler-backend/models"
    "errors"
)

type MemoryManager struct {
    Memory *models.MemoryManager
}

func NewMemoryManager(totalSize, osSize int) *MemoryManager {
    return &MemoryManager{
        Memory: &models.MemoryManager{
            TotalSize: totalSize,
            OSSize:    osSize,
            Blocks: []models.MemoryBlock{
                {
                    Start:  osSize,
                    Length: totalSize - osSize,
                    IsUsed: false,
                },
            },
        },
    }
}

func (mm *MemoryManager) Allocate(size int) (int, error) {
    for i, block := range mm.Memory.Blocks {
        if !block.IsUsed && block.Length >= size {
            // 首次适应算法
            mm.Memory.Blocks[i].IsUsed = true
            if block.Length > size {
                // 分割块
                newBlock := models.MemoryBlock{
                    Start:  block.Start + size,
                    Length: block.Length - size,
                    IsUsed: false,
                }
                mm.Memory.Blocks[i].Length = size
                mm.Memory.Blocks = append(mm.Memory.Blocks[:i+1], append([]models.MemoryBlock{newBlock}, mm.Memory.Blocks[i+1:]...)...)
            }
            return block.Start, nil
        }
    }
    return -1, errors.New("no suitable memory block found")
}

func (mm *MemoryManager) Free(start int) {
    for i, block := range mm.Memory.Blocks {
        if block.Start == start {
            mm.Memory.Blocks[i].IsUsed = false
            // 合并相邻空闲块
            mm.mergeBlocks()
            break
        }
    }
}

func (mm *MemoryManager) mergeBlocks() {
    for i := 0; i < len(mm.Memory.Blocks)-1; i++ {
        if !mm.Memory.Blocks[i].IsUsed && !mm.Memory.Blocks[i+1].IsUsed {
            mm.Memory.Blocks[i].Length += mm.Memory.Blocks[i+1].Length
            mm.Memory.Blocks = append(mm.Memory.Blocks[:i+1], mm.Memory.Blocks[i+2:]...)
            i--
        }
    }
}