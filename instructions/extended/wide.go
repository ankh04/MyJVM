package extended

import (
	"MyJVM/instructions/base"
	"MyJVM/instructions/loads"
	"MyJVM/instructions/math"
	"MyJVM/instructions/stores"
	"MyJVM/rtda"
)

// wide指令用于将其他指令扩展成支持16位局部变量表索引的指令

type WIDE struct {
	modifiedInstruction base.Instruction
}

func (wide *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15: // iload
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x16: // lload
		inst := &loads.LLOAD{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x17: // fload
		inst := &loads.FLOAD{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x18: // dload
		inst := &loads.DLOAD{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x19: // aload
		inst := &loads.ALOAD{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x36: // istore
		inst := &stores.ISTORE{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x37: // lstore
		inst := &stores.LSTORE{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x38: // fstore
		inst := &stores.FSTORE{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x39: // dstore
		inst := &stores.DSTORE{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x3a: // astore
		inst := &stores.ASTORE{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x84: // iinc
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadUint16())
	case 0xa9: // ret
		panic("ret current unsupported")
	}
}

func (wide *WIDE) Execute(frame *rtda.Frame) {
	// 执行扩展的命令
	wide.modifiedInstruction.Execute(frame)
}
