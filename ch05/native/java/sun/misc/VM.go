package misc

import (
	"jvmgo/ch05/native"
	"jvmgo/ch05/rtda"
	"jvmgo/ch05/rtda/heap"
	"jvmgo/ch05/instructions/base"
)

const jlVM string = "sun/misc/VM"

func init() {
	native.Register(jlVM, "initialize", "()V", initialize)
}

func initialize(frame *rtda.Frame) {
	vmClass := frame.Method().Class()
	savedProps := vmClass.GetRefVar("savedProps", "Ljava/util/Properties;")
	key := heap.JString(vmClass.Loader(), "foo")
	value := heap.JString(vmClass.Loader(), "bar")
	frame.OperandStack().PushRef(savedProps)
	frame.OperandStack().PushRef(key)
	frame.OperandStack().PushRef(value)
	propsClass := frame.Method().Class().Loader().LoadClass("java/util/Properties")
	setPropMethod := propsClass.GetInstanceMethod("setProperty",
		"(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	base.InvokeMethod(frame, setPropMethod)
}
