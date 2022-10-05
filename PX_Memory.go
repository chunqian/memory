package memory

import "unsafe"

func PX_MemoryInitialize(mp *PX_memorypool, memory *PX_memory) {
	memory.buffer = (*uint8)(nil)
	memory.allocsize = int32(0)
	memory.usedsize = int32(0)
	memory.mp = mp
}

func PX_MemoryCat(memory *PX_memory, buffer unsafe.Pointer, size PX_int) PX_bool {
	var old *uint8
	var length int32
	var shl int32
	if size == int32(0) {
		return int32(1)
	}
	if memory.usedsize+size > memory.allocsize {
		shl = int32(0)
		old = memory.buffer
		length = memory.usedsize + size
		for int32(int32(1)<<func() (_cgo_ret int32) {
			_cgo_addr := &shl
			*_cgo_addr++
			return *_cgo_addr
		}()) <= length {
		}
		memory.allocsize = int32(1) << shl
		memory.buffer = (*uint8)(MP_Malloc(memory.mp, uint32(memory.allocsize)))
		if !(memory.buffer != nil) {
			return int32(0)
		}
		if old != nil {
			PX_memcpy(unsafe.Pointer(memory.buffer), unsafe.Pointer(old), memory.usedsize)
		}
		PX_memcpy(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(memory.buffer))+uintptr(memory.usedsize)))), buffer, size)
		if old != nil {
			MP_Free(memory.mp, unsafe.Pointer(old))
		}
	} else {
		PX_memcpy(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(memory.buffer))+uintptr(memory.usedsize)))), buffer, size)
	}
	memory.usedsize += size
	return int32(1)
}

func PX_MemoryFree(memory *PX_memory) {
	if memory.allocsize == int32(0) || uintptr(unsafe.Pointer(memory.buffer)) == uintptr(unsafe.Pointer(nil)) {
		return
	}
	MP_Free(memory.mp, unsafe.Pointer(memory.buffer))
	memory.buffer = (*uint8)(nil)
	memory.usedsize = int32(0)
	memory.allocsize = int32(0)
}

func PX_MemoryData(memory *PX_memory) *PX_byte {
	return memory.buffer
}

func PX_MemoryAlloc(memory *PX_memory, size PX_int) PX_bool {
	PX_MemoryFree(memory)
	memory.allocsize = size
	memory.usedsize = int32(0)
	if size == int32(0) {
		memory.buffer = (*uint8)(nil)
		return int32(1)
	} else {
		return func() int32 {
			if uintptr(unsafe.Pointer(func() (_cgo_ret *uint8) {
				_cgo_addr := &memory.buffer
				*_cgo_addr = (*uint8)(MP_Malloc(memory.mp, uint32(size)))
				return *_cgo_addr
			}())) != uintptr(unsafe.Pointer(nil)) {
				return 1
			} else {
				return 0
			}
		}()
	}
	return 0
}

func PX_MemoryResize(memory *PX_memory, size PX_int) PX_bool {
	return PX_MemoryAlloc(memory, size)
}

func PX_MemoryFind(memory *PX_memory, buffer unsafe.Pointer, size PX_int) *PX_byte {
	var offest int32
	if memory.usedsize < size {
		return (*uint8)(nil)
	}
	for offest = int32(0); offest < memory.usedsize-size+int32(1); offest++ {
		if PX_memequ(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(memory.buffer))+uintptr(offest)))), buffer, size) != 0 {
			return (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(memory.buffer)) + uintptr(offest)))
		}
	}
	return (*uint8)(nil)
}

func PX_MemoryRemove(memory *PX_memory, start PX_int, end PX_int) {
	if start > end {
		var t int32 = end
		end = start
		start = t
	}
	if start < int32(0) {
		PX_ASSERT()
		return
	}
	if end >= memory.usedsize {
		PX_ASSERT()
		return
	}
	PX_memcpy(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(memory.buffer))+uintptr(start)))), unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(memory.buffer))+uintptr(end)))))+uintptr(int32(1))))), memory.usedsize-end-int32(1))
	memory.usedsize += start - end - int32(1)
}

