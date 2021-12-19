package generic

import (
	"constraints"

	"github.com/segmentio/fasthash/fnv1a"
)

// Equaler is a function that returns whether 'a' and 'b' are equal.
type Equaler[T any] func(a, b T) bool

// Lesser is a function that returns whether 'a' is less than 'b'.
type Lesser[T any] func(a, b T) bool

// Hasher is a function that returns the hash of 't'.
type Hasher[T any] func(t T) uint64

// Equals wraps the '==' operator for comparable types.
func Equals[T comparable](a, b T) bool {
	return a == b
}

// Less wraps the '<' operator for ordered types.
func Less[T constraints.Ordered](a, b T) bool {
	return a < b
}

// Compare uses a less function to determine the ordering of 'a' and 'b'. It returns:
//
// * -1 if a < b
//
// * 1 if a > b
//
// * 0 if a == b
func Compare[T any](a, b T, less Lesser[T]) int {
	if less(a, b) {
		return -1
	} else if less(b, a) {
		return 1
	}
	return 0
}

// Max returns the max of a and b.
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Min returns the min of a and b.
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// MaxFunc returns the max of a and b using the less func.
func MaxFunc[T any](a, b T, less Lesser[T]) T {
	if less(b, a) {
		return a
	}
	return b
}

// MinFunc returns the min of a and b using the less func.
func MinFunc[T any](a, b T, less Lesser[T]) T {
	if less(a, b) {
		return a
	}
	return b
}

func HashUint64(u uint64) uint64 {
	return hash(u)
}
func HashUint32(u uint32) uint64 {
	return hash(uint64(u))
}
func HashUint16(u uint16) uint64 {
	return hash(uint64(u))
}
func HashUint8(u uint8) uint64 {
	return hash(uint64(u))
}
func HashInt64(i int64) uint64 {
	return hash(uint64(i))
}
func HashInt32(i int32) uint64 {
	return hash(uint64(i))
}
func HashInt16(i int32) uint64 {
	return hash(uint64(i))
}
func HashInt8(i int32) uint64 {
	return hash(uint64(i))
}
func HashInt(i int) uint64 {
	return hash(uint64(i))
}
func HashUint(i uint) uint64 {
	return hash(uint64(i))
}
func HashString(s string) uint64 {
	return fnv1a.HashString64(s)
}
func HashBytes(b []byte) uint64 {
	return fnv1a.HashBytes64(b)
}

func hash(u uint64) uint64 {
	u ^= u >> 33
	u *= 0xff51afd7ed558ccd
	u ^= u >> 33
	u *= 0xc4ceb9fe1a85ec53
	u ^= u >> 33
	return u
}