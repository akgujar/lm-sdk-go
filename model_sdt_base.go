/*
 * logicmonitor_sdk
 *
 * LogicMonitor is a SaaS-based performance monitoring platform that provides full visibility into complex, hybrid infrastructures, offering granular performance monitoring and actionable data and insights. logicmonitor_sdk enables you to manage your LogicMonitor account programmatically.
 *
 * API version: 1.0.0
 * Contact: sdk@logicmonitor.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package logicmonitor

type SdtBase struct {
	EndDateTimeOnLocal string `json:"endDateTimeOnLocal,omitempty"`
	SdtType int32 `json:"sdtType,omitempty"`
	MonthDay int32 `json:"monthDay,omitempty"`
	WeekOfMonth string `json:"weekOfMonth,omitempty"`
	Admin string `json:"admin,omitempty"`
	EndDateTime int64 `json:"endDateTime,omitempty"`
	Type_ string `json:"type"`
	IsEffective bool `json:"isEffective,omitempty"`
	Minute int32 `json:"minute,omitempty"`
	Duration int32 `json:"duration,omitempty"`
	EndHour int32 `json:"endHour,omitempty"`
	StartDateTime int64 `json:"startDateTime,omitempty"`
	Hour int32 `json:"hour,omitempty"`
	StartDateTimeOnLocal string `json:"startDateTimeOnLocal,omitempty"`
	WeekDay int32 `json:"weekDay,omitempty"`
	Comment string `json:"comment,omitempty"`
	Id string `json:"id,omitempty"`
	EndMinute int32 `json:"endMinute,omitempty"`
}