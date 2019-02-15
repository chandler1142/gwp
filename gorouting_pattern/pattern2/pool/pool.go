package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

//Pool 管理一组可以安全地在多个goroutiine间
//共享的资源，被管理的资源必须
//实现 io.Closer接口

type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

//ErrPoolClosed 表示请求了一个已经关闭的池
var ErrPoolClosed = errors.New("Pool has been closed.")

//New 创建一个用来管理资源的池
//这个池需要一个可以分配新资源的函数，
//并规定池的大小
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New(" size value too small.")
	}
	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("Acquire: ", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire: ", "New Resource")
		return p.factory()
	}
}

