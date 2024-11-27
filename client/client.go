// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateServiceWorkOrderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 426556
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1477832102462645
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5
	DurationDay *string `json:"DurationDay,omitempty" xml:"DurationDay,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Y
	IsAttachment *string `json:"IsAttachment,omitempty" xml:"IsAttachment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Y
	IsWorkOrderNotify *string `json:"IsWorkOrderNotify,omitempty" xml:"IsWorkOrderNotify,omitempty"`
	// example:
	//
	// 5
	NotifyDay *string `json:"NotifyDay,omitempty" xml:"NotifyDay,omitempty"`
	// example:
	//
	// 10
	NotifyId *int64 `json:"NotifyId,omitempty" xml:"NotifyId,omitempty"`
	// This parameter is required.
	OperateRemark *string `json:"OperateRemark,omitempty" xml:"OperateRemark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CREATE
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 426556
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// This parameter is required.
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-01-21 15:25:25
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	WorkOrderDetail *string `json:"WorkOrderDetail,omitempty" xml:"WorkOrderDetail,omitempty"`
	// This parameter is required.
	WorkOrderName *string `json:"WorkOrderName,omitempty" xml:"WorkOrderName,omitempty"`
	// This parameter is required.
	WorkOrderSource *string `json:"WorkOrderSource,omitempty" xml:"WorkOrderSource,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// UNPROCESSED
	WorkOrderStatus *string `json:"WorkOrderStatus,omitempty" xml:"WorkOrderStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MONTH_REPORT
	WorkOrderType *string `json:"WorkOrderType,omitempty" xml:"WorkOrderType,omitempty"`
}

func (s CreateServiceWorkOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceWorkOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceWorkOrderRequest) SetCreator(v string) *CreateServiceWorkOrderRequest {
	s.Creator = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetCustomerId(v string) *CreateServiceWorkOrderRequest {
	s.CustomerId = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetDurationDay(v string) *CreateServiceWorkOrderRequest {
	s.DurationDay = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetIsAttachment(v string) *CreateServiceWorkOrderRequest {
	s.IsAttachment = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetIsWorkOrderNotify(v string) *CreateServiceWorkOrderRequest {
	s.IsWorkOrderNotify = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetNotifyDay(v string) *CreateServiceWorkOrderRequest {
	s.NotifyDay = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetNotifyId(v int64) *CreateServiceWorkOrderRequest {
	s.NotifyId = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetOperateRemark(v string) *CreateServiceWorkOrderRequest {
	s.OperateRemark = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetOperateType(v string) *CreateServiceWorkOrderRequest {
	s.OperateType = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetOperator(v string) *CreateServiceWorkOrderRequest {
	s.Operator = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetOwnerId(v string) *CreateServiceWorkOrderRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetStartTime(v int64) *CreateServiceWorkOrderRequest {
	s.StartTime = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetWorkOrderDetail(v string) *CreateServiceWorkOrderRequest {
	s.WorkOrderDetail = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetWorkOrderName(v string) *CreateServiceWorkOrderRequest {
	s.WorkOrderName = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetWorkOrderSource(v string) *CreateServiceWorkOrderRequest {
	s.WorkOrderSource = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetWorkOrderStatus(v string) *CreateServiceWorkOrderRequest {
	s.WorkOrderStatus = &v
	return s
}

func (s *CreateServiceWorkOrderRequest) SetWorkOrderType(v string) *CreateServiceWorkOrderRequest {
	s.WorkOrderType = &v
	return s
}

type CreateServiceWorkOrderResponseBody struct {
	// example:
	//
	// 200
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateServiceWorkOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7DC44321-7AAE-51CD-8E5F-CEB968569042
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateServiceWorkOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceWorkOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceWorkOrderResponseBody) SetCode(v string) *CreateServiceWorkOrderResponseBody {
	s.Code = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBody) SetData(v *CreateServiceWorkOrderResponseBodyData) *CreateServiceWorkOrderResponseBody {
	s.Data = v
	return s
}

func (s *CreateServiceWorkOrderResponseBody) SetHttpStatusCode(v int32) *CreateServiceWorkOrderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBody) SetMessage(v string) *CreateServiceWorkOrderResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBody) SetRequestId(v string) *CreateServiceWorkOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBody) SetSuccess(v bool) *CreateServiceWorkOrderResponseBody {
	s.Success = &v
	return s
}

type CreateServiceWorkOrderResponseBodyData struct {
	// example:
	//
	// 2024-03-07 16:45:01
	CompleteTime *int64 `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// example:
	//
	// 2024-10-04T02:19:55Z
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 426556
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 1477832102462645
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// example:
	//
	// 24-03-11 00:00:00
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1978941
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 426556
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 2021-01-21 15:25:25
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	WorkOrderDetail *string `json:"WorkOrderDetail,omitempty" xml:"WorkOrderDetail,omitempty"`
	WorkOrderName   *string `json:"WorkOrderName,omitempty" xml:"WorkOrderName,omitempty"`
	WorkOrderSource *string `json:"WorkOrderSource,omitempty" xml:"WorkOrderSource,omitempty"`
	// example:
	//
	// UNPROCESSED
	WorkOrderStatus *string `json:"WorkOrderStatus,omitempty" xml:"WorkOrderStatus,omitempty"`
	// example:
	//
	// MONTH_REPORT
	WorkOrderType *string `json:"WorkOrderType,omitempty" xml:"WorkOrderType,omitempty"`
}

