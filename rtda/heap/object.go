package heap

type Object struct {
	class *Class
	data  interface{}
	extra interface{}
}

func (self *Object) Class() *Class {
	return self.class
}

func (self *Object) Fields() Slots {
	return self.data.(Slots)
}


// class的实例
// class子类的实例


// s -> t
// 传入s, 调用t

func (self *Object) IsInstanceOf(class *Class) bool {
	res := class.IsAssignableFrom(self.class)
	return res
}

func (self *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetRef(field.slotId, ref)
}

func (self *Object) GetRefVar(name, descriptor string) *Object {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetRef(field.slotId)
}

func (self *Object) Extra() interface{} {
	return self.extra
}

func (self *Object) SetExtra(extra interface{}) {
	self.extra = extra
}

func (self *Object) Data() interface{} {
	return self.data
}
