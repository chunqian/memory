package memory

import unsafe "unsafe"

var _cgos_px_hex_to_dex_table_PX_Typedef [103]int32 = [103]int32{int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(1), int32(2), int32(3), int32(4), int32(5), int32(6), int32(7), int32(8), int32(9), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(10), int32(11), int32(12), int32(13), int32(14), int32(15), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(10), int32(11), int32(12), int32(13), int32(14), int32(15)}

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
type _cgoa_1_PX_Typedef struct {
	data [64]int8
}
type PX_RETURN_STRING = _cgoa_1_PX_Typedef

const (
	PX_ALIGN_LEFTTOP     int32 = 0
	PX_ALIGN_MIDTOP      int32 = 1
	PX_ALIGN_RIGHTTOP    int32 = 2
	PX_ALIGN_LEFTMID     int32 = 3
	PX_ALIGN_CENTER      int32 = 4
	PX_ALIGN_RIGHTMID    int32 = 5
	PX_ALIGN_LEFTBOTTOM  int32 = 6
	PX_ALIGN_MIDBOTTOM   int32 = 7
	PX_ALIGN_RIGHTBOTTOM int32 = 8
)

type PX_ALIGN = int32
type _cgoa_3_PX_Typedef struct {
	_11 float32
	_12 float32
	_13 float32
	_14 float32
	_21 float32
	_22 float32
	_23 float32
	_24 float32
	_31 float32
	_32 float32
	_33 float32
	_34 float32
	_41 float32
	_42 float32
	_43 float32
	_44 float32
}
type _cgoa_2_PX_Typedef struct {
	 _cgoa_3_PX_Typedef
}
type Struct__px_matrix struct {
	_cgoa_2_PX_Typedef
}
type PX_matrix = Struct__px_matrix
type _cgoa_5_PX_Typedef struct {
	r uint8
	g uint8
	b uint8
	a uint8
}
type _cgoa_4_PX_Typedef struct {
	 _cgoa_5_PX_Typedef
}
type Struct__px_color struct {
	_argb _cgoa_4_PX_Typedef
}
type PX_color = Struct__px_color
type Struct__px_color_hsl struct {
	a float32
	H float32
	S float32
	L float32
}
type PX_color_hsl = Struct__px_color_hsl
type Struct__px_point struct {
	x float32
	y float32
	z float32
}
type PX_point = Struct__px_point
type PX_point32 = Struct__px_point
type Struct__px_wpoint struct {
	x float64
	y float64
	z float64
}
type PX_wpoint = Struct__px_wpoint
type _cgoa_6_PX_Typedef struct {
	x float32
	y float32
}
type PX_point2D = _cgoa_6_PX_Typedef
type PX_vector2D = _cgoa_6_PX_Typedef
type PX_point3D = Struct__px_point
type PX_vector3D = Struct__px_point
type Struct__px_point4 struct {
	x float32
	y float32
	z float32
	w float32
}
type PX_point4D = Struct__px_point4
type PX_vertex = Struct__px_point4
type PX_vector4D = Struct__px_point4
type _cgoa_7_PX_Typedef struct {
	p0 Struct__px_point
	n  Struct__px_point
}
type PX_plane = _cgoa_7_PX_Typedef
type Struct__px_rect struct {
	x      float32
	y      float32
	width  float32
	height float32
}
type PX_rect = Struct__px_rect
type Struct___px_complex struct {
	re float64
	im float64
}
type PX_complex = Struct___px_complex
type Struct___px_timestamp struct {
	year   int16
	month  int16
	day    int16
	hour   int16
	minute int16
	second int16
}
type PX_timestamp = Struct___px_timestamp

const (
	PX_STRINGFORMAT_TYPE_INT    int32 = 0
	PX_STRINGFORMAT_TYPE_FLOAT  int32 = 1
	PX_STRINGFORMAT_TYPE_STRING int32 = 2
)

