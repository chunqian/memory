package memory

import unsafe "unsafe"

func PX_isBigEndianCPU() PX_bool {
	type _inner_px_isBigEndianCPU struct {
		i uint32
	}
	var c _inner_px_isBigEndianCPU
	c.i = uint32(305419896)
	return func() int32 {
		if int32(18) == int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&*(*[4]uint8)(unsafe.Pointer(&c)))))) + uintptr(int32(0))))) {
			return 1
		} else {
			return 0
		}
	}()
}

func PX_htoi(hex_str *PX_char) PX_uint {
	var ch int8
	var iret uint32 = uint32(0)
	for int32(func() (_cgo_ret int8) {
		_cgo_addr := &ch
		*_cgo_addr = *func() (_cgo_ret *int8) {
			_cgo_addr := &hex_str
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		return *_cgo_addr
	}()) != int32(0) {
		iret = iret<<int32(4) | uint32(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&px_hex_to_dex_table)))) + uintptr(ch)*4)))
	}
	return iret
}

func PX_atoi(s *PX_char) PX_int {
	var i int32
	var n int32
	var sign int32 = int32(1)
	for i = int32(0); int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i)))) == ' '; i++ {
	}
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i)))) == '+' || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i)))) == '-' {
		sign = func() int32 {
			if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i)))) == '-' {
				return -1
			} else {
				return int32(1)
			}
		}()
		i++
	}
	for n = int32(0); int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i)))) >= '0' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i)))) <= '9'; i++ {
		n = int32(10)*n + (int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i)))) - '0')
	}
	return sign * n
}

func PX_atof(fstr *PX_char) PX_float {
	var temp float64 = float64(int32(10))
	var ispnum int32 = int32(1)
	var ans float64 = float64(int32(0))
	var be int32 = int32(0)
	if int32(*fstr) == '-' {
		ispnum = int32(0)
		*(*uintptr)(unsafe.Pointer(&fstr))++
	} else if int32(*fstr) == '+' {
		*(*uintptr)(unsafe.Pointer(&fstr))++
	}
	for int32(*fstr) != '\x00' {
		if int32(*fstr) == '.' {
			*(*uintptr)(unsafe.Pointer(&fstr))++
			break
		}
		if int32(*fstr) == 'e' || int32(*fstr) == 'E' {
			*(*uintptr)(unsafe.Pointer(&fstr))++
			be = int32(1)
			break
		}
		ans = ans*float64(int32(10)) + float64(int32(*fstr)-'0')
		*(*uintptr)(unsafe.Pointer(&fstr))++
	}
	if be != 0 {
		var e int32 = PX_atoi(fstr)
		var e10 float32 = float32(int32(1))
		if e > int32(0) {
			for e != 0 {
				e10 *= float32(int32(10))
				e--
			}
		} else {
			for e != 0 {
				e10 /= float32(int32(10))
				e++
			}
		}
		ans *= float64(e10)
	} else {
		for int32(*fstr) != '\x00' {
			ans = ans + float64(int32(*fstr)-'0')/temp
			temp *= float64(int32(10))
			*(*uintptr)(unsafe.Pointer(&fstr))++
		}
	}
	if ispnum != 0 {
		return float32(ans)
	} else {
		return float32(ans) * float32(-1)
	}
	return 0
}

func PX_ftos(f PX_float, precision PX_int) PX_RETURN_STRING {
	var str PX_RETURN_STRING
	PX_ftoa(f, (*int8)(unsafe.Pointer(&str.data)), int32(64), precision)
	return str
}

func PX_itos(num PX_int, radix PX_int) PX_RETURN_STRING {
	var str PX_RETURN_STRING
	PX_itoa(num, (*int8)(unsafe.Pointer(&str.data)), int32(64), radix)
	return str
}

func PX_AscToWord(asc *PX_char, u16 *PX_word) {
	for *asc != 0 {
		*u16 = uint16(*asc)
		*(*uintptr)(unsafe.Pointer(&u16)) += 2
		*(*uintptr)(unsafe.Pointer(&asc))++
	}
	*u16 = uint16(0)
}

func PX_ftoa(f PX_float, outbuf *PX_char, maxlen PX_int, precision PX_int) PX_int {
	var i_value int32
	var f_value int32
	var shl int32 = int32(10)
	var len int32
	var zero_oft int32 = int32(0)
	if maxlen == int32(0) {
		return int32(0)
	}
	shl = PX_pow_ii(shl, precision)
	i_value = int32(f)
	f_value = int32(func() float32 {
		if float32(shl)*(f-float32(int32(f))) > float32(int32(0)) {
			return float32(shl) * (f - float32(int32(f)))
		} else {
			return -(float32(shl) * (f - float32(int32(f))))
		}
	}())
	if f_value < int32(0) {
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(outbuf)) + uintptr(int32(0)))) = int8('N')
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(outbuf)) + uintptr(int32(1)))) = int8('a')
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(outbuf)) + uintptr(int32(2)))) = int8('N')
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(outbuf)) + uintptr(int32(3)))) = int8('\x00')
		return int32(0)
	}
	if i_value == int32(0) && f < float32(int32(0)) {
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(outbuf)) + uintptr(int32(0)))) = int8('-')
		PX_itoa(i_value, (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(outbuf))+uintptr(int32(1)))), maxlen-int32(1), int32(10))
	} else {
		PX_itoa(i_value, outbuf, maxlen, int32(10))
	}
	len = PX_strlen(outbuf)
	if precision == int32(0) {
		return len
	}
	if len > maxlen-int32(3) {
		return len
	}
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(outbuf)) + uintptr(len))) = int8('.')
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(outbuf)) + uintptr(len+int32(1)))) = int8('\x00')
	for f_value < shl {
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(outbuf))+uintptr(len)))))+uintptr(int32(1)))))) + uintptr(zero_oft))) = int8('0')
		shl /= int32(10)
		zero_oft++
	}
	PX_itoa(f_value, (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(outbuf))+uintptr(len)))))+uintptr(zero_oft))), maxlen-len-int32(1), int32(10))
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(outbuf))+uintptr(len)))))+uintptr(int32(1)))))) + uintptr(precision))) = int8('\x00')
	len += PX_strlen((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(outbuf)) + uintptr(len))))
	return len
}