func (s CreateServiceWorkOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceWorkOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateServiceWorkOrderResponseBodyData) SetCompleteTime(v int64) *CreateServiceWorkOrderResponseBodyData {
	s.CompleteTime = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetCreateTime(v int64) *CreateServiceWorkOrderResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetCreator(v string) *CreateServiceWorkOrderResponseBodyData {
	s.Creator = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetCustomerId(v string) *CreateServiceWorkOrderResponseBodyData {
	s.CustomerId = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetEndTime(v int64) *CreateServiceWorkOrderResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetId(v int64) *CreateServiceWorkOrderResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetOwnerId(v string) *CreateServiceWorkOrderResponseBodyData {
	s.OwnerId = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetStartTime(v int64) *CreateServiceWorkOrderResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetWorkOrderDetail(v string) *CreateServiceWorkOrderResponseBodyData {
	s.WorkOrderDetail = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetWorkOrderName(v string) *CreateServiceWorkOrderResponseBodyData {
	s.WorkOrderName = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetWorkOrderSource(v string) *CreateServiceWorkOrderResponseBodyData {
	s.WorkOrderSource = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetWorkOrderStatus(v string) *CreateServiceWorkOrderResponseBodyData {
	s.WorkOrderStatus = &v
	return s
}

func (s *CreateServiceWorkOrderResponseBodyData) SetWorkOrderType(v string) *CreateServiceWorkOrderResponseBodyData {
	s.WorkOrderType = &v
	return s
}

type CreateServiceWorkOrderResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceWorkOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceWorkOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceWorkOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceWorkOrderResponse) SetHeaders(v map[string]*string) *CreateServiceWorkOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceWorkOrderResponse) SetStatusCode(v int32) *CreateServiceWorkOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceWorkOrderResponse) SetBody(v *CreateServiceWorkOrderResponseBody) *CreateServiceWorkOrderResponse {
	s.Body = v
	return s
}

type DisposeServiceWorkOrderRequest struct {
	AttachmentName *string `json:"AttachmentName,omitempty" xml:"AttachmentName,omitempty"`
	// example:
	//
	// 2024-04-14 00:00:00
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 405639
	ForwardOwnerId *string `json:"ForwardOwnerId,omitempty" xml:"ForwardOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 23172
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// Y
	IsAttachment *string `json:"IsAttachment,omitempty" xml:"IsAttachment,omitempty"`
	// example:
	//
	// Y
	IsWorkOrderNotify *string `json:"IsWorkOrderNotify,omitempty" xml:"IsWorkOrderNotify,omitempty"`
	// example:
	//
	// 10
	NotifyId *int64 `json:"NotifyId,omitempty" xml:"NotifyId,omitempty"`
	// This parameter is required.
	OperateRemark *string `json:"OperateRemark,omitempty" xml:"OperateRemark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROCESSED
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 396120
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 2024-04-02 00:00:00
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 336333
	UpgradeOwnerId  *string `json:"UpgradeOwnerId,omitempty" xml:"UpgradeOwnerId,omitempty"`
	WorkOrderDetail *string `json:"WorkOrderDetail,omitempty" xml:"WorkOrderDetail,omitempty"`
	// This parameter is required.
	WorkOrderName *string `json:"WorkOrderName,omitempty" xml:"WorkOrderName,omitempty"`
	// example:
	//
	// PROCESSED
	WorkOrderStatus *string `json:"WorkOrderStatus,omitempty" xml:"WorkOrderStatus,omitempty"`
}

func (s DisposeServiceWorkOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s DisposeServiceWorkOrderRequest) GoString() string {
	return s.String()
}

func (s *DisposeServiceWorkOrderRequest) SetAttachmentName(v string) *DisposeServiceWorkOrderRequest {
	s.AttachmentName = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetEndTime(v int64) *DisposeServiceWorkOrderRequest {
	s.EndTime = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetForwardOwnerId(v string) *DisposeServiceWorkOrderRequest {
	s.ForwardOwnerId = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetId(v int64) *DisposeServiceWorkOrderRequest {
	s.Id = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetIsAttachment(v string) *DisposeServiceWorkOrderRequest {
	s.IsAttachment = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetIsWorkOrderNotify(v string) *DisposeServiceWorkOrderRequest {
	s.IsWorkOrderNotify = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetNotifyId(v int64) *DisposeServiceWorkOrderRequest {
	s.NotifyId = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetOperateRemark(v string) *DisposeServiceWorkOrderRequest {
	s.OperateRemark = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetOperateType(v string) *DisposeServiceWorkOrderRequest {
	s.OperateType = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetOperator(v string) *DisposeServiceWorkOrderRequest {
	s.Operator = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetStartTime(v int64) *DisposeServiceWorkOrderRequest {
	s.StartTime = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetUpgradeOwnerId(v string) *DisposeServiceWorkOrderRequest {
	s.UpgradeOwnerId = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetWorkOrderDetail(v string) *DisposeServiceWorkOrderRequest {
	s.WorkOrderDetail = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetWorkOrderName(v string) *DisposeServiceWorkOrderRequest {
	s.WorkOrderName = &v
	return s
}

func (s *DisposeServiceWorkOrderRequest) SetWorkOrderStatus(v string) *DisposeServiceWorkOrderRequest {
	s.WorkOrderStatus = &v
	return s
}

type DisposeServiceWorkOrderResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ED520610-6231-5D80-BADD-A8CDC7BBC809
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisposeServiceWorkOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisposeServiceWorkOrderResponseBody) GoString() string {
	return s.String()
}

func (s *DisposeServiceWorkOrderResponseBody) SetCode(v string) *DisposeServiceWorkOrderResponseBody {
	s.Code = &v
	return s
}

func (s *DisposeServiceWorkOrderResponseBody) SetHttpStatusCode(v int32) *DisposeServiceWorkOrderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DisposeServiceWorkOrderResponseBody) SetMessage(v string) *DisposeServiceWorkOrderResponseBody {
	s.Message = &v
	return s
}

func (s *DisposeServiceWorkOrderResponseBody) SetRequestId(v string) *DisposeServiceWorkOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisposeServiceWorkOrderResponseBody) SetSuccess(v bool) *DisposeServiceWorkOrderResponseBody {
	s.Success = &v
	return s
}

type DisposeServiceWorkOrderResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisposeServiceWorkOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisposeServiceWorkOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s DisposeServiceWorkOrderResponse) GoString() string {
	return s.String()
}

func (s *DisposeServiceWorkOrderResponse) SetHeaders(v map[string]*string) *DisposeServiceWorkOrderResponse {
	s.Headers = v
	return s
}

func (s *DisposeServiceWorkOrderResponse) SetStatusCode(v int32) *DisposeServiceWorkOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *DisposeServiceWorkOrderResponse) SetBody(v *DisposeServiceWorkOrderResponseBody) *DisposeServiceWorkOrderResponse {
	s.Body = v
	return s
}

type DisposeWorkTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// WB01089929
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// This parameter is required.
	OptRemark *string `json:"OptRemark,omitempty" xml:"OptRemark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10310
	TaskIds *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
}

func (s DisposeWorkTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DisposeWorkTaskRequest) GoString() string {
	return s.String()
}

func (s *DisposeWorkTaskRequest) SetOperator(v string) *DisposeWorkTaskRequest {
	s.Operator = &v
	return s
}

func (s *DisposeWorkTaskRequest) SetOptRemark(v string) *DisposeWorkTaskRequest {
	s.OptRemark = &v
	return s
}

func (s *DisposeWorkTaskRequest) SetStatus(v int32) *DisposeWorkTaskRequest {
	s.Status = &v
	return s
}

func (s *DisposeWorkTaskRequest) SetTaskIds(v string) *DisposeWorkTaskRequest {
	s.TaskIds = &v
	return s
}

type DisposeWorkTaskResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 86786E4C-6416-55CF-9AB6-5E275B68801D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisposeWorkTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisposeWorkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DisposeWorkTaskResponseBody) SetCode(v string) *DisposeWorkTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DisposeWorkTaskResponseBody) SetHttpStatusCode(v int32) *DisposeWorkTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DisposeWorkTaskResponseBody) SetMessage(v string) *DisposeWorkTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DisposeWorkTaskResponseBody) SetRequestId(v string) *DisposeWorkTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisposeWorkTaskResponseBody) SetSuccess(v bool) *DisposeWorkTaskResponseBody {
	s.Success = &v
	return s
}

type DisposeWorkTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisposeWorkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisposeWorkTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DisposeWorkTaskResponse) GoString() string {
	return s.String()
}

func (s *DisposeWorkTaskResponse) SetHeaders(v map[string]*string) *DisposeWorkTaskResponse {
	s.Headers = v
	return s
}

func (s *DisposeWorkTaskResponse) SetStatusCode(v int32) *DisposeWorkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DisposeWorkTaskResponse) SetBody(v *DisposeWorkTaskResponseBody) *DisposeWorkTaskResponse {
	s.Body = v
	return s
}

type GetAttackedAssetDealRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// month
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1732268720000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1732268720000
	StartDate       *int64  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	SuspEventSource *string `json:"SuspEventSource,omitempty" xml:"SuspEventSource,omitempty"`
}

func (s GetAttackedAssetDealRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAttackedAssetDealRequest) GoString() string {
	return s.String()
}

func (s *GetAttackedAssetDealRequest) SetDateType(v string) *GetAttackedAssetDealRequest {
	s.DateType = &v
	return s
}

func (s *GetAttackedAssetDealRequest) SetEndDate(v int64) *GetAttackedAssetDealRequest {
	s.EndDate = &v
	return s
}

func (s *GetAttackedAssetDealRequest) SetStartDate(v int64) *GetAttackedAssetDealRequest {
	s.StartDate = &v
	return s
}

func (s *GetAttackedAssetDealRequest) SetSuspEventSource(v string) *GetAttackedAssetDealRequest {
	s.SuspEventSource = &v
	return s
}

type GetAttackedAssetDealResponseBody struct {
	// example:
	//
	// 200
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAttackedAssetDealResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1E74F11C-B4A8-5774-962C-02003BA8504E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAttackedAssetDealResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAttackedAssetDealResponseBody) GoString() string {
	return s.String()
}

func (s *GetAttackedAssetDealResponseBody) SetCode(v string) *GetAttackedAssetDealResponseBody {
	s.Code = &v
	return s
}

func (s *GetAttackedAssetDealResponseBody) SetData(v *GetAttackedAssetDealResponseBodyData) *GetAttackedAssetDealResponseBody {
	s.Data = v
	return s
}

func (s *GetAttackedAssetDealResponseBody) SetHttpStatusCode(v int32) *GetAttackedAssetDealResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAttackedAssetDealResponseBody) SetMessage(v string) *GetAttackedAssetDealResponseBody {
	s.Message = &v
	return s
}

func (s *GetAttackedAssetDealResponseBody) SetRequestId(v string) *GetAttackedAssetDealResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAttackedAssetDealResponseBody) SetSuccess(v bool) *GetAttackedAssetDealResponseBody {
	s.Success = &v
	return s
}

type GetAttackedAssetDealResponseBodyData struct {
	EcsTrendList []*GetAttackedAssetDealResponseBodyDataEcsTrendList `json:"EcsTrendList,omitempty" xml:"EcsTrendList,omitempty" type:"Repeated"`
}

func (s GetAttackedAssetDealResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAttackedAssetDealResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAttackedAssetDealResponseBodyData) SetEcsTrendList(v []*GetAttackedAssetDealResponseBodyDataEcsTrendList) *GetAttackedAssetDealResponseBodyData {
	s.EcsTrendList = v
	return s
}

type GetAttackedAssetDealResponseBodyDataEcsTrendList struct {
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// 2
	DealCount *int64 `json:"DealCount,omitempty" xml:"DealCount,omitempty"`
	FindCount *int64 `json:"FindCount,omitempty" xml:"FindCount,omitempty"`
}

func (s GetAttackedAssetDealResponseBodyDataEcsTrendList) String() string {
	return tea.Prettify(s)
}

func (s GetAttackedAssetDealResponseBodyDataEcsTrendList) GoString() string {
	return s.String()
}

func (s *GetAttackedAssetDealResponseBodyDataEcsTrendList) SetDate(v string) *GetAttackedAssetDealResponseBodyDataEcsTrendList {
	s.Date = &v
	return s
}

func (s *GetAttackedAssetDealResponseBodyDataEcsTrendList) SetDealCount(v int64) *GetAttackedAssetDealResponseBodyDataEcsTrendList {
	s.DealCount = &v
	return s
}

func (s *GetAttackedAssetDealResponseBodyDataEcsTrendList) SetFindCount(v int64) *GetAttackedAssetDealResponseBodyDataEcsTrendList {
	s.FindCount = &v
	return s
}

type GetAttackedAssetDealResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAttackedAssetDealResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAttackedAssetDealResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAttackedAssetDealResponse) GoString() string {
	return s.String()
}

func (s *GetAttackedAssetDealResponse) SetHeaders(v map[string]*string) *GetAttackedAssetDealResponse {
	s.Headers = v
	return s
}

func (s *GetAttackedAssetDealResponse) SetStatusCode(v int32) *GetAttackedAssetDealResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAttackedAssetDealResponse) SetBody(v *GetAttackedAssetDealResponseBody) *GetAttackedAssetDealResponse {
	s.Body = v
	return s
}

type GetBaselineSummaryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// month
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1732156885986
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1729478485000
	StartDate       *int64  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	SuspEventSource *string `json:"SuspEventSource,omitempty" xml:"SuspEventSource,omitempty"`
}

func (s GetBaselineSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBaselineSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetBaselineSummaryRequest) SetDateType(v string) *GetBaselineSummaryRequest {
	s.DateType = &v
	return s
}

func (s *GetBaselineSummaryRequest) SetEndDate(v int64) *GetBaselineSummaryRequest {
	s.EndDate = &v
	return s
}

func (s *GetBaselineSummaryRequest) SetStartDate(v int64) *GetBaselineSummaryRequest {
	s.StartDate = &v
	return s
}

func (s *GetBaselineSummaryRequest) SetSuspEventSource(v string) *GetBaselineSummaryRequest {
	s.SuspEventSource = &v
	return s
}

type GetBaselineSummaryResponseBody struct {
	// example:
	//
	// 200
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetBaselineSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 67D61738-5E38-5164-947A-34E3850D493A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBaselineSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBaselineSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetBaselineSummaryResponseBody) SetCode(v string) *GetBaselineSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetBaselineSummaryResponseBody) SetData(v *GetBaselineSummaryResponseBodyData) *GetBaselineSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetBaselineSummaryResponseBody) SetHttpStatusCode(v int32) *GetBaselineSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetBaselineSummaryResponseBody) SetMessage(v string) *GetBaselineSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetBaselineSummaryResponseBody) SetRequestId(v string) *GetBaselineSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBaselineSummaryResponseBody) SetSuccess(v bool) *GetBaselineSummaryResponseBody {
	s.Success = &v
	return s
}

type GetBaselineSummaryResponseBodyData struct {
	TrendDTOList []*GetBaselineSummaryResponseBodyDataTrendDTOList `json:"TrendDTOList,omitempty" xml:"TrendDTOList,omitempty" type:"Repeated"`
}

func (s GetBaselineSummaryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetBaselineSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBaselineSummaryResponseBodyData) SetTrendDTOList(v []*GetBaselineSummaryResponseBodyDataTrendDTOList) *GetBaselineSummaryResponseBodyData {
	s.TrendDTOList = v
	return s
}

type GetBaselineSummaryResponseBodyDataTrendDTOList struct {
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// 10
	DealCount *int64 `json:"DealCount,omitempty" xml:"DealCount,omitempty"`
	// example:
	//
	// 12
	FindCount *int64 `json:"FindCount,omitempty" xml:"FindCount,omitempty"`
}

func (s GetBaselineSummaryResponseBodyDataTrendDTOList) String() string {
	return tea.Prettify(s)
}

func (s GetBaselineSummaryResponseBodyDataTrendDTOList) GoString() string {
	return s.String()
}

func (s *GetBaselineSummaryResponseBodyDataTrendDTOList) SetDate(v string) *GetBaselineSummaryResponseBodyDataTrendDTOList {
	s.Date = &v
	return s
}

func (s *GetBaselineSummaryResponseBodyDataTrendDTOList) SetDealCount(v int64) *GetBaselineSummaryResponseBodyDataTrendDTOList {
	s.DealCount = &v
	return s
}

func (s *GetBaselineSummaryResponseBodyDataTrendDTOList) SetFindCount(v int64) *GetBaselineSummaryResponseBodyDataTrendDTOList {
	s.FindCount = &v
	return s
}

type GetBaselineSummaryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBaselineSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBaselineSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBaselineSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetBaselineSummaryResponse) SetHeaders(v map[string]*string) *GetBaselineSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetBaselineSummaryResponse) SetStatusCode(v int32) *GetBaselineSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBaselineSummaryResponse) SetBody(v *GetBaselineSummaryResponseBody) *GetBaselineSummaryResponse {
	s.Body = v
	return s
}

type GetDetailByIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 22
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDetailByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDetailByIdRequest) GoString() string {
	return s.String()
}

func (s *GetDetailByIdRequest) SetId(v int64) *GetDetailByIdRequest {
	s.Id = &v
	return s
}

type GetDetailByIdResponseBody struct {
	// example:
	//
	// 404
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetDetailByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DAB46EC5-3746-59C4-B6D2-469F442EC73F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDetailByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDetailByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetDetailByIdResponseBody) SetCode(v string) *GetDetailByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetDetailByIdResponseBody) SetData(v *GetDetailByIdResponseBodyData) *GetDetailByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetDetailByIdResponseBody) SetHttpStatusCode(v int32) *GetDetailByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDetailByIdResponseBody) SetMessage(v string) *GetDetailByIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetDetailByIdResponseBody) SetRequestId(v string) *GetDetailByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDetailByIdResponseBody) SetSuccess(v bool) *GetDetailByIdResponseBody {
	s.Success = &v
	return s
}

type GetDetailByIdResponseBodyData struct {
	VulDetails []*GetDetailByIdResponseBodyDataVulDetails `json:"VulDetails,omitempty" xml:"VulDetails,omitempty" type:"Repeated"`
}

func (s GetDetailByIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDetailByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDetailByIdResponseBodyData) SetVulDetails(v []*GetDetailByIdResponseBodyDataVulDetails) *GetDetailByIdResponseBodyData {
	s.VulDetails = v
	return s
}

type GetDetailByIdResponseBodyDataVulDetails struct {
	// example:
	//
	// CVE-2022-21291
	CveId *string `json:"CveId,omitempty" xml:"CveId,omitempty"`
	// example:
	//
	// 10.0
	CvssScore *string `json:"CvssScore,omitempty" xml:"CvssScore,omitempty"`
	// example:
	//
	// https://avd.aliyun.com/detail/CVE-2022-21291
	FixSuggestion *string `json:"FixSuggestion,omitempty" xml:"FixSuggestion,omitempty"`
	// example:
	//
	// Chanjet T-Plus SetupAccount/Upload. Aspx file upload vulnerability(CNVD-2022-60632)
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetDetailByIdResponseBodyDataVulDetails) String() string {
	return tea.Prettify(s)
}

func (s GetDetailByIdResponseBodyDataVulDetails) GoString() string {
	return s.String()
}

func (s *GetDetailByIdResponseBodyDataVulDetails) SetCveId(v string) *GetDetailByIdResponseBodyDataVulDetails {
	s.CveId = &v
	return s
}

func (s *GetDetailByIdResponseBodyDataVulDetails) SetCvssScore(v string) *GetDetailByIdResponseBodyDataVulDetails {
	s.CvssScore = &v
	return s
}

func (s *GetDetailByIdResponseBodyDataVulDetails) SetFixSuggestion(v string) *GetDetailByIdResponseBodyDataVulDetails {
	s.FixSuggestion = &v
	return s
}

func (s *GetDetailByIdResponseBodyDataVulDetails) SetTitle(v string) *GetDetailByIdResponseBodyDataVulDetails {
	s.Title = &v
	return s
}

type GetDetailByIdResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDetailByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDetailByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDetailByIdResponse) GoString() string {
	return s.String()
}

func (s *GetDetailByIdResponse) SetHeaders(v map[string]*string) *GetDetailByIdResponse {
	s.Headers = v
	return s
}

func (s *GetDetailByIdResponse) SetStatusCode(v int32) *GetDetailByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDetailByIdResponse) SetBody(v *GetDetailByIdResponseBody) *GetDetailByIdResponse {
	s.Body = v
	return s
}

type GetDocumentDownloadUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 175815
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDocumentDownloadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentDownloadUrlRequest) SetId(v int64) *GetDocumentDownloadUrlRequest {
	s.Id = &v
	return s
}

type GetDocumentDownloadUrlResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// https://oos-cn.ctyunapi.cn/example-bucket/test/1.jpg
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// C7BE80B4-7692-54FA-AB22-2A7DF08C4754
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDocumentDownloadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentDownloadUrlResponseBody) SetCode(v string) *GetDocumentDownloadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocumentDownloadUrlResponseBody) SetData(v string) *GetDocumentDownloadUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetDocumentDownloadUrlResponseBody) SetHttpStatusCode(v int32) *GetDocumentDownloadUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDocumentDownloadUrlResponseBody) SetMessage(v string) *GetDocumentDownloadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentDownloadUrlResponseBody) SetRequestId(v string) *GetDocumentDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentDownloadUrlResponseBody) SetSuccess(v bool) *GetDocumentDownloadUrlResponseBody {
	s.Success = &v
	return s
}

type GetDocumentDownloadUrlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentDownloadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentDownloadUrlResponse) SetHeaders(v map[string]*string) *GetDocumentDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentDownloadUrlResponse) SetStatusCode(v int32) *GetDocumentDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentDownloadUrlResponse) SetBody(v *GetDocumentDownloadUrlResponseBody) *GetDocumentDownloadUrlResponse {
	s.Body = v
	return s
}

type GetDocumentPageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DeliveredBy  *string `json:"DeliveredBy,omitempty" xml:"DeliveredBy,omitempty"`
	DocumentName *string `json:"DocumentName,omitempty" xml:"DocumentName,omitempty"`
	// example:
	//
	// 0
	DocumentType *string `json:"DocumentType,omitempty" xml:"DocumentType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
}

func (s GetDocumentPageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentPageRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentPageRequest) SetCurrentPage(v int32) *GetDocumentPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetDocumentPageRequest) SetDeliveredBy(v string) *GetDocumentPageRequest {
	s.DeliveredBy = &v
	return s
}

func (s *GetDocumentPageRequest) SetDocumentName(v string) *GetDocumentPageRequest {
	s.DocumentName = &v
	return s
}

func (s *GetDocumentPageRequest) SetDocumentType(v string) *GetDocumentPageRequest {
	s.DocumentType = &v
	return s
}

func (s *GetDocumentPageRequest) SetPageSize(v int32) *GetDocumentPageRequest {
	s.PageSize = &v
	return s
}

func (s *GetDocumentPageRequest) SetReportType(v string) *GetDocumentPageRequest {
	s.ReportType = &v
	return s
}

type GetDocumentPageResponseBody struct {
	// example:
	//
	// 200
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetDocumentPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	PageInfo *GetDocumentPageResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 04DAD7B4-E1DA-5C2C-8E5C-A1EDC880CF60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDocumentPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentPageResponseBody) SetCode(v string) *GetDocumentPageResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocumentPageResponseBody) SetData(v []*GetDocumentPageResponseBodyData) *GetDocumentPageResponseBody {
	s.Data = v
	return s
}

func (s *GetDocumentPageResponseBody) SetHttpStatusCode(v int32) *GetDocumentPageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDocumentPageResponseBody) SetMessage(v string) *GetDocumentPageResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentPageResponseBody) SetPageInfo(v *GetDocumentPageResponseBodyPageInfo) *GetDocumentPageResponseBody {
	s.PageInfo = v
	return s
}

func (s *GetDocumentPageResponseBody) SetRequestId(v string) *GetDocumentPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentPageResponseBody) SetSuccess(v bool) *GetDocumentPageResponseBody {
	s.Success = &v
	return s
}

type GetDocumentPageResponseBodyData struct {
	DeliveredBy  *string `json:"DeliveredBy,omitempty" xml:"DeliveredBy,omitempty"`
	DocumentName *string `json:"DocumentName,omitempty" xml:"DocumentName,omitempty"`
	// example:
	//
	// 3
	DocumentType *string `json:"DocumentType,omitempty" xml:"DocumentType,omitempty"`
	// example:
	//
	// 346409
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2023-03-21 17:26:34
	UploadTime *string `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
}

func (s GetDocumentPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocumentPageResponseBodyData) SetDeliveredBy(v string) *GetDocumentPageResponseBodyData {
	s.DeliveredBy = &v
	return s
}

func (s *GetDocumentPageResponseBodyData) SetDocumentName(v string) *GetDocumentPageResponseBodyData {
	s.DocumentName = &v
	return s
}

func (s *GetDocumentPageResponseBodyData) SetDocumentType(v string) *GetDocumentPageResponseBodyData {
	s.DocumentType = &v
	return s
}

func (s *GetDocumentPageResponseBodyData) SetId(v int64) *GetDocumentPageResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetDocumentPageResponseBodyData) SetUploadTime(v string) *GetDocumentPageResponseBodyData {
	s.UploadTime = &v
	return s
}

type GetDocumentPageResponseBodyPageInfo struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 3149
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetDocumentPageResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentPageResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *GetDocumentPageResponseBodyPageInfo) SetCurrentPage(v int32) *GetDocumentPageResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *GetDocumentPageResponseBodyPageInfo) SetPageSize(v int32) *GetDocumentPageResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *GetDocumentPageResponseBodyPageInfo) SetTotalCount(v int32) *GetDocumentPageResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

type GetDocumentPageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentPageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentPageResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentPageResponse) SetHeaders(v map[string]*string) *GetDocumentPageResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentPageResponse) SetStatusCode(v int32) *GetDocumentPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentPageResponse) SetBody(v *GetDocumentPageResponseBody) *GetDocumentPageResponse {
	s.Body = v
	return s
}

type GetDocumentSummaryRequest struct {
	// example:
	//
	// 1
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
}

func (s GetDocumentSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentSummaryRequest) SetReportType(v string) *GetDocumentSummaryRequest {
	s.ReportType = &v
	return s
}

type GetDocumentSummaryResponseBody struct {
	// example:
	//
	// 200
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetDocumentSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7903F2DE-D9EE-5D16-8A08-E9223E54B281
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDocumentSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentSummaryResponseBody) SetCode(v string) *GetDocumentSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocumentSummaryResponseBody) SetData(v *GetDocumentSummaryResponseBodyData) *GetDocumentSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetDocumentSummaryResponseBody) SetHttpStatusCode(v int32) *GetDocumentSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDocumentSummaryResponseBody) SetMessage(v string) *GetDocumentSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentSummaryResponseBody) SetRequestId(v string) *GetDocumentSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentSummaryResponseBody) SetSuccess(v bool) *GetDocumentSummaryResponseBody {
	s.Success = &v
	return s
}

type GetDocumentSummaryResponseBodyData struct {
	// example:
	//
	// 10
	DocumentCount *int64 `json:"DocumentCount,omitempty" xml:"DocumentCount,omitempty"`
	// example:
	//
	// 10
	Frequency *int64 `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
}

func (s GetDocumentSummaryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocumentSummaryResponseBodyData) SetDocumentCount(v int64) *GetDocumentSummaryResponseBodyData {
	s.DocumentCount = &v
	return s
}

func (s *GetDocumentSummaryResponseBodyData) SetFrequency(v int64) *GetDocumentSummaryResponseBodyData {
	s.Frequency = &v
	return s
}

type GetDocumentSummaryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentSummaryResponse) SetHeaders(v map[string]*string) *GetDocumentSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentSummaryResponse) SetStatusCode(v int32) *GetDocumentSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentSummaryResponse) SetBody(v *GetDocumentSummaryResponseBody) *GetDocumentSummaryResponse {
	s.Body = v
	return s
}

type GetRecentDocumentRequest struct {
	// This parameter is required.
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// This parameter is required.
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	StartDate       *int64  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	SuspEventSource *string `json:"SuspEventSource,omitempty" xml:"SuspEventSource,omitempty"`
}

func (s GetRecentDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRecentDocumentRequest) GoString() string {
	return s.String()
}

func (s *GetRecentDocumentRequest) SetDateType(v string) *GetRecentDocumentRequest {
	s.DateType = &v
	return s
}

func (s *GetRecentDocumentRequest) SetEndDate(v int64) *GetRecentDocumentRequest {
	s.EndDate = &v
	return s
}

func (s *GetRecentDocumentRequest) SetStartDate(v int64) *GetRecentDocumentRequest {
	s.StartDate = &v
	return s
}

func (s *GetRecentDocumentRequest) SetSuspEventSource(v string) *GetRecentDocumentRequest {
	s.SuspEventSource = &v
	return s
}

type GetRecentDocumentResponseBody struct {
	// example:
	//
	// 200
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetRecentDocumentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4916FA8D-F294-518D-B373-8B59D63CAB19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRecentDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRecentDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecentDocumentResponseBody) SetCode(v string) *GetRecentDocumentResponseBody {
	s.Code = &v
	return s
}

func (s *GetRecentDocumentResponseBody) SetData(v []*GetRecentDocumentResponseBodyData) *GetRecentDocumentResponseBody {
	s.Data = v
	return s
}

func (s *GetRecentDocumentResponseBody) SetHttpStatusCode(v int32) *GetRecentDocumentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRecentDocumentResponseBody) SetMessage(v string) *GetRecentDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *GetRecentDocumentResponseBody) SetRequestId(v string) *GetRecentDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRecentDocumentResponseBody) SetSuccess(v bool) *GetRecentDocumentResponseBody {
	s.Success = &v
	return s
}

