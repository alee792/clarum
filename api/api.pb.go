// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package clarum

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Action_Type int32

const (
	Action_CREATION    Action_Type = 0
	Action_TRANSPORT   Action_Type = 1
	Action_USE         Action_Type = 2
	Action_MAINTENANCE Action_Type = 3
	Action_ADMIN       Action_Type = 4
	Action_DISPOSAL    Action_Type = 5
	Action_OTHER       Action_Type = 6
)

var Action_Type_name = map[int32]string{
	0: "CREATION",
	1: "TRANSPORT",
	2: "USE",
	3: "MAINTENANCE",
	4: "ADMIN",
	5: "DISPOSAL",
	6: "OTHER",
}

var Action_Type_value = map[string]int32{
	"CREATION":    0,
	"TRANSPORT":   1,
	"USE":         2,
	"MAINTENANCE": 3,
	"ADMIN":       4,
	"DISPOSAL":    5,
	"OTHER":       6,
}

func (x Action_Type) String() string {
	return proto.EnumName(Action_Type_name, int32(x))
}

func (Action_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2, 0}
}

type Instrument struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	IssuedAt             int32    `protobuf:"varint,3,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
	Lineage              *Lineage `protobuf:"bytes,4,opt,name=lineage,proto3" json:"lineage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Instrument) Reset()         { *m = Instrument{} }
func (m *Instrument) String() string { return proto.CompactTextString(m) }
func (*Instrument) ProtoMessage()    {}
func (*Instrument) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *Instrument) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instrument.Unmarshal(m, b)
}
func (m *Instrument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instrument.Marshal(b, m, deterministic)
}
func (m *Instrument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instrument.Merge(m, src)
}
func (m *Instrument) XXX_Size() int {
	return xxx_messageInfo_Instrument.Size(m)
}
func (m *Instrument) XXX_DiscardUnknown() {
	xxx_messageInfo_Instrument.DiscardUnknown(m)
}

var xxx_messageInfo_Instrument proto.InternalMessageInfo

func (m *Instrument) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Instrument) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Instrument) GetIssuedAt() int32 {
	if m != nil {
		return m.IssuedAt
	}
	return 0
}

func (m *Instrument) GetLineage() *Lineage {
	if m != nil {
		return m.Lineage
	}
	return nil
}

type Lineage struct {
	Id                   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Actions              []*Action `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Lineage) Reset()         { *m = Lineage{} }
func (m *Lineage) String() string { return proto.CompactTextString(m) }
func (*Lineage) ProtoMessage()    {}
func (*Lineage) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *Lineage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Lineage.Unmarshal(m, b)
}
func (m *Lineage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Lineage.Marshal(b, m, deterministic)
}
func (m *Lineage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lineage.Merge(m, src)
}
func (m *Lineage) XXX_Size() int {
	return xxx_messageInfo_Lineage.Size(m)
}
func (m *Lineage) XXX_DiscardUnknown() {
	xxx_messageInfo_Lineage.DiscardUnknown(m)
}

var xxx_messageInfo_Lineage proto.InternalMessageInfo

func (m *Lineage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Lineage) GetActions() []*Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

type Action struct {
	Id                   string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	At                   int32       `protobuf:"varint,2,opt,name=at,proto3" json:"at,omitempty"`
	Type                 Action_Type `protobuf:"varint,3,opt,name=type,proto3,enum=Action_Type" json:"type,omitempty"`
	Event                *Event      `protobuf:"bytes,4,opt,name=event,proto3" json:"event,omitempty"`
	Description          string      `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Action) Reset()         { *m = Action{} }
func (m *Action) String() string { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()    {}
func (*Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *Action) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Action.Unmarshal(m, b)
}
func (m *Action) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Action.Marshal(b, m, deterministic)
}
func (m *Action) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Action.Merge(m, src)
}
func (m *Action) XXX_Size() int {
	return xxx_messageInfo_Action.Size(m)
}
func (m *Action) XXX_DiscardUnknown() {
	xxx_messageInfo_Action.DiscardUnknown(m)
}

var xxx_messageInfo_Action proto.InternalMessageInfo

func (m *Action) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Action) GetAt() int32 {
	if m != nil {
		return m.At
	}
	return 0
}

func (m *Action) GetType() Action_Type {
	if m != nil {
		return m.Type
	}
	return Action_CREATION
}