func PX_itoa(num PX_int, str *PX_char, MaxStrSize PX_int, radix PX_int) PX_int {
	var index [37]int8 = [37]int8{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '\x00'}
	var unum uint32
	var temp int8
	var i int32 = int32(0)
	var j int32
	var l int32
	if MaxStrSize == int32(0) {
		return int32(0)
	}
	if radix == int32(10) && num < int32(0) {
		unum = uint32(-num)
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(func() (_cgo_ret int32) {
			_cgo_addr := &i
			_cgo_ret = *_cgo_addr
			*_cgo_addr++
			return
		}()))) = int8('-')
	} else {
		unum = uint32(num)
	}
	for {
		if MaxStrSize <= i+int32(1) {
			return int32(0)
		}
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(func() (_cgo_ret int32) {
			_cgo_addr := &i
			_cgo_ret = *_cgo_addr
			*_cgo_addr++
			return
		}()))) = *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&index)))) + uintptr(unum%uint32(radix))))
		unum /= uint32(radix)
		if !(unum != 0) {
			break
		}
	}
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(i))) = int8('\x00')
	l = i
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(int32(0))))) == '-' {
		for func() int32 {
			j = int32(1)
			return func() (_cgo_ret int32) {
				_cgo_addr := &i
				*_cgo_addr = i - int32(1)
				return *_cgo_addr
			}()
		}(); j < i; func() int32 {
			j++
			return func() (_cgo_ret int32) {
				_cgo_addr := &i
				_cgo_ret = *_cgo_addr
				*_cgo_addr--
				return
			}()
		}() {
			temp = *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(j)))
			*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(j))) = *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(i)))
			*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(i))) = temp
		}
	} else {
		for func() int32 {
			j = int32(0)
			return func() (_cgo_ret int32) {
				_cgo_addr := &i
				_cgo_ret = *_cgo_addr
				*_cgo_addr--
				return
			}()
		}(); j < i; func() int32 {
			j++
			return func() (_cgo_ret int32) {
				_cgo_addr := &i
				_cgo_ret = *_cgo_addr
				*_cgo_addr--
				return
			}()
		}() {
			temp = *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(j)))
			*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(j))) = *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(i)))
			*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(i))) = temp
		}
	}
	return l
}

func PX_SwapEndian(val PX_dword) PX_dword {
	val = val<<int32(8)&uint32(4278255360) | val>>int32(8)&uint32(16711935)
	return val<<int32(16) | val>>int32(16)
}

func PX_BufferToHexString(data *PX_byte, size PX_int, hex_str *PX_char) {
	var i int32
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(hex_str)) + uintptr(int32(0)))) = int8(0)
	for i = int32(0); i < size; i++ {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(i)))) == int32(0) {
			PX_strcat((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(hex_str))+uintptr(i*int32(2)))), (*int8)(unsafe.Pointer(&[3]int8{'0', '0', '\x00'})))
		} else if int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(i)))) < int32(16) {
			*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(hex_str)) + uintptr(i*int32(2)))) = int8('0')
			PX_itoa(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(i)))), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(hex_str))+uintptr(i*int32(2))))))+uintptr(int32(1)))), int32(3), int32(16))
		} else {
			PX_itoa(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(i)))), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(hex_str))+uintptr(i*int32(2)))), int32(3), int32(16))
		}
	}
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(hex_str)) + uintptr(i*int32(2)))) = int8(0)
}

func PX_HexStringToBuffer(hex_str *PX_char, data *PX_byte) PX_int {
	var i int32 = int32(0)
	var hex [3]int8 = [3]int8{int8(0)}
	if PX_strlen(hex_str)&int32(1) != 0 {
		return int32(0)
	}
	for *hex_str != 0 {
		if int32(*hex_str) == ' ' {
			*(*uintptr)(unsafe.Pointer(&hex_str))++
			continue
		}
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&hex)))) + uintptr(int32(0)))) = *hex_str
		*(*uintptr)(unsafe.Pointer(&hex_str))++
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&hex)))) + uintptr(int32(1)))) = *hex_str
		*(*uintptr)(unsafe.Pointer(&hex_str))++
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(i))) = uint8(PX_htoi((*int8)(unsafe.Pointer(&hex))))
		i++
	}
	return i
}

func PX_STRINGFORMAT_INT(_i PX_int) PX_stringformat {
	var fmt PX_stringformat
	fmt.type_ = int32(0)
	*(*int32)(unsafe.Pointer(&fmt._innerPX_stringformat)) = _i
	return fmt
}

func PX_STRINGFORMAT_FLOAT(_f PX_float) PX_stringformat {
	var fmt PX_stringformat
	fmt.type_ = int32(1)
	*(*float32)(unsafe.Pointer(&fmt._innerPX_stringformat)) = _f
	return fmt
}

func PX_STRINGFORMAT_STRING(_s *PX_char) PX_stringformat {
	var fmt PX_stringformat
	fmt.type_ = int32(2)
	fmt._pstring = _s
	return fmt
}

