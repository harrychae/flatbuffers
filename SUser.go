// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package AniLand

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SUser struct {
	_tab flatbuffers.Table
}

func GetRootAsSUser(buf []byte, offset flatbuffers.UOffsetT) *SUser {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SUser{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *SUser) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SUser) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SUser) Uidx() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SUser) MutateUidx(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func (rcv *SUser) Uid() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *SUser) X() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *SUser) MutateX(n float64) bool {
	return rcv._tab.MutateFloat64Slot(8, n)
}

func (rcv *SUser) Y() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *SUser) MutateY(n float64) bool {
	return rcv._tab.MutateFloat64Slot(10, n)
}

func (rcv *SUser) Z() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *SUser) MutateZ(n float64) bool {
	return rcv._tab.MutateFloat64Slot(12, n)
}

func (rcv *SUser) AngleZ() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *SUser) MutateAngleZ(n float64) bool {
	return rcv._tab.MutateFloat64Slot(14, n)
}

func SUserStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func SUserAddUidx(builder *flatbuffers.Builder, uidx uint64) {
	builder.PrependUint64Slot(0, uidx, 0)
}
func SUserAddUid(builder *flatbuffers.Builder, uid flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(uid), 0)
}
func SUserAddX(builder *flatbuffers.Builder, x float64) {
	builder.PrependFloat64Slot(2, x, 0.0)
}
func SUserAddY(builder *flatbuffers.Builder, y float64) {
	builder.PrependFloat64Slot(3, y, 0.0)
}
func SUserAddZ(builder *flatbuffers.Builder, z float64) {
	builder.PrependFloat64Slot(4, z, 0.0)
}
func SUserAddAngleZ(builder *flatbuffers.Builder, angleZ float64) {
	builder.PrependFloat64Slot(5, angleZ, 0.0)
}
func SUserEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
