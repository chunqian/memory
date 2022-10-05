package memory

import "unsafe"

type PX_memoryNode struct {
	startAddr unsafe.Pointer
	endAddr   unsafe.Pointer
}

func PX_MemoryPool_GetFreeTableAddr(MP *PX_memorypool) *PX_memoryNode {
	return (*PX_memoryNode)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(MP.endAddr)))-uintptr(16*uint64(MP.freeTableCount)))))) + uintptr(int32(1))))))
}

func PX_MemoryPool_GetFreeTable(MP *PX_memorypool, Index PX_uint) *PX_memoryNode {
	Index++
	return (*PX_memoryNode)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(MP.endAddr)))-uintptr(16*uint64(Index)))))) + uintptr(int32(1))))))
}

func PX_AllocFromFree(MP *PX_memorypool, Size PX_uint) *PX_memoryNode {
	var Node *PX_memoryNode
	if uint64(Size)+32 > uint64(MP.freeSize) {
		return (*PX_memoryNode)(nil)
	}
	Node = (*PX_memoryNode)(MP.allocAddr)
	(*Node).startAddr = unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(MP.allocAddr))) + uintptr(16))))
	(*Node).endAddr = unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(Node.startAddr)))+uintptr(Size))))) - uintptr(int32(1)))))
	MP.freeSize -= uint32(uint64(Size) + 32)
	MP.allocAddr = unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(MP.allocAddr)))+uintptr(Size))))) + uintptr(16))))
	return Node
}

func PX_MemoryPoolRemoveFreeNode(MP *PX_memorypool, Index PX_uint) {
	var i uint32
	var pFreeNode *PX_memoryNode = PX_MemoryPool_GetFreeTable(MP, Index)
	for i = Index; i < MP.freeTableCount; i++ {
		*pFreeNode = *(*PX_memoryNode)(unsafe.Pointer(uintptr(unsafe.Pointer(pFreeNode)) - uintptr(int32(1))*16))
		*(*uintptr)(unsafe.Pointer(&pFreeNode)) -= 16
	}
	MP.freeTableCount--
	if MP.freeTableCount == uint32(0) {
		MP.maxMemoryfragSize = uint32(0)
	}
}

func PX_AllocFreeMemoryNode(MP *PX_memorypool) *PX_memoryNode {
	MP.freeTableCount++
	return PX_MemoryPool_GetFreeTable(MP, MP.freeTableCount-uint32(1))
}

func PX_UpdateMaxFreqSize(MP *PX_memorypool) {
	var itNode *PX_memoryNode
	var i uint32
	var Size uint32
	MP.maxMemoryfragSize = uint32(0)
	for i = uint32(0); i < MP.freeTableCount; i++ {
		itNode = PX_MemoryPool_GetFreeTable(MP, i)
		if func() (_cgo_ret uint32) {
			_cgo_addr := &Size
			*_cgo_addr = uint32(uintptr(unsafe.Pointer((*int8)(itNode.endAddr))) - uintptr(unsafe.Pointer((*int8)(itNode.startAddr))) + uintptr(int64(1)))
			return *_cgo_addr
		}() > MP.maxMemoryfragSize {
			MP.maxMemoryfragSize = Size
		}
	}
}

func PX_AllocFromFreq(MP *PX_memorypool, Size PX_uint) *PX_memoryNode {
	var i uint32
	var fSize uint32
	var itNode *PX_memoryNode
	var allocNode *PX_memoryNode
	Size += uint32(16)
	if MP.maxMemoryfragSize >= Size {
		for i = uint32(0); i < MP.freeTableCount; i++ {
			itNode = PX_MemoryPool_GetFreeTable(MP, i)
			fSize = uint32(uintptr(unsafe.Pointer((*int8)(itNode.endAddr))) - uintptr(unsafe.Pointer((*int8)(itNode.startAddr))) + uintptr(int64(1)))
			if Size <= fSize && uint64(Size)+16 >= uint64(fSize) {
				allocNode = (*PX_memoryNode)(itNode.startAddr)
				allocNode.startAddr = unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(itNode.startAddr))) + uintptr(16))))
				allocNode.endAddr = itNode.endAddr
				PX_MemoryPoolRemoveFreeNode(MP, i)
				PX_UpdateMaxFreqSize(MP)
				return allocNode
			} else if Size < fSize {
				if uint64(MP.freeSize) < 16 {
					return (*PX_memoryNode)(nil)
				}
				allocNode = (*PX_memoryNode)(itNode.startAddr)
				allocNode.startAddr = unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(itNode.startAddr))) + uintptr(16))))
				allocNode.endAddr = unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(itNode.startAddr)))+uintptr(Size))))) - uintptr(int32(1)))))
				itNode.startAddr = unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(allocNode.endAddr))) + uintptr(int32(1)))))
				MP.freeSize -= uint32(16)
				PX_UpdateMaxFreqSize(MP)
				return allocNode
			}
		}
		return (*PX_memoryNode)(nil)
	} else {
		return (*PX_memoryNode)(nil)
	}
	return nil
}