func PX_sprintf8(_out_str *PX_char, str_size PX_int, fmt *PX_char, _1 PX_stringformat, _2 PX_stringformat, _3 PX_stringformat, _4 PX_stringformat, _5 PX_stringformat, _6 PX_stringformat, _7 PX_stringformat, _8 PX_stringformat) PX_int {
	var length int32 = int32(0)
	var p *int8 = nil
	var tret PX_RETURN_STRING
	var pstringfmt PX_stringformat
	var precision int32 = int32(3)
	if !(_out_str != nil) || !(str_size != 0) {
		for p = fmt; *p != 0; *(*uintptr)(unsafe.Pointer(&p))++ {
			if int32(*p) != '%' {
				length++
				continue
			}
			switch int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))))) {
			case '1':
				pstringfmt = _1
				break
			case '2':
				pstringfmt = _2
				break
			case '3':
				pstringfmt = _3
				break
			case '4':
				pstringfmt = _4
				break
			case '5':
				pstringfmt = _5
				break
			case '6':
				pstringfmt = _6
				break
			case '7':
				pstringfmt = _7
				break
			case '8':
				pstringfmt = _8
				break
			default:
				length++
				continue
			}
			if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(2))))) == '.' && PX_charIsNumeric(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(3))))) != 0 {
				precision = int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(3))))) - '0'
				*(*uintptr)(unsafe.Pointer(&p)) += uintptr(int32(3))
			} else {
				*(*uintptr)(unsafe.Pointer(&p))++
			}
			switch uint32(pstringfmt.type_) {
			case uint32(0):
				tret = PX_itos(*(*int32)(unsafe.Pointer(&pstringfmt._innerPX_stringformat)), int32(10))
				length += PX_strlen((*int8)(unsafe.Pointer(&(&tret).data)))
				break
			case uint32(1):
				tret = PX_ftos(*(*float32)(unsafe.Pointer(&pstringfmt._innerPX_stringformat)), precision)
				length += PX_strlen((*int8)(unsafe.Pointer(&(&tret).data)))
				break
			case uint32(2):
				length += PX_strlen(pstringfmt._pstring)
				break
			default:
				return int32(0)
			}
		}
		return length
	}
	PX_memset(unsafe.Pointer(_out_str), uint8(0), str_size)
	for p = fmt; *p != 0; *(*uintptr)(unsafe.Pointer(&p))++ {
		if int32(*p) != '%' {
			*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(_out_str)) + uintptr(length))) = *p
			length++
			if length >= str_size {
				PX_ASSERT()
			}
			continue
		}
		switch int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))))) {
		case '1':
			pstringfmt = _1
			break
		case '2':
			pstringfmt = _2
			break
		case '3':
			pstringfmt = _3
			break
		case '4':
			pstringfmt = _4
			break
		case '5':
			pstringfmt = _5
			break
		case '6':
			pstringfmt = _6
			break
		case '7':
			pstringfmt = _7
			break
		case '8':
			pstringfmt = _8
			break
		default:
			*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(_out_str)) + uintptr(length))) = *p
			length++
			if length >= str_size {
				PX_ASSERT()
			}
			continue
		}
		if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(2))))) == '.' && PX_charIsNumeric(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(3))))) != 0 {
			precision = int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(3))))) - '0'
			*(*uintptr)(unsafe.Pointer(&p)) += uintptr(int32(3))
		} else {
			*(*uintptr)(unsafe.Pointer(&p))++
		}
		switch uint32(pstringfmt.type_) {
		case uint32(0):
			tret = PX_itos(*(*int32)(unsafe.Pointer(&pstringfmt._innerPX_stringformat)), int32(10))
			if length+PX_strlen((*int8)(unsafe.Pointer(&tret.data))) < str_size {
				PX_strcat(_out_str, (*int8)(unsafe.Pointer(&tret.data)))
				length += PX_strlen((*int8)(unsafe.Pointer(&tret.data)))
			} else {
				return length
			}
			break
		case uint32(1):
			tret = PX_ftos(*(*float32)(unsafe.Pointer(&pstringfmt._innerPX_stringformat)), precision)
			if length+PX_strlen((*int8)(unsafe.Pointer(&tret.data))) < str_size {
				PX_strcat(_out_str, (*int8)(unsafe.Pointer(&tret.data)))
				length += PX_strlen((*int8)(unsafe.Pointer(&tret.data)))
			} else {
				return length
			}
			break
		case uint32(2):
			if length+PX_strlen(pstringfmt._pstring) < str_size {
				PX_strcat(_out_str, pstringfmt._pstring)
				length += PX_strlen(pstringfmt._pstring)
			} else {
				return length
			}
			break
		default:
			return int32(0)
		}
	}
	if length >= str_size {
		PX_ASSERT()
	}
	return length
}

func PX_sprintf7(str *PX_char, str_size PX_int, fmt *PX_char, _1 PX_stringformat, _2 PX_stringformat, _3 PX_stringformat, _4 PX_stringformat, _5 PX_stringformat, _6 PX_stringformat, _7 PX_stringformat) PX_int {
	return PX_sprintf8(str, str_size, fmt, _1, _2, _3, _4, _5, _6, _7, PX_STRINGFORMAT_INT(int32(0)))
}

func PX_sprintf6(str *PX_char, str_size PX_int, fmt *PX_char, _1 PX_stringformat, _2 PX_stringformat, _3 PX_stringformat, _4 PX_stringformat, _5 PX_stringformat, _6 PX_stringformat) PX_int {
	return PX_sprintf8(str, str_size, fmt, _1, _2, _3, _4, _5, _6, PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)))
}

func PX_sprintf5(str *PX_char, str_size PX_int, fmt *PX_char, _1 PX_stringformat, _2 PX_stringformat, _3 PX_stringformat, _4 PX_stringformat, _5 PX_stringformat) PX_int {
	return PX_sprintf8(str, str_size, fmt, _1, _2, _3, _4, _5, PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)))
}

func PX_sprintf4(str *PX_char, str_size PX_int, fmt *PX_char, _1 PX_stringformat, _2 PX_stringformat, _3 PX_stringformat, _4 PX_stringformat) PX_int {
	return PX_sprintf8(str, str_size, fmt, _1, _2, _3, _4, PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)))
}

func PX_sprintf3(str *PX_char, str_size PX_int, fmt *PX_char, _1 PX_stringformat, _2 PX_stringformat, _3 PX_stringformat) PX_int {
	return PX_sprintf8(str, str_size, fmt, _1, _2, _3, PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)))
}

func PX_sprintf2(str *PX_char, str_size PX_int, fmt *PX_char, _1 PX_stringformat, _2 PX_stringformat) PX_int {
	return PX_sprintf8(str, str_size, fmt, _1, _2, PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)))
}

func PX_sprintf1(str *PX_char, str_size PX_int, fmt *PX_char, _1 PX_stringformat) PX_int {
	return PX_sprintf8(str, str_size, fmt, _1, PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)))
}

func PX_sprintf0(str *PX_char, str_size PX_int, fmt *PX_char) PX_int {
	return PX_sprintf8(str, str_size, fmt, PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)))
}

