package rtda

import "jvmgo/ch05/rtda/heap"

type Frame struct {
	lower        *Frame
	// 局部变量表
	localVars    LocalVars
	// 操作数栈
	operandStack *OperandStack
	thread       *Thread
	// 动态链接
	method       *heap.Method
	// 返回地址
	nextPC       int
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

func (self *Frame) RevertNextPC() {
	self.nextPC = self.thread.pc
}

func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) NextPC() int {
	return self.nextPC
}

func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}

func (self *Frame) Method() *heap.Method {
	return self.method
}
