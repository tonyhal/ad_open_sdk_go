/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
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

// QianchuanFileVideoEfficiencyGetV10ApiService QianchuanFileVideoEfficiencyGetV10Api service
type QianchuanFileVideoEfficiencyGetV10ApiService service

type ApiOpenApiV10QianchuanFileVideoEfficiencyGetGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanFileVideoEfficiencyGetV10ApiService
	advertiserId *string
	materialIds  *[]string
}

func (r *ApiOpenApiV10QianchuanFileVideoEfficiencyGetGetRequest) AdvertiserId(advertiserId string) *ApiOpenApiV10QianchuanFileVideoEfficiencyGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanFileVideoEfficiencyGetGetRequest) MaterialIds(materialIds []string) *ApiOpenApiV10QianchuanFileVideoEfficiencyGetGetRequest {
	r.materialIds = &materialIds
	return r
}

func (r *ApiOpenApiV10QianchuanFileVideoEfficiencyGetGetRequest) Execute() (*QianchuanFileVideoEfficiencyGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanFileVideoEfficiencyGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanFileVideoEfficiencyGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanFileVideoEfficiencyGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanFileVideoEfficiencyGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanFileVideoEfficiencyGetGet Method for OpenApiV10QianchuanFileVideoEfficiencyGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanFileVideoEfficiencyGetGetRequest
*/
func (a *QianchuanFileVideoEfficiencyGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanFileVideoEfficiencyGetGetRequest {
	return &ApiOpenApiV10QianchuanFileVideoEfficiencyGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanFileVideoEfficiencyGetV10Response
func (a *QianchuanFileVideoEfficiencyGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanFileVideoEfficiencyGetGetRequest) (*QianchuanFileVideoEfficiencyGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *QianchuanFileVideoEfficiencyGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/file/video/efficiency/get/"

	localVarHeaderParams := make(map[string]string)
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
