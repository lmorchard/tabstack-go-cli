// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tabstack

import (
	"context"
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
func (r *AgentService) AutomateStreaming(ctx context.Context, body AgentAutomateParams, opts ...option.RequestOption) (stream *ssestream.Stream[AutomateEvent]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	path := "automate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStreamWithSynthesizeEventData[AutomateEvent](ssestream.NewDecoder(raw), err)
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
// - `fast` - Quick answers with minimal web searches
// - `balanced` - Standard research with multiple iterations (default)
//
// **Use Cases:**
//
// - Answering complex questions with cited sources
// - Synthesizing information from multiple web sources
// - Research reports on specific topics
// - Fact-checking and verification tasks
func (r *AgentService) ResearchStreaming(ctx context.Context, body AgentResearchParams, opts ...option.RequestOption) (stream *ssestream.Stream[ResearchEvent]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	path := "research"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStreamWithSynthesizeEventData[ResearchEvent](ssestream.NewDecoder(raw), err)
}

type AutomateEvent struct {
	// Event payload data
	Data any `json:"data"`
	// The event type (e.g., start, agent:processing, complete)
	Event string `json:"event"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutomateEvent) RawJSON() string { return r.JSON.raw }
func (r *AutomateEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResearchEvent struct {
	// Event payload data
	Data any `json:"data"`
	// The event type (e.g., start, planning:start, searching:end, complete)
	Event string `json:"event"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchEvent) RawJSON() string { return r.JSON.raw }
func (r *ResearchEvent) UnmarshalJSON(data []byte) error {
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
	// The research query or question to answer
	Query string `json:"query" api:"required"`
	// Timeout in seconds for fetching web pages
	FetchTimeout param.Opt[int64] `json:"fetch_timeout,omitzero"`
	// Skip cache and force fresh research
	Nocache param.Opt[bool] `json:"nocache,omitzero"`
	// Research mode: fast (quick answers), balanced (standard research, default)
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

// Research mode: fast (quick answers), balanced (standard research, default)
type AgentResearchParamsMode string

const (
	AgentResearchParamsModeFast     AgentResearchParamsMode = "fast"
	AgentResearchParamsModeBalanced AgentResearchParamsMode = "balanced"
)
