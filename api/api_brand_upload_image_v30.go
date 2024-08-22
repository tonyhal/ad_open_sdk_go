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

// BrandUploadImageV30ApiService BrandUploadImageV30Api service
type BrandUploadImageV30ApiService service

type ApiOpenApiV30BrandUploadImagePostRequest struct {
	ctx          context.Context
	ApiService   *BrandUploadImageV30ApiService
	advertiserId *int64
	imageFile    *FormFileInfo
}

// 广告主ID
func (r *ApiOpenApiV30BrandUploadImagePostRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30BrandUploadImagePostRequest {
	r.advertiserId = &advertiserId
	return r
}

// 图片文件
func (r *ApiOpenApiV30BrandUploadImagePostRequest) ImageFile(imageFile *FormFileInfo) *ApiOpenApiV30BrandUploadImagePostRequest {
	r.imageFile = imageFile
	return r
}

func (r *ApiOpenApiV30BrandUploadImagePostRequest) Execute() (*BrandUploadImageV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30BrandUploadImagePostRequest) AccessToken(accessToken string) *ApiOpenApiV30BrandUploadImagePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30BrandUploadImagePostRequest) WithLog(enable bool) *ApiOpenApiV30BrandUploadImagePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30BrandUploadImagePost Method for OpenApiV30BrandUploadImagePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30BrandUploadImagePostRequest
*/
func (a *BrandUploadImageV30ApiService) Post(ctx context.Context) *ApiOpenApiV30BrandUploadImagePostRequest {
	return &ApiOpenApiV30BrandUploadImagePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BrandUploadImageV30Response
func (a *BrandUploadImageV30ApiService) postExecute(r *ApiOpenApiV30BrandUploadImagePostRequest) (*BrandUploadImageV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *BrandUploadImageV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/brand/upload_image/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.imageFile == nil {
		return localVarReturnValue, nil, ReportError("imageFile is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	parameterAddToHeaderOrQuery(localVarFormParams, "advertiser_id", r.advertiserId)
	if r.imageFile != nil {
		formFiles["image_file"] = r.imageFile
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
