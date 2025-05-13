package services

import (
	"fmt"
	"os-scheduler-backend/models"
	"sort"
	"sync"
)

type Scheduler struct {
	Queue          *models.ProcessQueue
	ProcessorCount int
	MaxProcesses   int
	mutex          sync.Mutex
	nextPID        int
}

func NewScheduler(processorCount, maxProcesses int) *Scheduler {
	return &Scheduler{
		Queue: &models.ProcessQueue{
			Ready:     make([]*models.PCB, 0),
			Running:   make([]*models.PCB, 0),
			Waiting:   make([]*models.PCB, 0),
			Backup:    make([]*models.PCB, 0),
			Suspended: make([]*models.PCB, 0),
		},
		ProcessorCount: processorCount,
		MaxProcesses:   maxProcesses,
		nextPID:        1,
	}
}

func (s *Scheduler) AddProcess(process *models.PCB) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	process.PID = s.nextPID
	s.nextPID++
	process.State = models.Ready
	process.ProcessorID = -1

	if len(s.Queue.Ready)+len(s.Queue.Running) < s.MaxProcesses {
		s.Queue.Ready = append(s.Queue.Ready, process)
		s.sortReadyQueue()
	} else {
		s.Queue.Backup = append(s.Queue.Backup, process)
	}
}

func (s *Scheduler) Schedule() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// 1. 处理运行中的进程
	for _, p := range s.Queue.Running {
		p.Priority--
		p.RequiredTime--
		
		if p.RequiredTime <= 0 {
			// 进程完成，移出运行队列
			s.removeFromRunning(p)
			p.State = models.Finished
		} else {
			// 进程未完成，放回就绪队列以便重新参与调度
			s.removeFromRunning(p)
			p.State = models.Ready
			p.ProcessorID = -1
			s.Queue.Ready = append(s.Queue.Ready, p)
		}
	}

	// 2. 重新排序就绪队列
	s.sortReadyQueue()

	// 3. 为所有处理机重新分配进程
	for i := 0; i < s.ProcessorCount; i++ {
		if len(s.Queue.Ready) > 0 {
			process := s.Queue.Ready[0]
			s.Queue.Ready = s.Queue.Ready[1:]
			process.State = models.Running
			process.ProcessorID = i
			s.Queue.Running = append(s.Queue.Running, process)
		}
	}

	// 4. 从后备队列调入新进程
	for len(s.Queue.Ready) + len(s.Queue.Running) < s.MaxProcesses && len(s.Queue.Backup) > 0 {
		process := s.Queue.Backup[0]
		s.Queue.Backup = s.Queue.Backup[1:]
		s.Queue.Ready = append(s.Queue.Ready, process)
		s.sortReadyQueue()
	}
}

func (s *Scheduler) sortReadyQueue() {
	sort.Slice(s.Queue.Ready, func(i, j int) bool {
		return s.Queue.Ready[i].Priority > s.Queue.Ready[j].Priority
	})
}

func (s *Scheduler) removeFromRunning(process *models.PCB) {
	for i, p := range s.Queue.Running {
		if p.PID == process.PID {
			s.Queue.Running = append(s.Queue.Running[:i], s.Queue.Running[i+1:]...)
			break
		}
	}
}

// SuspendProcess 将指定进程挂起
func (s *Scheduler) SuspendProcess(pid int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// 在就绪队列中查找进程
	for i, p := range s.Queue.Ready {
		if p.PID == pid {
			p.State = models.Suspended
			s.Queue.Suspended = append(s.Queue.Suspended, p)
			s.Queue.Ready = append(s.Queue.Ready[:i], s.Queue.Ready[i+1:]...)
			return nil
		}
	}

	// 在运行队列中查找进程
	for i, p := range s.Queue.Running {
		if p.PID == pid {
			p.State = models.Suspended
			s.Queue.Suspended = append(s.Queue.Suspended, p)
			s.Queue.Running = append(s.Queue.Running[:i], s.Queue.Running[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("找不到进程 %d", pid)
}

// ResumeProcess 恢复被挂起的进程
func (s *Scheduler) ResumeProcess(pid int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// 在挂起队列中查找进程
	for i, p := range s.Queue.Suspended {
		if p.PID == pid {
			p.State = models.Ready
			s.Queue.Ready = append(s.Queue.Ready, p)
			s.Queue.Suspended = append(s.Queue.Suspended[:i], s.Queue.Suspended[i+1:]...)
			s.sortReadyQueue() // 重新排序就绪队列
			return nil
		}
	}

	return fmt.Errorf("找不到被挂起的进程 %d", pid)
}
