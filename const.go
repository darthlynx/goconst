package goconst

// Math constants
const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1

	MaxInt32 = 1<<31 - 1
	MinInt32 = -1 << 31
)
