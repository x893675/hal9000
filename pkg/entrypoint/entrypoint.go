package entrypoint

import (
	"context"
	"errors"
	"os"
	"sync"
)

var (
	mu                    sync.Mutex
	reloadCh              = make(chan struct{})
	shutdownCtx           context.Context
	cancelFn              context.CancelFunc
	errAlreadyInitialized = errors.New("entrypoint already initialized")
	ep                    *EntryPoint
)

func init() {
	shutdownCtx, cancelFn = context.WithCancel(context.Background())
}

func Initialize() (*EntryPoint, error) {
	mu.Lock()
	defer mu.Unlock()
	if ep != nil {
		return nil, errAlreadyInitialized
	}
	ep = &EntryPoint{}
	return ep, nil
}

// OnShutdown subscribe on shutdown event for gracefully exit via context.
func OnShutdown() context.Context {
	return shutdownCtx
}

// OnReload subscribe on reload event.
func OnReload() <-chan struct{} {
	return reloadCh
}

// EntryPoint manager of single point of application
type EntryPoint struct {
}

// Shutdown raise shutdown event.
func (e *EntryPoint) Shutdown(ctx context.Context, code int) {
	mu.Lock()
	defer mu.Unlock()
	cancelFn()
	if _, ok := ctx.Deadline(); ok {
		<-ctx.Done()
	}
	os.Exit(code)
}

// Reload raise reload event.
func (e *EntryPoint) Reload() {
	mu.Lock()
	defer mu.Unlock()
	ch := reloadCh
	reloadCh = make(chan struct{})
	close(ch)
}
