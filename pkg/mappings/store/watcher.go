package store

import (
	"context"
	"sync"

	"github.com/loft-sh/vcluster/pkg/syncer/synccontext"
	"k8s.io/client-go/util/workqueue"
)

type watcher struct {
	m sync.Mutex

	addQueueFn synccontext.AddQueueFunc
	queue      workqueue.RateLimitingInterface
}

func (w *watcher) Dispatch(nameMapping synccontext.NameMapping) {
	w.m.Lock()
	defer w.m.Unlock()

	if w.queue == nil {
		return
	}

	w.addQueueFn(nameMapping, w.queue)
}

func (w *watcher) Start(_ context.Context, queue workqueue.RateLimitingInterface) error {
	w.m.Lock()
	defer w.m.Unlock()

	w.queue = queue
	return nil
}

func dispatchAll(watches []*watcher, nameMapping synccontext.NameMapping) {
	if len(watches) == 0 {
		return
	}

	go func(watches []*watcher) {
		for _, watch := range watches {
			watch.Dispatch(nameMapping)
		}
	}(watches)
}
