/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
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

// QianchuanFileVideoOriginalGetV10ApiService QianchuanFileVideoOriginalGetV10Api service
type QianchuanFileVideoOriginalGetV10ApiService service

type ApiOpenApiV10QianchuanFileVideoOriginalGetGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanFileVideoOriginalGetV10ApiService
	advertiserId *string
	materialIds  *[]string
}

func (r *ApiOpenApiV10QianchuanFileVideoOriginalGetGetRequest) AdvertiserId(advertiserId string) *ApiOpenApiV10QianchuanFileVideoOriginalGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanFileVideoOriginalGetGetRequest) MaterialIds(materialIds []string) *ApiOpenApiV10QianchuanFileVideoOriginalGetGetRequest {
	r.materialIds = &materialIds
	return r
}

func (r *ApiOpenApiV10QianchuanFileVideoOriginalGetGetRequest) Execute() (*QianchuanFileVideoOriginalGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanFileVideoOriginalGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanFileVideoOriginalGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanFileVideoOriginalGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanFileVideoOriginalGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanFileVideoOriginalGetGet Method for OpenApiV10QianchuanFileVideoOriginalGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanFileVideoOriginalGetGetRequest
*/
func (a *QianchuanFileVideoOriginalGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanFileVideoOriginalGetGetRequest {
	return &ApiOpenApiV10QianchuanFileVideoOriginalGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanFileVideoOriginalGetV10Response
func (a *QianchuanFileVideoOriginalGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanFileVideoOriginalGetGetRequest) (*QianchuanFileVideoOriginalGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanFileVideoOriginalGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/file/video/original/get/"

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
