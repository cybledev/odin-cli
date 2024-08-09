// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/cybledev/odin-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for CertificateCertCount

// register flags to command
func registerModelCertificateCertCountFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCertificateCertCountPropCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCertificateCertCountPropCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagCountDescription := ``

	var flagCountName string
	if cmdPrefix == "" {
		flagCountName = "count"
	} else {
		flagCountName = fmt.Sprintf("%v.count", cmdPrefix)
	}

	var flagCountDefault int64

	_ = cmd.PersistentFlags().Int64(flagCountName, flagCountDefault, flagCountDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCertificateCertCountFlags(depth int, m *models.CertificateCertCount, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, CountAdded := retrieveCertificateCertCountPropCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || CountAdded

	return nil, retAdded
}

func retrieveCertificateCertCountPropCountFlags(depth int, m *models.CertificateCertCount, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagCountName := fmt.Sprintf("%v.count", cmdPrefix)
	if cmd.Flags().Changed(flagCountName) {

		var flagCountName string
		if cmdPrefix == "" {
			flagCountName = "count"
		} else {
			flagCountName = fmt.Sprintf("%v.count", cmdPrefix)
		}

		flagCountValue, err := cmd.Flags().GetInt64(flagCountName)
		if err != nil {
			return err, false
		}
		m.Count = flagCountValue

		retAdded = true
	}

	return nil, retAdded
}
