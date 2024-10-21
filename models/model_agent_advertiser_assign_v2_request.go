/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentAdvertiserAssignV2Request struct for AgentAdvertiserAssignV2Request
type AgentAdvertiserAssignV2Request struct {
	// 广告主账户ID，最多支持100个
	AdvertiserIds []int64 `json:"advertiser_ids"`
	// 待分配广告主账户的所属代理商账户ID
	AgentId int64 `json:"agent_id"`
	// 待分配的员工id
	EmployeeId int64                       `json:"employee_id"`
	Role       AgentAdvertiserAssignV2Role `json:"role"`
	Type       AgentAdvertiserAssignV2Type `json:"type"`
}