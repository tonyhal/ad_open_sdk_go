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

// ToolsAppManagementShareAccountListV2ApiService ToolsAppManagementShareAccountListV2Api service
type ToolsAppManagementShareAccountListV2ApiService service

type ApiOpenApi2ToolsAppManagementShareAccountListGetRequest struct {
	ctx            context.Context
	ApiService     *ToolsAppManagementShareAccountListV2ApiService
	organizationId *int64
	packageId      *string
	page           *int64
	pageSize       *int64
	searchType     *ToolsAppManagementShareAccountListV2SearchType
}

func (r *ApiOpenApi2ToolsAppManagementShareAccountListGetRequest) OrganizationId(organizationId int64) *ApiOpenApi2ToolsAppManagementShareAccountListGetRequest {
	r.organizationId = &organizationId
	return r
}

func (r *ApiOpenApi2ToolsAppManagementShareAccountListGetRequest) PackageId(packageId string) *ApiOpenApi2ToolsAppManagementShareAccountListGetRequest {
	r.packageId = &packageId
	return r
}

func (r *ApiOpenApi2ToolsAppManagementShareAccountListGetRequest) Page(page int64) *ApiOpenApi2ToolsAppManagementShareAccountListGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ToolsAppManagementShareAccountListGetRequest) PageSize(pageSize int64) *ApiOpenApi2ToolsAppManagementShareAccountListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsAppManagementShareAccountListGetRequest) SearchType(searchType ToolsAppManagementShareAccountListV2SearchType) *ApiOpenApi2ToolsAppManagementShareAccountListGetRequest {
	r.searchType = &searchType
	return r
}

func (r *ApiOpenApi2ToolsAppManagementShareAccountListGetRequest) Execute() (*ToolsAppManagementShareAccountListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsAppManagementShareAccountListGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAppManagementShareAccountListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAppManagementShareAccountListGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsAppManagementShareAccountListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAppManagementShareAccountListGet Method for OpenApi2ToolsAppManagementShareAccountListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAppManagementShareAccountListGetRequest
*/
func (a *ToolsAppManagementShareAccountListV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsAppManagementShareAccountListGetRequest {
	return &ApiOpenApi2ToolsAppManagementShareAccountListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAppManagementShareAccountListV2Response
func (a *ToolsAppManagementShareAccountListV2ApiService) getExecute(r *ApiOpenApi2ToolsAppManagementShareAccountListGetRequest) (*ToolsAppManagementShareAccountListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsAppManagementShareAccountListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/app_management/share_account/list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.organizationId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "organization_id", r.organizationId)
	}
	if r.packageId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "package_id", r.packageId)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.searchType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search_type", r.searchType)
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