func PX_memset(dst unsafe.Pointer, byte PX_byte, size PX_int) {
	var dw uint32 = uint32(func() int32 {
		if int32(byte) != 0 {
			return int32(byte)<<int32(24) | int32(byte)<<int32(16) | int32(byte)<<int32(8) | int32(byte)
		} else {
			return int32(0)
		}
	}())
	var _4byteMovDst *uint32 = (*uint32)(dst)
	var pdst *uint8 = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(dst))) + uintptr(size&-4)))
	var _movTs uint32 = uint32(size >> int32(2))
	for func() (_cgo_ret uint32) {
		_cgo_addr := &_movTs
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0 {
		*func() (_cgo_ret *uint32) {
			_cgo_addr := &_4byteMovDst
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
			return
		}() = dw
	}
	_movTs = uint32(size & int32(3))
	for func() (_cgo_ret uint32) {
		_cgo_addr := &_movTs
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0 {
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &pdst
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = byte
	}
}

func PX_memdwordset(dst unsafe.Pointer, dw PX_dword, count PX_int) {
	var p *uint32 = (*uint32)(dst)
	for func() (_cgo_ret int32) {
		_cgo_addr := &count
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0 {
		*func() (_cgo_ret *uint32) {
			_cgo_addr := &p
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
			return
		}() = dw
	}
}

func PX_memequ(dst unsafe.Pointer, src unsafe.Pointer, size PX_int) PX_bool {
	var _4byteMovSrc *uint32 = (*uint32)(src)
	var _4byteMovDst *uint32 = (*uint32)(dst)
	var psrc *uint8 = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(src))) + uintptr(size&-4)))
	var pdst *uint8 = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(dst))) + uintptr(size&-4)))
	var _movTs uint32 = uint32(size >> int32(2))
	if uintptr(unsafe.Pointer(dst)) == uintptr(unsafe.Pointer(nil)) || uintptr(unsafe.Pointer(src)) == uintptr(unsafe.Pointer(nil)) {
		PX_ASSERT()
		return int32(0)
	}
	for func() (_cgo_ret uint32) {
		_cgo_addr := &_movTs
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0 {
		if *func() (_cgo_ret *uint32) {
			_cgo_addr := &_4byteMovDst
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
			return
		}() != *func() (_cgo_ret *uint32) {
			_cgo_addr := &_4byteMovSrc
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
			return
		}() {
			return int32(0)
		}
	}
	_movTs = uint32(size & int32(3))
	for func() (_cgo_ret uint32) {
		_cgo_addr := &_movTs
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0 {
		if int32(*func() (_cgo_ret *uint8) {
			_cgo_addr := &pdst
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()) != int32(*func() (_cgo_ret *uint8) {
			_cgo_addr := &psrc
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()) {
			return int32(0)
		}
	}
	return int32(1)
}

func PX_memcpy(dst unsafe.Pointer, src unsafe.Pointer, size PX_int) {
	type _px_memcpy_16 struct {
		m [16]uint8
	}
	type PX_MEMCPY_16 = _px_memcpy_16
	type _px_memcpy_32 struct {
		m [32]uint8
	}
	type PX_MEMCPY_32 = _px_memcpy_32
	type _px_memcpy_64 struct {
		m [64]uint8
	}
	type PX_MEMCPY_64 = _px_memcpy_64
	type _px_memcpy_128 struct {
		m [128]uint8
	}
	type PX_MEMCPY_128 = _px_memcpy_128
	type _px_memcpy_256 struct {
		m [256]uint8
	}
	type PX_MEMCPY_256 = _px_memcpy_256
	type _px_memcpy_512 struct {
		m [512]uint8
	}
	type PX_MEMCPY_512 = _px_memcpy_512
	type _px_memcpy_1024 struct {
		m [1024]uint8
	}
	type PX_MEMCPY_1024 = _px_memcpy_1024
	type _px_memcpy_2048 struct {
		m [2048]uint8
	}
	type PX_MEMCPY_2048 = _px_memcpy_2048
	type _px_memcpy_4096 struct {
		m [4096]uint8
	}
	type PX_MEMCPY_4096 = _px_memcpy_4096
	var _4byteMovSrc *uint32
	var _4byteMovDst *uint32
	var psrc *uint8
	var pdst *uint8
	var _4kbyteMovSrc *_px_memcpy_4096
	var _4kbyteMovDst *_px_memcpy_4096
	var _2kbyteMovSrc *_px_memcpy_2048
	var _2kbyteMovDst *_px_memcpy_2048
	var _1kbyteMovSrc *_px_memcpy_1024
	var _1kbyteMovDst *_px_memcpy_1024
	var _512byteMovSrc *_px_memcpy_512
	var _512byteMovDst *_px_memcpy_512
	var _256byteMovSrc *_px_memcpy_256
	var _256byteMovDst *_px_memcpy_256
	var _128byteMovSrc *_px_memcpy_128
	var _128byteMovDst *_px_memcpy_128
	var _64byteMovSrc *_px_memcpy_64
	var _64byteMovDst *_px_memcpy_64
	var _32byteMovSrc *_px_memcpy_32
	var _32byteMovDst *_px_memcpy_32
	var _16byteMovSrc *_px_memcpy_16
	var _16byteMovDst *_px_memcpy_16
	var _movTs uint32
	if size <= int32(0) {
		return
	}
	if uintptr(unsafe.Pointer(dst)) > uintptr(unsafe.Pointer(unsafe.Pointer(src))) && uintptr(unsafe.Pointer((*int8)(dst))) < uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(src)))+uintptr(size))))) {
		var _4byteMov uint32 = uint32(0)
		psrc = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(src)))+uintptr(size))))) - uintptr(int32(1))))
		pdst = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(dst)))+uintptr(size))))) - uintptr(int32(1))))
		_movTs = uint32(size & int32(3))
		for func() (_cgo_ret uint32) {
			_cgo_addr := &_movTs
			_cgo_ret = *_cgo_addr
			*_cgo_addr--
			return
		}() != 0 {
			*func() (_cgo_ret *uint8) {
				_cgo_addr := &pdst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))--
				return
			}() = *func() (_cgo_ret *uint8) {
				_cgo_addr := &psrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))--
				return
			}()
		}
		*(*uintptr)(unsafe.Pointer(&pdst)) -= uintptr(int32(3))
		*(*uintptr)(unsafe.Pointer(&psrc)) -= uintptr(int32(3))
		_4byteMovDst = (*uint32)(unsafe.Pointer(pdst))
		_4byteMovSrc = (*uint32)(unsafe.Pointer(psrc))
		_movTs = uint32(size >> int32(2))
		for func() (_cgo_ret uint32) {
			_cgo_addr := &_movTs
			_cgo_ret = *_cgo_addr
			*_cgo_addr--
			return
		}() != 0 {
			_4byteMov = *_4byteMovSrc
			*(*uintptr)(unsafe.Pointer(&_4byteMovSrc)) -= 4
			*_4byteMovDst = _4byteMov
			*(*uintptr)(unsafe.Pointer(&_4byteMovDst)) -= 4
		}
	} else if uintptr(unsafe.Pointer(src)) > uintptr(unsafe.Pointer(unsafe.Pointer(dst))) && uintptr(unsafe.Pointer((*int8)(src))) < uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(dst)))+uintptr(size))))) {
		var _4byteMov uint32 = uint32(0)
		psrc = (*uint8)(src)
		pdst = (*uint8)(dst)
		_4byteMovDst = (*uint32)(unsafe.Pointer(pdst))
		_4byteMovSrc = (*uint32)(unsafe.Pointer(psrc))
		_movTs = uint32(size >> int32(2))
		for func() (_cgo_ret uint32) {
			_cgo_addr := &_movTs
			_cgo_ret = *_cgo_addr
			*_cgo_addr--
			return
		}() != 0 {
			_4byteMov = *_4byteMovSrc
			*(*uintptr)(unsafe.Pointer(&_4byteMovSrc)) += 4
			*_4byteMovDst = _4byteMov
			*(*uintptr)(unsafe.Pointer(&_4byteMovDst)) += 4
		}
		_movTs = uint32(size & int32(3))
		psrc = (*uint8)(unsafe.Pointer(_4byteMovSrc))
		pdst = (*uint8)(unsafe.Pointer(_4byteMovDst))
		for func() (_cgo_ret uint32) {
			_cgo_addr := &_movTs
			_cgo_ret = *_cgo_addr
			*_cgo_addr--
			return
		}() != 0 {
			*func() (_cgo_ret *uint8) {
				_cgo_addr := &pdst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}() = *func() (_cgo_ret *uint8) {
				_cgo_addr := &psrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}()
		}
		*(*uintptr)(unsafe.Pointer(&pdst)) += uintptr(int32(3))
		*(*uintptr)(unsafe.Pointer(&psrc)) += uintptr(int32(3))
	} else {
		_movTs = uint32(size >> int32(12))
		if _movTs != 0 {
			_4kbyteMovSrc = (*_px_memcpy_4096)(src)
			_4kbyteMovDst = (*_px_memcpy_4096)(dst)
			for func() (_cgo_ret uint32) {
				_cgo_addr := &_movTs
				_cgo_ret = *_cgo_addr
				*_cgo_addr--
				return
			}() != 0 {
				*func() (_cgo_ret *_px_memcpy_4096) {
					_cgo_addr := &_4kbyteMovDst
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4096
					return
				}() = *func() (_cgo_ret *_px_memcpy_4096) {
					_cgo_addr := &_4kbyteMovSrc
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4096
					return
				}()
			}
			src = unsafe.Pointer(_4kbyteMovSrc)
			dst = unsafe.Pointer(_4kbyteMovDst)
		}
		_movTs = uint32(size & 2048)
		if _movTs != 0 {
			_2kbyteMovSrc = (*_px_memcpy_2048)(src)
			_2kbyteMovDst = (*_px_memcpy_2048)(dst)
			*func() (_cgo_ret *_px_memcpy_2048) {
				_cgo_addr := &_2kbyteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 2048
				return
			}() = *func() (_cgo_ret *_px_memcpy_2048) {
				_cgo_addr := &_2kbyteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 2048
				return
			}()
			src = unsafe.Pointer(_2kbyteMovSrc)
			dst = unsafe.Pointer(_2kbyteMovDst)
		}
		_movTs = uint32(size & 1024)
		if _movTs != 0 {
			_1kbyteMovSrc = (*_px_memcpy_1024)(src)
			_1kbyteMovDst = (*_px_memcpy_1024)(dst)
			*func() (_cgo_ret *_px_memcpy_1024) {
				_cgo_addr := &_1kbyteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 1024
				return
			}() = *func() (_cgo_ret *_px_memcpy_1024) {
				_cgo_addr := &_1kbyteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 1024
				return
			}()
			src = unsafe.Pointer(_1kbyteMovSrc)
			dst = unsafe.Pointer(_1kbyteMovDst)
		}
		_movTs = uint32(size & 512)
		if _movTs != 0 {
			_512byteMovSrc = (*_px_memcpy_512)(src)
			_512byteMovDst = (*_px_memcpy_512)(dst)
			*func() (_cgo_ret *_px_memcpy_512) {
				_cgo_addr := &_512byteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 512
				return
			}() = *func() (_cgo_ret *_px_memcpy_512) {
				_cgo_addr := &_512byteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 512
				return
			}()
			src = unsafe.Pointer(_512byteMovSrc)
			dst = unsafe.Pointer(_512byteMovDst)
		}
		_movTs = uint32(size & 256)
		if _movTs != 0 {
			_256byteMovSrc = (*_px_memcpy_256)(src)
			_256byteMovDst = (*_px_memcpy_256)(dst)
			*func() (_cgo_ret *_px_memcpy_256) {
				_cgo_addr := &_256byteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 256
				return
			}() = *func() (_cgo_ret *_px_memcpy_256) {
				_cgo_addr := &_256byteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 256
				return
			}()
			src = unsafe.Pointer(_256byteMovSrc)
			dst = unsafe.Pointer(_256byteMovDst)
		}
		_movTs = uint32(size & 128)
		if _movTs != 0 {
			_128byteMovSrc = (*_px_memcpy_128)(src)
			_128byteMovDst = (*_px_memcpy_128)(dst)
			*func() (_cgo_ret *_px_memcpy_128) {
				_cgo_addr := &_128byteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 128
				return
			}() = *func() (_cgo_ret *_px_memcpy_128) {
				_cgo_addr := &_128byteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 128
				return
			}()
			src = unsafe.Pointer(_128byteMovSrc)
			dst = unsafe.Pointer(_128byteMovDst)
		}
		_movTs = uint32(size & 64)
		if _movTs != 0 {
			_64byteMovSrc = (*_px_memcpy_64)(src)
			_64byteMovDst = (*_px_memcpy_64)(dst)
			*func() (_cgo_ret *_px_memcpy_64) {
				_cgo_addr := &_64byteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 64
				return
			}() = *func() (_cgo_ret *_px_memcpy_64) {
				_cgo_addr := &_64byteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 64
				return
			}()
			src = unsafe.Pointer(_64byteMovSrc)
			dst = unsafe.Pointer(_64byteMovDst)
		}
		_movTs = uint32(size & 32)
		if _movTs != 0 {
			_32byteMovSrc = (*_px_memcpy_32)(src)
			_32byteMovDst = (*_px_memcpy_32)(dst)
			*func() (_cgo_ret *_px_memcpy_32) {
				_cgo_addr := &_32byteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 32
				return
			}() = *func() (_cgo_ret *_px_memcpy_32) {
				_cgo_addr := &_32byteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 32
				return
			}()
			src = unsafe.Pointer(_32byteMovSrc)
			dst = unsafe.Pointer(_32byteMovDst)
		}
		_movTs = uint32(size & 16)
		if _movTs != 0 {
			_16byteMovSrc = (*_px_memcpy_16)(src)
			_16byteMovDst = (*_px_memcpy_16)(dst)
			*func() (_cgo_ret *_px_memcpy_16) {
				_cgo_addr := &_16byteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 16
				return
			}() = *func() (_cgo_ret *_px_memcpy_16) {
				_cgo_addr := &_16byteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 16
				return
			}()
			src = unsafe.Pointer(_16byteMovSrc)
			dst = unsafe.Pointer(_16byteMovDst)
		}
		_movTs = uint32(size & int32(15))
		if _movTs >= uint32(12) {
			_4byteMovSrc = (*uint32)(src)
			_4byteMovDst = (*uint32)(dst)
			*func() (_cgo_ret *uint32) {
				_cgo_addr := &_4byteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}() = *func() (_cgo_ret *uint32) {
				_cgo_addr := &_4byteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}()
			*func() (_cgo_ret *uint32) {
				_cgo_addr := &_4byteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}() = *func() (_cgo_ret *uint32) {
				_cgo_addr := &_4byteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}()
			*func() (_cgo_ret *uint32) {
				_cgo_addr := &_4byteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}() = *func() (_cgo_ret *uint32) {
				_cgo_addr := &_4byteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}()
			_movTs -= uint32(12)
			psrc = (*uint8)(unsafe.Pointer(_4byteMovSrc))
			pdst = (*uint8)(unsafe.Pointer(_4byteMovDst))
			for func() (_cgo_ret uint32) {
				_cgo_addr := &_movTs
				_cgo_ret = *_cgo_addr
				*_cgo_addr--
				return
			}() != 0 {
				*func() (_cgo_ret *uint8) {
					_cgo_addr := &pdst
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}() = *func() (_cgo_ret *uint8) {
					_cgo_addr := &psrc
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}()
			}
		} else if _movTs >= uint32(8) {
			_4byteMovSrc = (*uint32)(src)
			_4byteMovDst = (*uint32)(dst)
			*func() (_cgo_ret *uint32) {
				_cgo_addr := &_4byteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}() = *func() (_cgo_ret *uint32) {
				_cgo_addr := &_4byteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}()
			*func() (_cgo_ret *uint32) {
				_cgo_addr := &_4byteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}() = *func() (_cgo_ret *uint32) {
				_cgo_addr := &_4byteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}()
			_movTs -= uint32(8)
			psrc = (*uint8)(unsafe.Pointer(_4byteMovSrc))
			pdst = (*uint8)(unsafe.Pointer(_4byteMovDst))
			for func() (_cgo_ret uint32) {
				_cgo_addr := &_movTs
				_cgo_ret = *_cgo_addr
				*_cgo_addr--
				return
			}() != 0 {
				*func() (_cgo_ret *uint8) {
					_cgo_addr := &pdst
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}() = *func() (_cgo_ret *uint8) {
					_cgo_addr := &psrc
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}()
			}
		} else if _movTs >= uint32(4) {
			_4byteMovSrc = (*uint32)(src)
			_4byteMovDst = (*uint32)(dst)
			*func() (_cgo_ret *uint32) {
				_cgo_addr := &_4byteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}() = *func() (_cgo_ret *uint32) {
				_cgo_addr := &_4byteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}()
			_movTs -= uint32(4)
			psrc = (*uint8)(unsafe.Pointer(_4byteMovSrc))
			pdst = (*uint8)(unsafe.Pointer(_4byteMovDst))
			for func() (_cgo_ret uint32) {
				_cgo_addr := &_movTs
				_cgo_ret = *_cgo_addr
				*_cgo_addr--
				return
			}() != 0 {
				*func() (_cgo_ret *uint8) {
					_cgo_addr := &pdst
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}() = *func() (_cgo_ret *uint8) {
					_cgo_addr := &psrc
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}()
			}
		} else {
			psrc = (*uint8)(src)
			pdst = (*uint8)(dst)
			for func() (_cgo_ret uint32) {
				_cgo_addr := &_movTs
				_cgo_ret = *_cgo_addr
				*_cgo_addr--
				return
			}() != 0 {
				*func() (_cgo_ret *uint8) {
					_cgo_addr := &pdst
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}() = *func() (_cgo_ret *uint8) {
					_cgo_addr := &psrc
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}()
			}
		}
	}
}

