package monitor

import (
	"bytes"
	"fmt"
	"runtime"
	"runtime/pprof"
	"strconv"
	"strings"
	"sync"
	"time"
)

// MonitorWaitGo 监控阻塞情况
func MonitorWaitGo() {
	waitGoReason := []string{
		"GC assist marking",
		"IO wait",
		"chan receive (nil chan)",
		"chan send (nil chan)",
		"sleep",
		"chan receive",
		"chan send",
		"dumping heap",
		"garbage collection",
		"garbage collection scan",
		"panicwait",
		"select",
		"select (no cases)",
		"GC assist wait",
		"GC sweep wait",
		"GC scavenge wait",
		"finalizer wait",
		"force gc (idle)",
		"semacquire",
		"sync.Cond.Wait",
		"timer goroutine (idle)",
		"trace reader (blocked)",
		"wait for GC cycle",
		"GC worker (idle)",
		"preempted",
		"debug call",
	}
	fmt.Println("开始监控协程状态")
	var pool = sync.Pool{
		New: func() interface{} {

			return new(bytes.Buffer)
		},
	}
	for {
		waitGoReasonMap := map[string]int{}
		for _, i2 := range pprof.Profiles() {
			if i2.Name() == "goroutine" {
				buf := pool.Get().(*bytes.Buffer)
				err := i2.WriteTo(buf, 2)
				if err != nil {
					return
				}
				split := strings.Split(buf.String(), "\ngoroutine")
				for _, s := range split {
					littleSplit := strings.Split("goroutine"+s, "\n")
					for _, s2 := range waitGoReason {
						if strings.Contains(littleSplit[0], s2) {
							if s2 == "chan send" && strings.Contains(littleSplit[0], "chan send (nil chan)") {
								waitGoReasonMap["chan send (nil chan)"]++
								break
							}

							if s2 == "chan receive" && strings.Contains(littleSplit[0], "chan receive (nil chan)") {
								waitGoReasonMap["chan receive (nil chan)"]++
								break
							}

							waitGoReasonMap[s2]++
							if strings.Contains(littleSplit[0], "chan send (nil chan)") || strings.Contains(littleSplit[0], "chan receive (nil chan)") {
							}
							break
						}
					}
				}
				buf.Reset()
				var reasonRobot string
				reasonRobot = "协程状态: 协程总数:" + strconv.Itoa(runtime.NumGoroutine()) + "\n\n\n"
				if len(waitGoReasonMap) <= 0 {
					reasonRobot += "无被阻塞协程" + "\n\n\n"
					continue
				}
				for s, i := range waitGoReasonMap {
					reasonRobot += "阻塞协程:" + strconv.Itoa(i) + ",原因" + s + "\n\n"
				}
				fmt.Println(reasonRobot)
				//repository.RobotTestDemoConf(context.Background(), reasonRobot)
			}
		}
		time.Sleep(time.Second * 1)
	}
}
