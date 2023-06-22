// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package AniLand

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Proto struct {
	_tab flatbuffers.Table
}

func GetRootAsProto(buf []byte, offset flatbuffers.UOffsetT) *Proto {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Proto{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Proto) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Proto) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Proto) Ft() Ftype {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return Ftype(rcv._tab.GetInt16(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Proto) MutateFt(n Ftype) bool {
	return rcv._tab.MutateInt16Slot(4, int16(n))
}

func (rcv *Proto) Uidx() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Proto) MutateUidx(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

func (rcv *Proto) Pos(obj *Vec3) *Vec3 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Vec3)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Proto) User(obj *Uinfo) *Uinfo {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Uinfo)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func ProtoStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func ProtoAddFt(builder *flatbuffers.Builder, Ft Ftype) {
	builder.PrependInt16Slot(0, int16(Ft), 0)
}
func ProtoAddUidx(builder *flatbuffers.Builder, uidx uint64) {
	builder.PrependUint64Slot(1, uidx, 0)
}
func ProtoAddPos(builder *flatbuffers.Builder, pos flatbuffers.UOffsetT) {
	builder.PrependStructSlot(2, flatbuffers.UOffsetT(pos), 0)
}
func ProtoAddUser(builder *flatbuffers.Builder, user flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(user), 0)
}
func ProtoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
