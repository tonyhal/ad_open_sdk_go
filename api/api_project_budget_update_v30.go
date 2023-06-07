/*
API Title

巨量引擎开放平台

API version: 1.0.0
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

// ProjectBudgetUpdateV30ApiService ProjectBudgetUpdateV30Api service
type ProjectBudgetUpdateV30ApiService service

type ApiOpenApiV30ProjectBudgetUpdatePostRequest struct {
	ctx                           context.Context
	ApiService                    *ProjectBudgetUpdateV30ApiService
	projectBudgetUpdateV30Request *ProjectBudgetUpdateV30Request
}

func (r *ApiOpenApiV30ProjectBudgetUpdatePostRequest) ProjectBudgetUpdateV30Request(projectBudgetUpdateV30Request ProjectBudgetUpdateV30Request) *ApiOpenApiV30ProjectBudgetUpdatePostRequest {
	r.projectBudgetUpdateV30Request = &projectBudgetUpdateV30Request
	return r
}

func (r *ApiOpenApiV30ProjectBudgetUpdatePostRequest) Execute() (*ProjectBudgetUpdateV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30ProjectBudgetUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApiV30ProjectBudgetUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ProjectBudgetUpdatePostRequest) WithLog(enable bool) *ApiOpenApiV30ProjectBudgetUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ProjectBudgetUpdatePost Method for OpenApiV30ProjectBudgetUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ProjectBudgetUpdatePostRequest
*/
func (a *ProjectBudgetUpdateV30ApiService) Post(ctx context.Context) *ApiOpenApiV30ProjectBudgetUpdatePostRequest {
	return &ApiOpenApiV30ProjectBudgetUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ProjectBudgetUpdateV30Response
func (a *ProjectBudgetUpdateV30ApiService) postExecute(r *ApiOpenApiV30ProjectBudgetUpdatePostRequest) (*ProjectBudgetUpdateV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ProjectBudgetUpdateV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/project/budget/update/"

	localVarHeaderParams := make(map[string]string)
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
	localVarPostBody = r.projectBudgetUpdateV30Request
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
