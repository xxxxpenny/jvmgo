package classfile

type LineNumberTableAttribute struct {
	LineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	self.LineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range self.LineNumberTable {
		self.LineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}

	}
}

func (self *LineNumberTableAttribute) GetLineNumber(pc int) int {
	for i := len(self.LineNumberTable) - 1; i >= 0; i-- {
		entry := self.LineNumberTable[i]
		if pc >= int(entry.startPc) {
			return int(entry.lineNumber)
		}
	}
	return -1
}