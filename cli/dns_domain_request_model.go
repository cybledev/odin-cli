// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/cybledev/odin-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for DNSDomainRequest

// register flags to command
func registerModelDNSDomainRequestFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDNSDomainRequestPropDomain(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDNSDomainRequestPropKeyword(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDNSDomainRequestPropLimit(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDNSDomainRequestPropPublishedAfter(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDNSDomainRequestPropPublishedBefore(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDNSDomainRequestPropStart(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDNSDomainRequestPropDomain(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagDomainDescription := ``

	var flagDomainName string
	if cmdPrefix == "" {
		flagDomainName = "domain"
	} else {
		flagDomainName = fmt.Sprintf("%v.domain", cmdPrefix)
	}

	var flagDomainDefault string

	_ = cmd.PersistentFlags().String(flagDomainName, flagDomainDefault, flagDomainDescription)

	return nil
}

func registerDNSDomainRequestPropKeyword(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagKeywordDescription := ``

	var flagKeywordName string
	if cmdPrefix == "" {
		flagKeywordName = "keyword"
	} else {
		flagKeywordName = fmt.Sprintf("%v.keyword", cmdPrefix)
	}

	var flagKeywordDefault string

	_ = cmd.PersistentFlags().String(flagKeywordName, flagKeywordDefault, flagKeywordDescription)

	return nil
}

func registerDNSDomainRequestPropLimit(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagLimitDescription := `Required. `

	var flagLimitName string
	if cmdPrefix == "" {
		flagLimitName = "limit"
	} else {
		flagLimitName = fmt.Sprintf("%v.limit", cmdPrefix)
	}

	var flagLimitDefault int64

	_ = cmd.PersistentFlags().Int64(flagLimitName, flagLimitDefault, flagLimitDescription)

	return nil
}

func registerDNSDomainRequestPropPublishedAfter(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagPublishedAfterDescription := ``

	var flagPublishedAfterName string
	if cmdPrefix == "" {
		flagPublishedAfterName = "publishedAfter"
	} else {
		flagPublishedAfterName = fmt.Sprintf("%v.publishedAfter", cmdPrefix)
	}

	var flagPublishedAfterDefault string

	_ = cmd.PersistentFlags().String(flagPublishedAfterName, flagPublishedAfterDefault, flagPublishedAfterDescription)

	return nil
}

func registerDNSDomainRequestPropPublishedBefore(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagPublishedBeforeDescription := ``

	var flagPublishedBeforeName string
	if cmdPrefix == "" {
		flagPublishedBeforeName = "publishedBefore"
	} else {
		flagPublishedBeforeName = fmt.Sprintf("%v.publishedBefore", cmdPrefix)
	}

	var flagPublishedBeforeDefault string

	_ = cmd.PersistentFlags().String(flagPublishedBeforeName, flagPublishedBeforeDefault, flagPublishedBeforeDescription)

	return nil
}

func registerDNSDomainRequestPropStart(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: start []interface{} array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDNSDomainRequestFlags(depth int, m *models.DNSDomainRequest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, DomainAdded := retrieveDNSDomainRequestPropDomainFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DomainAdded

	err, KeywordAdded := retrieveDNSDomainRequestPropKeywordFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || KeywordAdded

	err, LimitAdded := retrieveDNSDomainRequestPropLimitFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || LimitAdded

	err, PublishedAfterAdded := retrieveDNSDomainRequestPropPublishedAfterFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || PublishedAfterAdded

	err, PublishedBeforeAdded := retrieveDNSDomainRequestPropPublishedBeforeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || PublishedBeforeAdded

	err, StartAdded := retrieveDNSDomainRequestPropStartFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || StartAdded

	return nil, retAdded
}

func retrieveDNSDomainRequestPropDomainFlags(depth int, m *models.DNSDomainRequest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagDomainName := fmt.Sprintf("%v.domain", cmdPrefix)
	if cmd.Flags().Changed(flagDomainName) {

		var flagDomainName string
		if cmdPrefix == "" {
			flagDomainName = "domain"
		} else {
			flagDomainName = fmt.Sprintf("%v.domain", cmdPrefix)
		}

		flagDomainValue, err := cmd.Flags().GetString(flagDomainName)
		if err != nil {
			return err, false
		}
		m.Domain = flagDomainValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDNSDomainRequestPropKeywordFlags(depth int, m *models.DNSDomainRequest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagKeywordName := fmt.Sprintf("%v.keyword", cmdPrefix)
	if cmd.Flags().Changed(flagKeywordName) {

		var flagKeywordName string
		if cmdPrefix == "" {
			flagKeywordName = "keyword"
		} else {
			flagKeywordName = fmt.Sprintf("%v.keyword", cmdPrefix)
		}

		flagKeywordValue, err := cmd.Flags().GetString(flagKeywordName)
		if err != nil {
			return err, false
		}
		m.Keyword = flagKeywordValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDNSDomainRequestPropLimitFlags(depth int, m *models.DNSDomainRequest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagLimitName := fmt.Sprintf("%v.limit", cmdPrefix)
	if cmd.Flags().Changed(flagLimitName) {

		var flagLimitName string
		if cmdPrefix == "" {
			flagLimitName = "limit"
		} else {
			flagLimitName = fmt.Sprintf("%v.limit", cmdPrefix)
		}

		flagLimitValue, err := cmd.Flags().GetInt64(flagLimitName)
		if err != nil {
			return err, false
		}
		m.Limit = &flagLimitValue

		retAdded = true
	} else {
		m.Limit = defaultLimit
		retAdded = true
	}

	return nil, retAdded
}

func retrieveDNSDomainRequestPropPublishedAfterFlags(depth int, m *models.DNSDomainRequest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagPublishedAfterName := fmt.Sprintf("%v.publishedAfter", cmdPrefix)
	if cmd.Flags().Changed(flagPublishedAfterName) {

		var flagPublishedAfterName string
		if cmdPrefix == "" {
			flagPublishedAfterName = "publishedAfter"
		} else {
			flagPublishedAfterName = fmt.Sprintf("%v.publishedAfter", cmdPrefix)
		}

		flagPublishedAfterValue, err := cmd.Flags().GetString(flagPublishedAfterName)
		if err != nil {
			return err, false
		}
		m.PublishedAfter = flagPublishedAfterValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDNSDomainRequestPropPublishedBeforeFlags(depth int, m *models.DNSDomainRequest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagPublishedBeforeName := fmt.Sprintf("%v.publishedBefore", cmdPrefix)
	if cmd.Flags().Changed(flagPublishedBeforeName) {

		var flagPublishedBeforeName string
		if cmdPrefix == "" {
			flagPublishedBeforeName = "publishedBefore"
		} else {
			flagPublishedBeforeName = fmt.Sprintf("%v.publishedBefore", cmdPrefix)
		}

		flagPublishedBeforeValue, err := cmd.Flags().GetString(flagPublishedBeforeName)
		if err != nil {
			return err, false
		}
		m.PublishedBefore = flagPublishedBeforeValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDNSDomainRequestPropStartFlags(depth int, m *models.DNSDomainRequest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagStartName := fmt.Sprintf("%v.start", cmdPrefix)
	if cmd.Flags().Changed(flagStartName) {
		// warning: start array type []interface{} is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