func MP_Create(MemoryAddr unsafe.Pointer, MemorySize PX_uint) PX_memorypool {
	var Index uint32 = uint32(0)
	var MP PX_memorypool
	PX_memset(unsafe.Pointer(&MP), uint8(0), int32(56))
	if MemorySize == uint32(0) {
		return MP
	}
	MP.startAddr = MemoryAddr
	MP.allocAddr = MemoryAddr
	if MemorySize != 0 {
		MP.endAddr = unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(MemoryAddr)))+uintptr(MemorySize))))) - uintptr(int32(1)))))
	} else {
		MP.endAddr = MP.startAddr
	}
	MP.size = MemorySize
	MP.freeSize = MemorySize
	MP.freeTableCount = uint32(0)
	MP.maxMemoryfragSize = uint32(0)
	MP.nodeCount = uint32(0)
	MP.errorCallPtr = (func(int32))(nil)
	PX_memset(MemoryAddr, uint8(0), int32(MemorySize))
	func() int {
		_ = Index
		return 0
	}()
	return MP
}

func MP_Malloc(MP *PX_memorypool, Size PX_uint) unsafe.Pointer {
	var MemNode *PX_memoryNode
	if Size == uint32(0) {
		return unsafe.Pointer(nil)
	}
	if uint64(Size)%8 != 0 {
		Size = uint32((uint64(Size)/8 + uint64(1)) * 8)
	}
	MemNode = PX_AllocFromFreq(MP, Size)
	if uintptr(unsafe.Pointer(MemNode)) != uintptr(unsafe.Pointer(nil)) {
		MP.nodeCount++
		return MemNode.startAddr
	}
	MemNode = PX_AllocFromFree(MP, Size)
	if uintptr(unsafe.Pointer(MemNode)) != uintptr(unsafe.Pointer(nil)) {
		MP.nodeCount++
		return MemNode.startAddr
	}
	if func(_cgo_fn func(int32)) uintptr {
		return *(*uintptr)(unsafe.Pointer(&_cgo_fn))
	}(MP.errorCallPtr) == uintptr(unsafe.Pointer(nil)) {
		PX_ERROR((*int8)(unsafe.Pointer(&[26]int8{'M', 'e', 'm', 'o', 'r', 'y', 'P', 'o', 'o', 'l', ' ', 'O', 'u', 't', ' ', 'O', 'f', ' ', 'M', 'e', 'm', 'o', 'r', 'y', '!', '\x00'})))
		PX_ASSERT()
	} else {
		MP.errorCallPtr(int32(0))
		PX_ASSERT()
	}
	return unsafe.Pointer(nil)
}

