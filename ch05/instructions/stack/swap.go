package stack

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type SWAP struct {
	base.NoOperandsInstruction
}

func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	first := stack.PopSlot()
	second := stack.PopSlot()
	stack.PushSlot(first)
	stack.PushSlot(second)
}