func PX_strcpy(dst *PX_char, src *PX_char, size PX_int) {
	for func() (_cgo_ret int32) {
		_cgo_addr := &size
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0 {
		if *src != 0 {
			*func() (_cgo_ret *int8) {
				_cgo_addr := &dst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}() = *func() (_cgo_ret *int8) {
				_cgo_addr := &src
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}()
		} else {
			*dst = int8('\x00')
			return
		}
	}
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(dst)) - uintptr(int32(1)))) = int8('\x00')
}

func PX_wstrcpy(dst *PX_word, src *PX_word, size PX_int) {
	for func() (_cgo_ret int32) {
		_cgo_addr := &size
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0 {
		if *src != 0 {
			*func() (_cgo_ret *uint16) {
				_cgo_addr := &dst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 2
				return
			}() = *func() (_cgo_ret *uint16) {
				_cgo_addr := &src
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 2
				return
			}()
		} else {
			*dst = uint16('\x00')
			return
		}
	}
	*(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(dst)) - uintptr(int32(1))*2)) = uint16('\x00')
}

func PX_strcat(src *PX_char, cat *PX_char) {
	var len int32 = PX_strlen(cat)
	for *src != 0 {
		*(*uintptr)(unsafe.Pointer(&src))++
	}
	for func() (_cgo_ret int32) {
		_cgo_addr := &len
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0 {
		*func() (_cgo_ret *int8) {
			_cgo_addr := &src
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *int8) {
			_cgo_addr := &cat
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
	}
	*src = int8('\x00')
}

func PX_strcat_s(src *PX_char, size PX_int, cat *PX_char) {
	if PX_strlen(src)+PX_strlen(cat) < size {
		PX_strcat(src, cat)
	}
}

func PX_wstrcat(src *PX_word, cat *PX_word) {
	var len int32 = PX_wstrlen(cat)
	for *src != 0 {
		*(*uintptr)(unsafe.Pointer(&src)) += 2
	}
	for func() (_cgo_ret int32) {
		_cgo_addr := &len
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0 {
		*func() (_cgo_ret *uint16) {
			_cgo_addr := &src
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 2
			return
		}() = *func() (_cgo_ret *uint16) {
			_cgo_addr := &cat
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 2
			return
		}()
	}
	*src = uint16('\x00')
}

func PX_strset(dst *PX_char, src *PX_char) {
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(dst)) + uintptr(int32(0)))) = int8(0)
	PX_strcat(dst, src)
}

