// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo_list_activity_manager.proto

package todolist

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Tags int32

const (
	Tags_PRIORITY    Tags = 0
	Tags_NORMAL      Tags = 1
	Tags_MONEY       Tags = 2
	Tags_SPORTS      Tags = 3
	Tags_COKKING     Tags = 4
	Tags_SOCAILIZING Tags = 5
	Tags_GENERAL     Tags = 6
)

var Tags_name = map[int32]string{
	0: "PRIORITY",
	1: "NORMAL",
	2: "MONEY",
	3: "SPORTS",
	4: "COKKING",
	5: "SOCAILIZING",
	6: "GENERAL",
}

var Tags_value = map[string]int32{
	"PRIORITY":    0,
	"NORMAL":      1,
	"MONEY":       2,
	"SPORTS":      3,
	"COKKING":     4,
	"SOCAILIZING": 5,
	"GENERAL":     6,
}

func (x Tags) String() string {
	return proto.EnumName(Tags_name, int32(x))
}

func (Tags) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_516e47f255818132, []int{0}
}

type CreateTodoListRequest struct {
	UserID               string        `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Name                 string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string        `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	DueDate              string        `protobuf:"bytes,4,opt,name=dueDate,proto3" json:"dueDate,omitempty"`
	CustomDueDate        string        `protobuf:"bytes,5,opt,name=customDueDate,proto3" json:"customDueDate,omitempty"`
	RepeatMode           bool          `protobuf:"varint,6,opt,name=repeatMode,proto3" json:"repeatMode,omitempty"`
	Tag                  Tags          `protobuf:"varint,7,opt,name=tag,proto3,enum=todolist.Tags" json:"tag,omitempty"`
	Notification         *Notification `protobuf:"bytes,8,opt,name=notification,proto3" json:"notification,omitempty"`
	SubTask              []*SubTask    `protobuf:"bytes,9,rep,name=subTask,proto3" json:"subTask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateTodoListRequest) Reset()         { *m = CreateTodoListRequest{} }
func (m *CreateTodoListRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTodoListRequest) ProtoMessage()    {}
func (*CreateTodoListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_516e47f255818132, []int{0}
}

func (m *CreateTodoListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoListRequest.Unmarshal(m, b)
}
func (m *CreateTodoListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoListRequest.Marshal(b, m, deterministic)
}
func (m *CreateTodoListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoListRequest.Merge(m, src)
}
func (m *CreateTodoListRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTodoListRequest.Size(m)
}
func (m *CreateTodoListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoListRequest proto.InternalMessageInfo

func (m *CreateTodoListRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *CreateTodoListRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateTodoListRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateTodoListRequest) GetDueDate() string {
	if m != nil {
		return m.DueDate
	}
	return ""
}

func (m *CreateTodoListRequest) GetCustomDueDate() string {
	if m != nil {
		return m.CustomDueDate
	}
	return ""
}

func (m *CreateTodoListRequest) GetRepeatMode() bool {
	if m != nil {
		return m.RepeatMode
	}
	return false
}

func (m *CreateTodoListRequest) GetTag() Tags {
	if m != nil {
		return m.Tag
	}
	return Tags_PRIORITY
}

func (m *CreateTodoListRequest) GetNotification() *Notification {
	if m != nil {
		return m.Notification
	}
	return nil
}

func (m *CreateTodoListRequest) GetSubTask() []*SubTask {
	if m != nil {
		return m.SubTask
	}
	return nil
}

type CreateTodoListResponse struct {
	Data                 []*CreateTodoListResponse_CreateTodoListUserInformationResponse `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Errors               []*CreateTodoListResponse_ErrorResponse                         `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                        `json:"-"`
	XXX_unrecognized     []byte                                                          `json:"-"`
	XXX_sizecache        int32                                                           `json:"-"`
}

func (m *CreateTodoListResponse) Reset()         { *m = CreateTodoListResponse{} }
func (m *CreateTodoListResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTodoListResponse) ProtoMessage()    {}
func (*CreateTodoListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_516e47f255818132, []int{1}
}

func (m *CreateTodoListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoListResponse.Unmarshal(m, b)
}
func (m *CreateTodoListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoListResponse.Marshal(b, m, deterministic)
}
func (m *CreateTodoListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoListResponse.Merge(m, src)
}
func (m *CreateTodoListResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTodoListResponse.Size(m)
}
func (m *CreateTodoListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoListResponse proto.InternalMessageInfo

func (m *CreateTodoListResponse) GetData() []*CreateTodoListResponse_CreateTodoListUserInformationResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *CreateTodoListResponse) GetErrors() []*CreateTodoListResponse_ErrorResponse {
	if m != nil {
		return m.Errors
	}
	return nil
}

type CreateTodoListResponse_CreateTodoListUserInformationResponse struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoListResponse_CreateTodoListUserInformationResponse) Reset() {
	*m = CreateTodoListResponse_CreateTodoListUserInformationResponse{}
}
func (m *CreateTodoListResponse_CreateTodoListUserInformationResponse) String() string {
	return proto.CompactTextString(m)
}
func (*CreateTodoListResponse_CreateTodoListUserInformationResponse) ProtoMessage() {}
func (*CreateTodoListResponse_CreateTodoListUserInformationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_516e47f255818132, []int{1, 0}
}

func (m *CreateTodoListResponse_CreateTodoListUserInformationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoListResponse_CreateTodoListUserInformationResponse.Unmarshal(m, b)
}
func (m *CreateTodoListResponse_CreateTodoListUserInformationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoListResponse_CreateTodoListUserInformationResponse.Marshal(b, m, deterministic)
}
func (m *CreateTodoListResponse_CreateTodoListUserInformationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoListResponse_CreateTodoListUserInformationResponse.Merge(m, src)
}
func (m *CreateTodoListResponse_CreateTodoListUserInformationResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTodoListResponse_CreateTodoListUserInformationResponse.Size(m)
}
func (m *CreateTodoListResponse_CreateTodoListUserInformationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoListResponse_CreateTodoListUserInformationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoListResponse_CreateTodoListUserInformationResponse proto.InternalMessageInfo

func (m *CreateTodoListResponse_CreateTodoListUserInformationResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type CreateTodoListResponse_ErrorResponse struct {
	ErrorMessage         string   `protobuf:"bytes,1,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	StatusCode           string   `protobuf:"bytes,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoListResponse_ErrorResponse) Reset()         { *m = CreateTodoListResponse_ErrorResponse{} }
