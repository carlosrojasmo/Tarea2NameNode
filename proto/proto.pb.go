// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: proto/proto.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Propuesta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk []*PropuestaChunk `protobuf:"bytes,1,rep,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *Propuesta) Reset() {
	*x = Propuesta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Propuesta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Propuesta) ProtoMessage() {}

func (x *Propuesta) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Propuesta.ProtoReflect.Descriptor instead.
func (*Propuesta) Descriptor() ([]byte, []int) {
	return file_proto_proto_proto_rawDescGZIP(), []int{0}
}

func (x *Propuesta) GetChunk() []*PropuestaChunk {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type PropuestaChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset      int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	IpMaquina   string `protobuf:"bytes,2,opt,name=ipMaquina,proto3" json:"ipMaquina,omitempty"`
	NombreLibro string `protobuf:"bytes,3,opt,name=nombreLibro,proto3" json:"nombreLibro,omitempty"`
}

func (x *PropuestaChunk) Reset() {
	*x = PropuestaChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropuestaChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropuestaChunk) ProtoMessage() {}

func (x *PropuestaChunk) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropuestaChunk.ProtoReflect.Descriptor instead.
func (*PropuestaChunk) Descriptor() ([]byte, []int) {
	return file_proto_proto_proto_rawDescGZIP(), []int{1}
}

func (x *PropuestaChunk) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *PropuestaChunk) GetIpMaquina() string {
	if x != nil {
		return x.IpMaquina
	}
	return ""
}

func (x *PropuestaChunk) GetNombreLibro() string {
	if x != nil {
		return x.NombreLibro
	}
	return ""
}

type SendChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk  []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
	Offset int64  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SendChunk) Reset() {
	*x = SendChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendChunk) ProtoMessage() {}

func (x *SendChunk) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendChunk.ProtoReflect.Descriptor instead.
func (*SendChunk) Descriptor() ([]byte, []int) {
	return file_proto_proto_proto_rawDescGZIP(), []int{2}
}

func (x *SendChunk) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

