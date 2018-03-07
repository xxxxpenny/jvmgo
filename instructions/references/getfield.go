package references

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

type GET_FIELD struct {
	base.Index16Instruction
}

func (self *GET_FIELD) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	slotId := field.SlotId()
	descriptor := field.Descriptor()
	if field.IsStatic() {
		panic("")
	}
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		panic("java.lang.NullPointException")
	}
	stack := frame.OperandStack()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'I', 'S':
		stack.PushInt(ref.Fields().GetInt(slotId))
	case 'J':
		stack.PushLong(ref.Fields().GetLong(slotId))
	case 'F':
		stack.PushFloat(ref.Fields().GetFloat(slotId))
	case 'D':
		stack.PushDouble(ref.Fields().GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(ref.Fields().GetRef(slotId))

	}
}
