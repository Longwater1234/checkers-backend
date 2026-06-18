---
name: "Fix sync.Pool buffer leak bugs"
about: "Two critical bugs in the buffer pool implementation that cause memory leaks"
title: "Fix: sync.Pool buffer leak - pointer and oversized buffer issues"
labels: ["bug"]
---

## 🐛 Issue: sync.Pool Buffer Leak Bugs

There are **2 subtle bugs** in the proposed sync.Pool buffer reuse implementation (#87) that will cause memory leaks:

### Bug #1: Return by Value, Not Pointer ⚠️
**Problem:**
```go
// ❌ WRONG - unsafe type assertion and returns copy
buf := bufferPool.Get().([]byte)
bufferPool.Put(buf)  // Returns value, not pointer!
```

**Why it's a problem:**
- Type assertion from `interface{}` creates an unsafe copy of the slice header
- Inconsistent with Go best practices for pooled objects
- Can cause panics or undefined behavior

**Fix:**
```go
// ✅ CORRECT - store pointer to []byte in pool
var bufferPool = sync.Pool{
	New: func() interface{} {
		return new([]byte)  // return pointer
	},
}

// Usage:
bufPtr := bufferPool.Get().(*[]byte)
*bufPtr = (*bufPtr)[:0]  // reset length, keep capacity
bufferPool.Put(bufPtr)
```

---

### Bug #2: Oversized Buffers Poison the Pool 💾
**Problem:**
```go
buf := bufferPool.Get().([]byte)  // capacity 4096
websocket.Message.Receive(hunter.Conn, &buf)  // if > 4096 bytes, buf grows!
bufferPool.Put(buf)  // ❌ Returns HUGE buffer to pool - MEMORY LEAK!
```

**Why it's a problem:**
- When a message exceeds the buffer capacity, `websocket.Message.Receive()` reallocates
- The oversized buffer gets returned to the pool
- Future `Get()` calls return huge buffers, defeating the purpose of pooling
- Memory gradually accumulates

**Fix:**
```go
const maxPooledSize = 4096

// After receiving:
if cap(*bufPtr) <= maxPooledSize {
	*bufPtr = (*bufPtr)[:0]
	bufferPool.Put(bufPtr)  // ✅ Only return normal-sized buffers
} else {
	// Let oversized buffer be garbage collected
	bufferPool.Put(&[]byte{})
}
```

---

## Solution Summary
1. Store `*[]byte` in the pool, not `[]byte`
2. Only return buffers of expected size back to the pool
3. Check capacity before returning; oversized buffers should be discarded

See issue #87 for context.
