package proto

import (
	date "google.golang.org/genproto/googleapis/type/date"
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

// Desribes the type of pets.
type PetType int32

const (
	// The type was not set.
	PetType_PTUnknown PetType = 0
	// The pet is a canine.
	PetType_PTCanine PetType = 1
	// The pet is a feline.
	PetType_PTFeline PetType = 2
	// The pet is a bird.
	PetType_PTBird PetType = 3
	// The pet is a reptile.
	PetType_PTReptile PetType = 4
)

// Enum value maps for PetType.
var (
	PetType_name = map[int32]string{
		0: "PTUnknown",
		1: "PTCanine",
		2: "PTFeline",
		3: "PTBird",
		4: "PTReptile",
	}
	PetType_value = map[string]int32{
		"PTUnknown": 0,
		"PTCanine":  1,
		"PTFeline":  2,
		"PTBird":    3,
		"PTReptile": 4,
	}
)

func (x PetType) Enum() *PetType {
	p := new(PetType)
	*p = x
	return p
}

func (x PetType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PetType) Descriptor() protoreflect.EnumDescriptor {
	return file_petstore_proto_enumTypes[0].Descriptor()
}

func (PetType) Type() protoreflect.EnumType {
	return &file_petstore_proto_enumTypes[0]
}

func (x PetType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PetType.Descriptor instead.
func (PetType) EnumDescriptor() ([]byte, []int) {
	return file_petstore_proto_rawDescGZIP(), []int{0}
}

// Types of OTEL sampling we support.
type SamplerType int32

const (
	SamplerType_STUnknown SamplerType = 0
	SamplerType_STNever   SamplerType = 1
	SamplerType_STAlways  SamplerType = 2
	SamplerType_STFloat   SamplerType = 3
)

// Enum value maps for SamplerType.
var (
	SamplerType_name = map[int32]string{
		0: "STUnknown",
		1: "STNever",
		2: "STAlways",
		3: "STFloat",
	}
	SamplerType_value = map[string]int32{
		"STUnknown": 0,
		"STNever":   1,
		"STAlways":  2,
		"STFloat":   3,
	}
)

func (x SamplerType) Enum() *SamplerType {
	p := new(SamplerType)
	*p = x
	return p
}

func (x SamplerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SamplerType) Descriptor() protoreflect.EnumDescriptor {
	return file_petstore_proto_enumTypes[1].Descriptor()
}

func (SamplerType) Type() protoreflect.EnumType {
	return &file_petstore_proto_enumTypes[1]
}

func (x SamplerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SamplerType.Descriptor instead.
func (SamplerType) EnumDescriptor() ([]byte, []int) {
	return file_petstore_proto_rawDescGZIP(), []int{1}
}

// Represents a range of dates.
type DateRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// When to start the range, this is inclusive.
	Start *date.Date `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	// When to end the range, this is exclusive.
	End *date.Date `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *DateRange) Reset() {
	*x = DateRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DateRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateRange) ProtoMessage() {}

func (x *DateRange) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateRange.ProtoReflect.Descriptor instead.
func (*DateRange) Descriptor() ([]byte, []int) {
	return file_petstore_proto_rawDescGZIP(), []int{0}
}

func (x *DateRange) GetStart() *date.Date {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *DateRange) GetEnd() *date.Date {
	if x != nil {
		return x.End
	}
	return nil
}

// Represents a unique pet.
type Pet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A UUIDv4 for this pet. This can never be set on an AddPet().
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the pet.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The type of pet.
	Type PetType `protobuf:"varint,3,opt,name=type,proto3,enum=petstore.PetType" json:"type,omitempty"`
	// The pet's birthday.
	Birthday *date.Date `protobuf:"bytes,4,opt,name=birthday,proto3" json:"birthday,omitempty"`
}

func (x *Pet) Reset() {
	*x = Pet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pet) ProtoMessage() {}

func (x *Pet) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pet.ProtoReflect.Descriptor instead.
func (*Pet) Descriptor() ([]byte, []int) {
	return file_petstore_proto_rawDescGZIP(), []int{1}
}

func (x *Pet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Pet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pet) GetType() PetType {
	if x != nil {
		return x.Type
	}
	return PetType_PTUnknown
}

func (x *Pet) GetBirthday() *date.Date {
	if x != nil {
		return x.Birthday
	}
	return nil
}

// The request used to add a pets to the system.
type AddPetsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The pet information to add. Pet.id must not be set.
	Pets []*Pet `protobuf:"bytes,1,rep,name=pets,proto3" json:"pets,omitempty"`
}