func (m *CreateTodoListResponse_ErrorResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTodoListResponse_ErrorResponse) ProtoMessage()    {}
func (*CreateTodoListResponse_ErrorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_516e47f255818132, []int{1, 1}
}

func (m *CreateTodoListResponse_ErrorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoListResponse_ErrorResponse.Unmarshal(m, b)
}
func (m *CreateTodoListResponse_ErrorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoListResponse_ErrorResponse.Marshal(b, m, deterministic)
}
func (m *CreateTodoListResponse_ErrorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoListResponse_ErrorResponse.Merge(m, src)
}
func (m *CreateTodoListResponse_ErrorResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTodoListResponse_ErrorResponse.Size(m)
}
func (m *CreateTodoListResponse_ErrorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoListResponse_ErrorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoListResponse_ErrorResponse proto.InternalMessageInfo

func (m *CreateTodoListResponse_ErrorResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *CreateTodoListResponse_ErrorResponse) GetStatusCode() string {
	if m != nil {
		return m.StatusCode
	}
	return ""
}

type SubTask struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Status               bool     `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Offset               int32    `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubTask) Reset()         { *m = SubTask{} }
func (m *SubTask) String() string { return proto.CompactTextString(m) }
func (*SubTask) ProtoMessage()    {}
func (*SubTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_516e47f255818132, []int{2}
}

