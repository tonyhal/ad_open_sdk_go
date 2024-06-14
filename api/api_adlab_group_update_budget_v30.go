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

// AdlabGroupUpdateBudgetV30ApiService AdlabGroupUpdateBudgetV30Api service
type AdlabGroupUpdateBudgetV30ApiService service

type ApiOpenApiV30AdlabGroupUpdateBudgetPostRequest struct {
	ctx                              context.Context
	ApiService                       *AdlabGroupUpdateBudgetV30ApiService
	adlabGroupUpdateBudgetV30Request *AdlabGroupUpdateBudgetV30Request
}

func (r *ApiOpenApiV30AdlabGroupUpdateBudgetPostRequest) AdlabGroupUpdateBudgetV30Request(adlabGroupUpdateBudgetV30Request AdlabGroupUpdateBudgetV30Request) *ApiOpenApiV30AdlabGroupUpdateBudgetPostRequest {
	r.adlabGroupUpdateBudgetV30Request = &adlabGroupUpdateBudgetV30Request
	return r
}

func (r *ApiOpenApiV30AdlabGroupUpdateBudgetPostRequest) Execute() (*AdlabGroupUpdateBudgetV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30AdlabGroupUpdateBudgetPostRequest) AccessToken(accessToken string) *ApiOpenApiV30AdlabGroupUpdateBudgetPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30AdlabGroupUpdateBudgetPostRequest) WithLog(enable bool) *ApiOpenApiV30AdlabGroupUpdateBudgetPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30AdlabGroupUpdateBudgetPost Method for OpenApiV30AdlabGroupUpdateBudgetPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30AdlabGroupUpdateBudgetPostRequest
*/
func (a *AdlabGroupUpdateBudgetV30ApiService) Post(ctx context.Context) *ApiOpenApiV30AdlabGroupUpdateBudgetPostRequest {
	return &ApiOpenApiV30AdlabGroupUpdateBudgetPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AdlabGroupUpdateBudgetV30Response
func (a *AdlabGroupUpdateBudgetV30ApiService) postExecute(r *ApiOpenApiV30AdlabGroupUpdateBudgetPostRequest) (*AdlabGroupUpdateBudgetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AdlabGroupUpdateBudgetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/adlab/group/update_budget/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.adlabGroupUpdateBudgetV30Request
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
