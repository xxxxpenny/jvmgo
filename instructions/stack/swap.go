package stack

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
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
