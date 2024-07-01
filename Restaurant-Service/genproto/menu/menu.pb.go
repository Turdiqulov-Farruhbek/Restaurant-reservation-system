// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.1
// source: submodule/reservation/menu.proto

package menu

import (
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

// Menu message definition
type Menu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RestaurantId string  `protobuf:"bytes,2,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
	Name         string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description  string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Price        float64 `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	CreatedAt    string  `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` // Timestamp with time zone
	UpdatedAt    string  `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"` // Timestamp with time zone
	DeletedAt    int64   `protobuf:"varint,8,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Menu) Reset() {
	*x = Menu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_reservation_menu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Menu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Menu) ProtoMessage() {}

func (x *Menu) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_reservation_menu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Menu.ProtoReflect.Descriptor instead.
func (*Menu) Descriptor() ([]byte, []int) {
	return file_submodule_reservation_menu_proto_rawDescGZIP(), []int{0}
}

func (x *Menu) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Menu) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

func (x *Menu) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Menu) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Menu) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Menu) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Menu) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Menu) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

// Requests and Responses for Menu CRUD
type CreateMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Menu *Menu `protobuf:"bytes,1,opt,name=menu,proto3" json:"menu,omitempty"`
}

func (x *CreateMenuRequest) Reset() {
	*x = CreateMenuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_reservation_menu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMenuRequest) ProtoMessage() {}

func (x *CreateMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_reservation_menu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMenuRequest.ProtoReflect.Descriptor instead.
func (*CreateMenuRequest) Descriptor() ([]byte, []int) {
	return file_submodule_reservation_menu_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMenuRequest) GetMenu() *Menu {
	if x != nil {
		return x.Menu
	}
	return nil
}

type CreateMenuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Menu *Menu `protobuf:"bytes,1,opt,name=menu,proto3" json:"menu,omitempty"`
}

func (x *CreateMenuResponse) Reset() {
	*x = CreateMenuResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_reservation_menu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMenuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMenuResponse) ProtoMessage() {}

func (x *CreateMenuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_reservation_menu_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMenuResponse.ProtoReflect.Descriptor instead.
func (*CreateMenuResponse) Descriptor() ([]byte, []int) {
	return file_submodule_reservation_menu_proto_rawDescGZIP(), []int{2}
}

func (x *CreateMenuResponse) GetMenu() *Menu {
	if x != nil {
		return x.Menu
	}
	return nil
}

type GetMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMenuRequest) Reset() {
	*x = GetMenuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_reservation_menu_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMenuRequest) ProtoMessage() {}

func (x *GetMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_reservation_menu_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMenuRequest.ProtoReflect.Descriptor instead.
func (*GetMenuRequest) Descriptor() ([]byte, []int) {
	return file_submodule_reservation_menu_proto_rawDescGZIP(), []int{3}
}

func (x *GetMenuRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetMenuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Menu *Menu `protobuf:"bytes,1,opt,name=menu,proto3" json:"menu,omitempty"`
}

func (x *GetMenuResponse) Reset() {
	*x = GetMenuResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_reservation_menu_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMenuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMenuResponse) ProtoMessage() {}

func (x *GetMenuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_reservation_menu_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMenuResponse.ProtoReflect.Descriptor instead.
func (*GetMenuResponse) Descriptor() ([]byte, []int) {
	return file_submodule_reservation_menu_proto_rawDescGZIP(), []int{4}
}

func (x *GetMenuResponse) GetMenu() *Menu {
	if x != nil {
		return x.Menu
	}
	return nil
}

type UpdateMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Menu *Menu `protobuf:"bytes,1,opt,name=menu,proto3" json:"menu,omitempty"`
}

func (x *UpdateMenuRequest) Reset() {
	*x = UpdateMenuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_reservation_menu_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMenuRequest) ProtoMessage() {}

func (x *UpdateMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_reservation_menu_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMenuRequest.ProtoReflect.Descriptor instead.
func (*UpdateMenuRequest) Descriptor() ([]byte, []int) {
	return file_submodule_reservation_menu_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateMenuRequest) GetMenu() *Menu {
	if x != nil {
		return x.Menu
	}
	return nil
}

type UpdateMenuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Menu *Menu `protobuf:"bytes,1,opt,name=menu,proto3" json:"menu,omitempty"`
}

func (x *UpdateMenuResponse) Reset() {
	*x = UpdateMenuResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_reservation_menu_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMenuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMenuResponse) ProtoMessage() {}

func (x *UpdateMenuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_reservation_menu_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMenuResponse.ProtoReflect.Descriptor instead.
func (*UpdateMenuResponse) Descriptor() ([]byte, []int) {
	return file_submodule_reservation_menu_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateMenuResponse) GetMenu() *Menu {
	if x != nil {
		return x.Menu
	}
	return nil
}

type DeleteMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteMenuRequest) Reset() {
	*x = DeleteMenuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_reservation_menu_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMenuRequest) ProtoMessage() {}

func (x *DeleteMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_reservation_menu_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMenuRequest.ProtoReflect.Descriptor instead.
func (*DeleteMenuRequest) Descriptor() ([]byte, []int) {
	return file_submodule_reservation_menu_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteMenuRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteMenuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteMenuResponse) Reset() {
	*x = DeleteMenuResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_reservation_menu_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMenuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMenuResponse) ProtoMessage() {}

func (x *DeleteMenuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_reservation_menu_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMenuResponse.ProtoReflect.Descriptor instead.
func (*DeleteMenuResponse) Descriptor() ([]byte, []int) {
	return file_submodule_reservation_menu_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteMenuResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Request and Response for GetAllMenus
type GetAllMenusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RestaurantId string  `protobuf:"bytes,1,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
	Name         string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description  string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	MinPrice     float64 `protobuf:"fixed64,4,opt,name=minPrice,proto3" json:"minPrice,omitempty"`
	MaxPrice     float64 `protobuf:"fixed64,5,opt,name=maxPrice,proto3" json:"maxPrice,omitempty"`
}

