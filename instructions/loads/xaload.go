package loads

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

func checkNotNIl(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}

func checkIndex(len int, index int32) {
	if index < 0 || index >= int32(len) {
		panic("ArrayIndexOutOfBoundsException")
	}
}

type AALOAD struct {
	base.NoOperandsInstruction
}

func (self *AALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNIl(arrRef)
	refs := arrRef.Refs()
	checkIndex(len(refs), index)
	stack.PushRef(refs[index])
}

type BALOAD struct {
	base.NoOperandsInstruction
}

func (self *BALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNIl(arrRef)
	refs := arrRef.Bytes()
	checkIndex(len(refs), index)
	stack.PushInt(int32(refs[index]))
}

type CALOAD struct {
	base.NoOperandsInstruction
}

func (self *CALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNIl(arrRef)
	refs := arrRef.Chars()
	checkIndex(len(refs), index)
	stack.PushInt(int32(refs[index]))
}

type DALOAD struct {
	base.NoOperandsInstruction
}

func (self *DALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNIl(arrRef)
	refs := arrRef.Doubles()
	checkIndex(len(refs), index)
	stack.PushDouble(refs[index])
}

type FALOAD struct {
	base.NoOperandsInstruction
}

func (self *FALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNIl(arrRef)
	refs := arrRef.Floats()
	checkIndex(len(refs), index)
	stack.PushFloat(refs[index])
}

type IALOAD struct {
	base.NoOperandsInstruction
}

func (self *IALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNIl(arrRef)
	refs := arrRef.Ints()
	checkIndex(len(refs), index)
	stack.PushInt(refs[index])
}

type LALOAD struct {
	base.NoOperandsInstruction
}

func (self *LALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNIl(arrRef)
	refs := arrRef.Longs()
	checkIndex(len(refs), index)
	stack.PushLong(refs[index])
}

type SALOAD struct {
	base.NoOperandsInstruction
}

func (self *SALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNIl(arrRef)
	refs := arrRef.Shorts()
	checkIndex(len(refs), index)
	stack.PushInt(int32(refs[index]))
}
