package references

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
	"jvmgo/ch05/rtda/heap"
	"reflect"
)

type ATHROW struct {
	base.NoOperandsInstruction
}

func (self *ATHROW) Execute(frame *rtda.Frame) {
	ex := frame.OperandStack().PopRef()
	if ex == nil {
		panic("java.lang.NullPointerException")
	}

	thread := frame.Thread()
	if !findAndGoToExceptionHandler(thread, ex) {
		handleUnCaughtException(thread, ex)
	}
}

func findAndGoToExceptionHandler(thread *rtda.Thread, ex *heap.Object) bool {
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC() - 1
		handlerPc := frame.Method().FindExceptionHandler(ex.Class(), pc)
		if handlerPc > 0 {
			stack := frame.OperandStack()
			stack.Clear()
			stack.PushRef(ex)
			frame.SetNextPC(handlerPc)
			return true
		}
		thread.PopFrame()
		if thread.IsEmpty() {
			break
		}
	}
	return false
}

func handleUnCaughtException(thread *rtda.Thread, ex *heap.Object) {
	thread.ClearStack()
	jMsg := ex.GetRefVar("detailMessage", "Ljava/lang/String;")
	goMsg := heap.GoString(jMsg)
	println(ex.Class().JavaName() + ": " + goMsg)
	stes := reflect.ValueOf(ex.Extra())
	for i := 0; i < stes.Len(); i++ {
		ste := stes.Index(i).Interface().(interface{
			String() string
		})
		println("\tat" + ste.String())
	}

}
