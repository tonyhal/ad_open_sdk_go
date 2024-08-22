/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// AdvertiserAttachmentUploadV30ApiService AdvertiserAttachmentUploadV30Api service
type AdvertiserAttachmentUploadV30ApiService service

type ApiOpenApiV30AdvertiserAttachmentUploadPostRequest struct {
	ctx            context.Context
	ApiService     *AdvertiserAttachmentUploadV30ApiService
	advertiserId   *int64
	attachmentType *AdvertiserAttachmentUploadV30AttachmentType
	filename       *string
	imageData      *FormFileInfo
}

func (r *ApiOpenApiV30AdvertiserAttachmentUploadPostRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30AdvertiserAttachmentUploadPostRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30AdvertiserAttachmentUploadPostRequest) AttachmentType(attachmentType AdvertiserAttachmentUploadV30AttachmentType) *ApiOpenApiV30AdvertiserAttachmentUploadPostRequest {
	r.attachmentType = &attachmentType
	return r
}

// 文件名 注意：不要包含文件路径，不要含有&#39;/&#39;等非法字符
func (r *ApiOpenApiV30AdvertiserAttachmentUploadPostRequest) Filename(filename string) *ApiOpenApiV30AdvertiserAttachmentUploadPostRequest {
	r.filename = &filename
	return r
}

// 图片数据
func (r *ApiOpenApiV30AdvertiserAttachmentUploadPostRequest) ImageData(imageData *FormFileInfo) *ApiOpenApiV30AdvertiserAttachmentUploadPostRequest {
	r.imageData = imageData
	return r
}

func (r *ApiOpenApiV30AdvertiserAttachmentUploadPostRequest) Execute() (*AdvertiserAttachmentUploadV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30AdvertiserAttachmentUploadPostRequest) AccessToken(accessToken string) *ApiOpenApiV30AdvertiserAttachmentUploadPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30AdvertiserAttachmentUploadPostRequest) WithLog(enable bool) *ApiOpenApiV30AdvertiserAttachmentUploadPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30AdvertiserAttachmentUploadPost Method for OpenApiV30AdvertiserAttachmentUploadPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30AdvertiserAttachmentUploadPostRequest
*/
func (a *AdvertiserAttachmentUploadV30ApiService) Post(ctx context.Context) *ApiOpenApiV30AdvertiserAttachmentUploadPostRequest {
	return &ApiOpenApiV30AdvertiserAttachmentUploadPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AdvertiserAttachmentUploadV30Response
func (a *AdvertiserAttachmentUploadV30ApiService) postExecute(r *ApiOpenApiV30AdvertiserAttachmentUploadPostRequest) (*AdvertiserAttachmentUploadV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AdvertiserAttachmentUploadV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/advertiser/attachment/upload/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.attachmentType == nil {
		return localVarReturnValue, nil, ReportError("attachmentType is required and must be specified")
	}
	if r.filename == nil {
		return localVarReturnValue, nil, ReportError("filename is required and must be specified")
	}
	if r.imageData == nil {
		return localVarReturnValue, nil, ReportError("imageData is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	parameterAddToHeaderOrQuery(localVarFormParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarFormParams, "attachment_type", r.attachmentType)
	parameterAddToHeaderOrQuery(localVarFormParams, "filename", r.filename)
	if r.imageData != nil {
		formFiles["image_data"] = r.imageData
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.call(r.ctx, req, &localVarReturnValue)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}
