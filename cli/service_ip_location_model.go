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

// Schema cli for ServiceIPLocation

// register flags to command
func registerModelServiceIPLocationFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerServiceIPLocationPropCity(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceIPLocationPropContinent(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceIPLocationPropCoordinates(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceIPLocationPropCountryCode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceIPLocationPropCountryName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceIPLocationPropGeoPoint(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceIPLocationPropLocaleCode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceIPLocationPropNetwork(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceIPLocationPropPostalCode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerServiceIPLocationPropCity(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagCityDescription := ``

	var flagCityName string
	if cmdPrefix == "" {
		flagCityName = "city"
	} else {
		flagCityName = fmt.Sprintf("%v.city", cmdPrefix)
	}

	var flagCityDefault string

	_ = cmd.PersistentFlags().String(flagCityName, flagCityDefault, flagCityDescription)

	return nil
}

func registerServiceIPLocationPropContinent(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagContinentDescription := ``

	var flagContinentName string
	if cmdPrefix == "" {
		flagContinentName = "continent"
	} else {
		flagContinentName = fmt.Sprintf("%v.continent", cmdPrefix)
	}

	var flagContinentDefault string

	_ = cmd.PersistentFlags().String(flagContinentName, flagContinentDefault, flagContinentDescription)

	return nil
}

func registerServiceIPLocationPropCoordinates(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var flagCoordinatesName string
	if cmdPrefix == "" {
		flagCoordinatesName = "coordinates"
	} else {
		flagCoordinatesName = fmt.Sprintf("%v.coordinates", cmdPrefix)
	}

	if err := registerModelServiceCoordinatesFlags(depth+1, flagCoordinatesName, cmd); err != nil {
		return err
	}

	return nil
}

func registerServiceIPLocationPropCountryCode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagCountryCodeDescription := ``

	var flagCountryCodeName string
	if cmdPrefix == "" {
		flagCountryCodeName = "country_code"
	} else {
		flagCountryCodeName = fmt.Sprintf("%v.country_code", cmdPrefix)
	}

	var flagCountryCodeDefault string

	_ = cmd.PersistentFlags().String(flagCountryCodeName, flagCountryCodeDefault, flagCountryCodeDescription)

	return nil
}

func registerServiceIPLocationPropCountryName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagCountryNameDescription := ``

	var flagCountryNameName string
	if cmdPrefix == "" {
		flagCountryNameName = "country_name"
	} else {
		flagCountryNameName = fmt.Sprintf("%v.country_name", cmdPrefix)
	}

	var flagCountryNameDefault string

	_ = cmd.PersistentFlags().String(flagCountryNameName, flagCountryNameDefault, flagCountryNameDescription)

	return nil
}

func registerServiceIPLocationPropGeoPoint(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagGeoPointDescription := ``

	var flagGeoPointName string
	if cmdPrefix == "" {
		flagGeoPointName = "geo_point"
	} else {
		flagGeoPointName = fmt.Sprintf("%v.geo_point", cmdPrefix)
	}

	var flagGeoPointDefault string

	_ = cmd.PersistentFlags().String(flagGeoPointName, flagGeoPointDefault, flagGeoPointDescription)

	return nil
}

func registerServiceIPLocationPropLocaleCode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagLocaleCodeDescription := ``

	var flagLocaleCodeName string
	if cmdPrefix == "" {
		flagLocaleCodeName = "locale_code"
	} else {
		flagLocaleCodeName = fmt.Sprintf("%v.locale_code", cmdPrefix)
	}

	var flagLocaleCodeDefault string

	_ = cmd.PersistentFlags().String(flagLocaleCodeName, flagLocaleCodeDefault, flagLocaleCodeDescription)

	return nil
}

func registerServiceIPLocationPropNetwork(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagNetworkDescription := ``

	var flagNetworkName string
	if cmdPrefix == "" {
		flagNetworkName = "network"
	} else {
		flagNetworkName = fmt.Sprintf("%v.network", cmdPrefix)
	}

	var flagNetworkDefault string

	_ = cmd.PersistentFlags().String(flagNetworkName, flagNetworkDefault, flagNetworkDescription)

	return nil
}

func registerServiceIPLocationPropPostalCode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagPostalCodeDescription := ``

	var flagPostalCodeName string
	if cmdPrefix == "" {
		flagPostalCodeName = "postal_code"
	} else {
		flagPostalCodeName = fmt.Sprintf("%v.postal_code", cmdPrefix)
	}

	var flagPostalCodeDefault string

	_ = cmd.PersistentFlags().String(flagPostalCodeName, flagPostalCodeDefault, flagPostalCodeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelServiceIPLocationFlags(depth int, m *models.ServiceIPLocation, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, CityAdded := retrieveServiceIPLocationPropCityFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || CityAdded

	err, ContinentAdded := retrieveServiceIPLocationPropContinentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ContinentAdded

	err, CoordinatesAdded := retrieveServiceIPLocationPropCoordinatesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || CoordinatesAdded

	err, CountryCodeAdded := retrieveServiceIPLocationPropCountryCodeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || CountryCodeAdded

	err, CountryNameAdded := retrieveServiceIPLocationPropCountryNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || CountryNameAdded

	err, GeoPointAdded := retrieveServiceIPLocationPropGeoPointFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || GeoPointAdded

	err, LocaleCodeAdded := retrieveServiceIPLocationPropLocaleCodeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || LocaleCodeAdded

	err, NetworkAdded := retrieveServiceIPLocationPropNetworkFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || NetworkAdded

	err, PostalCodeAdded := retrieveServiceIPLocationPropPostalCodeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || PostalCodeAdded

	return nil, retAdded
}

func retrieveServiceIPLocationPropCityFlags(depth int, m *models.ServiceIPLocation, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagCityName := fmt.Sprintf("%v.city", cmdPrefix)
	if cmd.Flags().Changed(flagCityName) {

		var flagCityName string
		if cmdPrefix == "" {
			flagCityName = "city"
		} else {
			flagCityName = fmt.Sprintf("%v.city", cmdPrefix)
		}

		flagCityValue, err := cmd.Flags().GetString(flagCityName)
		if err != nil {
			return err, false
		}
		m.City = flagCityValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServiceIPLocationPropContinentFlags(depth int, m *models.ServiceIPLocation, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagContinentName := fmt.Sprintf("%v.continent", cmdPrefix)
	if cmd.Flags().Changed(flagContinentName) {

		var flagContinentName string
		if cmdPrefix == "" {
			flagContinentName = "continent"
		} else {
			flagContinentName = fmt.Sprintf("%v.continent", cmdPrefix)
		}

		flagContinentValue, err := cmd.Flags().GetString(flagContinentName)
		if err != nil {
			return err, false
		}
		m.Continent = flagContinentValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServiceIPLocationPropCoordinatesFlags(depth int, m *models.ServiceIPLocation, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagCoordinatesName := fmt.Sprintf("%v.coordinates", cmdPrefix)
	if cmd.Flags().Changed(flagCoordinatesName) {
		// info: complex object coordinates ServiceCoordinates is retrieved outside this Changed() block
	}
	flagCoordinatesValue := m.Coordinates
	if swag.IsZero(flagCoordinatesValue) {
		flagCoordinatesValue = &models.ServiceCoordinates{}
	}

	err, CoordinatesAdded := retrieveModelServiceCoordinatesFlags(depth+1, flagCoordinatesValue, flagCoordinatesName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || CoordinatesAdded
	if CoordinatesAdded {
		m.Coordinates = flagCoordinatesValue
	}

	return nil, retAdded
}

func retrieveServiceIPLocationPropCountryCodeFlags(depth int, m *models.ServiceIPLocation, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagCountryCodeName := fmt.Sprintf("%v.country_code", cmdPrefix)
	if cmd.Flags().Changed(flagCountryCodeName) {

		var flagCountryCodeName string
		if cmdPrefix == "" {
			flagCountryCodeName = "country_code"
		} else {
			flagCountryCodeName = fmt.Sprintf("%v.country_code", cmdPrefix)
		}

		flagCountryCodeValue, err := cmd.Flags().GetString(flagCountryCodeName)
		if err != nil {
			return err, false
		}
		m.CountryCode = flagCountryCodeValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServiceIPLocationPropCountryNameFlags(depth int, m *models.ServiceIPLocation, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagCountryNameName := fmt.Sprintf("%v.country_name", cmdPrefix)
	if cmd.Flags().Changed(flagCountryNameName) {

		var flagCountryNameName string
		if cmdPrefix == "" {
			flagCountryNameName = "country_name"
		} else {
			flagCountryNameName = fmt.Sprintf("%v.country_name", cmdPrefix)
		}

		flagCountryNameValue, err := cmd.Flags().GetString(flagCountryNameName)
		if err != nil {
			return err, false
		}
		m.CountryName = flagCountryNameValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServiceIPLocationPropGeoPointFlags(depth int, m *models.ServiceIPLocation, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagGeoPointName := fmt.Sprintf("%v.geo_point", cmdPrefix)
	if cmd.Flags().Changed(flagGeoPointName) {

		var flagGeoPointName string
		if cmdPrefix == "" {
			flagGeoPointName = "geo_point"
		} else {
			flagGeoPointName = fmt.Sprintf("%v.geo_point", cmdPrefix)
		}

		flagGeoPointValue, err := cmd.Flags().GetString(flagGeoPointName)
		if err != nil {
			return err, false
		}
		m.GeoPoint = flagGeoPointValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServiceIPLocationPropLocaleCodeFlags(depth int, m *models.ServiceIPLocation, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagLocaleCodeName := fmt.Sprintf("%v.locale_code", cmdPrefix)
	if cmd.Flags().Changed(flagLocaleCodeName) {

		var flagLocaleCodeName string
		if cmdPrefix == "" {
			flagLocaleCodeName = "locale_code"
		} else {
			flagLocaleCodeName = fmt.Sprintf("%v.locale_code", cmdPrefix)
		}

		flagLocaleCodeValue, err := cmd.Flags().GetString(flagLocaleCodeName)
		if err != nil {
			return err, false
		}
		m.LocaleCode = flagLocaleCodeValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServiceIPLocationPropNetworkFlags(depth int, m *models.ServiceIPLocation, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagNetworkName := fmt.Sprintf("%v.network", cmdPrefix)
	if cmd.Flags().Changed(flagNetworkName) {

		var flagNetworkName string
		if cmdPrefix == "" {
			flagNetworkName = "network"
		} else {
			flagNetworkName = fmt.Sprintf("%v.network", cmdPrefix)
		}

		flagNetworkValue, err := cmd.Flags().GetString(flagNetworkName)
		if err != nil {
			return err, false
		}
		m.Network = flagNetworkValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServiceIPLocationPropPostalCodeFlags(depth int, m *models.ServiceIPLocation, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagPostalCodeName := fmt.Sprintf("%v.postal_code", cmdPrefix)
	if cmd.Flags().Changed(flagPostalCodeName) {

		var flagPostalCodeName string
		if cmdPrefix == "" {
			flagPostalCodeName = "postal_code"
		} else {
			flagPostalCodeName = fmt.Sprintf("%v.postal_code", cmdPrefix)
		}

		flagPostalCodeValue, err := cmd.Flags().GetString(flagPostalCodeName)
		if err != nil {
			return err, false
		}
		m.PostalCode = flagPostalCodeValue

		retAdded = true
	}

	return nil, retAdded
}
