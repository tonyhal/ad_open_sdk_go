/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
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

// EventManagerShareGetV30ApiService EventManagerShareGetV30Api service
type EventManagerShareGetV30ApiService service

type ApiOpenApiV30EventManagerShareGetGetRequest struct {
	ctx            context.Context
	ApiService     *EventManagerShareGetV30ApiService
	organizationId *int64
	assetId        *int64
	page           *int32
	pageSize       *int32
}

func (r *ApiOpenApiV30EventManagerShareGetGetRequest) OrganizationId(organizationId int64) *ApiOpenApiV30EventManagerShareGetGetRequest {
	r.organizationId = &organizationId
	return r
}

func (r *ApiOpenApiV30EventManagerShareGetGetRequest) AssetId(assetId int64) *ApiOpenApiV30EventManagerShareGetGetRequest {
	r.assetId = &assetId
	return r
}

func (r *ApiOpenApiV30EventManagerShareGetGetRequest) Page(page int32) *ApiOpenApiV30EventManagerShareGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV30EventManagerShareGetGetRequest) PageSize(pageSize int32) *ApiOpenApiV30EventManagerShareGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30EventManagerShareGetGetRequest) Execute() (*EventManagerShareGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30EventManagerShareGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30EventManagerShareGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30EventManagerShareGetGetRequest) WithLog(enable bool) *ApiOpenApiV30EventManagerShareGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30EventManagerShareGetGet Method for OpenApiV30EventManagerShareGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30EventManagerShareGetGetRequest
*/
func (a *EventManagerShareGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30EventManagerShareGetGetRequest {
	return &ApiOpenApiV30EventManagerShareGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return EventManagerShareGetV30Response
func (a *EventManagerShareGetV30ApiService) getExecute(r *ApiOpenApiV30EventManagerShareGetGetRequest) (*EventManagerShareGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *EventManagerShareGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/event_manager/share/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.organizationId == nil {
		return localVarReturnValue, nil, ReportError("organizationId is required and must be specified")
	}
	if r.assetId == nil {
		return localVarReturnValue, nil, ReportError("assetId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "organization_id", r.organizationId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "asset_id", r.assetId)
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