func (x *GetAllMenusRequest) Reset() {
	*x = GetAllMenusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_reservation_menu_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMenusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMenusRequest) ProtoMessage() {}

func (x *GetAllMenusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_reservation_menu_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMenusRequest.ProtoReflect.Descriptor instead.
func (*GetAllMenusRequest) Descriptor() ([]byte, []int) {
	return file_submodule_reservation_menu_proto_rawDescGZIP(), []int{9}
}

func (x *GetAllMenusRequest) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

func (x *GetAllMenusRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetAllMenusRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetAllMenusRequest) GetMinPrice() float64 {
	if x != nil {
		return x.MinPrice
	}
	return 0
}

func (x *GetAllMenusRequest) GetMaxPrice() float64 {
	if x != nil {
		return x.MaxPrice
	}
	return 0
}

type GetAllMenusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Menus []*Menu `protobuf:"bytes,1,rep,name=menus,proto3" json:"menus,omitempty"`
}

func (x *GetAllMenusResponse) Reset() {
	*x = GetAllMenusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_reservation_menu_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMenusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMenusResponse) ProtoMessage() {}

func (x *GetAllMenusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_reservation_menu_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMenusResponse.ProtoReflect.Descriptor instead.
func (*GetAllMenusResponse) Descriptor() ([]byte, []int) {
	return file_submodule_reservation_menu_proto_rawDescGZIP(), []int{10}
}

func (x *GetAllMenusResponse) GetMenus() []*Menu {
	if x != nil {
		return x.Menus
	}
	return nil
}

var File_submodule_reservation_menu_proto protoreflect.FileDescriptor

var file_submodule_reservation_menu_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x22, 0xe4, 0x01, 0x0a, 0x04, 0x4d, 0x65, 0x6e,
	0x75, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x33, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x04,
	0x6d, 0x65, 0x6e, 0x75, 0x22, 0x34, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x6d, 0x65,
	0x6e, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e,
	0x4d, 0x65, 0x6e, 0x75, 0x52, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x22,
	0x33, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x04,
	0x6d, 0x65, 0x6e, 0x75, 0x22, 0x34, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x6d, 0x65,
	0x6e, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e,
	0x4d, 0x65, 0x6e, 0x75, 0x52, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x2e, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0xa7, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x37, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x20, 0x0a, 0x05, 0x6d, 0x65, 0x6e, 0x75, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x05, 0x6d, 0x65, 0x6e,
	0x75, 0x73, 0x32, 0xcc, 0x02, 0x0a, 0x0b, 0x4d, 0x65, 0x6e, 0x75, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75,
	0x12, 0x17, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x65, 0x6e, 0x75,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x14,
	0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x17, 0x2e, 0x6d, 0x65, 0x6e, 0x75,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x17, 0x2e, 0x6d, 0x65, 0x6e,
	0x75, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x12, 0x18, 0x2e, 0x6d,
	0x65, 0x6e, 0x75, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_submodule_reservation_menu_proto_rawDescOnce sync.Once
	file_submodule_reservation_menu_proto_rawDescData = file_submodule_reservation_menu_proto_rawDesc
)

