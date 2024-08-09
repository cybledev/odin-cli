package odin_cli_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/cybledev/odin-cli/cli"
)

func TestHostsSearchCommand(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantErr  bool
		wantData bool
	}{
		{
			name:     "hosts invalid search",
			args:     []string{"hosts", "search", "--request.query='services.port:443'", "--request.limit=5"},
			wantErr:  false,
			wantData: true,
		},
		{
			name:     "hosts invalid query",
			args:     []string{"hosts", "search", "--request.query=", "--request.limit=5"},
			wantErr:  false,
			wantData: false,
		},
		{
			name:     "hosts invalid limit",
			args:     []string{"hosts", "search", "--request.query=provider:aws", "--request.limit=invalid"},
			wantErr:  true,
			wantData: false,
		},
		{
			name:     "buckets invalid search",
			args:     []string{"exposed-buckets", "search", "--request.query='provider: aws'", "--request.limit=5"},
			wantErr:  false,
			wantData: true,
		},
		{
			name:     "buckets invalid query",
			args:     []string{"exposed-buckets", "search", "--request.query=", "--request.limit=5"},
			wantErr:  false,
			wantData: false,
		},
		{
			name:     "buckets invalid limit",
			args:     []string{"exposed-bucket", "search", "--request.query=provider:aws", "--request.limit=invalid"},
			wantErr:  true,
			wantData: false,
		},
		{
			name:     "buckets invalid limit",
			args:     []string{"exposed-buckets", "search", "--request.query=provider:aws", "--request.limit=invalid"},
			wantErr:  true,
			wantData: false,
		},
		{
			name:     "certificates search",
			args:     []string{"certificate", "search"},
			wantErr:  false,
			wantData: true,
		},
		{
			name:     "v2 hosts search",
			args:     []string{"hosts", "v2-search"},
			wantErr:  false,
			wantData: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootCmd, err := cli.MakeRootCmd()
			if err != nil {
				t.Errorf("Failed to create root command: %v", err)
				return
			}

			buf := new(bytes.Buffer)
			rootCmd.SetOut(buf)
			rootCmd.SetErr(buf)
			rootCmd.SetArgs(tt.args)

			err = rootCmd.Execute()
			if (err != nil) != tt.wantErr {
				t.Errorf("hostsSearchCmd.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantData {
				var response map[string]interface{}
				err := json.Unmarshal(buf.Bytes(), &response)
				if err != nil {
					t.Errorf("Failed to parse JSON response: %v", err)
					return
				}

				if data, ok := response["data"]; !ok || data == nil {
					t.Errorf("Expected 'data' field in response, got: %v", response)
				}
			}
		})
	}
}