func PX_strlen(dst *PX_char) PX_int {
	var len int32 = int32(0)
	if !(dst != nil) {
		PX_ERROR((*int8)(unsafe.Pointer(&[21]int8{'N', 'U', 'L', 'L', ' ', 'p', 'o', 'i', 'n', 't', 'e', 'r', ' ', 'a', 's', 's', 'e', 'r', 't', '!', '\x00'})))
		return int32(0)
	}
	for *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(dst)) + uintptr(func() (_cgo_ret int32) {
		_cgo_addr := &len
		_cgo_ret = *_cgo_addr
		*_cgo_addr++
		return
	}()))) != 0 {
	}
	return len - int32(1)
}

func PX_wstrlen(dst *PX_word) PX_int {
	var len int32 = int32(0)
	if !(dst != nil) {
		PX_ERROR((*int8)(unsafe.Pointer(&[21]int8{'N', 'U', 'L', 'L', ' ', 'p', 'o', 'i', 'n', 't', 'e', 'r', ' ', 'a', 's', 's', 'e', 'r', 't', '!', '\x00'})))
		return int32(0)
	}
	for *(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(dst)) + uintptr(func() (_cgo_ret int32) {
		_cgo_addr := &len
		_cgo_ret = *_cgo_addr
		*_cgo_addr++
		return
	}())*2)) != 0 {
	}
	return len - int32(1)
}

