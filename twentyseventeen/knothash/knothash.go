package knothash

import "fmt"

func KnotHashSparse(input []byte, rounds int) [256]byte {
	var values [256]byte
	for i := range values {
		values[i] = byte(i)
	}

	var pos, skip byte
	for round := 0; round < rounds; round++ {
		for _, in := range input {
			rot := in / 2
			for i := byte(0); i < rot; i++ {
				left := (pos + i)             // rely on byte rollover
				right := (pos + (in - i) - 1) // rely on byte rollover
				values[left], values[right] = values[right], values[left]
			}
			pos = (pos + skip + in) // rely on byte rollover
			skip++
		}
	}
	return values
}

func KnotHash(input []byte) [16]byte {
	padded := make([]byte, len(input), len(input)+5)
	copy(padded, input)
	sparse := KnotHashSparse(append(padded, 17, 31, 73, 47, 23), 64)
	var blocks [16]byte
	for block := 0; block < 16; block++ {
		for i := 0; i < 16; i++ {
			blocks[block] ^= sparse[(block*16)+i]
		}
	}
	return blocks
}

func KnotHashString(input []byte) string {
	blocks := KnotHash(input)
	return fmt.Sprintf(
		"%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x",
		blocks[0], blocks[1], blocks[2], blocks[3],
		blocks[4], blocks[5], blocks[6], blocks[7],
		blocks[8], blocks[9], blocks[10], blocks[11],
		blocks[12], blocks[13], blocks[14], blocks[15])
}
