package heap

func LookupMethodInClass(class *Class, name, descriptor string) *Method {
	//for c := class; c != nil; c = c.superClass {
	//	for _, method := range c.methods {
	//		if method.name == name && method.descriptor == descriptor {
	//			return method
	//		}
	//	}
	//}
	//return nil

	if class == nil {
		return nil
	} else {
		for _, method := range class.methods {
			// 不需要判断方法是否可以接近,编译期错误
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
		return LookupMethodInClass(class.superClass, name, descriptor)
	}

}

func LookupMethodInInterfaces(ifaces []*Class, name, descriptor string) *Method {
	for _, iface := range ifaces {
		for _, method := range iface.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
		method := LookupMethodInInterfaces(iface.interfaces, name, descriptor)
		if method == nil {
			return nil
		}
	}
	return nil
}
