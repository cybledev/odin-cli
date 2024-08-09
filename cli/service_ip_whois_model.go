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

// Schema cli for ServiceIPWhois

// register flags to command
func registerModelServiceIPWhoisFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerServiceIPWhoisPropEncoding(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceIPWhoisPropDescr(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceIPWhoisPropNetwork(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceIPWhoisPropOrganization(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceIPWhoisPropRaw(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerServiceIPWhoisPropEncoding(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var flagEncodingName string
	if cmdPrefix == "" {
		flagEncodingName = "_encoding"
	} else {
		flagEncodingName = fmt.Sprintf("%v._encoding", cmdPrefix)
	}

	if err := registerModelServiceEncodingFlags(depth+1, flagEncodingName, cmd); err != nil {
		return err
	}

	return nil
}

func registerServiceIPWhoisPropDescr(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagDescrDescription := ``

	var flagDescrName string
	if cmdPrefix == "" {
		flagDescrName = "descr"
	} else {
		flagDescrName = fmt.Sprintf("%v.descr", cmdPrefix)
	}

	var flagDescrDefault string

	_ = cmd.PersistentFlags().String(flagDescrName, flagDescrDefault, flagDescrDescription)

	return nil
}

func registerServiceIPWhoisPropNetwork(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagNetworkDescription := ``

	var flagNetworkName string
	if cmdPrefix == "" {
		flagNetworkName = "network"
	} else {
		flagNetworkName = fmt.Sprintf("%v.network", cmdPrefix)
	}

	var flagNetworkDefault string

	_ = cmd.PersistentFlags().String(flagNetworkName, flagNetworkDefault, flagNetworkDescription)

	return nil
}

func registerServiceIPWhoisPropOrganization(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagOrganizationDescription := ``

	var flagOrganizationName string
	if cmdPrefix == "" {
		flagOrganizationName = "organization"
	} else {
		flagOrganizationName = fmt.Sprintf("%v.organization", cmdPrefix)
	}

	var flagOrganizationDefault string

	_ = cmd.PersistentFlags().String(flagOrganizationName, flagOrganizationDefault, flagOrganizationDescription)

	return nil
}

func registerServiceIPWhoisPropRaw(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagRawDescription := ``

	var flagRawName string
	if cmdPrefix == "" {
		flagRawName = "raw"
	} else {
		flagRawName = fmt.Sprintf("%v.raw", cmdPrefix)
	}

	var flagRawDefault string

	_ = cmd.PersistentFlags().String(flagRawName, flagRawDefault, flagRawDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelServiceIPWhoisFlags(depth int, m *models.ServiceIPWhois, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, EncodingAdded := retrieveServiceIPWhoisPropEncodingFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || EncodingAdded

	err, DescrAdded := retrieveServiceIPWhoisPropDescrFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DescrAdded

	err, NetworkAdded := retrieveServiceIPWhoisPropNetworkFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || NetworkAdded

	err, OrganizationAdded := retrieveServiceIPWhoisPropOrganizationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || OrganizationAdded

	err, RawAdded := retrieveServiceIPWhoisPropRawFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || RawAdded

	return nil, retAdded
}

func retrieveServiceIPWhoisPropEncodingFlags(depth int, m *models.ServiceIPWhois, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagEncodingName := fmt.Sprintf("%v._encoding", cmdPrefix)
	if cmd.Flags().Changed(flagEncodingName) {
		// info: complex object _encoding ServiceEncoding is retrieved outside this Changed() block
	}
	flagEncodingValue := m.Encoding
	if swag.IsZero(flagEncodingValue) {
		flagEncodingValue = &models.ServiceEncoding{}
	}

	err, EncodingAdded := retrieveModelServiceEncodingFlags(depth+1, flagEncodingValue, flagEncodingName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || EncodingAdded
	if EncodingAdded {
		m.Encoding = flagEncodingValue
	}

	return nil, retAdded
}

func retrieveServiceIPWhoisPropDescrFlags(depth int, m *models.ServiceIPWhois, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagDescrName := fmt.Sprintf("%v.descr", cmdPrefix)
	if cmd.Flags().Changed(flagDescrName) {

		var flagDescrName string
		if cmdPrefix == "" {
			flagDescrName = "descr"
		} else {
			flagDescrName = fmt.Sprintf("%v.descr", cmdPrefix)
		}

		flagDescrValue, err := cmd.Flags().GetString(flagDescrName)
		if err != nil {
			return err, false
		}
		m.Descr = flagDescrValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServiceIPWhoisPropNetworkFlags(depth int, m *models.ServiceIPWhois, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagNetworkName := fmt.Sprintf("%v.network", cmdPrefix)
	if cmd.Flags().Changed(flagNetworkName) {

		var flagNetworkName string
		if cmdPrefix == "" {
			flagNetworkName = "network"
		} else {
			flagNetworkName = fmt.Sprintf("%v.network", cmdPrefix)
		}

		flagNetworkValue, err := cmd.Flags().GetString(flagNetworkName)
		if err != nil {
			return err, false
		}
		m.Network = flagNetworkValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServiceIPWhoisPropOrganizationFlags(depth int, m *models.ServiceIPWhois, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagOrganizationName := fmt.Sprintf("%v.organization", cmdPrefix)
	if cmd.Flags().Changed(flagOrganizationName) {

		var flagOrganizationName string
		if cmdPrefix == "" {
			flagOrganizationName = "organization"
		} else {
			flagOrganizationName = fmt.Sprintf("%v.organization", cmdPrefix)
		}

		flagOrganizationValue, err := cmd.Flags().GetString(flagOrganizationName)
		if err != nil {
			return err, false
		}
		m.Organization = flagOrganizationValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServiceIPWhoisPropRawFlags(depth int, m *models.ServiceIPWhois, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagRawName := fmt.Sprintf("%v.raw", cmdPrefix)
	if cmd.Flags().Changed(flagRawName) {

		var flagRawName string
		if cmdPrefix == "" {
			flagRawName = "raw"
		} else {
			flagRawName = fmt.Sprintf("%v.raw", cmdPrefix)
		}

		flagRawValue, err := cmd.Flags().GetString(flagRawName)
		if err != nil {
			return err, false
		}
		m.Raw = flagRawValue

		retAdded = true
	}

	return nil, retAdded
}
