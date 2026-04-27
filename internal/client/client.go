// Package client constructs a Tabstack SDK client from CLI configuration.
package client

import (
	"fmt"
	"os"

	"github.com/lmorchard/tabstack-go-cli/internal/config"
	tabstack "github.com/stainless-sdks/tabstack-go"
	"github.com/stainless-sdks/tabstack-go/option"
)

// New constructs a Tabstack SDK client. Values from cfg are layered on top of
// the SDK's environment defaults (TABSTACK_API_KEY, TABSTACK_BASE_URL), so a
// non-empty cfg.APIKey or cfg.BaseURL overrides whatever was in the env.
//
// Returns an error if no API key resolves from cfg or env — the SDK itself
// would happily make unauthenticated requests and fail with a 401 mid-stream,
// which is a worse experience than failing fast here.
func New(cfg *config.Config) (*tabstack.Client, error) {
	if cfg.APIKey == "" && os.Getenv("TABSTACK_API_KEY") == "" {
		return nil, fmt.Errorf("no Tabstack API key set: provide --api-key, set api_key in config, or export TABSTACK_API_KEY")
	}

	var opts []option.RequestOption
	if cfg.APIKey != "" {
		opts = append(opts, option.WithAPIKey(cfg.APIKey))
	}
	if cfg.BaseURL != "" {
		opts = append(opts, option.WithBaseURL(cfg.BaseURL))
	}

	c := tabstack.NewClient(opts...)
	return &c, nil
}
