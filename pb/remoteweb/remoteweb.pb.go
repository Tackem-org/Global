// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: remoteweb.proto

package remoteweb

import (
	web "github.com/Tackem-org/Global/pb/web"
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

type PageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path            string    `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	BasePath        string    `protobuf:"bytes,2,opt,name=base_path,json=basePath,proto3" json:"base_path,omitempty"`
	User            *UserData `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Method          string    `protobuf:"bytes,4,opt,name=method,proto3" json:"method,omitempty"`
	QueryParamsJson string    `protobuf:"bytes,5,opt,name=query_params_json,json=queryParamsJson,proto3" json:"query_params_json,omitempty"`
	PostJson        string    `protobuf:"bytes,6,opt,name=post_json,json=postJson,proto3" json:"post_json,omitempty"`
	PathParamsJson  string    `protobuf:"bytes,7,opt,name=path_params_json,json=pathParamsJson,proto3" json:"path_params_json,omitempty"`
}

func (x *PageRequest) Reset() {
	*x = PageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remoteweb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageRequest) ProtoMessage() {}

func (x *PageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remoteweb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageRequest.ProtoReflect.Descriptor instead.
func (*PageRequest) Descriptor() ([]byte, []int) {
	return file_remoteweb_proto_rawDescGZIP(), []int{0}
}

func (x *PageRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *PageRequest) GetBasePath() string {
	if x != nil {
		return x.BasePath
	}
	return ""
}

func (x *PageRequest) GetUser() *UserData {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *PageRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *PageRequest) GetQueryParamsJson() string {
	if x != nil {
		return x.QueryParamsJson
	}
	return ""
}

func (x *PageRequest) GetPostJson() string {
	if x != nil {
		return x.PostJson
	}
	return ""
}

func (x *PageRequest) GetPathParamsJson() string {
	if x != nil {
		return x.PathParamsJson
	}
	return ""
}

type PageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode        uint32   `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	HideErrorFromUser bool     `protobuf:"varint,2,opt,name=hide_error_from_user,json=hideErrorFromUser,proto3" json:"hide_error_from_user,omitempty"`
	ErrorMessage      string   `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	RedirectUrl       string   `protobuf:"bytes,4,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	TemplateHtml      string   `protobuf:"bytes,5,opt,name=template_html,json=templateHtml,proto3" json:"template_html,omitempty"`
	PageVariablesJson string   `protobuf:"bytes,6,opt,name=page_variables_json,json=pageVariablesJson,proto3" json:"page_variables_json,omitempty"`
	CustomPageName    string   `protobuf:"bytes,7,opt,name=custom_page_name,json=customPageName,proto3" json:"custom_page_name,omitempty"`
	CustomCss         []string `protobuf:"bytes,8,rep,name=custom_css,json=customCss,proto3" json:"custom_css,omitempty"`
	CustomJs          []string `protobuf:"bytes,9,rep,name=custom_js,json=customJs,proto3" json:"custom_js,omitempty"`
}

func (x *PageResponse) Reset() {
	*x = PageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remoteweb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageResponse) ProtoMessage() {}

func (x *PageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_remoteweb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageResponse.ProtoReflect.Descriptor instead.
func (*PageResponse) Descriptor() ([]byte, []int) {
	return file_remoteweb_proto_rawDescGZIP(), []int{1}
}

func (x *PageResponse) GetStatusCode() uint32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *PageResponse) GetHideErrorFromUser() bool {
	if x != nil {
		return x.HideErrorFromUser
	}
	return false
}

func (x *PageResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *PageResponse) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

func (x *PageResponse) GetTemplateHtml() string {
	if x != nil {
		return x.TemplateHtml
	}
	return ""
}

func (x *PageResponse) GetPageVariablesJson() string {
	if x != nil {
		return x.PageVariablesJson
	}
	return ""
}

func (x *PageResponse) GetCustomPageName() string {
	if x != nil {
		return x.CustomPageName
	}
	return ""
}

func (x *PageResponse) GetCustomCss() []string {
	if x != nil {
		return x.CustomCss
	}
	return nil
}

func (x *PageResponse) GetCustomJs() []string {
	if x != nil {
		return x.CustomJs
	}
	return nil
}

type FileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *FileRequest) Reset() {
	*x = FileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remoteweb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRequest) ProtoMessage() {}

func (x *FileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remoteweb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRequest.ProtoReflect.Descriptor instead.
func (*FileRequest) Descriptor() ([]byte, []int) {
	return file_remoteweb_proto_rawDescGZIP(), []int{2}
}

func (x *FileRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type FileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode   uint32 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	File         []byte `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *FileResponse) Reset() {
	*x = FileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remoteweb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileResponse) ProtoMessage() {}

func (x *FileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_remoteweb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileResponse.ProtoReflect.Descriptor instead.
func (*FileResponse) Descriptor() ([]byte, []int) {
	return file_remoteweb_proto_rawDescGZIP(), []int{3}
}

func (x *FileResponse) GetStatusCode() uint32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *FileResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *FileResponse) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

type WebSocketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command  string    `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	User     *UserData `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	DataJson string    `protobuf:"bytes,3,opt,name=data_json,json=dataJson,proto3" json:"data_json,omitempty"`
}

