// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/cybledev/odin-cli/decorator"
	"github.com/cybledev/odin-cli/models"
	"github.com/cybledev/odin-cli/odin_cli/exposed_files"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationExposedFilesPostExposedFilesSearchCmd returns a command to handle operation postExposedFilesSearch
func makeOperationExposedFilesPostExposedFilesSearchCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "search",
		Example: `odin-cli exposed-files search --query '{"bucket":bucket:\"lit-link-prd.appspot.com\"}'`,
		Short: `returns the matching records according to provided filters`,
		RunE:   decorator.RunEColorWrapper(runOperationExposedFilesPostExposedFilesSearch),
	}

	if err := registerOperationExposedFilesPostExposedFilesSearchParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationExposedFilesPostExposedFilesSearch uses cmd flags to call endpoint api
func runOperationExposedFilesPostExposedFilesSearch(cmd *cobra.Command, args []string) (val string, err error) {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return
	}
	// retrieve flag values from cmd and fill params
	params := exposed_files.NewPostExposedFilesSearchParams()
	if err, _ = retrieveOperationExposedFilesPostExposedFilesSearchQueryFlag(params, "", cmd); err != nil {
		return
	}
	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return
	}
	// make request and then print result
	msgStr, err := parseOperationExposedFilesPostExposedFilesSearchResult(appCli.ExposedFiles.PostExposedFilesSearch(params, nil))
	if err != nil {
		return
	}

	if !debug {
		val = msgStr
	}
	return
}

// registerOperationExposedFilesPostExposedFilesSearchParamFlags registers all flags needed to fill params
func registerOperationExposedFilesPostExposedFilesSearchParamFlags(cmd *cobra.Command) error {
	if err := registerOperationExposedFilesPostExposedFilesSearchQueryParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationExposedFilesPostExposedFilesSearchQueryParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var flagQueryName string
	if cmdPrefix == "" {
		flagQueryName = "query"
	} else {
		flagQueryName = fmt.Sprintf("%v.query", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(flagQueryName, "", `Optional json string for [query]. Search Query`)

	// add flags for body
	if err := registerModelExposedSearchRequestFlags(0, "request", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationExposedFilesPostExposedFilesSearchQueryFlag(m *exposed_files.PostExposedFilesSearchParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("query") {
		// Read query string from cmd and unmarshal
		flagQueryValueStr, err := cmd.Flags().GetString("query")
		if err != nil {
			return err, false
		}

		flagQueryValue := models.ExposedSearchRequest{}
		if err := json.Unmarshal([]byte(flagQueryValueStr), &flagQueryValue); err != nil {
			return fmt.Errorf("cannot unmarshal query string in models.ExposedSearchRequest: %v", err), false
		}
		m.Query = &flagQueryValue
	}
	flagQueryModel := m.Query
	if swag.IsZero(flagQueryModel) {
		flagQueryModel = &models.ExposedSearchRequest{}
	}
	err, added := retrieveModelExposedSearchRequestFlags(0, flagQueryModel, "request", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Query = flagQueryModel
	}

	if dryRun && debug {
		flagQueryValueDebugBytes, err := json.Marshal(m.Query)
		if err != nil {
			return err, false
		}
		logDebugf("Query dry-run payload: %v", string(flagQueryValueDebugBytes))
	}

	retAdded = retAdded || added

	return nil, retAdded
}

// parseOperationExposedFilesPostExposedFilesSearchResult parses request result and return the string content
func parseOperationExposedFilesPostExposedFilesSearchResult(resp0 *exposed_files.PostExposedFilesSearchOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*exposed_files.PostExposedFilesSearchOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.MarshalIndent(resp0.Payload, "", "  ")
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*exposed_files.PostExposedFilesSearchBadRequest)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.MarshalIndent(resp1.Payload, "", "  ")
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*exposed_files.PostExposedFilesSearchInternalServerError)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.MarshalIndent(resp2.Payload, "", "  ")
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.MarshalIndent(resp0.Payload, "", "  ")
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