func (m *SubTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubTask.Unmarshal(m, b)
}
func (m *SubTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubTask.Marshal(b, m, deterministic)
}
func (m *SubTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubTask.Merge(m, src)
}
func (m *SubTask) XXX_Size() int {
	return xxx_messageInfo_SubTask.Size(m)
}
func (m *SubTask) XXX_DiscardUnknown() {
	xxx_messageInfo_SubTask.DiscardUnknown(m)
}

var xxx_messageInfo_SubTask proto.InternalMessageInfo

func (m *SubTask) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SubTask) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SubTask) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *SubTask) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type Notification struct {
	Email                *Email       `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Message              *Textmessage `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_516e47f255818132, []int{3}
}

func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetEmail() *Email {
	if m != nil {
		return m.Email
	}
	return nil
}

func (m *Notification) GetMessage() *Textmessage {
	if m != nil {
		return m.Message
	}
	return nil
}

type Email struct {
	Notification         bool     `protobuf:"varint,1,opt,name=notification,proto3" json:"notification,omitempty"`
	EmailID              string   `protobuf:"bytes,2,opt,name=emailID,proto3" json:"emailID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Email) Reset()         { *m = Email{} }
func (m *Email) String() string { return proto.CompactTextString(m) }
func (*Email) ProtoMessage()    {}
func (*Email) Descriptor() ([]byte, []int) {
	return fileDescriptor_516e47f255818132, []int{4}
}

func (m *Email) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Email.Unmarshal(m, b)
}
func (m *Email) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Email.Marshal(b, m, deterministic)
}
func (m *Email) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Email.Merge(m, src)
}
func (m *Email) XXX_Size() int {
	return xxx_messageInfo_Email.Size(m)
}
func (m *Email) XXX_DiscardUnknown() {
	xxx_messageInfo_Email.DiscardUnknown(m)
}

var xxx_messageInfo_Email proto.InternalMessageInfo

func (m *Email) GetNotification() bool {
	if m != nil {
		return m.Notification
	}
	return false
}

func (m *Email) GetEmailID() string {
	if m != nil {
		return m.EmailID
	}
	return ""
}

type Textmessage struct {
	Notification         bool     `protobuf:"varint,1,opt,name=notification,proto3" json:"notification,omitempty"`
	CountryCode          string   `protobuf:"bytes,2,opt,name=countryCode,proto3" json:"countryCode,omitempty"`
	MobileNumber         string   `protobuf:"bytes,3,opt,name=mobileNumber,proto3" json:"mobileNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Textmessage) Reset()         { *m = Textmessage{} }
func (m *Textmessage) String() string { return proto.CompactTextString(m) }
func (*Textmessage) ProtoMessage()    {}
func (*Textmessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_516e47f255818132, []int{5}
}

func (m *Textmessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Textmessage.Unmarshal(m, b)
}
func (m *Textmessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Textmessage.Marshal(b, m, deterministic)
}
func (m *Textmessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Textmessage.Merge(m, src)
}
func (m *Textmessage) XXX_Size() int {
	return xxx_messageInfo_Textmessage.Size(m)
}
func (m *Textmessage) XXX_DiscardUnknown() {
	xxx_messageInfo_Textmessage.DiscardUnknown(m)
}

var xxx_messageInfo_Textmessage proto.InternalMessageInfo

func (m *Textmessage) GetNotification() bool {
	if m != nil {
		return m.Notification
	}
	return false
}

func (m *Textmessage) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

func (m *Textmessage) GetMobileNumber() string {
	if m != nil {
		return m.MobileNumber
	}
	return ""
}

type UpdateTodoListRequest struct {
	UserID               string        `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Name                 string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string        `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	DueDate              string        `protobuf:"bytes,4,opt,name=dueDate,proto3" json:"dueDate,omitempty"`
	CustomDueDate        string        `protobuf:"bytes,5,opt,name=customDueDate,proto3" json:"customDueDate,omitempty"`
	RepeatMode           bool          `protobuf:"varint,6,opt,name=repeatMode,proto3" json:"repeatMode,omitempty"`
	Tag                  Tags          `protobuf:"varint,7,opt,name=tag,proto3,enum=todolist.Tags" json:"tag,omitempty"`
	Notification         *Notification `protobuf:"bytes,8,opt,name=notification,proto3" json:"notification,omitempty"`
	SubTask              []*SubTask    `protobuf:"bytes,9,rep,name=subTask,proto3" json:"subTask,omitempty"`
	ActivityID           string        `protobuf:"bytes,10,opt,name=activityID,proto3" json:"activityID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UpdateTodoListRequest) Reset()         { *m = UpdateTodoListRequest{} }
func (m *UpdateTodoListRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTodoListRequest) ProtoMessage()    {}
func (*UpdateTodoListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_516e47f255818132, []int{6}
}

func (m *UpdateTodoListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTodoListRequest.Unmarshal(m, b)
}
func (m *UpdateTodoListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTodoListRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTodoListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTodoListRequest.Merge(m, src)
}
func (m *UpdateTodoListRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTodoListRequest.Size(m)
}
func (m *UpdateTodoListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTodoListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTodoListRequest proto.InternalMessageInfo

func (m *UpdateTodoListRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UpdateTodoListRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateTodoListRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateTodoListRequest) GetDueDate() string {
	if m != nil {
		return m.DueDate
	}
	return ""
}

func (m *UpdateTodoListRequest) GetCustomDueDate() string {
	if m != nil {
		return m.CustomDueDate
	}
	return ""
}

func (m *UpdateTodoListRequest) GetRepeatMode() bool {
	if m != nil {
		return m.RepeatMode
	}
	return false
}

func (m *UpdateTodoListRequest) GetTag() Tags {
	if m != nil {
		return m.Tag
	}
	return Tags_PRIORITY
}

func (m *UpdateTodoListRequest) GetNotification() *Notification {
	if m != nil {
		return m.Notification
	}
	return nil
}

func (m *UpdateTodoListRequest) GetSubTask() []*SubTask {
	if m != nil {
		return m.SubTask
	}
	return nil
}

func (m *UpdateTodoListRequest) GetActivityID() string {
	if m != nil {
		return m.ActivityID
	}
	return ""
}

type UpdateTodoListResponse struct {
	Data                 []*UpdateTodoListResponse_UpdateTodoListUserInformationResponse `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Errors               []*UpdateTodoListResponse_ErrorResponse                         `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                        `json:"-"`
	XXX_unrecognized     []byte                                                          `json:"-"`
	XXX_sizecache        int32                                                           `json:"-"`
}

func (m *UpdateTodoListResponse) Reset()         { *m = UpdateTodoListResponse{} }
func (m *UpdateTodoListResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateTodoListResponse) ProtoMessage()    {}
func (*UpdateTodoListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_516e47f255818132, []int{7}
}

func (m *UpdateTodoListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTodoListResponse.Unmarshal(m, b)
}
func (m *UpdateTodoListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTodoListResponse.Marshal(b, m, deterministic)
}
func (m *UpdateTodoListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTodoListResponse.Merge(m, src)
}
func (m *UpdateTodoListResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateTodoListResponse.Size(m)
}
func (m *UpdateTodoListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTodoListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTodoListResponse proto.InternalMessageInfo

func (m *UpdateTodoListResponse) GetData() []*UpdateTodoListResponse_UpdateTodoListUserInformationResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *UpdateTodoListResponse) GetErrors() []*UpdateTodoListResponse_ErrorResponse {
	if m != nil {
		return m.Errors
	}
	return nil
}

type UpdateTodoListResponse_UpdateTodoListUserInformationResponse struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTodoListResponse_UpdateTodoListUserInformationResponse) Reset() {
	*m = UpdateTodoListResponse_UpdateTodoListUserInformationResponse{}
}
func (m *UpdateTodoListResponse_UpdateTodoListUserInformationResponse) String() string {
	return proto.CompactTextString(m)
}
func (*UpdateTodoListResponse_UpdateTodoListUserInformationResponse) ProtoMessage() {}
func (*UpdateTodoListResponse_UpdateTodoListUserInformationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_516e47f255818132, []int{7, 0}
}

func (m *UpdateTodoListResponse_UpdateTodoListUserInformationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTodoListResponse_UpdateTodoListUserInformationResponse.Unmarshal(m, b)
}
func (m *UpdateTodoListResponse_UpdateTodoListUserInformationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTodoListResponse_UpdateTodoListUserInformationResponse.Marshal(b, m, deterministic)
}
func (m *UpdateTodoListResponse_UpdateTodoListUserInformationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTodoListResponse_UpdateTodoListUserInformationResponse.Merge(m, src)
}
func (m *UpdateTodoListResponse_UpdateTodoListUserInformationResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateTodoListResponse_UpdateTodoListUserInformationResponse.Size(m)
}
func (m *UpdateTodoListResponse_UpdateTodoListUserInformationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTodoListResponse_UpdateTodoListUserInformationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTodoListResponse_UpdateTodoListUserInformationResponse proto.InternalMessageInfo

func (m *UpdateTodoListResponse_UpdateTodoListUserInformationResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type UpdateTodoListResponse_ErrorResponse struct {
	ErrorMessage         string   `protobuf:"bytes,1,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	StatusCode           string   `protobuf:"bytes,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTodoListResponse_ErrorResponse) Reset()         { *m = UpdateTodoListResponse_ErrorResponse{} }
func (m *UpdateTodoListResponse_ErrorResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateTodoListResponse_ErrorResponse) ProtoMessage()    {}
func (*UpdateTodoListResponse_ErrorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_516e47f255818132, []int{7, 1}
}

func (m *UpdateTodoListResponse_ErrorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTodoListResponse_ErrorResponse.Unmarshal(m, b)
}
func (m *UpdateTodoListResponse_ErrorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTodoListResponse_ErrorResponse.Marshal(b, m, deterministic)
}
func (m *UpdateTodoListResponse_ErrorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTodoListResponse_ErrorResponse.Merge(m, src)
}
func (m *UpdateTodoListResponse_ErrorResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateTodoListResponse_ErrorResponse.Size(m)
}
func (m *UpdateTodoListResponse_ErrorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTodoListResponse_ErrorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTodoListResponse_ErrorResponse proto.InternalMessageInfo

func (m *UpdateTodoListResponse_ErrorResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *UpdateTodoListResponse_ErrorResponse) GetStatusCode() string {
	if m != nil {
		return m.StatusCode
	}
	return ""
}

func init() {
	proto.RegisterEnum("todolist.Tags", Tags_name, Tags_value)
	proto.RegisterType((*CreateTodoListRequest)(nil), "todolist.CreateTodoListRequest")
	proto.RegisterType((*CreateTodoListResponse)(nil), "todolist.CreateTodoListResponse")
	proto.RegisterType((*CreateTodoListResponse_CreateTodoListUserInformationResponse)(nil), "todolist.CreateTodoListResponse.CreateTodoListUserInformationResponse")
	proto.RegisterType((*CreateTodoListResponse_ErrorResponse)(nil), "todolist.CreateTodoListResponse.ErrorResponse")
	proto.RegisterType((*SubTask)(nil), "todolist.subTask")
	proto.RegisterType((*Notification)(nil), "todolist.notification")
	proto.RegisterType((*Email)(nil), "todolist.email")
	proto.RegisterType((*Textmessage)(nil), "todolist.textmessage")
	proto.RegisterType((*UpdateTodoListRequest)(nil), "todolist.UpdateTodoListRequest")
	proto.RegisterType((*UpdateTodoListResponse)(nil), "todolist.UpdateTodoListResponse")
	proto.RegisterType((*UpdateTodoListResponse_UpdateTodoListUserInformationResponse)(nil), "todolist.UpdateTodoListResponse.UpdateTodoListUserInformationResponse")
	proto.RegisterType((*UpdateTodoListResponse_ErrorResponse)(nil), "todolist.UpdateTodoListResponse.ErrorResponse")
}

func init() { proto.RegisterFile("todo_list_activity_manager.proto", fileDescriptor_516e47f255818132) }

var fileDescriptor_516e47f255818132 = []byte{
	// 687 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x55, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0x5e, 0xd2, 0xcf, 0xbd, 0xd9, 0x47, 0x31, 0x5a, 0x15, 0xed, 0x00, 0x51, 0xc4, 0xa4, 0x0a,
	0xa4, 0x22, 0x95, 0x03, 0x12, 0xb7, 0xd2, 0x76, 0x53, 0xb4, 0x7e, 0x0c, 0xb7, 0x3b, 0x6c, 0x97,
	0xc9, 0x6d, 0xdd, 0x12, 0xb1, 0xc4, 0x25, 0x76, 0x80, 0xfd, 0x45, 0x4e, 0xfc, 0x14, 0x7e, 0x02,
	0xb2, 0xeb, 0x36, 0x49, 0x59, 0xb5, 0x09, 0x21, 0x4e, 0xdc, 0xe2, 0xe7, 0x79, 0xf2, 0xd8, 0x7e,
	0xfd, 0xfa, 0x31, 0x38, 0x82, 0x4d, 0xd9, 0xcd, 0xad, 0xcf, 0xc5, 0x0d, 0x99, 0x08, 0xff, 0x8b,
	0x2f, 0xee, 0x6e, 0x02, 0x12, 0x92, 0x39, 0x8d, 0xea, 0x8b, 0x88, 0x09, 0x86, 0xca, 0x52, 0x21,
	0x05, 0xee, 0x0f, 0x13, 0x8e, 0x5a, 0x11, 0x25, 0x82, 0x8e, 0xd8, 0x94, 0x75, 0x7d, 0x2e, 0x30,
	0xfd, 0x1c, 0x53, 0x2e, 0x50, 0x15, 0x8a, 0x31, 0xa7, 0x91, 0xd7, 0xb6, 0x0d, 0xc7, 0xa8, 0xed,
	0x62, 0x3d, 0x42, 0x08, 0xf2, 0x21, 0x09, 0xa8, 0x6d, 0x2a, 0x54, 0x7d, 0x23, 0x07, 0xac, 0x29,
	0xe5, 0x93, 0xc8, 0x5f, 0x08, 0x9f, 0x85, 0x76, 0x4e, 0x51, 0x69, 0x08, 0xd9, 0x50, 0x9a, 0xc6,
	0xb4, 0x4d, 0x04, 0xb5, 0xf3, 0x8a, 0x5d, 0x0d, 0xd1, 0x0b, 0xd8, 0x9f, 0xc4, 0x5c, 0xb0, 0xa0,
	0xad, 0xf9, 0x82, 0xe2, 0xb3, 0x20, 0x7a, 0x06, 0x10, 0xd1, 0x05, 0x25, 0xa2, 0xc7, 0xa6, 0xd4,
	0x2e, 0x3a, 0x46, 0xad, 0x8c, 0x53, 0x08, 0x72, 0x20, 0x27, 0xc8, 0xdc, 0x2e, 0x39, 0x46, 0xed,
	0xa0, 0x71, 0x50, 0x5f, 0xed, 0xaf, 0x3e, 0x22, 0x73, 0x8e, 0x25, 0x85, 0xde, 0xc1, 0x5e, 0xc8,
	0x84, 0x3f, 0xf3, 0x27, 0x44, 0x2d, 0xb2, 0xec, 0x18, 0x35, 0xab, 0x51, 0x4d, 0xa4, 0x69, 0x16,
	0x67, 0xb4, 0xe8, 0x15, 0x94, 0x78, 0x3c, 0x1e, 0x11, 0xfe, 0xc9, 0xde, 0x75, 0x72, 0x35, 0xab,
	0xf1, 0x24, 0xf9, 0x4d, 0x13, 0x78, 0xa5, 0x70, 0xbf, 0x9b, 0x50, 0xdd, 0x2c, 0x29, 0x5f, 0xb0,
	0x90, 0x53, 0x74, 0x0d, 0xf9, 0x29, 0x11, 0xc4, 0x36, 0x94, 0xc9, 0x69, 0x62, 0x72, 0xbf, 0x7e,
	0x03, 0xbe, 0x94, 0x07, 0x10, 0xce, 0x58, 0x14, 0x2c, 0xd7, 0xa8, 0x55, 0x58, 0x79, 0xa2, 0x53,
	0x28, 0xd2, 0x28, 0x62, 0x11, 0xb7, 0x4d, 0xe5, 0x5e, 0x7f, 0xd0, 0xbd, 0x23, 0xe5, 0x6b, 0x17,
	0xfd, 0xf7, 0xf1, 0x5b, 0x38, 0x79, 0xd4, 0xb4, 0xe8, 0x00, 0xcc, 0x75, 0x73, 0x98, 0x5e, 0xfb,
	0x78, 0x08, 0xfb, 0x19, 0x47, 0xe4, 0xc2, 0x9e, 0xf2, 0xec, 0x51, 0xce, 0xc9, 0x9c, 0x6a, 0x69,
	0x06, 0x93, 0xe7, 0xca, 0x05, 0x11, 0x31, 0x6f, 0xc9, 0x73, 0x5d, 0xf6, 0x54, 0x0a, 0x71, 0xd9,
	0xba, 0xf2, 0xeb, 0xc6, 0x33, 0xb6, 0x37, 0x9e, 0xf9, 0x7b, 0xe3, 0x55, 0xa1, 0xb8, 0xb4, 0x53,
	0x5d, 0x59, 0xc6, 0x7a, 0x24, 0x71, 0x36, 0x9b, 0x71, 0x2a, 0x54, 0x3f, 0x16, 0xb0, 0x1e, 0xb9,
	0xb3, 0x6c, 0x9b, 0xa0, 0x13, 0x28, 0xd0, 0x80, 0xf8, 0xb7, 0x6a, 0x5a, 0xab, 0x71, 0x98, 0x54,
	0x55, 0xc1, 0x78, 0xc9, 0xa2, 0xd7, 0x50, 0x0a, 0xf4, 0x36, 0x4d, 0x25, 0x3c, 0x4a, 0x84, 0x82,
	0x7e, 0x13, 0x9a, 0xc4, 0x2b, 0x95, 0xdb, 0xd1, 0xbe, 0xb2, 0x4a, 0x99, 0xbe, 0x34, 0xd4, 0x32,
	0xb3, 0x8b, 0xb0, 0xa1, 0xa4, 0xc4, 0x5e, 0x5b, 0x6f, 0x71, 0x35, 0x74, 0xbf, 0x82, 0x95, 0xb2,
	0x7f, 0x94, 0x99, 0x03, 0xd6, 0x84, 0xc5, 0xa1, 0x88, 0xee, 0x52, 0x35, 0x4f, 0x43, 0xd2, 0x25,
	0x60, 0x63, 0xff, 0x96, 0xf6, 0xe3, 0x60, 0x4c, 0x23, 0x7d, 0x9f, 0x33, 0x98, 0xfb, 0xd3, 0x84,
	0xa3, 0xcb, 0xc5, 0xf4, 0x7f, 0x70, 0xfc, 0xb5, 0xe0, 0x90, 0x4b, 0x5d, 0xe5, 0xb5, 0xd7, 0xb6,
	0x61, 0x79, 0x17, 0x12, 0x44, 0x05, 0xcb, 0x66, 0xc9, 0x1f, 0x0a, 0x96, 0xfb, 0xf5, 0x1b, 0xf0,
	0x1f, 0x07, 0xcb, 0x16, 0xf7, 0xad, 0xc1, 0xf2, 0xa8, 0x69, 0xff, 0x49, 0xb0, 0xbc, 0x1c, 0x43,
	0x5e, 0x1e, 0x31, 0xda, 0x83, 0xf2, 0x05, 0xf6, 0x06, 0xd8, 0x1b, 0x5d, 0x55, 0x76, 0x10, 0x40,
	0xb1, 0x3f, 0xc0, 0xbd, 0x66, 0xb7, 0x62, 0xa0, 0x5d, 0x28, 0xf4, 0x06, 0xfd, 0xce, 0x55, 0xc5,
	0x94, 0xf0, 0xf0, 0x62, 0x80, 0x47, 0xc3, 0x4a, 0x0e, 0x59, 0x50, 0x6a, 0x0d, 0xce, 0xcf, 0xbd,
	0xfe, 0x59, 0x25, 0x8f, 0x0e, 0xc1, 0x1a, 0x0e, 0x5a, 0x4d, 0xaf, 0xeb, 0x5d, 0x4b, 0xa0, 0x20,
	0xd9, 0xb3, 0x4e, 0xbf, 0x83, 0x9b, 0xdd, 0x4a, 0xb1, 0xf1, 0x11, 0x9e, 0xca, 0xbd, 0x36, 0xf5,
	0x11, 0xf6, 0x96, 0x6f, 0x30, 0xfa, 0x00, 0x90, 0x24, 0x2c, 0x7a, 0xbe, 0x3d, 0xa7, 0xd5, 0x7d,
	0x3a, 0x76, 0x1e, 0x0a, 0x72, 0x77, 0xe7, 0x3d, 0x5c, 0xaf, 0x9f, 0xf4, 0x71, 0x51, 0xbd, 0xf1,
	0x6f, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x20, 0xb6, 0xe7, 0x93, 0x07, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TodoActivityManagerClient is the client API for TodoActivityManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoActivityManagerClient interface {
	CreateTodo(ctx context.Context, in *CreateTodoListRequest, opts ...grpc.CallOption) (*CreateTodoListResponse, error)
}

type todoActivityManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoActivityManagerClient(cc grpc.ClientConnInterface) TodoActivityManagerClient {
	return &todoActivityManagerClient{cc}
}

func (c *todoActivityManagerClient) CreateTodo(ctx context.Context, in *CreateTodoListRequest, opts ...grpc.CallOption) (*CreateTodoListResponse, error) {
	out := new(CreateTodoListResponse)
	err := c.cc.Invoke(ctx, "/todolist.TodoActivityManager/CreateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoActivityManagerServer is the server API for TodoActivityManager service.
type TodoActivityManagerServer interface {
	CreateTodo(context.Context, *CreateTodoListRequest) (*CreateTodoListResponse, error)
}

// UnimplementedTodoActivityManagerServer can be embedded to have forward compatible implementations.
type UnimplementedTodoActivityManagerServer struct {
}

func (*UnimplementedTodoActivityManagerServer) CreateTodo(ctx context.Context, req *CreateTodoListRequest) (*CreateTodoListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}

func RegisterTodoActivityManagerServer(s *grpc.Server, srv TodoActivityManagerServer) {
	s.RegisterService(&_TodoActivityManager_serviceDesc, srv)
}

func _TodoActivityManager_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoActivityManagerServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todolist.TodoActivityManager/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoActivityManagerServer).CreateTodo(ctx, req.(*CreateTodoListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoActivityManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todolist.TodoActivityManager",
	HandlerType: (*TodoActivityManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodo",
			Handler:    _TodoActivityManager_CreateTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo_list_activity_manager.proto",
}