func (x *AddPetsReq) Reset() {
	*x = AddPetsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPetsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPetsReq) ProtoMessage() {}

func (x *AddPetsReq) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPetsReq.ProtoReflect.Descriptor instead.
func (*AddPetsReq) Descriptor() ([]byte, []int) {
	return file_petstore_proto_rawDescGZIP(), []int{2}
}

func (x *AddPetsReq) GetPets() []*Pet {
	if x != nil {
		return x.Pets
	}
	return nil
}

// The response do AddPets().
type AddPetsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The IDs of the pets that were added.
	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *AddPetsResp) Reset() {
	*x = AddPetsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPetsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPetsResp) ProtoMessage() {}

func (x *AddPetsResp) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPetsResp.ProtoReflect.Descriptor instead.
func (*AddPetsResp) Descriptor() ([]byte, []int) {
	return file_petstore_proto_rawDescGZIP(), []int{3}
}

func (x *AddPetsResp) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

// The request used to update pets in the system.
type UpdatePetsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The pet information to update. Pet.id must be set.
	Pets []*Pet `protobuf:"bytes,1,rep,name=pets,proto3" json:"pets,omitempty"`
}

func (x *UpdatePetsReq) Reset() {
	*x = UpdatePetsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePetsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePetsReq) ProtoMessage() {}

func (x *UpdatePetsReq) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePetsReq.ProtoReflect.Descriptor instead.
func (*UpdatePetsReq) Descriptor() ([]byte, []int) {
	return file_petstore_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePetsReq) GetPets() []*Pet {
	if x != nil {
		return x.Pets
	}
	return nil
}

// The response do UpdatePets().
type UpdatePetsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdatePetsResp) Reset() {
	*x = UpdatePetsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePetsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePetsResp) ProtoMessage() {}

func (x *UpdatePetsResp) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePetsResp.ProtoReflect.Descriptor instead.
func (*UpdatePetsResp) Descriptor() ([]byte, []int) {
	return file_petstore_proto_rawDescGZIP(), []int{5}
}

// Used to indicate which pets to delete. This is an all or nothing request.
type DeletePetsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The IDs of the pets to delete.
	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeletePetsReq) Reset() {
	*x = DeletePetsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePetsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePetsReq) ProtoMessage() {}

func (x *DeletePetsReq) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePetsReq.ProtoReflect.Descriptor instead.
func (*DeletePetsReq) Descriptor() ([]byte, []int) {
	return file_petstore_proto_rawDescGZIP(), []int{6}
}

func (x *DeletePetsReq) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

// The response to a DeletePet().
type DeletePetsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeletePetsResp) Reset() {
	*x = DeletePetsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePetsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePetsResp) ProtoMessage() {}

func (x *DeletePetsResp) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePetsResp.ProtoReflect.Descriptor instead.
func (*DeletePetsResp) Descriptor() ([]byte, []int) {
	return file_petstore_proto_rawDescGZIP(), []int{7}
}

// The request to search for pets.
type SearchPetsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Pet names to filter by.
	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	// Pet types to filter by.
	Types []PetType `protobuf:"varint,2,rep,packed,name=types,proto3,enum=petstore.PetType" json:"types,omitempty"`
	// Birthdays to filter by.
	BirthdateRange *DateRange `protobuf:"bytes,3,opt,name=birthdate_range,json=birthdateRange,proto3" json:"birthdate_range,omitempty"`
}

func (x *SearchPetsReq) Reset() {
	*x = SearchPetsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchPetsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchPetsReq) ProtoMessage() {}

func (x *SearchPetsReq) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchPetsReq.ProtoReflect.Descriptor instead.
func (*SearchPetsReq) Descriptor() ([]byte, []int) {
	return file_petstore_proto_rawDescGZIP(), []int{8}
}

func (x *SearchPetsReq) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *SearchPetsReq) GetTypes() []PetType {
	if x != nil {
		return x.Types
	}
	return nil
}

func (x *SearchPetsReq) GetBirthdateRange() *DateRange {
	if x != nil {
		return x.BirthdateRange
	}
	return nil
}

type Sampler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of sampling to change to.
	Type SamplerType `protobuf:"varint,1,opt,name=type,proto3,enum=petstore.SamplerType" json:"type,omitempty"`
	// This is the sampling rate if type == STFloat. Values must be
	// > 0 and <= 1.0 .
	FloatValue float64 `protobuf:"fixed64,2,opt,name=float_value,json=floatValue,proto3" json:"float_value,omitempty"`
}

