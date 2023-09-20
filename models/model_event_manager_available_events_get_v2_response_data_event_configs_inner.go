/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerAvailableEventsGetV2ResponseDataEventConfigsInner struct for EventManagerAvailableEventsGetV2ResponseDataEventConfigsInner
type EventManagerAvailableEventsGetV2ResponseDataEventConfigsInner struct {
	// 事件描述信息
	Description *string `json:"description,omitempty"`
	// 事件中文名称
	EventCnName *string `json:"event_cn_name,omitempty"`
	// 事件ID
	EventId *int64 `json:"event_id,omitempty"`
	// 事件类型
	EventType *string `json:"event_type,omitempty"`
	// 事件的附加属性
	Properties []*EventManagerAvailableEventsGetV2ResponseDataEventConfigsInnerPropertiesInner `json:"properties,omitempty"`
	// 事件回传方式列表，枚举值：`JSSDK` JS埋码 、`EXTERNAL_API` API回传 、`XPATH` XPath圈选
	TrackTypes []*EventManagerAvailableEventsGetV2DataEventConfigsTrackTypes `json:"track_types,omitempty"`
}