func (m *Action) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Action) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Person struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	FirstName            string   `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Person) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Person) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

type Role struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Role                 string   `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (m *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(m, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Role) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *Role) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Party struct {
	Person               *Person  `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
	Role                 *Role    `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Party) Reset()         { *m = Party{} }
func (m *Party) String() string { return proto.CompactTextString(m) }
func (*Party) ProtoMessage()    {}
func (*Party) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *Party) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Party.Unmarshal(m, b)
}
func (m *Party) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Party.Marshal(b, m, deterministic)
}
func (m *Party) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Party.Merge(m, src)
}
func (m *Party) XXX_Size() int {
	return xxx_messageInfo_Party.Size(m)
}
func (m *Party) XXX_DiscardUnknown() {
	xxx_messageInfo_Party.DiscardUnknown(m)
}

var xxx_messageInfo_Party proto.InternalMessageInfo

func (m *Party) GetPerson() *Person {
	if m != nil {
		return m.Person
	}
	return nil
}

func (m *Party) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

type Event struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Parties              []*Party `protobuf:"bytes,3,rep,name=parties,proto3" json:"parties,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Event) GetParties() []*Party {
	if m != nil {
		return m.Parties
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{7}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type CreateInstrumentRequest struct {
	Instrument           *Instrument `protobuf:"bytes,1,opt,name=instrument,proto3" json:"instrument,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateInstrumentRequest) Reset()         { *m = CreateInstrumentRequest{} }
func (m *CreateInstrumentRequest) String() string { return proto.CompactTextString(m) }
func (*CreateInstrumentRequest) ProtoMessage()    {}
func (*CreateInstrumentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{8}
}

func (m *CreateInstrumentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateInstrumentRequest.Unmarshal(m, b)
}
func (m *CreateInstrumentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateInstrumentRequest.Marshal(b, m, deterministic)
}
func (m *CreateInstrumentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateInstrumentRequest.Merge(m, src)
}
func (m *CreateInstrumentRequest) XXX_Size() int {
	return xxx_messageInfo_CreateInstrumentRequest.Size(m)
}
func (m *CreateInstrumentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateInstrumentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateInstrumentRequest proto.InternalMessageInfo

func (m *CreateInstrumentRequest) GetInstrument() *Instrument {
	if m != nil {
		return m.Instrument
	}
	return nil
}

type GetInstrumentRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInstrumentRequest) Reset()         { *m = GetInstrumentRequest{} }
func (m *GetInstrumentRequest) String() string { return proto.CompactTextString(m) }
func (*GetInstrumentRequest) ProtoMessage()    {}
func (*GetInstrumentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{9}
}