func (x *Sampler) Reset() {
	*x = Sampler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sampler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sampler) ProtoMessage() {}

func (x *Sampler) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sampler.ProtoReflect.Descriptor instead.
func (*Sampler) Descriptor() ([]byte, []int) {
	return file_petstore_proto_rawDescGZIP(), []int{9}
}

func (x *Sampler) GetType() SamplerType {
	if x != nil {
		return x.Type
	}
	return SamplerType_STUnknown
}

func (x *Sampler) GetFloatValue() float64 {
	if x != nil {
		return x.FloatValue
	}
	return 0
}

// Used to request we change the OTEL sampling.
type ChangeSamplerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sampler *Sampler `protobuf:"bytes,1,opt,name=sampler,proto3" json:"sampler,omitempty"`
}

func (x *ChangeSamplerReq) Reset() {
	*x = ChangeSamplerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeSamplerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeSamplerReq) ProtoMessage() {}

func (x *ChangeSamplerReq) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeSamplerReq.ProtoReflect.Descriptor instead.
func (*ChangeSamplerReq) Descriptor() ([]byte, []int) {
	return file_petstore_proto_rawDescGZIP(), []int{10}
}

func (x *ChangeSamplerReq) GetSampler() *Sampler {
	if x != nil {
		return x.Sampler
	}
	return nil
}

// The response to a sampling change.
type ChangeSamplerResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChangeSamplerResp) Reset() {
	*x = ChangeSamplerResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petstore_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeSamplerResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeSamplerResp) ProtoMessage() {}

func (x *ChangeSamplerResp) ProtoReflect() protoreflect.Message {
	mi := &file_petstore_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeSamplerResp.ProtoReflect.Descriptor instead.
func (*ChangeSamplerResp) Descriptor() ([]byte, []int) {
	return file_petstore_proto_rawDescGZIP(), []int{11}
}

var File_petstore_proto protoreflect.FileDescriptor

var file_petstore_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x59, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x27, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74,
	0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x23, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x7f, 0x0a,
	0x03, 0x50, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x50, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x2d, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x44, 0x61, 0x74, 0x65, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x22, 0x2f,
	0x0a, 0x0a, 0x41, 0x64, 0x64, 0x50, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x04,
	0x70, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x65, 0x74,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x65, 0x74, 0x52, 0x04, 0x70, 0x65, 0x74, 0x73, 0x22,
	0x1f, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x50, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x22, 0x32, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x12, 0x21, 0x0a, 0x04, 0x70, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x65, 0x74, 0x52, 0x04,
	0x70, 0x65, 0x74, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x21, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x8c, 0x01, 0x0a, 0x0d,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x65,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x0f,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0e, 0x62, 0x69, 0x72, 0x74,
	0x68, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x55, 0x0a, 0x07, 0x53, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x3f, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x2b, 0x0a, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x52, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x72, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x2a, 0x4f, 0x0a, 0x07, 0x50, 0x65, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x54, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x54, 0x43, 0x61, 0x6e, 0x69, 0x6e, 0x65, 0x10, 0x01, 0x12,
	0x0c, 0x0a, 0x08, 0x50, 0x54, 0x46, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x02, 0x12, 0x0a, 0x0a,
	0x06, 0x50, 0x54, 0x42, 0x69, 0x72, 0x64, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x54, 0x52,
	0x65, 0x70, 0x74, 0x69, 0x6c, 0x65, 0x10, 0x04, 0x2a, 0x44, 0x0a, 0x0b, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x54, 0x55, 0x6e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x4e, 0x65, 0x76, 0x65,
	0x72, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x41, 0x6c, 0x77, 0x61, 0x79, 0x73, 0x10,
	0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x10, 0x03, 0x32, 0xd0,
	0x02, 0x0a, 0x08, 0x50, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x50, 0x65, 0x74, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70,
	0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x65, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x65, 0x74, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x70,
	0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x65, 0x74, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a,
	0x18, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0a, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x65, 0x74, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x65, 0x74, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x65, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x65,
	0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4a, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x50, 0x61, 0x63, 0x6b, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2f,
	0x47, 0x6f, 0x2d, 0x66, 0x6f, 0x72, 0x2d, 0x44, 0x65, 0x76, 0x4f, 0x70, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_petstore_proto_rawDescOnce sync.Once
	file_petstore_proto_rawDescData = file_petstore_proto_rawDesc
)

