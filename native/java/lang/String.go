package lang

import (
	"jvmgo/native"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

const jlString string = "java/lang/String"

func init() {
	native.Register(jlString, "intern", "()Ljava/lang/String;", intern)
}

func intern(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}
