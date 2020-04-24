// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.8.0
// source: recipes-service.proto

package generated

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AddRecipeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recipe *Recipe `protobuf:"bytes,1,opt,name=recipe,proto3" json:"recipe,omitempty"`
}

func (x *AddRecipeRequest) Reset() {
	*x = AddRecipeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recipes_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRecipeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRecipeRequest) ProtoMessage() {}

func (x *AddRecipeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recipes_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRecipeRequest.ProtoReflect.Descriptor instead.
func (*AddRecipeRequest) Descriptor() ([]byte, []int) {
	return file_recipes_service_proto_rawDescGZIP(), []int{0}
}

func (x *AddRecipeRequest) GetRecipe() *Recipe {
	if x != nil {
		return x.Recipe
	}
	return nil
}

type AddRecipeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AddRecipeResponse) Reset() {
	*x = AddRecipeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recipes_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRecipeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRecipeResponse) ProtoMessage() {}

func (x *AddRecipeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_recipes_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRecipeResponse.ProtoReflect.Descriptor instead.
func (*AddRecipeResponse) Descriptor() ([]byte, []int) {
	return file_recipes_service_proto_rawDescGZIP(), []int{1}
}

func (x *AddRecipeResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ListAllRecipesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAllRecipesRequest) Reset() {
	*x = ListAllRecipesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recipes_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllRecipesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllRecipesRequest) ProtoMessage() {}

func (x *ListAllRecipesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recipes_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllRecipesRequest.ProtoReflect.Descriptor instead.
func (*ListAllRecipesRequest) Descriptor() ([]byte, []int) {
	return file_recipes_service_proto_rawDescGZIP(), []int{2}
}

type ListAllRecipesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recipe *Recipe `protobuf:"bytes,1,opt,name=recipe,proto3" json:"recipe,omitempty"`
}

func (x *ListAllRecipesResponse) Reset() {
	*x = ListAllRecipesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recipes_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllRecipesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllRecipesResponse) ProtoMessage() {}

func (x *ListAllRecipesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_recipes_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllRecipesResponse.ProtoReflect.Descriptor instead.
func (*ListAllRecipesResponse) Descriptor() ([]byte, []int) {
	return file_recipes_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListAllRecipesResponse) GetRecipe() *Recipe {
	if x != nil {
		return x.Recipe
	}
	return nil
}

type Recipe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cuisine string `protobuf:"bytes,2,opt,name=cuisine,proto3" json:"cuisine,omitempty"`
}

func (x *Recipe) Reset() {
	*x = Recipe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recipes_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recipe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recipe) ProtoMessage() {}

func (x *Recipe) ProtoReflect() protoreflect.Message {
	mi := &file_recipes_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recipe.ProtoReflect.Descriptor instead.
func (*Recipe) Descriptor() ([]byte, []int) {
	return file_recipes_service_proto_rawDescGZIP(), []int{4}
}

func (x *Recipe) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Recipe) GetCuisine() string {
	if x != nil {
		return x.Cuisine
	}
	return ""
}

type Ingredient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Quantity string `protobuf:"bytes,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Ingredient) Reset() {
	*x = Ingredient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recipes_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ingredient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ingredient) ProtoMessage() {}

func (x *Ingredient) ProtoReflect() protoreflect.Message {
	mi := &file_recipes_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ingredient.ProtoReflect.Descriptor instead.
func (*Ingredient) Descriptor() ([]byte, []int) {
	return file_recipes_service_proto_rawDescGZIP(), []int{5}
}

func (x *Ingredient) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Ingredient) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

type ListAllIngredientsAtHomeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ingredient *Ingredient `protobuf:"bytes,1,opt,name=ingredient,proto3" json:"ingredient,omitempty"`
}

func (x *ListAllIngredientsAtHomeRequest) Reset() {
	*x = ListAllIngredientsAtHomeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recipes_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllIngredientsAtHomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllIngredientsAtHomeRequest) ProtoMessage() {}

func (x *ListAllIngredientsAtHomeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recipes_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllIngredientsAtHomeRequest.ProtoReflect.Descriptor instead.
func (*ListAllIngredientsAtHomeRequest) Descriptor() ([]byte, []int) {
	return file_recipes_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListAllIngredientsAtHomeRequest) GetIngredient() *Ingredient {
	if x != nil {
		return x.Ingredient
	}
	return nil
}

type ListAllIngredientsAtHomeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *ListAllIngredientsAtHomeResponse) Reset() {
	*x = ListAllIngredientsAtHomeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recipes_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllIngredientsAtHomeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllIngredientsAtHomeResponse) ProtoMessage() {}

func (x *ListAllIngredientsAtHomeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_recipes_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllIngredientsAtHomeResponse.ProtoReflect.Descriptor instead.
func (*ListAllIngredientsAtHomeResponse) Descriptor() ([]byte, []int) {
	return file_recipes_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListAllIngredientsAtHomeResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetIngredientsForAllRecipesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recipe *Recipe `protobuf:"bytes,1,opt,name=recipe,proto3" json:"recipe,omitempty"`
}

func (x *GetIngredientsForAllRecipesRequest) Reset() {
	*x = GetIngredientsForAllRecipesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recipes_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIngredientsForAllRecipesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIngredientsForAllRecipesRequest) ProtoMessage() {}

func (x *GetIngredientsForAllRecipesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recipes_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIngredientsForAllRecipesRequest.ProtoReflect.Descriptor instead.
func (*GetIngredientsForAllRecipesRequest) Descriptor() ([]byte, []int) {
	return file_recipes_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetIngredientsForAllRecipesRequest) GetRecipe() *Recipe {
	if x != nil {
		return x.Recipe
	}
	return nil
}

type GetIngredientsForAllRecipesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ingredient *Ingredient `protobuf:"bytes,1,opt,name=ingredient,proto3" json:"ingredient,omitempty"`
}

func (x *GetIngredientsForAllRecipesResponse) Reset() {
	*x = GetIngredientsForAllRecipesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recipes_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIngredientsForAllRecipesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIngredientsForAllRecipesResponse) ProtoMessage() {}

func (x *GetIngredientsForAllRecipesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_recipes_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIngredientsForAllRecipesResponse.ProtoReflect.Descriptor instead.
func (*GetIngredientsForAllRecipesResponse) Descriptor() ([]byte, []int) {
	return file_recipes_service_proto_rawDescGZIP(), []int{9}
}

func (x *GetIngredientsForAllRecipesResponse) GetIngredient() *Ingredient {
	if x != nil {
		return x.Ingredient
	}
	return nil
}

var File_recipes_service_proto protoreflect.FileDescriptor

var file_recipes_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67,
	0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x10,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x27, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x52, 0x06, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x22, 0x2d, 0x0a, 0x11, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x17, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x41, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x63, 0x69,
	0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x06, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x22, 0x36, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x69, 0x73, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x75, 0x69, 0x73, 0x69, 0x6e, 0x65, 0x22, 0x3c, 0x0a, 0x0a,
	0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x56, 0x0a, 0x1f, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x41, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a,
	0x0a, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x67, 0x72,
	0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65,
	0x6e, 0x74, 0x22, 0x3c, 0x0a, 0x20, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x67,
	0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x41, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x4d, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x46, 0x6f, 0x72, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73,
	0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x06, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x22,
	0x5a, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x46, 0x6f, 0x72, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64,
	0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x0a, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x32, 0xc0, 0x03, 0x0a, 0x0e,
	0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68,
	0x0a, 0x09, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x12, 0x19, 0x2e, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73,
	0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f,
	0x68, 0x6f, 0x6d, 0x65, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x53, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x63, 0x69,
	0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x63, 0x69,
	0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x71, 0x0a,
	0x18, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x41, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x28, 0x2e, 0x72, 0x65, 0x63, 0x69,
	0x70, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x67, 0x72, 0x65,
	0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x41, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x41, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01,
	0x12, 0x7c, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x46, 0x6f, 0x72, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x12,
	0x2b, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x67,
	0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x6f, 0x72, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x6f, 0x72, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x29,
	0x5a, 0x11, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x92, 0x41, 0x13, 0x12, 0x0e, 0x0a, 0x0c, 0x68, 0x6f, 0x6d, 0x65, 0x2d, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x2a, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_recipes_service_proto_rawDescOnce sync.Once
	file_recipes_service_proto_rawDescData = file_recipes_service_proto_rawDesc
)

func file_recipes_service_proto_rawDescGZIP() []byte {
	file_recipes_service_proto_rawDescOnce.Do(func() {
		file_recipes_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_recipes_service_proto_rawDescData)
	})
	return file_recipes_service_proto_rawDescData
}

var file_recipes_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_recipes_service_proto_goTypes = []interface{}{
	(*AddRecipeRequest)(nil),                    // 0: recipes.AddRecipeRequest
	(*AddRecipeResponse)(nil),                   // 1: recipes.AddRecipeResponse
	(*ListAllRecipesRequest)(nil),               // 2: recipes.ListAllRecipesRequest
	(*ListAllRecipesResponse)(nil),              // 3: recipes.ListAllRecipesResponse
	(*Recipe)(nil),                              // 4: recipes.Recipe
	(*Ingredient)(nil),                          // 5: recipes.Ingredient
	(*ListAllIngredientsAtHomeRequest)(nil),     // 6: recipes.ListAllIngredientsAtHomeRequest
	(*ListAllIngredientsAtHomeResponse)(nil),    // 7: recipes.ListAllIngredientsAtHomeResponse
	(*GetIngredientsForAllRecipesRequest)(nil),  // 8: recipes.GetIngredientsForAllRecipesRequest
	(*GetIngredientsForAllRecipesResponse)(nil), // 9: recipes.GetIngredientsForAllRecipesResponse
}
var file_recipes_service_proto_depIdxs = []int32{
	4, // 0: recipes.AddRecipeRequest.recipe:type_name -> recipes.Recipe
	4, // 1: recipes.ListAllRecipesResponse.recipe:type_name -> recipes.Recipe
	5, // 2: recipes.ListAllIngredientsAtHomeRequest.ingredient:type_name -> recipes.Ingredient
	4, // 3: recipes.GetIngredientsForAllRecipesRequest.recipe:type_name -> recipes.Recipe
	5, // 4: recipes.GetIngredientsForAllRecipesResponse.ingredient:type_name -> recipes.Ingredient
	0, // 5: recipes.RecipesService.AddRecipe:input_type -> recipes.AddRecipeRequest
	2, // 6: recipes.RecipesService.ListAllRecipes:input_type -> recipes.ListAllRecipesRequest
	6, // 7: recipes.RecipesService.ListAllIngredientsAtHome:input_type -> recipes.ListAllIngredientsAtHomeRequest
	8, // 8: recipes.RecipesService.GetIngredientsForAllRecipes:input_type -> recipes.GetIngredientsForAllRecipesRequest
	1, // 9: recipes.RecipesService.AddRecipe:output_type -> recipes.AddRecipeResponse
	3, // 10: recipes.RecipesService.ListAllRecipes:output_type -> recipes.ListAllRecipesResponse
	7, // 11: recipes.RecipesService.ListAllIngredientsAtHome:output_type -> recipes.ListAllIngredientsAtHomeResponse
	9, // 12: recipes.RecipesService.GetIngredientsForAllRecipes:output_type -> recipes.GetIngredientsForAllRecipesResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_recipes_service_proto_init() }
func file_recipes_service_proto_init() {
	if File_recipes_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_recipes_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRecipeRequest); i {
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
		file_recipes_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRecipeResponse); i {
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
		file_recipes_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAllRecipesRequest); i {
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
		file_recipes_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAllRecipesResponse); i {
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
		file_recipes_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Recipe); i {
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
		file_recipes_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ingredient); i {
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
		file_recipes_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAllIngredientsAtHomeRequest); i {
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
		file_recipes_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAllIngredientsAtHomeResponse); i {
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
		file_recipes_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIngredientsForAllRecipesRequest); i {
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
		file_recipes_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIngredientsForAllRecipesResponse); i {
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
			RawDescriptor: file_recipes_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_recipes_service_proto_goTypes,
		DependencyIndexes: file_recipes_service_proto_depIdxs,
		MessageInfos:      file_recipes_service_proto_msgTypes,
	}.Build()
	File_recipes_service_proto = out.File
	file_recipes_service_proto_rawDesc = nil
	file_recipes_service_proto_goTypes = nil
	file_recipes_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RecipesServiceClient is the client API for RecipesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecipesServiceClient interface {
	// unary
	AddRecipe(ctx context.Context, in *AddRecipeRequest, opts ...grpc.CallOption) (*AddRecipeResponse, error)
	// server-streaming
	ListAllRecipes(ctx context.Context, in *ListAllRecipesRequest, opts ...grpc.CallOption) (RecipesService_ListAllRecipesClient, error)
	// client-streaming
	ListAllIngredientsAtHome(ctx context.Context, opts ...grpc.CallOption) (RecipesService_ListAllIngredientsAtHomeClient, error)
	// Bidirectional-streaming
	GetIngredientsForAllRecipes(ctx context.Context, opts ...grpc.CallOption) (RecipesService_GetIngredientsForAllRecipesClient, error)
}

type recipesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecipesServiceClient(cc grpc.ClientConnInterface) RecipesServiceClient {
	return &recipesServiceClient{cc}
}

func (c *recipesServiceClient) AddRecipe(ctx context.Context, in *AddRecipeRequest, opts ...grpc.CallOption) (*AddRecipeResponse, error) {
	out := new(AddRecipeResponse)
	err := c.cc.Invoke(ctx, "/recipes.RecipesService/AddRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recipesServiceClient) ListAllRecipes(ctx context.Context, in *ListAllRecipesRequest, opts ...grpc.CallOption) (RecipesService_ListAllRecipesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RecipesService_serviceDesc.Streams[0], "/recipes.RecipesService/ListAllRecipes", opts...)
	if err != nil {
		return nil, err
	}
	x := &recipesServiceListAllRecipesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RecipesService_ListAllRecipesClient interface {
	Recv() (*ListAllRecipesResponse, error)
	grpc.ClientStream
}

type recipesServiceListAllRecipesClient struct {
	grpc.ClientStream
}

func (x *recipesServiceListAllRecipesClient) Recv() (*ListAllRecipesResponse, error) {
	m := new(ListAllRecipesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *recipesServiceClient) ListAllIngredientsAtHome(ctx context.Context, opts ...grpc.CallOption) (RecipesService_ListAllIngredientsAtHomeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RecipesService_serviceDesc.Streams[1], "/recipes.RecipesService/ListAllIngredientsAtHome", opts...)
	if err != nil {
		return nil, err
	}
	x := &recipesServiceListAllIngredientsAtHomeClient{stream}
	return x, nil
}

type RecipesService_ListAllIngredientsAtHomeClient interface {
	Send(*ListAllIngredientsAtHomeRequest) error
	CloseAndRecv() (*ListAllIngredientsAtHomeResponse, error)
	grpc.ClientStream
}

type recipesServiceListAllIngredientsAtHomeClient struct {
	grpc.ClientStream
}

func (x *recipesServiceListAllIngredientsAtHomeClient) Send(m *ListAllIngredientsAtHomeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *recipesServiceListAllIngredientsAtHomeClient) CloseAndRecv() (*ListAllIngredientsAtHomeResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ListAllIngredientsAtHomeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *recipesServiceClient) GetIngredientsForAllRecipes(ctx context.Context, opts ...grpc.CallOption) (RecipesService_GetIngredientsForAllRecipesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RecipesService_serviceDesc.Streams[2], "/recipes.RecipesService/GetIngredientsForAllRecipes", opts...)
	if err != nil {
		return nil, err
	}
	x := &recipesServiceGetIngredientsForAllRecipesClient{stream}
	return x, nil
}

type RecipesService_GetIngredientsForAllRecipesClient interface {
	Send(*GetIngredientsForAllRecipesRequest) error
	Recv() (*GetIngredientsForAllRecipesResponse, error)
	grpc.ClientStream
}

type recipesServiceGetIngredientsForAllRecipesClient struct {
	grpc.ClientStream
}

func (x *recipesServiceGetIngredientsForAllRecipesClient) Send(m *GetIngredientsForAllRecipesRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *recipesServiceGetIngredientsForAllRecipesClient) Recv() (*GetIngredientsForAllRecipesResponse, error) {
	m := new(GetIngredientsForAllRecipesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RecipesServiceServer is the server API for RecipesService service.
type RecipesServiceServer interface {
	// unary
	AddRecipe(context.Context, *AddRecipeRequest) (*AddRecipeResponse, error)
	// server-streaming
	ListAllRecipes(*ListAllRecipesRequest, RecipesService_ListAllRecipesServer) error
	// client-streaming
	ListAllIngredientsAtHome(RecipesService_ListAllIngredientsAtHomeServer) error
	// Bidirectional-streaming
	GetIngredientsForAllRecipes(RecipesService_GetIngredientsForAllRecipesServer) error
}

// UnimplementedRecipesServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRecipesServiceServer struct {
}

func (*UnimplementedRecipesServiceServer) AddRecipe(context.Context, *AddRecipeRequest) (*AddRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRecipe not implemented")
}
func (*UnimplementedRecipesServiceServer) ListAllRecipes(*ListAllRecipesRequest, RecipesService_ListAllRecipesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListAllRecipes not implemented")
}
func (*UnimplementedRecipesServiceServer) ListAllIngredientsAtHome(RecipesService_ListAllIngredientsAtHomeServer) error {
	return status.Errorf(codes.Unimplemented, "method ListAllIngredientsAtHome not implemented")
}
func (*UnimplementedRecipesServiceServer) GetIngredientsForAllRecipes(RecipesService_GetIngredientsForAllRecipesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetIngredientsForAllRecipes not implemented")
}

func RegisterRecipesServiceServer(s *grpc.Server, srv RecipesServiceServer) {
	s.RegisterService(&_RecipesService_serviceDesc, srv)
}

func _RecipesService_AddRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecipesServiceServer).AddRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recipes.RecipesService/AddRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecipesServiceServer).AddRecipe(ctx, req.(*AddRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecipesService_ListAllRecipes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListAllRecipesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RecipesServiceServer).ListAllRecipes(m, &recipesServiceListAllRecipesServer{stream})
}

type RecipesService_ListAllRecipesServer interface {
	Send(*ListAllRecipesResponse) error
	grpc.ServerStream
}

type recipesServiceListAllRecipesServer struct {
	grpc.ServerStream
}

func (x *recipesServiceListAllRecipesServer) Send(m *ListAllRecipesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _RecipesService_ListAllIngredientsAtHome_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RecipesServiceServer).ListAllIngredientsAtHome(&recipesServiceListAllIngredientsAtHomeServer{stream})
}

type RecipesService_ListAllIngredientsAtHomeServer interface {
	SendAndClose(*ListAllIngredientsAtHomeResponse) error
	Recv() (*ListAllIngredientsAtHomeRequest, error)
	grpc.ServerStream
}

type recipesServiceListAllIngredientsAtHomeServer struct {
	grpc.ServerStream
}

func (x *recipesServiceListAllIngredientsAtHomeServer) SendAndClose(m *ListAllIngredientsAtHomeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *recipesServiceListAllIngredientsAtHomeServer) Recv() (*ListAllIngredientsAtHomeRequest, error) {
	m := new(ListAllIngredientsAtHomeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RecipesService_GetIngredientsForAllRecipes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RecipesServiceServer).GetIngredientsForAllRecipes(&recipesServiceGetIngredientsForAllRecipesServer{stream})
}

type RecipesService_GetIngredientsForAllRecipesServer interface {
	Send(*GetIngredientsForAllRecipesResponse) error
	Recv() (*GetIngredientsForAllRecipesRequest, error)
	grpc.ServerStream
}

type recipesServiceGetIngredientsForAllRecipesServer struct {
	grpc.ServerStream
}

func (x *recipesServiceGetIngredientsForAllRecipesServer) Send(m *GetIngredientsForAllRecipesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *recipesServiceGetIngredientsForAllRecipesServer) Recv() (*GetIngredientsForAllRecipesRequest, error) {
	m := new(GetIngredientsForAllRecipesRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _RecipesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "recipes.RecipesService",
	HandlerType: (*RecipesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddRecipe",
			Handler:    _RecipesService_AddRecipe_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListAllRecipes",
			Handler:       _RecipesService_ListAllRecipes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListAllIngredientsAtHome",
			Handler:       _RecipesService_ListAllIngredientsAtHome_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetIngredientsForAllRecipes",
			Handler:       _RecipesService_GetIngredientsForAllRecipes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "recipes-service.proto",
}
