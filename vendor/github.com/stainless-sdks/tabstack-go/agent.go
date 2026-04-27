// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tabstack

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/tabstack-go/internal/apijson"
	"github.com/stainless-sdks/tabstack-go/internal/requestconfig"
	"github.com/stainless-sdks/tabstack-go/option"
	"github.com/stainless-sdks/tabstack-go/packages/param"
	"github.com/stainless-sdks/tabstack-go/packages/respjson"
	"github.com/stainless-sdks/tabstack-go/packages/ssestream"
	"github.com/stainless-sdks/tabstack-go/shared/constant"
)

// AgentService contains methods and other services that help with interacting with
// the tabstack API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentService] method instead.
type AgentService struct {
	options []option.RequestOption
}

// NewAgentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAgentService(opts ...option.RequestOption) (r AgentService) {
	r = AgentService{}
	r.options = opts
	return
}

// Execute AI-powered browser automation tasks using natural language with optional
// geotargeting. This endpoint **always streams** responses using Server-Sent
// Events (SSE).
//
// **Streaming Response:**
//
// - All responses are streamed using Server-Sent Events (`text/event-stream`)
// - Real-time progress updates and results as they're generated
//
// **Geotargeting:**
//
// - Optionally specify a country code for geotargeted browsing
//
// **Use Cases:**
//
// - Web scraping and data extraction
// - Form filling and interaction
// - Navigation and information gathering
// - Multi-step web workflows
// - Content analysis from web pages
func (r *AgentService) AutomateStreaming(ctx context.Context, body AgentAutomateParams, opts ...option.RequestOption) (stream *ssestream.Stream[AutomateEventUnion]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	path := "automate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStreamWithSynthesizeEventData[AutomateEventUnion](ssestream.NewDecoder(raw), err)
}

