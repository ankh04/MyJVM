package base

type BytecodeReader struct {
	code []byte
	pc   int
}

func (br *BytecodeReader) PC() int {
	return br.pc
}

func (br *BytecodeReader) Reset(code []byte, pc int) {
	br.code = code
	br.pc = pc
}

func (br *BytecodeReader) ReadUint8() uint8 {
	res := br.code[br.pc] // 读取八位数据
	br.pc++
	return res
}

func (br *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(br.code[br.pc])   // 读取八位数据
	byte2 := uint16(br.code[br.pc+1]) // 继续读取八位数据
	br.pc += 2
	return (byte1 << 8) | byte2
}

func (br *BytecodeReader) ReadInt8() int8 {
	return int8(br.ReadUint8())
}

func (br *BytecodeReader) ReadInt16() int16 {
	return int16(br.ReadUint16())
}

func (br *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(br.ReadUint8())
	byte2 := int32(br.ReadUint8())
	byte3 := int32(br.ReadUint8())
	byte4 := int32(br.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

func (br *BytecodeReader) SkipPadding() {
	// 往后跳4个字节
	for br.pc%4 != 0 {
		br.ReadUint8()
	}
}

func (br *BytecodeReader) ReadInt32s(count int32) []int32 {
	ints := make([]int32, count)
	for i := range ints {
		ints[i] = br.ReadInt32()
	}
	return ints
}