func PX_MemoryClear(memory *PX_memory) {
	memory.usedsize = int32(0)
}

func PX_MemoryCopy(memory *PX_memory, buffer unsafe.Pointer, startoffset PX_int, size PX_int) PX_bool {
	var old *uint8
	var length int32
	var shl int32
	if startoffset+size > memory.allocsize {
		shl = int32(0)
		old = memory.buffer
		length = startoffset + size
		for int32(int32(1)<<func() (_cgo_ret int32) {
			_cgo_addr := &shl
			*_cgo_addr++
			return *_cgo_addr
		}()) <= length {
		}
		memory.allocsize = int32(1) << shl
		memory.buffer = (*uint8)(MP_Malloc(memory.mp, uint32(memory.allocsize)))
		if !(memory.buffer != nil) {
			MP_Free(memory.mp, unsafe.Pointer(old))
			return int32(0)
		}
		if old != nil {
			PX_memcpy(unsafe.Pointer(memory.buffer), unsafe.Pointer(old), memory.usedsize)
		}
		PX_memcpy(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(memory.buffer))+uintptr(startoffset)))), buffer, size)
		if old != nil {
			MP_Free(memory.mp, unsafe.Pointer(old))
		}
		memory.usedsize = startoffset + size
	} else {
		if startoffset+size > memory.usedsize {
			memory.usedsize = startoffset + size
		}
		PX_memcpy(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(memory.buffer))+uintptr(startoffset)))), buffer, size)
	}
	return int32(1)
}

func PX_CircularBufferInitialize(mp *PX_memorypool, pcbuffer *PX_circularBuffer, size PX_int) PX_bool {
	PX_memset(unsafe.Pointer(pcbuffer), uint8(0), int32(24))
	if size == int32(0) {
		PX_ASSERT()
	}
	pcbuffer.buffer = (*float64)(MP_Malloc(mp, uint32(8*uint64(size))))
	if !(pcbuffer != nil) {
		return int32(0)
	}
	PX_memset(unsafe.Pointer(pcbuffer.buffer), uint8(0), int32(8*uint64(size)))
	pcbuffer.mp = mp
	pcbuffer.pointer = int32(0)
	pcbuffer.size = size
	return int32(1)
}

func PX_CircularBufferPush(pcbuffer *PX_circularBuffer, v PX_double) {
	pcbuffer.pointer--
	if pcbuffer.pointer < int32(0) {
		pcbuffer.pointer = pcbuffer.size - int32(1)
	}
	*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(pcbuffer.buffer)) + uintptr(pcbuffer.pointer)*8)) = v
}

func PX_CircularBufferAdd(pcbuffer *PX_circularBuffer, pos PX_int, v PX_double) {
	pos = pos % pcbuffer.size
	if pos < int32(0) {
		pos += pcbuffer.size
	}
	*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(pcbuffer.buffer)) + uintptr(pos)*8)) += v
}

func PX_CircularBufferSet(pcbuffer *PX_circularBuffer, pos PX_int, v PX_double) {
	pos = pos % pcbuffer.size
	if pos < int32(0) {
		pos += pcbuffer.size
	}
	*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(pcbuffer.buffer)) + uintptr(pos)*8)) = v
}

func PX_CircularBufferGet(pcbuffer *PX_circularBuffer, pos PX_int) PX_double {
	pos = pos % pcbuffer.size
	if pos < int32(0) {
		pos += pcbuffer.size
	}
	return *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(pcbuffer.buffer)) + uintptr(pos)*8))
}

func PX_CircularBufferZeroClear(pcbuffer *PX_circularBuffer) {
	PX_memset(unsafe.Pointer(pcbuffer.buffer), uint8(0), int32(8*uint64(pcbuffer.size)))
}

