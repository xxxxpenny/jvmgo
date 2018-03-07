package base

type BytecodeReader struct {
	code []byte
	pc   int
}

func (self *BytecodeReader) PC() int {
	return self.pc
}

func (self *BytecodeReader) Reset(code []byte, pc int) {
	self.code = code
	self.pc = pc
}

func (self *BytecodeReader) ReadUint8() uint8 {
	i := self.code[self.pc]
	self.pc++
	return i
}

func (self *BytecodeReader) ReadInt8() int8 {
	return int8(self.ReadUint8())
}

func (self *BytecodeReader) ReadUint16() uint16 {
	byte1 := self.ReadUint8()
	byte2 := self.ReadUint8()
	return uint16(byte1)<<8 | uint16(byte2)
}

func (self *BytecodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}

func (self *BytecodeReader) ReadUint32() uint32 {
	byte1 := self.ReadUint16()
	byte2 := self.ReadUint16()
	return uint32(byte1)<<16 | uint32(byte2)
}

func (self *BytecodeReader) ReadInt32() int32 {
	return int32(self.ReadUint32())
}

func (self *BytecodeReader) SkipPadding() {
	for self.pc%4 != 0 {
		self.ReadUint8()
	}
}

func (self *BytecodeReader) ReadInt32s(count int32) []int32 {
	int32Array := make([]int32, count)
	for i := range int32Array {
		int32Array[i] = self.ReadInt32()
	}
	return int32Array
}
