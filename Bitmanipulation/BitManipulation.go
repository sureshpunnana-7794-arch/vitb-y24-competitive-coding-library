package main

type BitManipulation struct{}

func (b BitManipulation) GetBit(n int64, k int) int64 {
	return (n >> k) & 1
}

func (b BitManipulation) SetBit(n int64, k int) int64 {
	return (1 << k) | n
}

func (b BitManipulation) ClearBit(n int64, k int) int64 {
	return n & ^(1 << k) // Go uses ^ for bitwise NOT instead of ~
}

func (b BitManipulation) ToggleBit(n int64, k int) int64 {
	return n ^ (1 << k)
}

func (b BitManipulation) IsPowerOfTwo(n int64) bool {
	return n > 0 && (n & (n - 1)) == 0
}
