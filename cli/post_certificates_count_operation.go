// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/cybledev/odin-cli/models"
	"github.com/cybledev/odin-cli/odin_cli/certificate"

	"github.com/cybledev/odin-cli/decorator"
	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationCertificatePostCertificatesCountCmd returns a command to handle operation postCertificatesCount
func makeOperationCertificatePostCertificatesCountCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Args: cobra.NoArgs,
		Use:   "count",
		Short: `Returns the total no of records based on query`,
		RunE:  decorator.RunEColorWrapper( runOperationCertificatePostCertificatesCount),
	}

	if err := registerOperationCertificatePostCertificatesCountParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationCertificatePostCertificatesCount uses cmd flags to call endpoint api
func runOperationCertificatePostCertificatesCount(cmd *cobra.Command, args []string) (string, error) {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return "", err
	}
	// retrieve flag values from cmd and fill params
	params := certificate.NewPostCertificatesCountParams()
	if err, _ = retrieveOperationCertificatePostCertificatesCountQueryFlag(params, "", cmd); err != nil {
		return "", err
	}
	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return "", nil
	}
	// make request and then print result
	msgStr, err := parseOperationCertificatePostCertificatesCountResult(appCli.Certificate.PostCertificatesCount(params, nil))
	if err != nil {
		return "", err
	}

	if !debug {
		return msgStr, err
	}

	return "", nil
}

// registerOperationCertificatePostCertificatesCountParamFlags registers all flags needed to fill params
func registerOperationCertificatePostCertificatesCountParamFlags(cmd *cobra.Command) error {
	if err := registerOperationCertificatePostCertificatesCountQueryParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationCertificatePostCertificatesCountQueryParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var flagQueryName string
	if cmdPrefix == "" {
		flagQueryName = "query"
	} else {
		flagQueryName = fmt.Sprintf("%v.query", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(flagQueryName, "", `Optional json string for [query]. Count Query`)

	// add flags for body
	if err := registerModelCertificateCertCountRequestFlags(0, "request", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationCertificatePostCertificatesCountQueryFlag(m *certificate.PostCertificatesCountParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("query") {
		// Read query string from cmd and unmarshal
		flagQueryValueStr, err := cmd.Flags().GetString("query")
		if err != nil {
			return err, false
		}

		flagQueryValue := models.CertificateCertCountRequest{}
		if err := json.Unmarshal([]byte(flagQueryValueStr), &flagQueryValue); err != nil {
			return fmt.Errorf("cannot unmarshal query string in models.CertificateCertCountRequest: %v", err), false
		}
		m.Query = &flagQueryValue
	}
	flagQueryModel := m.Query
	if swag.IsZero(flagQueryModel) {
		flagQueryModel = &models.CertificateCertCountRequest{}
	}
	err, added := retrieveModelCertificateCertCountRequestFlags(0, flagQueryModel, "request", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Query = flagQueryModel
	}

	if dryRun && debug {
		flagQueryValueDebugBytes, err := json.Marshal(m.Query)
		if err != nil {
			return err, false
		}
		logDebugf("Query dry-run payload: %v", string(flagQueryValueDebugBytes))
	}

	retAdded = retAdded || added

	return nil, retAdded
}

// parseOperationCertificatePostCertificatesCountResult parses request result and return the string content
func parseOperationCertificatePostCertificatesCountResult(resp0 *certificate.PostCertificatesCountOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*certificate.PostCertificatesCountOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*certificate.PostCertificatesCountBadRequest)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*certificate.PostCertificatesCountRequestTimeout)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp3 interface{} = respErr
		resp3, ok := iResp3.(*certificate.PostCertificatesCountInternalServerError)
		if ok {
			if !swag.IsZero(resp3) && !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}

// register flags to command
func registerModelPostCertificatesCountOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	// register embedded models.CertificateAPIResponse flags

	if err := registerModelCertificateAPIResponseFlags(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	// register anonymous fields for postCertificatesCountOKBodyAO1

	if err := registerPostCertificatesCountOKBodyPropAnonPostCertificatesCountOKBodyAO1Data(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

// inline definition name postCertificatesCountOKBodyAO1, type

func registerPostCertificatesCountOKBodyPropAnonPostCertificatesCountOKBodyAO1Data(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var flagDataName string
	if cmdPrefix == "" {
		flagDataName = "data"
	} else {
		flagDataName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelCertificateCertCountFlags(depth+1, flagDataName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPostCertificatesCountOKBodyFlags(depth int, m *certificate.PostCertificatesCountOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	// retrieve model models.CertificateAPIResponse
	err, PostCertificatesCountOKBodyAO0Added := retrieveModelCertificateAPIResponseFlags(depth, &m.CertificateAPIResponse, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || PostCertificatesCountOKBodyAO0Added

	// retrieve allOf postCertificatesCountOKBodyAO1 fields

	err, DataAdded := retrievePostCertificatesCountOKBodyPropAnonPostCertificatesCountOKBodyAO1DataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DataAdded

	return nil, retAdded
}

// define retrieve functions for fields for inline definition name postCertificatesCountOKBodyAO1

func retrievePostCertificatesCountOKBodyPropAnonPostCertificatesCountOKBodyAO1DataFlags(depth int, m *certificate.PostCertificatesCountOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagDataName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(flagDataName) {
		// info: complex object data models.CertificateCertCount is retrieved outside this Changed() block
	}
	flagDataValue := m.Data
	if swag.IsZero(flagDataValue) {
		flagDataValue = &models.CertificateCertCount{}
	}

	err, DataAdded := retrieveModelCertificateCertCountFlags(depth+1, flagDataValue, flagDataName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DataAdded
	if DataAdded {
		m.Data = flagDataValue
	}

	return nil, retAdded
}
