// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/stainless-sdks/tabstack-go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type AgentAction string                    // Always "agent:action"
type AgentExtracted string                 // Always "agent:extracted"
type AgentProcessing string                // Always "agent:processing"
type AgentReasoned string                  // Always "agent:reasoned"
type AgentStatus string                    // Always "agent:status"
type AgentStep string                      // Always "agent:step"
type AgentWaiting string                   // Always "agent:waiting"
type AIGeneration string                   // Always "ai:generation"
type AIGenerationError string              // Always "ai:generation:error"
type AnalyzingEnd string                   // Always "analyzing:end"
type AnalyzingStart string                 // Always "analyzing:start"
type Assistant string                      // Always "assistant"
type BrowserActionCompleted string         // Always "browser:action_completed"
type BrowserActionStarted string           // Always "browser:action_started"
type BrowserNavigated string               // Always "browser:navigated"
type BrowserReconnected string             // Always "browser:reconnected"
type BrowserScreenshotCaptured string      // Always "browser:screenshot_captured"
type BrowserScreenshotCapturedImage string // Always "browser:screenshot_captured_image"
type CdpEndpointConnected string           // Always "cdp:endpoint_connected"
type CdpEndpointCycle string               // Always "cdp:endpoint_cycle"
type Complete string                       // Always "complete"
type Content string                        // Always "content"
type Custom string                         // Always "custom"
type Done string                           // Always "done"
type Error string                          // Always "error"
type ErrorJson string                      // Always "error-json"
type ErrorText string                      // Always "error-text"
type EvaluatingEnd string                  // Always "evaluating:end"
type EvaluatingStart string                // Always "evaluating:start"
type ExecutionDenied string                // Always "execution-denied"
type File string                           // Always "file"
type FileData string                       // Always "file-data"
type FileID string                         // Always "file-id"
type FileURL string                        // Always "file-url"
type FollowingEnd string                   // Always "following:end"
type FollowingStart string                 // Always "following:start"
type Image string                          // Always "image"
type ImageData string                      // Always "image-data"
type ImageFileID string                    // Always "image-file-id"
type ImageURL string                       // Always "image-url"
type InteractiveFormDataError string       // Always "interactive:form_data:error"
type InteractiveFormDataRequest string     // Always "interactive:form_data:request"
type IterationEnd string                   // Always "iteration:end"
type IterationStart string                 // Always "iteration:start"
type Json string                           // Always "json"
type JudgingEnd string                     // Always "judging:end"
type JudgingStart string                   // Always "judging:start"
type Media string                          // Always "media"
type OutliningEnd string                   // Always "outlining:end"
type OutliningStart string                 // Always "outlining:start"
type PlanningEnd string                    // Always "planning:end"
type PlanningStart string                  // Always "planning:start"
type PrefetchingEnd string                 // Always "prefetching:end"
type PrefetchingStart string               // Always "prefetching:start"
type Reasoning string                      // Always "reasoning"
type SearchingEnd string                   // Always "searching:end"
type SearchingStart string                 // Always "searching:start"
type Start string                          // Always "start"
type System string                         // Always "system"
type SystemDebugCompression string         // Always "system:debug_compression"
type SystemDebugMessage string             // Always "system:debug_message"
type TaskAborted string                    // Always "task:aborted"
type TaskCompleted string                  // Always "task:completed"
type TaskMetrics string                    // Always "task:metrics"
type TaskMetricsIncremental string         // Always "task:metrics_incremental"
type TaskSetup string                      // Always "task:setup"
type TaskStarted string                    // Always "task:started"
type TaskValidated string                  // Always "task:validated"
type TaskValidationError string            // Always "task:validation_error"
type Text string                           // Always "text"
type Tool string                           // Always "tool"
type ToolApprovalRequest string            // Always "tool-approval-request"
type ToolApprovalResponse string           // Always "tool-approval-response"
type ToolCall string                       // Always "tool-call"
type ToolResult string                     // Always "tool-result"
type User string                           // Always "user"
type WritingEnd string                     // Always "writing:end"
type WritingStart string                   // Always "writing:start"

