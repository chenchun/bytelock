package keylock

import (
	"sync/atomic"
	"time"
)

var (
	KIB31 = 8191     //takes 31KIB memory
	KIB511 = 131071  //takes 511KIB memory
    MIB2 = 524287    //takes 2MIB memory

	sleepTime = 10 * time.Millisecond
)

func NewKeylock() *Keylock {
	return &Keylock{locks: make([]uint32, MIB2), keyGen: Crc32Mod}
}

type Keylock struct {
	locks  []uint32
	keyGen KeyGen
}

func (l *Keylock) Lock(key []byte) {
	for {
		if atomic.CompareAndSwapUint32(&l.locks[l.keyGen(key, len(l.locks))], 0, 1) {
			return
		}
		time.Sleep(sleepTime)
	}
}

func (l *Keylock) Unlock(key []byte) {
	if atomic.CompareAndSwapUint32(&l.locks[l.keyGen(key, len(l.locks))], 1, 0) {
		return
	}
	panic("unlock of unlocked bytelock")
}