func file_petstore_proto_rawDescGZIP() []byte {
	file_petstore_proto_rawDescOnce.Do(func() {
		file_petstore_proto_rawDescData = protoimpl.X.CompressGZIP(file_petstore_proto_rawDescData)
	})
	return file_petstore_proto_rawDescData
}

var file_petstore_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_petstore_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_petstore_proto_goTypes = []interface{}{
	(PetType)(0),              // 0: petstore.PetType
	(SamplerType)(0),          // 1: petstore.SamplerType
	(*DateRange)(nil),         // 2: petstore.DateRange
	(*Pet)(nil),               // 3: petstore.Pet
	(*AddPetsReq)(nil),        // 4: petstore.AddPetsReq
	(*AddPetsResp)(nil),       // 5: petstore.AddPetsResp
	(*UpdatePetsReq)(nil),     // 6: petstore.UpdatePetsReq
	(*UpdatePetsResp)(nil),    // 7: petstore.UpdatePetsResp
	(*DeletePetsReq)(nil),     // 8: petstore.DeletePetsReq
	(*DeletePetsResp)(nil),    // 9: petstore.DeletePetsResp
	(*SearchPetsReq)(nil),     // 10: petstore.SearchPetsReq
	(*Sampler)(nil),           // 11: petstore.Sampler
	(*ChangeSamplerReq)(nil),  // 12: petstore.ChangeSamplerReq
	(*ChangeSamplerResp)(nil), // 13: petstore.ChangeSamplerResp
	(*date.Date)(nil),         // 14: google.type.Date
}
var file_petstore_proto_depIdxs = []int32{
	14, // 0: petstore.DateRange.start:type_name -> google.type.Date
	14, // 1: petstore.DateRange.end:type_name -> google.type.Date
	0,  // 2: petstore.Pet.type:type_name -> petstore.PetType
	14, // 3: petstore.Pet.birthday:type_name -> google.type.Date
	3,  // 4: petstore.AddPetsReq.pets:type_name -> petstore.Pet
	3,  // 5: petstore.UpdatePetsReq.pets:type_name -> petstore.Pet
	0,  // 6: petstore.SearchPetsReq.types:type_name -> petstore.PetType
	2,  // 7: petstore.SearchPetsReq.birthdate_range:type_name -> petstore.DateRange
	1,  // 8: petstore.Sampler.type:type_name -> petstore.SamplerType
	11, // 9: petstore.ChangeSamplerReq.sampler:type_name -> petstore.Sampler
	4,  // 10: petstore.PetStore.AddPets:input_type -> petstore.AddPetsReq
	6,  // 11: petstore.PetStore.UpdatePets:input_type -> petstore.UpdatePetsReq
	8,  // 12: petstore.PetStore.DeletePets:input_type -> petstore.DeletePetsReq
	10, // 13: petstore.PetStore.SearchPets:input_type -> petstore.SearchPetsReq
	12, // 14: petstore.PetStore.ChangeSampler:input_type -> petstore.ChangeSamplerReq
	5,  // 15: petstore.PetStore.AddPets:output_type -> petstore.AddPetsResp
	7,  // 16: petstore.PetStore.UpdatePets:output_type -> petstore.UpdatePetsResp
	9,  // 17: petstore.PetStore.DeletePets:output_type -> petstore.DeletePetsResp
	3,  // 18: petstore.PetStore.SearchPets:output_type -> petstore.Pet
	13, // 19: petstore.PetStore.ChangeSampler:output_type -> petstore.ChangeSamplerResp
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_petstore_proto_init() }
func file_petstore_proto_init() {
	if File_petstore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_petstore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DateRange); i {
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
		file_petstore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pet); i {
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
		file_petstore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPetsReq); i {
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
		file_petstore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPetsResp); i {
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
		file_petstore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePetsReq); i {
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
		file_petstore_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePetsResp); i {
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
		file_petstore_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePetsReq); i {
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
		file_petstore_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePetsResp); i {
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
		file_petstore_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchPetsReq); i {
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
		file_petstore_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sampler); i {
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
		file_petstore_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeSamplerReq); i {
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
		file_petstore_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeSamplerResp); i {
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
			RawDescriptor: file_petstore_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_petstore_proto_goTypes,
		DependencyIndexes: file_petstore_proto_depIdxs,
		EnumInfos:         file_petstore_proto_enumTypes,
		MessageInfos:      file_petstore_proto_msgTypes,
	}.Build()
	File_petstore_proto = out.File
	file_petstore_proto_rawDesc = nil
	file_petstore_proto_goTypes = nil
	file_petstore_proto_depIdxs = nil
}