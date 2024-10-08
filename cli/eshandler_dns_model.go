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

// Schema cli for EshandlerDNS

// register flags to command
func registerModelEshandlerDNSFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerEshandlerDNSPropAddedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEshandlerDNSPropDomain(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEshandlerDNSPropExtDNSName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEshandlerDNSPropSubdomain(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerEshandlerDNSPropAddedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagAddedAtDescription := ``

	var flagAddedAtName string
	if cmdPrefix == "" {
		flagAddedAtName = "added_at"
	} else {
		flagAddedAtName = fmt.Sprintf("%v.added_at", cmdPrefix)
	}

	var flagAddedAtDefault string

	_ = cmd.PersistentFlags().String(flagAddedAtName, flagAddedAtDefault, flagAddedAtDescription)

	return nil
}

func registerEshandlerDNSPropDomain(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerEshandlerDNSPropExtDNSName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var flagExtDNSNameName string
	if cmdPrefix == "" {
		flagExtDNSNameName = "ext_dns_name"
	} else {
		flagExtDNSNameName = fmt.Sprintf("%v.ext_dns_name", cmdPrefix)
	}

	if err := registerModelEshandlerEXTDNSNameFlags(depth+1, flagExtDNSNameName, cmd); err != nil {
		return err
	}

	return nil
}

func registerEshandlerDNSPropSubdomain(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelEshandlerDNSFlags(depth int, m *models.EshandlerDNS, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, AddedAtAdded := retrieveEshandlerDNSPropAddedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || AddedAtAdded

	err, DomainAdded := retrieveEshandlerDNSPropDomainFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DomainAdded

	err, ExtDNSNameAdded := retrieveEshandlerDNSPropExtDNSNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ExtDNSNameAdded

	err, SubdomainAdded := retrieveEshandlerDNSPropSubdomainFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || SubdomainAdded

	return nil, retAdded
}

func retrieveEshandlerDNSPropAddedAtFlags(depth int, m *models.EshandlerDNS, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagAddedAtName := fmt.Sprintf("%v.added_at", cmdPrefix)
	if cmd.Flags().Changed(flagAddedAtName) {

		var flagAddedAtName string
		if cmdPrefix == "" {
			flagAddedAtName = "added_at"
		} else {
			flagAddedAtName = fmt.Sprintf("%v.added_at", cmdPrefix)
		}

		flagAddedAtValue, err := cmd.Flags().GetString(flagAddedAtName)
		if err != nil {
			return err, false
		}
		m.AddedAt = flagAddedAtValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveEshandlerDNSPropDomainFlags(depth int, m *models.EshandlerDNS, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveEshandlerDNSPropExtDNSNameFlags(depth int, m *models.EshandlerDNS, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagExtDNSNameName := fmt.Sprintf("%v.ext_dns_name", cmdPrefix)
	if cmd.Flags().Changed(flagExtDNSNameName) {
		// info: complex object ext_dns_name EshandlerEXTDNSName is retrieved outside this Changed() block
	}
	flagExtDNSNameValue := m.ExtDNSName
	if swag.IsZero(flagExtDNSNameValue) {
		flagExtDNSNameValue = &models.EshandlerEXTDNSName{}
	}

	err, ExtDNSNameAdded := retrieveModelEshandlerEXTDNSNameFlags(depth+1, flagExtDNSNameValue, flagExtDNSNameName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ExtDNSNameAdded
	if ExtDNSNameAdded {
		m.ExtDNSName = flagExtDNSNameValue
	}

	return nil, retAdded
}

func retrieveEshandlerDNSPropSubdomainFlags(depth int, m *models.EshandlerDNS, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