type GetRecentDocumentResponseBodyData struct {
	// example:
	//
	// 360491
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2023-03-20 14:30:38
	UploadTime *string `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
}

func (s GetRecentDocumentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRecentDocumentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRecentDocumentResponseBodyData) SetId(v int64) *GetRecentDocumentResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetRecentDocumentResponseBodyData) SetName(v string) *GetRecentDocumentResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetRecentDocumentResponseBodyData) SetUploadTime(v string) *GetRecentDocumentResponseBodyData {
	s.UploadTime = &v
	return s
}

type GetRecentDocumentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRecentDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRecentDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRecentDocumentResponse) GoString() string {
	return s.String()
}

func (s *GetRecentDocumentResponse) SetHeaders(v map[string]*string) *GetRecentDocumentResponse {
	s.Headers = v
	return s
}

func (s *GetRecentDocumentResponse) SetStatusCode(v int32) *GetRecentDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecentDocumentResponse) SetBody(v *GetRecentDocumentResponseBody) *GetRecentDocumentResponse {
	s.Body = v
	return s
}

type GetSafetyCoverRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// month
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1732268720000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1732255620000
	StartDate       *int64  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	SuspEventSource *string `json:"SuspEventSource,omitempty" xml:"SuspEventSource,omitempty"`
}

func (s GetSafetyCoverRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSafetyCoverRequest) GoString() string {
	return s.String()
}

func (s *GetSafetyCoverRequest) SetDateType(v string) *GetSafetyCoverRequest {
	s.DateType = &v
	return s
}

func (s *GetSafetyCoverRequest) SetEndDate(v int64) *GetSafetyCoverRequest {
	s.EndDate = &v
	return s
}

func (s *GetSafetyCoverRequest) SetStartDate(v int64) *GetSafetyCoverRequest {
	s.StartDate = &v
	return s
}

func (s *GetSafetyCoverRequest) SetSuspEventSource(v string) *GetSafetyCoverRequest {
	s.SuspEventSource = &v
	return s
}

type GetSafetyCoverResponseBody struct {
	// example:
	//
	// 404
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetSafetyCoverResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// system error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 564f8bb9-df3c-42a0-877a-b35d48f66603
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSafetyCoverResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSafetyCoverResponseBody) GoString() string {
	return s.String()
}

func (s *GetSafetyCoverResponseBody) SetCode(v string) *GetSafetyCoverResponseBody {
	s.Code = &v
	return s
}

func (s *GetSafetyCoverResponseBody) SetData(v *GetSafetyCoverResponseBodyData) *GetSafetyCoverResponseBody {
	s.Data = v
	return s
}

func (s *GetSafetyCoverResponseBody) SetHttpStatusCode(v int32) *GetSafetyCoverResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSafetyCoverResponseBody) SetMessage(v string) *GetSafetyCoverResponseBody {
	s.Message = &v
	return s
}

func (s *GetSafetyCoverResponseBody) SetRequestId(v string) *GetSafetyCoverResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSafetyCoverResponseBody) SetSuccess(v bool) *GetSafetyCoverResponseBody {
	s.Success = &v
	return s
}

type GetSafetyCoverResponseBodyData struct {
	CfwProtection *GetSafetyCoverResponseBodyDataCfwProtection `json:"CfwProtection,omitempty" xml:"CfwProtection,omitempty" type:"Struct"`
	EcsProtection *GetSafetyCoverResponseBodyDataEcsProtection `json:"EcsProtection,omitempty" xml:"EcsProtection,omitempty" type:"Struct"`
	WafProtection *GetSafetyCoverResponseBodyDataWafProtection `json:"WafProtection,omitempty" xml:"WafProtection,omitempty" type:"Struct"`
}

func (s GetSafetyCoverResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSafetyCoverResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSafetyCoverResponseBodyData) SetCfwProtection(v *GetSafetyCoverResponseBodyDataCfwProtection) *GetSafetyCoverResponseBodyData {
	s.CfwProtection = v
	return s
}

func (s *GetSafetyCoverResponseBodyData) SetEcsProtection(v *GetSafetyCoverResponseBodyDataEcsProtection) *GetSafetyCoverResponseBodyData {
	s.EcsProtection = v
	return s
}

func (s *GetSafetyCoverResponseBodyData) SetWafProtection(v *GetSafetyCoverResponseBodyDataWafProtection) *GetSafetyCoverResponseBodyData {
	s.WafProtection = v
	return s
}

type GetSafetyCoverResponseBodyDataCfwProtection struct {
	// example:
	//
	// 5
	NoProtectionCount *int64 `json:"NoProtectionCount,omitempty" xml:"NoProtectionCount,omitempty"`
	// example:
	//
	// 5
	ProtectionCount *int64 `json:"ProtectionCount,omitempty" xml:"ProtectionCount,omitempty"`
	// example:
	//
	// 35.00
	ProtectionGrowthRate *string `json:"ProtectionGrowthRate,omitempty" xml:"ProtectionGrowthRate,omitempty"`
	// example:
	//
	// 50.00
	ProtectionRate *string `json:"ProtectionRate,omitempty" xml:"ProtectionRate,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetSafetyCoverResponseBodyDataCfwProtection) String() string {
	return tea.Prettify(s)
}

func (s GetSafetyCoverResponseBodyDataCfwProtection) GoString() string {
	return s.String()
}

func (s *GetSafetyCoverResponseBodyDataCfwProtection) SetNoProtectionCount(v int64) *GetSafetyCoverResponseBodyDataCfwProtection {
	s.NoProtectionCount = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataCfwProtection) SetProtectionCount(v int64) *GetSafetyCoverResponseBodyDataCfwProtection {
	s.ProtectionCount = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataCfwProtection) SetProtectionGrowthRate(v string) *GetSafetyCoverResponseBodyDataCfwProtection {
	s.ProtectionGrowthRate = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataCfwProtection) SetProtectionRate(v string) *GetSafetyCoverResponseBodyDataCfwProtection {
	s.ProtectionRate = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataCfwProtection) SetTotalCount(v int64) *GetSafetyCoverResponseBodyDataCfwProtection {
	s.TotalCount = &v
	return s
}

type GetSafetyCoverResponseBodyDataEcsProtection struct {
	// example:
	//
	// 5
	NoProtectionCount *int64 `json:"NoProtectionCount,omitempty" xml:"NoProtectionCount,omitempty"`
	// example:
	//
	// 5
	ProtectionCount *int64 `json:"ProtectionCount,omitempty" xml:"ProtectionCount,omitempty"`
	// example:
	//
	// 35.00
	ProtectionGrowthRate *string `json:"ProtectionGrowthRate,omitempty" xml:"ProtectionGrowthRate,omitempty"`
	// example:
	//
	// 50.00
	ProtectionRate *string `json:"ProtectionRate,omitempty" xml:"ProtectionRate,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetSafetyCoverResponseBodyDataEcsProtection) String() string {
	return tea.Prettify(s)
}

func (s GetSafetyCoverResponseBodyDataEcsProtection) GoString() string {
	return s.String()
}

func (s *GetSafetyCoverResponseBodyDataEcsProtection) SetNoProtectionCount(v int64) *GetSafetyCoverResponseBodyDataEcsProtection {
	s.NoProtectionCount = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataEcsProtection) SetProtectionCount(v int64) *GetSafetyCoverResponseBodyDataEcsProtection {
	s.ProtectionCount = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataEcsProtection) SetProtectionGrowthRate(v string) *GetSafetyCoverResponseBodyDataEcsProtection {
	s.ProtectionGrowthRate = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataEcsProtection) SetProtectionRate(v string) *GetSafetyCoverResponseBodyDataEcsProtection {
	s.ProtectionRate = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataEcsProtection) SetTotalCount(v int64) *GetSafetyCoverResponseBodyDataEcsProtection {
	s.TotalCount = &v
	return s
}

type GetSafetyCoverResponseBodyDataWafProtection struct {
	// example:
	//
	// 5
	NoProtectionCount *int64 `json:"NoProtectionCount,omitempty" xml:"NoProtectionCount,omitempty"`
	// example:
	//
	// 5
	ProtectionCount *int64 `json:"ProtectionCount,omitempty" xml:"ProtectionCount,omitempty"`
	// example:
	//
	// 35.00
	ProtectionGrowthRate *string `json:"ProtectionGrowthRate,omitempty" xml:"ProtectionGrowthRate,omitempty"`
	// example:
	//
	// 50.00
	ProtectionRate *string `json:"ProtectionRate,omitempty" xml:"ProtectionRate,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetSafetyCoverResponseBodyDataWafProtection) String() string {
	return tea.Prettify(s)
}

