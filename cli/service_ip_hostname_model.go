// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/cybledev/odin-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for ServiceIPHostname

// register flags to command
func registerModelServiceIPHostnameFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerServiceIPHostnamePropLastUpdatedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceIPHostnamePropName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerServiceIPHostnamePropLastUpdatedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagLastUpdatedAtDescription := ``

	var flagLastUpdatedAtName string
	if cmdPrefix == "" {
		flagLastUpdatedAtName = "last_updated_at"
	} else {
		flagLastUpdatedAtName = fmt.Sprintf("%v.last_updated_at", cmdPrefix)
	}

	var flagLastUpdatedAtDefault string

	_ = cmd.PersistentFlags().String(flagLastUpdatedAtName, flagLastUpdatedAtDefault, flagLastUpdatedAtDescription)

	return nil
}

func registerServiceIPHostnamePropName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagNameDescription := ``

	var flagNameName string
	if cmdPrefix == "" {
		flagNameName = "name"
	} else {
		flagNameName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var flagNameDefault string

	_ = cmd.PersistentFlags().String(flagNameName, flagNameDefault, flagNameDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelServiceIPHostnameFlags(depth int, m *models.ServiceIPHostname, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, LastUpdatedAtAdded := retrieveServiceIPHostnamePropLastUpdatedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || LastUpdatedAtAdded

	err, NameAdded := retrieveServiceIPHostnamePropNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || NameAdded

	return nil, retAdded
}

func retrieveServiceIPHostnamePropLastUpdatedAtFlags(depth int, m *models.ServiceIPHostname, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagLastUpdatedAtName := fmt.Sprintf("%v.last_updated_at", cmdPrefix)
	if cmd.Flags().Changed(flagLastUpdatedAtName) {

		var flagLastUpdatedAtName string
		if cmdPrefix == "" {
			flagLastUpdatedAtName = "last_updated_at"
		} else {
			flagLastUpdatedAtName = fmt.Sprintf("%v.last_updated_at", cmdPrefix)
		}

		flagLastUpdatedAtValue, err := cmd.Flags().GetString(flagLastUpdatedAtName)
		if err != nil {
			return err, false
		}
		m.LastUpdatedAt = flagLastUpdatedAtValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServiceIPHostnamePropNameFlags(depth int, m *models.ServiceIPHostname, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagNameName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(flagNameName) {

		var flagNameName string
		if cmdPrefix == "" {
			flagNameName = "name"
		} else {
			flagNameName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		flagNameValue, err := cmd.Flags().GetString(flagNameName)
		if err != nil {
			return err, false
		}
		m.Name = flagNameValue

		retAdded = true
	}

	return nil, retAdded
}
