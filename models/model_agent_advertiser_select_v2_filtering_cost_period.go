/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentAdvertiserSelectV2FilteringCostPeriod
type AgentAdvertiserSelectV2FilteringCostPeriod string

// List of agent_advertiser_select_v2_filtering_cost_period
const (
	TODAY_AgentAdvertiserSelectV2FilteringCostPeriod        AgentAdvertiserSelectV2FilteringCostPeriod = "TODAY"
	LAST_30_DAYS_AgentAdvertiserSelectV2FilteringCostPeriod AgentAdvertiserSelectV2FilteringCostPeriod = "LAST_30_DAYS"
	LAST_15_DAYS_AgentAdvertiserSelectV2FilteringCostPeriod AgentAdvertiserSelectV2FilteringCostPeriod = "LAST_15_DAYS"
	YESTERDAY_AgentAdvertiserSelectV2FilteringCostPeriod    AgentAdvertiserSelectV2FilteringCostPeriod = "YESTERDAY"
	LAST_7_DAYS_AgentAdvertiserSelectV2FilteringCostPeriod  AgentAdvertiserSelectV2FilteringCostPeriod = "LAST_7_DAYS"
)

// Ptr returns reference to agent_advertiser_select_v2_filtering_cost_period value
func (v AgentAdvertiserSelectV2FilteringCostPeriod) Ptr() *AgentAdvertiserSelectV2FilteringCostPeriod {
	return &v
}
