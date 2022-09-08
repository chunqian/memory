package memory

import unsafe "unsafe"

var _px_hex_to_dex_table [103]int32 = [103]int32{int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(1), int32(2), int32(3), int32(4), int32(5), int32(6), int32(7), int32(8), int32(9), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(10), int32(11), int32(12), int32(13), int32(14), int32(15), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(10), int32(11), int32(12), int32(13), int32(14), int32(15)}

type PX_bool = int32
type PX_dword = uint32
type PX_short = int16
type PX_int16 = int16
type PX_uint16 = uint16
type PX_word = uint16
type PX_ushort = uint16
type PX_uint = uint32
type PX_uint32 = uint32
type PX_int = int32
type PX_int32 = int32
type PX_char = int8
type PX_byte = uint8
type PX_uchar = uint8
type PX_ulong = uint64
type PX_long = int64
type PX_float = float32
type PX_float32 = float32
type PX_double = float64
type PX_double64 = float64
type PX_qword = uint64
type PX_uint64 = uint64
type PX_int64 = int64

type _px_return_string struct {
	data [64]int8
}
type PX_RETURN_STRING = _px_return_string

type PX_STRINGFORMAT_TYPE = int32

type _inner_px_stringformat struct {
	_pstring *int8
}
type _px_stringformat struct {
	type_ int32
	_inner_px_stringformat
}
type PX_stringformat = _px_stringformat

type PX_MEMORYPOOL_ERROR = int32
const (
	PX_MEMORYPOOL_ERROR_OUTOFMEMORY     PX_MEMORYPOOL_ERROR = 0
	PX_MEMORYPOOL_ERROR_INVALID_ACCESS  PX_MEMORYPOOL_ERROR = 1
	PX_MEMORYPOOL_ERROR_INVALID_ADDRESS PX_MEMORYPOOL_ERROR = 2
)

type PX_MP_ErrorCall = func(PX_MEMORYPOOL_ERROR)

type _memoryPool struct {
	AllocAddr         unsafe.Pointer
	StartAddr         unsafe.Pointer
	EndAddr           unsafe.Pointer
	Size              uint32
	FreeSize          uint32
	nodeCount         uint32
	FreeTableCount    uint32
	MaxMemoryfragSize uint32
	ErrorCall_Ptr     func(int32)
}
type PX_memorypool = _memoryPool

type _memroy struct {
	buffer    *uint8
	mp        *_memoryPool
	usedsize  int32
	allocsize int32
}
type PX_memory = _memroy
type PX_fifobuffer = _memroy
type PX_stack = _memroy

type _px_circularBuffer struct {
	mp      *_memoryPool
	buffer  *float64
	size    int32
	pointer int32
}
type PX_CircularBuffer = _px_circularBuffer
