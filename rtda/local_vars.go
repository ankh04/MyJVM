package rtda

import "math"

// LocalVars 使用Slot作为局部变量表的元素是因为
//   1. 方便对数值类型和引用类型数据的统一管理
//   2. 方便垃圾回收
type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (l LocalVars) SetInt(index uint, val int32) {
	l[index].num = val
}

func (l LocalVars) GetInt(index uint) int32 {
	return l[index].num
}

// float转成int bits存储
func (l LocalVars) SetFloat(index uint, float float32) {
	bits := math.Float32bits(float)
	l[index].num = int32(bits)
}

func (l LocalVars) GetFloat(index uint) float32 {
	bits := uint32(l[index].num)
	return math.Float32frombits(bits)
}

// Long转成int bits存储
func (l LocalVars) SetLong(index uint, long int64) {
	l[index].num = int32(long)
	l[index+1].num = int32(long >> 32)
}

func (l LocalVars) GetLong(index uint) int64 {
	low := uint32(l[index].num)
	high := uint32(l[index+1].num)
	return int64(high)<<32 | int64(low)
}

// Double先转成long型，再存储
func (l LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	l.SetLong(index, int64(bits))
}

func (l LocalVars) GetDouble(index uint) float64 {
	bits := uint64(l.GetLong(index))
	return math.Float64frombits(bits)
}

func (l LocalVars) SetRef(index uint, ref *Object) {
	l[index].ref = ref
}

func (l LocalVars) GetRef(index uint) *Object {
	return l[index].ref
}
