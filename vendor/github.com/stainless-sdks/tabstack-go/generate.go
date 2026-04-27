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
)

// GenerateService contains methods and other services that help with interacting
// with the tabstack API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGenerateService] method instead.
type GenerateService struct {
	options []option.RequestOption
}

// NewGenerateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGenerateService(opts ...option.RequestOption) (r GenerateService) {
	r = GenerateService{}
	r.options = opts
	return
}

// Fetches URL content, extracts data, and transforms it using AI based on custom
// instructions. Use this to generate new content, summaries, or restructured data.
func (r *GenerateService) Json(ctx context.Context, body GenerateJsonParams, opts ...option.RequestOption) (res *GenerateJsonResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "generate/json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type GenerateJsonResponse map[string]any

type GenerateJsonParams struct {
	// Instructions describing how to transform the data. Maximum 20,000 characters.
	Instructions string `json:"instructions" api:"required"`
	// JSON schema defining the structure of the transformed output
	JsonSchema any `json:"json_schema,omitzero" api:"required"`
	// URL to fetch content from
	URL string `json:"url" api:"required" format:"uri"`
	// Bypass cache and force fresh data retrieval
	Nocache param.Opt[bool] `json:"nocache,omitzero"`
	// Fetch effort level controlling speed vs. capability tradeoff. "min": fastest, no
	// fallback (1-5s). "standard": balanced with enhanced reliability (default,
	// 3-15s). "max": full browser rendering for JS-heavy sites (15-60s).
	//
	// Any of "min", "standard", "max".
	Effort GenerateJsonParamsEffort `json:"effort,omitzero"`
	// Optional geotargeting parameters for proxy requests
	GeoTarget GenerateJsonParamsGeoTarget `json:"geo_target,omitzero"`
	paramObj
}

func (r GenerateJsonParams) MarshalJSON() (data []byte, err error) {
	type shadow GenerateJsonParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GenerateJsonParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Fetch effort level controlling speed vs. capability tradeoff. "min": fastest, no
// fallback (1-5s). "standard": balanced with enhanced reliability (default,
// 3-15s). "max": full browser rendering for JS-heavy sites (15-60s).
type GenerateJsonParamsEffort string

const (
	GenerateJsonParamsEffortMin      GenerateJsonParamsEffort = "min"
	GenerateJsonParamsEffortStandard GenerateJsonParamsEffort = "standard"
	GenerateJsonParamsEffortMax      GenerateJsonParamsEffort = "max"
)

// Optional geotargeting parameters for proxy requests
type GenerateJsonParamsGeoTarget struct {
	// Country code using ISO 3166-1 alpha-2 standard (2 letters, e.g., "US", "GB",
	// "JP"). See: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
	Country param.Opt[string] `json:"country,omitzero"`
	paramObj
}

func (r GenerateJsonParamsGeoTarget) MarshalJSON() (data []byte, err error) {
	type shadow GenerateJsonParamsGeoTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GenerateJsonParamsGeoTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
