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

// makeOperationCertificatePostCertificatesScrollNextCmd returns a command to handle operation postCertificatesScrollNext
func makeOperationCertificatePostCertificatesScrollNextCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Args: cobra.NoArgs,
		Use:   "next",
		Short: `Returns the next batch of record based on query. It uses es scroll api for the pagination.`,
		RunE:  decorator.RunEColorWrapper(runOperationCertificatePostCertificatesScrollNext),
	}

	if err := registerOperationCertificatePostCertificatesScrollNextParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationCertificatePostCertificatesScrollNext uses cmd flags to call endpoint api
func runOperationCertificatePostCertificatesScrollNext(cmd *cobra.Command, args []string) (string, error) {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return "", err
	}
	// retrieve flag values from cmd and fill params
	params := certificate.NewPostCertificatesScrollNextParams()
	if err, _ = retrieveOperationCertificatePostCertificatesScrollNextQueryFlag(params, "", cmd); err != nil {
		return "", err
	}
	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return "", nil
	}
	// make request and then print result
	msgStr, err := parseOperationCertificatePostCertificatesScrollNextResult(appCli.Certificate.PostCertificatesScrollNext(params, nil))
	if err != nil {
		return "", err
	}

	if !debug {
		return msgStr, nil
	}

	return "", nil
}

// registerOperationCertificatePostCertificatesScrollNextParamFlags registers all flags needed to fill params
func registerOperationCertificatePostCertificatesScrollNextParamFlags(cmd *cobra.Command) error {
	if err := registerOperationCertificatePostCertificatesScrollNextQueryParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationCertificatePostCertificatesScrollNextQueryParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var flagQueryName string
	if cmdPrefix == "" {
		flagQueryName = "query"
	} else {
		flagQueryName = fmt.Sprintf("%v.query", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(flagQueryName, "", `Optional json string for [query]. Search Query`)

	// add flags for body
	if err := registerModelCertificateNextBatchRequestFlags(0, "request", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationCertificatePostCertificatesScrollNextQueryFlag(m *certificate.PostCertificatesScrollNextParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("query") {
		// Read query string from cmd and unmarshal
		flagQueryValueStr, err := cmd.Flags().GetString("query")
		if err != nil {
			return err, false
		}

		flagQueryValue := models.CertificateNextBatchRequest{}
		if err := json.Unmarshal([]byte(flagQueryValueStr), &flagQueryValue); err != nil {
			return fmt.Errorf("cannot unmarshal query string in models.CertificateNextBatchRequest: %v", err), false
		}
		m.Query = &flagQueryValue
	}
	flagQueryModel := m.Query
	if swag.IsZero(flagQueryModel) {
		flagQueryModel = &models.CertificateNextBatchRequest{}
	}
	err, added := retrieveModelCertificateNextBatchRequestFlags(0, flagQueryModel, "request", cmd)
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

// parseOperationCertificatePostCertificatesScrollNextResult parses request result and return the string content
func parseOperationCertificatePostCertificatesScrollNextResult(resp0 *certificate.PostCertificatesScrollNextOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*certificate.PostCertificatesScrollNextOK)
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
		resp1, ok := iResp1.(*certificate.PostCertificatesScrollNextBadRequest)
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
		resp2, ok := iResp2.(*certificate.PostCertificatesScrollNextRequestTimeout)
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
		resp3, ok := iResp3.(*certificate.PostCertificatesScrollNextInternalServerError)
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
func registerModelPostCertificatesScrollNextOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	// register embedded models.CertificateAPIResponse flags

	if err := registerModelCertificateAPIResponseFlags(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	// register anonymous fields for postCertificatesScrollNextOKBodyAO1

	if err := registerPostCertificatesScrollNextOKBodyPropAnonPostCertificatesScrollNextOKBodyAO1Data(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

// inline definition name postCertificatesScrollNextOKBodyAO1, type

func registerPostCertificatesScrollNextOKBodyPropAnonPostCertificatesScrollNextOKBodyAO1Data(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var flagDataName string
	if cmdPrefix == "" {
		flagDataName = "data"
	} else {
		flagDataName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelCertificateCertScrollFlags(depth+1, flagDataName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPostCertificatesScrollNextOKBodyFlags(depth int, m *certificate.PostCertificatesScrollNextOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	// retrieve model models.CertificateAPIResponse
	err, PostCertificatesScrollNextOKBodyAO0Added := retrieveModelCertificateAPIResponseFlags(depth, &m.CertificateAPIResponse, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || PostCertificatesScrollNextOKBodyAO0Added

	// retrieve allOf postCertificatesScrollNextOKBodyAO1 fields

	err, DataAdded := retrievePostCertificatesScrollNextOKBodyPropAnonPostCertificatesScrollNextOKBodyAO1DataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DataAdded

	return nil, retAdded
}

// define retrieve functions for fields for inline definition name postCertificatesScrollNextOKBodyAO1

func retrievePostCertificatesScrollNextOKBodyPropAnonPostCertificatesScrollNextOKBodyAO1DataFlags(depth int, m *certificate.PostCertificatesScrollNextOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagDataName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(flagDataName) {
		// info: complex object data models.CertificateCertScroll is retrieved outside this Changed() block
	}
	flagDataValue := m.Data
	if swag.IsZero(flagDataValue) {
		flagDataValue = &models.CertificateCertScroll{}
	}

	err, DataAdded := retrieveModelCertificateCertScrollFlags(depth+1, flagDataValue, flagDataName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DataAdded
	if DataAdded {
		m.Data = flagDataValue
	}

	return nil, retAdded
}
