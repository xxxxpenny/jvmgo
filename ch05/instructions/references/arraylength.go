package references

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type ARRAY_LENGTH struct {
	base.NoOperandsInstruction
}

func (self *ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}
	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}
