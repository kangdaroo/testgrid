/*
Copyright The TestGrid Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: summary.proto

package summary

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TestInfo_Trend int32

const (
	TestInfo_UNKNOWN   TestInfo_Trend = 0
	TestInfo_NO_CHANGE TestInfo_Trend = 1
	TestInfo_UP        TestInfo_Trend = 2
	TestInfo_DOWN      TestInfo_Trend = 3
)

var TestInfo_Trend_name = map[int32]string{
	0: "UNKNOWN",
	1: "NO_CHANGE",
	2: "UP",
	3: "DOWN",
}

var TestInfo_Trend_value = map[string]int32{
	"UNKNOWN":   0,
	"NO_CHANGE": 1,
	"UP":        2,
	"DOWN":      3,
}

func (x TestInfo_Trend) String() string {
	return proto.EnumName(TestInfo_Trend_name, int32(x))
}

func (TestInfo_Trend) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f7168d0e3f3f5589, []int{1, 0}
}

type DashboardTabSummary_TabStatus int32

const (
	DashboardTabSummary_NOT_SET DashboardTabSummary_TabStatus = 0
	DashboardTabSummary_UNKNOWN DashboardTabSummary_TabStatus = 1
	DashboardTabSummary_PASS    DashboardTabSummary_TabStatus = 2
	DashboardTabSummary_FAIL    DashboardTabSummary_TabStatus = 3
	DashboardTabSummary_FLAKY   DashboardTabSummary_TabStatus = 4
	DashboardTabSummary_STALE   DashboardTabSummary_TabStatus = 5
	DashboardTabSummary_BROKEN  DashboardTabSummary_TabStatus = 6
)

var DashboardTabSummary_TabStatus_name = map[int32]string{
	0: "NOT_SET",
	1: "UNKNOWN",
	2: "PASS",
	3: "FAIL",
	4: "FLAKY",
	5: "STALE",
	6: "BROKEN",
}

var DashboardTabSummary_TabStatus_value = map[string]int32{
	"NOT_SET": 0,
	"UNKNOWN": 1,
	"PASS":    2,
	"FAIL":    3,
	"FLAKY":   4,
	"STALE":   5,
	"BROKEN":  6,
}

func (x DashboardTabSummary_TabStatus) String() string {
	return proto.EnumName(DashboardTabSummary_TabStatus_name, int32(x))
}

func (DashboardTabSummary_TabStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f7168d0e3f3f5589, []int{4, 0}
}

// Summary of a failing test.
type FailingTestSummary struct {
	// Display name of the test.
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Name of the test. E.g., the target for tests in Sponge.
	TestName string `protobuf:"bytes,2,opt,name=test_name,json=testName,proto3" json:"test_name,omitempty"`
	// First build ID at which the test failed.
	FailBuildId string `protobuf:"bytes,3,opt,name=fail_build_id,json=failBuildId,proto3" json:"fail_build_id,omitempty"`
	// Timestamp for the first cycle in which the test failed.
	FailTimestamp float64 `protobuf:"fixed64,4,opt,name=fail_timestamp,json=failTimestamp,proto3" json:"fail_timestamp,omitempty"`
	// Last build ID at which the test passed.
	PassBuildId string `protobuf:"bytes,5,opt,name=pass_build_id,json=passBuildId,proto3" json:"pass_build_id,omitempty"`
	// Timestamp for the last cycle in which the test passed.
	PassTimestamp float64 `protobuf:"fixed64,6,opt,name=pass_timestamp,json=passTimestamp,proto3" json:"pass_timestamp,omitempty"`
	// Number of times the test has failed.
	FailCount int32 `protobuf:"varint,7,opt,name=fail_count,json=failCount,proto3" json:"fail_count,omitempty"`
	// Link to search for build changes.
	BuildLink string `protobuf:"bytes,8,opt,name=build_link,json=buildLink,proto3" json:"build_link,omitempty"`
	// Text for option to search for build changes.
	BuildLinkText string `protobuf:"bytes,9,opt,name=build_link_text,json=buildLinkText,proto3" json:"build_link_text,omitempty"`
	// Text to display for link to search for build changes.
	BuildUrlText string `protobuf:"bytes,10,opt,name=build_url_text,json=buildUrlText,proto3" json:"build_url_text,omitempty"`
	// Text for failure statuses associated with this test.
	FailureMessage string `protobuf:"bytes,11,opt,name=failure_message,json=failureMessage,proto3" json:"failure_message,omitempty"`
	// List of bug IDs for bugs associated with this test.
	LinkedBugs []string `protobuf:"bytes,12,rep,name=linked_bugs,json=linkedBugs,proto3" json:"linked_bugs,omitempty"`
	// A link to the first build in which the test failed.
	FailTestLink string `protobuf:"bytes,13,opt,name=fail_test_link,json=failTestLink,proto3" json:"fail_test_link,omitempty"`
	// The test ID for the latest test failure. (Does not indicate the failure is
	// 'over', just the latest test failure we found.)
	LatestFailBuildId string `protobuf:"bytes,14,opt,name=latest_fail_build_id,json=latestFailBuildId,proto3" json:"latest_fail_build_id,omitempty"`
	// Maps (property name):(property value) for arbitrary alert properties.
	Properties map[string]string `protobuf:"bytes,15,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A list of IDs for issue hotlists related to this failure.
	HotlistIds           []string `protobuf:"bytes,16,rep,name=hotlist_ids,json=hotlistIds,proto3" json:"hotlist_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FailingTestSummary) Reset()         { *m = FailingTestSummary{} }
func (m *FailingTestSummary) String() string { return proto.CompactTextString(m) }
func (*FailingTestSummary) ProtoMessage()    {}
func (*FailingTestSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7168d0e3f3f5589, []int{0}
}

func (m *FailingTestSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FailingTestSummary.Unmarshal(m, b)
}
func (m *FailingTestSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FailingTestSummary.Marshal(b, m, deterministic)
}
func (m *FailingTestSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FailingTestSummary.Merge(m, src)
}
func (m *FailingTestSummary) XXX_Size() int {
	return xxx_messageInfo_FailingTestSummary.Size(m)
}
func (m *FailingTestSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_FailingTestSummary.DiscardUnknown(m)
}

var xxx_messageInfo_FailingTestSummary proto.InternalMessageInfo

func (m *FailingTestSummary) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *FailingTestSummary) GetTestName() string {
	if m != nil {
		return m.TestName
	}
	return ""
}

func (m *FailingTestSummary) GetFailBuildId() string {
	if m != nil {
		return m.FailBuildId
	}
	return ""
}

func (m *FailingTestSummary) GetFailTimestamp() float64 {
	if m != nil {
		return m.FailTimestamp
	}
	return 0
}

func (m *FailingTestSummary) GetPassBuildId() string {
	if m != nil {
		return m.PassBuildId
	}
	return ""
}

func (m *FailingTestSummary) GetPassTimestamp() float64 {
	if m != nil {
		return m.PassTimestamp
	}
	return 0
}

func (m *FailingTestSummary) GetFailCount() int32 {
	if m != nil {
		return m.FailCount
	}
	return 0
}

func (m *FailingTestSummary) GetBuildLink() string {
	if m != nil {
		return m.BuildLink
	}
	return ""
}

func (m *FailingTestSummary) GetBuildLinkText() string {
	if m != nil {
		return m.BuildLinkText
	}
	return ""
}

func (m *FailingTestSummary) GetBuildUrlText() string {
	if m != nil {
		return m.BuildUrlText
	}
	return ""
}

func (m *FailingTestSummary) GetFailureMessage() string {
	if m != nil {
		return m.FailureMessage
	}
	return ""
}

func (m *FailingTestSummary) GetLinkedBugs() []string {
	if m != nil {
		return m.LinkedBugs
	}
	return nil
}

func (m *FailingTestSummary) GetFailTestLink() string {
	if m != nil {
		return m.FailTestLink
	}
	return ""
}

func (m *FailingTestSummary) GetLatestFailBuildId() string {
	if m != nil {
		return m.LatestFailBuildId
	}
	return ""
}

func (m *FailingTestSummary) GetProperties() map[string]string {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *FailingTestSummary) GetHotlistIds() []string {
	if m != nil {
		return m.HotlistIds
	}
	return nil
}

// Metrics about a specific test, i.e. passes, fails, total runs, etc.
// Next ID: 12
type TestInfo struct {
	// The display name of the test, typically what is shown for each row in TestGrid
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The total number of test runs not including runs failed due to infrastructure
	// failures.
	TotalNonInfraRuns int32 `protobuf:"varint,2,opt,name=total_non_infra_runs,json=totalNonInfraRuns,proto3" json:"total_non_infra_runs,omitempty"`
	// The number of passed test runs not including runs failed due to
	// infrastructure failures.
	PassedNonInfraRuns int32 `protobuf:"varint,3,opt,name=passed_non_infra_runs,json=passedNonInfraRuns,proto3" json:"passed_non_infra_runs,omitempty"`
	// The number of failed test runs not including runs failed due to
	// infrastructure failures.
	FailedNonInfraRuns int32 `protobuf:"varint,4,opt,name=failed_non_infra_runs,json=failedNonInfraRuns,proto3" json:"failed_non_infra_runs,omitempty"`
	// The number of failed test runs specifically due to infrastructure
	// failures.
	FailedInfraRuns int32 `protobuf:"varint,5,opt,name=failed_infra_runs,json=failedInfraRuns,proto3" json:"failed_infra_runs,omitempty"`
	// The total number of all runs, including failures due to infrastructure
	TotalRunsWithInfra int32 `protobuf:"varint,6,opt,name=total_runs_with_infra,json=totalRunsWithInfra,proto3" json:"total_runs_with_infra,omitempty"`
	// Any other type of runs not included above.
	OtherRuns int32 `protobuf:"varint,7,opt,name=other_runs,json=otherRuns,proto3" json:"other_runs,omitempty"`
	// The flakiness of the test, measured out of 100
	Flakiness float32 `protobuf:"fixed32,8,opt,name=flakiness,proto3" json:"flakiness,omitempty"`
	// The flakiness of the test from previous intervals
	PreviousFlakiness []float32 `protobuf:"fixed32,10,rep,packed,name=previous_flakiness,json=previousFlakiness,proto3" json:"previous_flakiness,omitempty"`
	// The change of flakiness based on the last interval's flakiness
	// e.g. if last interval the flakiness was 50, and now it's 75, the
	// trend is UP. A trend of NO_CHANGE means last week and this week were
	// exactly the same. The interval is set by each tab's config, with
	// a default of 7 days.
	ChangeFromLastInterval TestInfo_Trend `protobuf:"varint,9,opt,name=change_from_last_interval,json=changeFromLastInterval,proto3,enum=TestInfo_Trend" json:"change_from_last_interval,omitempty"`
	// A map of infra failure name to the count of that failure for the interval.
	InfraFailures        map[string]int32 `protobuf:"bytes,11,rep,name=infra_failures,json=infraFailures,proto3" json:"infra_failures,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TestInfo) Reset()         { *m = TestInfo{} }
func (m *TestInfo) String() string { return proto.CompactTextString(m) }
func (*TestInfo) ProtoMessage()    {}
func (*TestInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7168d0e3f3f5589, []int{1}
}

func (m *TestInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestInfo.Unmarshal(m, b)
}
func (m *TestInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestInfo.Marshal(b, m, deterministic)
}
func (m *TestInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestInfo.Merge(m, src)
}
func (m *TestInfo) XXX_Size() int {
	return xxx_messageInfo_TestInfo.Size(m)
}
func (m *TestInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TestInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TestInfo proto.InternalMessageInfo

func (m *TestInfo) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *TestInfo) GetTotalNonInfraRuns() int32 {
	if m != nil {
		return m.TotalNonInfraRuns
	}
	return 0
}

func (m *TestInfo) GetPassedNonInfraRuns() int32 {
	if m != nil {
		return m.PassedNonInfraRuns
	}
	return 0
}

func (m *TestInfo) GetFailedNonInfraRuns() int32 {
	if m != nil {
		return m.FailedNonInfraRuns
	}
	return 0
}

func (m *TestInfo) GetFailedInfraRuns() int32 {
	if m != nil {
		return m.FailedInfraRuns
	}
	return 0
}

func (m *TestInfo) GetTotalRunsWithInfra() int32 {
	if m != nil {
		return m.TotalRunsWithInfra
	}
	return 0
}

func (m *TestInfo) GetOtherRuns() int32 {
	if m != nil {
		return m.OtherRuns
	}
	return 0
}

func (m *TestInfo) GetFlakiness() float32 {
	if m != nil {
		return m.Flakiness
	}
	return 0
}

func (m *TestInfo) GetPreviousFlakiness() []float32 {
	if m != nil {
		return m.PreviousFlakiness
	}
	return nil
}

func (m *TestInfo) GetChangeFromLastInterval() TestInfo_Trend {
	if m != nil {
		return m.ChangeFromLastInterval
	}
	return TestInfo_UNKNOWN
}

func (m *TestInfo) GetInfraFailures() map[string]int32 {
	if m != nil {
		return m.InfraFailures
	}
	return nil
}

// Summary of the flakiness and overall healthiness of a dashboard tab
type HealthinessInfo struct {
	// The start of the time frame that the analysis was run for.
	// Represents the lower bound but does not guarantee that the earliest
	// test occurred at start
	Start *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	// The end of the time frame that the analysis was run for.
	// Same caveat as above but for upper bound.
	End *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	// A list of test entries associated with this tab + timeframe.
	Tests []*TestInfo `protobuf:"bytes,3,rep,name=tests,proto3" json:"tests,omitempty"`
	// The flakiness out of 100 (think percentage but drop the sign)
	AverageFlakiness float32 `protobuf:"fixed32,4,opt,name=average_flakiness,json=averageFlakiness,proto3" json:"average_flakiness,omitempty"`
	// The average flakiness for previous intervals
	PreviousFlakiness    []float32 `protobuf:"fixed32,5,rep,packed,name=previous_flakiness,json=previousFlakiness,proto3" json:"previous_flakiness,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *HealthinessInfo) Reset()         { *m = HealthinessInfo{} }
func (m *HealthinessInfo) String() string { return proto.CompactTextString(m) }
func (*HealthinessInfo) ProtoMessage()    {}
func (*HealthinessInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7168d0e3f3f5589, []int{2}
}

func (m *HealthinessInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthinessInfo.Unmarshal(m, b)
}
func (m *HealthinessInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthinessInfo.Marshal(b, m, deterministic)
}
func (m *HealthinessInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthinessInfo.Merge(m, src)
}
func (m *HealthinessInfo) XXX_Size() int {
	return xxx_messageInfo_HealthinessInfo.Size(m)
}
func (m *HealthinessInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthinessInfo.DiscardUnknown(m)
}

var xxx_messageInfo_HealthinessInfo proto.InternalMessageInfo

func (m *HealthinessInfo) GetStart() *timestamp.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *HealthinessInfo) GetEnd() *timestamp.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *HealthinessInfo) GetTests() []*TestInfo {
	if m != nil {
		return m.Tests
	}
	return nil
}

func (m *HealthinessInfo) GetAverageFlakiness() float32 {
	if m != nil {
		return m.AverageFlakiness
	}
	return 0
}

func (m *HealthinessInfo) GetPreviousFlakiness() []float32 {
	if m != nil {
		return m.PreviousFlakiness
	}
	return nil
}

// Information about alerts that have been sent
type AlertingData struct {
	// Seconds since epoch at which an email was last sent
	LastEmailTime        *timestamp.Timestamp `protobuf:"bytes,1,opt,name=last_email_time,json=lastEmailTime,proto3" json:"last_email_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AlertingData) Reset()         { *m = AlertingData{} }
func (m *AlertingData) String() string { return proto.CompactTextString(m) }
func (*AlertingData) ProtoMessage()    {}
func (*AlertingData) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7168d0e3f3f5589, []int{3}
}

func (m *AlertingData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertingData.Unmarshal(m, b)
}
func (m *AlertingData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertingData.Marshal(b, m, deterministic)
}
func (m *AlertingData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertingData.Merge(m, src)
}
func (m *AlertingData) XXX_Size() int {
	return xxx_messageInfo_AlertingData.Size(m)
}
func (m *AlertingData) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertingData.DiscardUnknown(m)
}

var xxx_messageInfo_AlertingData proto.InternalMessageInfo

func (m *AlertingData) GetLastEmailTime() *timestamp.Timestamp {
	if m != nil {
		return m.LastEmailTime
	}
	return nil
}

// Summary of a dashboard tab.
type DashboardTabSummary struct {
	// The name of the dashboard.
	DashboardName string `protobuf:"bytes,1,opt,name=dashboard_name,json=dashboardName,proto3" json:"dashboard_name,omitempty"`
	// The name of the dashboard tab.
	DashboardTabName string `protobuf:"bytes,2,opt,name=dashboard_tab_name,json=dashboardTabName,proto3" json:"dashboard_tab_name,omitempty"`
	// Any top-level alert on this dashboard tab.
	Alert string `protobuf:"bytes,3,opt,name=alert,proto3" json:"alert,omitempty"`
	// List of failing test summary information.
	FailingTestSummaries []*FailingTestSummary `protobuf:"bytes,4,rep,name=failing_test_summaries,json=failingTestSummaries,proto3" json:"failing_test_summaries,omitempty"`
	// Seconds since epoch at which the test group was last updated.
	LastUpdateTimestamp float64 `protobuf:"fixed64,5,opt,name=last_update_timestamp,json=lastUpdateTimestamp,proto3" json:"last_update_timestamp,omitempty"`
	// A summary of the status of this dashboard tab.
	Status string `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	// The overall status for this dashboard tab.
	OverallStatus DashboardTabSummary_TabStatus `protobuf:"varint,7,opt,name=overall_status,json=overallStatus,proto3,enum=DashboardTabSummary_TabStatus" json:"overall_status,omitempty"`
	// The ID for the latest passing build.
	LatestGreen string `protobuf:"bytes,8,opt,name=latest_green,json=latestGreen,proto3" json:"latest_green,omitempty"`
	// Seconds since epoch at which tests last ran.
	LastRunTimestamp float64 `protobuf:"fixed64,9,opt,name=last_run_timestamp,json=lastRunTimestamp,proto3" json:"last_run_timestamp,omitempty"`
	// String indicating the URL for linking to a bug.
	BugUrl string `protobuf:"bytes,10,opt,name=bug_url,json=bugUrl,proto3" json:"bug_url,omitempty"`
	// Metrics for the recent healthiness of a tab
	Healthiness *HealthinessInfo `protobuf:"bytes,12,opt,name=healthiness,proto3" json:"healthiness,omitempty"`
	// All the issue IDs linked to this tab.
	LinkedIssues []string `protobuf:"bytes,13,rep,name=linked_issues,json=linkedIssues,proto3" json:"linked_issues,omitempty"`
	// Metrics about alerts sent with respect to this summary
	// Maintained by alerter; does not need to be populated by summarizer
	AlertingData         *AlertingData `protobuf:"bytes,14,opt,name=alerting_data,json=alertingData,proto3" json:"alerting_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DashboardTabSummary) Reset()         { *m = DashboardTabSummary{} }
func (m *DashboardTabSummary) String() string { return proto.CompactTextString(m) }
func (*DashboardTabSummary) ProtoMessage()    {}
func (*DashboardTabSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7168d0e3f3f5589, []int{4}
}

func (m *DashboardTabSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DashboardTabSummary.Unmarshal(m, b)
}
func (m *DashboardTabSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DashboardTabSummary.Marshal(b, m, deterministic)
}
func (m *DashboardTabSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DashboardTabSummary.Merge(m, src)
}
func (m *DashboardTabSummary) XXX_Size() int {
	return xxx_messageInfo_DashboardTabSummary.Size(m)
}
func (m *DashboardTabSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_DashboardTabSummary.DiscardUnknown(m)
}

var xxx_messageInfo_DashboardTabSummary proto.InternalMessageInfo

func (m *DashboardTabSummary) GetDashboardName() string {
	if m != nil {
		return m.DashboardName
	}
	return ""
}

func (m *DashboardTabSummary) GetDashboardTabName() string {
	if m != nil {
		return m.DashboardTabName
	}
	return ""
}

func (m *DashboardTabSummary) GetAlert() string {
	if m != nil {
		return m.Alert
	}
	return ""
}

func (m *DashboardTabSummary) GetFailingTestSummaries() []*FailingTestSummary {
	if m != nil {
		return m.FailingTestSummaries
	}
	return nil
}

func (m *DashboardTabSummary) GetLastUpdateTimestamp() float64 {
	if m != nil {
		return m.LastUpdateTimestamp
	}
	return 0
}

func (m *DashboardTabSummary) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DashboardTabSummary) GetOverallStatus() DashboardTabSummary_TabStatus {
	if m != nil {
		return m.OverallStatus
	}
	return DashboardTabSummary_NOT_SET
}

func (m *DashboardTabSummary) GetLatestGreen() string {
	if m != nil {
		return m.LatestGreen
	}
	return ""
}

func (m *DashboardTabSummary) GetLastRunTimestamp() float64 {
	if m != nil {
		return m.LastRunTimestamp
	}
	return 0
}

func (m *DashboardTabSummary) GetBugUrl() string {
	if m != nil {
		return m.BugUrl
	}
	return ""
}

func (m *DashboardTabSummary) GetHealthiness() *HealthinessInfo {
	if m != nil {
		return m.Healthiness
	}
	return nil
}

func (m *DashboardTabSummary) GetLinkedIssues() []string {
	if m != nil {
		return m.LinkedIssues
	}
	return nil
}

func (m *DashboardTabSummary) GetAlertingData() *AlertingData {
	if m != nil {
		return m.AlertingData
	}
	return nil
}

// Summary state of a dashboard.
type DashboardSummary struct {
	// Summary of a dashboard tab; see config.proto.
	TabSummaries         []*DashboardTabSummary `protobuf:"bytes,1,rep,name=tab_summaries,json=tabSummaries,proto3" json:"tab_summaries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DashboardSummary) Reset()         { *m = DashboardSummary{} }
func (m *DashboardSummary) String() string { return proto.CompactTextString(m) }
func (*DashboardSummary) ProtoMessage()    {}
func (*DashboardSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7168d0e3f3f5589, []int{5}
}

func (m *DashboardSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DashboardSummary.Unmarshal(m, b)
}
func (m *DashboardSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DashboardSummary.Marshal(b, m, deterministic)
}
func (m *DashboardSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DashboardSummary.Merge(m, src)
}
func (m *DashboardSummary) XXX_Size() int {
	return xxx_messageInfo_DashboardSummary.Size(m)
}
func (m *DashboardSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_DashboardSummary.DiscardUnknown(m)
}

var xxx_messageInfo_DashboardSummary proto.InternalMessageInfo

func (m *DashboardSummary) GetTabSummaries() []*DashboardTabSummary {
	if m != nil {
		return m.TabSummaries
	}
	return nil
}

func init() {
	proto.RegisterEnum("TestInfo_Trend", TestInfo_Trend_name, TestInfo_Trend_value)
	proto.RegisterEnum("DashboardTabSummary_TabStatus", DashboardTabSummary_TabStatus_name, DashboardTabSummary_TabStatus_value)
	proto.RegisterType((*FailingTestSummary)(nil), "FailingTestSummary")
	proto.RegisterMapType((map[string]string)(nil), "FailingTestSummary.PropertiesEntry")
	proto.RegisterType((*TestInfo)(nil), "TestInfo")
	proto.RegisterMapType((map[string]int32)(nil), "TestInfo.InfraFailuresEntry")
	proto.RegisterType((*HealthinessInfo)(nil), "HealthinessInfo")
	proto.RegisterType((*AlertingData)(nil), "AlertingData")
	proto.RegisterType((*DashboardTabSummary)(nil), "DashboardTabSummary")
	proto.RegisterType((*DashboardSummary)(nil), "DashboardSummary")
}

func init() { proto.RegisterFile("summary.proto", fileDescriptor_f7168d0e3f3f5589) }

var fileDescriptor_f7168d0e3f3f5589 = []byte{
	// 1178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x61, 0x73, 0xdb, 0x44,
	0x13, 0xae, 0x63, 0xcb, 0x89, 0xd6, 0x96, 0xad, 0x5c, 0xd3, 0xbe, 0x7a, 0x43, 0xa1, 0xc1, 0xa5,
	0xe0, 0x81, 0xa2, 0x80, 0x19, 0x66, 0x80, 0x19, 0x66, 0x70, 0x52, 0xbb, 0x75, 0x9b, 0x3a, 0x1d,
	0xc5, 0x99, 0x0e, 0xc3, 0x07, 0xcd, 0xb9, 0x3a, 0xdb, 0x9a, 0xc8, 0x92, 0x47, 0x77, 0x0a, 0xcd,
	0x4f, 0xe1, 0x9f, 0xf0, 0x81, 0xdf, 0xc4, 0x6f, 0x60, 0x76, 0x4f, 0xb6, 0xd4, 0x34, 0xd0, 0x7e,
	0xd3, 0x3d, 0xfb, 0xec, 0xde, 0x6a, 0xf7, 0xd9, 0x5b, 0xb0, 0x64, 0xb6, 0x5c, 0xf2, 0xf4, 0xca,
	0x5d, 0xa5, 0x89, 0x4a, 0xf6, 0xef, 0xcf, 0x93, 0x64, 0x1e, 0x89, 0x43, 0x3a, 0x4d, 0xb3, 0xd9,
	0xa1, 0x0a, 0x97, 0x42, 0x2a, 0xbe, 0x5c, 0x69, 0x42, 0xe7, 0x2f, 0x03, 0xd8, 0x90, 0x87, 0x51,
	0x18, 0xcf, 0x27, 0x42, 0xaa, 0x33, 0xed, 0xcd, 0x3e, 0x85, 0x66, 0x10, 0xca, 0x55, 0xc4, 0xaf,
	0xfc, 0x98, 0x2f, 0x85, 0x53, 0x39, 0xa8, 0x74, 0x4d, 0xaf, 0x91, 0x63, 0x63, 0xbe, 0x14, 0xec,
	0x23, 0x30, 0x95, 0x90, 0x4a, 0xdb, 0xb7, 0xc8, 0xbe, 0x83, 0x00, 0x19, 0x3b, 0x60, 0xcd, 0x78,
	0x18, 0xf9, 0xd3, 0x2c, 0x8c, 0x02, 0x3f, 0x0c, 0x9c, 0xaa, 0x0e, 0x80, 0xe0, 0x11, 0x62, 0xa3,
	0x80, 0x3d, 0x84, 0x16, 0x71, 0x36, 0x29, 0x39, 0xb5, 0x83, 0x4a, 0xb7, 0xe2, 0x91, 0xe7, 0x64,
	0x0d, 0x62, 0xa8, 0x15, 0x97, 0xb2, 0x08, 0x65, 0xe8, 0x50, 0x08, 0x96, 0x42, 0x11, 0xa7, 0x08,
	0x55, 0xd7, 0xa1, 0x10, 0x2d, 0x42, 0x7d, 0x0c, 0x40, 0x37, 0xbe, 0x4e, 0xb2, 0x58, 0x39, 0xdb,
	0x07, 0x95, 0xae, 0xe1, 0x99, 0x88, 0x1c, 0x23, 0x80, 0x66, 0x7d, 0x49, 0x14, 0xc6, 0x17, 0xce,
	0x0e, 0x5d, 0x63, 0x12, 0x72, 0x12, 0xc6, 0x17, 0xec, 0x73, 0x68, 0x17, 0x66, 0x5f, 0x89, 0x37,
	0xca, 0x31, 0x89, 0x63, 0x6d, 0x38, 0x13, 0xf1, 0x46, 0xb1, 0xcf, 0xa0, 0xa5, 0x79, 0x59, 0x1a,
	0x69, 0x1a, 0x10, 0xad, 0x49, 0xe8, 0x79, 0x1a, 0x11, 0xeb, 0x0b, 0x68, 0xe3, 0xcd, 0x59, 0x2a,
	0xfc, 0xa5, 0x90, 0x92, 0xcf, 0x85, 0xd3, 0x20, 0x5a, 0x2b, 0x87, 0x5f, 0x68, 0x94, 0xdd, 0x87,
	0x06, 0x5e, 0x28, 0x02, 0x7f, 0x9a, 0xcd, 0xa5, 0xd3, 0x3c, 0xa8, 0x76, 0x4d, 0x0f, 0x34, 0x74,
	0x94, 0xcd, 0x25, 0xde, 0xa7, 0xeb, 0x88, 0xdd, 0xa0, 0xd4, 0x2d, 0x7d, 0x1f, 0xd5, 0x51, 0x48,
	0x45, 0xd9, 0x1f, 0xc2, 0x5e, 0xc4, 0x89, 0xf2, 0x76, 0x63, 0x5a, 0xc4, 0xdd, 0xd5, 0xb6, 0x61,
	0xa9, 0x3d, 0xc7, 0x00, 0xab, 0x34, 0x59, 0x89, 0x54, 0x85, 0x42, 0x3a, 0xed, 0x83, 0x6a, 0xb7,
	0xd1, 0x7b, 0xe0, 0xbe, 0xab, 0x15, 0xf7, 0xe5, 0x86, 0x35, 0x88, 0x55, 0x7a, 0xe5, 0x95, 0xdc,
	0x30, 0xf9, 0x45, 0xa2, 0xa2, 0x50, 0x2a, 0x3f, 0x0c, 0xa4, 0x63, 0xeb, 0xe4, 0x73, 0x68, 0x14,
	0xc8, 0xfd, 0x9f, 0xa1, 0x7d, 0xcd, 0x9f, 0xd9, 0x50, 0xbd, 0x10, 0x57, 0xb9, 0xe4, 0xf0, 0x93,
	0xed, 0x81, 0x71, 0xc9, 0xa3, 0x6c, 0x2d, 0x33, 0x7d, 0xf8, 0x69, 0xeb, 0x87, 0x4a, 0xe7, 0x0f,
	0x03, 0x76, 0x30, 0x97, 0x51, 0x3c, 0x4b, 0x3e, 0x44, 0xb4, 0x87, 0xb0, 0xa7, 0x12, 0xc5, 0x23,
	0x3f, 0x4e, 0x62, 0x3f, 0x8c, 0x67, 0x29, 0xf7, 0xd3, 0x2c, 0x96, 0x14, 0xd8, 0xf0, 0x76, 0xc9,
	0x36, 0x4e, 0xe2, 0x11, 0x5a, 0xbc, 0x2c, 0x96, 0xec, 0x5b, 0xb8, 0x83, 0x1a, 0x12, 0xc1, 0x75,
	0x8f, 0x2a, 0x79, 0x30, 0x6d, 0xbc, 0xee, 0x82, 0x25, 0x7e, 0xd7, 0xa5, 0xa6, 0x5d, 0xb4, 0xf1,
	0x2d, 0x97, 0x2f, 0x61, 0x37, 0x77, 0x29, 0xd1, 0x0d, 0xa2, 0xb7, 0xb5, 0xe1, 0xad, 0xf0, 0xfa,
	0x17, 0x90, 0xe4, 0xff, 0x1e, 0xaa, 0x85, 0x76, 0x22, 0xc9, 0x1b, 0x1e, 0x23, 0x23, 0x32, 0x5f,
	0x85, 0x6a, 0x41, 0x6e, 0x28, 0xec, 0x44, 0x2d, 0x44, 0xaa, 0xe3, 0xe6, 0xba, 0x27, 0x84, 0x22,
	0xde, 0x03, 0x73, 0x16, 0xf1, 0x8b, 0x30, 0x16, 0x52, 0x92, 0xec, 0xb7, 0xbc, 0x02, 0x60, 0x5f,
	0x03, 0x5b, 0xa5, 0xe2, 0x32, 0x4c, 0x32, 0xe9, 0x17, 0x34, 0x38, 0xa8, 0x76, 0xb7, 0xbc, 0xdd,
	0xb5, 0x65, 0xb8, 0xa1, 0x3f, 0x83, 0xff, 0xbf, 0x5e, 0xf0, 0x78, 0x2e, 0xfc, 0x59, 0x9a, 0x2c,
	0xfd, 0x88, 0x63, 0xeb, 0x63, 0x25, 0xd2, 0x4b, 0x1e, 0xd1, 0xbc, 0xb4, 0x7a, 0x6d, 0x77, 0xdd,
	0x32, 0x77, 0x92, 0x8a, 0x38, 0xf0, 0xee, 0x6a, 0x8f, 0x61, 0x9a, 0x2c, 0x4f, 0x38, 0x5a, 0x34,
	0x9d, 0x1d, 0x43, 0x4b, 0xd7, 0x23, 0x1f, 0x09, 0xe9, 0x34, 0x48, 0x86, 0xf7, 0x8a, 0x00, 0xf4,
	0x83, 0xc3, 0xdc, 0xac, 0xf5, 0x67, 0x85, 0x65, 0x6c, 0xff, 0x17, 0x60, 0xef, 0x92, 0xde, 0x27,
	0x32, 0xa3, 0x2c, 0xb2, 0xef, 0xc1, 0xa0, 0x3c, 0x59, 0x03, 0xb6, 0xcf, 0xc7, 0xcf, 0xc7, 0xa7,
	0xaf, 0xc6, 0xf6, 0x2d, 0x66, 0x81, 0x39, 0x3e, 0xf5, 0x8f, 0x9f, 0xf6, 0xc7, 0x4f, 0x06, 0x76,
	0x85, 0xd5, 0x61, 0xeb, 0xfc, 0xa5, 0xbd, 0xc5, 0x76, 0xa0, 0xf6, 0x18, 0x09, 0xd5, 0xce, 0xdf,
	0x15, 0x68, 0x3f, 0x15, 0x3c, 0x52, 0x0b, 0xaa, 0x0c, 0x49, 0xf4, 0x1b, 0x30, 0xa4, 0xe2, 0xa9,
	0xa2, 0x8b, 0x1b, 0xbd, 0x7d, 0x57, 0xbf, 0xcf, 0xee, 0xfa, 0x7d, 0x76, 0x37, 0x8f, 0x95, 0xa7,
	0x89, 0xec, 0x11, 0x54, 0x45, 0x1c, 0x50, 0x52, 0xff, 0xcd, 0x47, 0x1a, 0xbb, 0x0f, 0x06, 0xce,
	0x31, 0xca, 0x13, 0x0b, 0x65, 0x6e, 0x0a, 0xe5, 0x69, 0x9c, 0x7d, 0x05, 0xbb, 0xfc, 0x52, 0xa4,
	0x1c, 0xfb, 0xb3, 0x69, 0x66, 0x8d, 0x7a, 0x6e, 0xe7, 0x86, 0xe1, 0x7b, 0x5a, 0x6f, 0xfc, 0x4b,
	0xeb, 0x3b, 0x1e, 0x34, 0xfb, 0x11, 0x4e, 0x72, 0x3c, 0x7f, 0xcc, 0x15, 0x67, 0x47, 0xd0, 0xa6,
	0xf6, 0x8b, 0xe5, 0xfa, 0x99, 0xff, 0x80, 0xdf, 0xb6, 0xd0, 0x65, 0xb0, 0xcc, 0x57, 0x40, 0xe7,
	0x4f, 0x03, 0x6e, 0x3f, 0xe6, 0x72, 0x31, 0x4d, 0x78, 0x1a, 0x4c, 0xf8, 0x74, 0xbd, 0xa0, 0x1e,
	0x42, 0x2b, 0x58, 0xc3, 0xe5, 0x69, 0xb7, 0x36, 0x28, 0xcd, 0xfb, 0x23, 0x60, 0x05, 0x4d, 0xf1,
	0x69, 0x79, 0x5b, 0xd9, 0x41, 0x29, 0x2e, 0xb1, 0xf7, 0xc0, 0xe0, 0xf8, 0x03, 0xf9, 0xb6, 0xd2,
	0x07, 0x36, 0x82, 0xbb, 0x33, 0xfd, 0xea, 0xe9, 0x27, 0x56, 0x6f, 0x58, 0x7c, 0x14, 0x6b, 0x54,
	0xe4, 0xdb, 0x37, 0x3c, 0x8a, 0xde, 0xde, 0xec, 0x3a, 0x86, 0xcf, 0x61, 0x0f, 0xee, 0x50, 0x45,
	0xb2, 0x55, 0xc0, 0x95, 0x28, 0xad, 0x2b, 0x83, 0xd6, 0xd5, 0x6d, 0x34, 0x9e, 0x93, 0xad, 0x58,
	0x5a, 0x77, 0xa1, 0x2e, 0x15, 0x57, 0x99, 0xa4, 0x01, 0x37, 0xbd, 0xfc, 0xc4, 0x06, 0xd0, 0x4a,
	0xb0, 0x61, 0x51, 0xe4, 0xe7, 0xf6, 0x6d, 0x9a, 0xae, 0x4f, 0xdc, 0x1b, 0xea, 0xe5, 0xe2, 0x27,
	0xb1, 0x3c, 0x2b, 0xf7, 0xd2, 0x47, 0x7c, 0x34, 0xf3, 0xbd, 0x30, 0x4f, 0x85, 0x88, 0xf3, 0xb5,
	0xd7, 0xd0, 0xd8, 0x13, 0x84, 0xb0, 0x88, 0x94, 0x75, 0x9a, 0xc5, 0xa5, 0x94, 0x4d, 0x4a, 0xd9,
	0x46, 0x8b, 0x97, 0xc5, 0x45, 0xbe, 0xff, 0x83, 0xed, 0x69, 0x36, 0xc7, 0xe5, 0x97, 0xef, 0xbd,
	0xfa, 0x34, 0x9b, 0x9f, 0xa7, 0x11, 0xeb, 0x41, 0x63, 0x51, 0x8c, 0x83, 0xd3, 0x24, 0x29, 0xd8,
	0xee, 0xb5, 0x11, 0xf1, 0xca, 0x24, 0xf6, 0x00, 0xac, 0x7c, 0xf9, 0x85, 0x52, 0x66, 0x42, 0x3a,
	0x16, 0x6d, 0x90, 0xa6, 0x06, 0x47, 0x84, 0xb1, 0x1e, 0x58, 0x3c, 0xd7, 0x9d, 0x1f, 0x70, 0xc5,
	0x69, 0xa7, 0x35, 0x7a, 0x96, 0x5b, 0x56, 0xa3, 0xd7, 0xe4, 0xa5, 0x53, 0xe7, 0x37, 0x30, 0x37,
	0x25, 0xc1, 0xb9, 0x1e, 0x9f, 0x4e, 0xfc, 0xb3, 0xc1, 0xc4, 0xbe, 0x55, 0x1e, 0xf2, 0x0a, 0x4e,
	0xf3, 0xcb, 0xfe, 0xd9, 0x99, 0x9e, 0xeb, 0x61, 0x7f, 0x74, 0x62, 0x57, 0x99, 0x09, 0xc6, 0xf0,
	0xa4, 0xff, 0xfc, 0x57, 0xbb, 0x86, 0x9f, 0x67, 0x93, 0xfe, 0xc9, 0xc0, 0x36, 0x18, 0x40, 0xfd,
	0xc8, 0x3b, 0x7d, 0x3e, 0x18, 0xdb, 0xf5, 0x67, 0xb5, 0x9d, 0x86, 0xdd, 0xec, 0xbc, 0x00, 0x7b,
	0xd3, 0x89, 0xb5, 0x6c, 0x7f, 0x04, 0x0b, 0x55, 0x58, 0x48, 0xa8, 0x42, 0x12, 0xda, 0xbb, 0xa9,
	0x67, 0x5e, 0x53, 0xad, 0xbf, 0x43, 0x21, 0xa7, 0x75, 0x1a, 0x96, 0xef, 0xfe, 0x09, 0x00, 0x00,
	0xff, 0xff, 0xf7, 0x9a, 0x99, 0x54, 0xe2, 0x09, 0x00, 0x00,
}
