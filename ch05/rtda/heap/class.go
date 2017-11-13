package heap

import (
	"jvmgo/ch05/classfile"
	"strings"
)

type Class struct {
	accessFlags       uint16
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
	initStarted       bool
	jClass            *Object
	sourceFile        string
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfacesNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	class.sourceFile = func() string {
		if sfAttr := cf.SourceFileAttribute(); sfAttr != nil {
			return sfAttr.FileName()
		}

		return "Unknown"
	}()
	return class
}

func (self *Class) SourceFile() string {
	return self.sourceFile
}

func (self *Class) JClass() *Object {
	return self.jClass
}

func (self *Class) JavaName() string {
	return strings.Replace(self.name, "/", ".", -1)
}

func (self *Class) GetClientMethod() *Method {
	return self.getStaticMethod("<clinit>", "()V")
}

func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (self *Class) InitStarted() bool {
	return self.initStarted
}

func (self *Class) StartInit() {
	self.initStarted = true
}

func (self *Class) Name() string {
	return self.name
}

func (self *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		//fmt.Printf("name = %s, descriptor = %s\n", method.name, method.descriptor)
		if method.IsStatic() && method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return nil
}

func (self *Class) StaticVars() Slots {
	return self.staticVars
}

func (self *Class) SuperClass() *Class {
	return self.superClass
}

func (self *Class) isAccessibleTo(class *Class) bool {

	return self.IsPublic() || self.GetPackageName() == class.GetPackageName()
}

func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}

func (self *Class) GetPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

func (self *Class) Loader() *ClassLoader {
	return self.loader
}

func (self *Class) NewObject() *Object {
	object := &Object{}
	object.class = self
	object.data = newSlots(self.instanceSlotCount)
	return object
}

func (self *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := self; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic && field.name == name &&
				field.descriptor == descriptor {
				return field
			}
		}
	}

	return nil
}

func (self *Class) GetInstanceMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		if !method.IsStatic() && method.name == name && method.descriptor == descriptor {
			return method
		}
	}

	return nil
}

func (self *Class) GetRefVar(name, descriptor string) *Object {
	field := self.getField(name, descriptor, true)
	return self.staticVars.GetRef(field.slotId)
}

func (self *Class) isJlObject() bool {
	return self.name == "java/lang/Object"
}

func (self *Class) isJlCloneable() bool {
	return self.name == "java/lang/Cloneable"
}

func (self *Class) isJioSerializable() bool {
	return self.name == "java/io/Serializable"
}

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}

func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}

func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}

func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}

func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}

func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}

func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}

func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

func (self *Class) IsArray() bool {
	return self.name[0] == '['
}

func (self *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[self.name]

	return ok
}