func MP_Free(MP *PX_memorypool, pAddress unsafe.Pointer) {
	var i uint32
	var sIndex uint32
	var itNode *PX_memoryNode
	var FreeNode PX_memoryNode
	var pcTempStart *uint8
	var pcTempEnd *uint8
	var bExist uint8
	var TempPointer unsafe.Pointer
	var TempNode *PX_memoryNode
	MP.nodeCount--
	bExist = uint8(0)
	TempPointer = unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(pAddress))) - uintptr(16))))
	TempNode = (*PX_memoryNode)(TempPointer)
	FreeNode.startAddr = TempNode.startAddr
	FreeNode.endAddr = TempNode.endAddr
	FreeNode.startAddr = unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(FreeNode.startAddr))) - uintptr(16))))
	if uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(FreeNode.endAddr)))+uintptr(int32(1)))))) == uintptr(unsafe.Pointer((*int8)(MP.allocAddr))) {
		for i = uint32(0); i < MP.freeTableCount; i++ {
			itNode = PX_MemoryPool_GetFreeTable(MP, i)
			if uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(itNode.endAddr)))+uintptr(int32(1)))))) == uintptr(unsafe.Pointer((*int8)(FreeNode.startAddr))) {
				MP.allocAddr = itNode.startAddr
				MP.freeSize += uint32(uint64(uintptr(unsafe.Pointer((*int8)(FreeNode.endAddr)))-uintptr(unsafe.Pointer((*int8)(FreeNode.startAddr)))) + 16 + uint64(1))
				MP.freeSize += uint32(uint64(uintptr(unsafe.Pointer((*int8)(itNode.endAddr)))-uintptr(unsafe.Pointer((*int8)(itNode.startAddr)))+uintptr(int64(1))) + 16)
				PX_MemoryPoolRemoveFreeNode(MP, i)
				goto _END
			}
		}
		MP.allocAddr = unsafe.Pointer((*int8)(FreeNode.startAddr))
		MP.freeSize += uint32(uint64(uintptr(unsafe.Pointer((*int8)(FreeNode.endAddr)))-uintptr(unsafe.Pointer((*int8)(FreeNode.startAddr)))) + 16 + uint64(1))
		goto _END
	}
	sIndex = uint32(4294967295)
	for i = uint32(0); i < MP.freeTableCount; i++ {
		itNode = PX_MemoryPool_GetFreeTable(MP, i)
		pcTempStart = (*uint8)(itNode.startAddr)
		pcTempEnd = (*uint8)(itNode.endAddr)
		if uintptr(unsafe.Pointer(pcTempStart)) == uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(FreeNode.endAddr)))+uintptr(int32(1)))))) {
			if sIndex == uint32(4294967295) {
				sIndex = i
				itNode.startAddr = FreeNode.startAddr
				FreeNode = *itNode
				MP.freeSize += uint32(16)
			} else {
				MP.freeSize += uint32(16)
				itNode.startAddr = FreeNode.startAddr
				PX_MemoryPoolRemoveFreeNode(MP, sIndex)
			}
			bExist = uint8(1)
		}
		if uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pcTempEnd))+uintptr(int32(1)))))) == uintptr(unsafe.Pointer((*uint8)(FreeNode.startAddr))) {
			if sIndex == uint32(4294967295) {
				sIndex = i
				itNode.endAddr = FreeNode.endAddr
				FreeNode = *itNode
				MP.freeSize += uint32(16)
			} else {
				itNode.endAddr = FreeNode.endAddr
				MP.freeSize += uint32(16)
				PX_MemoryPoolRemoveFreeNode(MP, sIndex)
			}
			bExist = uint8(1)
		}
	}
	if int32(bExist) == int32(0) {
		*PX_AllocFreeMemoryNode(MP) = FreeNode
	}
_END:
	PX_UpdateMaxFreqSize(MP)
	return
}

func MP_Release(Node *PX_memorypool) {
}

func MP_ErrorCatch(Pool *PX_memorypool, ErrorCall func(PX_MEMORYPOOL_ERROR)) {
	Pool.errorCallPtr = ErrorCall
}

func MP_Size(Pool *PX_memorypool, pAddress unsafe.Pointer) PX_uint {
	var TempPointer unsafe.Pointer
	var TempNode *PX_memoryNode
	if uintptr(unsafe.Pointer(pAddress)) == uintptr(unsafe.Pointer(nil)) {
		PX_ASSERT()
		return uint32(0)
	}
	TempPointer = unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(pAddress))) - uintptr(16))))
	TempNode = (*PX_memoryNode)(TempPointer)
	return uint32(uintptr(unsafe.Pointer((*int8)(TempNode.endAddr)))-uintptr(unsafe.Pointer((*int8)(TempNode.startAddr)))) + uint32(1)
}

func MP_Reset(Pool *PX_memorypool) {
	Pool.allocAddr = Pool.startAddr
	Pool.endAddr = unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(Pool.startAddr)))+uintptr(Pool.size))))) - uintptr(int32(1)))))
	Pool.freeSize = Pool.size
	Pool.freeTableCount = uint32(0)
	Pool.maxMemoryfragSize = uint32(0)
	Pool.nodeCount = uint32(0)
}
