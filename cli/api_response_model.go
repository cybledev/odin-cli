// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/cybledev/odin-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for APIResponse

// register flags to command
func registerModelAPIResponseFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerAPIResponsePropData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAPIResponsePropMessage(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAPIResponsePropPagination(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAPIResponsePropSuccess(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerAPIResponsePropData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: data interface{} map type is not supported by go-swagger cli yet

	return nil
}

func registerAPIResponsePropMessage(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagMessageDescription := ``

	var flagMessageName string
	if cmdPrefix == "" {
		flagMessageName = "message"
	} else {
		flagMessageName = fmt.Sprintf("%v.message", cmdPrefix)
	}

	var flagMessageDefault string

	_ = cmd.PersistentFlags().String(flagMessageName, flagMessageDefault, flagMessageDescription)

	return nil
}

func registerAPIResponsePropPagination(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: pagination interface{} map type is not supported by go-swagger cli yet

	return nil
}

func registerAPIResponsePropSuccess(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagSuccessDescription := ``

	var flagSuccessName string
	if cmdPrefix == "" {
		flagSuccessName = "success"
	} else {
		flagSuccessName = fmt.Sprintf("%v.success", cmdPrefix)
	}

	var flagSuccessDefault bool

	_ = cmd.PersistentFlags().Bool(flagSuccessName, flagSuccessDefault, flagSuccessDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelAPIResponseFlags(depth int, m *models.APIResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, DataAdded := retrieveAPIResponsePropDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DataAdded

	err, MessageAdded := retrieveAPIResponsePropMessageFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || MessageAdded

	err, PaginationAdded := retrieveAPIResponsePropPaginationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || PaginationAdded

	err, SuccessAdded := retrieveAPIResponsePropSuccessFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || SuccessAdded

	return nil, retAdded
}

func retrieveAPIResponsePropDataFlags(depth int, m *models.APIResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagDataName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(flagDataName) {
		// warning: data map type interface{} is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveAPIResponsePropMessageFlags(depth int, m *models.APIResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagMessageName := fmt.Sprintf("%v.message", cmdPrefix)
	if cmd.Flags().Changed(flagMessageName) {

		var flagMessageName string
		if cmdPrefix == "" {
			flagMessageName = "message"
		} else {
			flagMessageName = fmt.Sprintf("%v.message", cmdPrefix)
		}

		flagMessageValue, err := cmd.Flags().GetString(flagMessageName)
		if err != nil {
			return err, false
		}
		m.Message = flagMessageValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAPIResponsePropPaginationFlags(depth int, m *models.APIResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagPaginationName := fmt.Sprintf("%v.pagination", cmdPrefix)
	if cmd.Flags().Changed(flagPaginationName) {
		// warning: pagination map type interface{} is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveAPIResponsePropSuccessFlags(depth int, m *models.APIResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagSuccessName := fmt.Sprintf("%v.success", cmdPrefix)
	if cmd.Flags().Changed(flagSuccessName) {

		var flagSuccessName string
		if cmdPrefix == "" {
			flagSuccessName = "success"
		} else {
			flagSuccessName = fmt.Sprintf("%v.success", cmdPrefix)
		}

		flagSuccessValue, err := cmd.Flags().GetBool(flagSuccessName)
		if err != nil {
			return err, false
		}
		m.Success = flagSuccessValue

		retAdded = true
	}

	return nil, retAdded
}