func (x *WebSocketRequest) Reset() {
	*x = WebSocketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remoteweb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebSocketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebSocketRequest) ProtoMessage() {}

func (x *WebSocketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remoteweb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebSocketRequest.ProtoReflect.Descriptor instead.
func (*WebSocketRequest) Descriptor() ([]byte, []int) {
	return file_remoteweb_proto_rawDescGZIP(), []int{4}
}

func (x *WebSocketRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *WebSocketRequest) GetUser() *UserData {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *WebSocketRequest) GetDataJson() string {
	if x != nil {
		return x.DataJson
	}
	return ""
}

type WebSocketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode        uint32 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	HideErrorFromUser bool   `protobuf:"varint,2,opt,name=hide_error_from_user,json=hideErrorFromUser,proto3" json:"hide_error_from_user,omitempty"`
	ErrorMessage      string `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	TellAll           bool   `protobuf:"varint,4,opt,name=tell_all,json=tellAll,proto3" json:"tell_all,omitempty"`
	DataJson          string `protobuf:"bytes,5,opt,name=data_json,json=dataJson,proto3" json:"data_json,omitempty"`
}

func (x *WebSocketResponse) Reset() {
	*x = WebSocketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remoteweb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebSocketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebSocketResponse) ProtoMessage() {}

func (x *WebSocketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_remoteweb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebSocketResponse.ProtoReflect.Descriptor instead.
func (*WebSocketResponse) Descriptor() ([]byte, []int) {
	return file_remoteweb_proto_rawDescGZIP(), []int{5}
}

func (x *WebSocketResponse) GetStatusCode() uint32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *WebSocketResponse) GetHideErrorFromUser() bool {
	if x != nil {
		return x.HideErrorFromUser
	}
	return false
}

func (x *WebSocketResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *WebSocketResponse) GetTellAll() bool {
	if x != nil {
		return x.TellAll
	}
	return false
}

func (x *WebSocketResponse) GetDataJson() string {
	if x != nil {
		return x.DataJson
	}
	return ""
}

type UserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      uint64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Icon        string   `protobuf:"bytes,3,opt,name=Icon,proto3" json:"Icon,omitempty"`
	IsAdmin     bool     `protobuf:"varint,4,opt,name=IsAdmin,proto3" json:"IsAdmin,omitempty"`
	Permissions []string `protobuf:"bytes,5,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *UserData) Reset() {
	*x = UserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remoteweb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_remoteweb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_remoteweb_proto_rawDescGZIP(), []int{6}
}

func (x *UserData) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserData) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *UserData) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

func (x *UserData) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type TasksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TasksRequest) Reset() {
	*x = TasksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remoteweb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TasksRequest) ProtoMessage() {}

