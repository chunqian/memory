package memory

import unsafe "unsafe"

type PX_memoryNode struct {
	StartAddr unsafe.Pointer
	EndAddr   unsafe.Pointer
}

func PX_MemoryPool_GetFreeTableAddr(MP *PX_memorypool) *PX_memoryNode {
	return (*PX_memoryNode)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(MP.EndAddr)))-uintptr(16*uint64(MP.FreeTableCount)))))) + uintptr(int32(1))))))
}

func PX_MemoryPool_GetFreeTable(MP *PX_memorypool, Index PX_uint) *PX_memoryNode {
	Index++
	return (*PX_memoryNode)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(MP.EndAddr)))-uintptr(16*uint64(Index)))))) + uintptr(int32(1))))))
}

func PX_AllocFromFree(MP *PX_memorypool, Size PX_uint) *PX_memoryNode {
	var Node *PX_memoryNode
	if uint64(Size)+32 > uint64(MP.FreeSize) {
		return (*PX_memoryNode)(nil)
	}
	Node = (*PX_memoryNode)(MP.AllocAddr)
	(*Node).StartAddr = unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(MP.AllocAddr))) + uintptr(16))))
	(*Node).EndAddr = unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(Node.StartAddr)))+uintptr(Size))))) - uintptr(int32(1)))))
	MP.FreeSize -= uint32(uint64(Size) + 32)
	MP.AllocAddr = unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(MP.AllocAddr)))+uintptr(Size))))) + uintptr(16))))
	return Node
}

func PX_MemoryPoolRemoveFreeNode(MP *PX_memorypool, Index PX_uint) {
	var i uint32
	var pFreeNode *PX_memoryNode = PX_MemoryPool_GetFreeTable(MP, Index)
	for i = Index; i < MP.FreeTableCount; i++ {
		*pFreeNode = *(*PX_memoryNode)(unsafe.Pointer(uintptr(unsafe.Pointer(pFreeNode)) - uintptr(int32(1))*16))
		*(*uintptr)(unsafe.Pointer(&pFreeNode)) -= 16
	}
	MP.FreeTableCount--
	if MP.FreeTableCount == uint32(0) {
		MP.MaxMemoryfragSize = uint32(0)
	}
}

func PX_AllocFreeMemoryNode(MP *PX_memorypool) *PX_memoryNode {
	MP.FreeTableCount++
	return PX_MemoryPool_GetFreeTable(MP, MP.FreeTableCount-uint32(1))
}

func PX_UpdateMaxFreqSize(MP *PX_memorypool) {
	var itNode *PX_memoryNode
	var i uint32
	var Size uint32
	MP.MaxMemoryfragSize = uint32(0)
	for i = uint32(0); i < MP.FreeTableCount; i++ {
		itNode = PX_MemoryPool_GetFreeTable(MP, i)
		if func() (_cgo_ret uint32) {
			_cgo_addr := &Size
			*_cgo_addr = uint32(uintptr(unsafe.Pointer((*int8)(itNode.EndAddr))) - uintptr(unsafe.Pointer((*int8)(itNode.StartAddr))) + uintptr(int64(1)))
			return *_cgo_addr
		}() > MP.MaxMemoryfragSize {
			MP.MaxMemoryfragSize = Size
		}
	}
}

