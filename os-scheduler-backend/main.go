package main

import (
	"fmt"
	"net/http"
	"os-scheduler-backend/models"
	"os-scheduler-backend/services"

	_ "os-scheduler-backend/docs" // 这里会引入自动生成的 docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Response represents a generic response structure
type Response struct {
	Code    int         `json:"code"`    // 状态码，0表示成功，非0表示错误
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 数据，可以为空
}

// StatusResponse represents the system status response
type StatusResponse struct {
	Queue  *models.ProcessQueue  `json:"queue"`
	Memory *models.MemoryManager `json:"memory"`
}

// ProcessorStatusResponse 表示处理机状态响应
type ProcessorStatusResponse struct {
    Processors []*models.PCB `json:"processors"` // 每个处理机当前运行的进程，如果没有则为nil
}

// @title 操作系统调度器 API
// @version 1.0
// @description 这是一个操作系统调度实验的后端API服务
// @host localhost:8080
// @BasePath /
var (
	scheduler     *services.Scheduler
	memoryManager *services.MemoryManager
)

func main() {
	// 初始化调度器和内存管理器
	scheduler = services.NewScheduler(2, 5)              // 2个处理机，最大5个进程
	memoryManager = services.NewMemoryManager(4096, 128) // 总内存4096，操作系统占128

	r := gin.Default()

	// 允许跨域
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// API路由
	r.POST("/process", addProcess)
	r.GET("/status", getStatus)
	r.POST("/schedule", runSchedule)
	r.POST("/suspend/:pid", suspendProcess)
	r.POST("/resume/:pid", resumeProcess)
	r.GET("/processor-status", getProcessorStatus)  // 添加新路由

	// 添加 swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

// @Summary 添加新进程
// @Description 添加一个新的进程到系统中
// @Accept json
// @Produce json
// @Param process body models.PCB true "进程信息"
// @Success 200 {object} models.PCB
// @Failure 400 {object} Response
// @Router /process [post]
func addProcess(c *gin.Context) {
	var process models.PCB
	if err := c.BindJSON(&process); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "请求参数错误",
			Data:    err.Error(),
		})
		return
	}

	start, err := memoryManager.Allocate(process.MemorySize)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "内存分配失败",
			Data:    err.Error(),
		})
		return
	}
	process.MemoryStart = start
	process.TotalRequiredTime = process.RequiredTime

	scheduler.AddProcess(&process)
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "进程添加成功",
		Data:    process,
	})
}

func getStatus(c *gin.Context) {
	status := StatusResponse{
		Queue:  scheduler.Queue,
		Memory: memoryManager.Memory,
	}
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "获取状态成功",
		Data:    status,
	})
}

func runSchedule(c *gin.Context) {
	scheduler.Schedule()
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "调度执行成功",
		Data:    scheduler.Queue,
	})
}

func suspendProcess(c *gin.Context) {
	pid := c.Param("pid")
	if pid == "" {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "进程ID不能为空",
		})
		return
	}

	var processID int
	if _, err := fmt.Sscanf(pid, "%d", &processID); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "无效的进程ID",
			Data:    err.Error(),
		})
		return
	}

	if err := scheduler.SuspendProcess(processID); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "进程挂起失败",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: fmt.Sprintf("进程 %d 已挂起", processID),
	})
}

func resumeProcess(c *gin.Context) {
	pid := c.Param("pid")
	if pid == "" {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "进程ID不能为空",
		})
		return
	}

	var processID int
	if _, err := fmt.Sscanf(pid, "%d", &processID); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "无效的进程ID",
			Data:    err.Error(),
		})
		return
	}

	if err := scheduler.ResumeProcess(processID); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "进程恢复失败",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: fmt.Sprintf("进程 %d 已恢复", processID),
	})
}

func getProcessorStatus(c *gin.Context) {
    processorCount := scheduler.ProcessorCount
    processors := make([]*models.PCB, processorCount)

    for _, p := range scheduler.Queue.Running {
        if p.ProcessorID >= 0 && p.ProcessorID < processorCount {
            processors[p.ProcessorID] = p
        }
    }

    c.JSON(http.StatusOK, Response{
        Code:    0,
        Message: "获取处理机状态成功",
        Data: ProcessorStatusResponse{
            Processors: processors,
        },
    })
}