type PX_STRINGFORMAT_TYPE = int32
type _cgoa_9_PX_Typedef struct {
	_pstring *int8
}
type _cgoa_8_PX_Typedef struct {
	type_ int32
	_cgoa_9_PX_Typedef
}
type PX_stringformat = _cgoa_8_PX_Typedef
type _cgoa_10_PX_Typedef struct {
	Vertex_1 Struct__px_point
	Vertex_2 Struct__px_point
	Vertex_3 Struct__px_point
}
type PX_triangle = _cgoa_10_PX_Typedef

const (
	PX_CEPTRUM_TYPE_REAL     int32 = 0
	PX_CEPSTRUM_TYPE_COMPLEX int32 = 1
)

type PX_CEPSTRUM_TYPE = int32
type _cgoa_12_PX_Typedef struct {
	ipv6 [4]uint32
}
type _cgoa_11_PX_Typedef struct {
	port uint32
	_cgoa_12_PX_Typedef
}
type PX_Network_Addr = _cgoa_11_PX_Typedef
type _cgoa_13_PX_Typedef struct {
	A float64
	p float64
	f float64
}
type PX_sine = _cgoa_13_PX_Typedef

const (
	PX_FIRFILTER_TYPE_LOWPASS        int32 = 0
	PX_FIRFILTER_TYPE_HIGHPASS       int32 = 1
	PX_FIRFILTER_TYPE_BANDPASS       int32 = 2
	PX_FIRFILTER_TYPE_STOPBANDFILTER int32 = 3
)

type PX_FIRFILTER_TYPE = int32

const (
	PX_FIRFILTER_WINDOW_TYPE_RECT       int32 = 0
	PX_FIRFILTER_WINDOW_TYPE_TUKEY      int32 = 1
	PX_FIRFILTER_WINDOW_TYPE_TRIANGULAR int32 = 2
	PX_FIRFILTER_WINDOW_TYPE_HANNING    int32 = 3
	PX_FIRFILTER_WINDOW_TYPE_HAMMING    int32 = 4
	PX_FIRFILTER_WINDOW_TYPE_BLACKMAN   int32 = 5
	PX_FIRFILTER_WINDOW_TYPE_KAISER     int32 = 6
)

type PX_FIRFILTER_WINDOW_TYPE = int32

var _cgos_px_hex_to_dex_table_PX_Log [103]int32 = [103]int32{int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(1), int32(2), int32(3), int32(4), int32(5), int32(6), int32(7), int32(8), int32(9), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(10), int32(11), int32(12), int32(13), int32(14), int32(15), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(10), int32(11), int32(12), int32(13), int32(14), int32(15)}
var _cgos_px_hex_to_dex_table_PX_MemoryPool [103]int32 = [103]int32{int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(1), int32(2), int32(3), int32(4), int32(5), int32(6), int32(7), int32(8), int32(9), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(10), int32(11), int32(12), int32(13), int32(14), int32(15), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(10), int32(11), int32(12), int32(13), int32(14), int32(15)}

const (
	PX_MEMORYPOOL_ERROR_OUTOFMEMORY     int32 = 0
	PX_MEMORYPOOL_ERROR_INVALID_ACCESS  int32 = 1
	PX_MEMORYPOOL_ERROR_INVALID_ADDRESS int32 = 2
)

type PX_MEMORYPOOL_ERROR = int32
type PX_MP_ErrorCall = func(int32)
type Struct__memoryPool struct {
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
type PX_memorypool = Struct__memoryPool

var _cgos_px_hex_to_dex_table_PX_Memory [103]int32 = [103]int32{int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(1), int32(2), int32(3), int32(4), int32(5), int32(6), int32(7), int32(8), int32(9), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(10), int32(11), int32(12), int32(13), int32(14), int32(15), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0), int32(10), int32(11), int32(12), int32(13), int32(14), int32(15)}

type Struct___PX_memroy struct {
	buffer    *uint8
	mp        *Struct__memoryPool
	usedsize  int32
	allocsize int32
}
type PX_memory = Struct___PX_memroy
type _cgoa_10_PX_Memory struct {
	mp      *Struct__memoryPool
	buffer  *float64
	size    int32
	pointer int32
}
type PX_CircularBuffer = _cgoa_10_PX_Memory
type PX_fifobuffer = Struct___PX_memroy
type PX_stack = Struct___PX_memroy
