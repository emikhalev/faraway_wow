package closer

import (
	"context"
	"io"
	"os"
	"os/signal"
	"sync"
	"time"
)

type Closer interface {
	AddCloser(closer io.Closer, name string)
	Wait()
}

type closerFn struct {
	name string
	fn   func() error
}

type closer struct {
	ctx context.Context

	muClosers sync.RWMutex
	closers   []closerFn

	logger Logger

	deadline time.Duration

	done chan struct{}
}

func New(ctx context.Context, signals ...os.Signal) *closer {
	closer := new(closer)
	closer.ctx = ctx
	closer.logger = newLogger()
	closer.done = make(chan struct{})

	// Process terminate signals (graceful shutdown)
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, signals...)
		<-ch
		signal.Stop(ch)
		closer.closeAllWithDeadline()
	}()

	return closer
}

func (c *closer) WithLogger(logger Logger) *closer {
	c.logger = logger
	return c
}

func (c *closer) WithDeadline(duration time.Duration) *closer {
	c.deadline = duration
	return c
}

func (c *closer) AddCloser(closer io.Closer, name string) {
	c.muClosers.Lock()
	defer c.muClosers.Unlock()

	c.closers = append(c.closers, closerFn{
		name: name,
		fn:   closer.Close,
	})
}

func (c *closer) Wait() {
	<-c.done
}

func (c *closer) closeAll() {
	c.muClosers.Lock()
	defer c.muClosers.Unlock()

	for idx := len(c.closers) - 1; idx >= 0; idx-- {
		cf := c.closers[idx]
		c.logger.Infof(c.ctx, "%s stopping...", cf.name)
		if err := cf.fn(); err != nil {
			c.logger.Error(c.ctx, err.Error())
		} else {
			c.logger.Infof(c.ctx, "%s gracefully stopped", cf.name)
		}
	}
}

func (c *closer) closeAllWithDeadline() {
	// without deadline
	if c.deadline == 0 {
		c.closeAll()
		c.logger.Infof(c.ctx, "Server gracefully stopped")
		close(c.done)
		return
	}

	// with deadline
	done := make(chan struct{})
	ctx, cancel := context.WithTimeout(c.ctx, c.deadline)
	defer cancel()

	go func() {
		c.closeAll()
		close(done)
	}()

	select {
	case <-done:
		c.logger.Infof(c.ctx, "Server gracefully stopped")
	case <-ctx.Done():
		c.logger.Infof(c.ctx, "Cannot stop server gracefully: timeout")
	}

	close(c.done)
}