func (c AgentAction) Default() AgentAction                       { return "agent:action" }
func (c AgentExtracted) Default() AgentExtracted                 { return "agent:extracted" }
func (c AgentProcessing) Default() AgentProcessing               { return "agent:processing" }
func (c AgentReasoned) Default() AgentReasoned                   { return "agent:reasoned" }
func (c AgentStatus) Default() AgentStatus                       { return "agent:status" }
func (c AgentStep) Default() AgentStep                           { return "agent:step" }
func (c AgentWaiting) Default() AgentWaiting                     { return "agent:waiting" }
func (c AIGeneration) Default() AIGeneration                     { return "ai:generation" }
func (c AIGenerationError) Default() AIGenerationError           { return "ai:generation:error" }
func (c AnalyzingEnd) Default() AnalyzingEnd                     { return "analyzing:end" }
func (c AnalyzingStart) Default() AnalyzingStart                 { return "analyzing:start" }
func (c Assistant) Default() Assistant                           { return "assistant" }
func (c BrowserActionCompleted) Default() BrowserActionCompleted { return "browser:action_completed" }
func (c BrowserActionStarted) Default() BrowserActionStarted     { return "browser:action_started" }
func (c BrowserNavigated) Default() BrowserNavigated             { return "browser:navigated" }
func (c BrowserReconnected) Default() BrowserReconnected         { return "browser:reconnected" }
func (c BrowserScreenshotCaptured) Default() BrowserScreenshotCaptured {
	return "browser:screenshot_captured"
}
func (c BrowserScreenshotCapturedImage) Default() BrowserScreenshotCapturedImage {
	return "browser:screenshot_captured_image"
}
func (c CdpEndpointConnected) Default() CdpEndpointConnected { return "cdp:endpoint_connected" }
func (c CdpEndpointCycle) Default() CdpEndpointCycle         { return "cdp:endpoint_cycle" }
func (c Complete) Default() Complete                         { return "complete" }
func (c Content) Default() Content                           { return "content" }
func (c Custom) Default() Custom                             { return "custom" }
func (c Done) Default() Done                                 { return "done" }
func (c Error) Default() Error                               { return "error" }
func (c ErrorJson) Default() ErrorJson                       { return "error-json" }
func (c ErrorText) Default() ErrorText                       { return "error-text" }
func (c EvaluatingEnd) Default() EvaluatingEnd               { return "evaluating:end" }
func (c EvaluatingStart) Default() EvaluatingStart           { return "evaluating:start" }
func (c ExecutionDenied) Default() ExecutionDenied           { return "execution-denied" }
func (c File) Default() File                                 { return "file" }
func (c FileData) Default() FileData                         { return "file-data" }
func (c FileID) Default() FileID                             { return "file-id" }
func (c FileURL) Default() FileURL                           { return "file-url" }
func (c FollowingEnd) Default() FollowingEnd                 { return "following:end" }
func (c FollowingStart) Default() FollowingStart             { return "following:start" }
func (c Image) Default() Image                               { return "image" }
func (c ImageData) Default() ImageData                       { return "image-data" }
func (c ImageFileID) Default() ImageFileID                   { return "image-file-id" }
func (c ImageURL) Default() ImageURL                         { return "image-url" }
func (c InteractiveFormDataError) Default() InteractiveFormDataError {
	return "interactive:form_data:error"
}
func (c InteractiveFormDataRequest) Default() InteractiveFormDataRequest {
	return "interactive:form_data:request"
}
func (c IterationEnd) Default() IterationEnd                     { return "iteration:end" }
func (c IterationStart) Default() IterationStart                 { return "iteration:start" }
func (c Json) Default() Json                                     { return "json" }
func (c JudgingEnd) Default() JudgingEnd                         { return "judging:end" }
func (c JudgingStart) Default() JudgingStart                     { return "judging:start" }
func (c Media) Default() Media                                   { return "media" }
func (c OutliningEnd) Default() OutliningEnd                     { return "outlining:end" }
func (c OutliningStart) Default() OutliningStart                 { return "outlining:start" }
func (c PlanningEnd) Default() PlanningEnd                       { return "planning:end" }
func (c PlanningStart) Default() PlanningStart                   { return "planning:start" }
func (c PrefetchingEnd) Default() PrefetchingEnd                 { return "prefetching:end" }
func (c PrefetchingStart) Default() PrefetchingStart             { return "prefetching:start" }
func (c Reasoning) Default() Reasoning                           { return "reasoning" }
func (c SearchingEnd) Default() SearchingEnd                     { return "searching:end" }
func (c SearchingStart) Default() SearchingStart                 { return "searching:start" }
func (c Start) Default() Start                                   { return "start" }
func (c System) Default() System                                 { return "system" }
func (c SystemDebugCompression) Default() SystemDebugCompression { return "system:debug_compression" }
func (c SystemDebugMessage) Default() SystemDebugMessage         { return "system:debug_message" }
func (c TaskAborted) Default() TaskAborted                       { return "task:aborted" }
func (c TaskCompleted) Default() TaskCompleted                   { return "task:completed" }
func (c TaskMetrics) Default() TaskMetrics                       { return "task:metrics" }
func (c TaskMetricsIncremental) Default() TaskMetricsIncremental { return "task:metrics_incremental" }
func (c TaskSetup) Default() TaskSetup                           { return "task:setup" }
func (c TaskStarted) Default() TaskStarted                       { return "task:started" }
func (c TaskValidated) Default() TaskValidated                   { return "task:validated" }
func (c TaskValidationError) Default() TaskValidationError       { return "task:validation_error" }
func (c Text) Default() Text                                     { return "text" }
func (c Tool) Default() Tool                                     { return "tool" }
func (c ToolApprovalRequest) Default() ToolApprovalRequest       { return "tool-approval-request" }
func (c ToolApprovalResponse) Default() ToolApprovalResponse     { return "tool-approval-response" }
func (c ToolCall) Default() ToolCall                             { return "tool-call" }
func (c ToolResult) Default() ToolResult                         { return "tool-result" }
func (c User) Default() User                                     { return "user" }
func (c WritingEnd) Default() WritingEnd                         { return "writing:end" }
func (c WritingStart) Default() WritingStart                     { return "writing:start" }

