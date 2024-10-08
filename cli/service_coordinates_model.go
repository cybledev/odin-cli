// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/cybledev/odin-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for ServiceCoordinates

// register flags to command
func registerModelServiceCoordinatesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerServiceCoordinatesPropLatitude(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceCoordinatesPropLongitude(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerServiceCoordinatesPropLatitude(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagLatitudeDescription := ``

	var flagLatitudeName string
	if cmdPrefix == "" {
		flagLatitudeName = "latitude"
	} else {
		flagLatitudeName = fmt.Sprintf("%v.latitude", cmdPrefix)
	}

	var flagLatitudeDefault string

	_ = cmd.PersistentFlags().String(flagLatitudeName, flagLatitudeDefault, flagLatitudeDescription)

	return nil
}

func registerServiceCoordinatesPropLongitude(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagLongitudeDescription := ``

	var flagLongitudeName string
	if cmdPrefix == "" {
		flagLongitudeName = "longitude"
	} else {
		flagLongitudeName = fmt.Sprintf("%v.longitude", cmdPrefix)
	}

	var flagLongitudeDefault string

	_ = cmd.PersistentFlags().String(flagLongitudeName, flagLongitudeDefault, flagLongitudeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelServiceCoordinatesFlags(depth int, m *models.ServiceCoordinates, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, LatitudeAdded := retrieveServiceCoordinatesPropLatitudeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || LatitudeAdded

	err, LongitudeAdded := retrieveServiceCoordinatesPropLongitudeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || LongitudeAdded

	return nil, retAdded
}

func retrieveServiceCoordinatesPropLatitudeFlags(depth int, m *models.ServiceCoordinates, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagLatitudeName := fmt.Sprintf("%v.latitude", cmdPrefix)
	if cmd.Flags().Changed(flagLatitudeName) {

		var flagLatitudeName string
		if cmdPrefix == "" {
			flagLatitudeName = "latitude"
		} else {
			flagLatitudeName = fmt.Sprintf("%v.latitude", cmdPrefix)
		}

		flagLatitudeValue, err := cmd.Flags().GetString(flagLatitudeName)
		if err != nil {
			return err, false
		}
		m.Latitude = flagLatitudeValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServiceCoordinatesPropLongitudeFlags(depth int, m *models.ServiceCoordinates, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagLongitudeName := fmt.Sprintf("%v.longitude", cmdPrefix)
	if cmd.Flags().Changed(flagLongitudeName) {

		var flagLongitudeName string
		if cmdPrefix == "" {
			flagLongitudeName = "longitude"
		} else {
			flagLongitudeName = fmt.Sprintf("%v.longitude", cmdPrefix)
		}

		flagLongitudeValue, err := cmd.Flags().GetString(flagLongitudeName)
		if err != nil {
			return err, false
		}
		m.Longitude = flagLongitudeValue

		retAdded = true
	}

	return nil, retAdded
}
