// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/cybledev/odin-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for IPCVE

// register flags to command
func registerModelIPCVEFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerIPCVEPropExploit(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIPCVEPropID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIPCVEPropReferences(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIPCVEPropScore(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIPCVEPropServices(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIPCVEPropSeverity(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIPCVEPropSummary(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIPCVEPropVectorString(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIPCVEPropWeakness(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerIPCVEPropExploit(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: exploit []*IPExploitDetails array type is not supported by go-swagger cli yet

	return nil
}

func registerIPCVEPropID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagIDDescription := ``

	var flagIDName string
	if cmdPrefix == "" {
		flagIDName = "id"
	} else {
		flagIDName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var flagIDDefault string

	_ = cmd.PersistentFlags().String(flagIDName, flagIDDefault, flagIDDescription)

	return nil
}

func registerIPCVEPropReferences(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: references []string array type is not supported by go-swagger cli yet

	return nil
}

func registerIPCVEPropScore(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagScoreDescription := ``

	var flagScoreName string
	if cmdPrefix == "" {
		flagScoreName = "score"
	} else {
		flagScoreName = fmt.Sprintf("%v.score", cmdPrefix)
	}

	var flagScoreDefault float64

	_ = cmd.PersistentFlags().Float64(flagScoreName, flagScoreDefault, flagScoreDescription)

	return nil
}

func registerIPCVEPropServices(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: services []string array type is not supported by go-swagger cli yet

	return nil
}

func registerIPCVEPropSeverity(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagSeverityDescription := ``

	var flagSeverityName string
	if cmdPrefix == "" {
		flagSeverityName = "severity"
	} else {
		flagSeverityName = fmt.Sprintf("%v.severity", cmdPrefix)
	}

	var flagSeverityDefault string

	_ = cmd.PersistentFlags().String(flagSeverityName, flagSeverityDefault, flagSeverityDescription)

	return nil
}

func registerIPCVEPropSummary(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagSummaryDescription := ``

	var flagSummaryName string
	if cmdPrefix == "" {
		flagSummaryName = "summary"
	} else {
		flagSummaryName = fmt.Sprintf("%v.summary", cmdPrefix)
	}

	var flagSummaryDefault string

	_ = cmd.PersistentFlags().String(flagSummaryName, flagSummaryDefault, flagSummaryDescription)

	return nil
}

func registerIPCVEPropVectorString(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagVectorStringDescription := ``

	var flagVectorStringName string
	if cmdPrefix == "" {
		flagVectorStringName = "vector_string"
	} else {
		flagVectorStringName = fmt.Sprintf("%v.vector_string", cmdPrefix)
	}

	var flagVectorStringDefault string

	_ = cmd.PersistentFlags().String(flagVectorStringName, flagVectorStringDefault, flagVectorStringDescription)

	return nil
}

func registerIPCVEPropWeakness(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagWeaknessDescription := ``

	var flagWeaknessName string
	if cmdPrefix == "" {
		flagWeaknessName = "weakness"
	} else {
		flagWeaknessName = fmt.Sprintf("%v.weakness", cmdPrefix)
	}

	var flagWeaknessDefault string

	_ = cmd.PersistentFlags().String(flagWeaknessName, flagWeaknessDefault, flagWeaknessDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelIPCVEFlags(depth int, m *models.IPCVE, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, ExploitAdded := retrieveIPCVEPropExploitFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ExploitAdded

	err, IDAdded := retrieveIPCVEPropIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || IDAdded

	err, ReferencesAdded := retrieveIPCVEPropReferencesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ReferencesAdded

	err, ScoreAdded := retrieveIPCVEPropScoreFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ScoreAdded

	err, ServicesAdded := retrieveIPCVEPropServicesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ServicesAdded

	err, SeverityAdded := retrieveIPCVEPropSeverityFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || SeverityAdded

	err, SummaryAdded := retrieveIPCVEPropSummaryFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || SummaryAdded

	err, VectorStringAdded := retrieveIPCVEPropVectorStringFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || VectorStringAdded

	err, WeaknessAdded := retrieveIPCVEPropWeaknessFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || WeaknessAdded

	return nil, retAdded
}

func retrieveIPCVEPropExploitFlags(depth int, m *models.IPCVE, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagExploitName := fmt.Sprintf("%v.exploit", cmdPrefix)
	if cmd.Flags().Changed(flagExploitName) {
		// warning: exploit array type []*IPExploitDetails is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveIPCVEPropIDFlags(depth int, m *models.IPCVE, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagIDName := fmt.Sprintf("%v.id", cmdPrefix)
	if cmd.Flags().Changed(flagIDName) {

		var flagIDName string
		if cmdPrefix == "" {
			flagIDName = "id"
		} else {
			flagIDName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		flagIDValue, err := cmd.Flags().GetString(flagIDName)
		if err != nil {
			return err, false
		}
		m.ID = flagIDValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIPCVEPropReferencesFlags(depth int, m *models.IPCVE, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagReferencesName := fmt.Sprintf("%v.references", cmdPrefix)
	if cmd.Flags().Changed(flagReferencesName) {
		// warning: references array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveIPCVEPropScoreFlags(depth int, m *models.IPCVE, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagScoreName := fmt.Sprintf("%v.score", cmdPrefix)
	if cmd.Flags().Changed(flagScoreName) {

		var flagScoreName string
		if cmdPrefix == "" {
			flagScoreName = "score"
		} else {
			flagScoreName = fmt.Sprintf("%v.score", cmdPrefix)
		}

		flagScoreValue, err := cmd.Flags().GetFloat64(flagScoreName)
		if err != nil {
			return err, false
		}
		m.Score = flagScoreValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIPCVEPropServicesFlags(depth int, m *models.IPCVE, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagServicesName := fmt.Sprintf("%v.services", cmdPrefix)
	if cmd.Flags().Changed(flagServicesName) {
		// warning: services array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveIPCVEPropSeverityFlags(depth int, m *models.IPCVE, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagSeverityName := fmt.Sprintf("%v.severity", cmdPrefix)
	if cmd.Flags().Changed(flagSeverityName) {

		var flagSeverityName string
		if cmdPrefix == "" {
			flagSeverityName = "severity"
		} else {
			flagSeverityName = fmt.Sprintf("%v.severity", cmdPrefix)
		}

		flagSeverityValue, err := cmd.Flags().GetString(flagSeverityName)
		if err != nil {
			return err, false
		}
		m.Severity = flagSeverityValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIPCVEPropSummaryFlags(depth int, m *models.IPCVE, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagSummaryName := fmt.Sprintf("%v.summary", cmdPrefix)
	if cmd.Flags().Changed(flagSummaryName) {

		var flagSummaryName string
		if cmdPrefix == "" {
			flagSummaryName = "summary"
		} else {
			flagSummaryName = fmt.Sprintf("%v.summary", cmdPrefix)
		}

		flagSummaryValue, err := cmd.Flags().GetString(flagSummaryName)
		if err != nil {
			return err, false
		}
		m.Summary = flagSummaryValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIPCVEPropVectorStringFlags(depth int, m *models.IPCVE, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagVectorStringName := fmt.Sprintf("%v.vector_string", cmdPrefix)
	if cmd.Flags().Changed(flagVectorStringName) {

		var flagVectorStringName string
		if cmdPrefix == "" {
			flagVectorStringName = "vector_string"
		} else {
			flagVectorStringName = fmt.Sprintf("%v.vector_string", cmdPrefix)
		}

		flagVectorStringValue, err := cmd.Flags().GetString(flagVectorStringName)
		if err != nil {
			return err, false
		}
		m.VectorString = flagVectorStringValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIPCVEPropWeaknessFlags(depth int, m *models.IPCVE, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagWeaknessName := fmt.Sprintf("%v.weakness", cmdPrefix)
	if cmd.Flags().Changed(flagWeaknessName) {

		var flagWeaknessName string
		if cmdPrefix == "" {
			flagWeaknessName = "weakness"
		} else {
			flagWeaknessName = fmt.Sprintf("%v.weakness", cmdPrefix)
		}

		flagWeaknessValue, err := cmd.Flags().GetString(flagWeaknessName)
		if err != nil {
			return err, false
		}
		m.Weakness = flagWeaknessValue

		retAdded = true
	}

	return nil, retAdded
}
