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

// makeOperationHostsPostV2HostsCountCmd returns a command to handle operation postV2HostsCount
func makeOperationHostsPostV2HostsCountCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "v2-count",
		Short: `Returns the total no of records based on query`,
		RunE:  decorator.RunEColorWrapper(runOperationHostsPostV2HostsCount),
	}

	if err := registerOperationHostsPostV2HostsCountParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationHostsPostV2HostsCount uses cmd flags to call endpoint api
func runOperationHostsPostV2HostsCount(cmd *cobra.Command, args []string) (string, error) {
	args = append(args, apiVersionPrefix + "/v2/")
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return "", err
	}
	// retrieve flag values from cmd and fill params
	params := hosts.NewPostV2HostsCountParams()
	if err, _ = retrieveOperationHostsPostV2HostsCountQueryFlag(params, "", cmd); err != nil {
		return "", err
	}
	if dryRun {
		logDebugf("dry-run flag specified. Skip sending request.")
		return "", nil
	}
	// make request and then print result
	msgStr, err := parseOperationHostsPostV2HostsCountResult(appCli.Hosts.PostV2HostsCount(params, nil))
	if err != nil {
		return "", err
	}

	if !debug {
		return msgStr, nil
	}

	return "", nil
}

// registerOperationHostsPostV2HostsCountParamFlags registers all flags needed to fill params
func registerOperationHostsPostV2HostsCountParamFlags(cmd *cobra.Command) error {
	if err := registerOperationHostsPostV2HostsCountQueryParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationHostsPostV2HostsCountQueryParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var flagQueryName string
	if cmdPrefix == "" {
		flagQueryName = "query"
	} else {
		flagQueryName = fmt.Sprintf("%v.query", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(flagQueryName, "", `Optional json string for [query]. Count Query`)

	// add flags for body
	if err := registerModelCybleComOdinAPIControllersV2IpservicesCountRequestFlags(0, "request", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationHostsPostV2HostsCountQueryFlag(m *hosts.PostV2HostsCountParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("query") {
		// Read query string from cmd and unmarshal
		flagQueryValueStr, err := cmd.Flags().GetString("query")
		if err != nil {
			return err, false
		}

		flagQueryValue := models.CybleComOdinAPIControllersV2IpservicesCountRequest{}
		if err := json.Unmarshal([]byte(flagQueryValueStr), &flagQueryValue); err != nil {
			return fmt.Errorf("cannot unmarshal query string in models.CybleComOdinAPIControllersV2IpservicesCountRequest: %v", err), false
		}
		m.Query = &flagQueryValue
	}
	flagQueryModel := m.Query
	if swag.IsZero(flagQueryModel) {
		flagQueryModel = &models.CybleComOdinAPIControllersV2IpservicesCountRequest{}
	}
	err, added := retrieveModelCybleComOdinAPIControllersV2IpservicesCountRequestFlags(0, flagQueryModel, "request", cmd)
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

// parseOperationHostsPostV2HostsCountResult parses request result and return the string content
func parseOperationHostsPostV2HostsCountResult(resp0 *hosts.PostV2HostsCountOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*hosts.PostV2HostsCountOK)
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
		resp1, ok := iResp1.(*hosts.PostV2HostsCountBadRequest)
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
		resp2, ok := iResp2.(*hosts.PostV2HostsCountPaymentRequired)
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
		resp3, ok := iResp3.(*hosts.PostV2HostsCountRequestTimeout)
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
		resp4, ok := iResp4.(*hosts.PostV2HostsCountInternalServerError)
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
func registerModelPostV2HostsCountOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	// register embedded models.CybleComOdinAPIControllersV2IpservicesAPIResponse flags

	if err := registerModelCybleComOdinAPIControllersV2IpservicesAPIResponseFlags(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	// register anonymous fields for postV2HostsCountOKBodyAO1

	if err := registerPostV2HostsCountOKBodyPropAnonPostV2HostsCountOKBodyAO1Data(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

// inline definition name postV2HostsCountOKBodyAO1, type

func registerPostV2HostsCountOKBodyPropAnonPostV2HostsCountOKBodyAO1Data(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var flagDataName string
	if cmdPrefix == "" {
		flagDataName = "data"
	} else {
		flagDataName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelCybleComOdinAPIControllersV2IpservicesCertCountFlags(depth+1, flagDataName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPostV2HostsCountOKBodyFlags(depth int, m *hosts.PostV2HostsCountOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	// retrieve model models.CybleComOdinAPIControllersV2IpservicesAPIResponse
	err, PostV2HostsCountOKBodyAO0Added := retrieveModelCybleComOdinAPIControllersV2IpservicesAPIResponseFlags(depth, &m.CybleComOdinAPIControllersV2IpservicesAPIResponse, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || PostV2HostsCountOKBodyAO0Added

	// retrieve allOf postV2HostsCountOKBodyAO1 fields

	err, DataAdded := retrievePostV2HostsCountOKBodyPropAnonPostV2HostsCountOKBodyAO1DataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DataAdded

	return nil, retAdded
}

// define retrieve functions for fields for inline definition name postV2HostsCountOKBodyAO1

func retrievePostV2HostsCountOKBodyPropAnonPostV2HostsCountOKBodyAO1DataFlags(depth int, m *hosts.PostV2HostsCountOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagDataName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(flagDataName) {
		// info: complex object data models.CybleComOdinAPIControllersV2IpservicesCertCount is retrieved outside this Changed() block
	}
	flagDataValue := m.Data
	if swag.IsZero(flagDataValue) {
		flagDataValue = &models.CybleComOdinAPIControllersV2IpservicesCertCount{}
	}

	err, DataAdded := retrieveModelCybleComOdinAPIControllersV2IpservicesCertCountFlags(depth+1, flagDataValue, flagDataName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DataAdded
	if DataAdded {
		m.Data = flagDataValue
	}

	return nil, retAdded
}