func PX_strequ2(src *PX_char, dst *PX_char) PX_bool {
	var _l int8
	var _r int8
	for int32(1) != 0 {
		_l = *src
		_r = *dst
		if int32(_l) >= 'a' && int32(_l) <= 'z' {
			_l += int8(-32)
		}
		if int32(_r) >= 'a' && int32(_r) <= 'z' {
			_r += int8(-32)
		}
		if int32(_l) == int32(_r) {
			if int32(_l) == int32(0) {
				return int32(1)
			}
		} else {
			return int32(0)
		}
		*(*uintptr)(unsafe.Pointer(&src))++
		*(*uintptr)(unsafe.Pointer(&dst))++
	}
	return int32(0)
}

func PX_strupr(src *PX_char) {
	for int32(*src) != '\x00' {
		if int32(*src) >= 'a' && int32(*src) <= 'z' {
			*src -= int8(32)
		}
		*(*uintptr)(unsafe.Pointer(&src))++
	}
}

func PX_strlwr(src *PX_char) {
	for int32(*src) != '\x00' {
		if int32(*src) > 'A' && int32(*src) <= 'Z' {
			*src += int8(32)
		}
		*(*uintptr)(unsafe.Pointer(&src))++
	}
}

func PX_strequ(src *PX_char, dst *PX_char) PX_bool {
	var ret int32 = int32(0)
	for !(func() (_cgo_ret int32) {
		_cgo_addr := &ret
		*_cgo_addr = int32(*(*uint8)(unsafe.Pointer(src))) - int32(*(*uint8)(unsafe.Pointer(dst)))
		return *_cgo_addr
	}() != 0) && int32(*src) != 0 {
		*(*uintptr)(unsafe.Pointer(&src))++
		*(*uintptr)(unsafe.Pointer(&dst))++
	}
	return func() int32 {
		if !(ret != 0) {
			return 1
		} else {
			return 0
		}
	}()
}

func PX_strIsNumeric(str *PX_char) PX_bool {
	var Dot int32 = int32(0)
	var CurrentCharIndex int32
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(int32(0))))) == '\x00' {
		return int32(0)
	}
	if !(PX_charIsNumeric(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(int32(0))))) != 0) && !(int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(int32(0))))) == '-') {
		if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(int32(0))))) != '.' {
			return int32(0)
		}
	}
	for CurrentCharIndex = int32(1); CurrentCharIndex < PX_strlen(str); CurrentCharIndex++ {
		if !(PX_charIsNumeric(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(CurrentCharIndex)))) != 0) {
			if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(CurrentCharIndex)))) == 'e' && (int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(CurrentCharIndex+int32(1))))) == '+' || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(CurrentCharIndex+int32(1))))) == '-') {
				CurrentCharIndex++
				continue
			}
			if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(CurrentCharIndex)))) == '.' {
				if Dot != 0 {
					return int32(0)
				} else {
					Dot = int32(1)
					continue
				}
			}
			return int32(0)
		}
	}
	return int32(1)
}

