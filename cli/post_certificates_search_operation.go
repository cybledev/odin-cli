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

// makeOperationCertificatePostCertificatesSearchCmd returns a command to handle operation postCertificatesSearch
func makeOperationCertificatePostCertificatesSearchCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Args: cobra.NoArgs,
		Use:   "search",
		Short: `Returns the record based on query and blank query will return all records. It uses es searchafter for the pagination.`,
		RunE:  decorator.RunEColorWrapper(runOperationCertificatePostCertificatesSearch),
	}

	if err := registerOperationCertificatePostCertificatesSearchParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationCertificatePostCertificatesSearch uses cmd flags to call endpoint api
func runOperationCertificatePostCertificatesSearch(cmd *cobra.Command, args []string) (string, error) {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return "", err
	}
	// retrieve flag values from cmd and fill params
	params := certificate.NewPostCertificatesSearchParams()
	if err, _ = retrieveOperationCertificatePostCertificatesSearchQueryFlag(params, "", cmd); err != nil {
		return "", err
	}
	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return "", nil
	}
	// make request and then print result
	msgStr, err := parseOperationCertificatePostCertificatesSearchResult(appCli.Certificate.PostCertificatesSearch(params, nil))
	if err != nil {
		return "", err
	}

	if !debug {
		return msgStr, nil
	}

	return "", nil
}

// registerOperationCertificatePostCertificatesSearchParamFlags registers all flags needed to fill params
func registerOperationCertificatePostCertificatesSearchParamFlags(cmd *cobra.Command) error {
	if err := registerOperationCertificatePostCertificatesSearchQueryParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationCertificatePostCertificatesSearchQueryParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var flagQueryName string
	if cmdPrefix == "" {
		flagQueryName = "query"
	} else {
		flagQueryName = fmt.Sprintf("%v.query", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(flagQueryName, "", `Optional json string for [query]. Search Query`)

	// add flags for body
	if err := registerModelCertificateCertSearchRequestFlags(0, "request", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationCertificatePostCertificatesSearchQueryFlag(m *certificate.PostCertificatesSearchParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("query") {
		// Read query string from cmd and unmarshal
		flagQueryValueStr, err := cmd.Flags().GetString("query")
		if err != nil {
			return err, false
		}

		flagQueryValue := models.CertificateCertSearchRequest{}
		if err := json.Unmarshal([]byte(flagQueryValueStr), &flagQueryValue); err != nil {
			return fmt.Errorf("cannot unmarshal query string in models.CertificateCertSearchRequest: %v", err), false
		}
		m.Query = &flagQueryValue
	}
	flagQueryModel := m.Query
	if swag.IsZero(flagQueryModel) {
		flagQueryModel = &models.CertificateCertSearchRequest{}
	}
	err, added := retrieveModelCertificateCertSearchRequestFlags(0, flagQueryModel, "request", cmd)
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

// parseOperationCertificatePostCertificatesSearchResult parses request result and return the string content
func parseOperationCertificatePostCertificatesSearchResult(resp0 *certificate.PostCertificatesSearchOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*certificate.PostCertificatesSearchOK)
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
		resp1, ok := iResp1.(*certificate.PostCertificatesSearchBadRequest)
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
		resp2, ok := iResp2.(*certificate.PostCertificatesSearchRequestTimeout)
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
		resp3, ok := iResp3.(*certificate.PostCertificatesSearchInternalServerError)
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
func registerModelPostCertificatesSearchOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	// register embedded models.CertificateAPIResponse flags

	if err := registerModelCertificateAPIResponseFlags(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	// register anonymous fields for postCertificatesSearchOKBodyAO1

	if err := registerPostCertificatesSearchOKBodyPropAnonPostCertificatesSearchOKBodyAO1Data(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPostCertificatesSearchOKBodyPropAnonPostCertificatesSearchOKBodyAO1Pagination(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

// inline definition name postCertificatesSearchOKBodyAO1, type

func registerPostCertificatesSearchOKBodyPropAnonPostCertificatesSearchOKBodyAO1Data(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: data []interface{} array type is not supported by go-swagger cli yet

	return nil
}

func registerPostCertificatesSearchOKBodyPropAnonPostCertificatesSearchOKBodyAO1Pagination(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var flagPaginationName string
	if cmdPrefix == "" {
		flagPaginationName = "pagination"
	} else {
		flagPaginationName = fmt.Sprintf("%v.pagination", cmdPrefix)
	}

	if err := registerModelCertificateSearchPaginationFlags(depth+1, flagPaginationName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPostCertificatesSearchOKBodyFlags(depth int, m *certificate.PostCertificatesSearchOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	// retrieve model models.CertificateAPIResponse
	err, PostCertificatesSearchOKBodyAO0Added := retrieveModelCertificateAPIResponseFlags(depth, &m.CertificateAPIResponse, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || PostCertificatesSearchOKBodyAO0Added

	// retrieve allOf postCertificatesSearchOKBodyAO1 fields

	err, DataAdded := retrievePostCertificatesSearchOKBodyPropAnonPostCertificatesSearchOKBodyAO1DataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DataAdded

	err, PaginationAdded := retrievePostCertificatesSearchOKBodyPropAnonPostCertificatesSearchOKBodyAO1PaginationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || PaginationAdded

	return nil, retAdded
}

// define retrieve functions for fields for inline definition name postCertificatesSearchOKBodyAO1

func retrievePostCertificatesSearchOKBodyPropAnonPostCertificatesSearchOKBodyAO1DataFlags(depth int, m *certificate.PostCertificatesSearchOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagDataName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(flagDataName) {
		// warning: data array type []interface{} is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrievePostCertificatesSearchOKBodyPropAnonPostCertificatesSearchOKBodyAO1PaginationFlags(depth int, m *certificate.PostCertificatesSearchOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagPaginationName := fmt.Sprintf("%v.pagination", cmdPrefix)
	if cmd.Flags().Changed(flagPaginationName) {
		// info: complex object pagination models.CertificateSearchPagination is retrieved outside this Changed() block
	}
	flagPaginationValue := m.Pagination
	if swag.IsZero(flagPaginationValue) {
		flagPaginationValue = &models.CertificateSearchPagination{}
	}

	err, PaginationAdded := retrieveModelCertificateSearchPaginationFlags(depth+1, flagPaginationValue, flagPaginationName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || PaginationAdded
	if PaginationAdded {
		m.Pagination = flagPaginationValue
	}

	return nil, retAdded
}
