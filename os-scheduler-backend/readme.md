## 关于maxProcesses

在代码中 maxProcesses 参数设置为5，它起到以下作用：

1. 限制系统中同时运行的进程数量：
   
   - 系统最多只能同时容纳5个进程（包括就绪队列和运行队列中的进程）
   - 这是一个多道程序设计中的重要参数，也称为"道数"
2. 后备队列管理：
   
   - 当系统中的进程数达到最大值（5个）时，新提交的进程会被放入后备队列
   - 只有当系统中的进程数小于5时，才会从后备队列中调入新的进程
3. 资源管理：
   
   - 通过限制同时存在的进程数，可以更好地管理系统资源
   - 防止系统因进程过多而导致性能下降
4. 调度效率：
   
   - 合理的道数设置可以提高CPU利用率
   - 在这个系统中，有2个处理机和5个最大进程数，意味着：
     - 即使两个处理机都在运行进程
     - 仍然可以有3个进程在就绪队列中等待调度
     - 这样可以保证处理机始终有进程可调度，提高系统吞吐量