func PX_AllocFromFreq(MP *PX_memorypool, Size PX_uint) *PX_memoryNode {
	var i uint32
	var fSize uint32
	var itNode *PX_memoryNode
	var allocNode *PX_memoryNode
	Size += uint32(16)
	if MP.MaxMemoryfragSize >= Size {
		for i = uint32(0); i < MP.FreeTableCount; i++ {
			itNode = PX_MemoryPool_GetFreeTable(MP, i)
			fSize = uint32(uintptr(unsafe.Pointer((*int8)(itNode.EndAddr))) - uintptr(unsafe.Pointer((*int8)(itNode.StartAddr))) + uintptr(int64(1)))
			if Size <= fSize && uint64(Size)+16 >= uint64(fSize) {
				allocNode = (*PX_memoryNode)(itNode.StartAddr)
				allocNode.StartAddr = unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(itNode.StartAddr))) + uintptr(16))))
				allocNode.EndAddr = itNode.EndAddr
				PX_MemoryPoolRemoveFreeNode(MP, i)
				PX_UpdateMaxFreqSize(MP)
				return allocNode
			} else if Size < fSize {
				if uint64(MP.FreeSize) < 16 {
					return (*PX_memoryNode)(nil)
				}
				allocNode = (*PX_memoryNode)(itNode.StartAddr)
				allocNode.StartAddr = unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(itNode.StartAddr))) + uintptr(16))))
				allocNode.EndAddr = unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(itNode.StartAddr)))+uintptr(Size))))) - uintptr(int32(1)))))
				itNode.StartAddr = unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(allocNode.EndAddr))) + uintptr(int32(1)))))
				MP.FreeSize -= uint32(16)
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
	MP.StartAddr = MemoryAddr
	MP.AllocAddr = MemoryAddr
	if MemorySize != 0 {
		MP.EndAddr = unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(MemoryAddr)))+uintptr(MemorySize))))) - uintptr(int32(1)))))
	} else {
		MP.EndAddr = MP.StartAddr
	}
	MP.Size = MemorySize
	MP.FreeSize = MemorySize
	MP.FreeTableCount = uint32(0)
	MP.MaxMemoryfragSize = uint32(0)
	MP.nodeCount = uint32(0)
	MP.ErrorCall_Ptr = (func(int32))(nil)
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
		return MemNode.StartAddr
	}
	MemNode = PX_AllocFromFree(MP, Size)
	if uintptr(unsafe.Pointer(MemNode)) != uintptr(unsafe.Pointer(nil)) {
		MP.nodeCount++
		return MemNode.StartAddr
	}
	if func(_cgo_fn func(int32)) uintptr {
		return *(*uintptr)(unsafe.Pointer(&_cgo_fn))
	}(MP.ErrorCall_Ptr) == uintptr(unsafe.Pointer(nil)) {
		PX_ERROR((*int8)(unsafe.Pointer(&[26]int8{'M', 'e', 'm', 'o', 'r', 'y', 'P', 'o', 'o', 'l', ' ', 'O', 'u', 't', ' ', 'O', 'f', ' ', 'M', 'e', 'm', 'o', 'r', 'y', '!', '\x00'})))
		PX_ASSERT()
	} else {
		MP.ErrorCall_Ptr(int32(0))
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
	FreeNode.StartAddr = TempNode.StartAddr
	FreeNode.EndAddr = TempNode.EndAddr
	FreeNode.StartAddr = unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(FreeNode.StartAddr))) - uintptr(16))))
	if uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(FreeNode.EndAddr)))+uintptr(int32(1)))))) == uintptr(unsafe.Pointer((*int8)(MP.AllocAddr))) {
		for i = uint32(0); i < MP.FreeTableCount; i++ {
			itNode = PX_MemoryPool_GetFreeTable(MP, i)
			if uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(itNode.EndAddr)))+uintptr(int32(1)))))) == uintptr(unsafe.Pointer((*int8)(FreeNode.StartAddr))) {
				MP.AllocAddr = itNode.StartAddr
				MP.FreeSize += uint32(uint64(uintptr(unsafe.Pointer((*int8)(FreeNode.EndAddr)))-uintptr(unsafe.Pointer((*int8)(FreeNode.StartAddr)))) + 16 + uint64(1))
				MP.FreeSize += uint32(uint64(uintptr(unsafe.Pointer((*int8)(itNode.EndAddr)))-uintptr(unsafe.Pointer((*int8)(itNode.StartAddr)))+uintptr(int64(1))) + 16)
				PX_MemoryPoolRemoveFreeNode(MP, i)
				goto _END
			}
		}
		MP.AllocAddr = unsafe.Pointer((*int8)(FreeNode.StartAddr))
		MP.FreeSize += uint32(uint64(uintptr(unsafe.Pointer((*int8)(FreeNode.EndAddr)))-uintptr(unsafe.Pointer((*int8)(FreeNode.StartAddr)))) + 16 + uint64(1))
		goto _END
	}
	sIndex = uint32(4294967295)
	for i = uint32(0); i < MP.FreeTableCount; i++ {
		itNode = PX_MemoryPool_GetFreeTable(MP, i)
		pcTempStart = (*uint8)(itNode.StartAddr)
		pcTempEnd = (*uint8)(itNode.EndAddr)
		if uintptr(unsafe.Pointer(pcTempStart)) == uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(FreeNode.EndAddr)))+uintptr(int32(1)))))) {
			if sIndex == uint32(4294967295) {
				sIndex = i
				itNode.StartAddr = FreeNode.StartAddr
				FreeNode = *itNode
				MP.FreeSize += uint32(16)
			} else {
				MP.FreeSize += uint32(16)
				itNode.StartAddr = FreeNode.StartAddr
				PX_MemoryPoolRemoveFreeNode(MP, sIndex)
			}
			bExist = uint8(1)
		}
		if uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pcTempEnd))+uintptr(int32(1)))))) == uintptr(unsafe.Pointer((*uint8)(FreeNode.StartAddr))) {
			if sIndex == uint32(4294967295) {
				sIndex = i
				itNode.EndAddr = FreeNode.EndAddr
				FreeNode = *itNode
				MP.FreeSize += uint32(16)
			} else {
				itNode.EndAddr = FreeNode.EndAddr
				MP.FreeSize += uint32(16)
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
	Pool.ErrorCall_Ptr = ErrorCall
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
	return uint32(uintptr(unsafe.Pointer((*int8)(TempNode.EndAddr)))-uintptr(unsafe.Pointer((*int8)(TempNode.StartAddr)))) + uint32(1)
}

func MP_Reset(Pool *PX_memorypool) {
	Pool.AllocAddr = Pool.StartAddr
	Pool.EndAddr = unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(Pool.StartAddr)))+uintptr(Pool.Size))))) - uintptr(int32(1)))))
	Pool.FreeSize = Pool.Size
	Pool.FreeTableCount = uint32(0)
	Pool.MaxMemoryfragSize = uint32(0)
	Pool.nodeCount = uint32(0)
}
