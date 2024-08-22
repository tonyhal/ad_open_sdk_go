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

type CommonApiService service

type CommonApiRequest struct {
	ctx          context.Context
	path         string
	ApiService   *CommonApiService
	requestBody  interface{}
	requestQuery map[string]interface{}
	requestFile  map[string]FormFile
	requestForm  map[string]interface{}
	httpMethod   string
	contentType  string
}

func (r *CommonApiRequest) WithLog(enable bool) *CommonApiRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

func (r *CommonApiRequest) AccessToken(accessToken string) *CommonApiRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *CommonApiRequest) RequestBody(requestBody interface{}) *CommonApiRequest {
	r.requestBody = requestBody
	return r
}

func (r *CommonApiRequest) RequestQuery(requestQuery map[string]interface{}) *CommonApiRequest {
	r.requestQuery = requestQuery
	return r
}

func (r *CommonApiRequest) RequestFile(requestFile map[string]FormFile) *CommonApiRequest {
	r.requestFile = requestFile
	return r
}

func (r *CommonApiRequest) RequestForm(requestForm map[string]interface{}) *CommonApiRequest {
	r.requestForm = requestForm
	return r
}

func (r *CommonApiRequest) Execute() (*CommonResponse, *http.Response, error) {
	return r.ApiService.execute(r)
}

func (a *CommonApiService) Get(ctx context.Context, path string) *CommonApiRequest {
	return &CommonApiRequest{
		ctx:        ctx,
		path:       path,
		ApiService: a,
		httpMethod: http.MethodGet,
	}
}

func (a *CommonApiService) Post(ctx context.Context, path string) *CommonApiRequest {
	return &CommonApiRequest{
		ctx:         ctx,
		path:        path,
		ApiService:  a,
		httpMethod:  http.MethodPost,
		contentType: "application/json",
	}
}

func (a *CommonApiService) PostMultipart(ctx context.Context, path string) *CommonApiRequest {
	return &CommonApiRequest{
		ctx:         ctx,
		path:        path,
		ApiService:  a,
		httpMethod:  http.MethodPost,
		contentType: "multipart/form-data",
	}
}

func (a *CommonApiService) execute(r *CommonApiRequest) (*CommonResponse, *http.Response, error) {
	var localVarReturnValue *CommonResponse

	r.ctx = a.client.prepareCtx(r.ctx)
	localBasePath := a.client.Cfg.GetBasePath()
	localVarPath := localBasePath + r.path
	localVarHeaderParams := make(map[string]string)
	localVarFormFiles := make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.contentType != "" {
		localVarHeaderParams["Content-Type"] = r.contentType
	}
	if len(r.requestQuery) > 0 {
		for k, v := range r.requestQuery {
			parameterAddToHeaderOrQuery(localVarQueryParams, k, v)
		}
	}
	if len(r.requestForm) > 0 {
		for k, v := range r.requestForm {
			parameterAddToHeaderOrQuery(localVarFormParams, k, v)
		}
	}

	if len(r.requestFile) > 0 {
		for k, v := range r.requestFile {
			localVarFormFiles[k] = &FormFileInfo{FileName: v.FileName, FileBytes: v.FileBytes}
		}
	}

	req, err := a.client.prepareRequest(r.ctx, localVarPath, r.httpMethod, r.requestBody,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.call(r.ctx, req, &localVarReturnValue)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}