// Submit a response to an interactive form data request from an in-progress
// automation task. When the AI agent encounters a form requiring user data, it
// emits an `interactive:form_data:request` or `interactive:form_data:error` SSE
// event containing a `requestId`. Use this endpoint to provide the requested data
// or cancel the request.
//
// **Lifecycle:**
//
//   - Input requests expire after 2 minutes by default
//   - Expired or already-answered requests return `410 Gone`
//   - Successful submissions return `202 Accepted` (fire-and-forget from caller's
//     perspective)
func (r *AgentService) AutomateInput(ctx context.Context, requestID string, body AgentAutomateInputParams, opts ...option.RequestOption) (res *AgentAutomateInputResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if requestID == "" {
		err = errors.New("missing required requestID parameter")
		return nil, err
	}
	path := fmt.Sprintf("automate/%s/input", url.PathEscape(requestID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Execute AI-powered research queries that search the web, analyze sources, and
// synthesize comprehensive answers. This endpoint **always streams** responses
// using Server-Sent Events (SSE).
//
// **Streaming Response:**
//
// - All responses are streamed using Server-Sent Events (`text/event-stream`)
// - Real-time progress updates as research progresses through phases
//
// **Research Modes:**
//
// - `fast` - Quick answers with minimal web searches (default)
// - `balanced` - Standard research with multiple iterations
//
// **Use Cases:**
//
// - Answering complex questions with cited sources
// - Synthesizing information from multiple web sources
// - Research reports on specific topics
// - Fact-checking and verification tasks
func (r *AgentService) ResearchStreaming(ctx context.Context, body AgentResearchParams, opts ...option.RequestOption) (stream *ssestream.Stream[ResearchEventUnion]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	path := "research"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStreamWithSynthesizeEventData[ResearchEventUnion](ssestream.NewDecoder(raw), err)
}

// AutomateEventUnion contains all possible properties and values from
// [AutomateEventAgentAction], [AutomateEventAgentExtracted],
// [AutomateEventAgentProcessing], [AutomateEventAgentReasoned],
// [AutomateEventAgentStatus], [AutomateEventAgentStep],
// [AutomateEventAgentWaiting], [AutomateEventAIGeneration],
// [AutomateEventAIGenerationError], [AutomateEventBrowserActionCompleted],
// [AutomateEventBrowserActionStarted], [AutomateEventBrowserNavigated],
// [AutomateEventBrowserReconnected], [AutomateEventBrowserScreenshotCaptured],
// [AutomateEventBrowserScreenshotCapturedImage],
// [AutomateEventCdpEndpointConnected], [AutomateEventCdpEndpointCycle],
// [AutomateEventComplete], [AutomateEventDone], [AutomateEventError],
// [AutomateEventInteractiveFormDataError],
// [AutomateEventInteractiveFormDataRequest],
// [AutomateEventSystemDebugCompression], [AutomateEventSystemDebugMessage],
// [AutomateEventTaskAborted], [AutomateEventTaskCompleted],
// [AutomateEventTaskMetrics], [AutomateEventTaskMetricsIncremental],
// [AutomateEventTaskSetup], [AutomateEventTaskStarted],
// [AutomateEventTaskValidated], [AutomateEventTaskValidationError].
//
// Use the [AutomateEventUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomateEventUnion struct {
	// This field is a union of [AutomateEventAgentActionData],
	// [AutomateEventAgentExtractedData], [AutomateEventAgentProcessingData],
	// [AutomateEventAgentReasonedData], [AutomateEventAgentStatusData],
	// [AutomateEventAgentStepData], [AutomateEventAgentWaitingData],
	// [AutomateEventAIGenerationData], [AutomateEventAIGenerationErrorData],
	// [AutomateEventBrowserActionCompletedData],
	// [AutomateEventBrowserActionStartedData], [AutomateEventBrowserNavigatedData],
	// [AutomateEventBrowserReconnectedData],
	// [AutomateEventBrowserScreenshotCapturedData],
	// [AutomateEventBrowserScreenshotCapturedImageData],
	// [AutomateEventCdpEndpointConnectedData], [AutomateEventCdpEndpointCycleData],
	// [AutomateEventCompleteData], [map[string]any], [AutomateEventErrorData],
	// [AutomateEventInteractiveFormDataErrorData],
	// [AutomateEventInteractiveFormDataRequestData],
	// [AutomateEventSystemDebugCompressionData],
	// [AutomateEventSystemDebugMessageData], [AutomateEventTaskAbortedData],
	// [AutomateEventTaskCompletedData], [AutomateEventTaskMetricsData],
	// [AutomateEventTaskMetricsIncrementalData], [AutomateEventTaskSetupData],
	// [AutomateEventTaskStartedData], [AutomateEventTaskValidatedData],
	// [AutomateEventTaskValidationErrorData]
	Data AutomateEventUnionData `json:"data"`
	// Any of "agent:action", "agent:extracted", "agent:processing", "agent:reasoned",
	// "agent:status", "agent:step", "agent:waiting", "ai:generation",
	// "ai:generation:error", "browser:action_completed", "browser:action_started",
	// "browser:navigated", "browser:reconnected", "browser:screenshot_captured",
	// "browser:screenshot_captured_image", "cdp:endpoint_connected",
	// "cdp:endpoint_cycle", "complete", "done", "error",
	// "interactive:form_data:error", "interactive:form_data:request",
	// "system:debug_compression", "system:debug_message", "task:aborted",
	// "task:completed", "task:metrics", "task:metrics_incremental", "task:setup",
	// "task:started", "task:validated", "task:validation_error".
	Event string `json:"event"`
	JSON  struct {
		Data  respjson.Field
		Event respjson.Field
		raw   string
	} `json:"-"`
}

// anyAutomateEvent is implemented by each variant of [AutomateEventUnion] to add
// type safety for the return type of [AutomateEventUnion.AsAny]
type anyAutomateEvent interface {
	implAutomateEventUnion()
}

func (AutomateEventAgentAction) implAutomateEventUnion()                    {}
func (AutomateEventAgentExtracted) implAutomateEventUnion()                 {}
func (AutomateEventAgentProcessing) implAutomateEventUnion()                {}
func (AutomateEventAgentReasoned) implAutomateEventUnion()                  {}
func (AutomateEventAgentStatus) implAutomateEventUnion()                    {}
func (AutomateEventAgentStep) implAutomateEventUnion()                      {}
func (AutomateEventAgentWaiting) implAutomateEventUnion()                   {}
func (AutomateEventAIGeneration) implAutomateEventUnion()                   {}
func (AutomateEventAIGenerationError) implAutomateEventUnion()              {}
func (AutomateEventBrowserActionCompleted) implAutomateEventUnion()         {}
func (AutomateEventBrowserActionStarted) implAutomateEventUnion()           {}
func (AutomateEventBrowserNavigated) implAutomateEventUnion()               {}
func (AutomateEventBrowserReconnected) implAutomateEventUnion()             {}
func (AutomateEventBrowserScreenshotCaptured) implAutomateEventUnion()      {}
func (AutomateEventBrowserScreenshotCapturedImage) implAutomateEventUnion() {}
func (AutomateEventCdpEndpointConnected) implAutomateEventUnion()           {}
func (AutomateEventCdpEndpointCycle) implAutomateEventUnion()               {}
func (AutomateEventComplete) implAutomateEventUnion()                       {}
func (AutomateEventDone) implAutomateEventUnion()                           {}
func (AutomateEventError) implAutomateEventUnion()                          {}
func (AutomateEventInteractiveFormDataError) implAutomateEventUnion()       {}
func (AutomateEventInteractiveFormDataRequest) implAutomateEventUnion()     {}
func (AutomateEventSystemDebugCompression) implAutomateEventUnion()         {}
func (AutomateEventSystemDebugMessage) implAutomateEventUnion()             {}
func (AutomateEventTaskAborted) implAutomateEventUnion()                    {}
func (AutomateEventTaskCompleted) implAutomateEventUnion()                  {}
func (AutomateEventTaskMetrics) implAutomateEventUnion()                    {}
func (AutomateEventTaskMetricsIncremental) implAutomateEventUnion()         {}
func (AutomateEventTaskSetup) implAutomateEventUnion()                      {}
func (AutomateEventTaskStarted) implAutomateEventUnion()                    {}
func (AutomateEventTaskValidated) implAutomateEventUnion()                  {}
func (AutomateEventTaskValidationError) implAutomateEventUnion()            {}

// Use the following switch statement to find the correct variant
//
//	switch variant := AutomateEventUnion.AsAny().(type) {
//	case tabstack.AutomateEventAgentAction:
//	case tabstack.AutomateEventAgentExtracted:
//	case tabstack.AutomateEventAgentProcessing:
//	case tabstack.AutomateEventAgentReasoned:
//	case tabstack.AutomateEventAgentStatus:
//	case tabstack.AutomateEventAgentStep:
//	case tabstack.AutomateEventAgentWaiting:
//	case tabstack.AutomateEventAIGeneration:
//	case tabstack.AutomateEventAIGenerationError:
//	case tabstack.AutomateEventBrowserActionCompleted:
//	case tabstack.AutomateEventBrowserActionStarted:
//	case tabstack.AutomateEventBrowserNavigated:
//	case tabstack.AutomateEventBrowserReconnected:
//	case tabstack.AutomateEventBrowserScreenshotCaptured:
//	case tabstack.AutomateEventBrowserScreenshotCapturedImage:
//	case tabstack.AutomateEventCdpEndpointConnected:
//	case tabstack.AutomateEventCdpEndpointCycle:
//	case tabstack.AutomateEventComplete:
//	case tabstack.AutomateEventDone:
//	case tabstack.AutomateEventError:
//	case tabstack.AutomateEventInteractiveFormDataError:
//	case tabstack.AutomateEventInteractiveFormDataRequest:
//	case tabstack.AutomateEventSystemDebugCompression:
//	case tabstack.AutomateEventSystemDebugMessage:
//	case tabstack.AutomateEventTaskAborted:
//	case tabstack.AutomateEventTaskCompleted:
//	case tabstack.AutomateEventTaskMetrics:
//	case tabstack.AutomateEventTaskMetricsIncremental:
//	case tabstack.AutomateEventTaskSetup:
//	case tabstack.AutomateEventTaskStarted:
//	case tabstack.AutomateEventTaskValidated:
//	case tabstack.AutomateEventTaskValidationError:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u AutomateEventUnion) AsAny() anyAutomateEvent {
	switch u.Event {
	case "agent:action":
		return u.AsAgentAction()
	case "agent:extracted":
		return u.AsAgentExtracted()
	case "agent:processing":
		return u.AsAgentProcessing()
	case "agent:reasoned":
		return u.AsAgentReasoned()
	case "agent:status":
		return u.AsAgentStatus()
	case "agent:step":
		return u.AsAgentStep()
	case "agent:waiting":
		return u.AsAgentWaiting()
	case "ai:generation":
		return u.AsAIGeneration()
	case "ai:generation:error":
		return u.AsAIGenerationError()
	case "browser:action_completed":
		return u.AsBrowserActionCompleted()
	case "browser:action_started":
		return u.AsBrowserActionStarted()
	case "browser:navigated":
		return u.AsBrowserNavigated()
	case "browser:reconnected":
		return u.AsBrowserReconnected()
	case "browser:screenshot_captured":
		return u.AsBrowserScreenshotCaptured()
	case "browser:screenshot_captured_image":
		return u.AsBrowserScreenshotCapturedImage()
	case "cdp:endpoint_connected":
		return u.AsCdpEndpointConnected()
	case "cdp:endpoint_cycle":
		return u.AsCdpEndpointCycle()
	case "complete":
		return u.AsComplete()
	case "done":
		return u.AsDone()
	case "error":
		return u.AsError()
	case "interactive:form_data:error":
		return u.AsInteractiveFormDataError()
	case "interactive:form_data:request":
		return u.AsInteractiveFormDataRequest()
	case "system:debug_compression":
		return u.AsSystemDebugCompression()
	case "system:debug_message":
		return u.AsSystemDebugMessage()
	case "task:aborted":
		return u.AsTaskAborted()
	case "task:completed":
		return u.AsTaskCompleted()
	case "task:metrics":
		return u.AsTaskMetrics()
	case "task:metrics_incremental":
		return u.AsTaskMetricsIncremental()
	case "task:setup":
		return u.AsTaskSetup()
	case "task:started":
		return u.AsTaskStarted()
	case "task:validated":
		return u.AsTaskValidated()
	case "task:validation_error":
		return u.AsTaskValidationError()
	}
	return nil
}

func (u AutomateEventUnion) AsAgentAction() (v AutomateEventAgentAction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsAgentExtracted() (v AutomateEventAgentExtracted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsAgentProcessing() (v AutomateEventAgentProcessing) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsAgentReasoned() (v AutomateEventAgentReasoned) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsAgentStatus() (v AutomateEventAgentStatus) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsAgentStep() (v AutomateEventAgentStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsAgentWaiting() (v AutomateEventAgentWaiting) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsAIGeneration() (v AutomateEventAIGeneration) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsAIGenerationError() (v AutomateEventAIGenerationError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsBrowserActionCompleted() (v AutomateEventBrowserActionCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsBrowserActionStarted() (v AutomateEventBrowserActionStarted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsBrowserNavigated() (v AutomateEventBrowserNavigated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsBrowserReconnected() (v AutomateEventBrowserReconnected) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsBrowserScreenshotCaptured() (v AutomateEventBrowserScreenshotCaptured) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsBrowserScreenshotCapturedImage() (v AutomateEventBrowserScreenshotCapturedImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsCdpEndpointConnected() (v AutomateEventCdpEndpointConnected) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsCdpEndpointCycle() (v AutomateEventCdpEndpointCycle) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsComplete() (v AutomateEventComplete) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsDone() (v AutomateEventDone) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsError() (v AutomateEventError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsInteractiveFormDataError() (v AutomateEventInteractiveFormDataError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsInteractiveFormDataRequest() (v AutomateEventInteractiveFormDataRequest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsSystemDebugCompression() (v AutomateEventSystemDebugCompression) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsSystemDebugMessage() (v AutomateEventSystemDebugMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsTaskAborted() (v AutomateEventTaskAborted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsTaskCompleted() (v AutomateEventTaskCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsTaskMetrics() (v AutomateEventTaskMetrics) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsTaskMetricsIncremental() (v AutomateEventTaskMetricsIncremental) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsTaskSetup() (v AutomateEventTaskSetup) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsTaskStarted() (v AutomateEventTaskStarted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsTaskValidated() (v AutomateEventTaskValidated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventUnion) AsTaskValidationError() (v AutomateEventTaskValidationError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomateEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventUnionData is an implicit subunion of [AutomateEventUnion].
// AutomateEventUnionData provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [AutomateEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAutomateEventDoneData]
type AutomateEventUnionData struct {
	// This field will be present if the value is a [any] instead of an object.
	OfAutomateEventDoneData any     `json:",inline"`
	Action                  string  `json:"action"`
	IterationID             string  `json:"iterationId"`
	Timestamp               float64 `json:"timestamp"`
	Ref                     string  `json:"ref"`
	Value                   string  `json:"value"`
	// This field is from variant [AutomateEventAgentExtractedData].
	ExtractedData string `json:"extractedData"`
	// This field is from variant [AutomateEventAgentProcessingData].
	HasScreenshot bool `json:"hasScreenshot"`
	// This field is from variant [AutomateEventAgentProcessingData].
	Operation string `json:"operation"`
	// This field is from variant [AutomateEventAgentReasonedData].
	Reasoning string `json:"reasoning"`
	// This field is from variant [AutomateEventAgentStatusData].
	Message string `json:"message"`
	// This field is from variant [AutomateEventAgentStepData].
	CurrentIteration float64 `json:"currentIteration"`
	// This field is from variant [AutomateEventAgentWaitingData].
	Seconds float64 `json:"seconds"`
	// This field is from variant [AutomateEventAIGenerationData].
	FinishReason string `json:"finishReason"`
	Prompt       string `json:"prompt"`
	Schema       any    `json:"schema"`
	// This field is from variant [AutomateEventAIGenerationData].
	Usage AutomateEventAIGenerationDataUsage `json:"usage"`
	// This field is a union of [[]AutomateEventAIGenerationDataMessageUnion], [[]any],
	// [[]any]
	Messages AutomateEventUnionDataMessages `json:"messages"`
	// This field is from variant [AutomateEventAIGenerationData].
	Object any `json:"object"`
	// This field is from variant [AutomateEventAIGenerationData].
	ProviderMetadata map[string]any `json:"providerMetadata"`
	// This field is from variant [AutomateEventAIGenerationData].
	Temperature float64 `json:"temperature"`
	// This field is from variant [AutomateEventAIGenerationData].
	Warnings []any `json:"warnings"`
	// This field is a union of [string], [string], [string],
	// [AutomateEventCompleteDataError], [AutomateEventErrorDataError]
	Error   AutomateEventUnionDataError `json:"error"`
	Success bool                        `json:"success"`
	// This field is from variant [AutomateEventBrowserNavigatedData].
	Title         string  `json:"title"`
	URL           string  `json:"url"`
	EndpointIndex float64 `json:"endpointIndex"`
	// This field is from variant [AutomateEventBrowserReconnectedData].
	StartingURL string  `json:"startingUrl"`
	Total       float64 `json:"total"`
	// This field is from variant [AutomateEventBrowserScreenshotCapturedData].
	Format string `json:"format"`
	// This field is from variant [AutomateEventBrowserScreenshotCapturedData].
	Size float64 `json:"size"`
	// This field is from variant [AutomateEventBrowserScreenshotCapturedImageData].
	Image string `json:"image"`
	// This field is from variant [AutomateEventBrowserScreenshotCapturedImageData].
	MediaType string `json:"mediaType"`
	// This field is from variant [AutomateEventCdpEndpointCycleData].
	Attempt     float64 `json:"attempt"`
	FinalAnswer string  `json:"finalAnswer"`
	// This field is from variant [AutomateEventCompleteData].
	Stats AutomateEventCompleteDataStats `json:"stats"`
	// This field is from variant [AutomateEventInteractiveFormDataErrorData].
	FieldErrors map[string]string `json:"fieldErrors"`
	// This field is a union of [[]AutomateEventInteractiveFormDataErrorDataField],
	// [[]AutomateEventInteractiveFormDataRequestDataField]
	Fields          AutomateEventUnionDataFields `json:"fields"`
	FormDescription string                       `json:"formDescription"`
	PageTitle       string                       `json:"pageTitle"`
	PageURL         string                       `json:"pageUrl"`
	RequestID       string                       `json:"requestId"`
	// This field is from variant [AutomateEventSystemDebugCompressionData].
	CompressedSize float64 `json:"compressedSize"`
	// This field is from variant [AutomateEventSystemDebugCompressionData].
	CompressionPercent float64 `json:"compressionPercent"`
	// This field is from variant [AutomateEventSystemDebugCompressionData].
	OriginalSize float64 `json:"originalSize"`
	// This field is from variant [AutomateEventTaskAbortedData].
	Reason                 string  `json:"reason"`
	AIGenerationCount      float64 `json:"aiGenerationCount"`
	AIGenerationErrorCount float64 `json:"aiGenerationErrorCount"`
	EventCounts            float64 `json:"eventCounts"`
	StepCount              float64 `json:"stepCount"`
	TotalInputTokens       float64 `json:"totalInputTokens"`
	TotalOutputTokens      float64 `json:"totalOutputTokens"`
	// This field is from variant [AutomateEventTaskSetupData].
	BrowserName string `json:"browserName"`
	Task        string `json:"task"`
	// This field is from variant [AutomateEventTaskSetupData].
	Data any `json:"data"`
	// This field is from variant [AutomateEventTaskSetupData].
	Guardrails string `json:"guardrails"`
	// This field is from variant [AutomateEventTaskSetupData].
	HasAPIKey bool `json:"hasApiKey"`
	// This field is from variant [AutomateEventTaskSetupData].
	KeySource string `json:"keySource"`
	// This field is from variant [AutomateEventTaskSetupData].
	Model string `json:"model"`
	// This field is from variant [AutomateEventTaskSetupData].
	Provider string `json:"provider"`
	// This field is from variant [AutomateEventTaskSetupData].
	Proxy string `json:"proxy"`
	// This field is from variant [AutomateEventTaskSetupData].
	PwCdpEndpoint string `json:"pwCdpEndpoint"`
	// This field is from variant [AutomateEventTaskSetupData].
	PwCdpEndpointCount float64 `json:"pwCdpEndpointCount"`
	// This field is from variant [AutomateEventTaskSetupData].
	PwCdpEndpoints []string `json:"pwCdpEndpoints"`
	// This field is from variant [AutomateEventTaskSetupData].
	PwEndpoint string `json:"pwEndpoint"`
	// This field is from variant [AutomateEventTaskSetupData].
	Vision bool `json:"vision"`
	// This field is from variant [AutomateEventTaskStartedData].
	Plan string `json:"plan"`
	// This field is from variant [AutomateEventTaskStartedData].
	SuccessCriteria string `json:"successCriteria"`
	// This field is from variant [AutomateEventTaskStartedData].
	ActionItems []string `json:"actionItems"`
	// This field is from variant [AutomateEventTaskValidatedData].
	CompletionQuality string `json:"completionQuality"`
	// This field is from variant [AutomateEventTaskValidatedData].
	Observation string `json:"observation"`
	// This field is from variant [AutomateEventTaskValidatedData].
	Feedback string `json:"feedback"`
	// This field is from variant [AutomateEventTaskValidationErrorData].
	Errors []string `json:"errors"`
	// This field is from variant [AutomateEventTaskValidationErrorData].
	RawResponse any `json:"rawResponse"`
	// This field is from variant [AutomateEventTaskValidationErrorData].
	RetryCount float64 `json:"retryCount"`
	JSON       struct {
		OfAutomateEventDoneData respjson.Field
		Action                  respjson.Field
		IterationID             respjson.Field
		Timestamp               respjson.Field
		Ref                     respjson.Field
		Value                   respjson.Field
		ExtractedData           respjson.Field
		HasScreenshot           respjson.Field
		Operation               respjson.Field
		Reasoning               respjson.Field
		Message                 respjson.Field
		CurrentIteration        respjson.Field
		Seconds                 respjson.Field
		FinishReason            respjson.Field
		Prompt                  respjson.Field
		Schema                  respjson.Field
		Usage                   respjson.Field
		Messages                respjson.Field
		Object                  respjson.Field
		ProviderMetadata        respjson.Field
		Temperature             respjson.Field
		Warnings                respjson.Field
		Error                   respjson.Field
		Success                 respjson.Field
		Title                   respjson.Field
		URL                     respjson.Field
		EndpointIndex           respjson.Field
		StartingURL             respjson.Field
		Total                   respjson.Field
		Format                  respjson.Field
		Size                    respjson.Field
		Image                   respjson.Field
		MediaType               respjson.Field
		Attempt                 respjson.Field
		FinalAnswer             respjson.Field
		Stats                   respjson.Field
		FieldErrors             respjson.Field
		Fields                  respjson.Field
		FormDescription         respjson.Field
		PageTitle               respjson.Field
		PageURL                 respjson.Field
		RequestID               respjson.Field
		CompressedSize          respjson.Field
		CompressionPercent      respjson.Field
		OriginalSize            respjson.Field
		Reason                  respjson.Field
		AIGenerationCount       respjson.Field
		AIGenerationErrorCount  respjson.Field
		EventCounts             respjson.Field
		StepCount               respjson.Field
		TotalInputTokens        respjson.Field
		TotalOutputTokens       respjson.Field
		BrowserName             respjson.Field
		Task                    respjson.Field
		Data                    respjson.Field
		Guardrails              respjson.Field
		HasAPIKey               respjson.Field
		KeySource               respjson.Field
		Model                   respjson.Field
		Provider                respjson.Field
		Proxy                   respjson.Field
		PwCdpEndpoint           respjson.Field
		PwCdpEndpointCount      respjson.Field
		PwCdpEndpoints          respjson.Field
		PwEndpoint              respjson.Field
		Vision                  respjson.Field
		Plan                    respjson.Field
		SuccessCriteria         respjson.Field
		ActionItems             respjson.Field
		CompletionQuality       respjson.Field
		Observation             respjson.Field
		Feedback                respjson.Field
		Errors                  respjson.Field
		RawResponse             respjson.Field
		RetryCount              respjson.Field
		raw                     string
	} `json:"-"`
}

func (r *AutomateEventUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventUnionDataMessages is an implicit subunion of [AutomateEventUnion].
// AutomateEventUnionDataMessages provides convenient access to the sub-properties
// of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AutomateEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAutomateEventAIGenerationDataMessages OfAnyArray]
type AutomateEventUnionDataMessages struct {
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageUnion] instead of an object.
	OfAutomateEventAIGenerationDataMessages []AutomateEventAIGenerationDataMessageUnion `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfAutomateEventAIGenerationDataMessages respjson.Field
		OfAnyArray                              respjson.Field
		raw                                     string
	} `json:"-"`
}

func (r *AutomateEventUnionDataMessages) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventUnionDataError is an implicit subunion of [AutomateEventUnion].
// AutomateEventUnionDataError provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [AutomateEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type AutomateEventUnionDataError struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	Code     string `json:"code"`
	Message  string `json:"message"`
	// This field is from variant [AutomateEventErrorDataError].
	Timestamp string `json:"timestamp"`
	JSON      struct {
		OfString  respjson.Field
		Code      respjson.Field
		Message   respjson.Field
		Timestamp respjson.Field
		raw       string
	} `json:"-"`
}

func (r *AutomateEventUnionDataError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventUnionDataFields is an implicit subunion of [AutomateEventUnion].
// AutomateEventUnionDataFields provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [AutomateEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAutomateEventInteractiveFormDataErrorDataFields
// OfAutomateEventInteractiveFormDataRequestDataFields]
type AutomateEventUnionDataFields struct {
	// This field will be present if the value is a
	// [[]AutomateEventInteractiveFormDataErrorDataField] instead of an object.
	OfAutomateEventInteractiveFormDataErrorDataFields []AutomateEventInteractiveFormDataErrorDataField `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventInteractiveFormDataRequestDataField] instead of an object.
	OfAutomateEventInteractiveFormDataRequestDataFields []AutomateEventInteractiveFormDataRequestDataField `json:",inline"`
	JSON                                                struct {
		OfAutomateEventInteractiveFormDataErrorDataFields   respjson.Field
		OfAutomateEventInteractiveFormDataRequestDataFields respjson.Field
		raw                                                 string
	} `json:"-"`
}

func (r *AutomateEventUnionDataFields) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "agent:action" event from /v1/automate.
type AutomateEventAgentAction struct {
	// Event data for action execution
	Data  AutomateEventAgentActionData `json:"data" api:"required"`
	Event constant.AgentAction         `json:"event" default:"agent:action"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAgentAction) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventAgentAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data for action execution
type AutomateEventAgentActionData struct {
	Action      string  `json:"action" api:"required"`
	IterationID string  `json:"iterationId" api:"required"`
	Timestamp   float64 `json:"timestamp" api:"required"`
	Ref         string  `json:"ref" api:"nullable"`
	Value       string  `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		IterationID respjson.Field
		Timestamp   respjson.Field
		Ref         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAgentActionData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventAgentActionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "agent:extracted" event from /v1/automate.
type AutomateEventAgentExtracted struct {
	// Event data for extracted data
	Data  AutomateEventAgentExtractedData `json:"data" api:"required"`
	Event constant.AgentExtracted         `json:"event" default:"agent:extracted"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAgentExtracted) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventAgentExtracted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data for extracted data
type AutomateEventAgentExtractedData struct {
	ExtractedData string  `json:"extractedData" api:"required"`
	IterationID   string  `json:"iterationId" api:"required"`
	Timestamp     float64 `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtractedData respjson.Field
		IterationID   respjson.Field
		Timestamp     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAgentExtractedData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventAgentExtractedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "agent:processing" event from /v1/automate.
type AutomateEventAgentProcessing struct {
	// Event data for when the agent is waiting for model generation
	Data  AutomateEventAgentProcessingData `json:"data" api:"required"`
	Event constant.AgentProcessing         `json:"event" default:"agent:processing"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAgentProcessing) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventAgentProcessing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data for when the agent is waiting for model generation
type AutomateEventAgentProcessingData struct {
	HasScreenshot bool    `json:"hasScreenshot" api:"required"`
	IterationID   string  `json:"iterationId" api:"required"`
	Operation     string  `json:"operation" api:"required"`
	Timestamp     float64 `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasScreenshot respjson.Field
		IterationID   respjson.Field
		Operation     respjson.Field
		Timestamp     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAgentProcessingData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventAgentProcessingData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "agent:reasoned" event from /v1/automate.
type AutomateEventAgentReasoned struct {
	// Event data for agent reasoning
	Data  AutomateEventAgentReasonedData `json:"data" api:"required"`
	Event constant.AgentReasoned         `json:"event" default:"agent:reasoned"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAgentReasoned) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventAgentReasoned) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data for agent reasoning
type AutomateEventAgentReasonedData struct {
	IterationID string  `json:"iterationId" api:"required"`
	Reasoning   string  `json:"reasoning" api:"required"`
	Timestamp   float64 `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IterationID respjson.Field
		Reasoning   respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAgentReasonedData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventAgentReasonedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "agent:status" event from /v1/automate.
type AutomateEventAgentStatus struct {
	// Event data for status messages
	Data  AutomateEventAgentStatusData `json:"data" api:"required"`
	Event constant.AgentStatus         `json:"event" default:"agent:status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAgentStatus) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventAgentStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data for status messages
type AutomateEventAgentStatusData struct {
	IterationID string  `json:"iterationId" api:"required"`
	Message     string  `json:"message" api:"required"`
	Timestamp   float64 `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IterationID respjson.Field
		Message     respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAgentStatusData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventAgentStatusData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "agent:step" event from /v1/automate.
type AutomateEventAgentStep struct {
	// Event data for agent step tracking (each loop iteration)
	Data  AutomateEventAgentStepData `json:"data" api:"required"`
	Event constant.AgentStep         `json:"event" default:"agent:step"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAgentStep) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventAgentStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data for agent step tracking (each loop iteration)
type AutomateEventAgentStepData struct {
	CurrentIteration float64 `json:"currentIteration" api:"required"`
	IterationID      string  `json:"iterationId" api:"required"`
	Timestamp        float64 `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentIteration respjson.Field
		IterationID      respjson.Field
		Timestamp        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAgentStepData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventAgentStepData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "agent:waiting" event from /v1/automate.
type AutomateEventAgentWaiting struct {
	// Event data for waiting notifications
	Data  AutomateEventAgentWaitingData `json:"data" api:"required"`
	Event constant.AgentWaiting         `json:"event" default:"agent:waiting"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAgentWaiting) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventAgentWaiting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data for waiting notifications
type AutomateEventAgentWaitingData struct {
	IterationID string  `json:"iterationId" api:"required"`
	Seconds     float64 `json:"seconds" api:"required"`
	Timestamp   float64 `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IterationID respjson.Field
		Seconds     respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAgentWaitingData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventAgentWaitingData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "ai:generation" event from /v1/automate.
type AutomateEventAIGeneration struct {
	// Event data when AI generation occurs
	Data  AutomateEventAIGenerationData `json:"data" api:"required"`
	Event constant.AIGeneration         `json:"event" default:"ai:generation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGeneration) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventAIGeneration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data when AI generation occurs
type AutomateEventAIGenerationData struct {
	// Any of "stop", "length", "content-filter", "tool-calls", "error", "other".
	FinishReason     string                                      `json:"finishReason" api:"required"`
	IterationID      string                                      `json:"iterationId" api:"required"`
	Prompt           string                                      `json:"prompt" api:"required"`
	Schema           any                                         `json:"schema" api:"required"`
	Timestamp        float64                                     `json:"timestamp" api:"required"`
	Usage            AutomateEventAIGenerationDataUsage          `json:"usage" api:"required"`
	Messages         []AutomateEventAIGenerationDataMessageUnion `json:"messages"`
	Object           any                                         `json:"object"`
	ProviderMetadata map[string]any                              `json:"providerMetadata"`
	Temperature      float64                                     `json:"temperature"`
	Warnings         []any                                       `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FinishReason     respjson.Field
		IterationID      respjson.Field
		Prompt           respjson.Field
		Schema           respjson.Field
		Timestamp        respjson.Field
		Usage            respjson.Field
		Messages         respjson.Field
		Object           respjson.Field
		ProviderMetadata respjson.Field
		Temperature      respjson.Field
		Warnings         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventAIGenerationData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataUsage struct {
	InputTokens  float64 `json:"inputTokens"`
	OutputTokens float64 `json:"outputTokens"`
	TotalTokens  float64 `json:"totalTokens"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputTokens  respjson.Field
		OutputTokens respjson.Field
		TotalTokens  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataUsage) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventAIGenerationDataUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageUnion contains all possible properties and
// values from [AutomateEventAIGenerationDataMessageSystem],
// [AutomateEventAIGenerationDataMessageUser],
// [AutomateEventAIGenerationDataMessageAssistant],
// [AutomateEventAIGenerationDataMessageTool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomateEventAIGenerationDataMessageUnion struct {
	// This field is a union of [string],
	// [AutomateEventAIGenerationDataMessageUserContentUnion],
	// [AutomateEventAIGenerationDataMessageAssistantContentUnion],
	// [[]AutomateEventAIGenerationDataMessageToolContentUnion]
	Content AutomateEventAIGenerationDataMessageUnionContent `json:"content"`
	Role    string                                           `json:"role"`
	// This field is a union of
	// [map[string]map[string]AutomateEventAIGenerationDataMessageSystemProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageUserProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageAssistantProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageToolProviderOptionUnion]
	ProviderOptions AutomateEventAIGenerationDataMessageUnionProviderOptions `json:"providerOptions"`
	JSON            struct {
		Content         respjson.Field
		Role            respjson.Field
		ProviderOptions respjson.Field
		raw             string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageUnion) AsSystem() (v AutomateEventAIGenerationDataMessageSystem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUnion) AsUser() (v AutomateEventAIGenerationDataMessageUser) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUnion) AsAssistant() (v AutomateEventAIGenerationDataMessageAssistant) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUnion) AsTool() (v AutomateEventAIGenerationDataMessageTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomateEventAIGenerationDataMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageUnionContent is an implicit subunion of
// [AutomateEventAIGenerationDataMessageUnion].
// AutomateEventAIGenerationDataMessageUnionContent provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AutomateEventAIGenerationDataMessageUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfAutomateEventAIGenerationDataMessageUserContentArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArray
// OfAutomateEventAIGenerationDataMessageToolContentArray]
type AutomateEventAIGenerationDataMessageUnionContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageUserContentArrayItemUnion] instead of an
	// object.
	OfAutomateEventAIGenerationDataMessageUserContentArray []AutomateEventAIGenerationDataMessageUserContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemUnion] instead
	// of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentUnion] instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentArray []AutomateEventAIGenerationDataMessageToolContentUnion `json:",inline"`
	JSON                                                   struct {
		OfString                                                    respjson.Field
		OfAutomateEventAIGenerationDataMessageUserContentArray      respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArray respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentArray      respjson.Field
		raw                                                         string
	} `json:"-"`
}

func (r *AutomateEventAIGenerationDataMessageUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageUnionProviderOptions is an implicit subunion
// of [AutomateEventAIGenerationDataMessageUnion].
// AutomateEventAIGenerationDataMessageUnionProviderOptions provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AutomateEventAIGenerationDataMessageUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageSystemProviderOptionArray
// OfAutomateEventAIGenerationDataMessageUserProviderOptionArray
// OfAutomateEventAIGenerationDataMessageAssistantProviderOptionArray
// OfAutomateEventAIGenerationDataMessageToolProviderOptionArray]
type AutomateEventAIGenerationDataMessageUnionProviderOptions struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageSystemProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageSystemProviderOptionArray []AutomateEventAIGenerationDataMessageSystemProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageUserProviderOptionArrayItemUnion] instead
	// of an object.
	OfAutomateEventAIGenerationDataMessageUserProviderOptionArray []AutomateEventAIGenerationDataMessageUserProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolProviderOptionArrayItemUnion] instead
	// of an object.
	OfAutomateEventAIGenerationDataMessageToolProviderOptionArray []AutomateEventAIGenerationDataMessageToolProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                          struct {
		OfString                                                           respjson.Field
		OfFloat                                                            respjson.Field
		OfBool                                                             respjson.Field
		OfAnyArray                                                         respjson.Field
		OfAutomateEventAIGenerationDataMessageSystemProviderOptionArray    respjson.Field
		OfAutomateEventAIGenerationDataMessageUserProviderOptionArray      respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantProviderOptionArray respjson.Field
		OfAutomateEventAIGenerationDataMessageToolProviderOptionArray      respjson.Field
		raw                                                                string
	} `json:"-"`
}

func (r *AutomateEventAIGenerationDataMessageUnionProviderOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A system message. It can contain system information.
//
// Note: using the "system" part of the prompt is strongly preferred to increase
// the resilience against prompt injection attacks, and because not all providers
// support several system messages.
type AutomateEventAIGenerationDataMessageSystem struct {
	Content string          `json:"content" api:"required"`
	Role    constant.System `json:"role" default:"system"`
	// Additional provider-specific metadata. They are passed through to the provider
	// from the AI SDK and enable provider-specific functionality that can be fully
	// encapsulated in the provider.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageSystemProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content         respjson.Field
		Role            respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageSystem) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventAIGenerationDataMessageSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageSystemProviderOptionUnion contains all
// possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageSystemProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageSystemProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageSystemProviderOptionArray]
type AutomateEventAIGenerationDataMessageSystemProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageSystemProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageSystemProviderOptionArray []AutomateEventAIGenerationDataMessageSystemProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                            struct {
		OfString                                                        respjson.Field
		OfFloat                                                         respjson.Field
		OfBool                                                          respjson.Field
		OfAnyArray                                                      respjson.Field
		OfAutomateEventAIGenerationDataMessageSystemProviderOptionArray respjson.Field
		raw                                                             string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageSystemProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageSystemProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageSystemProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageSystemProviderOptionUnion) AsAutomateEventAIGenerationDataMessageSystemProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageSystemProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageSystemProviderOptionUnion) AsAutomateEventAIGenerationDataMessageSystemProviderOptionArray() (v []AutomateEventAIGenerationDataMessageSystemProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageSystemProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageSystemProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageSystemProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageSystemProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageSystemProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageSystemProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageSystemProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageSystemProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageSystemProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageSystemProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageSystemProviderOptionArrayItemUnion contains
// all possible properties and values from [string], [float64], [bool], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageSystemProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageSystemProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageSystemProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageSystemProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageSystemProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageSystemProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageSystemProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A user message. It can contain text or a combination of text and images.
type AutomateEventAIGenerationDataMessageUser struct {
	// Content of a user message. It can be a string or an array of text and image
	// parts.
	Content AutomateEventAIGenerationDataMessageUserContentUnion `json:"content" api:"required"`
	Role    constant.User                                        `json:"role" default:"user"`
	// Additional provider-specific metadata. They are passed through to the provider
	// from the AI SDK and enable provider-specific functionality that can be fully
	// encapsulated in the provider.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageUserProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content         respjson.Field
		Role            respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageUser) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventAIGenerationDataMessageUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageUserContentUnion contains all possible
// properties and values from [string],
// [[]AutomateEventAIGenerationDataMessageUserContentArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfAutomateEventAIGenerationDataMessageUserContentArray]
type AutomateEventAIGenerationDataMessageUserContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageUserContentArrayItemUnion] instead of an
	// object.
	OfAutomateEventAIGenerationDataMessageUserContentArray []AutomateEventAIGenerationDataMessageUserContentArrayItemUnion `json:",inline"`
	JSON                                                   struct {
		OfString                                               respjson.Field
		OfAutomateEventAIGenerationDataMessageUserContentArray respjson.Field
		raw                                                    string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageUserContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentUnion) AsAutomateEventAIGenerationDataMessageUserContentArray() (v []AutomateEventAIGenerationDataMessageUserContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageUserContentUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomateEventAIGenerationDataMessageUserContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageUserContentArrayItemUnion contains all
// possible properties and values from
// [AutomateEventAIGenerationDataMessageUserContentArrayItemText],
// [AutomateEventAIGenerationDataMessageUserContentArrayItemImage],
// [AutomateEventAIGenerationDataMessageUserContentArrayItemFile].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomateEventAIGenerationDataMessageUserContentArrayItemUnion struct {
	// This field is from variant
	// [AutomateEventAIGenerationDataMessageUserContentArrayItemText].
	Text string `json:"text"`
	Type string `json:"type"`
	// This field is a union of
	// [map[string]map[string]AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionUnion]
	ProviderOptions AutomateEventAIGenerationDataMessageUserContentArrayItemUnionProviderOptions `json:"providerOptions"`
	// This field is from variant
	// [AutomateEventAIGenerationDataMessageUserContentArrayItemImage].
	Image     AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageUnion `json:"image"`
	MediaType string                                                                  `json:"mediaType"`
	// This field is from variant
	// [AutomateEventAIGenerationDataMessageUserContentArrayItemFile].
	Data AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataUnion `json:"data"`
	// This field is from variant
	// [AutomateEventAIGenerationDataMessageUserContentArrayItemFile].
	Filename string `json:"filename"`
	JSON     struct {
		Text            respjson.Field
		Type            respjson.Field
		ProviderOptions respjson.Field
		Image           respjson.Field
		MediaType       respjson.Field
		Data            respjson.Field
		Filename        respjson.Field
		raw             string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemUnion) AsText() (v AutomateEventAIGenerationDataMessageUserContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemUnion) AsImage() (v AutomateEventAIGenerationDataMessageUserContentArrayItemImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemUnion) AsFile() (v AutomateEventAIGenerationDataMessageUserContentArrayItemFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageUserContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageUserContentArrayItemUnionProviderOptions is
// an implicit subunion of
// [AutomateEventAIGenerationDataMessageUserContentArrayItemUnion].
// AutomateEventAIGenerationDataMessageUserContentArrayItemUnionProviderOptions
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AutomateEventAIGenerationDataMessageUserContentArrayItemUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionArray
// OfAutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionArray
// OfAutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionArray]
type AutomateEventAIGenerationDataMessageUserContentArrayItemUnionProviderOptions struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionArray []AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionArray []AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionArray []AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                              struct {
		OfString                                                                           respjson.Field
		OfFloat                                                                            respjson.Field
		OfBool                                                                             respjson.Field
		OfAnyArray                                                                         respjson.Field
		OfAutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionArray  respjson.Field
		OfAutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionArray respjson.Field
		OfAutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionArray  respjson.Field
		raw                                                                                string
	} `json:"-"`
}

func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemUnionProviderOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part of a prompt. It contains a string of text.
type AutomateEventAIGenerationDataMessageUserContentArrayItemText struct {
	// The text content.
	Text string        `json:"text" api:"required"`
	Type constant.Text `json:"type" default:"text"`
	// Additional provider-specific metadata. They are passed through to the provider
	// from the AI SDK and enable provider-specific functionality that can be fully
	// encapsulated in the provider.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text            respjson.Field
		Type            respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageUserContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionArray]
type AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionArray []AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                              struct {
		OfString                                                                          respjson.Field
		OfFloat                                                                           respjson.Field
		OfBool                                                                            respjson.Field
		OfAnyArray                                                                        respjson.Field
		OfAutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionArray respjson.Field
		raw                                                                               string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionUnion) AsAutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionUnion) AsAutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionArray() (v []AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemTextProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content part of a prompt. It contains an image.
type AutomateEventAIGenerationDataMessageUserContentArrayItemImage struct {
	// Image data. Can either be:
	//
	// - data: a base64-encoded string, a Uint8Array, an ArrayBuffer, or a Buffer
	// - URL: a URL that points to the image
	Image AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageUnion `json:"image" api:"required"`
	Type  constant.Image                                                          `json:"type" default:"image"`
	// Optional IANA media type of the image.
	MediaType string `json:"mediaType"`
	// Additional provider-specific metadata. They are passed through to the provider
	// from the AI SDK and enable provider-specific functionality that can be fully
	// encapsulated in the provider.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image           respjson.Field
		Type            respjson.Field
		MediaType       respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageUserContentArrayItemImage) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageUnion contains
// all possible properties and values from [string],
// [AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageObject],
// [AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageByteLength],
// [AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageV1GlobalBuffer].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is a union of
	// [AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageObjectBuffer],
	// [AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageV1GlobalBufferBuffer]
	Buffer          AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageUnionBuffer `json:"buffer"`
	ByteLength      float64                                                                       `json:"byteLength"`
	ByteOffset      float64                                                                       `json:"byteOffset"`
	BytesPerElement float64                                                                       `json:"BYTES_PER_ELEMENT"`
	Length          float64                                                                       `json:"length"`
	JSON            struct {
		OfString        respjson.Field
		Buffer          respjson.Field
		ByteLength      respjson.Field
		ByteOffset      respjson.Field
		BytesPerElement respjson.Field
		Length          respjson.Field
		raw             string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageUnion) AsAutomateEventAIGenerationDataMessageUserContentArrayItemImageImageObject() (v AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageUnion) AsAutomateEventAIGenerationDataMessageUserContentArrayItemImageImageByteLength() (v AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageByteLength) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageUnion) AsAutomateEventAIGenerationDataMessageUserContentArrayItemImageImageV1GlobalBuffer() (v AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageV1GlobalBuffer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageUnionBuffer is
// an implicit subunion of
// [AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageUnion].
// AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageUnionBuffer
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageUnion].
type AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageUnionBuffer struct {
	ByteLength float64 `json:"byteLength"`
	JSON       struct {
		ByteLength respjson.Field
		raw        string
	} `json:"-"`
}

func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageUnionBuffer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageObject struct {
	Buffer          AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageObjectBuffer `json:"buffer" api:"required"`
	ByteLength      float64                                                                        `json:"byteLength" api:"required"`
	ByteOffset      float64                                                                        `json:"byteOffset" api:"required"`
	BytesPerElement float64                                                                        `json:"BYTES_PER_ELEMENT" api:"required"`
	Length          float64                                                                        `json:"length" api:"required"`
	ExtraFields     map[string]float64                                                             `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Buffer          respjson.Field
		ByteLength      respjson.Field
		ByteOffset      respjson.Field
		BytesPerElement respjson.Field
		Length          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageObject) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageObjectBuffer struct {
	ByteLength float64 `json:"byteLength" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByteLength  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageObjectBuffer) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageObjectBuffer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageByteLength struct {
	ByteLength float64 `json:"byteLength" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByteLength  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageByteLength) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageByteLength) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageV1GlobalBuffer struct {
	Buffer          AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageV1GlobalBufferBuffer `json:"buffer" api:"required"`
	ByteLength      float64                                                                                `json:"byteLength" api:"required"`
	ByteOffset      float64                                                                                `json:"byteOffset" api:"required"`
	BytesPerElement float64                                                                                `json:"BYTES_PER_ELEMENT" api:"required"`
	Length          float64                                                                                `json:"length" api:"required"`
	ExtraFields     map[string]float64                                                                     `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Buffer          respjson.Field
		ByteLength      respjson.Field
		ByteOffset      respjson.Field
		BytesPerElement respjson.Field
		Length          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageV1GlobalBuffer) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageV1GlobalBuffer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageV1GlobalBufferBuffer struct {
	ByteLength float64 `json:"byteLength" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByteLength  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageV1GlobalBufferBuffer) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemImageImageV1GlobalBufferBuffer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionArray]
type AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionArray []AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                               struct {
		OfString                                                                           respjson.Field
		OfFloat                                                                            respjson.Field
		OfBool                                                                             respjson.Field
		OfAnyArray                                                                         respjson.Field
		OfAutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionArray respjson.Field
		raw                                                                                string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionUnion) AsAutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionUnion) AsAutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionArray() (v []AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemImageProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File content part of a prompt. It contains a file.
type AutomateEventAIGenerationDataMessageUserContentArrayItemFile struct {
	// File data. Can either be:
	//
	// - data: a base64-encoded string, a Uint8Array, an ArrayBuffer, or a Buffer
	// - URL: a URL that points to the image
	Data AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataUnion `json:"data" api:"required"`
	// IANA media type of the file.
	MediaType string        `json:"mediaType" api:"required"`
	Type      constant.File `json:"type" default:"file"`
	// Optional filename of the file.
	Filename string `json:"filename"`
	// Additional provider-specific metadata. They are passed through to the provider
	// from the AI SDK and enable provider-specific functionality that can be fully
	// encapsulated in the provider.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data            respjson.Field
		MediaType       respjson.Field
		Type            respjson.Field
		Filename        respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageUserContentArrayItemFile) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataUnion contains
// all possible properties and values from [string],
// [AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataObject],
// [AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataByteLength],
// [AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataV1GlobalBuffer].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is a union of
	// [AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataObjectBuffer],
	// [AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataV1GlobalBufferBuffer]
	Buffer          AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataUnionBuffer `json:"buffer"`
	ByteLength      float64                                                                     `json:"byteLength"`
	ByteOffset      float64                                                                     `json:"byteOffset"`
	BytesPerElement float64                                                                     `json:"BYTES_PER_ELEMENT"`
	Length          float64                                                                     `json:"length"`
	JSON            struct {
		OfString        respjson.Field
		Buffer          respjson.Field
		ByteLength      respjson.Field
		ByteOffset      respjson.Field
		BytesPerElement respjson.Field
		Length          respjson.Field
		raw             string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataUnion) AsAutomateEventAIGenerationDataMessageUserContentArrayItemFileDataObject() (v AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataUnion) AsAutomateEventAIGenerationDataMessageUserContentArrayItemFileDataByteLength() (v AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataByteLength) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataUnion) AsAutomateEventAIGenerationDataMessageUserContentArrayItemFileDataV1GlobalBuffer() (v AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataV1GlobalBuffer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataUnionBuffer is
// an implicit subunion of
// [AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataUnion].
// AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataUnionBuffer
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataUnion].
type AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataUnionBuffer struct {
	ByteLength float64 `json:"byteLength"`
	JSON       struct {
		ByteLength respjson.Field
		raw        string
	} `json:"-"`
}

func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataUnionBuffer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataObject struct {
	Buffer          AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataObjectBuffer `json:"buffer" api:"required"`
	ByteLength      float64                                                                      `json:"byteLength" api:"required"`
	ByteOffset      float64                                                                      `json:"byteOffset" api:"required"`
	BytesPerElement float64                                                                      `json:"BYTES_PER_ELEMENT" api:"required"`
	Length          float64                                                                      `json:"length" api:"required"`
	ExtraFields     map[string]float64                                                           `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Buffer          respjson.Field
		ByteLength      respjson.Field
		ByteOffset      respjson.Field
		BytesPerElement respjson.Field
		Length          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataObject) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataObjectBuffer struct {
	ByteLength float64 `json:"byteLength" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByteLength  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataObjectBuffer) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataObjectBuffer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataByteLength struct {
	ByteLength float64 `json:"byteLength" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByteLength  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataByteLength) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataByteLength) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataV1GlobalBuffer struct {
	Buffer          AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataV1GlobalBufferBuffer `json:"buffer" api:"required"`
	ByteLength      float64                                                                              `json:"byteLength" api:"required"`
	ByteOffset      float64                                                                              `json:"byteOffset" api:"required"`
	BytesPerElement float64                                                                              `json:"BYTES_PER_ELEMENT" api:"required"`
	Length          float64                                                                              `json:"length" api:"required"`
	ExtraFields     map[string]float64                                                                   `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Buffer          respjson.Field
		ByteLength      respjson.Field
		ByteOffset      respjson.Field
		BytesPerElement respjson.Field
		Length          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataV1GlobalBuffer) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataV1GlobalBuffer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataV1GlobalBufferBuffer struct {
	ByteLength float64 `json:"byteLength" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByteLength  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataV1GlobalBufferBuffer) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemFileDataV1GlobalBufferBuffer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionArray]
type AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionArray []AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                              struct {
		OfString                                                                          respjson.Field
		OfFloat                                                                           respjson.Field
		OfBool                                                                            respjson.Field
		OfAnyArray                                                                        respjson.Field
		OfAutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionArray respjson.Field
		raw                                                                               string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionUnion) AsAutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionUnion) AsAutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionArray() (v []AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageUserContentArrayItemFileProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageUserProviderOptionUnion contains all
// possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageUserProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageUserProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageUserProviderOptionArray]
type AutomateEventAIGenerationDataMessageUserProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageUserProviderOptionArrayItemUnion] instead
	// of an object.
	OfAutomateEventAIGenerationDataMessageUserProviderOptionArray []AutomateEventAIGenerationDataMessageUserProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                          struct {
		OfString                                                      respjson.Field
		OfFloat                                                       respjson.Field
		OfBool                                                        respjson.Field
		OfAnyArray                                                    respjson.Field
		OfAutomateEventAIGenerationDataMessageUserProviderOptionArray respjson.Field
		raw                                                           string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageUserProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserProviderOptionUnion) AsAutomateEventAIGenerationDataMessageUserProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageUserProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserProviderOptionUnion) AsAutomateEventAIGenerationDataMessageUserProviderOptionArray() (v []AutomateEventAIGenerationDataMessageUserProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageUserProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageUserProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageUserProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageUserProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageUserProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageUserProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageUserProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageUserProviderOptionArrayItemUnion contains
// all possible properties and values from [string], [float64], [bool], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageUserProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageUserProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageUserProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageUserProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageUserProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An assistant message. It can contain text, tool calls, or a combination of text
// and tool calls.
type AutomateEventAIGenerationDataMessageAssistant struct {
	// Content of an assistant message. It can be a string or an array of text, image,
	// reasoning, redacted reasoning, and tool call parts.
	Content AutomateEventAIGenerationDataMessageAssistantContentUnion `json:"content" api:"required"`
	Role    constant.Assistant                                        `json:"role" default:"assistant"`
	// Additional provider-specific metadata. They are passed through to the provider
	// from the AI SDK and enable provider-specific functionality that can be fully
	// encapsulated in the provider.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageAssistantProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content         respjson.Field
		Role            respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistant) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventAIGenerationDataMessageAssistant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentUnion contains all possible
// properties and values from [string],
// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfAutomateEventAIGenerationDataMessageAssistantContentArray]
type AutomateEventAIGenerationDataMessageAssistantContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemUnion] instead
	// of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemUnion `json:",inline"`
	JSON                                                        struct {
		OfString                                                    respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArray respjson.Field
		raw                                                         string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArray() (v []AutomateEventAIGenerationDataMessageAssistantContentArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemUnion contains all
// possible properties and values from
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemText],
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemFile],
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoning],
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCall],
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResult],
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolApprovalRequest].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemUnion struct {
	Text string `json:"text"`
	Type string `json:"type"`
	// This field is a union of
	// [map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionUnion]
	ProviderOptions AutomateEventAIGenerationDataMessageAssistantContentArrayItemUnionProviderOptions `json:"providerOptions"`
	// This field is from variant
	// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemFile].
	Data AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataUnion `json:"data"`
	// This field is from variant
	// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemFile].
	MediaType string `json:"mediaType"`
	// This field is from variant
	// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemFile].
	Filename string `json:"filename"`
	// This field is from variant
	// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCall].
	Input      any    `json:"input"`
	ToolCallID string `json:"toolCallId"`
	ToolName   string `json:"toolName"`
	// This field is from variant
	// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCall].
	ProviderExecuted bool `json:"providerExecuted"`
	// This field is from variant
	// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResult].
	Output AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputUnion `json:"output"`
	// This field is from variant
	// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolApprovalRequest].
	ApprovalID string `json:"approvalId"`
	JSON       struct {
		Text             respjson.Field
		Type             respjson.Field
		ProviderOptions  respjson.Field
		Data             respjson.Field
		MediaType        respjson.Field
		Filename         respjson.Field
		Input            respjson.Field
		ToolCallID       respjson.Field
		ToolName         respjson.Field
		ProviderExecuted respjson.Field
		Output           respjson.Field
		ApprovalID       respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemUnion) AsText() (v AutomateEventAIGenerationDataMessageAssistantContentArrayItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemUnion) AsFile() (v AutomateEventAIGenerationDataMessageAssistantContentArrayItemFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemUnion) AsReasoning() (v AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoning) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemUnion) AsToolCall() (v AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemUnion) AsToolResult() (v AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemUnion) AsToolApprovalRequest() (v AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolApprovalRequest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemUnionProviderOptions
// is an implicit subunion of
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemUnion].
// AutomateEventAIGenerationDataMessageAssistantContentArrayItemUnionProviderOptions
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemUnionProviderOptions struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                         struct {
		OfString                                                                                     respjson.Field
		OfFloat                                                                                      respjson.Field
		OfBool                                                                                       respjson.Field
		OfAnyArray                                                                                   respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionArray       respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionArray       respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionArray  respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionArray   respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionArray respjson.Field
		raw                                                                                          string
	} `json:"-"`
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemUnionProviderOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text content part of a prompt. It contains a string of text.
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemText struct {
	// The text content.
	Text string        `json:"text" api:"required"`
	Type constant.Text `json:"type" default:"text"`
	// Additional provider-specific metadata. They are passed through to the provider
	// from the AI SDK and enable provider-specific functionality that can be fully
	// encapsulated in the provider.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text            respjson.Field
		Type            respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistantContentArrayItemText) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                   struct {
		OfString                                                                               respjson.Field
		OfFloat                                                                                respjson.Field
		OfBool                                                                                 respjson.Field
		OfAnyArray                                                                             respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionArray respjson.Field
		raw                                                                                    string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionArray() (v []AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemTextProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// File content part of a prompt. It contains a file.
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemFile struct {
	// File data. Can either be:
	//
	// - data: a base64-encoded string, a Uint8Array, an ArrayBuffer, or a Buffer
	// - URL: a URL that points to the image
	Data AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataUnion `json:"data" api:"required"`
	// IANA media type of the file.
	MediaType string        `json:"mediaType" api:"required"`
	Type      constant.File `json:"type" default:"file"`
	// Optional filename of the file.
	Filename string `json:"filename"`
	// Additional provider-specific metadata. They are passed through to the provider
	// from the AI SDK and enable provider-specific functionality that can be fully
	// encapsulated in the provider.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data            respjson.Field
		MediaType       respjson.Field
		Type            respjson.Field
		Filename        respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistantContentArrayItemFile) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataUnion
// contains all possible properties and values from [string],
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataObject],
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataByteLength],
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataV1GlobalBuffer].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is a union of
	// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataObjectBuffer],
	// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataV1GlobalBufferBuffer]
	Buffer          AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataUnionBuffer `json:"buffer"`
	ByteLength      float64                                                                          `json:"byteLength"`
	ByteOffset      float64                                                                          `json:"byteOffset"`
	BytesPerElement float64                                                                          `json:"BYTES_PER_ELEMENT"`
	Length          float64                                                                          `json:"length"`
	JSON            struct {
		OfString        respjson.Field
		Buffer          respjson.Field
		ByteLength      respjson.Field
		ByteOffset      respjson.Field
		BytesPerElement respjson.Field
		Length          respjson.Field
		raw             string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataObject() (v AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataByteLength() (v AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataByteLength) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataV1GlobalBuffer() (v AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataV1GlobalBuffer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataUnionBuffer
// is an implicit subunion of
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataUnion].
// AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataUnionBuffer
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataUnion].
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataUnionBuffer struct {
	ByteLength float64 `json:"byteLength"`
	JSON       struct {
		ByteLength respjson.Field
		raw        string
	} `json:"-"`
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataUnionBuffer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataObject struct {
	Buffer          AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataObjectBuffer `json:"buffer" api:"required"`
	ByteLength      float64                                                                           `json:"byteLength" api:"required"`
	ByteOffset      float64                                                                           `json:"byteOffset" api:"required"`
	BytesPerElement float64                                                                           `json:"BYTES_PER_ELEMENT" api:"required"`
	Length          float64                                                                           `json:"length" api:"required"`
	ExtraFields     map[string]float64                                                                `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Buffer          respjson.Field
		ByteLength      respjson.Field
		ByteOffset      respjson.Field
		BytesPerElement respjson.Field
		Length          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataObject) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataObjectBuffer struct {
	ByteLength float64 `json:"byteLength" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByteLength  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataObjectBuffer) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataObjectBuffer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataByteLength struct {
	ByteLength float64 `json:"byteLength" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByteLength  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataByteLength) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataByteLength) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataV1GlobalBuffer struct {
	Buffer          AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataV1GlobalBufferBuffer `json:"buffer" api:"required"`
	ByteLength      float64                                                                                   `json:"byteLength" api:"required"`
	ByteOffset      float64                                                                                   `json:"byteOffset" api:"required"`
	BytesPerElement float64                                                                                   `json:"BYTES_PER_ELEMENT" api:"required"`
	Length          float64                                                                                   `json:"length" api:"required"`
	ExtraFields     map[string]float64                                                                        `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Buffer          respjson.Field
		ByteLength      respjson.Field
		ByteOffset      respjson.Field
		BytesPerElement respjson.Field
		Length          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataV1GlobalBuffer) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataV1GlobalBuffer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataV1GlobalBufferBuffer struct {
	ByteLength float64 `json:"byteLength" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByteLength  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataV1GlobalBufferBuffer) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileDataV1GlobalBufferBuffer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                   struct {
		OfString                                                                               respjson.Field
		OfFloat                                                                                respjson.Field
		OfBool                                                                                 respjson.Field
		OfAnyArray                                                                             respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionArray respjson.Field
		raw                                                                                    string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionArray() (v []AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemFileProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reasoning content part of a prompt. It contains a reasoning.
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoning struct {
	// The reasoning text.
	Text string             `json:"text" api:"required"`
	Type constant.Reasoning `json:"type" default:"reasoning"`
	// Additional provider-specific metadata. They are passed through to the provider
	// from the AI SDK and enable provider-specific functionality that can be fully
	// encapsulated in the provider.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text            respjson.Field
		Type            respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoning) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                        struct {
		OfString                                                                                    respjson.Field
		OfFloat                                                                                     respjson.Field
		OfBool                                                                                      respjson.Field
		OfAnyArray                                                                                  respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionArray respjson.Field
		raw                                                                                         string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionArray() (v []AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemReasoningProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool call content part of a prompt. It contains a tool call (usually generated
// by the AI model).
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCall struct {
	// Arguments of the tool call. This is a JSON-serializable object that matches the
	// tool's input schema.
	Input any `json:"input" api:"required"`
	// ID of the tool call. This ID is used to match the tool call with the tool
	// result.
	ToolCallID string `json:"toolCallId" api:"required"`
	// Name of the tool that is being called.
	ToolName string            `json:"toolName" api:"required"`
	Type     constant.ToolCall `json:"type" default:"tool-call"`
	// Whether the tool call was executed by the provider.
	ProviderExecuted bool `json:"providerExecuted"`
	// Additional provider-specific metadata. They are passed through to the provider
	// from the AI SDK and enable provider-specific functionality that can be fully
	// encapsulated in the provider.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Input            respjson.Field
		ToolCallID       respjson.Field
		ToolName         respjson.Field
		Type             respjson.Field
		ProviderExecuted respjson.Field
		ProviderOptions  respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCall) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                       struct {
		OfString                                                                                   respjson.Field
		OfFloat                                                                                    respjson.Field
		OfBool                                                                                     respjson.Field
		OfAnyArray                                                                                 respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionArray respjson.Field
		raw                                                                                        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionArray() (v []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolCallProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool result content part of a prompt. It contains the result of the tool call
// with the matching ID.
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResult struct {
	// Result of the tool call. This is a JSON-serializable object.
	Output AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputUnion `json:"output" api:"required"`
	// ID of the tool call that this result is associated with.
	ToolCallID string `json:"toolCallId" api:"required"`
	// Name of the tool that generated this result.
	ToolName string              `json:"toolName" api:"required"`
	Type     constant.ToolResult `json:"type" default:"tool-result"`
	// Additional provider-specific metadata. They are passed through to the provider
	// from the AI SDK and enable provider-specific functionality that can be fully
	// encapsulated in the provider.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Output          respjson.Field
		ToolCallID      respjson.Field
		ToolName        respjson.Field
		Type            respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResult) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputUnion
// contains all possible properties and values from
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputText],
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJson],
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDenied],
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorText],
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJson],
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContent].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputUnion struct {
	Type string `json:"type"`
	// This field is a union of [string],
	// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueUnion],
	// [string],
	// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueUnion],
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueUnion]
	Value AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputUnionValue `json:"value"`
	// This field is a union of
	// [map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionUnion]
	ProviderOptions AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputUnionProviderOptions `json:"providerOptions"`
	// This field is from variant
	// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDenied].
	Reason string `json:"reason"`
	JSON   struct {
		Type            respjson.Field
		Value           respjson.Field
		ProviderOptions respjson.Field
		Reason          respjson.Field
		raw             string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputUnion) AsText() (v AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputUnion) AsJson() (v AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJson) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputUnion) AsExecutionDenied() (v AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDenied) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputUnion) AsErrorText() (v AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputUnion) AsErrorJson() (v AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJson) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputUnion) AsContent() (v AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputUnionValue
// is an implicit subunion of
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputUnion].
// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputUnionValue
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputUnionValue struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueUnion `json:",inline"`
	JSON                                                                                             struct {
		OfString                                                                                                  respjson.Field
		OfFloat                                                                                                   respjson.Field
		OfBool                                                                                                    respjson.Field
		OfAnyArray                                                                                                respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemArray      respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueArray             respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemArray respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueArray        respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueArray          respjson.Field
		raw                                                                                                       string
	} `json:"-"`
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputUnionValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputUnionProviderOptions
// is an implicit subunion of
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputUnion].
// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputUnionProviderOptions
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputUnionProviderOptions struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                                        struct {
		OfString                                                                                                          respjson.Field
		OfFloat                                                                                                           respjson.Field
		OfBool                                                                                                            respjson.Field
		OfAnyArray                                                                                                        respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionArray            respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionArray            respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionArray respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionArray       respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionArray       respjson.Field
		raw                                                                                                               string
	} `json:"-"`
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputUnionProviderOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputText struct {
	// Text tool output that should be directly sent to the API.
	Type  constant.Text `json:"type" default:"text"`
	Value string        `json:"value" api:"required"`
	// Provider-specific options.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type            respjson.Field
		Value           respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputText) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                                   struct {
		OfString                                                                                               respjson.Field
		OfFloat                                                                                                respjson.Field
		OfBool                                                                                                 respjson.Field
		OfAnyArray                                                                                             respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionArray respjson.Field
		raw                                                                                                    string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionArray() (v []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputTextProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJson struct {
	Type constant.Json `json:"type" default:"json"`
	// A JSON value can be a string, number, boolean, object, array, or null. JSON
	// values can be serialized and deserialized by the JSON.stringify and JSON.parse
	// methods.
	Value AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueUnion `json:"value" api:"required"`
	// Provider-specific options.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type            respjson.Field
		Value           respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJson) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemUnion],
// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueArrayItemUnion `json:",inline"`
	JSON                                                                                          struct {
		OfString                                                                                             respjson.Field
		OfFloat                                                                                              respjson.Field
		OfBool                                                                                               respjson.Field
		OfAnyArray                                                                                           respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemArray respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueArray        respjson.Field
		raw                                                                                                  string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapMap() (v map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueArray() (v []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemArrayItemUnion `json:",inline"`
	JSON                                                                                                 struct {
		OfString                                                                                             respjson.Field
		OfFloat                                                                                              respjson.Field
		OfBool                                                                                               respjson.Field
		OfAnyArray                                                                                           respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemArray respjson.Field
		raw                                                                                                  string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemArray() (v []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                                   struct {
		OfString                                                                                               respjson.Field
		OfFloat                                                                                                respjson.Field
		OfBool                                                                                                 respjson.Field
		OfAnyArray                                                                                             respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionArray respjson.Field
		raw                                                                                                    string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionArray() (v []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputJsonProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDenied struct {
	// Type when the user has denied the execution of the tool call.
	Type constant.ExecutionDenied `json:"type" default:"execution-denied"`
	// Provider-specific options.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionUnion `json:"providerOptions"`
	// Optional reason for the execution denial.
	Reason string `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type            respjson.Field
		ProviderOptions respjson.Field
		Reason          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDenied) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDenied) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                                              struct {
		OfString                                                                                                          respjson.Field
		OfFloat                                                                                                           respjson.Field
		OfBool                                                                                                            respjson.Field
		OfAnyArray                                                                                                        respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionArray respjson.Field
		raw                                                                                                               string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionArray() (v []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputExecutionDeniedProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorText struct {
	Type  constant.ErrorText `json:"type" default:"error-text"`
	Value string             `json:"value" api:"required"`
	// Provider-specific options.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type            respjson.Field
		Value           respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorText) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                                        struct {
		OfString                                                                                                    respjson.Field
		OfFloat                                                                                                     respjson.Field
		OfBool                                                                                                      respjson.Field
		OfAnyArray                                                                                                  respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionArray respjson.Field
		raw                                                                                                         string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionArray() (v []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorTextProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJson struct {
	Type constant.ErrorJson `json:"type" default:"error-json"`
	// A JSON value can be a string, number, boolean, object, array, or null. JSON
	// values can be serialized and deserialized by the JSON.stringify and JSON.parse
	// methods.
	Value AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueUnion `json:"value" api:"required"`
	// Provider-specific options.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type            respjson.Field
		Value           respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJson) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemUnion],
// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueArrayItemUnion `json:",inline"`
	JSON                                                                                               struct {
		OfString                                                                                                  respjson.Field
		OfFloat                                                                                                   respjson.Field
		OfBool                                                                                                    respjson.Field
		OfAnyArray                                                                                                respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemArray respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueArray        respjson.Field
		raw                                                                                                       string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapMap() (v map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueArray() (v []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemArrayItemUnion `json:",inline"`
	JSON                                                                                                      struct {
		OfString                                                                                                  respjson.Field
		OfFloat                                                                                                   respjson.Field
		OfBool                                                                                                    respjson.Field
		OfAnyArray                                                                                                respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemArray respjson.Field
		raw                                                                                                       string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemArray() (v []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                                        struct {
		OfString                                                                                                    respjson.Field
		OfFloat                                                                                                     respjson.Field
		OfBool                                                                                                      respjson.Field
		OfAnyArray                                                                                                  respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionArray respjson.Field
		raw                                                                                                         string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionArray() (v []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputErrorJsonProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContent struct {
	Type  constant.Content                                                                                 `json:"type" default:"content"`
	Value []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContent) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueUnion
// contains all possible properties and values from
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueText],
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueMedia],
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileData],
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURL],
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileID],
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageData],
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURL],
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileID],
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustom].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueUnion struct {
	// This field is from variant
	// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueText].
	Text string `json:"text"`
	Type string `json:"type"`
	// This field is a union of
	// [map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionUnion]
	ProviderOptions AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueUnionProviderOptions `json:"providerOptions"`
	Data            string                                                                                                        `json:"data"`
	MediaType       string                                                                                                        `json:"mediaType"`
	// This field is from variant
	// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileData].
	Filename string `json:"filename"`
	URL      string `json:"url"`
	FileID   string `json:"fileId"`
	JSON     struct {
		Text            respjson.Field
		Type            respjson.Field
		ProviderOptions respjson.Field
		Data            respjson.Field
		MediaType       respjson.Field
		Filename        respjson.Field
		URL             respjson.Field
		FileID          respjson.Field
		raw             string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueUnion) AsText() (v AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueUnion) AsMedia() (v AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueMedia) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueUnion) AsFileData() (v AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueUnion) AsFileURL() (v AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueUnion) AsFileID() (v AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileID) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueUnion) AsImageData() (v AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueUnion) AsImageURL() (v AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueUnion) AsImageFileID() (v AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileID) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueUnion) AsCustom() (v AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustom) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueUnionProviderOptions
// is an implicit subunion of
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueUnion].
// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueUnionProviderOptions
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueUnionProviderOptions struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                                                 struct {
		OfString                                                                                                                  respjson.Field
		OfFloat                                                                                                                   respjson.Field
		OfBool                                                                                                                    respjson.Field
		OfAnyArray                                                                                                                respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionArray        respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionArray    respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionArray     respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionArray      respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionArray   respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionArray    respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionArray respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionArray      respjson.Field
		raw                                                                                                                       string
	} `json:"-"`
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueUnionProviderOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueText struct {
	// Text content.
	Text string        `json:"text" api:"required"`
	Type constant.Text `json:"type" default:"text"`
	// Provider-specific options.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text            respjson.Field
		Type            respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueText) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                                               struct {
		OfString                                                                                                           respjson.Field
		OfFloat                                                                                                            respjson.Field
		OfBool                                                                                                             respjson.Field
		OfAnyArray                                                                                                         respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionArray respjson.Field
		raw                                                                                                                string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionArray() (v []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueTextProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueMedia struct {
	Data      string `json:"data" api:"required"`
	MediaType string `json:"mediaType" api:"required"`
	// Deprecated. Use image-data or file-data instead.
	//
	// Deprecated: Deprecated by the upstream schema.
	Type constant.Media `json:"type" default:"media"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		MediaType   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueMedia) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileData struct {
	// Base-64 encoded media data.
	Data string `json:"data" api:"required"`
	// IANA media type.
	MediaType string            `json:"mediaType" api:"required"`
	Type      constant.FileData `json:"type" default:"file-data"`
	// Optional filename of the file.
	Filename string `json:"filename"`
	// Provider-specific options.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data            respjson.Field
		MediaType       respjson.Field
		Type            respjson.Field
		Filename        respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileData) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                                                   struct {
		OfString                                                                                                               respjson.Field
		OfFloat                                                                                                                respjson.Field
		OfBool                                                                                                                 respjson.Field
		OfAnyArray                                                                                                             respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionArray respjson.Field
		raw                                                                                                                    string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionArray() (v []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileDataProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURL struct {
	Type constant.FileURL `json:"type" default:"file-url"`
	// URL of the file.
	URL string `json:"url" api:"required"`
	// Provider-specific options.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type            respjson.Field
		URL             respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURL) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                                                  struct {
		OfString                                                                                                              respjson.Field
		OfFloat                                                                                                               respjson.Field
		OfBool                                                                                                                respjson.Field
		OfAnyArray                                                                                                            respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionArray respjson.Field
		raw                                                                                                                   string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionArray() (v []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileURLProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileID struct {
	// ID of the file.
	//
	// If you use multiple providers, you need to specify the provider specific ids
	// using the Record option. The key is the provider name, e.g. 'openai' or
	// 'anthropic'.
	FileID AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDFileIDUnion `json:"fileId" api:"required"`
	Type   constant.FileID                                                                                            `json:"type" default:"file-id"`
	// Provider-specific options.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID          respjson.Field
		Type            respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileID) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                                                 struct {
		OfString                                                                                                             respjson.Field
		OfFloat                                                                                                              respjson.Field
		OfBool                                                                                                               respjson.Field
		OfAnyArray                                                                                                           respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionArray respjson.Field
		raw                                                                                                                  string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionArray() (v []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueFileIDProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageData struct {
	// Base-64 encoded image data.
	Data string `json:"data" api:"required"`
	// IANA media type.
	MediaType string `json:"mediaType" api:"required"`
	// Images that are referenced using base64 encoded data.
	Type constant.ImageData `json:"type" default:"image-data"`
	// Provider-specific options.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data            respjson.Field
		MediaType       respjson.Field
		Type            respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageData) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                                                    struct {
		OfString                                                                                                                respjson.Field
		OfFloat                                                                                                                 respjson.Field
		OfBool                                                                                                                  respjson.Field
		OfAnyArray                                                                                                              respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionArray respjson.Field
		raw                                                                                                                     string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionArray() (v []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageDataProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURL struct {
	// Images that are referenced using a URL.
	Type constant.ImageURL `json:"type" default:"image-url"`
	// URL of the image.
	URL string `json:"url" api:"required"`
	// Provider-specific options.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type            respjson.Field
		URL             respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                                                   struct {
		OfString                                                                                                               respjson.Field
		OfFloat                                                                                                                respjson.Field
		OfBool                                                                                                                 respjson.Field
		OfAnyArray                                                                                                             respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionArray respjson.Field
		raw                                                                                                                    string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionArray() (v []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageURLProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileID struct {
	// Image that is referenced using a provider file id.
	//
	// If you use multiple providers, you need to specify the provider specific ids
	// using the Record option. The key is the provider name, e.g. 'openai' or
	// 'anthropic'.
	FileID AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDFileIDUnion `json:"fileId" api:"required"`
	// Images that are referenced using a provider file id.
	Type constant.ImageFileID `json:"type" default:"image-file-id"`
	// Provider-specific options.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID          respjson.Field
		Type            respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileID) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                                                      struct {
		OfString                                                                                                                  respjson.Field
		OfFloat                                                                                                                   respjson.Field
		OfBool                                                                                                                    respjson.Field
		OfAnyArray                                                                                                                respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionArray respjson.Field
		raw                                                                                                                       string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionArray() (v []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustom struct {
	// Custom content part. This can be used to implement provider-specific content
	// parts.
	Type constant.Custom `json:"type" default:"custom"`
	// Provider-specific options.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type            respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustom) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                                                 struct {
		OfString                                                                                                             respjson.Field
		OfFloat                                                                                                              respjson.Field
		OfBool                                                                                                               respjson.Field
		OfAnyArray                                                                                                           respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionArray respjson.Field
		raw                                                                                                                  string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionArray() (v []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultOutputContentValueCustomProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                         struct {
		OfString                                                                                     respjson.Field
		OfFloat                                                                                      respjson.Field
		OfBool                                                                                       respjson.Field
		OfAnyArray                                                                                   respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionArray respjson.Field
		raw                                                                                          string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionArray() (v []AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolResultProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool approval request prompt part.
type AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolApprovalRequest struct {
	// ID of the tool approval.
	ApprovalID string `json:"approvalId" api:"required"`
	// ID of the tool call that the approval request is for.
	ToolCallID string                       `json:"toolCallId" api:"required"`
	Type       constant.ToolApprovalRequest `json:"type" default:"tool-approval-request"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApprovalID  respjson.Field
		ToolCallID  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolApprovalRequest) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageAssistantContentArrayItemToolApprovalRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantProviderOptionUnion contains all
// possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageAssistantProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageAssistantProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageAssistantProviderOptionArray]
type AutomateEventAIGenerationDataMessageAssistantProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageAssistantProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageAssistantProviderOptionArray []AutomateEventAIGenerationDataMessageAssistantProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                               struct {
		OfString                                                           respjson.Field
		OfFloat                                                            respjson.Field
		OfBool                                                             respjson.Field
		OfAnyArray                                                         respjson.Field
		OfAutomateEventAIGenerationDataMessageAssistantProviderOptionArray respjson.Field
		raw                                                                string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageAssistantProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantProviderOptionUnion) AsAutomateEventAIGenerationDataMessageAssistantProviderOptionArray() (v []AutomateEventAIGenerationDataMessageAssistantProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageAssistantProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageAssistantProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageAssistantProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageAssistantProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageAssistantProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageAssistantProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A tool message. It contains the result of one or more tool calls.
type AutomateEventAIGenerationDataMessageTool struct {
	// Content of a tool message. It is an array of tool result parts.
	Content []AutomateEventAIGenerationDataMessageToolContentUnion `json:"content" api:"required"`
	Role    constant.Tool                                          `json:"role" default:"tool"`
	// Additional provider-specific metadata. They are passed through to the provider
	// from the AI SDK and enable provider-specific functionality that can be fully
	// encapsulated in the provider.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageToolProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content         respjson.Field
		Role            respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageTool) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventAIGenerationDataMessageTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentUnion contains all possible
// properties and values from
// [AutomateEventAIGenerationDataMessageToolContentToolResult],
// [AutomateEventAIGenerationDataMessageToolContentToolApprovalResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomateEventAIGenerationDataMessageToolContentUnion struct {
	// This field is from variant
	// [AutomateEventAIGenerationDataMessageToolContentToolResult].
	Output AutomateEventAIGenerationDataMessageToolContentToolResultOutputUnion `json:"output"`
	// This field is from variant
	// [AutomateEventAIGenerationDataMessageToolContentToolResult].
	ToolCallID string `json:"toolCallId"`
	// This field is from variant
	// [AutomateEventAIGenerationDataMessageToolContentToolResult].
	ToolName string `json:"toolName"`
	Type     string `json:"type"`
	// This field is from variant
	// [AutomateEventAIGenerationDataMessageToolContentToolResult].
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionUnion `json:"providerOptions"`
	// This field is from variant
	// [AutomateEventAIGenerationDataMessageToolContentToolApprovalResponse].
	ApprovalID string `json:"approvalId"`
	// This field is from variant
	// [AutomateEventAIGenerationDataMessageToolContentToolApprovalResponse].
	Approved bool `json:"approved"`
	// This field is from variant
	// [AutomateEventAIGenerationDataMessageToolContentToolApprovalResponse].
	ProviderExecuted bool `json:"providerExecuted"`
	// This field is from variant
	// [AutomateEventAIGenerationDataMessageToolContentToolApprovalResponse].
	Reason string `json:"reason"`
	JSON   struct {
		Output           respjson.Field
		ToolCallID       respjson.Field
		ToolName         respjson.Field
		Type             respjson.Field
		ProviderOptions  respjson.Field
		ApprovalID       respjson.Field
		Approved         respjson.Field
		ProviderExecuted respjson.Field
		Reason           respjson.Field
		raw              string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentUnion) AsToolResult() (v AutomateEventAIGenerationDataMessageToolContentToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentUnion) AsToolApprovalResponse() (v AutomateEventAIGenerationDataMessageToolContentToolApprovalResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentUnion) RawJSON() string { return u.JSON.raw }

func (r *AutomateEventAIGenerationDataMessageToolContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool result content part of a prompt. It contains the result of the tool call
// with the matching ID.
type AutomateEventAIGenerationDataMessageToolContentToolResult struct {
	// Result of the tool call. This is a JSON-serializable object.
	Output AutomateEventAIGenerationDataMessageToolContentToolResultOutputUnion `json:"output" api:"required"`
	// ID of the tool call that this result is associated with.
	ToolCallID string `json:"toolCallId" api:"required"`
	// Name of the tool that generated this result.
	ToolName string              `json:"toolName" api:"required"`
	Type     constant.ToolResult `json:"type" default:"tool-result"`
	// Additional provider-specific metadata. They are passed through to the provider
	// from the AI SDK and enable provider-specific functionality that can be fully
	// encapsulated in the provider.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Output          respjson.Field
		ToolCallID      respjson.Field
		ToolName        respjson.Field
		Type            respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageToolContentToolResult) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageToolContentToolResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputUnion contains
// all possible properties and values from
// [AutomateEventAIGenerationDataMessageToolContentToolResultOutputText],
// [AutomateEventAIGenerationDataMessageToolContentToolResultOutputJson],
// [AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDenied],
// [AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorText],
// [AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJson],
// [AutomateEventAIGenerationDataMessageToolContentToolResultOutputContent].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputUnion struct {
	Type string `json:"type"`
	// This field is a union of [string],
	// [AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueUnion],
	// [string],
	// [AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueUnion],
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueUnion]
	Value AutomateEventAIGenerationDataMessageToolContentToolResultOutputUnionValue `json:"value"`
	// This field is a union of
	// [map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionUnion]
	ProviderOptions AutomateEventAIGenerationDataMessageToolContentToolResultOutputUnionProviderOptions `json:"providerOptions"`
	// This field is from variant
	// [AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDenied].
	Reason string `json:"reason"`
	JSON   struct {
		Type            respjson.Field
		Value           respjson.Field
		ProviderOptions respjson.Field
		Reason          respjson.Field
		raw             string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputUnion) AsText() (v AutomateEventAIGenerationDataMessageToolContentToolResultOutputText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputUnion) AsJson() (v AutomateEventAIGenerationDataMessageToolContentToolResultOutputJson) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputUnion) AsExecutionDenied() (v AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDenied) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputUnion) AsErrorText() (v AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputUnion) AsErrorJson() (v AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJson) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputUnion) AsContent() (v AutomateEventAIGenerationDataMessageToolContentToolResultOutputContent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputUnionValue is an
// implicit subunion of
// [AutomateEventAIGenerationDataMessageToolContentToolResultOutputUnion].
// AutomateEventAIGenerationDataMessageToolContentToolResultOutputUnionValue
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AutomateEventAIGenerationDataMessageToolContentToolResultOutputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputUnionValue struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueUnion `json:",inline"`
	JSON                                                                               struct {
		OfString                                                                                    respjson.Field
		OfFloat                                                                                     respjson.Field
		OfBool                                                                                      respjson.Field
		OfAnyArray                                                                                  respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemArray      respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueArray             respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemArray respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueArray        respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueArray          respjson.Field
		raw                                                                                         string
	} `json:"-"`
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputUnionValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputUnionProviderOptions
// is an implicit subunion of
// [AutomateEventAIGenerationDataMessageToolContentToolResultOutputUnion].
// AutomateEventAIGenerationDataMessageToolContentToolResultOutputUnionProviderOptions
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AutomateEventAIGenerationDataMessageToolContentToolResultOutputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputUnionProviderOptions struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                          struct {
		OfString                                                                                            respjson.Field
		OfFloat                                                                                             respjson.Field
		OfBool                                                                                              respjson.Field
		OfAnyArray                                                                                          respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionArray            respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionArray            respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionArray respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionArray       respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionArray       respjson.Field
		raw                                                                                                 string
	} `json:"-"`
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputUnionProviderOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageToolContentToolResultOutputText struct {
	// Text tool output that should be directly sent to the API.
	Type  constant.Text `json:"type" default:"text"`
	Value string        `json:"value" api:"required"`
	// Provider-specific options.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type            respjson.Field
		Value           respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageToolContentToolResultOutputText) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                     struct {
		OfString                                                                                 respjson.Field
		OfFloat                                                                                  respjson.Field
		OfBool                                                                                   respjson.Field
		OfAnyArray                                                                               respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionArray respjson.Field
		raw                                                                                      string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionArray() (v []AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputTextProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageToolContentToolResultOutputJson struct {
	Type constant.Json `json:"type" default:"json"`
	// A JSON value can be a string, number, boolean, object, array, or null. JSON
	// values can be serialized and deserialized by the JSON.stringify and JSON.parse
	// methods.
	Value AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueUnion `json:"value" api:"required"`
	// Provider-specific options.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type            respjson.Field
		Value           respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageToolContentToolResultOutputJson) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputJson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemUnion],
// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueArrayItemUnion `json:",inline"`
	JSON                                                                            struct {
		OfString                                                                               respjson.Field
		OfFloat                                                                                respjson.Field
		OfBool                                                                                 respjson.Field
		OfAnyArray                                                                             respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemArray respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueArray        respjson.Field
		raw                                                                                    string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapMap() (v map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueArray() (v []AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemArrayItemUnion `json:",inline"`
	JSON                                                                                   struct {
		OfString                                                                               respjson.Field
		OfFloat                                                                                respjson.Field
		OfBool                                                                                 respjson.Field
		OfAnyArray                                                                             respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemArray respjson.Field
		raw                                                                                    string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemArray() (v []AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                     struct {
		OfString                                                                                 respjson.Field
		OfFloat                                                                                  respjson.Field
		OfBool                                                                                   respjson.Field
		OfAnyArray                                                                               respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionArray respjson.Field
		raw                                                                                      string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionArray() (v []AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputJsonProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDenied struct {
	// Type when the user has denied the execution of the tool call.
	Type constant.ExecutionDenied `json:"type" default:"execution-denied"`
	// Provider-specific options.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionUnion `json:"providerOptions"`
	// Optional reason for the execution denial.
	Reason string `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type            respjson.Field
		ProviderOptions respjson.Field
		Reason          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDenied) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDenied) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                                struct {
		OfString                                                                                            respjson.Field
		OfFloat                                                                                             respjson.Field
		OfBool                                                                                              respjson.Field
		OfAnyArray                                                                                          respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionArray respjson.Field
		raw                                                                                                 string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionArray() (v []AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputExecutionDeniedProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorText struct {
	Type  constant.ErrorText `json:"type" default:"error-text"`
	Value string             `json:"value" api:"required"`
	// Provider-specific options.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type            respjson.Field
		Value           respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorText) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                          struct {
		OfString                                                                                      respjson.Field
		OfFloat                                                                                       respjson.Field
		OfBool                                                                                        respjson.Field
		OfAnyArray                                                                                    respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionArray respjson.Field
		raw                                                                                           string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionArray() (v []AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorTextProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJson struct {
	Type constant.ErrorJson `json:"type" default:"error-json"`
	// A JSON value can be a string, number, boolean, object, array, or null. JSON
	// values can be serialized and deserialized by the JSON.stringify and JSON.parse
	// methods.
	Value AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueUnion `json:"value" api:"required"`
	// Provider-specific options.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type            respjson.Field
		Value           respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJson) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemUnion],
// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueArrayItemUnion `json:",inline"`
	JSON                                                                                 struct {
		OfString                                                                                    respjson.Field
		OfFloat                                                                                     respjson.Field
		OfBool                                                                                      respjson.Field
		OfAnyArray                                                                                  respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemArray respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueArray        respjson.Field
		raw                                                                                         string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapMap() (v map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueArray() (v []AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemArrayItemUnion `json:",inline"`
	JSON                                                                                        struct {
		OfString                                                                                    respjson.Field
		OfFloat                                                                                     respjson.Field
		OfBool                                                                                      respjson.Field
		OfAnyArray                                                                                  respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemArray respjson.Field
		raw                                                                                         string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemArray() (v []AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                          struct {
		OfString                                                                                      respjson.Field
		OfFloat                                                                                       respjson.Field
		OfBool                                                                                        respjson.Field
		OfAnyArray                                                                                    respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionArray respjson.Field
		raw                                                                                           string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionArray() (v []AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputErrorJsonProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContent struct {
	Type  constant.Content                                                                   `json:"type" default:"content"`
	Value []AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageToolContentToolResultOutputContent) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueUnion
// contains all possible properties and values from
// [AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueText],
// [AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueMedia],
// [AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileData],
// [AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURL],
// [AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileID],
// [AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageData],
// [AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURL],
// [AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileID],
// [AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustom].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueUnion struct {
	// This field is from variant
	// [AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueText].
	Text string `json:"text"`
	Type string `json:"type"`
	// This field is a union of
	// [map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionUnion],
	// [map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionUnion]
	ProviderOptions AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueUnionProviderOptions `json:"providerOptions"`
	Data            string                                                                                          `json:"data"`
	MediaType       string                                                                                          `json:"mediaType"`
	// This field is from variant
	// [AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileData].
	Filename string `json:"filename"`
	URL      string `json:"url"`
	FileID   string `json:"fileId"`
	JSON     struct {
		Text            respjson.Field
		Type            respjson.Field
		ProviderOptions respjson.Field
		Data            respjson.Field
		MediaType       respjson.Field
		Filename        respjson.Field
		URL             respjson.Field
		FileID          respjson.Field
		raw             string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueUnion) AsText() (v AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueUnion) AsMedia() (v AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueMedia) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueUnion) AsFileData() (v AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueUnion) AsFileURL() (v AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueUnion) AsFileID() (v AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileID) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueUnion) AsImageData() (v AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueUnion) AsImageURL() (v AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueUnion) AsImageFileID() (v AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileID) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueUnion) AsCustom() (v AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustom) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueUnionProviderOptions
// is an implicit subunion of
// [AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueUnion].
// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueUnionProviderOptions
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueUnionProviderOptions struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                                   struct {
		OfString                                                                                                    respjson.Field
		OfFloat                                                                                                     respjson.Field
		OfBool                                                                                                      respjson.Field
		OfAnyArray                                                                                                  respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionArray        respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionArray    respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionArray     respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionArray      respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionArray   respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionArray    respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionArray respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionArray      respjson.Field
		raw                                                                                                         string
	} `json:"-"`
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueUnionProviderOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueText struct {
	// Text content.
	Text string        `json:"text" api:"required"`
	Type constant.Text `json:"type" default:"text"`
	// Provider-specific options.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text            respjson.Field
		Type            respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueText) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                                 struct {
		OfString                                                                                             respjson.Field
		OfFloat                                                                                              respjson.Field
		OfBool                                                                                               respjson.Field
		OfAnyArray                                                                                           respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionArray respjson.Field
		raw                                                                                                  string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionArray() (v []AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueTextProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueMedia struct {
	Data      string `json:"data" api:"required"`
	MediaType string `json:"mediaType" api:"required"`
	// Deprecated. Use image-data or file-data instead.
	//
	// Deprecated: Deprecated by the upstream schema.
	Type constant.Media `json:"type" default:"media"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		MediaType   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueMedia) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileData struct {
	// Base-64 encoded media data.
	Data string `json:"data" api:"required"`
	// IANA media type.
	MediaType string            `json:"mediaType" api:"required"`
	Type      constant.FileData `json:"type" default:"file-data"`
	// Optional filename of the file.
	Filename string `json:"filename"`
	// Provider-specific options.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data            respjson.Field
		MediaType       respjson.Field
		Type            respjson.Field
		Filename        respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileData) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                                     struct {
		OfString                                                                                                 respjson.Field
		OfFloat                                                                                                  respjson.Field
		OfBool                                                                                                   respjson.Field
		OfAnyArray                                                                                               respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionArray respjson.Field
		raw                                                                                                      string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionArray() (v []AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileDataProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURL struct {
	Type constant.FileURL `json:"type" default:"file-url"`
	// URL of the file.
	URL string `json:"url" api:"required"`
	// Provider-specific options.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type            respjson.Field
		URL             respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURL) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                                    struct {
		OfString                                                                                                respjson.Field
		OfFloat                                                                                                 respjson.Field
		OfBool                                                                                                  respjson.Field
		OfAnyArray                                                                                              respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionArray respjson.Field
		raw                                                                                                     string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionArray() (v []AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileURLProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileID struct {
	// ID of the file.
	//
	// If you use multiple providers, you need to specify the provider specific ids
	// using the Record option. The key is the provider name, e.g. 'openai' or
	// 'anthropic'.
	FileID AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDFileIDUnion `json:"fileId" api:"required"`
	Type   constant.FileID                                                                              `json:"type" default:"file-id"`
	// Provider-specific options.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID          respjson.Field
		Type            respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileID) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                                   struct {
		OfString                                                                                               respjson.Field
		OfFloat                                                                                                respjson.Field
		OfBool                                                                                                 respjson.Field
		OfAnyArray                                                                                             respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionArray respjson.Field
		raw                                                                                                    string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionArray() (v []AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueFileIDProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageData struct {
	// Base-64 encoded image data.
	Data string `json:"data" api:"required"`
	// IANA media type.
	MediaType string `json:"mediaType" api:"required"`
	// Images that are referenced using base64 encoded data.
	Type constant.ImageData `json:"type" default:"image-data"`
	// Provider-specific options.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data            respjson.Field
		MediaType       respjson.Field
		Type            respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageData) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                                      struct {
		OfString                                                                                                  respjson.Field
		OfFloat                                                                                                   respjson.Field
		OfBool                                                                                                    respjson.Field
		OfAnyArray                                                                                                respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionArray respjson.Field
		raw                                                                                                       string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionArray() (v []AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageDataProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURL struct {
	// Images that are referenced using a URL.
	Type constant.ImageURL `json:"type" default:"image-url"`
	// URL of the image.
	URL string `json:"url" api:"required"`
	// Provider-specific options.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type            respjson.Field
		URL             respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURL) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                                     struct {
		OfString                                                                                                 respjson.Field
		OfFloat                                                                                                  respjson.Field
		OfBool                                                                                                   respjson.Field
		OfAnyArray                                                                                               respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionArray respjson.Field
		raw                                                                                                      string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionArray() (v []AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageURLProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileID struct {
	// Image that is referenced using a provider file id.
	//
	// If you use multiple providers, you need to specify the provider specific ids
	// using the Record option. The key is the provider name, e.g. 'openai' or
	// 'anthropic'.
	FileID AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDFileIDUnion `json:"fileId" api:"required"`
	// Images that are referenced using a provider file id.
	Type constant.ImageFileID `json:"type" default:"image-file-id"`
	// Provider-specific options.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID          respjson.Field
		Type            respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileID) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                                        struct {
		OfString                                                                                                    respjson.Field
		OfFloat                                                                                                     respjson.Field
		OfBool                                                                                                      respjson.Field
		OfAnyArray                                                                                                  respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionArray respjson.Field
		raw                                                                                                         string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionArray() (v []AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueImageFileIDProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustom struct {
	// Custom content part. This can be used to implement provider-specific content
	// parts.
	Type constant.Custom `json:"type" default:"custom"`
	// Provider-specific options.
	ProviderOptions map[string]map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionUnion `json:"providerOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type            respjson.Field
		ProviderOptions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustom) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                                                   struct {
		OfString                                                                                               respjson.Field
		OfFloat                                                                                                respjson.Field
		OfBool                                                                                                 respjson.Field
		OfAnyArray                                                                                             respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionArray respjson.Field
		raw                                                                                                    string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionArray() (v []AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultOutputContentValueCustomProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionUnion
// contains all possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionArrayItemUnion]
	// instead of an object.
	OfAutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionArray []AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                                           struct {
		OfString                                                                       respjson.Field
		OfFloat                                                                        respjson.Field
		OfBool                                                                         respjson.Field
		OfAnyArray                                                                     respjson.Field
		OfAutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionArray respjson.Field
		raw                                                                            string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionArray() (v []AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolContentToolResultProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool approval response prompt part.
type AutomateEventAIGenerationDataMessageToolContentToolApprovalResponse struct {
	// ID of the tool approval.
	ApprovalID string `json:"approvalId" api:"required"`
	// Flag indicating whether the approval was granted or denied.
	Approved bool                          `json:"approved" api:"required"`
	Type     constant.ToolApprovalResponse `json:"type" default:"tool-approval-response"`
	// Flag indicating whether the tool call is provider-executed. Only
	// provider-executed tool approval responses should be sent to the model.
	ProviderExecuted bool `json:"providerExecuted"`
	// Optional reason for the approval or denial.
	Reason string `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApprovalID       respjson.Field
		Approved         respjson.Field
		Type             respjson.Field
		ProviderExecuted respjson.Field
		Reason           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationDataMessageToolContentToolApprovalResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *AutomateEventAIGenerationDataMessageToolContentToolApprovalResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolProviderOptionUnion contains all
// possible properties and values from [string], [float64], [bool],
// [map[string]AutomateEventAIGenerationDataMessageToolProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion],
// [[]AutomateEventAIGenerationDataMessageToolProviderOptionArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfAutomateEventAIGenerationDataMessageToolProviderOptionArray]
type AutomateEventAIGenerationDataMessageToolProviderOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a
	// [[]AutomateEventAIGenerationDataMessageToolProviderOptionArrayItemUnion] instead
	// of an object.
	OfAutomateEventAIGenerationDataMessageToolProviderOptionArray []AutomateEventAIGenerationDataMessageToolProviderOptionArrayItemUnion `json:",inline"`
	JSON                                                          struct {
		OfString                                                      respjson.Field
		OfFloat                                                       respjson.Field
		OfBool                                                        respjson.Field
		OfAnyArray                                                    respjson.Field
		OfAutomateEventAIGenerationDataMessageToolProviderOptionArray respjson.Field
		raw                                                           string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolProviderOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolProviderOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolProviderOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolProviderOptionV1Alias922155206_411_476_922155206_0_129727Map() (v map[string]AutomateEventAIGenerationDataMessageToolProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolProviderOptionUnion) AsAutomateEventAIGenerationDataMessageToolProviderOptionArray() (v []AutomateEventAIGenerationDataMessageToolProviderOptionArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolProviderOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolProviderOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolProviderOptionV1Alias922155206_411_476_922155206_0_129727ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutomateEventAIGenerationDataMessageToolProviderOptionArrayItemUnion contains
// all possible properties and values from [string], [float64], [bool], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type AutomateEventAIGenerationDataMessageToolProviderOptionArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u AutomateEventAIGenerationDataMessageToolProviderOptionArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolProviderOptionArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolProviderOptionArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutomateEventAIGenerationDataMessageToolProviderOptionArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutomateEventAIGenerationDataMessageToolProviderOptionArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutomateEventAIGenerationDataMessageToolProviderOptionArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "ai:generation:error" event from /v1/automate.
type AutomateEventAIGenerationError struct {
	// Event data when AI generation error occurs
	Data  AutomateEventAIGenerationErrorData `json:"data" api:"required"`
	Event constant.AIGenerationError         `json:"event" default:"ai:generation:error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationError) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventAIGenerationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data when AI generation error occurs
type AutomateEventAIGenerationErrorData struct {
	Error       string  `json:"error" api:"required"`
	IterationID string  `json:"iterationId" api:"required"`
	Prompt      string  `json:"prompt" api:"required"`
	Schema      any     `json:"schema" api:"required"`
	Timestamp   float64 `json:"timestamp" api:"required"`
	Messages    []any   `json:"messages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		IterationID respjson.Field
		Prompt      respjson.Field
		Schema      respjson.Field
		Timestamp   respjson.Field
		Messages    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventAIGenerationErrorData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventAIGenerationErrorData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "browser:action_completed" event from /v1/automate.
type AutomateEventBrowserActionCompleted struct {
	// Event data for action results
	Data  AutomateEventBrowserActionCompletedData `json:"data" api:"required"`
	Event constant.BrowserActionCompleted         `json:"event" default:"browser:action_completed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventBrowserActionCompleted) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventBrowserActionCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data for action results
type AutomateEventBrowserActionCompletedData struct {
	IterationID string  `json:"iterationId" api:"required"`
	Success     bool    `json:"success" api:"required"`
	Timestamp   float64 `json:"timestamp" api:"required"`
	Error       string  `json:"error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IterationID respjson.Field
		Success     respjson.Field
		Timestamp   respjson.Field
		Error       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventBrowserActionCompletedData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventBrowserActionCompletedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "browser:action_started" event from /v1/automate.
type AutomateEventBrowserActionStarted struct {
	// Event data for action execution
	Data  AutomateEventBrowserActionStartedData `json:"data" api:"required"`
	Event constant.BrowserActionStarted         `json:"event" default:"browser:action_started"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventBrowserActionStarted) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventBrowserActionStarted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data for action execution
type AutomateEventBrowserActionStartedData struct {
	Action      string  `json:"action" api:"required"`
	IterationID string  `json:"iterationId" api:"required"`
	Timestamp   float64 `json:"timestamp" api:"required"`
	Ref         string  `json:"ref" api:"nullable"`
	Value       string  `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		IterationID respjson.Field
		Timestamp   respjson.Field
		Ref         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventBrowserActionStartedData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventBrowserActionStartedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "browser:navigated" event from /v1/automate.
type AutomateEventBrowserNavigated struct {
	// Event data when navigating to a page
	Data  AutomateEventBrowserNavigatedData `json:"data" api:"required"`
	Event constant.BrowserNavigated         `json:"event" default:"browser:navigated"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventBrowserNavigated) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventBrowserNavigated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data when navigating to a page
type AutomateEventBrowserNavigatedData struct {
	IterationID string  `json:"iterationId" api:"required"`
	Timestamp   float64 `json:"timestamp" api:"required"`
	Title       string  `json:"title" api:"required"`
	URL         string  `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IterationID respjson.Field
		Timestamp   respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventBrowserNavigatedData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventBrowserNavigatedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "browser:reconnected" event from /v1/automate.
type AutomateEventBrowserReconnected struct {
	// Event data when the browser reconnects after a mid-task disconnect
	Data  AutomateEventBrowserReconnectedData `json:"data" api:"required"`
	Event constant.BrowserReconnected         `json:"event" default:"browser:reconnected"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventBrowserReconnected) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventBrowserReconnected) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data when the browser reconnects after a mid-task disconnect
type AutomateEventBrowserReconnectedData struct {
	// 1-based index of the CDP endpoint now in use
	EndpointIndex float64 `json:"endpointIndex" api:"required"`
	IterationID   string  `json:"iterationId" api:"required"`
	// The original starting URL the agent is restarting execution from
	StartingURL string  `json:"startingUrl" api:"required"`
	Timestamp   float64 `json:"timestamp" api:"required"`
	// Total number of configured CDP endpoints
	Total float64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndpointIndex respjson.Field
		IterationID   respjson.Field
		StartingURL   respjson.Field
		Timestamp     respjson.Field
		Total         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventBrowserReconnectedData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventBrowserReconnectedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "browser:screenshot_captured" event from /v1/automate.
type AutomateEventBrowserScreenshotCaptured struct {
	// Event data for screenshot capture
	Data  AutomateEventBrowserScreenshotCapturedData `json:"data" api:"required"`
	Event constant.BrowserScreenshotCaptured         `json:"event" default:"browser:screenshot_captured"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventBrowserScreenshotCaptured) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventBrowserScreenshotCaptured) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data for screenshot capture
type AutomateEventBrowserScreenshotCapturedData struct {
	// Any of "jpeg", "png".
	Format      string  `json:"format" api:"required"`
	IterationID string  `json:"iterationId" api:"required"`
	Size        float64 `json:"size" api:"required"`
	Timestamp   float64 `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Format      respjson.Field
		IterationID respjson.Field
		Size        respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventBrowserScreenshotCapturedData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventBrowserScreenshotCapturedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "browser:screenshot_captured_image" event from /v1/automate.
type AutomateEventBrowserScreenshotCapturedImage struct {
	// Event data for screenshot image capture with full image data This event contains
	// the complete screenshot and can be very large
	Data  AutomateEventBrowserScreenshotCapturedImageData `json:"data" api:"required"`
	Event constant.BrowserScreenshotCapturedImage         `json:"event" default:"browser:screenshot_captured_image"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventBrowserScreenshotCapturedImage) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventBrowserScreenshotCapturedImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data for screenshot image capture with full image data This event contains
// the complete screenshot and can be very large
type AutomateEventBrowserScreenshotCapturedImageData struct {
	Image       string `json:"image" api:"required"`
	IterationID string `json:"iterationId" api:"required"`
	// Any of "image/jpeg", "image/png".
	MediaType string  `json:"mediaType" api:"required"`
	Timestamp float64 `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image       respjson.Field
		IterationID respjson.Field
		MediaType   respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventBrowserScreenshotCapturedImageData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventBrowserScreenshotCapturedImageData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "cdp:endpoint_connected" event from /v1/automate.
type AutomateEventCdpEndpointConnected struct {
	// Event data when a CDP endpoint is successfully connected to
	Data  AutomateEventCdpEndpointConnectedData `json:"data" api:"required"`
	Event constant.CdpEndpointConnected         `json:"event" default:"cdp:endpoint_connected"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventCdpEndpointConnected) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventCdpEndpointConnected) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data when a CDP endpoint is successfully connected to
type AutomateEventCdpEndpointConnectedData struct {
	// 1-based index of the endpoint that connected
	EndpointIndex float64 `json:"endpointIndex" api:"required"`
	IterationID   string  `json:"iterationId" api:"required"`
	Timestamp     float64 `json:"timestamp" api:"required"`
	// Total number of configured CDP endpoints
	Total float64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndpointIndex respjson.Field
		IterationID   respjson.Field
		Timestamp     respjson.Field
		Total         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventCdpEndpointConnectedData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventCdpEndpointConnectedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "cdp:endpoint_cycle" event from /v1/automate.
type AutomateEventCdpEndpointCycle struct {
	// Event data when a CDP endpoint fails and the next one is being tried
	Data  AutomateEventCdpEndpointCycleData `json:"data" api:"required"`
	Event constant.CdpEndpointCycle         `json:"event" default:"cdp:endpoint_cycle"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventCdpEndpointCycle) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventCdpEndpointCycle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data when a CDP endpoint fails and the next one is being tried
type AutomateEventCdpEndpointCycleData struct {
	// 1-based index of the endpoint attempt that failed
	Attempt float64 `json:"attempt" api:"required"`
	// Sanitized error identifier from the failed connection attempt (error.name, not
	// error.message — full messages may contain endpoint URLs)
	Error       string  `json:"error" api:"required"`
	IterationID string  `json:"iterationId" api:"required"`
	Timestamp   float64 `json:"timestamp" api:"required"`
	// Total number of configured CDP endpoints
	Total float64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attempt     respjson.Field
		Error       respjson.Field
		IterationID respjson.Field
		Timestamp   respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventCdpEndpointCycleData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventCdpEndpointCycleData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "complete" event from /v1/automate.
type AutomateEventComplete struct {
	// Payload for the `complete` stream event. Structurally identical to
	// TaskExecutionResult from webAgent.ts — the `complete` event's data is the
	// agent's final TaskExecutionResult, stringified onto the SSE stream.
	Data  AutomateEventCompleteData `json:"data" api:"required"`
	Event constant.Complete         `json:"event" default:"complete"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventComplete) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventComplete) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the `complete` stream event. Structurally identical to
// TaskExecutionResult from webAgent.ts — the `complete` event's data is the
// agent's final TaskExecutionResult, stringified onto the SSE stream.
type AutomateEventCompleteData struct {
	// Final answer or result from the agent
	FinalAnswer string `json:"finalAnswer" api:"required"`
	// Execution statistics
	Stats AutomateEventCompleteDataStats `json:"stats" api:"required"`
	// Whether the task completed successfully
	Success bool `json:"success" api:"required"`
	// Structured error information for failed tasks
	Error AutomateEventCompleteDataError `json:"error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FinalAnswer respjson.Field
		Stats       respjson.Field
		Success     respjson.Field
		Error       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventCompleteData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventCompleteData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Execution statistics
type AutomateEventCompleteDataStats struct {
	Actions    float64 `json:"actions" api:"required"`
	DurationMs float64 `json:"durationMs" api:"required"`
	EndTime    float64 `json:"endTime" api:"required"`
	Iterations float64 `json:"iterations" api:"required"`
	StartTime  float64 `json:"startTime" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actions     respjson.Field
		DurationMs  respjson.Field
		EndTime     respjson.Field
		Iterations  respjson.Field
		StartTime   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventCompleteDataStats) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventCompleteDataStats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structured error information for failed tasks
type AutomateEventCompleteDataError struct {
	// Error codes for task failures
	//
	// Any of "TASK_ABORTED", "MAX_ITERATIONS", "MAX_ERRORS", "TASK_FAILED".
	Code string `json:"code" api:"required"`
	// Human-readable error message
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventCompleteDataError) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventCompleteDataError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "done" event from /v1/automate.
type AutomateEventDone struct {
	// Payload for the `done` stream terminator event. Empty today; reserved for future
	// metadata.
	Data  map[string]any `json:"data" api:"required"`
	Event constant.Done  `json:"event" default:"done"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventDone) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventDone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "error" event from /v1/automate.
type AutomateEventError struct {
	// Payload for the top-level `error` stream event. Emitted when an uncaught error
	// escapes the task runner. Mirrors `ErrorResponse` from the server package's
	// `taskRunner.ts` — kept structurally aligned so schema and runtime stay
	// consistent. Distinct from agent-level error events like `ai:generation:error`
	// and `task:validation_error`, which are emitted through the normal event emitter
	// during the agent loop.
	Data  AutomateEventErrorData `json:"data" api:"required"`
	Event constant.Error         `json:"event" default:"error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventError) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for the top-level `error` stream event. Emitted when an uncaught error
// escapes the task runner. Mirrors `ErrorResponse` from the server package's
// `taskRunner.ts` — kept structurally aligned so schema and runtime stay
// consistent. Distinct from agent-level error events like `ai:generation:error`
// and `task:validation_error`, which are emitted through the normal event emitter
// during the agent loop.
type AutomateEventErrorData struct {
	Error   AutomateEventErrorDataError `json:"error" api:"required"`
	Success bool                        `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventErrorData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventErrorData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventErrorDataError struct {
	Code    string `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// ISO-8601 timestamp
	Timestamp string `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventErrorDataError) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventErrorDataError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "interactive:form_data:error" event from /v1/automate.
type AutomateEventInteractiveFormDataError struct {
	// Event data when form validation fails and the agent re-requests data. Carries
	// both the error context and the fields that need new values. Callers respond to
	// this the same way as a request event.
	Data  AutomateEventInteractiveFormDataErrorData `json:"data" api:"required"`
	Event constant.InteractiveFormDataError         `json:"event" default:"interactive:form_data:error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventInteractiveFormDataError) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventInteractiveFormDataError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data when form validation fails and the agent re-requests data. Carries
// both the error context and the fields that need new values. Callers respond to
// this the same way as a request event.
type AutomateEventInteractiveFormDataErrorData struct {
	// Per-field error messages from validation (field ref -> error text)
	FieldErrors     map[string]string                                `json:"fieldErrors" api:"required"`
	Fields          []AutomateEventInteractiveFormDataErrorDataField `json:"fields" api:"required"`
	FormDescription string                                           `json:"formDescription" api:"required"`
	IterationID     string                                           `json:"iterationId" api:"required"`
	PageTitle       string                                           `json:"pageTitle" api:"required"`
	PageURL         string                                           `json:"pageUrl" api:"required"`
	RequestID       string                                           `json:"requestId" api:"required"`
	Timestamp       float64                                          `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldErrors     respjson.Field
		Fields          respjson.Field
		FormDescription respjson.Field
		IterationID     respjson.Field
		PageTitle       respjson.Field
		PageURL         respjson.Field
		RequestID       respjson.Field
		Timestamp       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventInteractiveFormDataErrorData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventInteractiveFormDataErrorData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single form field the agent needs data for.
type AutomateEventInteractiveFormDataErrorDataField struct {
	// Semantic field type
	//
	// Any of "text", "email", "phone", "date", "number", "select", "checkbox",
	// "radio", "textarea", "password", "other".
	FieldType string `json:"fieldType" api:"required"`
	// The field's visible label
	Label string `json:"label" api:"required"`
	// Element ref from the accessibility tree (e.g., "E42")
	Ref string `json:"ref" api:"required"`
	// Whether this field must be filled
	Required bool `json:"required" api:"required"`
	// Current value if already partially filled
	CurrentValue string `json:"currentValue"`
	// Additional context (e.g., validation error message on re-request)
	Description string `json:"description"`
	// Available options for select/radio fields
	Options []string `json:"options"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldType    respjson.Field
		Label        respjson.Field
		Ref          respjson.Field
		Required     respjson.Field
		CurrentValue respjson.Field
		Description  respjson.Field
		Options      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventInteractiveFormDataErrorDataField) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventInteractiveFormDataErrorDataField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "interactive:form_data:request" event from /v1/automate.
type AutomateEventInteractiveFormDataRequest struct {
	// Event data when the agent requests user data for form fields
	Data  AutomateEventInteractiveFormDataRequestData `json:"data" api:"required"`
	Event constant.InteractiveFormDataRequest         `json:"event" default:"interactive:form_data:request"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventInteractiveFormDataRequest) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventInteractiveFormDataRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data when the agent requests user data for form fields
type AutomateEventInteractiveFormDataRequestData struct {
	Fields          []AutomateEventInteractiveFormDataRequestDataField `json:"fields" api:"required"`
	FormDescription string                                             `json:"formDescription" api:"required"`
	IterationID     string                                             `json:"iterationId" api:"required"`
	PageTitle       string                                             `json:"pageTitle" api:"required"`
	PageURL         string                                             `json:"pageUrl" api:"required"`
	RequestID       string                                             `json:"requestId" api:"required"`
	Timestamp       float64                                            `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fields          respjson.Field
		FormDescription respjson.Field
		IterationID     respjson.Field
		PageTitle       respjson.Field
		PageURL         respjson.Field
		RequestID       respjson.Field
		Timestamp       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventInteractiveFormDataRequestData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventInteractiveFormDataRequestData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single form field the agent needs data for.
type AutomateEventInteractiveFormDataRequestDataField struct {
	// Semantic field type
	//
	// Any of "text", "email", "phone", "date", "number", "select", "checkbox",
	// "radio", "textarea", "password", "other".
	FieldType string `json:"fieldType" api:"required"`
	// The field's visible label
	Label string `json:"label" api:"required"`
	// Element ref from the accessibility tree (e.g., "E42")
	Ref string `json:"ref" api:"required"`
	// Whether this field must be filled
	Required bool `json:"required" api:"required"`
	// Current value if already partially filled
	CurrentValue string `json:"currentValue"`
	// Additional context (e.g., validation error message on re-request)
	Description string `json:"description"`
	// Available options for select/radio fields
	Options []string `json:"options"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldType    respjson.Field
		Label        respjson.Field
		Ref          respjson.Field
		Required     respjson.Field
		CurrentValue respjson.Field
		Description  respjson.Field
		Options      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventInteractiveFormDataRequestDataField) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventInteractiveFormDataRequestDataField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "system:debug_compression" event from /v1/automate.
type AutomateEventSystemDebugCompression struct {
	// Event data for compression debug info
	Data  AutomateEventSystemDebugCompressionData `json:"data" api:"required"`
	Event constant.SystemDebugCompression         `json:"event" default:"system:debug_compression"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventSystemDebugCompression) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventSystemDebugCompression) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data for compression debug info
type AutomateEventSystemDebugCompressionData struct {
	CompressedSize     float64 `json:"compressedSize" api:"required"`
	CompressionPercent float64 `json:"compressionPercent" api:"required"`
	IterationID        string  `json:"iterationId" api:"required"`
	OriginalSize       float64 `json:"originalSize" api:"required"`
	Timestamp          float64 `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompressedSize     respjson.Field
		CompressionPercent respjson.Field
		IterationID        respjson.Field
		OriginalSize       respjson.Field
		Timestamp          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventSystemDebugCompressionData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventSystemDebugCompressionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "system:debug_message" event from /v1/automate.
type AutomateEventSystemDebugMessage struct {
	// Event data for message debug info
	Data  AutomateEventSystemDebugMessageData `json:"data" api:"required"`
	Event constant.SystemDebugMessage         `json:"event" default:"system:debug_message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventSystemDebugMessage) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventSystemDebugMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data for message debug info
type AutomateEventSystemDebugMessageData struct {
	IterationID string  `json:"iterationId" api:"required"`
	Messages    []any   `json:"messages" api:"required"`
	Timestamp   float64 `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IterationID respjson.Field
		Messages    respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventSystemDebugMessageData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventSystemDebugMessageData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "task:aborted" event from /v1/automate.
type AutomateEventTaskAborted struct {
	// Event data when a task is aborted
	Data  AutomateEventTaskAbortedData `json:"data" api:"required"`
	Event constant.TaskAborted         `json:"event" default:"task:aborted"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventTaskAborted) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventTaskAborted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data when a task is aborted
type AutomateEventTaskAbortedData struct {
	FinalAnswer string  `json:"finalAnswer" api:"required"`
	IterationID string  `json:"iterationId" api:"required"`
	Reason      string  `json:"reason" api:"required"`
	Timestamp   float64 `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FinalAnswer respjson.Field
		IterationID respjson.Field
		Reason      respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventTaskAbortedData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventTaskAbortedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "task:completed" event from /v1/automate.
type AutomateEventTaskCompleted struct {
	// Event data when a task is completed
	Data  AutomateEventTaskCompletedData `json:"data" api:"required"`
	Event constant.TaskCompleted         `json:"event" default:"task:completed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventTaskCompleted) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventTaskCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data when a task is completed
type AutomateEventTaskCompletedData struct {
	FinalAnswer string  `json:"finalAnswer" api:"required"`
	IterationID string  `json:"iterationId" api:"required"`
	Timestamp   float64 `json:"timestamp" api:"required"`
	Success     bool    `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FinalAnswer respjson.Field
		IterationID respjson.Field
		Timestamp   respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventTaskCompletedData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventTaskCompletedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "task:metrics" event from /v1/automate.
type AutomateEventTaskMetrics struct {
	Data  AutomateEventTaskMetricsData `json:"data" api:"required"`
	Event constant.TaskMetrics         `json:"event" default:"task:metrics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventTaskMetrics) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventTaskMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventTaskMetricsData struct {
	AIGenerationCount      float64            `json:"aiGenerationCount" api:"required"`
	AIGenerationErrorCount float64            `json:"aiGenerationErrorCount" api:"required"`
	EventCounts            map[string]float64 `json:"eventCounts" api:"required"`
	IterationID            string             `json:"iterationId" api:"required"`
	StepCount              float64            `json:"stepCount" api:"required"`
	Timestamp              float64            `json:"timestamp" api:"required"`
	TotalInputTokens       float64            `json:"totalInputTokens" api:"required"`
	TotalOutputTokens      float64            `json:"totalOutputTokens" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AIGenerationCount      respjson.Field
		AIGenerationErrorCount respjson.Field
		EventCounts            respjson.Field
		IterationID            respjson.Field
		StepCount              respjson.Field
		Timestamp              respjson.Field
		TotalInputTokens       respjson.Field
		TotalOutputTokens      respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventTaskMetricsData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventTaskMetricsData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "task:metrics_incremental" event from /v1/automate.
type AutomateEventTaskMetricsIncremental struct {
	Data  AutomateEventTaskMetricsIncrementalData `json:"data" api:"required"`
	Event constant.TaskMetricsIncremental         `json:"event" default:"task:metrics_incremental"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventTaskMetricsIncremental) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventTaskMetricsIncremental) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutomateEventTaskMetricsIncrementalData struct {
	AIGenerationCount      float64            `json:"aiGenerationCount" api:"required"`
	AIGenerationErrorCount float64            `json:"aiGenerationErrorCount" api:"required"`
	EventCounts            map[string]float64 `json:"eventCounts" api:"required"`
	IterationID            string             `json:"iterationId" api:"required"`
	StepCount              float64            `json:"stepCount" api:"required"`
	Timestamp              float64            `json:"timestamp" api:"required"`
	TotalInputTokens       float64            `json:"totalInputTokens" api:"required"`
	TotalOutputTokens      float64            `json:"totalOutputTokens" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AIGenerationCount      respjson.Field
		AIGenerationErrorCount respjson.Field
		EventCounts            respjson.Field
		IterationID            respjson.Field
		StepCount              respjson.Field
		Timestamp              respjson.Field
		TotalInputTokens       respjson.Field
		TotalOutputTokens      respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventTaskMetricsIncrementalData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventTaskMetricsIncrementalData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "task:setup" event from /v1/automate.
type AutomateEventTaskSetup struct {
	// Event data when a task is setup
	Data  AutomateEventTaskSetupData `json:"data" api:"required"`
	Event constant.TaskSetup         `json:"event" default:"task:setup"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventTaskSetup) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventTaskSetup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data when a task is setup
type AutomateEventTaskSetupData struct {
	BrowserName string  `json:"browserName" api:"required"`
	IterationID string  `json:"iterationId" api:"required"`
	Task        string  `json:"task" api:"required"`
	Timestamp   float64 `json:"timestamp" api:"required"`
	Data        any     `json:"data"`
	Guardrails  string  `json:"guardrails"`
	HasAPIKey   bool    `json:"hasApiKey"`
	// Any of "global", "env", "not_set".
	KeySource     string `json:"keySource"`
	Model         string `json:"model"`
	Provider      string `json:"provider"`
	Proxy         string `json:"proxy"`
	PwCdpEndpoint string `json:"pwCdpEndpoint"`
	// Total number of CDP endpoints configured (index, not URLs)
	PwCdpEndpointCount float64  `json:"pwCdpEndpointCount"`
	PwCdpEndpoints     []string `json:"pwCdpEndpoints"`
	PwEndpoint         string   `json:"pwEndpoint"`
	URL                string   `json:"url"`
	Vision             bool     `json:"vision"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrowserName        respjson.Field
		IterationID        respjson.Field
		Task               respjson.Field
		Timestamp          respjson.Field
		Data               respjson.Field
		Guardrails         respjson.Field
		HasAPIKey          respjson.Field
		KeySource          respjson.Field
		Model              respjson.Field
		Provider           respjson.Field
		Proxy              respjson.Field
		PwCdpEndpoint      respjson.Field
		PwCdpEndpointCount respjson.Field
		PwCdpEndpoints     respjson.Field
		PwEndpoint         respjson.Field
		URL                respjson.Field
		Vision             respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventTaskSetupData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventTaskSetupData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "task:started" event from /v1/automate.
type AutomateEventTaskStarted struct {
	// Event data when a task is started
	Data  AutomateEventTaskStartedData `json:"data" api:"required"`
	Event constant.TaskStarted         `json:"event" default:"task:started"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventTaskStarted) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventTaskStarted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data when a task is started
type AutomateEventTaskStartedData struct {
	IterationID     string   `json:"iterationId" api:"required"`
	Plan            string   `json:"plan" api:"required"`
	SuccessCriteria string   `json:"successCriteria" api:"required"`
	Task            string   `json:"task" api:"required"`
	Timestamp       float64  `json:"timestamp" api:"required"`
	URL             string   `json:"url" api:"required"`
	ActionItems     []string `json:"actionItems"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IterationID     respjson.Field
		Plan            respjson.Field
		SuccessCriteria respjson.Field
		Task            respjson.Field
		Timestamp       respjson.Field
		URL             respjson.Field
		ActionItems     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventTaskStartedData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventTaskStartedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "task:validated" event from /v1/automate.
type AutomateEventTaskValidated struct {
	// Event data for task validation
	Data  AutomateEventTaskValidatedData `json:"data" api:"required"`
	Event constant.TaskValidated         `json:"event" default:"task:validated"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventTaskValidated) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventTaskValidated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data for task validation
type AutomateEventTaskValidatedData struct {
	// Any of "failed", "partial", "complete", "excellent".
	CompletionQuality string  `json:"completionQuality" api:"required"`
	FinalAnswer       string  `json:"finalAnswer" api:"required"`
	IterationID       string  `json:"iterationId" api:"required"`
	Observation       string  `json:"observation" api:"required"`
	Timestamp         float64 `json:"timestamp" api:"required"`
	Feedback          string  `json:"feedback"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletionQuality respjson.Field
		FinalAnswer       respjson.Field
		IterationID       respjson.Field
		Observation       respjson.Field
		Timestamp         respjson.Field
		Feedback          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventTaskValidatedData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventTaskValidatedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "task:validation_error" event from /v1/automate.
type AutomateEventTaskValidationError struct {
	// Event data for validation errors during action response processing
	Data  AutomateEventTaskValidationErrorData `json:"data" api:"required"`
	Event constant.TaskValidationError         `json:"event" default:"task:validation_error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventTaskValidationError) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventTaskValidationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event data for validation errors during action response processing
type AutomateEventTaskValidationErrorData struct {
	Errors      []string `json:"errors" api:"required"`
	IterationID string   `json:"iterationId" api:"required"`
	RawResponse any      `json:"rawResponse" api:"required"`
	RetryCount  float64  `json:"retryCount" api:"required"`
	Timestamp   float64  `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Errors      respjson.Field
		IterationID respjson.Field
		RawResponse respjson.Field
		RetryCount  respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEventTaskValidationErrorData) RawJSON() string { return r.JSON.raw }
func (r *AutomateEventTaskValidationErrorData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResearchEventUnion contains all possible properties and values from
// [ResearchEventAnalyzingEnd], [ResearchEventAnalyzingStart],
// [ResearchEventComplete], [ResearchEventError], [ResearchEventEvaluatingEnd],
// [ResearchEventEvaluatingStart], [ResearchEventFollowingEnd],
// [ResearchEventFollowingStart], [ResearchEventIterationEnd],
// [ResearchEventIterationStart], [ResearchEventJudgingEnd],
// [ResearchEventJudgingStart], [ResearchEventOutliningEnd],
// [ResearchEventOutliningStart], [ResearchEventPlanningEnd],
// [ResearchEventPlanningStart], [ResearchEventPrefetchingEnd],
// [ResearchEventPrefetchingStart], [ResearchEventSearchingEnd],
// [ResearchEventSearchingStart], [ResearchEventStart], [ResearchEventWritingEnd],
// [ResearchEventWritingStart].
//
// Use the [ResearchEventUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResearchEventUnion struct {
	// This field is a union of [ResearchEventAnalyzingEndData],
	// [ResearchEventAnalyzingStartData], [ResearchEventCompleteData],
	// [ResearchEventErrorData], [ResearchEventEvaluatingEndData],
	// [ResearchEventEvaluatingStartData], [ResearchEventFollowingEndData],
	// [ResearchEventFollowingStartData], [ResearchEventIterationEndData],
	// [ResearchEventIterationStartData], [ResearchEventJudgingEndData],
	// [ResearchEventJudgingStartData], [ResearchEventOutliningEndData],
	// [ResearchEventOutliningStartData], [ResearchEventPlanningEndData],
	// [ResearchEventPlanningStartData], [ResearchEventPrefetchingEndData],
	// [ResearchEventPrefetchingStartData], [ResearchEventSearchingEndData],
	// [ResearchEventSearchingStartData], [ResearchEventStartData],
	// [ResearchEventWritingEndData], [ResearchEventWritingStartData]
	Data ResearchEventUnionData `json:"data"`
	// Any of "analyzing:end", "analyzing:start", "complete", "error",
	// "evaluating:end", "evaluating:start", "following:end", "following:start",
	// "iteration:end", "iteration:start", "judging:end", "judging:start",
	// "outlining:end", "outlining:start", "planning:end", "planning:start",
	// "prefetching:end", "prefetching:start", "searching:end", "searching:start",
	// "start", "writing:end", "writing:start".
	Event string `json:"event"`
	JSON  struct {
		Data  respjson.Field
		Event respjson.Field
		raw   string
	} `json:"-"`
}

// anyResearchEvent is implemented by each variant of [ResearchEventUnion] to add
// type safety for the return type of [ResearchEventUnion.AsAny]
type anyResearchEvent interface {
	implResearchEventUnion()
}

func (ResearchEventAnalyzingEnd) implResearchEventUnion()     {}
func (ResearchEventAnalyzingStart) implResearchEventUnion()   {}
func (ResearchEventComplete) implResearchEventUnion()         {}
func (ResearchEventError) implResearchEventUnion()            {}
func (ResearchEventEvaluatingEnd) implResearchEventUnion()    {}
func (ResearchEventEvaluatingStart) implResearchEventUnion()  {}
func (ResearchEventFollowingEnd) implResearchEventUnion()     {}
func (ResearchEventFollowingStart) implResearchEventUnion()   {}
func (ResearchEventIterationEnd) implResearchEventUnion()     {}
func (ResearchEventIterationStart) implResearchEventUnion()   {}
func (ResearchEventJudgingEnd) implResearchEventUnion()       {}
func (ResearchEventJudgingStart) implResearchEventUnion()     {}
func (ResearchEventOutliningEnd) implResearchEventUnion()     {}
func (ResearchEventOutliningStart) implResearchEventUnion()   {}
func (ResearchEventPlanningEnd) implResearchEventUnion()      {}
func (ResearchEventPlanningStart) implResearchEventUnion()    {}
func (ResearchEventPrefetchingEnd) implResearchEventUnion()   {}
func (ResearchEventPrefetchingStart) implResearchEventUnion() {}
func (ResearchEventSearchingEnd) implResearchEventUnion()     {}
func (ResearchEventSearchingStart) implResearchEventUnion()   {}
func (ResearchEventStart) implResearchEventUnion()            {}
func (ResearchEventWritingEnd) implResearchEventUnion()       {}
func (ResearchEventWritingStart) implResearchEventUnion()     {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResearchEventUnion.AsAny().(type) {
//	case tabstack.ResearchEventAnalyzingEnd:
//	case tabstack.ResearchEventAnalyzingStart:
//	case tabstack.ResearchEventComplete:
//	case tabstack.ResearchEventError:
//	case tabstack.ResearchEventEvaluatingEnd:
//	case tabstack.ResearchEventEvaluatingStart:
//	case tabstack.ResearchEventFollowingEnd:
//	case tabstack.ResearchEventFollowingStart:
//	case tabstack.ResearchEventIterationEnd:
//	case tabstack.ResearchEventIterationStart:
//	case tabstack.ResearchEventJudgingEnd:
//	case tabstack.ResearchEventJudgingStart:
//	case tabstack.ResearchEventOutliningEnd:
//	case tabstack.ResearchEventOutliningStart:
//	case tabstack.ResearchEventPlanningEnd:
//	case tabstack.ResearchEventPlanningStart:
//	case tabstack.ResearchEventPrefetchingEnd:
//	case tabstack.ResearchEventPrefetchingStart:
//	case tabstack.ResearchEventSearchingEnd:
//	case tabstack.ResearchEventSearchingStart:
//	case tabstack.ResearchEventStart:
//	case tabstack.ResearchEventWritingEnd:
//	case tabstack.ResearchEventWritingStart:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResearchEventUnion) AsAny() anyResearchEvent {
	switch u.Event {
	case "analyzing:end":
		return u.AsAnalyzingEnd()
	case "analyzing:start":
		return u.AsAnalyzingStart()
	case "complete":
		return u.AsComplete()
	case "error":
		return u.AsError()
	case "evaluating:end":
		return u.AsEvaluatingEnd()
	case "evaluating:start":
		return u.AsEvaluatingStart()
	case "following:end":
		return u.AsFollowingEnd()
	case "following:start":
		return u.AsFollowingStart()
	case "iteration:end":
		return u.AsIterationEnd()
	case "iteration:start":
		return u.AsIterationStart()
	case "judging:end":
		return u.AsJudgingEnd()
	case "judging:start":
		return u.AsJudgingStart()
	case "outlining:end":
		return u.AsOutliningEnd()
	case "outlining:start":
		return u.AsOutliningStart()
	case "planning:end":
		return u.AsPlanningEnd()
	case "planning:start":
		return u.AsPlanningStart()
	case "prefetching:end":
		return u.AsPrefetchingEnd()
	case "prefetching:start":
		return u.AsPrefetchingStart()
	case "searching:end":
		return u.AsSearchingEnd()
	case "searching:start":
		return u.AsSearchingStart()
	case "start":
		return u.AsStart()
	case "writing:end":
		return u.AsWritingEnd()
	case "writing:start":
		return u.AsWritingStart()
	}
	return nil
}

func (u ResearchEventUnion) AsAnalyzingEnd() (v ResearchEventAnalyzingEnd) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResearchEventUnion) AsAnalyzingStart() (v ResearchEventAnalyzingStart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResearchEventUnion) AsComplete() (v ResearchEventComplete) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResearchEventUnion) AsError() (v ResearchEventError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResearchEventUnion) AsEvaluatingEnd() (v ResearchEventEvaluatingEnd) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResearchEventUnion) AsEvaluatingStart() (v ResearchEventEvaluatingStart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResearchEventUnion) AsFollowingEnd() (v ResearchEventFollowingEnd) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResearchEventUnion) AsFollowingStart() (v ResearchEventFollowingStart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResearchEventUnion) AsIterationEnd() (v ResearchEventIterationEnd) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResearchEventUnion) AsIterationStart() (v ResearchEventIterationStart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResearchEventUnion) AsJudgingEnd() (v ResearchEventJudgingEnd) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResearchEventUnion) AsJudgingStart() (v ResearchEventJudgingStart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResearchEventUnion) AsOutliningEnd() (v ResearchEventOutliningEnd) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResearchEventUnion) AsOutliningStart() (v ResearchEventOutliningStart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResearchEventUnion) AsPlanningEnd() (v ResearchEventPlanningEnd) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResearchEventUnion) AsPlanningStart() (v ResearchEventPlanningStart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResearchEventUnion) AsPrefetchingEnd() (v ResearchEventPrefetchingEnd) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResearchEventUnion) AsPrefetchingStart() (v ResearchEventPrefetchingStart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResearchEventUnion) AsSearchingEnd() (v ResearchEventSearchingEnd) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResearchEventUnion) AsSearchingStart() (v ResearchEventSearchingStart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResearchEventUnion) AsStart() (v ResearchEventStart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResearchEventUnion) AsWritingEnd() (v ResearchEventWritingEnd) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResearchEventUnion) AsWritingStart() (v ResearchEventWritingStart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResearchEventUnion) RawJSON() string { return u.JSON.raw }

func (r *ResearchEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResearchEventUnionData is an implicit subunion of [ResearchEventUnion].
// ResearchEventUnionData provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [ResearchEventUnion].
type ResearchEventUnionData struct {
	// This field is from variant [ResearchEventAnalyzingEndData].
	Analyzed  float64 `json:"analyzed"`
	Failed    float64 `json:"failed"`
	Iteration float64 `json:"iteration"`
	Message   string  `json:"message"`
	// This field is a union of [[]ResearchEventAnalyzingEndDataSample],
	// [[]ResearchEventFollowingEndDataSample]
	Samples   ResearchEventUnionDataSamples `json:"samples"`
	Timestamp float64                       `json:"timestamp"`
	// This field is from variant [ResearchEventAnalyzingStartData].
	PageCount float64 `json:"pageCount"`
	// This field is from variant [ResearchEventCompleteData].
	Metadata ResearchEventCompleteDataMetadata `json:"metadata"`
	// This field is from variant [ResearchEventCompleteData].
	Report string `json:"report"`
	// This field is from variant [ResearchEventErrorData].
	Error ResearchEventErrorDataError `json:"error"`
	// This field is from variant [ResearchEventErrorData].
	Activity string `json:"activity"`
	// This field is from variant [ResearchEventEvaluatingEndData].
	Coverage string `json:"coverage"`
	// This field is from variant [ResearchEventEvaluatingEndData].
	Gaps string `json:"gaps"`
	// This field is from variant [ResearchEventEvaluatingEndData].
	NextQueries []string `json:"nextQueries"`
	// This field is from variant [ResearchEventEvaluatingEndData].
	QuestionAssessments []ResearchEventEvaluatingEndDataQuestionAssessment `json:"questionAssessments"`
	// This field is from variant [ResearchEventEvaluatingEndData].
	ShouldContinue bool    `json:"shouldContinue"`
	PagesAnalyzed  float64 `json:"pagesAnalyzed"`
	// This field is from variant [ResearchEventEvaluatingStartData].
	QuestionCount float64 `json:"questionCount"`
	// This field is from variant [ResearchEventFollowingEndData].
	Followed float64 `json:"followed"`
	// This field is from variant [ResearchEventFollowingStartData].
	LinkCount float64 `json:"linkCount"`
	// This field is from variant [ResearchEventIterationEndData].
	IsLast bool `json:"isLast"`
	// This field is from variant [ResearchEventIterationEndData].
	StopReason string `json:"stopReason"`
	// This field is from variant [ResearchEventIterationStartData].
	MaxIterations float64  `json:"maxIterations"`
	Queries       []string `json:"queries"`
	// This field is from variant [ResearchEventJudgingEndData].
	Approved bool    `json:"approved"`
	Attempt  float64 `json:"attempt"`
	// This field is from variant [ResearchEventJudgingEndData].
	Score float64 `json:"score"`
	// This field is from variant [ResearchEventJudgingEndData].
	Feedback    string  `json:"feedback"`
	MaxAttempts float64 `json:"maxAttempts"`
	// This field is from variant [ResearchEventOutliningEndData].
	SourcesSelected float64 `json:"sourcesSelected"`
	// This field is from variant [ResearchEventOutliningStartData].
	QualityPageCount float64 `json:"qualityPageCount"`
	// This field is from variant [ResearchEventPlanningEndData].
	Complexity string `json:"complexity"`
	// This field is from variant [ResearchEventPlanningEndData].
	Objective string `json:"objective"`
	// This field is from variant [ResearchEventPlanningEndData].
	Plan string `json:"plan"`
	// This field is from variant [ResearchEventPlanningEndData].
	Questions []string `json:"questions"`
	// This field is from variant [ResearchEventPlanningStartData].
	HasPrefetchedContext bool `json:"hasPrefetchedContext"`
	// This field is from variant [ResearchEventPrefetchingEndData].
	Fetched float64 `json:"fetched"`
	// This field is from variant [ResearchEventPrefetchingStartData].
	URLCount float64 `json:"urlCount"`
	// This field is from variant [ResearchEventPrefetchingStartData].
	URLs []string `json:"urls"`
	// This field is from variant [ResearchEventSearchingEndData].
	URLsFound float64 `json:"urlsFound"`
	// This field is from variant [ResearchEventSearchingEndData].
	URLsNew float64 `json:"urlsNew"`
	// This field is from variant [ResearchEventWritingStartData].
	IsRevision bool `json:"isRevision"`
	// This field is from variant [ResearchEventWritingStartData].
	PreviousScore float64 `json:"previousScore"`
	JSON          struct {
		Analyzed             respjson.Field
		Failed               respjson.Field
		Iteration            respjson.Field
		Message              respjson.Field
		Samples              respjson.Field
		Timestamp            respjson.Field
		PageCount            respjson.Field
		Metadata             respjson.Field
		Report               respjson.Field
		Error                respjson.Field
		Activity             respjson.Field
		Coverage             respjson.Field
		Gaps                 respjson.Field
		NextQueries          respjson.Field
		QuestionAssessments  respjson.Field
		ShouldContinue       respjson.Field
		PagesAnalyzed        respjson.Field
		QuestionCount        respjson.Field
		Followed             respjson.Field
		LinkCount            respjson.Field
		IsLast               respjson.Field
		StopReason           respjson.Field
		MaxIterations        respjson.Field
		Queries              respjson.Field
		Approved             respjson.Field
		Attempt              respjson.Field
		Score                respjson.Field
		Feedback             respjson.Field
		MaxAttempts          respjson.Field
		SourcesSelected      respjson.Field
		QualityPageCount     respjson.Field
		Complexity           respjson.Field
		Objective            respjson.Field
		Plan                 respjson.Field
		Questions            respjson.Field
		HasPrefetchedContext respjson.Field
		Fetched              respjson.Field
		URLCount             respjson.Field
		URLs                 respjson.Field
		URLsFound            respjson.Field
		URLsNew              respjson.Field
		IsRevision           respjson.Field
		PreviousScore        respjson.Field
		raw                  string
	} `json:"-"`
}

func (r *ResearchEventUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResearchEventUnionDataSamples is an implicit subunion of [ResearchEventUnion].
// ResearchEventUnionDataSamples provides convenient access to the sub-properties
// of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ResearchEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfResearchEventAnalyzingEndDataSamples
// OfResearchEventFollowingEndDataSamples]
type ResearchEventUnionDataSamples struct {
	// This field will be present if the value is a
	// [[]ResearchEventAnalyzingEndDataSample] instead of an object.
	OfResearchEventAnalyzingEndDataSamples []ResearchEventAnalyzingEndDataSample `json:",inline"`
	// This field will be present if the value is a
	// [[]ResearchEventFollowingEndDataSample] instead of an object.
	OfResearchEventFollowingEndDataSamples []ResearchEventFollowingEndDataSample `json:",inline"`
	JSON                                   struct {
		OfResearchEventAnalyzingEndDataSamples respjson.Field
		OfResearchEventFollowingEndDataSamples respjson.Field
		raw                                    string
	} `json:"-"`
}

func (r *ResearchEventUnionDataSamples) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "analyzing:end" event from /v1/research.
type ResearchEventAnalyzingEnd struct {
	Data  ResearchEventAnalyzingEndData `json:"data" api:"required"`
	Event constant.AnalyzingEnd         `json:"event" default:"analyzing:end"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventAnalyzingEnd) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventAnalyzingEnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchEventAnalyzingEndData struct {
	Analyzed  float64                               `json:"analyzed" api:"required"`
	Failed    float64                               `json:"failed" api:"required"`
	Iteration float64                               `json:"iteration" api:"required"`
	Message   string                                `json:"message" api:"required"`
	Samples   []ResearchEventAnalyzingEndDataSample `json:"samples" api:"required"`
	Timestamp float64                               `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Analyzed    respjson.Field
		Failed      respjson.Field
		Iteration   respjson.Field
		Message     respjson.Field
		Samples     respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventAnalyzingEndData) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventAnalyzingEndData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Page sample - lightweight representation for event payloads
type ResearchEventAnalyzingEndDataSample struct {
	Domain string `json:"domain" api:"required"`
	Title  string `json:"title" api:"required"`
	URL    string `json:"url" api:"required"`
	// URL source tracking - where a URL came from
	//
	// Any of "user-input", "search-result", "extracted-link".
	URLSource string `json:"urlSource" api:"required"`
	// Any of "low", "medium", "high".
	Relevance string `json:"relevance"`
	// Any of "low", "medium", "high".
	Reliability string `json:"reliability"`
	Summary     string `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		URLSource   respjson.Field
		Relevance   respjson.Field
		Reliability respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventAnalyzingEndDataSample) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventAnalyzingEndDataSample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "analyzing:start" event from /v1/research.
type ResearchEventAnalyzingStart struct {
	Data  ResearchEventAnalyzingStartData `json:"data" api:"required"`
	Event constant.AnalyzingStart         `json:"event" default:"analyzing:start"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventAnalyzingStart) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventAnalyzingStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchEventAnalyzingStartData struct {
	Iteration float64 `json:"iteration" api:"required"`
	Message   string  `json:"message" api:"required"`
	PageCount float64 `json:"pageCount" api:"required"`
	Timestamp float64 `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Iteration   respjson.Field
		Message     respjson.Field
		PageCount   respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventAnalyzingStartData) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventAnalyzingStartData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "complete" event from /v1/research.
type ResearchEventComplete struct {
	// complete - Research finished successfully
	Data  ResearchEventCompleteData `json:"data" api:"required"`
	Event constant.Complete         `json:"event" default:"complete"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventComplete) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventComplete) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// complete - Research finished successfully
type ResearchEventCompleteData struct {
	Message string `json:"message" api:"required"`
	// Research metadata
	//
	// Note: citedPages, gapEvaluations, outline, and judgments are optional to support
	// fast mode, which skips these phases for maximum speed.
	Metadata  ResearchEventCompleteDataMetadata `json:"metadata" api:"required"`
	Report    string                            `json:"report" api:"required"`
	Timestamp float64                           `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Metadata    respjson.Field
		Report      respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventCompleteData) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventCompleteData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Research metadata
//
// Note: citedPages, gapEvaluations, outline, and judgments are optional to support
// fast mode, which skips these phases for maximum speed.
type ResearchEventCompleteDataMetadata struct {
	ExecutedQueries [][]string `json:"executedQueries" api:"required"`
	// Research mode determines depth, thinking budget, and quality controls
	//
	// Modes (in order of cost/thoroughness):
	//
	//   - **fast**: Quick answers with minimal validation (~$2, 1 iteration, no judge)
	//   - **balanced**: Standard research with moderate depth (~$8, 3 iterations, Flash
	//     models, no judge)
	//   - **deep**: Thorough research with judge review (~$15, 5 iterations, Flash
	//     models, with judge)
	//   - **max**: Maximum quality with Pro models (~$40, 5 iterations, Pro models, with
	//     judge)
	//   - **ultra**: Ultimate tier - all Pro models, 10 iterations (expensive, for when
	//     accuracy is paramount)
	//
	// Any of "fast", "balanced", "deep", "max", "ultra".
	Mode   string `json:"mode" api:"required"`
	Prompt string `json:"prompt" api:"required"`
	// Any of "simple", "moderate", "complex".
	QueryComplexity   string   `json:"queryComplexity" api:"required"`
	ResearchObjective string   `json:"researchObjective" api:"required"`
	ResearchPlan      string   `json:"researchPlan" api:"required"`
	ResearchQuestions []string `json:"researchQuestions" api:"required"`
	// Total pages analyzed across all iterations
	TotalPagesAnalyzed float64 `json:"totalPagesAnalyzed" api:"required"`
	// Pages cited in the report, ordered by first citation appearance
	CitedPages     []ResearchEventCompleteDataMetadataCitedPage     `json:"citedPages"`
	GapEvaluations []ResearchEventCompleteDataMetadataGapEvaluation `json:"gapEvaluations"`
	Judgments      []ResearchEventCompleteDataMetadataJudgment      `json:"judgments"`
	// Complete research metrics
	Metrics ResearchEventCompleteDataMetadataMetrics `json:"metrics"`
	// Report outline from research writer
	Outline    ResearchEventCompleteDataMetadataOutline    `json:"outline"`
	URLSources ResearchEventCompleteDataMetadataURLSources `json:"urlSources"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExecutedQueries    respjson.Field
		Mode               respjson.Field
		Prompt             respjson.Field
		QueryComplexity    respjson.Field
		ResearchObjective  respjson.Field
		ResearchPlan       respjson.Field
		ResearchQuestions  respjson.Field
		TotalPagesAnalyzed respjson.Field
		CitedPages         respjson.Field
		GapEvaluations     respjson.Field
		Judgments          respjson.Field
		Metrics            respjson.Field
		Outline            respjson.Field
		URLSources         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventCompleteDataMetadata) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventCompleteDataMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchEventCompleteDataMetadataCitedPage struct {
	ID            string   `json:"id" api:"required"`
	Claims        []string `json:"claims" api:"required"`
	SourceQueries []string `json:"sourceQueries" api:"required"`
	URL           string   `json:"url" api:"required"`
	Depth         float64  `json:"depth"`
	// Full page text (fetched markdown or search excerpts). Only populated when
	// `includeFullText: true` in ResearchOptions.
	//
	// - Fast mode: Parallel API excerpts (~5000 chars)
	// - Other modes: Fetched page markdown
	FullText  string `json:"fullText"`
	ParentURL string `json:"parentUrl"`
	// Any of "low", "medium", "high".
	Relevance string `json:"relevance"`
	// Any of "low", "medium", "high".
	Reliability string `json:"reliability"`
	// LLM-generated summary. Undefined in fast mode (no content analysis).
	Summary string `json:"summary"`
	Title   string `json:"title"`
	// URL source tracking - where a URL came from
	//
	// Any of "user-input", "search-result", "extracted-link".
	URLSource string `json:"urlSource"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Claims        respjson.Field
		SourceQueries respjson.Field
		URL           respjson.Field
		Depth         respjson.Field
		FullText      respjson.Field
		ParentURL     respjson.Field
		Relevance     respjson.Field
		Reliability   respjson.Field
		Summary       respjson.Field
		Title         respjson.Field
		URLSource     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventCompleteDataMetadataCitedPage) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventCompleteDataMetadataCitedPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Gap evaluation results from research strategist
type ResearchEventCompleteDataMetadataGapEvaluation struct {
	// Based on unanswered/partial questions, what specific information is still
	// needed?
	GapDescription string `json:"gapDescription" api:"required"`
	// Assessment of each research question's status and findings
	QuestionAssessments []ResearchEventCompleteDataMetadataGapEvaluationQuestionAssessment `json:"questionAssessments" api:"required"`
	// Research coverage level - assesses quality across all questions.
	//
	// Hierarchy: Light < Moderate < Solid < Comprehensive
	//
	//   - **Light**: Basic info on some questions, most need more depth → Continue
	//   - **Moderate**: Multiple questions answered, some remain partial → Continue
	//   - **Solid**: Most questions well-answered with validated sources → Sufficient to
	//     stop
	//   - **Comprehensive**: All questions thoroughly answered, exceptional depth →
	//     Definitely stop
	//
	// Any of "Light", "Moderate", "Solid", "Comprehensive".
	ResearchCoverage string `json:"researchCoverage" api:"required"`
	// Explicit decision: should research continue with another iteration?
	//
	//   - Considers: how many questions unanswered/partial, coverage for mode, remaining
	//     iterations
	//   - Drives query generation: true → generate queries, false → stop researching
	ShouldContinueResearch bool `json:"shouldContinueResearch" api:"required"`
	// New research questions to add (optional, use sparingly)
	//
	// - Only if original decomposition missed something critical
	// - Maximum 2-3 new questions total across all iterations
	// - Most iterations should return empty array or omit this field
	NewResearchQuestions []string `json:"newResearchQuestions"`
	// Search queries to address identified gaps (only when shouldContinueResearch is
	// true)
	//
	// - Target unanswered questions first, then partial questions
	// - 3-10 targeted queries if shouldContinueResearch is true
	// - Omit or provide empty array if shouldContinueResearch is false
	SearchQueries []string `json:"searchQueries"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GapDescription         respjson.Field
		QuestionAssessments    respjson.Field
		ResearchCoverage       respjson.Field
		ShouldContinueResearch respjson.Field
		NewResearchQuestions   respjson.Field
		SearchQueries          respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventCompleteDataMetadataGapEvaluation) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventCompleteDataMetadataGapEvaluation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Assessment of a single research question
type ResearchEventCompleteDataMetadataGapEvaluationQuestionAssessment struct {
	// What we learned (if answered/partial) or what's missing (if unanswered)
	Findings string `json:"findings" api:"required"`
	// The research question being assessed
	Question string `json:"question" api:"required"`
	// Status: answered (clear info), partial (some info, gaps remain), unanswered (no
	// relevant info)
	//
	// Any of "answered", "partial", "unanswered".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Findings    respjson.Field
		Question    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventCompleteDataMetadataGapEvaluationQuestionAssessment) RawJSON() string {
	return r.JSON.raw
}
func (r *ResearchEventCompleteDataMetadataGapEvaluationQuestionAssessment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Judgment result from research judge
type ResearchEventCompleteDataMetadataJudgment struct {
	Approved    bool    `json:"approved" api:"required"`
	Observation string  `json:"observation" api:"required"`
	Score       float64 `json:"score" api:"required"`
	Feedback    string  `json:"feedback"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Approved    respjson.Field
		Observation respjson.Field
		Score       respjson.Field
		Feedback    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventCompleteDataMetadataJudgment) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventCompleteDataMetadataJudgment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete research metrics
type ResearchEventCompleteDataMetadataMetrics struct {
	// Cached fetch count (subset of fetches)
	CachedFetches float64 `json:"cachedFetches" api:"required"`
	// Cached search count by provider name (subset of searches)
	CachedSearches map[string]float64 `json:"cachedSearches" api:"required"`
	// Fetch count (number of pages fetched)
	Fetches float64 `json:"fetches" api:"required"`
	// Number of research iterations performed
	Iterations float64 `json:"iterations" api:"required"`
	// Phase timings with duration in milliseconds
	Phases map[string]ResearchEventCompleteDataMetadataMetricsPhase `json:"phases" api:"required"`
	// Number of URLs blocked by robots.txt
	RobotsBlocked float64 `json:"robotsBlocked" api:"required"`
	// Search count by provider name (e.g., "bright-data", "parallel")
	Searches map[string]float64 `json:"searches" api:"required"`
	// Success rate metrics
	SuccessRates ResearchEventCompleteDataMetadataMetricsSuccessRates `json:"successRates" api:"required"`
	// Token usage by model ID (e.g., "gemini-2.5-flash")
	Tokens map[string]ResearchEventCompleteDataMetadataMetricsToken `json:"tokens" api:"required"`
	// Total duration in milliseconds
	TotalDuration float64 `json:"totalDuration" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CachedFetches  respjson.Field
		CachedSearches respjson.Field
		Fetches        respjson.Field
		Iterations     respjson.Field
		Phases         respjson.Field
		RobotsBlocked  respjson.Field
		Searches       respjson.Field
		SuccessRates   respjson.Field
		Tokens         respjson.Field
		TotalDuration  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventCompleteDataMetadataMetrics) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventCompleteDataMetadataMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchEventCompleteDataMetadataMetricsPhase struct {
	Duration float64 `json:"duration" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventCompleteDataMetadataMetricsPhase) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventCompleteDataMetadataMetricsPhase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Success rate metrics
type ResearchEventCompleteDataMetadataMetricsSuccessRates struct {
	Analyzes float64 `json:"analyzes" api:"required"`
	Fetches  float64 `json:"fetches" api:"required"`
	Searches float64 `json:"searches" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Analyzes    respjson.Field
		Fetches     respjson.Field
		Searches    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventCompleteDataMetadataMetricsSuccessRates) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventCompleteDataMetadataMetricsSuccessRates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token usage for a specific model
type ResearchEventCompleteDataMetadataMetricsToken struct {
	Input  float64 `json:"input" api:"required"`
	Output float64 `json:"output" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Input       respjson.Field
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventCompleteDataMetadataMetricsToken) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventCompleteDataMetadataMetricsToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Report outline from research writer
type ResearchEventCompleteDataMetadataOutline struct {
	DirectAnswer      string   `json:"directAnswer" api:"required"`
	KeyTakeaways      []string `json:"keyTakeaways" api:"required"`
	Outline           string   `json:"outline" api:"required"`
	RelevantSourceIDs []string `json:"relevantSourceIds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DirectAnswer      respjson.Field
		KeyTakeaways      respjson.Field
		Outline           respjson.Field
		RelevantSourceIDs respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventCompleteDataMetadataOutline) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventCompleteDataMetadataOutline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchEventCompleteDataMetadataURLSources struct {
	ExtractedLinks float64 `json:"extractedLinks" api:"required"`
	SearchResults  float64 `json:"searchResults" api:"required"`
	UserProvided   float64 `json:"userProvided" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtractedLinks respjson.Field
		SearchResults  respjson.Field
		UserProvided   respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventCompleteDataMetadataURLSources) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventCompleteDataMetadataURLSources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "error" event from /v1/research.
type ResearchEventError struct {
	// error - Research failed
	Data  ResearchEventErrorData `json:"data" api:"required"`
	Event constant.Error         `json:"event" default:"error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventError) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// error - Research failed
type ResearchEventErrorData struct {
	Error     ResearchEventErrorDataError `json:"error" api:"required"`
	Message   string                      `json:"message" api:"required"`
	Timestamp float64                     `json:"timestamp" api:"required"`
	// Activity types for research workflow
	//
	// Any of "prefetching", "planning", "iteration", "searching", "analyzing",
	// "following", "evaluating", "outlining", "writing", "judging".
	Activity  string  `json:"activity"`
	Iteration float64 `json:"iteration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		Message     respjson.Field
		Timestamp   respjson.Field
		Activity    respjson.Field
		Iteration   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventErrorData) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventErrorData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchEventErrorDataError struct {
	Message string `json:"message" api:"required"`
	Name    string `json:"name" api:"required"`
	Stack   string `json:"stack"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Name        respjson.Field
		Stack       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventErrorDataError) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventErrorDataError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "evaluating:end" event from /v1/research.
type ResearchEventEvaluatingEnd struct {
	Data  ResearchEventEvaluatingEndData `json:"data" api:"required"`
	Event constant.EvaluatingEnd         `json:"event" default:"evaluating:end"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventEvaluatingEnd) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventEvaluatingEnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchEventEvaluatingEndData struct {
	// Any of "Light", "Moderate", "Solid", "Comprehensive".
	Coverage            string                                             `json:"coverage" api:"required"`
	Gaps                string                                             `json:"gaps" api:"required"`
	Iteration           float64                                            `json:"iteration" api:"required"`
	Message             string                                             `json:"message" api:"required"`
	NextQueries         []string                                           `json:"nextQueries" api:"required"`
	QuestionAssessments []ResearchEventEvaluatingEndDataQuestionAssessment `json:"questionAssessments" api:"required"`
	ShouldContinue      bool                                               `json:"shouldContinue" api:"required"`
	Timestamp           float64                                            `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Coverage            respjson.Field
		Gaps                respjson.Field
		Iteration           respjson.Field
		Message             respjson.Field
		NextQueries         respjson.Field
		QuestionAssessments respjson.Field
		ShouldContinue      respjson.Field
		Timestamp           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventEvaluatingEndData) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventEvaluatingEndData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Assessment of a single research question
type ResearchEventEvaluatingEndDataQuestionAssessment struct {
	// What we learned (if answered/partial) or what's missing (if unanswered)
	Findings string `json:"findings" api:"required"`
	// The research question being assessed
	Question string `json:"question" api:"required"`
	// Status: answered (clear info), partial (some info, gaps remain), unanswered (no
	// relevant info)
	//
	// Any of "answered", "partial", "unanswered".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Findings    respjson.Field
		Question    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventEvaluatingEndDataQuestionAssessment) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventEvaluatingEndDataQuestionAssessment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "evaluating:start" event from /v1/research.
type ResearchEventEvaluatingStart struct {
	Data  ResearchEventEvaluatingStartData `json:"data" api:"required"`
	Event constant.EvaluatingStart         `json:"event" default:"evaluating:start"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventEvaluatingStart) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventEvaluatingStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchEventEvaluatingStartData struct {
	Iteration float64 `json:"iteration" api:"required"`
	Message   string  `json:"message" api:"required"`
	// Total pages analyzed so far (including this iteration)
	PagesAnalyzed float64 `json:"pagesAnalyzed" api:"required"`
	// Number of research questions being assessed
	QuestionCount float64 `json:"questionCount" api:"required"`
	Timestamp     float64 `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Iteration     respjson.Field
		Message       respjson.Field
		PagesAnalyzed respjson.Field
		QuestionCount respjson.Field
		Timestamp     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventEvaluatingStartData) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventEvaluatingStartData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "following:end" event from /v1/research.
type ResearchEventFollowingEnd struct {
	Data  ResearchEventFollowingEndData `json:"data" api:"required"`
	Event constant.FollowingEnd         `json:"event" default:"following:end"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventFollowingEnd) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventFollowingEnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchEventFollowingEndData struct {
	Failed    float64                               `json:"failed" api:"required"`
	Followed  float64                               `json:"followed" api:"required"`
	Iteration float64                               `json:"iteration" api:"required"`
	Message   string                                `json:"message" api:"required"`
	Samples   []ResearchEventFollowingEndDataSample `json:"samples" api:"required"`
	Timestamp float64                               `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Failed      respjson.Field
		Followed    respjson.Field
		Iteration   respjson.Field
		Message     respjson.Field
		Samples     respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventFollowingEndData) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventFollowingEndData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Page sample - lightweight representation for event payloads
type ResearchEventFollowingEndDataSample struct {
	Domain string `json:"domain" api:"required"`
	Title  string `json:"title" api:"required"`
	URL    string `json:"url" api:"required"`
	// URL source tracking - where a URL came from
	//
	// Any of "user-input", "search-result", "extracted-link".
	URLSource string `json:"urlSource" api:"required"`
	// Any of "low", "medium", "high".
	Relevance string `json:"relevance"`
	// Any of "low", "medium", "high".
	Reliability string `json:"reliability"`
	Summary     string `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		URLSource   respjson.Field
		Relevance   respjson.Field
		Reliability respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventFollowingEndDataSample) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventFollowingEndDataSample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "following:start" event from /v1/research.
type ResearchEventFollowingStart struct {
	Data  ResearchEventFollowingStartData `json:"data" api:"required"`
	Event constant.FollowingStart         `json:"event" default:"following:start"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventFollowingStart) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventFollowingStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchEventFollowingStartData struct {
	Iteration float64 `json:"iteration" api:"required"`
	LinkCount float64 `json:"linkCount" api:"required"`
	Message   string  `json:"message" api:"required"`
	Timestamp float64 `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Iteration   respjson.Field
		LinkCount   respjson.Field
		Message     respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventFollowingStartData) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventFollowingStartData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "iteration:end" event from /v1/research.
type ResearchEventIterationEnd struct {
	Data  ResearchEventIterationEndData `json:"data" api:"required"`
	Event constant.IterationEnd         `json:"event" default:"iteration:end"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventIterationEnd) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventIterationEnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchEventIterationEndData struct {
	// Whether this is the final iteration
	IsLast    bool    `json:"isLast" api:"required"`
	Iteration float64 `json:"iteration" api:"required"`
	Message   string  `json:"message" api:"required"`
	Timestamp float64 `json:"timestamp" api:"required"`
	// Why research iterations stopped (only present when isLast is true)
	//
	// Any of "max_iterations", "coverage_sufficient".
	StopReason string `json:"stopReason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsLast      respjson.Field
		Iteration   respjson.Field
		Message     respjson.Field
		Timestamp   respjson.Field
		StopReason  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventIterationEndData) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventIterationEndData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "iteration:start" event from /v1/research.
type ResearchEventIterationStart struct {
	Data  ResearchEventIterationStartData `json:"data" api:"required"`
	Event constant.IterationStart         `json:"event" default:"iteration:start"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventIterationStart) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventIterationStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchEventIterationStartData struct {
	Iteration float64 `json:"iteration" api:"required"`
	// Maximum iterations for this research mode
	MaxIterations float64 `json:"maxIterations" api:"required"`
	Message       string  `json:"message" api:"required"`
	// Search queries to execute in this iteration
	Queries   []string `json:"queries" api:"required"`
	Timestamp float64  `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Iteration     respjson.Field
		MaxIterations respjson.Field
		Message       respjson.Field
		Queries       respjson.Field
		Timestamp     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventIterationStartData) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventIterationStartData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "judging:end" event from /v1/research.
type ResearchEventJudgingEnd struct {
	Data  ResearchEventJudgingEndData `json:"data" api:"required"`
	Event constant.JudgingEnd         `json:"event" default:"judging:end"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventJudgingEnd) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventJudgingEnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchEventJudgingEndData struct {
	Approved  bool    `json:"approved" api:"required"`
	Attempt   float64 `json:"attempt" api:"required"`
	Message   string  `json:"message" api:"required"`
	Score     float64 `json:"score" api:"required"`
	Timestamp float64 `json:"timestamp" api:"required"`
	Feedback  string  `json:"feedback"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Approved    respjson.Field
		Attempt     respjson.Field
		Message     respjson.Field
		Score       respjson.Field
		Timestamp   respjson.Field
		Feedback    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventJudgingEndData) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventJudgingEndData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "judging:start" event from /v1/research.
type ResearchEventJudgingStart struct {
	Data  ResearchEventJudgingStartData `json:"data" api:"required"`
	Event constant.JudgingStart         `json:"event" default:"judging:start"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventJudgingStart) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventJudgingStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchEventJudgingStartData struct {
	Attempt float64 `json:"attempt" api:"required"`
	// Maximum attempts allowed (1 + maxRevisions)
	MaxAttempts float64 `json:"maxAttempts" api:"required"`
	Message     string  `json:"message" api:"required"`
	Timestamp   float64 `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attempt     respjson.Field
		MaxAttempts respjson.Field
		Message     respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventJudgingStartData) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventJudgingStartData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "outlining:end" event from /v1/research.
type ResearchEventOutliningEnd struct {
	Data  ResearchEventOutliningEndData `json:"data" api:"required"`
	Event constant.OutliningEnd         `json:"event" default:"outlining:end"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventOutliningEnd) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventOutliningEnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchEventOutliningEndData struct {
	Message         string  `json:"message" api:"required"`
	SourcesSelected float64 `json:"sourcesSelected" api:"required"`
	Timestamp       float64 `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message         respjson.Field
		SourcesSelected respjson.Field
		Timestamp       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventOutliningEndData) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventOutliningEndData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "outlining:start" event from /v1/research.
type ResearchEventOutliningStart struct {
	Data  ResearchEventOutliningStartData `json:"data" api:"required"`
	Event constant.OutliningStart         `json:"event" default:"outlining:start"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventOutliningStart) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventOutliningStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchEventOutliningStartData struct {
	Message string `json:"message" api:"required"`
	// Total pages analyzed across all iterations
	PagesAnalyzed float64 `json:"pagesAnalyzed" api:"required"`
	// Pages that meet quality threshold (medium+ relevance and reliability)
	QualityPageCount float64 `json:"qualityPageCount" api:"required"`
	Timestamp        float64 `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message          respjson.Field
		PagesAnalyzed    respjson.Field
		QualityPageCount respjson.Field
		Timestamp        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventOutliningStartData) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventOutliningStartData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "planning:end" event from /v1/research.
type ResearchEventPlanningEnd struct {
	Data  ResearchEventPlanningEndData `json:"data" api:"required"`
	Event constant.PlanningEnd         `json:"event" default:"planning:end"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventPlanningEnd) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventPlanningEnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchEventPlanningEndData struct {
	// Any of "simple", "moderate", "complex".
	Complexity string   `json:"complexity" api:"required"`
	Message    string   `json:"message" api:"required"`
	Objective  string   `json:"objective" api:"required"`
	Plan       string   `json:"plan" api:"required"`
	Queries    []string `json:"queries" api:"required"`
	Questions  []string `json:"questions" api:"required"`
	Timestamp  float64  `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Complexity  respjson.Field
		Message     respjson.Field
		Objective   respjson.Field
		Plan        respjson.Field
		Queries     respjson.Field
		Questions   respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventPlanningEndData) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventPlanningEndData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "planning:start" event from /v1/research.
type ResearchEventPlanningStart struct {
	Data  ResearchEventPlanningStartData `json:"data" api:"required"`
	Event constant.PlanningStart         `json:"event" default:"planning:start"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventPlanningStart) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventPlanningStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchEventPlanningStartData struct {
	// Whether prefetched user-provided URLs exist for context
	HasPrefetchedContext bool    `json:"hasPrefetchedContext" api:"required"`
	Message              string  `json:"message" api:"required"`
	Timestamp            float64 `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasPrefetchedContext respjson.Field
		Message              respjson.Field
		Timestamp            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventPlanningStartData) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventPlanningStartData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "prefetching:end" event from /v1/research.
type ResearchEventPrefetchingEnd struct {
	Data  ResearchEventPrefetchingEndData `json:"data" api:"required"`
	Event constant.PrefetchingEnd         `json:"event" default:"prefetching:end"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventPrefetchingEnd) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventPrefetchingEnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchEventPrefetchingEndData struct {
	Failed    float64 `json:"failed" api:"required"`
	Fetched   float64 `json:"fetched" api:"required"`
	Message   string  `json:"message" api:"required"`
	Timestamp float64 `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Failed      respjson.Field
		Fetched     respjson.Field
		Message     respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventPrefetchingEndData) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventPrefetchingEndData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "prefetching:start" event from /v1/research.
type ResearchEventPrefetchingStart struct {
	Data  ResearchEventPrefetchingStartData `json:"data" api:"required"`
	Event constant.PrefetchingStart         `json:"event" default:"prefetching:start"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventPrefetchingStart) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventPrefetchingStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchEventPrefetchingStartData struct {
	Message   string   `json:"message" api:"required"`
	Timestamp float64  `json:"timestamp" api:"required"`
	URLCount  float64  `json:"urlCount" api:"required"`
	URLs      []string `json:"urls" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Timestamp   respjson.Field
		URLCount    respjson.Field
		URLs        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventPrefetchingStartData) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventPrefetchingStartData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "searching:end" event from /v1/research.
type ResearchEventSearchingEnd struct {
	Data  ResearchEventSearchingEndData `json:"data" api:"required"`
	Event constant.SearchingEnd         `json:"event" default:"searching:end"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventSearchingEnd) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventSearchingEnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchEventSearchingEndData struct {
	Iteration float64 `json:"iteration" api:"required"`
	Message   string  `json:"message" api:"required"`
	Timestamp float64 `json:"timestamp" api:"required"`
	URLsFound float64 `json:"urlsFound" api:"required"`
	URLsNew   float64 `json:"urlsNew" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Iteration   respjson.Field
		Message     respjson.Field
		Timestamp   respjson.Field
		URLsFound   respjson.Field
		URLsNew     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventSearchingEndData) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventSearchingEndData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "searching:start" event from /v1/research.
type ResearchEventSearchingStart struct {
	Data  ResearchEventSearchingStartData `json:"data" api:"required"`
	Event constant.SearchingStart         `json:"event" default:"searching:start"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventSearchingStart) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventSearchingStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchEventSearchingStartData struct {
	Iteration float64  `json:"iteration" api:"required"`
	Message   string   `json:"message" api:"required"`
	Queries   []string `json:"queries" api:"required"`
	Timestamp float64  `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Iteration   respjson.Field
		Message     respjson.Field
		Queries     respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventSearchingStartData) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventSearchingStartData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "start" event from /v1/research.
type ResearchEventStart struct {
	// start - Research begins
	Data  ResearchEventStartData `json:"data" api:"required"`
	Event constant.Start         `json:"event" default:"start"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventStart) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// start - Research begins
type ResearchEventStartData struct {
	Message   string  `json:"message" api:"required"`
	Timestamp float64 `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventStartData) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventStartData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "writing:end" event from /v1/research.
type ResearchEventWritingEnd struct {
	Data  ResearchEventWritingEndData `json:"data" api:"required"`
	Event constant.WritingEnd         `json:"event" default:"writing:end"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventWritingEnd) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventWritingEnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchEventWritingEndData struct {
	Attempt   float64 `json:"attempt" api:"required"`
	Message   string  `json:"message" api:"required"`
	Timestamp float64 `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attempt     respjson.Field
		Message     respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventWritingEndData) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventWritingEndData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Envelope for the "writing:start" event from /v1/research.
type ResearchEventWritingStart struct {
	Data  ResearchEventWritingStartData `json:"data" api:"required"`
	Event constant.WritingStart         `json:"event" default:"writing:start"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventWritingStart) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventWritingStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchEventWritingStartData struct {
	Attempt float64 `json:"attempt" api:"required"`
	// Whether this is a revision attempt (attempt > 1)
	IsRevision bool `json:"isRevision" api:"required"`
	// Maximum attempts allowed (1 + maxRevisions)
	MaxAttempts float64 `json:"maxAttempts" api:"required"`
	Message     string  `json:"message" api:"required"`
	Timestamp   float64 `json:"timestamp" api:"required"`
	// Previous judgment score if this is a revision
	PreviousScore float64 `json:"previousScore"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attempt       respjson.Field
		IsRevision    respjson.Field
		MaxAttempts   respjson.Field
		Message       respjson.Field
		Timestamp     respjson.Field
		PreviousScore respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEventWritingStartData) RawJSON() string { return r.JSON.raw }
func (r *ResearchEventWritingStartData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentAutomateInputResponse struct {
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentAutomateInputResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentAutomateInputResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentAutomateParams struct {
	// The task description in natural language
	Task string `json:"task" api:"required"`
	// Safety constraints for execution
	Guardrails param.Opt[string] `json:"guardrails,omitzero"`
	// Enable interactive mode to allow human-in-the-loop input during task execution
	Interactive param.Opt[bool] `json:"interactive,omitzero"`
	// Maximum task iterations
	MaxIterations param.Opt[int64] `json:"maxIterations,omitzero"`
	// Maximum validation attempts
	MaxValidationAttempts param.Opt[int64] `json:"maxValidationAttempts,omitzero"`
	// Starting URL for the task
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	// JSON data to provide context for form filling or complex tasks
	Data any `json:"data,omitzero"`
	// Optional geotargeting parameters for proxy requests
	GeoTarget AgentAutomateParamsGeoTarget `json:"geo_target,omitzero"`
	paramObj
}

func (r AgentAutomateParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentAutomateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentAutomateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional geotargeting parameters for proxy requests
type AgentAutomateParamsGeoTarget struct {
	// Country code using ISO 3166-1 alpha-2 standard (2 letters, e.g., "US", "GB",
	// "JP"). See: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
	Country param.Opt[string] `json:"country,omitzero"`
	paramObj
}

func (r AgentAutomateParamsGeoTarget) MarshalJSON() (data []byte, err error) {
	type shadow AgentAutomateParamsGeoTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentAutomateParamsGeoTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentAutomateInputParams struct {
	// Set to true to cancel/decline the request
	Cancelled param.Opt[bool] `json:"cancelled,omitzero"`
	// Field values as array of {ref, value} pairs (required when not cancelled)
	Fields []AgentAutomateInputParamsField `json:"fields,omitzero"`
	paramObj
}

func (r AgentAutomateInputParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentAutomateInputParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentAutomateInputParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentAutomateInputParamsField struct {
	Ref   param.Opt[string] `json:"ref,omitzero"`
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r AgentAutomateInputParamsField) MarshalJSON() (data []byte, err error) {
	type shadow AgentAutomateInputParamsField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentAutomateInputParamsField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentResearchParams struct {
	// The research query or question to answer. Maximum 10,000 characters.
	Query string `json:"query" api:"required"`
	// Timeout in seconds for fetching web pages
	FetchTimeout param.Opt[int64] `json:"fetch_timeout,omitzero"`
	// Skip cache and force fresh research
	Nocache param.Opt[bool] `json:"nocache,omitzero"`
	// Research mode: fast (quick answers, default), balanced (standard research)
	//
	// Any of "fast", "balanced".
	Mode AgentResearchParamsMode `json:"mode,omitzero"`
	paramObj
}

func (r AgentResearchParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentResearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentResearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Research mode: fast (quick answers, default), balanced (standard research)
type AgentResearchParamsMode string

const (
	AgentResearchParamsModeFast     AgentResearchParamsMode = "fast"
	AgentResearchParamsModeBalanced AgentResearchParamsMode = "balanced"
)
