package cp10

const bitSize = 1 << 5

// NewBitMap ...
func NewBitMap(rangE int) *bitmap {
	return &bitmap{
		bitTable: make([]int, rangE>>bitSize),
	}
}

type bitmap struct {
	bitTable []int
}

// SetBit ...
func (b *bitmap) SetBit(x int) {
	b.bitTable[x>>bitSize] |= 1 << (x % bitSize)
}

// SetBitOnce 只设置一次
func (b *bitmap) SetBitOnce(x int) bool {
	prog, mask := x>>bitSize, 1<<(x%bitSize)
	if b.bitTable[prog]&mask != 1 {
		b.bitTable[prog] |= mask
		return true
	}

	return false
}

// RemoveBit ...
func (b *bitmap) RemoveBit(x int) {
	// 取反 保留其他位信息
	b.bitTable[x>>bitSize] &= ^(1 << (x % bitSize))
}

// Exist ...
func (b *bitmap) Exist(x int) bool {
	return b.bitTable[x>>bitSize]&1<<(x%bitSize) == 1
}
