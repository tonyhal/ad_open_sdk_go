/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.3
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

// DouplusOrderListV30ApiService DouplusOrderListV30Api service
type DouplusOrderListV30ApiService service

type ApiOpenApiV30DouplusOrderListGetRequest struct {
	ctx         context.Context
	ApiService  *DouplusOrderListV30ApiService
	awemeSecUid *string
	filter      *DouplusOrderListV30Filter
	pageSize    *int64
	page        *int64
}

func (r *ApiOpenApiV30DouplusOrderListGetRequest) AwemeSecUid(awemeSecUid string) *ApiOpenApiV30DouplusOrderListGetRequest {
	r.awemeSecUid = &awemeSecUid
	return r
}

func (r *ApiOpenApiV30DouplusOrderListGetRequest) Filter(filter DouplusOrderListV30Filter) *ApiOpenApiV30DouplusOrderListGetRequest {
	r.filter = &filter
	return r
}

func (r *ApiOpenApiV30DouplusOrderListGetRequest) PageSize(pageSize int64) *ApiOpenApiV30DouplusOrderListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30DouplusOrderListGetRequest) Page(page int64) *ApiOpenApiV30DouplusOrderListGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV30DouplusOrderListGetRequest) Execute() (*DouplusOrderListV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30DouplusOrderListGetRequest) AccessToken(accessToken string) *ApiOpenApiV30DouplusOrderListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30DouplusOrderListGetRequest) WithLog(enable bool) *ApiOpenApiV30DouplusOrderListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30DouplusOrderListGet Method for OpenApiV30DouplusOrderListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30DouplusOrderListGetRequest
*/
func (a *DouplusOrderListV30ApiService) Get(ctx context.Context) *ApiOpenApiV30DouplusOrderListGetRequest {
	return &ApiOpenApiV30DouplusOrderListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DouplusOrderListV30Response
func (a *DouplusOrderListV30ApiService) getExecute(r *ApiOpenApiV30DouplusOrderListGetRequest) (*DouplusOrderListV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *DouplusOrderListV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/douplus/order/list/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.awemeSecUid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_sec_uid", r.awemeSecUid)
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
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
