// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/cybledev/odin-cli/decorator"
	"github.com/cybledev/odin-cli/models"
	"github.com/cybledev/odin-cli/odin_cli/hosts"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationHostsPostV2HostsSearchCmd returns a command to handle operation postV2HostsSearch
func makeOperationHostsPostV2HostsSearchCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "v2-search",
		Short: `Returns the v2 record based on query and blank query will return all records. It uses es searchafter for the pagination.`,
		RunE:  decorator.RunEColorWrapper(runOperationHostsPostV2HostsSearch),
	}

	if err := registerOperationHostsPostV2HostsSearchParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationHostsPostV2HostsSearch uses cmd flags to call endpoint api
func runOperationHostsPostV2HostsSearch(cmd *cobra.Command, args []string) (string, error) {
	args = append(args, apiVersionPrefix + "/v2/")
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return "", err
	}
	// retrieve flag values from cmd and fill params
	params := hosts.NewPostV2HostsSearchParams()
	if err, _ = retrieveOperationHostsPostV2HostsSearchQueryFlag(params, "", cmd); err != nil {
		return "", err
	}
	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return "", nil
	}
	// make request and then print result
	msgStr, err := parseOperationHostsPostV2HostsSearchResult(appCli.Hosts.PostV2HostsSearch(params, nil))
	if err != nil {
		return "", err
	}

	if !debug {
		return msgStr, err
	}

	return "", nil
}

// registerOperationHostsPostV2HostsSearchParamFlags registers all flags needed to fill params
func registerOperationHostsPostV2HostsSearchParamFlags(cmd *cobra.Command) error {
	if err := registerOperationHostsPostV2HostsSearchQueryParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationHostsPostV2HostsSearchQueryParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var flagQueryName string
	if cmdPrefix == "" {
		flagQueryName = "query"
	} else {
		flagQueryName = fmt.Sprintf("%v.query", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(flagQueryName, "", `Optional json string for [query]. Query Request Payload`)

	// add flags for body
	if err := registerModelCybleComOdinAPIControllersV2IpservicesSearchRequestFlags(0, "request", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationHostsPostV2HostsSearchQueryFlag(m *hosts.PostV2HostsSearchParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("query") {
		// Read query string from cmd and unmarshal
		flagQueryValueStr, err := cmd.Flags().GetString("query")
		if err != nil {
			return err, false
		}

		flagQueryValue := models.CybleComOdinAPIControllersV2IpservicesSearchRequest{}
		if err := json.Unmarshal([]byte(flagQueryValueStr), &flagQueryValue); err != nil {
			return fmt.Errorf("cannot unmarshal query string in models.CybleComOdinAPIControllersV2IpservicesSearchRequest: %v", err), false
		}
		m.Query = &flagQueryValue
	}
	flagQueryModel := m.Query
	if swag.IsZero(flagQueryModel) {
		flagQueryModel = &models.CybleComOdinAPIControllersV2IpservicesSearchRequest{}
	}
	err, added := retrieveModelCybleComOdinAPIControllersV2IpservicesSearchRequestFlags(0, flagQueryModel, "request", cmd)
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

// parseOperationHostsPostV2HostsSearchResult parses request result and return the string content
func parseOperationHostsPostV2HostsSearchResult(resp0 *hosts.PostV2HostsSearchOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*hosts.PostV2HostsSearchOK)
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
		resp1, ok := iResp1.(*hosts.PostV2HostsSearchBadRequest)
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
		resp2, ok := iResp2.(*hosts.PostV2HostsSearchPaymentRequired)
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
		resp3, ok := iResp3.(*hosts.PostV2HostsSearchRequestTimeout)
		if ok {
			if !swag.IsZero(resp3) && !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp4 interface{} = respErr
		resp4, ok := iResp4.(*hosts.PostV2HostsSearchInternalServerError)
		if ok {
			if !swag.IsZero(resp4) && !swag.IsZero(resp4.Payload) {
				msgStr, err := json.Marshal(resp4.Payload)
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
func registerModelPostV2HostsSearchOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	// register embedded models.CybleComOdinAPIControllersV2IpservicesAPIResponse flags

	if err := registerModelCybleComOdinAPIControllersV2IpservicesAPIResponseFlags(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	// register anonymous fields for postV2HostsSearchOKBodyAO1

	if err := registerPostV2HostsSearchOKBodyPropAnonPostV2HostsSearchOKBodyAO1Pagination(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPostV2HostsSearchOKBodyPropAnonPostV2HostsSearchOKBodyAO1Data(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

// inline definition name postV2HostsSearchOKBodyAO1, type

func registerPostV2HostsSearchOKBodyPropAnonPostV2HostsSearchOKBodyAO1Pagination(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var flagPaginationName string
	if cmdPrefix == "" {
		flagPaginationName = " pagination"
	} else {
		flagPaginationName = fmt.Sprintf("%v. pagination", cmdPrefix)
	}

	if err := registerModelCybleComOdinAPIControllersV2IpservicesSearchPaginationFlags(depth+1, flagPaginationName, cmd); err != nil {
		return err
	}

	return nil
}

func registerPostV2HostsSearchOKBodyPropAnonPostV2HostsSearchOKBodyAO1Data(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: data []*models.ServiceService array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPostV2HostsSearchOKBodyFlags(depth int, m *hosts.PostV2HostsSearchOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	// retrieve model models.CybleComOdinAPIControllersV2IpservicesAPIResponse
	err, PostV2HostsSearchOKBodyAO0Added := retrieveModelCybleComOdinAPIControllersV2IpservicesAPIResponseFlags(depth, &m.CybleComOdinAPIControllersV2IpservicesAPIResponse, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || PostV2HostsSearchOKBodyAO0Added

	// retrieve allOf postV2HostsSearchOKBodyAO1 fields

	err, PaginationAdded := retrievePostV2HostsSearchOKBodyPropAnonPostV2HostsSearchOKBodyAO1PaginationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || PaginationAdded

	err, DataAdded := retrievePostV2HostsSearchOKBodyPropAnonPostV2HostsSearchOKBodyAO1DataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DataAdded

	return nil, retAdded
}

// define retrieve functions for fields for inline definition name postV2HostsSearchOKBodyAO1

func retrievePostV2HostsSearchOKBodyPropAnonPostV2HostsSearchOKBodyAO1PaginationFlags(depth int, m *hosts.PostV2HostsSearchOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagPaginationName := fmt.Sprintf("%v. pagination", cmdPrefix)
	if cmd.Flags().Changed(flagPaginationName) {
		// info: complex object  pagination models.CybleComOdinAPIControllersV2IpservicesSearchPagination is retrieved outside this Changed() block
	}
	flagPaginationValue := m.Pagination
	if swag.IsZero(flagPaginationValue) {
		flagPaginationValue = &models.CybleComOdinAPIControllersV2IpservicesSearchPagination{}
	}

	err, PaginationAdded := retrieveModelCybleComOdinAPIControllersV2IpservicesSearchPaginationFlags(depth+1, flagPaginationValue, flagPaginationName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || PaginationAdded
	if PaginationAdded {
		m.Pagination = flagPaginationValue
	}

	return nil, retAdded
}

func retrievePostV2HostsSearchOKBodyPropAnonPostV2HostsSearchOKBodyAO1DataFlags(depth int, m *hosts.PostV2HostsSearchOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagDataName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(flagDataName) {
		// warning: data array type []*models.ServiceService is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