func (s GetSafetyCoverResponseBodyDataWafProtection) GoString() string {
	return s.String()
}

func (s *GetSafetyCoverResponseBodyDataWafProtection) SetNoProtectionCount(v int64) *GetSafetyCoverResponseBodyDataWafProtection {
	s.NoProtectionCount = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataWafProtection) SetProtectionCount(v int64) *GetSafetyCoverResponseBodyDataWafProtection {
	s.ProtectionCount = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataWafProtection) SetProtectionGrowthRate(v string) *GetSafetyCoverResponseBodyDataWafProtection {
	s.ProtectionGrowthRate = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataWafProtection) SetProtectionRate(v string) *GetSafetyCoverResponseBodyDataWafProtection {
	s.ProtectionRate = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataWafProtection) SetTotalCount(v int64) *GetSafetyCoverResponseBodyDataWafProtection {
	s.TotalCount = &v
	return s
}

type GetSafetyCoverResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSafetyCoverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSafetyCoverResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSafetyCoverResponse) GoString() string {
	return s.String()
}

func (s *GetSafetyCoverResponse) SetHeaders(v map[string]*string) *GetSafetyCoverResponse {
	s.Headers = v
	return s
}

func (s *GetSafetyCoverResponse) SetStatusCode(v int32) *GetSafetyCoverResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSafetyCoverResponse) SetBody(v *GetSafetyCoverResponseBody) *GetSafetyCoverResponse {
	s.Body = v
	return s
}

