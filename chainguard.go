package main

import (
	"encoding/binary"
	"fmt"
)

func main() {

	// Define the constants from the assembly code
	/* LC1:
	.quad 8388875062886221384
	.quad 15534430226834279
	*/
	constant1 := []uint64{
		8388875062886221384, // 0x7465737465737465 "testeteste"
		15534430226834279,   // 0x03706F6C6C6548 containing "Hello" backwards
	}
	constant2 := []byte{2, 2, 2, 2, 2, 2, 2, 2}

	// Create a buffer to hold our data (16 bytes)
	buffer := make([]byte, 16)

	// Copy constant1 to buffer
	binary.LittleEndian.PutUint64(buffer[0:8], constant1[0])
	binary.LittleEndian.PutUint64(buffer[8:16], constant1[1])

	// Perform the transformations

	// 1. XOR with 33686018 (0x02020202) at offset 8
	// xorl $33686018, 8(%rsp)
	value := binary.LittleEndian.Uint32(buffer[8:12])
	value ^= 33686018 // 0x02020202
	binary.LittleEndian.PutUint32(buffer[8:12], value)

	// 2. XOR with constant2
	// The movq instruction is used to load and store 8 bytes (64 bits)
	// movq _constant2(%rip), %xmm0  # Loads 8 bytes
	// pxor %xmm1, %xmm0 # XOR with original data
	originalData := make([]byte, 8)
	copy(originalData, buffer[0:8])

	for i := 0; i < 8; i++ {
		buffer[i] = originalData[i] ^ constant2[i]
	}

	// 3. XOR with 514 (0x0202) at offset 12
	// xorw $514, 12(%rsp)
	value16 := binary.LittleEndian.Uint16(buffer[12:14])
	value16 ^= 514 // 0x0202
	binary.LittleEndian.PutUint16(buffer[12:14], value16)

	// 4. Set byte at offset 14 to 53 (ASCII '5')
	// movb $53, 14(%rsp)
	buffer[14] = 53

	// Print the results
	fmt.Print("Hex: ")
	for _, b := range buffer {
		fmt.Printf("%02x ", b)
	}
	fmt.Println()

	fmt.Print("ASCII: ")
	for _, b := range buffer {
		if b >= 32 && b <= 126 {
			fmt.Printf("%c", b)
		}
	}
	fmt.Println()

	fmt.Printf("As string: %s\n", string(buffer))
}
