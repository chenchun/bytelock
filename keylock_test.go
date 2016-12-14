package keylock

import "testing"

//race test
func TestLockUnLock(t *testing.T) {
	l := NewKeylock()
	key := []byte("pod1-22-31")
	a1 := 0
	a2 := 0
	go func() {
		l.Lock(key)
		a1 = 1
		a2 = 1
		l.Unlock(key)
	}()
	l.Lock(key)
	a1 = 2
	a2 = 2
	l.Unlock(key)
	if a1 != a2 {
		t.Fatal()
	}
}

func TestUnLockUnlock(t *testing.T) {
	l := NewKeylock()
	key := []byte("pod1-22-31")
	assertPanic(t, func() {l.Unlock(key)})
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}