func (x *TasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remoteweb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TasksRequest.ProtoReflect.Descriptor instead.
func (*TasksRequest) Descriptor() ([]byte, []int) {
	return file_remoteweb_proto_rawDescGZIP(), []int{7}
}

type TasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success      bool               `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ErrorMessage string             `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	Tasks        []*web.TaskMessage `protobuf:"bytes,3,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *TasksResponse) Reset() {
	*x = TasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remoteweb_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TasksResponse) ProtoMessage() {}

func (x *TasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_remoteweb_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TasksResponse.ProtoReflect.Descriptor instead.
func (*TasksResponse) Descriptor() ([]byte, []int) {
	return file_remoteweb_proto_rawDescGZIP(), []int{8}
}

func (x *TasksResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *TasksResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *TasksResponse) GetTasks() []*web.TaskMessage {
	if x != nil {
		return x.Tasks
	}
	return nil
}

var File_remoteweb_proto protoreflect.FileDescriptor

var file_remoteweb_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x77, 0x65, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x77, 0x65, 0x62, 0x1a, 0x09, 0x77, 0x65,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2, 0x01, 0x0a, 0x0b, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x62, 0x61, 0x73, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x27, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x77,
	0x65, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x6a, 0x73,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x4a, 0x73,
	0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x61,
	0x74, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4a, 0x73, 0x6f, 0x6e, 0x22, 0xe3, 0x02, 0x0a,
	0x0c, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2f,
	0x0a, 0x14, 0x68, 0x69, 0x64, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x66, 0x72, 0x6f,
	0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x68, 0x69,
	0x64, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x5f, 0x68, 0x74, 0x6d, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x48, 0x74, 0x6d, 0x6c, 0x12, 0x2e, 0x0a, 0x13,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x5f, 0x6a,
	0x73, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x61, 0x67, 0x65, 0x56,
	0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x61,
	0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x5f, 0x63, 0x73, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x43, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f,
	0x6a, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x4a, 0x73, 0x22, 0x21, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x68, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22,
	0x72, 0x0a, 0x10, 0x57, 0x65, 0x62, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x27, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x77, 0x65, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6a,
	0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x4a,
	0x73, 0x6f, 0x6e, 0x22, 0xc2, 0x01, 0x0a, 0x11, 0x57, 0x65, 0x62, 0x53, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2f, 0x0a, 0x14, 0x68, 0x69,
	0x64, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x68, 0x69, 0x64, 0x65, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x65, 0x6c, 0x6c, 0x5f, 0x61, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x74, 0x65, 0x6c, 0x6c, 0x41, 0x6c, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x4a, 0x73, 0x6f, 0x6e, 0x22, 0x87, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x76, 0x0a, 0x0d, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a,
	0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x32, 0xc9, 0x02, 0x0a, 0x09, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x57, 0x65, 0x62, 0x12, 0x39, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65,
	0x12, 0x16, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x77, 0x65, 0x62, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x77, 0x65, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x09, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x61, 0x67, 0x65,
	0x12, 0x16, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x77, 0x65, 0x62, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x77, 0x65, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x77, 0x65, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x77, 0x65, 0x62, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48,
	0x0a, 0x09, 0x57, 0x65, 0x62, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x77, 0x65, 0x62, 0x2e, 0x57, 0x65, 0x62, 0x53, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x77, 0x65, 0x62, 0x2e, 0x57, 0x65, 0x62, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x05, 0x54, 0x61, 0x73, 0x6b,
	0x73, 0x12, 0x17, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x77, 0x65, 0x62, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x77, 0x65, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x61, 0x63, 0x6b, 0x65, 0x6d, 0x2d, 0x6f, 0x72, 0x67, 0x2f,
	0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x77, 0x65, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_remoteweb_proto_rawDescOnce sync.Once
	file_remoteweb_proto_rawDescData = file_remoteweb_proto_rawDesc
)

func file_remoteweb_proto_rawDescGZIP() []byte {
	file_remoteweb_proto_rawDescOnce.Do(func() {
		file_remoteweb_proto_rawDescData = protoimpl.X.CompressGZIP(file_remoteweb_proto_rawDescData)
	})
	return file_remoteweb_proto_rawDescData
}

var file_remoteweb_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_remoteweb_proto_goTypes = []interface{}{
	(*PageRequest)(nil),       // 0: remoteweb.PageRequest
	(*PageResponse)(nil),      // 1: remoteweb.PageResponse
	(*FileRequest)(nil),       // 2: remoteweb.FileRequest
	(*FileResponse)(nil),      // 3: remoteweb.FileResponse
	(*WebSocketRequest)(nil),  // 4: remoteweb.WebSocketRequest
	(*WebSocketResponse)(nil), // 5: remoteweb.WebSocketResponse
	(*UserData)(nil),          // 6: remoteweb.UserData
	(*TasksRequest)(nil),      // 7: remoteweb.TasksRequest
	(*TasksResponse)(nil),     // 8: remoteweb.TasksResponse
	(*web.TaskMessage)(nil),   // 9: web.TaskMessage
}
var file_remoteweb_proto_depIdxs = []int32{
	6, // 0: remoteweb.PageRequest.user:type_name -> remoteweb.UserData
	6, // 1: remoteweb.WebSocketRequest.user:type_name -> remoteweb.UserData
	9, // 2: remoteweb.TasksResponse.tasks:type_name -> web.TaskMessage
	0, // 3: remoteweb.RemoteWeb.Page:input_type -> remoteweb.PageRequest
	0, // 4: remoteweb.RemoteWeb.AdminPage:input_type -> remoteweb.PageRequest
	2, // 5: remoteweb.RemoteWeb.File:input_type -> remoteweb.FileRequest
	4, // 6: remoteweb.RemoteWeb.WebSocket:input_type -> remoteweb.WebSocketRequest
	7, // 7: remoteweb.RemoteWeb.Tasks:input_type -> remoteweb.TasksRequest
	1, // 8: remoteweb.RemoteWeb.Page:output_type -> remoteweb.PageResponse
	1, // 9: remoteweb.RemoteWeb.AdminPage:output_type -> remoteweb.PageResponse
	3, // 10: remoteweb.RemoteWeb.File:output_type -> remoteweb.FileResponse
	5, // 11: remoteweb.RemoteWeb.WebSocket:output_type -> remoteweb.WebSocketResponse
	8, // 12: remoteweb.RemoteWeb.Tasks:output_type -> remoteweb.TasksResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_remoteweb_proto_init() }
func file_remoteweb_proto_init() {
	if File_remoteweb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_remoteweb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageRequest); i {
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
		file_remoteweb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageResponse); i {
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
		file_remoteweb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRequest); i {
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
		file_remoteweb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileResponse); i {
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
		file_remoteweb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebSocketRequest); i {
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
		file_remoteweb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebSocketResponse); i {
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
		file_remoteweb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserData); i {
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
		file_remoteweb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TasksRequest); i {
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
		file_remoteweb_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TasksResponse); i {
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
			RawDescriptor: file_remoteweb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_remoteweb_proto_goTypes,
		DependencyIndexes: file_remoteweb_proto_depIdxs,
		MessageInfos:      file_remoteweb_proto_msgTypes,
	}.Build()
	File_remoteweb_proto = out.File
	file_remoteweb_proto_rawDesc = nil
	file_remoteweb_proto_goTypes = nil
	file_remoteweb_proto_depIdxs = nil
}