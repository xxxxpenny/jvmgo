package rtda

import "jvmgo/ch05/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
