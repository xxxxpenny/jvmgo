package comparisons

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type LCMP struct {
	base.NoOperandsInstruction
}

func (self *LCMP) Execute(frame *rtda.Frame) {
	v2 := frame.OperandStack().PopLong()
	v1 := frame.OperandStack().PopLong()
	if v1 > v2 {
		frame.OperandStack().PushInt(1)
	} else if v1 < v2 {
		frame.OperandStack().PushInt(-1)
	} else {
		frame.OperandStack().PushInt(0)
	}
}
