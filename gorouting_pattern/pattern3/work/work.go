package work

import "sync"

//定义一个接口类型
type Worker interface {
	Task()
}

//Pool提供一个goroutine池，这个池可以完成
//任何已提交的worker任务
type Pool struct {
	work chan Worker
	wg sync.WaitGroup
}

//创建一个新的链接池
func New(maxGoroutine int) *Pool {
	p := Pool{
		work:make(chan Worker),
	}

	p.wg.Add(maxGoroutine)
	for i:=0;i<maxGoroutine;i++ {
		go func() {
			for w:= range p.work {
				w.Task()
			}
			//等待每一个goroutine做完
			p.wg.Done()
		}()
	}

	return &p
}

//提交工作到工作池
func (p *Pool) Run(w Worker) {
	p.work <- w
}

//Shutdown 等待所有的goroutine停止工作
func (p *Pool) Shutdown()  {
	close(p.work)
	p.wg.Wait()
}
