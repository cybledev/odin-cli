// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/cybledev/odin-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for DNSAPIResponse

// register flags to command
func registerModelDNSAPIResponseFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDNSAPIResponsePropData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDNSAPIResponsePropPagination(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDNSAPIResponsePropSuccess(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDNSAPIResponsePropData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: data interface{} map type is not supported by go-swagger cli yet

	return nil
}

func registerDNSAPIResponsePropPagination(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: pagination interface{} map type is not supported by go-swagger cli yet

	return nil
}

func registerDNSAPIResponsePropSuccess(depth int, cmdPrefix string, cmd *cobra.Command) error {
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
func retrieveModelDNSAPIResponseFlags(depth int, m *models.DNSAPIResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, DataAdded := retrieveDNSAPIResponsePropDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DataAdded

	err, PaginationAdded := retrieveDNSAPIResponsePropPaginationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || PaginationAdded

	err, SuccessAdded := retrieveDNSAPIResponsePropSuccessFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || SuccessAdded

	return nil, retAdded
}

func retrieveDNSAPIResponsePropDataFlags(depth int, m *models.DNSAPIResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDNSAPIResponsePropPaginationFlags(depth int, m *models.DNSAPIResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDNSAPIResponsePropSuccessFlags(depth int, m *models.DNSAPIResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
