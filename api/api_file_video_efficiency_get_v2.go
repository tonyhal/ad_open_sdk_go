/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
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

// FileVideoEfficiencyGetV2ApiService FileVideoEfficiencyGetV2Api service
type FileVideoEfficiencyGetV2ApiService service

type ApiOpenApi2FileVideoEfficiencyGetGetRequest struct {
	ctx          context.Context
	ApiService   *FileVideoEfficiencyGetV2ApiService
	advertiserId *string
	materialIds  *[]string
}

// 广告主id
func (r *ApiOpenApi2FileVideoEfficiencyGetGetRequest) AdvertiserId(advertiserId string) *ApiOpenApi2FileVideoEfficiencyGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 素材列表，单次最多可查询100个
func (r *ApiOpenApi2FileVideoEfficiencyGetGetRequest) MaterialIds(materialIds []string) *ApiOpenApi2FileVideoEfficiencyGetGetRequest {
	r.materialIds = &materialIds
	return r
}

func (r *ApiOpenApi2FileVideoEfficiencyGetGetRequest) Execute() (*FileVideoEfficiencyGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2FileVideoEfficiencyGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2FileVideoEfficiencyGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2FileVideoEfficiencyGetGetRequest) WithLog(enable bool) *ApiOpenApi2FileVideoEfficiencyGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2FileVideoEfficiencyGetGet Method for OpenApi2FileVideoEfficiencyGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2FileVideoEfficiencyGetGetRequest
*/
func (a *FileVideoEfficiencyGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2FileVideoEfficiencyGetGetRequest {
	return &ApiOpenApi2FileVideoEfficiencyGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FileVideoEfficiencyGetV2Response
func (a *FileVideoEfficiencyGetV2ApiService) getExecute(r *ApiOpenApi2FileVideoEfficiencyGetGetRequest) (*FileVideoEfficiencyGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *FileVideoEfficiencyGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/file/video/efficiency/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.materialIds == nil {
		return localVarReturnValue, nil, ReportError("materialIds is required and must be specified")
	}
	if len(*r.materialIds) < 1 {
		return localVarReturnValue, nil, ReportError("materialIds must have at least 1 elements")
	}
	if len(*r.materialIds) > 100 {
		return localVarReturnValue, nil, ReportError("materialIds must have less than 100 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "material_ids", r.materialIds)
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
