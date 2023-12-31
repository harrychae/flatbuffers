// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package AniLand

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Item struct {
	_tab flatbuffers.Struct
}

func (rcv *Item) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Item) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *Item) Iidx() uint32 {
	return rcv._tab.GetUint32(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *Item) MutateIidx(n uint32) bool {
	return rcv._tab.MutateUint32(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func (rcv *Item) Itype() byte {
	return rcv._tab.GetByte(rcv._tab.Pos + flatbuffers.UOffsetT(4))
}
func (rcv *Item) MutateItype(n byte) bool {
	return rcv._tab.MutateByte(rcv._tab.Pos+flatbuffers.UOffsetT(4), n)
}

func (rcv *Item) Itype2() byte {
	return rcv._tab.GetByte(rcv._tab.Pos + flatbuffers.UOffsetT(5))
}
func (rcv *Item) MutateItype2(n byte) bool {
	return rcv._tab.MutateByte(rcv._tab.Pos+flatbuffers.UOffsetT(5), n)
}

func CreateItem(builder *flatbuffers.Builder, iidx uint32, itype byte, itype2 byte) flatbuffers.UOffsetT {
	builder.Prep(4, 8)
	builder.Pad(2)
	builder.PrependByte(itype2)
	builder.PrependByte(itype)
	builder.PrependUint32(iidx)
	return builder.Offset()
}
