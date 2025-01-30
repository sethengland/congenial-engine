package storage

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"
)

var (
	inMem = make(map[uuid.UUID]int)
	mu    = &sync.RWMutex{}

	errContextCanceled = errors.New("context canceled")
	errNotFound        = errors.New("no records for the given id")
)

func SaveRecord(ctx context.Context, points int) (uuid.UUID, error) {
	select {
	case <-ctx.Done():
		return uuid.Nil, errContextCanceled
	default:
		mu.Lock()
		defer mu.Unlock()
		id := uuid.New()
		inMem[id] = points
		return id, nil
	}

}
func GetRecord(ctx context.Context, id uuid.UUID) (int, error) {
	select {
	case <-ctx.Done():
		return 0, errContextCanceled
	default:
		mu.RLock()
		defer mu.RUnlock()
		if points, ok := inMem[id]; ok {
			return points, nil
		}
		return 0, errNotFound
	}
}
