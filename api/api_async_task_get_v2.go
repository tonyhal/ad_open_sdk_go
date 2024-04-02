/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
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

// AsyncTaskGetV2ApiService AsyncTaskGetV2Api service
type AsyncTaskGetV2ApiService service

type ApiOpenApi2AsyncTaskGetGetRequest struct {
	ctx          context.Context
	ApiService   *AsyncTaskGetV2ApiService
	advertiserId *int64
	filtering    *AsyncTaskGetV2Filtering
	page         *int64
	pageSize     *int64
}

func (r *ApiOpenApi2AsyncTaskGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2AsyncTaskGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2AsyncTaskGetGetRequest) Filtering(filtering AsyncTaskGetV2Filtering) *ApiOpenApi2AsyncTaskGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2AsyncTaskGetGetRequest) Page(page int64) *ApiOpenApi2AsyncTaskGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2AsyncTaskGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2AsyncTaskGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2AsyncTaskGetGetRequest) Execute() (*AsyncTaskGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2AsyncTaskGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2AsyncTaskGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2AsyncTaskGetGetRequest) WithLog(enable bool) *ApiOpenApi2AsyncTaskGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2AsyncTaskGetGet Method for OpenApi2AsyncTaskGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2AsyncTaskGetGetRequest
*/
func (a *AsyncTaskGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2AsyncTaskGetGetRequest {
	return &ApiOpenApi2AsyncTaskGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AsyncTaskGetV2Response
func (a *AsyncTaskGetV2ApiService) getExecute(r *ApiOpenApi2AsyncTaskGetGetRequest) (*AsyncTaskGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AsyncTaskGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/async_task/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
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
