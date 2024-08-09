// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/cybledev/odin-cli/models"
	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// Schema cli for ExposedFileAPIResponse

// register flags to command
func registerModelExposedFileAPIResponseFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerExposedFileAPIResponsePropData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExposedFileAPIResponsePropPagination(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExposedFileAPIResponsePropSuccess(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerExposedFileAPIResponsePropData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: data []*ExposedFile array type is not supported by go-swagger cli yet

	return nil
}

func registerExposedFileAPIResponsePropPagination(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var flagPaginationName string
	if cmdPrefix == "" {
		flagPaginationName = "pagination"
	} else {
		flagPaginationName = fmt.Sprintf("%v.pagination", cmdPrefix)
	}

	if err := registerModelExposedSearchPaginationFlags(depth+1, flagPaginationName, cmd); err != nil {
		return err
	}

	return nil
}

func registerExposedFileAPIResponsePropSuccess(depth int, cmdPrefix string, cmd *cobra.Command) error {
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
func retrieveModelExposedFileAPIResponseFlags(depth int, m *models.ExposedFileAPIResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, DataAdded := retrieveExposedFileAPIResponsePropDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DataAdded

	err, PaginationAdded := retrieveExposedFileAPIResponsePropPaginationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || PaginationAdded

	err, SuccessAdded := retrieveExposedFileAPIResponsePropSuccessFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || SuccessAdded

	return nil, retAdded
}

func retrieveExposedFileAPIResponsePropDataFlags(depth int, m *models.ExposedFileAPIResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagDataName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(flagDataName) {
		// warning: data array type []*ExposedFile is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveExposedFileAPIResponsePropPaginationFlags(depth int, m *models.ExposedFileAPIResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagPaginationName := fmt.Sprintf("%v.pagination", cmdPrefix)
	if cmd.Flags().Changed(flagPaginationName) {
		// info: complex object pagination ExposedSearchPagination is retrieved outside this Changed() block
	}
	flagPaginationValue := m.Pagination
	if swag.IsZero(flagPaginationValue) {
		flagPaginationValue = &models.ExposedSearchPagination{}
	}

	err, PaginationAdded := retrieveModelExposedSearchPaginationFlags(depth+1, flagPaginationValue, flagPaginationName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || PaginationAdded
	if PaginationAdded {
		m.Pagination = flagPaginationValue
	}

	return nil, retAdded
}

func retrieveExposedFileAPIResponsePropSuccessFlags(depth int, m *models.ExposedFileAPIResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