func (m *GetInstrumentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInstrumentRequest.Unmarshal(m, b)
}
func (m *GetInstrumentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInstrumentRequest.Marshal(b, m, deterministic)
}
func (m *GetInstrumentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInstrumentRequest.Merge(m, src)
}
func (m *GetInstrumentRequest) XXX_Size() int {
	return xxx_messageInfo_GetInstrumentRequest.Size(m)
}
func (m *GetInstrumentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInstrumentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInstrumentRequest proto.InternalMessageInfo

func (m *GetInstrumentRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListInstrumentsRequest struct {
	Ids                  []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListInstrumentsRequest) Reset()         { *m = ListInstrumentsRequest{} }
func (m *ListInstrumentsRequest) String() string { return proto.CompactTextString(m) }
func (*ListInstrumentsRequest) ProtoMessage()    {}
func (*ListInstrumentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{10}
}

func (m *ListInstrumentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListInstrumentsRequest.Unmarshal(m, b)
}
func (m *ListInstrumentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListInstrumentsRequest.Marshal(b, m, deterministic)
}
func (m *ListInstrumentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListInstrumentsRequest.Merge(m, src)
}
func (m *ListInstrumentsRequest) XXX_Size() int {
	return xxx_messageInfo_ListInstrumentsRequest.Size(m)
}
func (m *ListInstrumentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListInstrumentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListInstrumentsRequest proto.InternalMessageInfo

func (m *ListInstrumentsRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type ListInstrumentsResponse struct {
	Instruments          []*Instrument `protobuf:"bytes,1,rep,name=instruments,proto3" json:"instruments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListInstrumentsResponse) Reset()         { *m = ListInstrumentsResponse{} }
func (m *ListInstrumentsResponse) String() string { return proto.CompactTextString(m) }
func (*ListInstrumentsResponse) ProtoMessage()    {}
func (*ListInstrumentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{11}
}

func (m *ListInstrumentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListInstrumentsResponse.Unmarshal(m, b)
}
func (m *ListInstrumentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListInstrumentsResponse.Marshal(b, m, deterministic)
}
func (m *ListInstrumentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListInstrumentsResponse.Merge(m, src)
}
func (m *ListInstrumentsResponse) XXX_Size() int {
	return xxx_messageInfo_ListInstrumentsResponse.Size(m)
}
func (m *ListInstrumentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListInstrumentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListInstrumentsResponse proto.InternalMessageInfo

func (m *ListInstrumentsResponse) GetInstruments() []*Instrument {
	if m != nil {
		return m.Instruments
	}
	return nil
}

type DeleteInstrumentRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteInstrumentRequest) Reset()         { *m = DeleteInstrumentRequest{} }
func (m *DeleteInstrumentRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteInstrumentRequest) ProtoMessage()    {}
func (*DeleteInstrumentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{12}
}

func (m *DeleteInstrumentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteInstrumentRequest.Unmarshal(m, b)
}
func (m *DeleteInstrumentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteInstrumentRequest.Marshal(b, m, deterministic)
}
func (m *DeleteInstrumentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteInstrumentRequest.Merge(m, src)
}
func (m *DeleteInstrumentRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteInstrumentRequest.Size(m)
}
func (m *DeleteInstrumentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteInstrumentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteInstrumentRequest proto.InternalMessageInfo

func (m *DeleteInstrumentRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UpdateInstrumentRequest struct {
	Id                   string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Instrument           *Instrument `protobuf:"bytes,2,opt,name=instrument,proto3" json:"instrument,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpdateInstrumentRequest) Reset()         { *m = UpdateInstrumentRequest{} }
func (m *UpdateInstrumentRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateInstrumentRequest) ProtoMessage()    {}
func (*UpdateInstrumentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{13}
}

func (m *UpdateInstrumentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateInstrumentRequest.Unmarshal(m, b)
}
func (m *UpdateInstrumentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateInstrumentRequest.Marshal(b, m, deterministic)
}
func (m *UpdateInstrumentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateInstrumentRequest.Merge(m, src)
}
func (m *UpdateInstrumentRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateInstrumentRequest.Size(m)
}
func (m *UpdateInstrumentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateInstrumentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateInstrumentRequest proto.InternalMessageInfo

func (m *UpdateInstrumentRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateInstrumentRequest) GetInstrument() *Instrument {
	if m != nil {
		return m.Instrument
	}
	return nil
}

func init() {
	proto.RegisterEnum("Action_Type", Action_Type_name, Action_Type_value)
	proto.RegisterType((*Instrument)(nil), "Instrument")
	proto.RegisterType((*Lineage)(nil), "Lineage")
	proto.RegisterType((*Action)(nil), "Action")
	proto.RegisterType((*Person)(nil), "Person")
	proto.RegisterType((*Role)(nil), "Role")
	proto.RegisterType((*Party)(nil), "Party")
	proto.RegisterType((*Event)(nil), "Event")
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*CreateInstrumentRequest)(nil), "CreateInstrumentRequest")
	proto.RegisterType((*GetInstrumentRequest)(nil), "GetInstrumentRequest")
	proto.RegisterType((*ListInstrumentsRequest)(nil), "ListInstrumentsRequest")
	proto.RegisterType((*ListInstrumentsResponse)(nil), "ListInstrumentsResponse")
	proto.RegisterType((*DeleteInstrumentRequest)(nil), "DeleteInstrumentRequest")
	proto.RegisterType((*UpdateInstrumentRequest)(nil), "UpdateInstrumentRequest")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 649 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x4d, 0xec, 0xd8, 0x89, 0xaf, 0xfb, 0xe3, 0x6f, 0xf4, 0x41, 0x4c, 0x0b, 0xc2, 0xcc, 0x02,
	0x05, 0x10, 0xb3, 0x08, 0x82, 0x0d, 0x6c, 0x4c, 0xea, 0xd2, 0x48, 0xa9, 0x13, 0x4d, 0x5c, 0x16,
	0xb0, 0xa8, 0x86, 0x64, 0x40, 0x96, 0x12, 0xc7, 0xf5, 0x4c, 0x2a, 0xe5, 0xfd, 0x78, 0x1d, 0xde,
	0x01, 0xf9, 0x27, 0xad, 0xb1, 0x13, 0xc4, 0xce, 0x73, 0xce, 0x9d, 0x7b, 0xcf, 0x3d, 0x67, 0x64,
	0x30, 0x58, 0x1c, 0x92, 0x38, 0x59, 0xc9, 0x15, 0xbe, 0x01, 0x18, 0x46, 0x42, 0x26, 0xeb, 0x25,
	0x8f, 0x24, 0x3a, 0x02, 0x25, 0x9c, 0xdb, 0x4d, 0xa7, 0xd9, 0x33, 0xa8, 0x12, 0xce, 0x11, 0x82,
	0x96, 0xdc, 0xc4, 0xdc, 0x56, 0x32, 0x24, 0xfb, 0x46, 0xa7, 0x60, 0x84, 0x42, 0xac, 0xf9, 0xfc,
	0x9a, 0x49, 0x5b, 0x75, 0x9a, 0x3d, 0x8d, 0x76, 0x72, 0xc0, 0x95, 0x08, 0x43, 0x7b, 0x11, 0x46,
	0x9c, 0xfd, 0xe0, 0x76, 0xcb, 0x69, 0xf6, 0xcc, 0x7e, 0x87, 0x8c, 0xf2, 0x33, 0xdd, 0x12, 0xf8,
	0x03, 0xb4, 0x0b, 0xac, 0x36, 0xef, 0x19, 0xb4, 0xd9, 0x4c, 0x86, 0xab, 0x48, 0xd8, 0x8a, 0xa3,
	0xf6, 0xcc, 0x7e, 0x9b, 0xb8, 0xd9, 0x99, 0x6e, 0x71, 0xfc, 0xab, 0x09, 0x7a, 0x8e, 0xd5, 0x6e,
	0x1f, 0x81, 0xc2, 0x64, 0xa6, 0x55, 0xa3, 0x0a, 0x93, 0xc8, 0x29, 0xd4, 0xa7, 0x22, 0x8f, 0xfa,
	0x07, 0x45, 0x2b, 0x12, 0x6c, 0x62, 0x5e, 0xec, 0xf2, 0x18, 0x34, 0x7e, 0xcb, 0x23, 0x59, 0x88,
	0xd5, 0x89, 0x97, 0x9e, 0x68, 0x0e, 0x22, 0x07, 0xcc, 0x39, 0x17, 0xb3, 0x24, 0x8c, 0xd3, 0x7b,
	0xb6, 0x96, 0x0d, 0x2a, 0x43, 0x98, 0x41, 0x2b, 0xed, 0x86, 0x0e, 0xa0, 0x33, 0xa0, 0x9e, 0x1b,
	0x0c, 0xc7, 0xbe, 0xd5, 0x40, 0x87, 0x60, 0x04, 0xd4, 0xf5, 0xa7, 0x93, 0x31, 0x0d, 0xac, 0x26,
	0x6a, 0x83, 0x7a, 0x35, 0xf5, 0x2c, 0x05, 0x1d, 0x83, 0x79, 0xe9, 0x0e, 0xfd, 0xc0, 0xf3, 0x5d,
	0x7f, 0xe0, 0x59, 0x2a, 0x32, 0x40, 0x73, 0xcf, 0x2e, 0x87, 0xbe, 0xd5, 0x4a, 0x3b, 0x9c, 0x0d,
	0xa7, 0x93, 0xf1, 0xd4, 0x1d, 0x59, 0x5a, 0x4a, 0x8c, 0x83, 0x0b, 0x8f, 0x5a, 0x3a, 0x0e, 0x40,
	0x9f, 0xf0, 0x44, 0xec, 0x58, 0xf7, 0x14, 0x8c, 0x05, 0x13, 0xf2, 0x3a, 0x62, 0xcb, 0x6d, 0x42,
	0x9d, 0x14, 0xf0, 0xd9, 0x92, 0xa3, 0x27, 0x00, 0xdf, 0xc3, 0x64, 0xcb, 0xaa, 0x19, 0x6b, 0x64,
	0x48, 0x4a, 0xe3, 0x11, 0xb4, 0xe8, 0x6a, 0xc1, 0x77, 0x05, 0x9e, 0xac, 0x16, 0x77, 0x81, 0xa7,
	0xdf, 0x55, 0x1b, 0xd4, 0xba, 0x0d, 0x03, 0xd0, 0x26, 0x2c, 0x91, 0x1b, 0xf4, 0x14, 0xf4, 0x38,
	0x13, 0x9b, 0xb5, 0x4c, 0xe3, 0xcb, 0xb5, 0xd3, 0x02, 0x46, 0x8f, 0x4a, 0xfd, 0xcd, 0xbe, 0x46,
	0x52, 0x11, 0xf9, 0x18, 0xfc, 0x15, 0xb4, 0xcc, 0xfd, 0x9a, 0xa6, 0xca, 0x7c, 0xa5, 0x36, 0x1f,
	0x39, 0xd0, 0x8e, 0x59, 0x22, 0x43, 0x2e, 0x6c, 0x35, 0x7b, 0x36, 0x3a, 0xc9, 0xf4, 0xd0, 0x2d,
	0x8c, 0xdb, 0xa0, 0x79, 0xcb, 0x58, 0x6e, 0xf0, 0x39, 0x74, 0x07, 0x09, 0x67, 0x92, 0xdf, 0xbf,
	0x7a, 0xca, 0x6f, 0xd6, 0x5c, 0x48, 0xf4, 0x0a, 0x20, 0xbc, 0x03, 0x8b, 0x05, 0x4c, 0x52, 0xaa,
	0x2b, 0xd1, 0xf8, 0x39, 0xfc, 0xff, 0x89, 0xcb, 0x7a, 0x93, 0x8a, 0x78, 0xfc, 0x12, 0x1e, 0x8e,
	0x42, 0x51, 0x2a, 0x14, 0xdb, 0x4a, 0x0b, 0xd4, 0x70, 0x2e, 0xec, 0xa6, 0xa3, 0xf6, 0x0c, 0x9a,
	0x7e, 0xe2, 0x0b, 0xe8, 0xd6, 0x6a, 0x45, 0xbc, 0x8a, 0x04, 0x47, 0xaf, 0xc1, 0xbc, 0x1f, 0x9e,
	0x5f, 0xaa, 0x88, 0x2b, 0xf3, 0xf8, 0x05, 0x74, 0xcf, 0xf8, 0x82, 0xef, 0xda, 0xb2, 0x2a, 0xf0,
	0x33, 0x74, 0xaf, 0xe2, 0x39, 0xfb, 0x87, 0xd2, 0x8a, 0x41, 0xca, 0x5f, 0x0d, 0xea, 0xff, 0x54,
	0xe0, 0xbf, 0x7b, 0x6a, 0xca, 0x93, 0xdb, 0x70, 0xc6, 0xd1, 0x7b, 0xb0, 0xaa, 0xf6, 0x23, 0x9b,
	0xec, 0x49, 0xe4, 0xa4, 0xdc, 0x1c, 0x37, 0xd0, 0x5b, 0x38, 0xfc, 0xc3, 0x73, 0xf4, 0x80, 0xec,
	0xca, 0xa0, 0x7a, 0xed, 0x1c, 0x8e, 0x2b, 0xb6, 0xa2, 0x2e, 0xd9, 0x1d, 0xca, 0x89, 0x4d, 0xf6,
	0x24, 0x80, 0x1b, 0xe8, 0x1d, 0x58, 0x55, 0x53, 0x91, 0x4d, 0xf6, 0xf8, 0x7c, 0xa2, 0x93, 0xfc,
	0xc1, 0x35, 0xd2, 0x9d, 0xab, 0x0e, 0x23, 0x9b, 0xec, 0x31, 0xbd, 0x22, 0xfe, 0x63, 0xe7, 0x8b,
	0x3e, 0x5b, 0xb0, 0x64, 0xbd, 0xfc, 0xa6, 0x67, 0x3f, 0xec, 0x37, 0xbf, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xc6, 0x0c, 0x62, 0xab, 0xbd, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InstrumentServiceClient is the client API for InstrumentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InstrumentServiceClient interface {
	CreateInstrument(ctx context.Context, in *CreateInstrumentRequest, opts ...grpc.CallOption) (*Instrument, error)
	GetInstrument(ctx context.Context, in *GetInstrumentRequest, opts ...grpc.CallOption) (*Instrument, error)
	ListInstruments(ctx context.Context, in *ListInstrumentsRequest, opts ...grpc.CallOption) (*ListInstrumentsResponse, error)
	DeleteInstrument(ctx context.Context, in *DeleteInstrumentRequest, opts ...grpc.CallOption) (*Empty, error)
	UpdateInstrument(ctx context.Context, in *UpdateInstrumentRequest, opts ...grpc.CallOption) (*Instrument, error)
}

type instrumentServiceClient struct {
	cc *grpc.ClientConn
}

func NewInstrumentServiceClient(cc *grpc.ClientConn) InstrumentServiceClient {
	return &instrumentServiceClient{cc}
}

func (c *instrumentServiceClient) CreateInstrument(ctx context.Context, in *CreateInstrumentRequest, opts ...grpc.CallOption) (*Instrument, error) {
	out := new(Instrument)
	err := c.cc.Invoke(ctx, "/InstrumentService/CreateInstrument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instrumentServiceClient) GetInstrument(ctx context.Context, in *GetInstrumentRequest, opts ...grpc.CallOption) (*Instrument, error) {
	out := new(Instrument)
	err := c.cc.Invoke(ctx, "/InstrumentService/GetInstrument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instrumentServiceClient) ListInstruments(ctx context.Context, in *ListInstrumentsRequest, opts ...grpc.CallOption) (*ListInstrumentsResponse, error) {
	out := new(ListInstrumentsResponse)
	err := c.cc.Invoke(ctx, "/InstrumentService/ListInstruments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instrumentServiceClient) DeleteInstrument(ctx context.Context, in *DeleteInstrumentRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/InstrumentService/DeleteInstrument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instrumentServiceClient) UpdateInstrument(ctx context.Context, in *UpdateInstrumentRequest, opts ...grpc.CallOption) (*Instrument, error) {
	out := new(Instrument)
	err := c.cc.Invoke(ctx, "/InstrumentService/UpdateInstrument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InstrumentServiceServer is the server API for InstrumentService service.
type InstrumentServiceServer interface {
	CreateInstrument(context.Context, *CreateInstrumentRequest) (*Instrument, error)
	GetInstrument(context.Context, *GetInstrumentRequest) (*Instrument, error)
	ListInstruments(context.Context, *ListInstrumentsRequest) (*ListInstrumentsResponse, error)
	DeleteInstrument(context.Context, *DeleteInstrumentRequest) (*Empty, error)
	UpdateInstrument(context.Context, *UpdateInstrumentRequest) (*Instrument, error)
}

func RegisterInstrumentServiceServer(s *grpc.Server, srv InstrumentServiceServer) {
	s.RegisterService(&_InstrumentService_serviceDesc, srv)
}

func _InstrumentService_CreateInstrument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInstrumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstrumentServiceServer).CreateInstrument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InstrumentService/CreateInstrument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstrumentServiceServer).CreateInstrument(ctx, req.(*CreateInstrumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstrumentService_GetInstrument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstrumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstrumentServiceServer).GetInstrument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InstrumentService/GetInstrument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstrumentServiceServer).GetInstrument(ctx, req.(*GetInstrumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstrumentService_ListInstruments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstrumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstrumentServiceServer).ListInstruments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InstrumentService/ListInstruments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstrumentServiceServer).ListInstruments(ctx, req.(*ListInstrumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstrumentService_DeleteInstrument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInstrumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstrumentServiceServer).DeleteInstrument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InstrumentService/DeleteInstrument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstrumentServiceServer).DeleteInstrument(ctx, req.(*DeleteInstrumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstrumentService_UpdateInstrument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInstrumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstrumentServiceServer).UpdateInstrument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InstrumentService/UpdateInstrument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstrumentServiceServer).UpdateInstrument(ctx, req.(*UpdateInstrumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InstrumentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "InstrumentService",
	HandlerType: (*InstrumentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInstrument",
			Handler:    _InstrumentService_CreateInstrument_Handler,
		},
		{
			MethodName: "GetInstrument",
			Handler:    _InstrumentService_GetInstrument_Handler,
		},
		{
			MethodName: "ListInstruments",
			Handler:    _InstrumentService_ListInstruments_Handler,
		},
		{
			MethodName: "DeleteInstrument",
			Handler:    _InstrumentService_DeleteInstrument_Handler,
		},
		{
			MethodName: "UpdateInstrument",
			Handler:    _InstrumentService_UpdateInstrument_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}