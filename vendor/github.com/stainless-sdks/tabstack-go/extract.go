// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tabstack

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/tabstack-go/internal/apijson"
	"github.com/stainless-sdks/tabstack-go/internal/requestconfig"
	"github.com/stainless-sdks/tabstack-go/option"
	"github.com/stainless-sdks/tabstack-go/packages/param"
	"github.com/stainless-sdks/tabstack-go/packages/respjson"
)

// ExtractService contains methods and other services that help with interacting
// with the tabstack API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtractService] method instead.
type ExtractService struct {
	options []option.RequestOption
}

// NewExtractService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewExtractService(opts ...option.RequestOption) (r ExtractService) {
	r = ExtractService{}
	r.options = opts
	return
}

// Fetches a URL and extracts structured data according to a provided JSON schema
func (r *ExtractService) Json(ctx context.Context, body ExtractJsonParams, opts ...option.RequestOption) (res *ExtractJsonResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "extract/json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetches a URL and converts its HTML content to clean Markdown format with
// optional metadata extraction
func (r *ExtractService) Markdown(ctx context.Context, body ExtractMarkdownParams, opts ...option.RequestOption) (res *ExtractMarkdownResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "extract/markdown"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ExtractJsonResponse map[string]any

type ExtractMarkdownResponse struct {
	// The markdown content (includes metadata as YAML frontmatter by default)
	Content string `json:"content" api:"required"`
	// The URL that was converted to markdown
	URL string `json:"url" api:"required" format:"uri"`
	// Extracted metadata from the page (only included when metadata parameter is true)
	Metadata ExtractMarkdownResponseMetadata `json:"metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		URL         respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractMarkdownResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractMarkdownResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Extracted metadata from the page (only included when metadata parameter is true)
type ExtractMarkdownResponseMetadata struct {
	// Author information from HTML metadata
	Author string `json:"author"`
	// Document creation date (ISO 8601)
	CreatedAt string `json:"created_at"`
	// Creator application (e.g., "Microsoft Word")
	Creator string `json:"creator"`
	// Page description from Open Graph or HTML
	Description string `json:"description"`
	// Featured image URL from Open Graph
	Image string `json:"image" format:"uri"`
	// PDF keywords as array
	Keywords []string `json:"keywords"`
	// Document modification date (ISO 8601)
	ModifiedAt string `json:"modified_at"`
	// Number of pages (PDF documents)
	PageCount int64 `json:"page_count"`
	// PDF version (e.g., "1.5")
	PdfVersion string `json:"pdf_version"`
	// PDF producer software (e.g., "Adobe PDF Library")
	Producer string `json:"producer"`
	// Publisher information from Open Graph
	Publisher string `json:"publisher"`
	// Site name from Open Graph
	SiteName string `json:"site_name"`
	// PDF-specific metadata fields (populated for PDF documents) PDF subject or
	// summary
	Subject string `json:"subject"`
	// Page title from Open Graph or HTML
	Title string `json:"title"`
	// Content type from Open Graph (e.g., article, website)
	Type string `json:"type"`
	// Canonical URL from Open Graph
	URL string `json:"url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Author      respjson.Field
		CreatedAt   respjson.Field
		Creator     respjson.Field
		Description respjson.Field
		Image       respjson.Field
		Keywords    respjson.Field
		ModifiedAt  respjson.Field
		PageCount   respjson.Field
		PdfVersion  respjson.Field
		Producer    respjson.Field
		Publisher   respjson.Field
		SiteName    respjson.Field
		Subject     respjson.Field
		Title       respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractMarkdownResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *ExtractMarkdownResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractJsonParams struct {
	// JSON schema definition that describes the structure of data to extract.
	JsonSchema any `json:"json_schema,omitzero" api:"required"`
	// URL to fetch and extract data from
	URL string `json:"url" api:"required" format:"uri"`
	// Bypass cache and force fresh data retrieval
	Nocache param.Opt[bool] `json:"nocache,omitzero"`
	// Fetch effort level controlling speed vs. capability tradeoff. "min": fastest, no
	// fallback (1-5s). "standard": balanced with enhanced reliability (default,
	// 3-15s). "max": full browser rendering for JS-heavy sites (15-60s).
	//
	// Any of "min", "standard", "max".
	Effort ExtractJsonParamsEffort `json:"effort,omitzero"`
	// Optional geotargeting parameters for proxy requests
	GeoTarget ExtractJsonParamsGeoTarget `json:"geo_target,omitzero"`
	paramObj
}

func (r ExtractJsonParams) MarshalJSON() (data []byte, err error) {
	type shadow ExtractJsonParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractJsonParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Fetch effort level controlling speed vs. capability tradeoff. "min": fastest, no
// fallback (1-5s). "standard": balanced with enhanced reliability (default,
// 3-15s). "max": full browser rendering for JS-heavy sites (15-60s).
type ExtractJsonParamsEffort string

const (
	ExtractJsonParamsEffortMin      ExtractJsonParamsEffort = "min"
	ExtractJsonParamsEffortStandard ExtractJsonParamsEffort = "standard"
	ExtractJsonParamsEffortMax      ExtractJsonParamsEffort = "max"
)

// Optional geotargeting parameters for proxy requests
type ExtractJsonParamsGeoTarget struct {
	// Country code using ISO 3166-1 alpha-2 standard (2 letters, e.g., "US", "GB",
	// "JP"). See: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
	Country param.Opt[string] `json:"country,omitzero"`
	paramObj
}

func (r ExtractJsonParamsGeoTarget) MarshalJSON() (data []byte, err error) {
	type shadow ExtractJsonParamsGeoTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractJsonParamsGeoTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractMarkdownParams struct {
	// URL to fetch and convert to markdown
	URL string `json:"url" api:"required" format:"uri"`
	// Include extracted metadata (Open Graph and HTML metadata) as a separate field in
	// the response
	Metadata param.Opt[bool] `json:"metadata,omitzero"`
	// Bypass cache and force fresh data retrieval
	Nocache param.Opt[bool] `json:"nocache,omitzero"`
	// Fetch effort level controlling speed vs. capability tradeoff. "min": fastest, no
	// fallback (1-5s). "standard": balanced with enhanced reliability (default,
	// 3-15s). "max": full browser rendering for JS-heavy sites (15-60s).
	//
	// Any of "min", "standard", "max".
	Effort ExtractMarkdownParamsEffort `json:"effort,omitzero"`
	// Optional geotargeting parameters for proxy requests
	GeoTarget ExtractMarkdownParamsGeoTarget `json:"geo_target,omitzero"`
	paramObj
}

func (r ExtractMarkdownParams) MarshalJSON() (data []byte, err error) {
	type shadow ExtractMarkdownParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractMarkdownParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Fetch effort level controlling speed vs. capability tradeoff. "min": fastest, no
// fallback (1-5s). "standard": balanced with enhanced reliability (default,
// 3-15s). "max": full browser rendering for JS-heavy sites (15-60s).
type ExtractMarkdownParamsEffort string

const (
	ExtractMarkdownParamsEffortMin      ExtractMarkdownParamsEffort = "min"
	ExtractMarkdownParamsEffortStandard ExtractMarkdownParamsEffort = "standard"
	ExtractMarkdownParamsEffortMax      ExtractMarkdownParamsEffort = "max"
)

// Optional geotargeting parameters for proxy requests
type ExtractMarkdownParamsGeoTarget struct {
	// Country code using ISO 3166-1 alpha-2 standard (2 letters, e.g., "US", "GB",
	// "JP"). See: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
	Country param.Opt[string] `json:"country,omitzero"`
	paramObj
}

func (r ExtractMarkdownParamsGeoTarget) MarshalJSON() (data []byte, err error) {
	type shadow ExtractMarkdownParamsGeoTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractMarkdownParamsGeoTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
