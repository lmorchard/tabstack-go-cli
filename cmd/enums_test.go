package cmd

import (
	"strings"
	"testing"
)

func TestValidateEnum(t *testing.T) {
	tests := []struct {
		name         string
		flag, value  string
		valid        []string
		wantErr      bool
		wantContains string
	}{
		{
			name:  "empty value passes (means SDK default)",
			flag:  "effort",
			value: "",
			valid: validEfforts,
		},
		{
			name:  "valid value",
			flag:  "effort",
			value: "min",
			valid: validEfforts,
		},
		{
			name:  "another valid value",
			flag:  "effort",
			value: "standard",
			valid: validEfforts,
		},
		{
			name:         "invalid effort",
			flag:         "effort",
			value:        "turbo",
			valid:        validEfforts,
			wantErr:      true,
			wantContains: `invalid --effort "turbo": must be one of min, standard, max`,
		},
		{
			name:         "invalid mode",
			flag:         "mode",
			value:        "chaos",
			valid:        validResearchModes,
			wantErr:      true,
			wantContains: `invalid --mode "chaos": must be one of fast, balanced`,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateEnum(tc.flag, tc.value, tc.valid)
			switch {
			case tc.wantErr && err == nil:
				t.Fatal("want error, got nil")
			case !tc.wantErr && err != nil:
				t.Fatalf("want nil, got %v", err)
			case tc.wantContains != "" && !strings.Contains(err.Error(), tc.wantContains):
				t.Errorf("error %q should contain %q", err.Error(), tc.wantContains)
			}
		})
	}
}
