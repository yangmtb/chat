package store

import (
	"container/list"
	"sync"
	"time"
)

// Store an object implementing store interface can be registered with setcustomstore
// function to handle storage and retrieval of captcha ids and solutions for them,
// replacing the default memory store.
//
// it is the responsibility of an object to delete expired and used captchas when
// necessary (for example, the default memory store collects them is set method after
// the certain amount of captchas has been stored.)
type Store interface {
	// Set sets the captcha for the captcha id
	Set(id string, value string)

	// Get returns stored captcha for the id. clear indicates
	// whether the captcha must be deleted from the store.
	Get(id string, clear bool) string
}

// expValue stores timestamp and id of captchas. it is used in the list inside
// memoryStore for indexing generated captchas by timestamp by timestamp to enable garbage
// collection of expired captchas.
type idTimeValue struct {
	timestamp time.Time
	id        string
}

// memoryStore is an internal store for captcha ids and their values
type memoryStore struct {
	sync.RWMutex
	// id : captchas
	captchas map[string]string
	// ids
	ids *list.List
	// number of items stored since last collection
	storedNum int
	// number of saved items that triggers collection
	collectNum int
	// expiration time of captchas
	expiration time.Duration
}

// NewMemoryStore return a new standard memory store for captchas with the given
// collection threshold and expiration time (duration). the returned store must
// be registered with set customstore to replace the default one.
func NewMemoryStore(collectNum int, expiration time.Duration) Store {
	s := new(memoryStore)
	s.captchas = make(map[string]string)
	s.ids = list.New()
	s.storedNum = 0
	s.collectNum = collectNum
	s.expiration = expiration
	return s
}

func (s *memoryStore) Set(id, value string) {
	s.Lock()
	s.captchas[id] = value
	s.ids.PushBack(idTimeValue{time.Now(), id})
	s.storedNum++
	s.Unlock()
	if s.storedNum > s.collectNum {
		go s.collect()
	}
}

func (s *memoryStore) Get(id string, clear bool) (value string) {
	if !clear {
		// when we don't need to clear captcha, acquire read lock.
		s.RLock()
		defer s.RUnlock()
	} else {
		s.Lock()
		defer s.Unlock()
	}
	value, ok := s.captchas[id]
	if !ok {
		return
	}
	if clear {
		delete(s.captchas, id)
	}
	return
}

func (s *memoryStore) collect() {
	now := time.Now()
	s.Lock()
	defer s.Unlock()
	s.storedNum = 0
	for e := s.ids.Front(); nil != e; {
		ev, ok := e.Value.(idTimeValue)
		if !ok {
			return
		}
		if ev.timestamp.Add(s.expiration).Before(now) {
			delete(s.captchas, ev.id)
			next := e.Next()
			s.ids.Remove(e)
			e = next
		} else {
			return
		}
	}
}
