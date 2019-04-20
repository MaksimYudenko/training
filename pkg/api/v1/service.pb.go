// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Main task
type Trainee struct {
	// Unique integer identifier
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Trainee first name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Trainee last name
	Surname string `protobuf:"bytes,3,opt,name=surname,proto3" json:"surname,omitempty"`
	// Trainee e-mail
	Email string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	// Trainee attendance
	HasAttend bool `protobuf:"varint,5,opt,name=hasAttend,proto3" json:"hasAttend,omitempty"`
	// Points number received for the task
	Score                int32    `protobuf:"varint,6,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Trainee) Reset()         { *m = Trainee{} }
func (m *Trainee) String() string { return proto.CompactTextString(m) }
func (*Trainee) ProtoMessage()    {}
func (*Trainee) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *Trainee) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trainee.Unmarshal(m, b)
}
func (m *Trainee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trainee.Marshal(b, m, deterministic)
}
func (m *Trainee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trainee.Merge(m, src)
}
func (m *Trainee) XXX_Size() int {
	return xxx_messageInfo_Trainee.Size(m)
}
func (m *Trainee) XXX_DiscardUnknown() {
	xxx_messageInfo_Trainee.DiscardUnknown(m)
}

var xxx_messageInfo_Trainee proto.InternalMessageInfo

func (m *Trainee) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Trainee) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Trainee) GetSurname() string {
	if m != nil {
		return m.Surname
	}
	return ""
}

func (m *Trainee) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Trainee) GetHasAttend() bool {
	if m != nil {
		return m.HasAttend
	}
	return false
}

func (m *Trainee) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

// Request data to create new trainee
type CreateRequest struct {
	// API versioning: specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Task entity to add
	Trainee              *Trainee `protobuf:"bytes,2,opt,name=trainee,proto3" json:"trainee,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateRequest) GetTrainee() *Trainee {
	if m != nil {
		return m.Trainee
	}
	return nil
}

