/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// NativeAnchorQrcodePreviewGetV30ApiService NativeAnchorQrcodePreviewGetV30Api service
type NativeAnchorQrcodePreviewGetV30ApiService service

type ApiOpenApiV30NativeAnchorQrcodePreviewGetGetRequest struct {
	ctx          context.Context
	ApiService   *NativeAnchorQrcodePreviewGetV30ApiService
	advertiserId *int64
	anchorIds    *[]string
	anchorType   *NativeAnchorQrcodePreviewGetV30AnchorType
}

func (r *ApiOpenApiV30NativeAnchorQrcodePreviewGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30NativeAnchorQrcodePreviewGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30NativeAnchorQrcodePreviewGetGetRequest) AnchorIds(anchorIds []string) *ApiOpenApiV30NativeAnchorQrcodePreviewGetGetRequest {
	r.anchorIds = &anchorIds
	return r
}

func (r *ApiOpenApiV30NativeAnchorQrcodePreviewGetGetRequest) AnchorType(anchorType NativeAnchorQrcodePreviewGetV30AnchorType) *ApiOpenApiV30NativeAnchorQrcodePreviewGetGetRequest {
	r.anchorType = &anchorType
	return r
}

func (r *ApiOpenApiV30NativeAnchorQrcodePreviewGetGetRequest) Execute() (*NativeAnchorQrcodePreviewGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30NativeAnchorQrcodePreviewGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30NativeAnchorQrcodePreviewGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30NativeAnchorQrcodePreviewGetGetRequest) WithLog(enable bool) *ApiOpenApiV30NativeAnchorQrcodePreviewGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30NativeAnchorQrcodePreviewGetGet Method for OpenApiV30NativeAnchorQrcodePreviewGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30NativeAnchorQrcodePreviewGetGetRequest
*/
func (a *NativeAnchorQrcodePreviewGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30NativeAnchorQrcodePreviewGetGetRequest {
	return &ApiOpenApiV30NativeAnchorQrcodePreviewGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return NativeAnchorQrcodePreviewGetV30Response
func (a *NativeAnchorQrcodePreviewGetV30ApiService) getExecute(r *ApiOpenApiV30NativeAnchorQrcodePreviewGetGetRequest) (*NativeAnchorQrcodePreviewGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *NativeAnchorQrcodePreviewGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/native_anchor/qrcode_preview/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.anchorIds == nil {
		return localVarReturnValue, nil, ReportError("anchorIds is required and must be specified")
	}
	if len(*r.anchorIds) < 1 {
		return localVarReturnValue, nil, ReportError("anchorIds must have at least 1 elements")
	}
	if len(*r.anchorIds) > 10 {
		return localVarReturnValue, nil, ReportError("anchorIds must have less than 10 elements")
	}
	if r.anchorType == nil {
		return localVarReturnValue, nil, ReportError("anchorType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "anchor_ids", r.anchorIds)
	parameterAddToHeaderOrQuery(localVarQueryParams, "anchor_type", r.anchorType)
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
