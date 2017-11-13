package heap

// jvms8 6.5.instanceof
// jvms8 6.5.checkcast

//other -> self
//s -> t
// other -> self

//
func (self *Class) IsAssignableFrom(other *Class) bool {
	s, t := other, self

	if s == t {

		return true
	}

	if !s.IsArray() {
		if !s.IsInterface() {
			if !t.IsInterface() {
				return s.IsSubClassOf(t)
			} else {
				return s.IsImplements(t)
			}
		} else {
			if !t.IsInterface() {
				return t.isJlObject()
			} else {
				return s.IsSubInterfaceOf(t)
			}
		}
	} else {
		if !t.IsArray() {
			if !t.IsInterface() {
				return t.isJlObject()
			} else {
				return t.isJlCloneable() || t.isJioSerializable()
			}
		} else {
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.IsAssignableFrom(sc)
		}
	}

	return false
}

// self extends c
func (self *Class) IsSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

// self implements iface
func (self *Class) IsImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.IsSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

// self extends iface
func (self *Class) IsSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == iface || superInterface.IsSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

// c extends self
func (self *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(self)
}
