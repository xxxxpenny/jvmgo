package heap

type SymRef struct {
	cp *ConstantPool
	// 符号引用的类名
	className string
	class     *Class
}

func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

func (self *SymRef) resolveClassRef() {
	// 取当前类的加载器,加载符号引用的类
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}