func PX_CircularBufferFree(pcbuffer *PX_circularBuffer) {
	if pcbuffer.mp != nil {
		MP_Free(pcbuffer.mp, unsafe.Pointer(pcbuffer.buffer))
		PX_memset(unsafe.Pointer(pcbuffer), uint8(0), int32(24))
	}
}

func PX_CircularBufferDelay(pcbuffer *PX_circularBuffer, pos PX_int) PX_double {
	return *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(pcbuffer.buffer)) + uintptr((pcbuffer.pointer+pos)%pcbuffer.size)*8))
}

func PX_FifoBufferInitialize(mp *PX_memorypool, pfifo *PX_memory) {
	PX_MemoryInitialize(mp, pfifo)
}

func PX_FifoBufferPop(pfifo *PX_fifobuffer, data unsafe.Pointer, size PX_int) PX_int {
	if pfifo.usedsize != 0 {
		var rsize int32 = *(*int32)(unsafe.Pointer(pfifo.buffer))
		if rsize > pfifo.usedsize-int32(4) {
			PX_ASSERT()
			return int32(0)
		}
		if size < rsize {
			return int32(0)
		}
		PX_memcpy(data, unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pfifo.buffer))+uintptr(4)))), rsize)
		PX_MemoryRemove(pfifo, int32(0), int32(uint64(rsize)+4-uint64(1)))
		return rsize
	}
	return int32(0)
}

func PX_FifoBufferPush(pfifo *PX_fifobuffer, data unsafe.Pointer, size PX_int) PX_bool {
	var wsize int32 = size
	if wsize < int32(0) {
		PX_ASSERT()
	}
	if wsize == int32(0) {
		return int32(1)
	}
	if PX_MemoryCat(pfifo, unsafe.Pointer(&wsize), int32(4)) != 0 {
		if PX_MemoryCat(pfifo, data, size) != 0 {
			return int32(1)
		}
	}
	return int32(0)
}

func PX_FifoBufferGetPopSize(pfifo *PX_fifobuffer) PX_int {
	if uint64(pfifo.usedsize) > 4 {
		return *(*int32)(unsafe.Pointer(pfifo.buffer))
	}
	return int32(0)
}

func PX_FifoBufferFree(pfifo *PX_fifobuffer) {
	PX_MemoryFree(pfifo)
}

func PX_StackInitialize(mp *PX_memorypool, pstack *PX_memory) {
	PX_MemoryInitialize(mp, pstack)
}

func PX_StackPop(pstack *PX_stack, data unsafe.Pointer, size PX_int) PX_int {
	if uint64(pstack.usedsize) > 4 {
		var rsize int32 = *(*int32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pstack.buffer))+uintptr(pstack.usedsize))))) - uintptr(4)))))
		if rsize > pstack.usedsize-int32(4) {
			PX_ASSERT()
			return int32(0)
		}
		if size < rsize {
			return int32(0)
		}
		PX_memcpy(data, unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pstack.buffer))+uintptr(pstack.usedsize)))))-uintptr(4)))))-uintptr(rsize)))), rsize)
		PX_MemoryRemove(pstack, int32(uint64(pstack.usedsize)-4-uint64(rsize)), pstack.usedsize-int32(1))
		return rsize
	}
	return int32(0)
}

func PX_StackPush(pstack *PX_stack, data unsafe.Pointer, size PX_int) PX_bool {
	var wsize int32 = size
	if wsize < int32(0) {
		PX_ASSERT()
	}
	if wsize == int32(0) {
		return int32(1)
	}
	if PX_MemoryCat(pstack, data, size) != 0 {
		if PX_MemoryCat(pstack, unsafe.Pointer(&wsize), int32(4)) != 0 {
			return int32(1)
		}
	}
	return int32(0)
}

func PX_StackGetPopSize(pstack *PX_stack) PX_int {
	return *(*int32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pstack.buffer))+uintptr(pstack.usedsize))))) - uintptr(4)))))
}

func PX_StackFree(pstack *PX_stack) {
	PX_MemoryFree(pstack)
}
