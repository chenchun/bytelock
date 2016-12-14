# keylock

`keylock` allocates a block of memory to create a lock pool.
The convenience of `keylock` is to use a lock without worry about if it the lock is created or deleted. 
It keeps you away from the complexity of when to create or delete a lock.

```
l := NewKeylock()
key := []byte("pod1-22-31")
l.Lock(key)
dosomething()...
l.Unlock(key)
```