func (x *SendChunk) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *SendChunk) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ReplyEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok int64 `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *ReplyEmpty) Reset() {
	*x = ReplyEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyEmpty) ProtoMessage() {}

func (x *ReplyEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyEmpty.ProtoReflect.Descriptor instead.
func (*ReplyEmpty) Descriptor() ([]byte, []int) {
	return file_proto_proto_proto_rawDescGZIP(), []int{3}
}

func (x *ReplyEmpty) GetOk() int64 {
	if x != nil {
		return x.Ok
	}
	return 0
}

type BookName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *BookName) Reset() {
	*x = BookName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookName) ProtoMessage() {}

func (x *BookName) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookName.ProtoReflect.Descriptor instead.
func (*BookName) Descriptor() ([]byte, []int) {
	return file_proto_proto_proto_rawDescGZIP(), []int{4}
}

func (x *BookName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SendUbicacion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ubicacion string `protobuf:"bytes,1,opt,name=ubicacion,proto3" json:"ubicacion,omitempty"`
	Id        string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SendUbicacion) Reset() {
	*x = SendUbicacion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendUbicacion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendUbicacion) ProtoMessage() {}

func (x *SendUbicacion) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendUbicacion.ProtoReflect.Descriptor instead.
func (*SendUbicacion) Descriptor() ([]byte, []int) {
	return file_proto_proto_proto_rawDescGZIP(), []int{5}
}

func (x *SendUbicacion) GetUbicacion() string {
	if x != nil {
		return x.Ubicacion
	}
	return ""
}

func (x *SendUbicacion) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ChunkId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ChunkId) Reset() {
	*x = ChunkId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkId) ProtoMessage() {}

func (x *ChunkId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkId.ProtoReflect.Descriptor instead.
func (*ChunkId) Descriptor() ([]byte, []int) {
	return file_proto_proto_proto_rawDescGZIP(), []int{6}
}

func (x *ChunkId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_proto_proto_proto protoreflect.FileDescriptor

var file_proto_proto_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x09, 0x50, 0x72,
	0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x12, 0x2b, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x05, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x22, 0x68, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74,
	0x61, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x69, 0x70, 0x4d, 0x61, 0x71, 0x75, 0x69, 0x6e, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x69, 0x70, 0x4d, 0x61, 0x71, 0x75, 0x69, 0x6e, 0x61, 0x12, 0x20, 0x0a, 0x0b,
	0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x22, 0x4d,
	0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1c, 0x0a,
	0x0a, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x6f,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x1e, 0x0a, 0x08, 0x62,
	0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3d, 0x0a, 0x0d, 0x73,
	0x65, 0x6e, 0x64, 0x55, 0x62, 0x69, 0x63, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x75, 0x62, 0x69, 0x63, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x62, 0x69, 0x63, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x19, 0x0a, 0x07, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x9d, 0x02, 0x0a, 0x0c, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x6e,
	0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x28, 0x01, 0x12, 0x3b, 0x0a, 0x10, 0x67,
	0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d, 0x65,
	0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x55, 0x62, 0x69,
	0x63, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x30, 0x01, 0x12, 0x31, 0x0a, 0x0d, 0x64, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x33, 0x0a, 0x0d, 0x73,
	0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x12, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x1a, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61,
	0x12, 0x33, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x61, 0x72, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x79,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_proto_proto_rawDescOnce sync.Once
	file_proto_proto_proto_rawDescData = file_proto_proto_proto_rawDesc
)

func file_proto_proto_proto_rawDescGZIP() []byte {
	file_proto_proto_proto_rawDescOnce.Do(func() {
		file_proto_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_proto_proto_rawDescData)
	})
	return file_proto_proto_proto_rawDescData
}

var file_proto_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_proto_proto_goTypes = []interface{}{
	(*Propuesta)(nil),      // 0: proto.Propuesta
	(*PropuestaChunk)(nil), // 1: proto.PropuestaChunk
	(*SendChunk)(nil),      // 2: proto.sendChunk
	(*ReplyEmpty)(nil),     // 3: proto.replyEmpty
	(*BookName)(nil),       // 4: proto.bookName
	(*SendUbicacion)(nil),  // 5: proto.sendUbicacion
	(*ChunkId)(nil),        // 6: proto.chunkId
}
var file_proto_proto_proto_depIdxs = []int32{
	1, // 0: proto.Propuesta.chunk:type_name -> proto.PropuestaChunk
	2, // 1: proto.LibroService.uploadBook:input_type -> proto.sendChunk
	4, // 2: proto.LibroService.getAddressChunks:input_type -> proto.bookName
	6, // 3: proto.LibroService.downloadChunk:input_type -> proto.chunkId
	0, // 4: proto.LibroService.sendPropuesta:input_type -> proto.Propuesta
	2, // 5: proto.LibroService.OrdenarChunk:input_type -> proto.sendChunk
	3, // 6: proto.LibroService.uploadBook:output_type -> proto.replyEmpty
	5, // 7: proto.LibroService.getAddressChunks:output_type -> proto.sendUbicacion
	2, // 8: proto.LibroService.downloadChunk:output_type -> proto.sendChunk
	0, // 9: proto.LibroService.sendPropuesta:output_type -> proto.Propuesta
	3, // 10: proto.LibroService.OrdenarChunk:output_type -> proto.replyEmpty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_proto_proto_init() }
func file_proto_proto_proto_init() {
	if File_proto_proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Propuesta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_proto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropuestaChunk); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_proto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendChunk); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_proto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyEmpty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_proto_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookName); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_proto_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendUbicacion); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_proto_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_proto_proto_goTypes,
		DependencyIndexes: file_proto_proto_proto_depIdxs,
		MessageInfos:      file_proto_proto_proto_msgTypes,
	}.Build()
	File_proto_proto_proto = out.File
	file_proto_proto_proto_rawDesc = nil
	file_proto_proto_proto_goTypes = nil
	file_proto_proto_proto_depIdxs = nil
}
