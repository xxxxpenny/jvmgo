package classfile

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributeInfo := make([]AttributeInfo, attributesCount)
	for i := range attributeInfo {
		attributeInfo[i] = readAttribute(reader, cp)
	}
	return attributeInfo
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attributeNameIndex := reader.readUint16()
	attributeName := cp.getUtf8(attributeNameIndex)
	attributeLength := reader.readUint32()
	attributeInfo := newAttribute(attributeName, attributeLength, cp)
	attributeInfo.readInfo(reader)
	return attributeInfo
}

func newAttribute(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