type GetSuspEventPageRequest struct {
	// example:
	//
	// 1732515522000
	AlarmEndTime *int64 `json:"AlarmEndTime,omitempty" xml:"AlarmEndTime,omitempty"`
	// example:
	//
	// 1722515522000
	AlarmStartTime *int64 `json:"AlarmStartTime,omitempty" xml:"AlarmStartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// SUSP_EVENT
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetSuspEventPageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventPageRequest) GoString() string {
	return s.String()
}

func (s *GetSuspEventPageRequest) SetAlarmEndTime(v int64) *GetSuspEventPageRequest {
	s.AlarmEndTime = &v
	return s
}

func (s *GetSuspEventPageRequest) SetAlarmStartTime(v int64) *GetSuspEventPageRequest {
	s.AlarmStartTime = &v
	return s
}

func (s *GetSuspEventPageRequest) SetCurrentPage(v int32) *GetSuspEventPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSuspEventPageRequest) SetPageSize(v int32) *GetSuspEventPageRequest {
	s.PageSize = &v
	return s
}

func (s *GetSuspEventPageRequest) SetSource(v string) *GetSuspEventPageRequest {
	s.Source = &v
	return s
}

func (s *GetSuspEventPageRequest) SetStatus(v int32) *GetSuspEventPageRequest {
	s.Status = &v
	return s
}

type GetSuspEventPageResponseBody struct {
	// example:
	//
	// OK
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetSuspEventPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// system error
	Message  *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	PageInfo *GetSuspEventPageResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// AFA6F7B7-7C4B-58BB-B8FB-E0FFA4483561
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSuspEventPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetSuspEventPageResponseBody) SetCode(v string) *GetSuspEventPageResponseBody {
	s.Code = &v
	return s
}

func (s *GetSuspEventPageResponseBody) SetData(v []*GetSuspEventPageResponseBodyData) *GetSuspEventPageResponseBody {
	s.Data = v
	return s
}

func (s *GetSuspEventPageResponseBody) SetHttpStatusCode(v int32) *GetSuspEventPageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSuspEventPageResponseBody) SetMessage(v string) *GetSuspEventPageResponseBody {
	s.Message = &v
	return s
}

func (s *GetSuspEventPageResponseBody) SetPageInfo(v *GetSuspEventPageResponseBodyPageInfo) *GetSuspEventPageResponseBody {
	s.PageInfo = v
	return s
}

func (s *GetSuspEventPageResponseBody) SetRequestId(v string) *GetSuspEventPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSuspEventPageResponseBody) SetSuccess(v bool) *GetSuspEventPageResponseBody {
	s.Success = &v
	return s
}

type GetSuspEventPageResponseBodyData struct {
	AlarmEventType *string `json:"AlarmEventType,omitempty" xml:"AlarmEventType,omitempty"`
	// example:
	//
	// 5b1eeebe4f22daa2b177298234214fa3
	AlarmId   *int64  `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	AlarmName *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	// example:
	//
	// SUSP_EVENT
	AlarmSource *string `json:"AlarmSource,omitempty" xml:"AlarmSource,omitempty"`
	// example:
	//
	// 1722515522000
	AlarmTime      *string `json:"AlarmTime,omitempty" xml:"AlarmTime,omitempty"`
	AnalysisResult *string `json:"AnalysisResult,omitempty" xml:"AnalysisResult,omitempty"`
	// example:
	//
	// 1732515522000
	DealTime *string `json:"DealTime,omitempty" xml:"DealTime,omitempty"`
	// example:
	//
	// suspicious
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// example:
	//
	// 9947
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// shells-azhou
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 47.99.188.31
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// example:
	//
	// 172.16.109.130
	IntranetIp     *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	OccurrenceTime *string `json:"OccurrenceTime,omitempty" xml:"OccurrenceTime,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Remark         *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetSuspEventPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSuspEventPageResponseBodyData) SetAlarmEventType(v string) *GetSuspEventPageResponseBodyData {
	s.AlarmEventType = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetAlarmId(v int64) *GetSuspEventPageResponseBodyData {
	s.AlarmId = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetAlarmName(v string) *GetSuspEventPageResponseBodyData {
	s.AlarmName = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetAlarmSource(v string) *GetSuspEventPageResponseBodyData {
	s.AlarmSource = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetAlarmTime(v string) *GetSuspEventPageResponseBodyData {
	s.AlarmTime = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetAnalysisResult(v string) *GetSuspEventPageResponseBodyData {
	s.AnalysisResult = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetDealTime(v string) *GetSuspEventPageResponseBodyData {
	s.DealTime = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetEventLevel(v string) *GetSuspEventPageResponseBodyData {
	s.EventLevel = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetId(v int64) *GetSuspEventPageResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetInstanceName(v string) *GetSuspEventPageResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetInternetIp(v string) *GetSuspEventPageResponseBodyData {
	s.InternetIp = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetIntranetIp(v string) *GetSuspEventPageResponseBodyData {
	s.IntranetIp = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetOccurrenceTime(v string) *GetSuspEventPageResponseBodyData {
	s.OccurrenceTime = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetOwnerId(v string) *GetSuspEventPageResponseBodyData {
	s.OwnerId = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetRemark(v string) *GetSuspEventPageResponseBodyData {
	s.Remark = &v
	return s
}

func (s *GetSuspEventPageResponseBodyData) SetStatus(v string) *GetSuspEventPageResponseBodyData {
	s.Status = &v
	return s
}

type GetSuspEventPageResponseBodyPageInfo struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetSuspEventPageResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventPageResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *GetSuspEventPageResponseBodyPageInfo) SetCurrentPage(v int32) *GetSuspEventPageResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *GetSuspEventPageResponseBodyPageInfo) SetPageSize(v int32) *GetSuspEventPageResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *GetSuspEventPageResponseBodyPageInfo) SetTotalCount(v int32) *GetSuspEventPageResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

type GetSuspEventPageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSuspEventPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSuspEventPageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventPageResponse) GoString() string {
	return s.String()
}

func (s *GetSuspEventPageResponse) SetHeaders(v map[string]*string) *GetSuspEventPageResponse {
	s.Headers = v
	return s
}

func (s *GetSuspEventPageResponse) SetStatusCode(v int32) *GetSuspEventPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSuspEventPageResponse) SetBody(v *GetSuspEventPageResponseBody) *GetSuspEventPageResponse {
	s.Body = v
	return s
}

type GetSuspEventSummaryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// month
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1732156885986
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1729478485000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// example:
	//
	// SUSP_EVENT
	SuspEventSource *string `json:"SuspEventSource,omitempty" xml:"SuspEventSource,omitempty"`
}

func (s GetSuspEventSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryRequest) SetDateType(v string) *GetSuspEventSummaryRequest {
	s.DateType = &v
	return s
}

func (s *GetSuspEventSummaryRequest) SetEndDate(v int64) *GetSuspEventSummaryRequest {
	s.EndDate = &v
	return s
}

func (s *GetSuspEventSummaryRequest) SetStartDate(v int64) *GetSuspEventSummaryRequest {
	s.StartDate = &v
	return s
}

func (s *GetSuspEventSummaryRequest) SetSuspEventSource(v string) *GetSuspEventSummaryRequest {
	s.SuspEventSource = &v
	return s
}

type GetSuspEventSummaryResponseBody struct {
	// example:
	//
	// 200
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetSuspEventSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9B2DAE9B-B901-5818-AFEF-E5637D938280
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSuspEventSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBody) SetCode(v string) *GetSuspEventSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetSuspEventSummaryResponseBody) SetData(v *GetSuspEventSummaryResponseBodyData) *GetSuspEventSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetSuspEventSummaryResponseBody) SetHttpStatusCode(v int32) *GetSuspEventSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSuspEventSummaryResponseBody) SetMessage(v string) *GetSuspEventSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetSuspEventSummaryResponseBody) SetRequestId(v string) *GetSuspEventSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSuspEventSummaryResponseBody) SetSuccess(v bool) *GetSuspEventSummaryResponseBody {
	s.Success = &v
	return s
}

type GetSuspEventSummaryResponseBodyData struct {
	NetworkAttackTrendDTO   *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTO   `json:"NetworkAttackTrendDTO,omitempty" xml:"NetworkAttackTrendDTO,omitempty" type:"Struct"`
	SuspEventDealSummaryDTO *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO `json:"SuspEventDealSummaryDTO,omitempty" xml:"SuspEventDealSummaryDTO,omitempty" type:"Struct"`
	SuspEventTopDTO         *GetSuspEventSummaryResponseBodyDataSuspEventTopDTO         `json:"SuspEventTopDTO,omitempty" xml:"SuspEventTopDTO,omitempty" type:"Struct"`
	SuspEventTrendDTO       *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTO       `json:"SuspEventTrendDTO,omitempty" xml:"SuspEventTrendDTO,omitempty" type:"Struct"`
}

func (s GetSuspEventSummaryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBodyData) SetNetworkAttackTrendDTO(v *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTO) *GetSuspEventSummaryResponseBodyData {
	s.NetworkAttackTrendDTO = v
	return s
}

func (s *GetSuspEventSummaryResponseBodyData) SetSuspEventDealSummaryDTO(v *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) *GetSuspEventSummaryResponseBodyData {
	s.SuspEventDealSummaryDTO = v
	return s
}

func (s *GetSuspEventSummaryResponseBodyData) SetSuspEventTopDTO(v *GetSuspEventSummaryResponseBodyDataSuspEventTopDTO) *GetSuspEventSummaryResponseBodyData {
	s.SuspEventTopDTO = v
	return s
}

func (s *GetSuspEventSummaryResponseBodyData) SetSuspEventTrendDTO(v *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTO) *GetSuspEventSummaryResponseBodyData {
	s.SuspEventTrendDTO = v
	return s
}

type GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTO struct {
	TrendList []*GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList `json:"TrendList,omitempty" xml:"TrendList,omitempty" type:"Repeated"`
}

func (s GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTO) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTO) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTO) SetTrendList(v []*GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList) *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTO {
	s.TrendList = v
	return s
}

type GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList struct {
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// 10
	DdosCount *int64 `json:"DdosCount,omitempty" xml:"DdosCount,omitempty"`
	// example:
	//
	// 10
	EipCount *int64 `json:"EipCount,omitempty" xml:"EipCount,omitempty"`
	// example:
	//
	// 10
	WafCount *int64 `json:"WafCount,omitempty" xml:"WafCount,omitempty"`
}

func (s GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList) SetDate(v string) *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList {
	s.Date = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList) SetDdosCount(v int64) *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList {
	s.DdosCount = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList) SetEipCount(v int64) *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList {
	s.EipCount = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList) SetWafCount(v int64) *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList {
	s.WafCount = &v
	return s
}

type GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO struct {
	// example:
	//
	// 20
	CompletedCount *int64 `json:"CompletedCount,omitempty" xml:"CompletedCount,omitempty"`
	// example:
	//
	// 5
	HandingCount *int64 `json:"HandingCount,omitempty" xml:"HandingCount,omitempty"`
	// example:
	//
	// 90
	HandingRate *string `json:"HandingRate,omitempty" xml:"HandingRate,omitempty"`
	// example:
	//
	// 35
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 10
	TotalGrowthRate *string `json:"TotalGrowthRate,omitempty" xml:"TotalGrowthRate,omitempty"`
	// example:
	//
	// 10
	WaitHandleCount *int64 `json:"WaitHandleCount,omitempty" xml:"WaitHandleCount,omitempty"`
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) SetCompletedCount(v int64) *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO {
	s.CompletedCount = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) SetHandingCount(v int64) *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO {
	s.HandingCount = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) SetHandingRate(v string) *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO {
	s.HandingRate = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) SetTotalCount(v int64) *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO {
	s.TotalCount = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) SetTotalGrowthRate(v string) *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO {
	s.TotalGrowthRate = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) SetWaitHandleCount(v int64) *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO {
	s.WaitHandleCount = &v
	return s
}

type GetSuspEventSummaryResponseBodyDataSuspEventTopDTO struct {
	SuspEventList []*GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList `json:"SuspEventList,omitempty" xml:"SuspEventList,omitempty" type:"Repeated"`
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventTopDTO) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventTopDTO) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTopDTO) SetSuspEventList(v []*GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList) *GetSuspEventSummaryResponseBodyDataSuspEventTopDTO {
	s.SuspEventList = v
	return s
}

type GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList struct {
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// example:
	//
	// 7
	TaskCount *int64 `json:"TaskCount,omitempty" xml:"TaskCount,omitempty"`
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList) SetEventName(v string) *GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList {
	s.EventName = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList) SetTaskCount(v int64) *GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList {
	s.TaskCount = &v
	return s
}

type GetSuspEventSummaryResponseBodyDataSuspEventTrendDTO struct {
	TrendList []*GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList `json:"TrendList,omitempty" xml:"TrendList,omitempty" type:"Repeated"`
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventTrendDTO) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventTrendDTO) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTO) SetTrendList(v []*GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList) *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTO {
	s.TrendList = v
	return s
}

type GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList struct {
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// 10
	DealCount *int64 `json:"DealCount,omitempty" xml:"DealCount,omitempty"`
	// example:
	//
	// 15
	FindCount *int64 `json:"FindCount,omitempty" xml:"FindCount,omitempty"`
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList) SetDate(v string) *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList {
	s.Date = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList) SetDealCount(v int64) *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList {
	s.DealCount = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList) SetFindCount(v int64) *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList {
	s.FindCount = &v
	return s
}

type GetSuspEventSummaryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSuspEventSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSuspEventSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSuspEventSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponse) SetHeaders(v map[string]*string) *GetSuspEventSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetSuspEventSummaryResponse) SetStatusCode(v int32) *GetSuspEventSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSuspEventSummaryResponse) SetBody(v *GetSuspEventSummaryResponseBody) *GetSuspEventSummaryResponse {
	s.Body = v
	return s
}

type GetSuspPageSummaryResponseBody struct {
	Code           *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetSuspPageSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// This parameter is required.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSuspPageSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSuspPageSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetSuspPageSummaryResponseBody) SetCode(v string) *GetSuspPageSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetSuspPageSummaryResponseBody) SetData(v *GetSuspPageSummaryResponseBodyData) *GetSuspPageSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetSuspPageSummaryResponseBody) SetHttpStatusCode(v int32) *GetSuspPageSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSuspPageSummaryResponseBody) SetMessage(v string) *GetSuspPageSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetSuspPageSummaryResponseBody) SetRequestId(v string) *GetSuspPageSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSuspPageSummaryResponseBody) SetSuccess(v bool) *GetSuspPageSummaryResponseBody {
	s.Success = &v
	return s
}

type GetSuspPageSummaryResponseBodyData struct {
	CompletedCount  *int64 `json:"CompletedCount,omitempty" xml:"CompletedCount,omitempty"`
	HandingCount    *int64 `json:"HandingCount,omitempty" xml:"HandingCount,omitempty"`
	HighCount       *int64 `json:"HighCount,omitempty" xml:"HighCount,omitempty"`
	LowCount        *int64 `json:"LowCount,omitempty" xml:"LowCount,omitempty"`
	MediumCount     *int64 `json:"MediumCount,omitempty" xml:"MediumCount,omitempty"`
	TotalCount      *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	WaitHandleCount *int64 `json:"WaitHandleCount,omitempty" xml:"WaitHandleCount,omitempty"`
}

func (s GetSuspPageSummaryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSuspPageSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSuspPageSummaryResponseBodyData) SetCompletedCount(v int64) *GetSuspPageSummaryResponseBodyData {
	s.CompletedCount = &v
	return s
}

func (s *GetSuspPageSummaryResponseBodyData) SetHandingCount(v int64) *GetSuspPageSummaryResponseBodyData {
	s.HandingCount = &v
	return s
}

func (s *GetSuspPageSummaryResponseBodyData) SetHighCount(v int64) *GetSuspPageSummaryResponseBodyData {
	s.HighCount = &v
	return s
}

func (s *GetSuspPageSummaryResponseBodyData) SetLowCount(v int64) *GetSuspPageSummaryResponseBodyData {
	s.LowCount = &v
	return s
}

func (s *GetSuspPageSummaryResponseBodyData) SetMediumCount(v int64) *GetSuspPageSummaryResponseBodyData {
	s.MediumCount = &v
	return s
}

func (s *GetSuspPageSummaryResponseBodyData) SetTotalCount(v int64) *GetSuspPageSummaryResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetSuspPageSummaryResponseBodyData) SetWaitHandleCount(v int64) *GetSuspPageSummaryResponseBodyData {
	s.WaitHandleCount = &v
	return s
}

type GetSuspPageSummaryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSuspPageSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSuspPageSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSuspPageSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetSuspPageSummaryResponse) SetHeaders(v map[string]*string) *GetSuspPageSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetSuspPageSummaryResponse) SetStatusCode(v int32) *GetSuspPageSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSuspPageSummaryResponse) SetBody(v *GetSuspPageSummaryResponseBody) *GetSuspPageSummaryResponse {
	s.Body = v
	return s
}

type GetUserStatusResponseBody struct {
	// example:
	//
	// 200
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetUserStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// D8DBD769-613E-5E6B-A9FD-B622375B152D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserStatusResponseBody) SetCode(v string) *GetUserStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserStatusResponseBody) SetData(v *GetUserStatusResponseBodyData) *GetUserStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetUserStatusResponseBody) SetHttpStatusCode(v int32) *GetUserStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUserStatusResponseBody) SetMessage(v string) *GetUserStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserStatusResponseBody) SetRequestId(v string) *GetUserStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserStatusResponseBody) SetSuccess(v bool) *GetUserStatusResponseBody {
	s.Success = &v
	return s
}

type GetUserStatusResponseBodyData struct {
	// example:
	//
	// official
	CustomerType *string `json:"CustomerType,omitempty" xml:"CustomerType,omitempty"`
	// example:
	//
	// 2023-09-28 00:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 726cec3c-4887-4354-8c21-c0ad12e10fc2
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 2023-09-20 00:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// example:
	//
	// FirstLogin
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// mdrjichu
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetUserStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUserStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserStatusResponseBodyData) SetCustomerType(v string) *GetUserStatusResponseBodyData {
	s.CustomerType = &v
	return s
}

func (s *GetUserStatusResponseBodyData) SetEndDate(v string) *GetUserStatusResponseBodyData {
	s.EndDate = &v
	return s
}

func (s *GetUserStatusResponseBodyData) SetInstanceId(v string) *GetUserStatusResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetUserStatusResponseBodyData) SetStartDate(v string) *GetUserStatusResponseBodyData {
	s.StartDate = &v
	return s
}

func (s *GetUserStatusResponseBodyData) SetStatus(v string) *GetUserStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetUserStatusResponseBodyData) SetVersion(v string) *GetUserStatusResponseBodyData {
	s.Version = &v
	return s
}

type GetUserStatusResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserStatusResponse) GoString() string {
	return s.String()
}

func (s *GetUserStatusResponse) SetHeaders(v map[string]*string) *GetUserStatusResponse {
	s.Headers = v
	return s
}

func (s *GetUserStatusResponse) SetStatusCode(v int32) *GetUserStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserStatusResponse) SetBody(v *GetUserStatusResponseBody) *GetUserStatusResponse {
	s.Body = v
	return s
}

type GetVulItemPageRequest struct {
	// example:
	//
	// RHSA-2018:3665-Important: NetworkManager security update
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// n
	Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	// example:
	//
	// later
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// oval:com.redhat.rhsa:def:20183665
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// sca
	ScanType *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
}

func (s GetVulItemPageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVulItemPageRequest) GoString() string {
	return s.String()
}

func (s *GetVulItemPageRequest) SetAliasName(v string) *GetVulItemPageRequest {
	s.AliasName = &v
	return s
}

func (s *GetVulItemPageRequest) SetCurrentPage(v int32) *GetVulItemPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetVulItemPageRequest) SetDealed(v string) *GetVulItemPageRequest {
	s.Dealed = &v
	return s
}

func (s *GetVulItemPageRequest) SetLevel(v string) *GetVulItemPageRequest {
	s.Level = &v
	return s
}

func (s *GetVulItemPageRequest) SetName(v string) *GetVulItemPageRequest {
	s.Name = &v
	return s
}

func (s *GetVulItemPageRequest) SetPageSize(v int32) *GetVulItemPageRequest {
	s.PageSize = &v
	return s
}

func (s *GetVulItemPageRequest) SetScanType(v string) *GetVulItemPageRequest {
	s.ScanType = &v
	return s
}

type GetVulItemPageResponseBody struct {
	// example:
	//
	// 200
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetVulItemPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	PageInfo *GetVulItemPageResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 02F8BBF3-2D61-5982-8911-EEB387BE3AF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetVulItemPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVulItemPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetVulItemPageResponseBody) SetCode(v string) *GetVulItemPageResponseBody {
	s.Code = &v
	return s
}

func (s *GetVulItemPageResponseBody) SetData(v []*GetVulItemPageResponseBodyData) *GetVulItemPageResponseBody {
	s.Data = v
	return s
}

func (s *GetVulItemPageResponseBody) SetHttpStatusCode(v int32) *GetVulItemPageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetVulItemPageResponseBody) SetMessage(v string) *GetVulItemPageResponseBody {
	s.Message = &v
	return s
}

func (s *GetVulItemPageResponseBody) SetPageInfo(v *GetVulItemPageResponseBodyPageInfo) *GetVulItemPageResponseBody {
	s.PageInfo = v
	return s
}

func (s *GetVulItemPageResponseBody) SetRequestId(v string) *GetVulItemPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVulItemPageResponseBody) SetSuccess(v bool) *GetVulItemPageResponseBody {
	s.Success = &v
	return s
}

type GetVulItemPageResponseBodyData struct {
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// example:
	//
	// 74
	AsapCount *int32 `json:"AsapCount,omitempty" xml:"AsapCount,omitempty"`
	// example:
	//
	// 1940494487193744
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// example:
	//
	// https://avd.aliyun.com/detail/
	CveUrlPrefix *string `json:"CveUrlPrefix,omitempty" xml:"CveUrlPrefix,omitempty"`
	// example:
	//
	// y
	Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	// example:
	//
	// 2023-04-23 14:47:34
	FindTime *string `json:"FindTime,omitempty" xml:"FindTime,omitempty"`
	// example:
	//
	// 20
	HandledCount *int32 `json:"HandledCount,omitempty" xml:"HandledCount,omitempty"`
	// example:
	//
	// 353845
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 10
	LaterCount *int32 `json:"LaterCount,omitempty" xml:"LaterCount,omitempty"`
	// example:
	//
	// later
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// oval:com.redhat.rhsa:def:20205002
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 8
	NntfCount *int32 `json:"NntfCount,omitempty" xml:"NntfCount,omitempty"`
	// example:
	//
	// CVE-2019-20907
	Related *string `json:"Related,omitempty" xml:"Related,omitempty"`
	// example:
	//
	// 20
	RelatedCveCount *int32 `json:"RelatedCveCount,omitempty" xml:"RelatedCveCount,omitempty"`
	// example:
	//
	// sca
	ScanType *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
	// example:
	//
	// Elevation of Privilege
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// example:
	//
	// 50
	TotalFixCount *int64 `json:"TotalFixCount,omitempty" xml:"TotalFixCount,omitempty"`
}

func (s GetVulItemPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetVulItemPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVulItemPageResponseBodyData) SetAliasName(v string) *GetVulItemPageResponseBodyData {
	s.AliasName = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetAsapCount(v int32) *GetVulItemPageResponseBodyData {
	s.AsapCount = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetCustomerId(v string) *GetVulItemPageResponseBodyData {
	s.CustomerId = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetCveUrlPrefix(v string) *GetVulItemPageResponseBodyData {
	s.CveUrlPrefix = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetDealed(v string) *GetVulItemPageResponseBodyData {
	s.Dealed = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetFindTime(v string) *GetVulItemPageResponseBodyData {
	s.FindTime = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetHandledCount(v int32) *GetVulItemPageResponseBodyData {
	s.HandledCount = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetId(v int64) *GetVulItemPageResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetLaterCount(v int32) *GetVulItemPageResponseBodyData {
	s.LaterCount = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetLevel(v string) *GetVulItemPageResponseBodyData {
	s.Level = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetName(v string) *GetVulItemPageResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetNntfCount(v int32) *GetVulItemPageResponseBodyData {
	s.NntfCount = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetRelated(v string) *GetVulItemPageResponseBodyData {
	s.Related = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetRelatedCveCount(v int32) *GetVulItemPageResponseBodyData {
	s.RelatedCveCount = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetScanType(v string) *GetVulItemPageResponseBodyData {
	s.ScanType = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetTags(v string) *GetVulItemPageResponseBodyData {
	s.Tags = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetTotalFixCount(v int64) *GetVulItemPageResponseBodyData {
	s.TotalFixCount = &v
	return s
}

type GetVulItemPageResponseBodyPageInfo struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 163
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetVulItemPageResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s GetVulItemPageResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *GetVulItemPageResponseBodyPageInfo) SetCurrentPage(v int32) *GetVulItemPageResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *GetVulItemPageResponseBodyPageInfo) SetPageSize(v int32) *GetVulItemPageResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *GetVulItemPageResponseBodyPageInfo) SetTotalCount(v int32) *GetVulItemPageResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

type GetVulItemPageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVulItemPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVulItemPageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVulItemPageResponse) GoString() string {
	return s.String()
}

func (s *GetVulItemPageResponse) SetHeaders(v map[string]*string) *GetVulItemPageResponse {
	s.Headers = v
	return s
}

func (s *GetVulItemPageResponse) SetStatusCode(v int32) *GetVulItemPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVulItemPageResponse) SetBody(v *GetVulItemPageResponseBody) *GetVulItemPageResponse {
	s.Body = v
	return s
}

type GetVulPageSummaryResponseBody struct {
	// example:
	//
	// 200
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetVulPageSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// A3A575C8-80F9-5F04-AA24-CCAC246884A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetVulPageSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVulPageSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetVulPageSummaryResponseBody) SetCode(v string) *GetVulPageSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetVulPageSummaryResponseBody) SetData(v *GetVulPageSummaryResponseBodyData) *GetVulPageSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetVulPageSummaryResponseBody) SetHttpStatusCode(v int32) *GetVulPageSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetVulPageSummaryResponseBody) SetMessage(v string) *GetVulPageSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetVulPageSummaryResponseBody) SetRequestId(v string) *GetVulPageSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVulPageSummaryResponseBody) SetSuccess(v bool) *GetVulPageSummaryResponseBody {
	s.Success = &v
	return s
}

type GetVulPageSummaryResponseBodyData struct {
	// example:
	//
	// 1990
	CompletedCount *int64 `json:"CompletedCount,omitempty" xml:"CompletedCount,omitempty"`
	// example:
	//
	// 6
	HandingCount *int64 `json:"HandingCount,omitempty" xml:"HandingCount,omitempty"`
	// example:
	//
	// 500
	HighCount *int64 `json:"HighCount,omitempty" xml:"HighCount,omitempty"`
	// example:
	//
	// 1000
	LowCount *int64 `json:"LowCount,omitempty" xml:"LowCount,omitempty"`
	// example:
	//
	// 500
	MediumCount *int64 `json:"MediumCount,omitempty" xml:"MediumCount,omitempty"`
	// example:
	//
	// 2000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 4
	WaitHandleCount *int64 `json:"WaitHandleCount,omitempty" xml:"WaitHandleCount,omitempty"`
}

func (s GetVulPageSummaryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetVulPageSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVulPageSummaryResponseBodyData) SetCompletedCount(v int64) *GetVulPageSummaryResponseBodyData {
	s.CompletedCount = &v
	return s
}

func (s *GetVulPageSummaryResponseBodyData) SetHandingCount(v int64) *GetVulPageSummaryResponseBodyData {
	s.HandingCount = &v
	return s
}

func (s *GetVulPageSummaryResponseBodyData) SetHighCount(v int64) *GetVulPageSummaryResponseBodyData {
	s.HighCount = &v
	return s
}

func (s *GetVulPageSummaryResponseBodyData) SetLowCount(v int64) *GetVulPageSummaryResponseBodyData {
	s.LowCount = &v
	return s
}

func (s *GetVulPageSummaryResponseBodyData) SetMediumCount(v int64) *GetVulPageSummaryResponseBodyData {
	s.MediumCount = &v
	return s
}

func (s *GetVulPageSummaryResponseBodyData) SetTotalCount(v int64) *GetVulPageSummaryResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetVulPageSummaryResponseBodyData) SetWaitHandleCount(v int64) *GetVulPageSummaryResponseBodyData {
	s.WaitHandleCount = &v
	return s
}

type GetVulPageSummaryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVulPageSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVulPageSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVulPageSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetVulPageSummaryResponse) SetHeaders(v map[string]*string) *GetVulPageSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetVulPageSummaryResponse) SetStatusCode(v int32) *GetVulPageSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVulPageSummaryResponse) SetBody(v *GetVulPageSummaryResponseBody) *GetVulPageSummaryResponse {
	s.Body = v
	return s
}

type GetVulSummaryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// month
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1732156885986
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1729478485000
	StartDate       *int64  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	SuspEventSource *string `json:"SuspEventSource,omitempty" xml:"SuspEventSource,omitempty"`
}

func (s GetVulSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVulSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetVulSummaryRequest) SetDateType(v string) *GetVulSummaryRequest {
	s.DateType = &v
	return s
}

func (s *GetVulSummaryRequest) SetEndDate(v int64) *GetVulSummaryRequest {
	s.EndDate = &v
	return s
}

func (s *GetVulSummaryRequest) SetStartDate(v int64) *GetVulSummaryRequest {
	s.StartDate = &v
	return s
}

func (s *GetVulSummaryRequest) SetSuspEventSource(v string) *GetVulSummaryRequest {
	s.SuspEventSource = &v
	return s
}

type GetVulSummaryResponseBody struct {
	// example:
	//
	// OK
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetVulSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// system error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// EF801DD1-D934-51B3-92D4-776CE17B184F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetVulSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVulSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetVulSummaryResponseBody) SetCode(v string) *GetVulSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetVulSummaryResponseBody) SetData(v *GetVulSummaryResponseBodyData) *GetVulSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetVulSummaryResponseBody) SetHttpStatusCode(v int32) *GetVulSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetVulSummaryResponseBody) SetMessage(v string) *GetVulSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetVulSummaryResponseBody) SetRequestId(v string) *GetVulSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVulSummaryResponseBody) SetSuccess(v bool) *GetVulSummaryResponseBody {
	s.Success = &v
	return s
}

type GetVulSummaryResponseBodyData struct {
	// example:
	//
	// 10
	CompletedCount *int64 `json:"CompletedCount,omitempty" xml:"CompletedCount,omitempty"`
	// example:
	//
	// 50
	DealRate  *string                                   `json:"DealRate,omitempty" xml:"DealRate,omitempty"`
	TrendList []*GetVulSummaryResponseBodyDataTrendList `json:"TrendList,omitempty" xml:"TrendList,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	WaitHandleCount *int64 `json:"WaitHandleCount,omitempty" xml:"WaitHandleCount,omitempty"`
}