func PX_strIsFloat(str *PX_char) PX_bool {
	var Dot int32 = int32(0)
	var CurrentCharIndex int32
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(int32(0))))) == '\x00' {
		return int32(0)
	}
	if !(PX_charIsNumeric(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(int32(0))))) != 0) && !(int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(int32(0))))) == '-') {
		if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(int32(0))))) != '.' {
			return int32(0)
		}
	}
	for CurrentCharIndex = int32(1); CurrentCharIndex < PX_strlen(str); CurrentCharIndex++ {
		if !(PX_charIsNumeric(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(CurrentCharIndex)))) != 0) {
			if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(CurrentCharIndex)))) == 'e' && (int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(CurrentCharIndex+int32(1))))) == '+' || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(CurrentCharIndex+int32(1))))) == '-') {
				CurrentCharIndex++
				continue
			}
			if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(CurrentCharIndex)))) == '.' {
				if Dot != 0 {
					return int32(0)
				} else {
					Dot = int32(1)
					continue
				}
			}
			return int32(0)
		}
	}
	return func() int32 {
		if Dot != 0 {
			return int32(1)
		} else {
			return int32(0)
		}
	}()
}

func PX_charIsNumeric(chr PX_char) PX_bool {
	if int32(chr) >= '0' && int32(chr) <= '9' {
		return int32(1)
	}
	return int32(0)
}

func PX_strIsInt(str *PX_char) PX_bool {
	if PX_strIsNumeric(str) != 0 {
		if PX_strIsFloat(str) != 0 {
			return int32(0)
		}
		return int32(1)
	}
	return int32(0)
}

func PX_pow_ii(i PX_int, n PX_int) PX_int {
	var fin int32 = int32(1)
	if n == int32(0) {
		return int32(1)
	}
	if n < int32(0) {
		return int32(0)
	}
	for func() (_cgo_ret int32) {
		_cgo_addr := &n
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0 {
		fin *= i
	}
	return fin
}

func PX_strchr(s *PX_char, ch PX_int) *PX_char {
	for int32(*s) != '\x00' {
		if int32(*s)-ch == int32(0) {
			return (*int8)(unsafe.Pointer(s))
		}
		*(*uintptr)(unsafe.Pointer(&s))++
	}
	return (*int8)(nil)
}

func PX_strcmp(str1 *PX_char, str2 *PX_char) PX_int {
	var ret int32 = int32(0)
	for !(func() (_cgo_ret int32) {
		_cgo_addr := &ret
		*_cgo_addr = int32(*(*uint8)(unsafe.Pointer(str1))) - int32(*(*uint8)(unsafe.Pointer(str2)))
		return *_cgo_addr
	}() != 0) && int32(*str1) != 0 {
		*(*uintptr)(unsafe.Pointer(&str1))++
		*(*uintptr)(unsafe.Pointer(&str2))++
	}
	return ret
}

func PX_strstr(dest *PX_char, src *PX_char) *PX_char {
	var start *int8 = (*int8)(unsafe.Pointer(dest))
	var substart *int8 = (*int8)(unsafe.Pointer(src))
	var cp *int8 = (*int8)(unsafe.Pointer(dest))
	for *cp != 0 {
		start = cp
		for int32(*start) != '\x00' && int32(*substart) != '\x00' && int32(*start) == int32(*substart) {
			*(*uintptr)(unsafe.Pointer(&start))++
			*(*uintptr)(unsafe.Pointer(&substart))++
		}
		if int32(*substart) == '\x00' {
			return cp
		}
		substart = (*int8)(unsafe.Pointer(src))
		*(*uintptr)(unsafe.Pointer(&cp))++
	}
	return (*int8)(nil)
}

func PX_strcut(dest *PX_char, left PX_int, right PX_int) {
	var len int32 = PX_strlen(dest)
	if len == int32(0) {
		return
	}
	if right < left {
		var t int32 = left
		left = right
		right = t
	}
	if left < int32(0) {
		left = int32(0)
	}
	if left > len-int32(1) {
		left = len - int32(1)
	}
	if right < int32(0) {
		right = int32(0)
	}
	if right > len-int32(1) {
		right = len - int32(1)
	}
	PX_memcpy(unsafe.Pointer(dest), unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(dest))+uintptr(left)))), right-left+int32(1))
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(dest)) + uintptr(right-left+int32(1)))) = int8(0)
}

func PX_htonl(h PX_dword) PX_dword {
	return func() uint32 {
		if PX_isBigEndianCPU() != 0 {
			return h
		} else {
			return uint32(h)&uint32(4278190080)>>int32(24) | uint32(h)&uint32(16711680)>>int32(8) | uint32(h)&uint32(65280)<<int32(8) | uint32(h)&uint32(255)<<int32(24)
		}
	}()
}

func PX_ntohl(n PX_dword) PX_dword {
	return func() uint32 {
		if PX_isBigEndianCPU() != 0 {
			return n
		} else {
			return uint32(n)&uint32(4278190080)>>int32(24) | uint32(n)&uint32(16711680)>>int32(8) | uint32(n)&uint32(65280)<<int32(8) | uint32(n)&uint32(255)<<int32(24)
		}
	}()
}

func PX_htons(h PX_dword) PX_word {
	return uint16(func() int32 {
		if PX_isBigEndianCPU() != 0 {
			return int32(h)
		} else {
			return int32(uint16(h))&int32(65280)>>int32(8) | int32(uint16(h))&int32(255)<<int32(8)
		}
	}())
}

func PX_ntohs(n PX_dword) PX_word {
	return uint16(func() int32 {
		if PX_isBigEndianCPU() != 0 {
			return int32(n)
		} else {
			return int32(uint16(n))&int32(65280)>>int32(8) | int32(uint16(n))&int32(255)<<int32(8)
		}
	}())
}
