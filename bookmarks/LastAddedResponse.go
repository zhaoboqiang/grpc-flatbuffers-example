// automatically generated by the FlatBuffers compiler, do not modify

package bookmarks

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type LastAddedResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsLastAddedResponse(buf []byte, offset flatbuffers.UOffsetT) *LastAddedResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &LastAddedResponse{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *LastAddedResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *LastAddedResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *LastAddedResponse) ID() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *LastAddedResponse) URL() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *LastAddedResponse) Title() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func LastAddedResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func LastAddedResponseAddID(builder *flatbuffers.Builder, ID flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(ID), 0)
}
func LastAddedResponseAddURL(builder *flatbuffers.Builder, URL flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(URL), 0)
}
func LastAddedResponseAddTitle(builder *flatbuffers.Builder, title flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(title), 0)
}
func LastAddedResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
