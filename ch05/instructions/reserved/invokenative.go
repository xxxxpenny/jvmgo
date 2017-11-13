package reserved

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
	"jvmgo/ch05/native"

	_ "jvmgo/ch05/native/java/lang"
	_ "jvmgo/ch05/native/java/sun/misc"
)

type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()
	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + "." + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}
	nativeMethod(frame)
}