// Contains data of created trainee
type CreateResponse struct {
	// API versioning: specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// ID of created trainee
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Request data to read trainee
type ReadRequest struct {
	// API versioning: specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Unique integer identifier
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Contains trainee data specified in by ID request
type ReadResponse struct {
	// API versioning: specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Task entity read by ID
	Trainee              *Trainee `protobuf:"bytes,2,opt,name=trainee,proto3" json:"trainee,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadResponse) GetTrainee() *Trainee {
	if m != nil {
		return m.Trainee
	}
	return nil
}

// Request data to update trainee
type UpdateRequest struct {
	// API versioning: specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Task entity to update
	Trainee              *Trainee `protobuf:"bytes,2,opt,name=trainee,proto3" json:"trainee,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *UpdateRequest) GetTrainee() *Trainee {
	if m != nil {
		return m.Trainee
	}
	return nil
}

// Contains status of update operation
type UpdateResponse struct {
	// API versioning: specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Contains number of entities have been updated
	// Equals 1 in case of successful update
	Updated              int64    `protobuf:"varint,2,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *UpdateResponse) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

// Request data to delete trainee
type DeleteRequest struct {
	// API versioning: specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Unique integer identifier to delete
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Contains status of delete operation
type DeleteResponse struct {
	// API versioning: specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Contains number of entities have been deleted
	// Equals 1 in case of successful delete
	Deleted              int64    `protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{8}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteResponse) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

// Request data to read all trainees
type ReadAllRequest struct {
	// API versioning: specify version explicitly
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAllRequest) Reset()         { *m = ReadAllRequest{} }
func (m *ReadAllRequest) String() string { return proto.CompactTextString(m) }
func (*ReadAllRequest) ProtoMessage()    {}
func (*ReadAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{9}
}

func (m *ReadAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAllRequest.Unmarshal(m, b)
}
func (m *ReadAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAllRequest.Marshal(b, m, deterministic)
}
func (m *ReadAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAllRequest.Merge(m, src)
}
func (m *ReadAllRequest) XXX_Size() int {
	return xxx_messageInfo_ReadAllRequest.Size(m)
}
func (m *ReadAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAllRequest proto.InternalMessageInfo

func (m *ReadAllRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

// Contains list of all trainees
type ReadAllResponse struct {
	// API versioning: specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// List of all trainees
	Trainees             []*Trainee `protobuf:"bytes,2,rep,name=trainees,proto3" json:"trainees,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReadAllResponse) Reset()         { *m = ReadAllResponse{} }
func (m *ReadAllResponse) String() string { return proto.CompactTextString(m) }
func (*ReadAllResponse) ProtoMessage()    {}
func (*ReadAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{10}
}

func (m *ReadAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAllResponse.Unmarshal(m, b)
}
func (m *ReadAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAllResponse.Marshal(b, m, deterministic)
}
func (m *ReadAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAllResponse.Merge(m, src)
}
func (m *ReadAllResponse) XXX_Size() int {
	return xxx_messageInfo_ReadAllResponse.Size(m)
}
func (m *ReadAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAllResponse proto.InternalMessageInfo

func (m *ReadAllResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadAllResponse) GetTrainees() []*Trainee {
	if m != nil {
		return m.Trainees
	}
	return nil
}

func init() {
	proto.RegisterType((*Trainee)(nil), "v1.Trainee")
	proto.RegisterType((*CreateRequest)(nil), "v1.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "v1.CreateResponse")
	proto.RegisterType((*ReadRequest)(nil), "v1.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "v1.ReadResponse")
	proto.RegisterType((*UpdateRequest)(nil), "v1.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "v1.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "v1.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "v1.DeleteResponse")
	proto.RegisterType((*ReadAllRequest)(nil), "v1.ReadAllRequest")
	proto.RegisterType((*ReadAllResponse)(nil), "v1.ReadAllResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6a, 0xdb, 0x40,
	0x10, 0x46, 0xb2, 0x2d, 0xc5, 0xe3, 0x58, 0x89, 0x27, 0x81, 0x2a, 0xa2, 0x14, 0xb1, 0x50, 0x2a,
	0x7c, 0xb0, 0xb0, 0x7b, 0x0b, 0xa1, 0x10, 0x6a, 0xa8, 0x29, 0x3d, 0xa9, 0xed, 0xa5, 0x97, 0xb2,
	0x8d, 0x16, 0x67, 0x41, 0x91, 0x54, 0x69, 0x6d, 0x28, 0x25, 0x97, 0xbe, 0x41, 0xe9, 0x73, 0xf4,
	0x69, 0xfa, 0x0a, 0x7d, 0x90, 0xa2, 0x5d, 0xad, 0x6d, 0x99, 0xc8, 0xf4, 0x90, 0x9b, 0x66, 0x76,
	0xe6, 0xfb, 0x99, 0x19, 0x1b, 0x86, 0x25, 0x2b, 0xd6, 0xfc, 0x86, 0x4d, 0xf2, 0x22, 0x13, 0x19,
	0x9a, 0xeb, 0xa9, 0x17, 0x88, 0x5b, 0x5e, 0xc4, 0x9f, 0x73, 0x5a, 0x88, 0x6f, 0xe1, 0x32, 0xcb,
	0x96, 0x09, 0x0b, 0x69, 0xce, 0x43, 0x9a, 0xa6, 0x99, 0xa0, 0x82, 0x67, 0x69, 0xa9, 0xaa, 0xc9,
	0x4f, 0x03, 0xec, 0x0f, 0x05, 0xe5, 0x29, 0x63, 0xe8, 0x80, 0xc9, 0x63, 0xd7, 0xf0, 0x8d, 0xa0,
	0x13, 0x99, 0x3c, 0x46, 0x84, 0x6e, 0x4a, 0xef, 0x98, 0x6b, 0xfa, 0x46, 0xd0, 0x8f, 0xe4, 0x37,
	0xba, 0x60, 0x97, 0xab, 0x42, 0xa6, 0x3b, 0x32, 0xad, 0x43, 0x3c, 0x87, 0x1e, 0xbb, 0xa3, 0x3c,
	0x71, 0xbb, 0x32, 0xaf, 0x02, 0x7c, 0x0a, 0xfd, 0x5b, 0x5a, 0x5e, 0x0b, 0xc1, 0xd2, 0xd8, 0xed,
	0xf9, 0x46, 0x70, 0x14, 0x6d, 0x13, 0x55, 0x4f, 0x79, 0x93, 0x15, 0xcc, 0xb5, 0x7c, 0x23, 0xe8,
	0x45, 0x2a, 0x20, 0x0b, 0x18, 0xbe, 0x2e, 0x18, 0x15, 0x2c, 0x62, 0x5f, 0x57, 0xac, 0x14, 0x78,
	0x0a, 0x1d, 0x9a, 0x73, 0xa9, 0xac, 0x1f, 0x55, 0x9f, 0xf8, 0x1c, 0x6c, 0xa1, 0x54, 0x4b, 0x75,
	0x83, 0xd9, 0x60, 0xb2, 0x9e, 0x4e, 0x6a, 0x23, 0x91, 0x7e, 0x23, 0x33, 0x70, 0x34, 0x52, 0x99,
	0x67, 0x69, 0xc9, 0x1e, 0x80, 0x52, 0xae, 0x4d, 0xed, 0x9a, 0x84, 0x30, 0x88, 0x18, 0x8d, 0xdb,
	0xb9, 0xf7, 0x1b, 0xde, 0xc0, 0xb1, 0x6a, 0x68, 0xa5, 0xf8, 0x4f, 0xb5, 0x0b, 0x18, 0x7e, 0xcc,
	0xe3, 0xc7, 0xf0, 0x7d, 0x05, 0x8e, 0x46, 0x6a, 0x15, 0xe5, 0x82, 0xbd, 0x92, 0x35, 0xda, 0x8b,
	0x0e, 0xc9, 0x14, 0x86, 0x73, 0x96, 0xb0, 0x43, 0x3a, 0xf6, 0x67, 0x70, 0x05, 0x8e, 0x6e, 0x39,
	0x44, 0x18, 0xcb, 0x9a, 0x0d, 0x61, 0x1d, 0x12, 0x02, 0x4e, 0x35, 0xc1, 0xeb, 0x24, 0x69, 0x65,
	0x24, 0xef, 0xe0, 0x64, 0x53, 0xd3, 0x4a, 0xf1, 0x02, 0x8e, 0xea, 0x11, 0x94, 0xae, 0xe9, 0x77,
	0xf6, 0xe7, 0xb3, 0x79, 0x9c, 0xfd, 0xee, 0x80, 0x53, 0x67, 0xdf, 0xab, 0x5f, 0x0f, 0x2e, 0xc0,
	0x52, 0xb7, 0x82, 0xa3, 0xaa, 0xa7, 0x71, 0x81, 0x1e, 0xee, 0xa6, 0x14, 0x3d, 0x79, 0xf2, 0xe3,
	0xcf, 0xdf, 0x5f, 0xe6, 0x88, 0x1c, 0x87, 0xeb, 0x69, 0x28, 0x91, 0x79, 0xba, 0xbc, 0x34, 0xc6,
	0x38, 0x87, 0x6e, 0x25, 0x15, 0x4f, 0xaa, 0xa6, 0x9d, 0x5b, 0xf2, 0x4e, 0xb7, 0x89, 0x1a, 0xe3,
	0x42, 0x62, 0x9c, 0xe1, 0x68, 0x17, 0x23, 0xfc, 0xce, 0xe3, 0x7b, 0x5c, 0x81, 0xa5, 0x76, 0xa8,
	0xf4, 0x34, 0x2e, 0x43, 0xe9, 0x69, 0xae, 0x98, 0xcc, 0x25, 0xd6, 0x2b, 0xef, 0xa2, 0x89, 0x55,
	0x7b, 0x9e, 0xf0, 0xf8, 0xfe, 0xd2, 0x18, 0x7f, 0x7a, 0x36, 0x3b, 0xf8, 0x8e, 0x6f, 0xc1, 0x52,
	0x9b, 0x54, 0xb4, 0x8d, 0x43, 0x50, 0xb4, 0xcd, 0x45, 0x6b, 0x0b, 0xe3, 0x07, 0x2c, 0x2c, 0xc0,
	0xae, 0x77, 0x86, 0xa8, 0xad, 0x6f, 0x97, 0xec, 0x9d, 0x35, 0x72, 0x35, 0xdc, 0xb9, 0x84, 0x73,
	0xb0, 0x31, 0xd5, 0x2f, 0x96, 0xfc, 0xb7, 0x7a, 0xf9, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x25, 0xee,
	0x9b, 0x2a, 0xec, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TraineeServiceClient is the client API for TraineeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TraineeServiceClient interface {
	// Create new trainee
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Read trainee
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	// Update trainee
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete trainee
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// Read all trainees
	ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error)
}

type traineeServiceClient struct {
	cc *grpc.ClientConn
}

func NewTraineeServiceClient(cc *grpc.ClientConn) TraineeServiceClient {
	return &traineeServiceClient{cc}
}

func (c *traineeServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/v1.TraineeService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traineeServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/v1.TraineeService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traineeServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/v1.TraineeService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traineeServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.TraineeService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traineeServiceClient) ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error) {
	out := new(ReadAllResponse)
	err := c.cc.Invoke(ctx, "/v1.TraineeService/ReadAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TraineeServiceServer is the server API for TraineeService service.
type TraineeServiceServer interface {
	// Create new trainee
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Read trainee
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	// Update trainee
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Delete trainee
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// Read all trainees
	ReadAll(context.Context, *ReadAllRequest) (*ReadAllResponse, error)
}

// UnimplementedTraineeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTraineeServiceServer struct {
}

func (*UnimplementedTraineeServiceServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedTraineeServiceServer) Read(ctx context.Context, req *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedTraineeServiceServer) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedTraineeServiceServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedTraineeServiceServer) ReadAll(ctx context.Context, req *ReadAllRequest) (*ReadAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAll not implemented")
}

func RegisterTraineeServiceServer(s *grpc.Server, srv TraineeServiceServer) {
	s.RegisterService(&_TraineeService_serviceDesc, srv)
}

func _TraineeService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraineeServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.TraineeService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraineeServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraineeService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraineeServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.TraineeService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraineeServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraineeService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraineeServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.TraineeService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraineeServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraineeService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraineeServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.TraineeService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraineeServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraineeService_ReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraineeServiceServer).ReadAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.TraineeService/ReadAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraineeServiceServer).ReadAll(ctx, req.(*ReadAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TraineeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.TraineeService",
	HandlerType: (*TraineeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TraineeService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _TraineeService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TraineeService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TraineeService_Delete_Handler,
		},
		{
			MethodName: "ReadAll",
			Handler:    _TraineeService_ReadAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
