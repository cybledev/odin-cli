// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/cybledev/odin-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for EshandlerEXTDNSName

// register flags to command
func registerModelEshandlerEXTDNSNameFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerEshandlerEXTDNSNamePropDomain(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEshandlerEXTDNSNamePropFld(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEshandlerEXTDNSNamePropSubdomain(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEshandlerEXTDNSNamePropTld(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerEshandlerEXTDNSNamePropDomain(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerEshandlerEXTDNSNamePropFld(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagFldDescription := ``

	var flagFldName string
	if cmdPrefix == "" {
		flagFldName = "fld"
	} else {
		flagFldName = fmt.Sprintf("%v.fld", cmdPrefix)
	}

	var flagFldDefault string

	_ = cmd.PersistentFlags().String(flagFldName, flagFldDefault, flagFldDescription)

	return nil
}

func registerEshandlerEXTDNSNamePropSubdomain(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagSubdomainDescription := ``

	var flagSubdomainName string
	if cmdPrefix == "" {
		flagSubdomainName = "subdomain"
	} else {
		flagSubdomainName = fmt.Sprintf("%v.subdomain", cmdPrefix)
	}

	var flagSubdomainDefault string

	_ = cmd.PersistentFlags().String(flagSubdomainName, flagSubdomainDefault, flagSubdomainDescription)

	return nil
}

func registerEshandlerEXTDNSNamePropTld(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagTldDescription := ``

	var flagTldName string
	if cmdPrefix == "" {
		flagTldName = "tld"
	} else {
		flagTldName = fmt.Sprintf("%v.tld", cmdPrefix)
	}

	var flagTldDefault string

	_ = cmd.PersistentFlags().String(flagTldName, flagTldDefault, flagTldDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelEshandlerEXTDNSNameFlags(depth int, m *models.EshandlerEXTDNSName, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, DomainAdded := retrieveEshandlerEXTDNSNamePropDomainFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DomainAdded

	err, FldAdded := retrieveEshandlerEXTDNSNamePropFldFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || FldAdded

	err, SubdomainAdded := retrieveEshandlerEXTDNSNamePropSubdomainFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || SubdomainAdded

	err, TldAdded := retrieveEshandlerEXTDNSNamePropTldFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || TldAdded

	return nil, retAdded
}

func retrieveEshandlerEXTDNSNamePropDomainFlags(depth int, m *models.EshandlerEXTDNSName, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveEshandlerEXTDNSNamePropFldFlags(depth int, m *models.EshandlerEXTDNSName, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagFldName := fmt.Sprintf("%v.fld", cmdPrefix)
	if cmd.Flags().Changed(flagFldName) {

		var flagFldName string
		if cmdPrefix == "" {
			flagFldName = "fld"
		} else {
			flagFldName = fmt.Sprintf("%v.fld", cmdPrefix)
		}

		flagFldValue, err := cmd.Flags().GetString(flagFldName)
		if err != nil {
			return err, false
		}
		m.Fld = flagFldValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveEshandlerEXTDNSNamePropSubdomainFlags(depth int, m *models.EshandlerEXTDNSName, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagSubdomainName := fmt.Sprintf("%v.subdomain", cmdPrefix)
	if cmd.Flags().Changed(flagSubdomainName) {

		var flagSubdomainName string
		if cmdPrefix == "" {
			flagSubdomainName = "subdomain"
		} else {
			flagSubdomainName = fmt.Sprintf("%v.subdomain", cmdPrefix)
		}

		flagSubdomainValue, err := cmd.Flags().GetString(flagSubdomainName)
		if err != nil {
			return err, false
		}
		m.Subdomain = flagSubdomainValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveEshandlerEXTDNSNamePropTldFlags(depth int, m *models.EshandlerEXTDNSName, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagTldName := fmt.Sprintf("%v.tld", cmdPrefix)
	if cmd.Flags().Changed(flagTldName) {

		var flagTldName string
		if cmdPrefix == "" {
			flagTldName = "tld"
		} else {
			flagTldName = fmt.Sprintf("%v.tld", cmdPrefix)
		}

		flagTldValue, err := cmd.Flags().GetString(flagTldName)
		if err != nil {
			return err, false
		}
		m.Tld = flagTldValue

		retAdded = true
	}

	return nil, retAdded
}
