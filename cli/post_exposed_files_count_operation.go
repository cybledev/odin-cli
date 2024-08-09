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

// makeOperationExposedFilesPostExposedFilesCountCmd returns a command to handle operation postExposedFilesCount
func makeOperationExposedFilesPostExposedFilesCountCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "count",
		Short: `Returns overall count of exposed bucket files according to provided filters`,
		RunE:   decorator.RunEColorWrapper(runOperationExposedFilesPostExposedFilesCount),
	}

	if err := registerOperationExposedFilesPostExposedFilesCountParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationExposedFilesPostExposedFilesCount uses cmd flags to call endpoint api
func runOperationExposedFilesPostExposedFilesCount(cmd *cobra.Command, args []string) (string, error) {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return "", err
	}
	// retrieve flag values from cmd and fill params
	params := exposed_files.NewPostExposedFilesCountParams()
	if err, _ = retrieveOperationExposedFilesPostExposedFilesCountQueryFlag(params, "", cmd); err != nil {
		return "", err
	}
	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return "", nil
	}
	// make request and then print result
	msgStr, err := parseOperationExposedFilesPostExposedFilesCountResult(appCli.ExposedFiles.PostExposedFilesCount(params, nil))
	if err != nil {
		return "", err
	}

	if !debug {
		return msgStr, nil
	}

	return "", nil
}

// registerOperationExposedFilesPostExposedFilesCountParamFlags registers all flags needed to fill params
func registerOperationExposedFilesPostExposedFilesCountParamFlags(cmd *cobra.Command) error {
	if err := registerOperationExposedFilesPostExposedFilesCountQueryParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationExposedFilesPostExposedFilesCountQueryParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var flagQueryName string
	if cmdPrefix == "" {
		flagQueryName = "query"
	} else {
		flagQueryName = fmt.Sprintf("%v.query", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(flagQueryName, "", `Optional json string for [query]. Count Request`)

	// add flags for body
	if err := registerModelExposedCountRequestFlags(0, "request", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationExposedFilesPostExposedFilesCountQueryFlag(m *exposed_files.PostExposedFilesCountParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("query") {
		// Read query string from cmd and unmarshal
		flagQueryValueStr, err := cmd.Flags().GetString("query")
		if err != nil {
			return err, false
		}

		flagQueryValue := models.ExposedCountRequest{}
		if err := json.Unmarshal([]byte(flagQueryValueStr), &flagQueryValue); err != nil {
			return fmt.Errorf("cannot unmarshal query string in models.ExposedCountRequest: %v", err), false
		}
		m.Query = &flagQueryValue
	}
	flagQueryModel := m.Query
	if swag.IsZero(flagQueryModel) {
		flagQueryModel = &models.ExposedCountRequest{}
	}
	err, added := retrieveModelExposedCountRequestFlags(0, flagQueryModel, "request", cmd)
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

// parseOperationExposedFilesPostExposedFilesCountResult parses request result and return the string content
func parseOperationExposedFilesPostExposedFilesCountResult(resp0 *exposed_files.PostExposedFilesCountOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*exposed_files.PostExposedFilesCountOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*exposed_files.PostExposedFilesCountBadRequest)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*exposed_files.PostExposedFilesCountInternalServerError)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}

// register flags to command
func registerModelPostExposedFilesCountOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	// register embedded models.ExposedAPIResponse flags

	if err := registerModelExposedAPIResponseFlags(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	// register anonymous fields for postExposedFilesCountOKBodyAO1

	if err := registerPostExposedFilesCountOKBodyPropAnonPostExposedFilesCountOKBodyAO1Data(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

// inline definition name postExposedFilesCountOKBodyAO1, type

func registerPostExposedFilesCountOKBodyPropAnonPostExposedFilesCountOKBodyAO1Data(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var flagDataName string
	if cmdPrefix == "" {
		flagDataName = "data"
	} else {
		flagDataName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelExposedSearchCountFlags(depth+1, flagDataName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPostExposedFilesCountOKBodyFlags(depth int, m *exposed_files.PostExposedFilesCountOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	// retrieve model models.ExposedAPIResponse
	err, PostExposedFilesCountOKBodyAO0Added := retrieveModelExposedAPIResponseFlags(depth, &m.ExposedAPIResponse, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || PostExposedFilesCountOKBodyAO0Added

	// retrieve allOf postExposedFilesCountOKBodyAO1 fields

	err, DataAdded := retrievePostExposedFilesCountOKBodyPropAnonPostExposedFilesCountOKBodyAO1DataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DataAdded

	return nil, retAdded
}

// define retrieve functions for fields for inline definition name postExposedFilesCountOKBodyAO1

func retrievePostExposedFilesCountOKBodyPropAnonPostExposedFilesCountOKBodyAO1DataFlags(depth int, m *exposed_files.PostExposedFilesCountOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagDataName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(flagDataName) {
		// info: complex object data models.ExposedSearchCount is retrieved outside this Changed() block
	}
	flagDataValue := m.Data
	if swag.IsZero(flagDataValue) {
		flagDataValue = &models.ExposedSearchCount{}
	}

	err, DataAdded := retrieveModelExposedSearchCountFlags(depth+1, flagDataValue, flagDataName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DataAdded
	if DataAdded {
		m.Data = flagDataValue
	}

	return nil, retAdded
}