func (c AgentAction) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c AgentExtracted) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c AgentProcessing) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c AgentReasoned) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c AgentStatus) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c AgentStep) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c AgentWaiting) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c AIGeneration) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c AIGenerationError) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c AnalyzingEnd) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c AnalyzingStart) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c Assistant) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c BrowserActionCompleted) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c BrowserActionStarted) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c BrowserNavigated) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c BrowserReconnected) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c BrowserScreenshotCaptured) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c BrowserScreenshotCapturedImage) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c CdpEndpointConnected) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c CdpEndpointCycle) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c Complete) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c Content) MarshalJSON() ([]byte, error)                        { return marshalString(c) }
func (c Custom) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c Done) MarshalJSON() ([]byte, error)                           { return marshalString(c) }
func (c Error) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c ErrorJson) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c ErrorText) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c EvaluatingEnd) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c EvaluatingStart) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c ExecutionDenied) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c File) MarshalJSON() ([]byte, error)                           { return marshalString(c) }
func (c FileData) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c FileID) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c FileURL) MarshalJSON() ([]byte, error)                        { return marshalString(c) }
func (c FollowingEnd) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c FollowingStart) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c Image) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c ImageData) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c ImageFileID) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c ImageURL) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c InteractiveFormDataError) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c InteractiveFormDataRequest) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c IterationEnd) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c IterationStart) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c Json) MarshalJSON() ([]byte, error)                           { return marshalString(c) }
func (c JudgingEnd) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c JudgingStart) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c Media) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c OutliningEnd) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c OutliningStart) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c PlanningEnd) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c PlanningStart) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c PrefetchingEnd) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c PrefetchingStart) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c Reasoning) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c SearchingEnd) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c SearchingStart) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c Start) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c System) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c SystemDebugCompression) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c SystemDebugMessage) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c TaskAborted) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c TaskCompleted) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c TaskMetrics) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c TaskMetricsIncremental) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c TaskSetup) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c TaskStarted) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c TaskValidated) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c TaskValidationError) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c Text) MarshalJSON() ([]byte, error)                           { return marshalString(c) }
func (c Tool) MarshalJSON() ([]byte, error)                           { return marshalString(c) }
func (c ToolApprovalRequest) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c ToolApprovalResponse) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c ToolCall) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c ToolResult) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c User) MarshalJSON() ([]byte, error)                           { return marshalString(c) }
func (c WritingEnd) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c WritingStart) MarshalJSON() ([]byte, error)                   { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