func (s GetVulSummaryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetVulSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVulSummaryResponseBodyData) SetCompletedCount(v int64) *GetVulSummaryResponseBodyData {
	s.CompletedCount = &v
	return s
}

func (s *GetVulSummaryResponseBodyData) SetDealRate(v string) *GetVulSummaryResponseBodyData {
	s.DealRate = &v
	return s
}

func (s *GetVulSummaryResponseBodyData) SetTrendList(v []*GetVulSummaryResponseBodyDataTrendList) *GetVulSummaryResponseBodyData {
	s.TrendList = v
	return s
}

func (s *GetVulSummaryResponseBodyData) SetWaitHandleCount(v int64) *GetVulSummaryResponseBodyData {
	s.WaitHandleCount = &v
	return s
}

type GetVulSummaryResponseBodyDataTrendList struct {
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// 10
	DealCount *int64 `json:"DealCount,omitempty" xml:"DealCount,omitempty"`
	// example:
	//
	// 15
	FindCount *int64 `json:"FindCount,omitempty" xml:"FindCount,omitempty"`
}

func (s GetVulSummaryResponseBodyDataTrendList) String() string {
	return tea.Prettify(s)
}

func (s GetVulSummaryResponseBodyDataTrendList) GoString() string {
	return s.String()
}

func (s *GetVulSummaryResponseBodyDataTrendList) SetDate(v string) *GetVulSummaryResponseBodyDataTrendList {
	s.Date = &v
	return s
}

func (s *GetVulSummaryResponseBodyDataTrendList) SetDealCount(v int64) *GetVulSummaryResponseBodyDataTrendList {
	s.DealCount = &v
	return s
}

func (s *GetVulSummaryResponseBodyDataTrendList) SetFindCount(v int64) *GetVulSummaryResponseBodyDataTrendList {
	s.FindCount = &v
	return s
}

type GetVulSummaryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVulSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVulSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVulSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetVulSummaryResponse) SetHeaders(v map[string]*string) *GetVulSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetVulSummaryResponse) SetStatusCode(v int32) *GetVulSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVulSummaryResponse) SetBody(v *GetVulSummaryResponseBody) *GetVulSummaryResponse {
	s.Body = v
	return s
}

type GetWorkTaskSummaryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// month
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1732156885986
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1729478485000
	StartDate       *int64  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	SuspEventSource *string `json:"SuspEventSource,omitempty" xml:"SuspEventSource,omitempty"`
}

func (s GetWorkTaskSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWorkTaskSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetWorkTaskSummaryRequest) SetDateType(v string) *GetWorkTaskSummaryRequest {
	s.DateType = &v
	return s
}

func (s *GetWorkTaskSummaryRequest) SetEndDate(v int64) *GetWorkTaskSummaryRequest {
	s.EndDate = &v
	return s
}

func (s *GetWorkTaskSummaryRequest) SetStartDate(v int64) *GetWorkTaskSummaryRequest {
	s.StartDate = &v
	return s
}

func (s *GetWorkTaskSummaryRequest) SetSuspEventSource(v string) *GetWorkTaskSummaryRequest {
	s.SuspEventSource = &v
	return s
}

type GetWorkTaskSummaryResponseBody struct {
	// example:
	//
	// OK
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetWorkTaskSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// EF801DD1-D934-51B3-92D4-776CE17B184F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWorkTaskSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkTaskSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkTaskSummaryResponseBody) SetCode(v string) *GetWorkTaskSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBody) SetData(v *GetWorkTaskSummaryResponseBodyData) *GetWorkTaskSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkTaskSummaryResponseBody) SetHttpStatusCode(v int32) *GetWorkTaskSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBody) SetMessage(v string) *GetWorkTaskSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBody) SetRequestId(v string) *GetWorkTaskSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBody) SetSuccess(v bool) *GetWorkTaskSummaryResponseBody {
	s.Success = &v
	return s
}

type GetWorkTaskSummaryResponseBodyData struct {
	// example:
	//
	// 60
	DealAverageDuration *int64 `json:"DealAverageDuration,omitempty" xml:"DealAverageDuration,omitempty"`
	// example:
	//
	// 20
	DealAverageDurationGrowthRate *string `json:"DealAverageDurationGrowthRate,omitempty" xml:"DealAverageDurationGrowthRate,omitempty"`
	// example:
	//
	// 100
	DealWorkTaskCount *int64 `json:"DealWorkTaskCount,omitempty" xml:"DealWorkTaskCount,omitempty"`
	// example:
	//
	// 20
	DealWorkTaskCountRate *string `json:"DealWorkTaskCountRate,omitempty" xml:"DealWorkTaskCountRate,omitempty"`
	// example:
	//
	// 10
	WorkTaskCount *int64 `json:"WorkTaskCount,omitempty" xml:"WorkTaskCount,omitempty"`
	// example:
	//
	// 90
	WorkTaskDealRate *string `json:"WorkTaskDealRate,omitempty" xml:"WorkTaskDealRate,omitempty"`
	// example:
	//
	// 20
	WorkTaskDealRateGrowthRate *string `json:"WorkTaskDealRateGrowthRate,omitempty" xml:"WorkTaskDealRateGrowthRate,omitempty"`
	// example:
	//
	// 20
	WorkTaskGrowthRate *string `json:"WorkTaskGrowthRate,omitempty" xml:"WorkTaskGrowthRate,omitempty"`
}

func (s GetWorkTaskSummaryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetWorkTaskSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkTaskSummaryResponseBodyData) SetDealAverageDuration(v int64) *GetWorkTaskSummaryResponseBodyData {
	s.DealAverageDuration = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBodyData) SetDealAverageDurationGrowthRate(v string) *GetWorkTaskSummaryResponseBodyData {
	s.DealAverageDurationGrowthRate = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBodyData) SetDealWorkTaskCount(v int64) *GetWorkTaskSummaryResponseBodyData {
	s.DealWorkTaskCount = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBodyData) SetDealWorkTaskCountRate(v string) *GetWorkTaskSummaryResponseBodyData {
	s.DealWorkTaskCountRate = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBodyData) SetWorkTaskCount(v int64) *GetWorkTaskSummaryResponseBodyData {
	s.WorkTaskCount = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBodyData) SetWorkTaskDealRate(v string) *GetWorkTaskSummaryResponseBodyData {
	s.WorkTaskDealRate = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBodyData) SetWorkTaskDealRateGrowthRate(v string) *GetWorkTaskSummaryResponseBodyData {
	s.WorkTaskDealRateGrowthRate = &v
	return s
}

func (s *GetWorkTaskSummaryResponseBodyData) SetWorkTaskGrowthRate(v string) *GetWorkTaskSummaryResponseBodyData {
	s.WorkTaskGrowthRate = &v
	return s
}

type GetWorkTaskSummaryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkTaskSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkTaskSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkTaskSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetWorkTaskSummaryResponse) SetHeaders(v map[string]*string) *GetWorkTaskSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetWorkTaskSummaryResponse) SetStatusCode(v int32) *GetWorkTaskSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkTaskSummaryResponse) SetBody(v *GetWorkTaskSummaryResponseBody) *GetWorkTaskSummaryResponse {
	s.Body = v
	return s
}

type PageServiceCustomerRequest struct {
	AuthStatus   *int32 `json:"AuthStatus,omitempty" xml:"AuthStatus,omitempty"`
	CmAuthStatus *int32 `json:"CmAuthStatus,omitempty" xml:"CmAuthStatus,omitempty"`
	// This parameter is required.
	CurrentPage       *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime           *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MonitorAuthStatus *int32 `json:"MonitorAuthStatus,omitempty" xml:"MonitorAuthStatus,omitempty"`
	// This parameter is required.
	PageSize  *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s PageServiceCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s PageServiceCustomerRequest) GoString() string {
	return s.String()
}

func (s *PageServiceCustomerRequest) SetAuthStatus(v int32) *PageServiceCustomerRequest {
	s.AuthStatus = &v
	return s
}

func (s *PageServiceCustomerRequest) SetCmAuthStatus(v int32) *PageServiceCustomerRequest {
	s.CmAuthStatus = &v
	return s
}

func (s *PageServiceCustomerRequest) SetCurrentPage(v int32) *PageServiceCustomerRequest {
	s.CurrentPage = &v
	return s
}

func (s *PageServiceCustomerRequest) SetEndTime(v int64) *PageServiceCustomerRequest {
	s.EndTime = &v
	return s
}

func (s *PageServiceCustomerRequest) SetMonitorAuthStatus(v int32) *PageServiceCustomerRequest {
	s.MonitorAuthStatus = &v
	return s
}

func (s *PageServiceCustomerRequest) SetPageSize(v int32) *PageServiceCustomerRequest {
	s.PageSize = &v
	return s
}

func (s *PageServiceCustomerRequest) SetStartTime(v int64) *PageServiceCustomerRequest {
	s.StartTime = &v
	return s
}

