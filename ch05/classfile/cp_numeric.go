package classfile

import "math"

type ConstantIntegerInfo struct {
	val int32
}

type ConstantFloatInfo struct {
	val float32
}

type ConstantLongInfo struct {
	val int64
}

type ConstantDoubleInfo struct {
	val float64
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	val := int32(reader.readUint32())
	self.val = val
}

func (self *ConstantIntegerInfo) Value() int32 {
	return self.val
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	val := math.Float32frombits(reader.readUint32())
	self.val = val
}

func (self *ConstantFloatInfo) Value() float32 {
	return self.val
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	val := int64(reader.readUint64())
	self.val = val
}

func (self *ConstantLongInfo) Value() int64 {
	return self.val
}

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	val := math.Float64frombits(reader.readUint64())
	self.val = val
}

func (self *ConstantDoubleInfo) Value() float64 {
	return self.val
}
