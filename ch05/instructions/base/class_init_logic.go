package base

import (
	"jvmgo/ch05/rtda"
	"jvmgo/ch05/rtda/heap"
)

func InitClass(thread *rtda.Thread, class *heap.Class) {
	class.StartInit()
	scheduleClient(thread, class)
	initSuperClass(thread, class)

}

func scheduleClient(thread *rtda.Thread, class *heap.Class) {
	clinit := class.GetClientMethod()
	if clinit != nil {
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}

func initSuperClass(thread *rtda.Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.SuperClass()
		if superClass != nil && !superClass.InitStarted() {
			InitClass(thread, superClass)
		}
	}
}