type PageServiceCustomerResponseBody struct {
	Code           *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*PageServiceCustomerResponseBodyData   `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	PageInfo       *PageServiceCustomerResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// Id of the request
	//
	// This parameter is required.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PageServiceCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageServiceCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *PageServiceCustomerResponseBody) SetCode(v string) *PageServiceCustomerResponseBody {
	s.Code = &v
	return s
}

func (s *PageServiceCustomerResponseBody) SetData(v []*PageServiceCustomerResponseBodyData) *PageServiceCustomerResponseBody {
	s.Data = v
	return s
}

func (s *PageServiceCustomerResponseBody) SetHttpStatusCode(v int32) *PageServiceCustomerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PageServiceCustomerResponseBody) SetMessage(v string) *PageServiceCustomerResponseBody {
	s.Message = &v
	return s
}

func (s *PageServiceCustomerResponseBody) SetPageInfo(v *PageServiceCustomerResponseBodyPageInfo) *PageServiceCustomerResponseBody {
	s.PageInfo = v
	return s
}

func (s *PageServiceCustomerResponseBody) SetRequestId(v string) *PageServiceCustomerResponseBody {
	s.RequestId = &v
	return s
}

func (s *PageServiceCustomerResponseBody) SetSuccess(v bool) *PageServiceCustomerResponseBody {
	s.Success = &v
	return s
}

type PageServiceCustomerResponseBodyData struct {
	Aliuid            *string `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	AuthStatus        *int32  `json:"AuthStatus,omitempty" xml:"AuthStatus,omitempty"`
	CmAuthStatus      *int32  `json:"CmAuthStatus,omitempty" xml:"CmAuthStatus,omitempty"`
	EndTime           *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Level             *string `json:"Level,omitempty" xml:"Level,omitempty"`
	MonitorAuthStatus *int32  `json:"MonitorAuthStatus,omitempty" xml:"MonitorAuthStatus,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnId             *string `json:"OwnId,omitempty" xml:"OwnId,omitempty"`
	StartTime         *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UserId            *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Version           *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s PageServiceCustomerResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PageServiceCustomerResponseBodyData) GoString() string {
	return s.String()
}

func (s *PageServiceCustomerResponseBodyData) SetAliuid(v string) *PageServiceCustomerResponseBodyData {
	s.Aliuid = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetAuthStatus(v int32) *PageServiceCustomerResponseBodyData {
	s.AuthStatus = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetCmAuthStatus(v int32) *PageServiceCustomerResponseBodyData {
	s.CmAuthStatus = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetEndTime(v int64) *PageServiceCustomerResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetLevel(v string) *PageServiceCustomerResponseBodyData {
	s.Level = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetMonitorAuthStatus(v int32) *PageServiceCustomerResponseBodyData {
	s.MonitorAuthStatus = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetName(v string) *PageServiceCustomerResponseBodyData {
	s.Name = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetOwnId(v string) *PageServiceCustomerResponseBodyData {
	s.OwnId = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetStartTime(v int64) *PageServiceCustomerResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetUserId(v string) *PageServiceCustomerResponseBodyData {
	s.UserId = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetVersion(v string) *PageServiceCustomerResponseBodyData {
	s.Version = &v
	return s
}

type PageServiceCustomerResponseBodyPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s PageServiceCustomerResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s PageServiceCustomerResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *PageServiceCustomerResponseBodyPageInfo) SetCurrentPage(v int32) *PageServiceCustomerResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *PageServiceCustomerResponseBodyPageInfo) SetPageSize(v int32) *PageServiceCustomerResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *PageServiceCustomerResponseBodyPageInfo) SetTotalCount(v int32) *PageServiceCustomerResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

type PageServiceCustomerResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PageServiceCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PageServiceCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s PageServiceCustomerResponse) GoString() string {
	return s.String()
}

func (s *PageServiceCustomerResponse) SetHeaders(v map[string]*string) *PageServiceCustomerResponse {
	s.Headers = v
	return s
}

func (s *PageServiceCustomerResponse) SetStatusCode(v int32) *PageServiceCustomerResponse {
	s.StatusCode = &v
	return s
}

func (s *PageServiceCustomerResponse) SetBody(v *PageServiceCustomerResponseBody) *PageServiceCustomerResponse {
	s.Body = v
	return s
}

type SendCustomEventRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1214484929940219
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// example:
	//
	// aegis_suspicious_event
	DataSource       *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	EventDescription *string `json:"EventDescription,omitempty" xml:"EventDescription,omitempty"`
	EventDetails     *string `json:"EventDetails,omitempty" xml:"EventDetails,omitempty"`
	// This parameter is required.
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SUSP_CUSTOM_CFW
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-04-23 14:47:34
	FindTime *int64 `json:"FindTime,omitempty" xml:"FindTime,omitempty"`
	// example:
	//
	// i-uf60h3ns25bzq9eyf8ps
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	IsSend *string `json:"IsSend,omitempty" xml:"IsSend,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// serious
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// 1724956996000
	OccurrenceTime *int64  `json:"OccurrenceTime,omitempty" xml:"OccurrenceTime,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// CloudSecCenter
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 68888f02-98f2-492b-a2b2-5b13295755b7
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	// UUID。
	//
	// example:
	//
	// 93B6CDAB-7D2E-33D2-9EBA-25D561A2E95F
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s SendCustomEventRequest) String() string {
	return tea.Prettify(s)
}

func (s SendCustomEventRequest) GoString() string {
	return s.String()
}

func (s *SendCustomEventRequest) SetCustomerId(v string) *SendCustomEventRequest {
	s.CustomerId = &v
	return s
}

func (s *SendCustomEventRequest) SetDataSource(v string) *SendCustomEventRequest {
	s.DataSource = &v
	return s
}

func (s *SendCustomEventRequest) SetEventDescription(v string) *SendCustomEventRequest {
	s.EventDescription = &v
	return s
}

func (s *SendCustomEventRequest) SetEventDetails(v string) *SendCustomEventRequest {
	s.EventDetails = &v
	return s
}

func (s *SendCustomEventRequest) SetEventName(v string) *SendCustomEventRequest {
	s.EventName = &v
	return s
}

func (s *SendCustomEventRequest) SetEventType(v string) *SendCustomEventRequest {
	s.EventType = &v
	return s
}

func (s *SendCustomEventRequest) SetFindTime(v int64) *SendCustomEventRequest {
	s.FindTime = &v
	return s
}

func (s *SendCustomEventRequest) SetInstanceId(v string) *SendCustomEventRequest {
	s.InstanceId = &v
	return s
}

func (s *SendCustomEventRequest) SetInstanceName(v string) *SendCustomEventRequest {
	s.InstanceName = &v
	return s
}

func (s *SendCustomEventRequest) SetIsSend(v string) *SendCustomEventRequest {
	s.IsSend = &v
	return s
}

func (s *SendCustomEventRequest) SetLevel(v string) *SendCustomEventRequest {
	s.Level = &v
	return s
}

func (s *SendCustomEventRequest) SetOccurrenceTime(v int64) *SendCustomEventRequest {
	s.OccurrenceTime = &v
	return s
}

func (s *SendCustomEventRequest) SetOwnerId(v string) *SendCustomEventRequest {
	s.OwnerId = &v
	return s
}

func (s *SendCustomEventRequest) SetProduct(v string) *SendCustomEventRequest {
	s.Product = &v
	return s
}

func (s *SendCustomEventRequest) SetUniqueId(v string) *SendCustomEventRequest {
	s.UniqueId = &v
	return s
}

func (s *SendCustomEventRequest) SetUuid(v string) *SendCustomEventRequest {
	s.Uuid = &v
	return s
}

type SendCustomEventResponseBody struct {
	// example:
	//
	// 200
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SendCustomEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 606EB377-155D-5AEB-AC4F-F013444A4C45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendCustomEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendCustomEventResponseBody) GoString() string {
	return s.String()
}

func (s *SendCustomEventResponseBody) SetCode(v string) *SendCustomEventResponseBody {
	s.Code = &v
	return s
}

func (s *SendCustomEventResponseBody) SetData(v *SendCustomEventResponseBodyData) *SendCustomEventResponseBody {
	s.Data = v
	return s
}

func (s *SendCustomEventResponseBody) SetHttpStatusCode(v int32) *SendCustomEventResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SendCustomEventResponseBody) SetMessage(v string) *SendCustomEventResponseBody {
	s.Message = &v
	return s
}

func (s *SendCustomEventResponseBody) SetRequestId(v string) *SendCustomEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendCustomEventResponseBody) SetSuccess(v bool) *SendCustomEventResponseBody {
	s.Success = &v
	return s
}

type SendCustomEventResponseBodyData struct {
	// example:
	//
	// 1601097845544644
	CustomerId   *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	CustomerName *string `json:"CustomerName,omitempty" xml:"CustomerName,omitempty"`
	// example:
	//
	// c0dc71d1-8a1d-4043-9767-f6c420e34901-81bd
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// example:
	//
	// SUSP_CUSTOM_WAF
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// example:
	//
	// 1914348
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 352675
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	OwnerName    *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	WorkTaskName *string `json:"WorkTaskName,omitempty" xml:"WorkTaskName,omitempty"`
}

func (s SendCustomEventResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SendCustomEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendCustomEventResponseBodyData) SetCustomerId(v string) *SendCustomEventResponseBodyData {
	s.CustomerId = &v
	return s
}

func (s *SendCustomEventResponseBodyData) SetCustomerName(v string) *SendCustomEventResponseBodyData {
	s.CustomerName = &v
	return s
}

func (s *SendCustomEventResponseBodyData) SetEventId(v string) *SendCustomEventResponseBodyData {
	s.EventId = &v
	return s
}

func (s *SendCustomEventResponseBodyData) SetEventType(v string) *SendCustomEventResponseBodyData {
	s.EventType = &v
	return s
}

func (s *SendCustomEventResponseBodyData) SetId(v int64) *SendCustomEventResponseBodyData {
	s.Id = &v
	return s
}

func (s *SendCustomEventResponseBodyData) SetOwnerId(v string) *SendCustomEventResponseBodyData {
	s.OwnerId = &v
	return s
}

func (s *SendCustomEventResponseBodyData) SetOwnerName(v string) *SendCustomEventResponseBodyData {
	s.OwnerName = &v
	return s
}

func (s *SendCustomEventResponseBodyData) SetWorkTaskName(v string) *SendCustomEventResponseBodyData {
	s.WorkTaskName = &v
	return s
}

type SendCustomEventResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendCustomEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendCustomEventResponse) String() string {
	return tea.Prettify(s)
}

func (s SendCustomEventResponse) GoString() string {
	return s.String()
}

func (s *SendCustomEventResponse) SetHeaders(v map[string]*string) *SendCustomEventResponse {
	s.Headers = v
	return s
}

func (s *SendCustomEventResponse) SetStatusCode(v int32) *SendCustomEventResponse {
	s.StatusCode = &v
	return s
}

func (s *SendCustomEventResponse) SetBody(v *SendCustomEventResponseBody) *SendCustomEventResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("mssp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建服务工单
//
// @param request - CreateServiceWorkOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceWorkOrderResponse
func (client *Client) CreateServiceWorkOrderWithOptions(request *CreateServiceWorkOrderRequest, runtime *util.RuntimeOptions) (_result *CreateServiceWorkOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Creator)) {
		body["Creator"] = request.Creator
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerId)) {
		body["CustomerId"] = request.CustomerId
	}

	if !tea.BoolValue(util.IsUnset(request.DurationDay)) {
		body["DurationDay"] = request.DurationDay
	}

	if !tea.BoolValue(util.IsUnset(request.IsAttachment)) {
		body["IsAttachment"] = request.IsAttachment
	}

	if !tea.BoolValue(util.IsUnset(request.IsWorkOrderNotify)) {
		body["IsWorkOrderNotify"] = request.IsWorkOrderNotify
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyDay)) {
		body["NotifyDay"] = request.NotifyDay
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyId)) {
		body["NotifyId"] = request.NotifyId
	}

	if !tea.BoolValue(util.IsUnset(request.OperateRemark)) {
		body["OperateRemark"] = request.OperateRemark
	}

	if !tea.BoolValue(util.IsUnset(request.OperateType)) {
		body["OperateType"] = request.OperateType
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["Operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		body["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.WorkOrderDetail)) {
		body["WorkOrderDetail"] = request.WorkOrderDetail
	}

	if !tea.BoolValue(util.IsUnset(request.WorkOrderName)) {
		body["WorkOrderName"] = request.WorkOrderName
	}

	if !tea.BoolValue(util.IsUnset(request.WorkOrderSource)) {
		body["WorkOrderSource"] = request.WorkOrderSource
	}

	if !tea.BoolValue(util.IsUnset(request.WorkOrderStatus)) {
		body["WorkOrderStatus"] = request.WorkOrderStatus
	}

	if !tea.BoolValue(util.IsUnset(request.WorkOrderType)) {
		body["WorkOrderType"] = request.WorkOrderType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceWorkOrder"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceWorkOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建服务工单
//
// @param request - CreateServiceWorkOrderRequest
//
// @return CreateServiceWorkOrderResponse
func (client *Client) CreateServiceWorkOrder(request *CreateServiceWorkOrderRequest) (_result *CreateServiceWorkOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceWorkOrderResponse{}
	_body, _err := client.CreateServiceWorkOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 处理服务工单
//
// @param request - DisposeServiceWorkOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisposeServiceWorkOrderResponse
func (client *Client) DisposeServiceWorkOrderWithOptions(request *DisposeServiceWorkOrderRequest, runtime *util.RuntimeOptions) (_result *DisposeServiceWorkOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttachmentName)) {
		body["AttachmentName"] = request.AttachmentName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ForwardOwnerId)) {
		body["ForwardOwnerId"] = request.ForwardOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.IsAttachment)) {
		body["IsAttachment"] = request.IsAttachment
	}

	if !tea.BoolValue(util.IsUnset(request.IsWorkOrderNotify)) {
		body["IsWorkOrderNotify"] = request.IsWorkOrderNotify
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyId)) {
		body["NotifyId"] = request.NotifyId
	}

	if !tea.BoolValue(util.IsUnset(request.OperateRemark)) {
		body["OperateRemark"] = request.OperateRemark
	}

	if !tea.BoolValue(util.IsUnset(request.OperateType)) {
		body["OperateType"] = request.OperateType
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["Operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UpgradeOwnerId)) {
		body["UpgradeOwnerId"] = request.UpgradeOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkOrderDetail)) {
		body["WorkOrderDetail"] = request.WorkOrderDetail
	}

	if !tea.BoolValue(util.IsUnset(request.WorkOrderName)) {
		body["WorkOrderName"] = request.WorkOrderName
	}

	if !tea.BoolValue(util.IsUnset(request.WorkOrderStatus)) {
		body["WorkOrderStatus"] = request.WorkOrderStatus
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DisposeServiceWorkOrder"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisposeServiceWorkOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 处理服务工单
//
// @param request - DisposeServiceWorkOrderRequest
//
// @return DisposeServiceWorkOrderResponse
func (client *Client) DisposeServiceWorkOrder(request *DisposeServiceWorkOrderRequest) (_result *DisposeServiceWorkOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisposeServiceWorkOrderResponse{}
	_body, _err := client.DisposeServiceWorkOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 处理告警工单
//
// @param request - DisposeWorkTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisposeWorkTaskResponse
func (client *Client) DisposeWorkTaskWithOptions(request *DisposeWorkTaskRequest, runtime *util.RuntimeOptions) (_result *DisposeWorkTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["Operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.OptRemark)) {
		body["OptRemark"] = request.OptRemark
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TaskIds)) {
		body["TaskIds"] = request.TaskIds
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DisposeWorkTask"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisposeWorkTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 处理告警工单
//
// @param request - DisposeWorkTaskRequest
//
// @return DisposeWorkTaskResponse
func (client *Client) DisposeWorkTask(request *DisposeWorkTaskRequest) (_result *DisposeWorkTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisposeWorkTaskResponse{}
	_body, _err := client.DisposeWorkTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 被攻击资产收敛趋势
//
// @param request - GetAttackedAssetDealRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAttackedAssetDealResponse
func (client *Client) GetAttackedAssetDealWithOptions(request *GetAttackedAssetDealRequest, runtime *util.RuntimeOptions) (_result *GetAttackedAssetDealResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DateType)) {
		body["DateType"] = request.DateType
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.SuspEventSource)) {
		body["SuspEventSource"] = request.SuspEventSource
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAttackedAssetDeal"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAttackedAssetDealResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 被攻击资产收敛趋势
//
// @param request - GetAttackedAssetDealRequest
//
// @return GetAttackedAssetDealResponse
func (client *Client) GetAttackedAssetDeal(request *GetAttackedAssetDealRequest) (_result *GetAttackedAssetDealResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAttackedAssetDealResponse{}
	_body, _err := client.GetAttackedAssetDealWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 合规风险收敛趋势
//
// @param request - GetBaselineSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBaselineSummaryResponse
func (client *Client) GetBaselineSummaryWithOptions(request *GetBaselineSummaryRequest, runtime *util.RuntimeOptions) (_result *GetBaselineSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DateType)) {
		body["DateType"] = request.DateType
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.SuspEventSource)) {
		body["SuspEventSource"] = request.SuspEventSource
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBaselineSummary"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBaselineSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 合规风险收敛趋势
//
// @param request - GetBaselineSummaryRequest
//
// @return GetBaselineSummaryResponse
func (client *Client) GetBaselineSummary(request *GetBaselineSummaryRequest) (_result *GetBaselineSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBaselineSummaryResponse{}
	_body, _err := client.GetBaselineSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询风险详情
//
// @param request - GetDetailByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDetailByIdResponse
func (client *Client) GetDetailByIdWithOptions(request *GetDetailByIdRequest, runtime *util.RuntimeOptions) (_result *GetDetailByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDetailById"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDetailByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询风险详情
//
// @param request - GetDetailByIdRequest
//
// @return GetDetailByIdResponse
func (client *Client) GetDetailById(request *GetDetailByIdRequest) (_result *GetDetailByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDetailByIdResponse{}
	_body, _err := client.GetDetailByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务报告单个下载
//
// @param request - GetDocumentDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentDownloadUrlResponse
func (client *Client) GetDocumentDownloadUrlWithOptions(request *GetDocumentDownloadUrlRequest, runtime *util.RuntimeOptions) (_result *GetDocumentDownloadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDocumentDownloadUrl"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDocumentDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务报告单个下载
//
// @param request - GetDocumentDownloadUrlRequest
//
// @return GetDocumentDownloadUrlResponse
func (client *Client) GetDocumentDownloadUrl(request *GetDocumentDownloadUrlRequest) (_result *GetDocumentDownloadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDocumentDownloadUrlResponse{}
	_body, _err := client.GetDocumentDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务报告查询
//
// @param request - GetDocumentPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentPageResponse
func (client *Client) GetDocumentPageWithOptions(request *GetDocumentPageRequest, runtime *util.RuntimeOptions) (_result *GetDocumentPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveredBy)) {
		body["DeliveredBy"] = request.DeliveredBy
	}

	if !tea.BoolValue(util.IsUnset(request.DocumentName)) {
		body["DocumentName"] = request.DocumentName
	}

	if !tea.BoolValue(util.IsUnset(request.DocumentType)) {
		body["DocumentType"] = request.DocumentType
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ReportType)) {
		body["ReportType"] = request.ReportType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDocumentPage"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDocumentPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务报告查询
//
// @param request - GetDocumentPageRequest
//
// @return GetDocumentPageResponse
func (client *Client) GetDocumentPage(request *GetDocumentPageRequest) (_result *GetDocumentPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDocumentPageResponse{}
	_body, _err := client.GetDocumentPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务报告首页统计项获取
//
// @param request - GetDocumentSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentSummaryResponse
func (client *Client) GetDocumentSummaryWithOptions(request *GetDocumentSummaryRequest, runtime *util.RuntimeOptions) (_result *GetDocumentSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReportType)) {
		body["ReportType"] = request.ReportType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDocumentSummary"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDocumentSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务报告首页统计项获取
//
// @param request - GetDocumentSummaryRequest
//
// @return GetDocumentSummaryResponse
func (client *Client) GetDocumentSummary(request *GetDocumentSummaryRequest) (_result *GetDocumentSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDocumentSummaryResponse{}
	_body, _err := client.GetDocumentSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 得到最近上传的服务报告
//
// @param request - GetRecentDocumentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRecentDocumentResponse
func (client *Client) GetRecentDocumentWithOptions(request *GetRecentDocumentRequest, runtime *util.RuntimeOptions) (_result *GetRecentDocumentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DateType)) {
		body["DateType"] = request.DateType
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.SuspEventSource)) {
		body["SuspEventSource"] = request.SuspEventSource
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRecentDocument"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRecentDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 得到最近上传的服务报告
//
// @param request - GetRecentDocumentRequest
//
// @return GetRecentDocumentResponse
func (client *Client) GetRecentDocument(request *GetRecentDocumentRequest) (_result *GetRecentDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRecentDocumentResponse{}
	_body, _err := client.GetRecentDocumentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 得到安全防护覆盖度
//
// @param request - GetSafetyCoverRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSafetyCoverResponse
func (client *Client) GetSafetyCoverWithOptions(request *GetSafetyCoverRequest, runtime *util.RuntimeOptions) (_result *GetSafetyCoverResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DateType)) {
		body["DateType"] = request.DateType
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.SuspEventSource)) {
		body["SuspEventSource"] = request.SuspEventSource
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSafetyCover"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSafetyCoverResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 得到安全防护覆盖度
//
// @param request - GetSafetyCoverRequest
//
// @return GetSafetyCoverResponse
func (client *Client) GetSafetyCover(request *GetSafetyCoverRequest) (_result *GetSafetyCoverResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSafetyCoverResponse{}
	_body, _err := client.GetSafetyCoverWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 告警处置查询
//
// @param request - GetSuspEventPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSuspEventPageResponse
func (client *Client) GetSuspEventPageWithOptions(request *GetSuspEventPageRequest, runtime *util.RuntimeOptions) (_result *GetSuspEventPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmEndTime)) {
		body["AlarmEndTime"] = request.AlarmEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.AlarmStartTime)) {
		body["AlarmStartTime"] = request.AlarmStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSuspEventPage"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSuspEventPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 告警处置查询
//
// @param request - GetSuspEventPageRequest
//
// @return GetSuspEventPageResponse
func (client *Client) GetSuspEventPage(request *GetSuspEventPageRequest) (_result *GetSuspEventPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSuspEventPageResponse{}
	_body, _err := client.GetSuspEventPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 得到告警的统计项
//
// @param request - GetSuspEventSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSuspEventSummaryResponse
func (client *Client) GetSuspEventSummaryWithOptions(request *GetSuspEventSummaryRequest, runtime *util.RuntimeOptions) (_result *GetSuspEventSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DateType)) {
		body["DateType"] = request.DateType
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.SuspEventSource)) {
		body["SuspEventSource"] = request.SuspEventSource
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSuspEventSummary"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSuspEventSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 得到告警的统计项
//
// @param request - GetSuspEventSummaryRequest
//
// @return GetSuspEventSummaryResponse
func (client *Client) GetSuspEventSummary(request *GetSuspEventSummaryRequest) (_result *GetSuspEventSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSuspEventSummaryResponse{}
	_body, _err := client.GetSuspEventSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 告警页统计
//
// @param request - GetSuspPageSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSuspPageSummaryResponse
func (client *Client) GetSuspPageSummaryWithOptions(runtime *util.RuntimeOptions) (_result *GetSuspPageSummaryResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetSuspPageSummary"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSuspPageSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 告警页统计
//
// @return GetSuspPageSummaryResponse
func (client *Client) GetSuspPageSummary() (_result *GetSuspPageSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSuspPageSummaryResponse{}
	_body, _err := client.GetSuspPageSummaryWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户开通状态
//
// @param request - GetUserStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserStatusResponse
func (client *Client) GetUserStatusWithOptions(runtime *util.RuntimeOptions) (_result *GetUserStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetUserStatus"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户开通状态
//
// @return GetUserStatusResponse
func (client *Client) GetUserStatus() (_result *GetUserStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserStatusResponse{}
	_body, _err := client.GetUserStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 风险查询
//
// @param request - GetVulItemPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVulItemPageResponse
func (client *Client) GetVulItemPageWithOptions(request *GetVulItemPageRequest, runtime *util.RuntimeOptions) (_result *GetVulItemPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliasName)) {
		body["AliasName"] = request.AliasName
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Dealed)) {
		body["Dealed"] = request.Dealed
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		body["Level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ScanType)) {
		body["ScanType"] = request.ScanType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVulItemPage"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVulItemPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 风险查询
//
// @param request - GetVulItemPageRequest
//
// @return GetVulItemPageResponse
func (client *Client) GetVulItemPage(request *GetVulItemPageRequest) (_result *GetVulItemPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVulItemPageResponse{}
	_body, _err := client.GetVulItemPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 风险页统计
//
// @param request - GetVulPageSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVulPageSummaryResponse
func (client *Client) GetVulPageSummaryWithOptions(runtime *util.RuntimeOptions) (_result *GetVulPageSummaryResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetVulPageSummary"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVulPageSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 风险页统计
//
// @return GetVulPageSummaryResponse
func (client *Client) GetVulPageSummary() (_result *GetVulPageSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVulPageSummaryResponse{}
	_body, _err := client.GetVulPageSummaryWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 得到风险的统计项
//
// @param request - GetVulSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVulSummaryResponse
func (client *Client) GetVulSummaryWithOptions(request *GetVulSummaryRequest, runtime *util.RuntimeOptions) (_result *GetVulSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DateType)) {
		body["DateType"] = request.DateType
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.SuspEventSource)) {
		body["SuspEventSource"] = request.SuspEventSource
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVulSummary"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVulSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 得到风险的统计项
//
// @param request - GetVulSummaryRequest
//
// @return GetVulSummaryResponse
func (client *Client) GetVulSummary(request *GetVulSummaryRequest) (_result *GetVulSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVulSummaryResponse{}
	_body, _err := client.GetVulSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 得到首行工单的统计项
//
// @param request - GetWorkTaskSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkTaskSummaryResponse
func (client *Client) GetWorkTaskSummaryWithOptions(request *GetWorkTaskSummaryRequest, runtime *util.RuntimeOptions) (_result *GetWorkTaskSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DateType)) {
		body["DateType"] = request.DateType
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.SuspEventSource)) {
		body["SuspEventSource"] = request.SuspEventSource
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWorkTaskSummary"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWorkTaskSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 得到首行工单的统计项
//
// @param request - GetWorkTaskSummaryRequest
//
// @return GetWorkTaskSummaryResponse
func (client *Client) GetWorkTaskSummary(request *GetWorkTaskSummaryRequest) (_result *GetWorkTaskSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWorkTaskSummaryResponse{}
	_body, _err := client.GetWorkTaskSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务客户信息查询
//
// @param request - PageServiceCustomerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PageServiceCustomerResponse
func (client *Client) PageServiceCustomerWithOptions(request *PageServiceCustomerRequest, runtime *util.RuntimeOptions) (_result *PageServiceCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthStatus)) {
		body["AuthStatus"] = request.AuthStatus
	}

	if !tea.BoolValue(util.IsUnset(request.CmAuthStatus)) {
		body["CmAuthStatus"] = request.CmAuthStatus
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MonitorAuthStatus)) {
		body["MonitorAuthStatus"] = request.MonitorAuthStatus
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PageServiceCustomer"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PageServiceCustomerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务客户信息查询
//
// @param request - PageServiceCustomerRequest
//
// @return PageServiceCustomerResponse
func (client *Client) PageServiceCustomer(request *PageServiceCustomerRequest) (_result *PageServiceCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PageServiceCustomerResponse{}
	_body, _err := client.PageServiceCustomerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发送自定义告警事件
//
// @param request - SendCustomEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendCustomEventResponse
func (client *Client) SendCustomEventWithOptions(request *SendCustomEventRequest, runtime *util.RuntimeOptions) (_result *SendCustomEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomerId)) {
		body["CustomerId"] = request.CustomerId
	}

	if !tea.BoolValue(util.IsUnset(request.DataSource)) {
		body["DataSource"] = request.DataSource
	}

	if !tea.BoolValue(util.IsUnset(request.EventDescription)) {
		body["EventDescription"] = request.EventDescription
	}

	if !tea.BoolValue(util.IsUnset(request.EventDetails)) {
		body["EventDetails"] = request.EventDetails
	}

	if !tea.BoolValue(util.IsUnset(request.EventName)) {
		body["EventName"] = request.EventName
	}

	if !tea.BoolValue(util.IsUnset(request.EventType)) {
		body["EventType"] = request.EventType
	}

	if !tea.BoolValue(util.IsUnset(request.FindTime)) {
		body["FindTime"] = request.FindTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.IsSend)) {
		body["IsSend"] = request.IsSend
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		body["Level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.OccurrenceTime)) {
		body["OccurrenceTime"] = request.OccurrenceTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		body["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		body["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.UniqueId)) {
		body["UniqueId"] = request.UniqueId
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendCustomEvent"),
		Version:     tea.String("2016-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendCustomEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送自定义告警事件
//
// @param request - SendCustomEventRequest
//
// @return SendCustomEventResponse
func (client *Client) SendCustomEvent(request *SendCustomEventRequest) (_result *SendCustomEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendCustomEventResponse{}
	_body, _err := client.SendCustomEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}