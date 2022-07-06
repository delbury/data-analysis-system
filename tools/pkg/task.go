package pkg

import (
	"fmt"
	"time"
)

// task 参数
type TaskTime struct {
	// 当天需要执行的时刻
	taskHour   int
	taskMinute int
	taskSecond int
}
type TaskParam struct {
	// 当前时间
	curTime time.Time
	// 定时触发的任务
	job func()
	// 是否当天已完成
	done *bool
	// 当天需要执行的时刻
	taskTime TaskTime
}

// 获取当前的 hour 和 minute
func GetCurrentHourMinute() (int, int) {
	curTime := time.Now()
	return curTime.Hour(), curTime.Minute()
}

// 每天执行一次指定时间段的任务
func DailyTask(job func(), taskHour int, taskMinute int, taskSecond int) {
	// 执行间隔
	ticker := time.NewTicker(time.Minute * 30)
	defer ticker.Stop()

	// 今天任务已完成
	todayIsDone := false

	for curTime := range ticker.C {
		taskTime := TaskTime{taskHour: taskHour, taskMinute: taskMinute, taskSecond: taskSecond}
		taskParam := TaskParam{curTime: curTime, job: job, done: &todayIsDone, taskTime: taskTime}
		Task(taskParam)
	}
}

// 在每天的某个时间段内执行一次
func Task(taskParam TaskParam) {
	curTime, done, job := taskParam.curTime, taskParam.done, taskParam.job
	taskHour, taskMinute, taskSecond := taskParam.taskTime.taskHour, taskParam.taskTime.taskMinute, taskParam.taskTime.taskSecond

	// 构造任务执行的时间
	hour, minute, second, nanosecond := taskHour, taskMinute, taskSecond, 0
	taskTime := time.Date(curTime.Year(), curTime.Month(), curTime.Day(), hour, minute, second, nanosecond, curTime.Location())

	// 当前时间与任务时间相差多少小时
	distanceHours := curTime.Sub(taskTime).Hours()

	if distanceHours >= 0 && distanceHours < 1 {
		// 在任务时间区间内
		if !*done {
			fmt.Println("do job")
			*done = true
			job()
			fmt.Println("job done")
		}
	} else {
		*done = false
	}
}
