package convert

type Integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

func StrToInt[T Integer](s string) T {
	var val T

	if len(s) == 0 {
		return T(0)
	}

	// negative num
	if s[0] == '-' {
		//s = s[1:]
		// return -StrToInt[T](s)
		// recursive call is slower than loop

		for _, c := range s[1:] {
			val = val*10 + T(c-'0')
		}

		return -val
	}

	for _, c := range s {
		val = val*10 + T(c-'0')
	}

	return val
}
