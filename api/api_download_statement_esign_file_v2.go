/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// DownloadStatementEsignFileV2ApiService DownloadStatementEsignFileV2Api service
type DownloadStatementEsignFileV2ApiService service

type ApiOpenApi2DownloadStatementEsignFileGetRequest struct {
	ctx         context.Context
	ApiService  *DownloadStatementEsignFileV2ApiService
	agentId     *int64
	statementId *int64
}

// 代理商ID
func (r *ApiOpenApi2DownloadStatementEsignFileGetRequest) AgentId(agentId int64) *ApiOpenApi2DownloadStatementEsignFileGetRequest {
	r.agentId = &agentId
	return r
}

// 结算单ID
func (r *ApiOpenApi2DownloadStatementEsignFileGetRequest) StatementId(statementId int64) *ApiOpenApi2DownloadStatementEsignFileGetRequest {
	r.statementId = &statementId
	return r
}

func (r *ApiOpenApi2DownloadStatementEsignFileGetRequest) Execute() (*DownloadStatementEsignFileV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2DownloadStatementEsignFileGetRequest) AccessToken(accessToken string) *ApiOpenApi2DownloadStatementEsignFileGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2DownloadStatementEsignFileGetRequest) WithLog(enable bool) *ApiOpenApi2DownloadStatementEsignFileGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2DownloadStatementEsignFileGet Method for OpenApi2DownloadStatementEsignFileGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2DownloadStatementEsignFileGetRequest
*/
func (a *DownloadStatementEsignFileV2ApiService) Get(ctx context.Context) *ApiOpenApi2DownloadStatementEsignFileGetRequest {
	return &ApiOpenApi2DownloadStatementEsignFileGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DownloadStatementEsignFileV2Response
func (a *DownloadStatementEsignFileV2ApiService) getExecute(r *ApiOpenApi2DownloadStatementEsignFileGetRequest) (*DownloadStatementEsignFileV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *DownloadStatementEsignFileV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/download/statement/esign_file/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.agentId == nil {
		return localVarReturnValue, nil, ReportError("agentId is required and must be specified")
	}
	if r.statementId == nil {
		return localVarReturnValue, nil, ReportError("statementId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "agent_id", r.agentId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "statement_id", r.statementId)
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.call(r.ctx, req, &localVarReturnValue)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}