func file_submodule_reservation_menu_proto_rawDescGZIP() []byte {
	file_submodule_reservation_menu_proto_rawDescOnce.Do(func() {
		file_submodule_reservation_menu_proto_rawDescData = protoimpl.X.CompressGZIP(file_submodule_reservation_menu_proto_rawDescData)
	})
	return file_submodule_reservation_menu_proto_rawDescData
}

var file_submodule_reservation_menu_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_submodule_reservation_menu_proto_goTypes = []any{
	(*Menu)(nil),                // 0: menu.Menu
	(*CreateMenuRequest)(nil),   // 1: menu.CreateMenuRequest
	(*CreateMenuResponse)(nil),  // 2: menu.CreateMenuResponse
	(*GetMenuRequest)(nil),      // 3: menu.GetMenuRequest
	(*GetMenuResponse)(nil),     // 4: menu.GetMenuResponse
	(*UpdateMenuRequest)(nil),   // 5: menu.UpdateMenuRequest
	(*UpdateMenuResponse)(nil),  // 6: menu.UpdateMenuResponse
	(*DeleteMenuRequest)(nil),   // 7: menu.DeleteMenuRequest
	(*DeleteMenuResponse)(nil),  // 8: menu.DeleteMenuResponse
	(*GetAllMenusRequest)(nil),  // 9: menu.GetAllMenusRequest
	(*GetAllMenusResponse)(nil), // 10: menu.GetAllMenusResponse
}
var file_submodule_reservation_menu_proto_depIdxs = []int32{
	0,  // 0: menu.CreateMenuRequest.menu:type_name -> menu.Menu
	0,  // 1: menu.CreateMenuResponse.menu:type_name -> menu.Menu
	0,  // 2: menu.GetMenuResponse.menu:type_name -> menu.Menu
	0,  // 3: menu.UpdateMenuRequest.menu:type_name -> menu.Menu
	0,  // 4: menu.UpdateMenuResponse.menu:type_name -> menu.Menu
	0,  // 5: menu.GetAllMenusResponse.menus:type_name -> menu.Menu
	1,  // 6: menu.MenuService.CreateMenu:input_type -> menu.CreateMenuRequest
	3,  // 7: menu.MenuService.GetMenu:input_type -> menu.GetMenuRequest
	5,  // 8: menu.MenuService.UpdateMenu:input_type -> menu.UpdateMenuRequest
	7,  // 9: menu.MenuService.DeleteMenu:input_type -> menu.DeleteMenuRequest
	9,  // 10: menu.MenuService.GetAllMenus:input_type -> menu.GetAllMenusRequest
	2,  // 11: menu.MenuService.CreateMenu:output_type -> menu.CreateMenuResponse
	4,  // 12: menu.MenuService.GetMenu:output_type -> menu.GetMenuResponse
	6,  // 13: menu.MenuService.UpdateMenu:output_type -> menu.UpdateMenuResponse
	8,  // 14: menu.MenuService.DeleteMenu:output_type -> menu.DeleteMenuResponse
	10, // 15: menu.MenuService.GetAllMenus:output_type -> menu.GetAllMenusResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_submodule_reservation_menu_proto_init() }
func file_submodule_reservation_menu_proto_init() {
	if File_submodule_reservation_menu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_submodule_reservation_menu_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Menu); i {
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
		file_submodule_reservation_menu_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateMenuRequest); i {
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
		file_submodule_reservation_menu_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateMenuResponse); i {
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
		file_submodule_reservation_menu_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetMenuRequest); i {
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
		file_submodule_reservation_menu_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetMenuResponse); i {
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
		file_submodule_reservation_menu_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateMenuRequest); i {
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
		file_submodule_reservation_menu_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateMenuResponse); i {
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
		file_submodule_reservation_menu_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteMenuRequest); i {
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
		file_submodule_reservation_menu_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteMenuResponse); i {
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
		file_submodule_reservation_menu_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllMenusRequest); i {
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
		file_submodule_reservation_menu_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllMenusResponse); i {
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
			RawDescriptor: file_submodule_reservation_menu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_submodule_reservation_menu_proto_goTypes,
		DependencyIndexes: file_submodule_reservation_menu_proto_depIdxs,
		MessageInfos:      file_submodule_reservation_menu_proto_msgTypes,
	}.Build()
	File_submodule_reservation_menu_proto = out.File
	file_submodule_reservation_menu_proto_rawDesc = nil
	file_submodule_reservation_menu_proto_goTypes = nil
	file_submodule_reservation_menu_proto_depIdxs = nil
}
