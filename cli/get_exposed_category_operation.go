// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/cybledev/odin-cli/decorator"
	"github.com/cybledev/odin-cli/odin_cli/fields"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationFieldsGetExposedCategoryCmd returns a command to handle operation getExposedCategory
func makeOperationFieldsGetExposedCategoryCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Args: cobra.NoArgs,
		Use:   "exposed",
		Short: `Returns the fields data for exposed files and buckets`,
		RunE:  decorator.RunEColorWrapper(runOperationFieldsGetExposedCategory),
	}

	if err := registerOperationFieldsGetExposedCategoryParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationFieldsGetExposedCategory uses cmd flags to call endpoint api
func runOperationFieldsGetExposedCategory(cmd *cobra.Command, args []string) (string, error) {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return "", err
	}
	// retrieve flag values from cmd and fill params
	params := fields.NewGetExposedCategoryParams()
	if err, _ = retrieveOperationFieldsGetExposedCategoryCategoryFlag(params, "", cmd); err != nil {
		return "", err
	}
	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return "", nil
	}
	// make request and then print result
	msgStr, err := parseOperationFieldsGetExposedCategoryResult(appCli.Fields.GetExposedCategory(params))
	if err != nil {
		return "", err
	}

	if !debug {
		return msgStr, nil
	}
	return "", err
}

// registerOperationFieldsGetExposedCategoryParamFlags registers all flags needed to fill params
func registerOperationFieldsGetExposedCategoryParamFlags(cmd *cobra.Command) error {
	if err := registerOperationFieldsGetExposedCategoryCategoryParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationFieldsGetExposedCategoryCategoryParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	flagCategoryDescription := `Required. get the category`

	var flagCategoryName string
	if cmdPrefix == "" {
		flagCategoryName = "category"
	} else {
		flagCategoryName = fmt.Sprintf("%v.category", cmdPrefix)
	}

	var flagCategoryDefault string

	_ = cmd.PersistentFlags().String(flagCategoryName, flagCategoryDefault, flagCategoryDescription)

	return nil
}

func retrieveOperationFieldsGetExposedCategoryCategoryFlag(m *fields.GetExposedCategoryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("category") {

		var flagCategoryName string
		if cmdPrefix == "" {
			flagCategoryName = "category"
		} else {
			flagCategoryName = fmt.Sprintf("%v.category", cmdPrefix)
		}

		flagCategoryValue, err := cmd.Flags().GetString(flagCategoryName)
		if err != nil {
			return err, false
		}
		m.Category = flagCategoryValue

	}

	return nil, retAdded
}

// parseOperationFieldsGetExposedCategoryResult parses request result and return the string content
func parseOperationFieldsGetExposedCategoryResult(resp0 *fields.GetExposedCategoryOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*fields.GetExposedCategoryOK)
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
		resp1, ok := iResp1.(*fields.GetExposedCategoryBadRequest)
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
		resp2, ok := iResp2.(*fields.GetExposedCategoryRequestTimeout)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp3 interface{} = respErr
		resp3, ok := iResp3.(*fields.GetExposedCategoryInternalServerError)
		if ok {
			if !swag.IsZero(resp3) && !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
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
func registerModelGetExposedCategoryOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	// register embedded models.FieldsAPIResponse flags

	if err := registerModelFieldsAPIResponseFlags(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	// register anonymous fields for getExposedCategoryOKBodyAO1

	if err := registerGetExposedCategoryOKBodyPropAnonGetExposedCategoryOKBodyAO1Data(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

// inline definition name getExposedCategoryOKBodyAO1, type

func registerGetExposedCategoryOKBodyPropAnonGetExposedCategoryOKBodyAO1Data(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Data []*models.FieldsField array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelGetExposedCategoryOKBodyFlags(depth int, m *fields.GetExposedCategoryOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	// retrieve model models.FieldsAPIResponse
	err, GetExposedCategoryOKBodyAO0Added := retrieveModelFieldsAPIResponseFlags(depth, &m.FieldsAPIResponse, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || GetExposedCategoryOKBodyAO0Added

	// retrieve allOf getExposedCategoryOKBodyAO1 fields

	err, DataAdded := retrieveGetExposedCategoryOKBodyPropAnonGetExposedCategoryOKBodyAO1DataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DataAdded

	return nil, retAdded
}

// define retrieve functions for fields for inline definition name getExposedCategoryOKBodyAO1

func retrieveGetExposedCategoryOKBodyPropAnonGetExposedCategoryOKBodyAO1DataFlags(depth int, m *fields.GetExposedCategoryOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagDataName := fmt.Sprintf("%v.Data", cmdPrefix)
	if cmd.Flags().Changed(flagDataName) {
		// warning: Data array type []*models.FieldsField is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
