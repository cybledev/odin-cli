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

// Schema cli for CertificateCertificateSearchData

// register flags to command
func registerModelCertificateCertificateSearchDataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCertificateCertificateSearchDataPropFingerprintMd5(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCertificateCertificateSearchDataPropFingerprintSha1(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCertificateCertificateSearchDataPropFingerprintSha256(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCertificateCertificateSearchDataPropIssuer(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCertificateCertificateSearchDataPropSubject(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCertificateCertificateSearchDataPropSubjectAltName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCertificateCertificateSearchDataPropTags(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCertificateCertificateSearchDataPropValidity(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCertificateCertificateSearchDataPropFingerprintMd5(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagFingerprintMd5Description := ``

	var flagFingerprintMd5Name string
	if cmdPrefix == "" {
		flagFingerprintMd5Name = "fingerprint_md5"
	} else {
		flagFingerprintMd5Name = fmt.Sprintf("%v.fingerprint_md5", cmdPrefix)
	}

	var flagFingerprintMd5Default string

	_ = cmd.PersistentFlags().String(flagFingerprintMd5Name, flagFingerprintMd5Default, flagFingerprintMd5Description)

	return nil
}

func registerCertificateCertificateSearchDataPropFingerprintSha1(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagFingerprintSha1Description := ``

	var flagFingerprintSha1Name string
	if cmdPrefix == "" {
		flagFingerprintSha1Name = "fingerprint_sha1"
	} else {
		flagFingerprintSha1Name = fmt.Sprintf("%v.fingerprint_sha1", cmdPrefix)
	}

	var flagFingerprintSha1Default string

	_ = cmd.PersistentFlags().String(flagFingerprintSha1Name, flagFingerprintSha1Default, flagFingerprintSha1Description)

	return nil
}

func registerCertificateCertificateSearchDataPropFingerprintSha256(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagFingerprintSha256Description := ``

	var flagFingerprintSha256Name string
	if cmdPrefix == "" {
		flagFingerprintSha256Name = "fingerprint_sha256"
	} else {
		flagFingerprintSha256Name = fmt.Sprintf("%v.fingerprint_sha256", cmdPrefix)
	}

	var flagFingerprintSha256Default string

	_ = cmd.PersistentFlags().String(flagFingerprintSha256Name, flagFingerprintSha256Default, flagFingerprintSha256Description)

	return nil
}

func registerCertificateCertificateSearchDataPropIssuer(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var flagIssuerName string
	if cmdPrefix == "" {
		flagIssuerName = "issuer"
	} else {
		flagIssuerName = fmt.Sprintf("%v.issuer", cmdPrefix)
	}

	if err := registerModelCertificateCertificateSearchDataIssuerFlags(depth+1, flagIssuerName, cmd); err != nil {
		return err
	}

	return nil
}

func registerCertificateCertificateSearchDataPropSubject(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var flagSubjectName string
	if cmdPrefix == "" {
		flagSubjectName = "subject"
	} else {
		flagSubjectName = fmt.Sprintf("%v.subject", cmdPrefix)
	}

	if err := registerModelCertificateCertificateSearchDataSubjectFlags(depth+1, flagSubjectName, cmd); err != nil {
		return err
	}

	return nil
}

func registerCertificateCertificateSearchDataPropSubjectAltName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var flagSubjectAltNameName string
	if cmdPrefix == "" {
		flagSubjectAltNameName = "subject_alt_name"
	} else {
		flagSubjectAltNameName = fmt.Sprintf("%v.subject_alt_name", cmdPrefix)
	}

	if err := registerModelCertificateCertificateSearchDataSubjectAltNameFlags(depth+1, flagSubjectAltNameName, cmd); err != nil {
		return err
	}

	return nil
}

func registerCertificateCertificateSearchDataPropTags(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: tags []string array type is not supported by go-swagger cli yet

	return nil
}

func registerCertificateCertificateSearchDataPropValidity(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var flagValidityName string
	if cmdPrefix == "" {
		flagValidityName = "validity"
	} else {
		flagValidityName = fmt.Sprintf("%v.validity", cmdPrefix)
	}

	if err := registerModelCertificateCertificateSearchDataValidityFlags(depth+1, flagValidityName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCertificateCertificateSearchDataFlags(depth int, m *models.CertificateCertificateSearchData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, FingerprintMd5Added := retrieveCertificateCertificateSearchDataPropFingerprintMd5Flags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || FingerprintMd5Added

	err, FingerprintSha1Added := retrieveCertificateCertificateSearchDataPropFingerprintSha1Flags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || FingerprintSha1Added

	err, FingerprintSha256Added := retrieveCertificateCertificateSearchDataPropFingerprintSha256Flags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || FingerprintSha256Added

	err, IssuerAdded := retrieveCertificateCertificateSearchDataPropIssuerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || IssuerAdded

	err, SubjectAdded := retrieveCertificateCertificateSearchDataPropSubjectFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || SubjectAdded

	err, SubjectAltNameAdded := retrieveCertificateCertificateSearchDataPropSubjectAltNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || SubjectAltNameAdded

	err, TagsAdded := retrieveCertificateCertificateSearchDataPropTagsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || TagsAdded

	err, ValidityAdded := retrieveCertificateCertificateSearchDataPropValidityFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ValidityAdded

	return nil, retAdded
}

func retrieveCertificateCertificateSearchDataPropFingerprintMd5Flags(depth int, m *models.CertificateCertificateSearchData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagFingerprintMd5Name := fmt.Sprintf("%v.fingerprint_md5", cmdPrefix)
	if cmd.Flags().Changed(flagFingerprintMd5Name) {

		var flagFingerprintMd5Name string
		if cmdPrefix == "" {
			flagFingerprintMd5Name = "fingerprint_md5"
		} else {
			flagFingerprintMd5Name = fmt.Sprintf("%v.fingerprint_md5", cmdPrefix)
		}

		flagFingerprintMd5Value, err := cmd.Flags().GetString(flagFingerprintMd5Name)
		if err != nil {
			return err, false
		}
		m.FingerprintMd5 = flagFingerprintMd5Value

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCertificateCertificateSearchDataPropFingerprintSha1Flags(depth int, m *models.CertificateCertificateSearchData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagFingerprintSha1Name := fmt.Sprintf("%v.fingerprint_sha1", cmdPrefix)
	if cmd.Flags().Changed(flagFingerprintSha1Name) {

		var flagFingerprintSha1Name string
		if cmdPrefix == "" {
			flagFingerprintSha1Name = "fingerprint_sha1"
		} else {
			flagFingerprintSha1Name = fmt.Sprintf("%v.fingerprint_sha1", cmdPrefix)
		}

		flagFingerprintSha1Value, err := cmd.Flags().GetString(flagFingerprintSha1Name)
		if err != nil {
			return err, false
		}
		m.FingerprintSha1 = flagFingerprintSha1Value

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCertificateCertificateSearchDataPropFingerprintSha256Flags(depth int, m *models.CertificateCertificateSearchData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagFingerprintSha256Name := fmt.Sprintf("%v.fingerprint_sha256", cmdPrefix)
	if cmd.Flags().Changed(flagFingerprintSha256Name) {

		var flagFingerprintSha256Name string
		if cmdPrefix == "" {
			flagFingerprintSha256Name = "fingerprint_sha256"
		} else {
			flagFingerprintSha256Name = fmt.Sprintf("%v.fingerprint_sha256", cmdPrefix)
		}

		flagFingerprintSha256Value, err := cmd.Flags().GetString(flagFingerprintSha256Name)
		if err != nil {
			return err, false
		}
		m.FingerprintSha256 = flagFingerprintSha256Value

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCertificateCertificateSearchDataPropIssuerFlags(depth int, m *models.CertificateCertificateSearchData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagIssuerName := fmt.Sprintf("%v.issuer", cmdPrefix)
	if cmd.Flags().Changed(flagIssuerName) {
		// info: complex object issuer CertificateCertificateSearchDataIssuer is retrieved outside this Changed() block
	}
	flagIssuerValue := m.Issuer
	if swag.IsZero(flagIssuerValue) {
		flagIssuerValue = &models.CertificateCertificateSearchDataIssuer{}
	}

	err, IssuerAdded := retrieveModelCertificateCertificateSearchDataIssuerFlags(depth+1, flagIssuerValue, flagIssuerName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || IssuerAdded
	if IssuerAdded {
		m.Issuer = flagIssuerValue
	}

	return nil, retAdded
}

func retrieveCertificateCertificateSearchDataPropSubjectFlags(depth int, m *models.CertificateCertificateSearchData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagSubjectName := fmt.Sprintf("%v.subject", cmdPrefix)
	if cmd.Flags().Changed(flagSubjectName) {
		// info: complex object subject CertificateCertificateSearchDataSubject is retrieved outside this Changed() block
	}
	flagSubjectValue := m.Subject
	if swag.IsZero(flagSubjectValue) {
		flagSubjectValue = &models.CertificateCertificateSearchDataSubject{}
	}

	err, SubjectAdded := retrieveModelCertificateCertificateSearchDataSubjectFlags(depth+1, flagSubjectValue, flagSubjectName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || SubjectAdded
	if SubjectAdded {
		m.Subject = flagSubjectValue
	}

	return nil, retAdded
}

func retrieveCertificateCertificateSearchDataPropSubjectAltNameFlags(depth int, m *models.CertificateCertificateSearchData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagSubjectAltNameName := fmt.Sprintf("%v.subject_alt_name", cmdPrefix)
	if cmd.Flags().Changed(flagSubjectAltNameName) {
		// info: complex object subject_alt_name CertificateCertificateSearchDataSubjectAltName is retrieved outside this Changed() block
	}
	flagSubjectAltNameValue := m.SubjectAltName
	if swag.IsZero(flagSubjectAltNameValue) {
		flagSubjectAltNameValue = &models.CertificateCertificateSearchDataSubjectAltName{}
	}

	err, SubjectAltNameAdded := retrieveModelCertificateCertificateSearchDataSubjectAltNameFlags(depth+1, flagSubjectAltNameValue, flagSubjectAltNameName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || SubjectAltNameAdded
	if SubjectAltNameAdded {
		m.SubjectAltName = flagSubjectAltNameValue
	}

	return nil, retAdded
}

func retrieveCertificateCertificateSearchDataPropTagsFlags(depth int, m *models.CertificateCertificateSearchData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagTagsName := fmt.Sprintf("%v.tags", cmdPrefix)
	if cmd.Flags().Changed(flagTagsName) {
		// warning: tags array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveCertificateCertificateSearchDataPropValidityFlags(depth int, m *models.CertificateCertificateSearchData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagValidityName := fmt.Sprintf("%v.validity", cmdPrefix)
	if cmd.Flags().Changed(flagValidityName) {
		// info: complex object validity CertificateCertificateSearchDataValidity is retrieved outside this Changed() block
	}
	flagValidityValue := m.Validity
	if swag.IsZero(flagValidityValue) {
		flagValidityValue = &models.CertificateCertificateSearchDataValidity{}
	}

	err, ValidityAdded := retrieveModelCertificateCertificateSearchDataValidityFlags(depth+1, flagValidityValue, flagValidityName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ValidityAdded
	if ValidityAdded {
		m.Validity = flagValidityValue
	}

	return nil, retAdded
}

// Extra schema cli for CertificateCertificateSearchDataIssuer

// register flags to command
func registerModelCertificateCertificateSearchDataIssuerFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCertificateCertificateSearchDataIssuerPropCommonName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCertificateCertificateSearchDataIssuerPropCountry(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCertificateCertificateSearchDataIssuerPropOrganization(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCertificateCertificateSearchDataIssuerPropCommonName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: common_name []string array type is not supported by go-swagger cli yet

	return nil
}

func registerCertificateCertificateSearchDataIssuerPropCountry(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: country []string array type is not supported by go-swagger cli yet

	return nil
}

func registerCertificateCertificateSearchDataIssuerPropOrganization(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: organization []string array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCertificateCertificateSearchDataIssuerFlags(depth int, m *models.CertificateCertificateSearchDataIssuer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, CommonNameAdded := retrieveCertificateCertificateSearchDataIssuerPropCommonNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || CommonNameAdded

	err, CountryAdded := retrieveCertificateCertificateSearchDataIssuerPropCountryFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || CountryAdded

	err, OrganizationAdded := retrieveCertificateCertificateSearchDataIssuerPropOrganizationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || OrganizationAdded

	return nil, retAdded
}

func retrieveCertificateCertificateSearchDataIssuerPropCommonNameFlags(depth int, m *models.CertificateCertificateSearchDataIssuer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagCommonNameName := fmt.Sprintf("%v.common_name", cmdPrefix)
	if cmd.Flags().Changed(flagCommonNameName) {
		// warning: common_name array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveCertificateCertificateSearchDataIssuerPropCountryFlags(depth int, m *models.CertificateCertificateSearchDataIssuer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagCountryName := fmt.Sprintf("%v.country", cmdPrefix)
	if cmd.Flags().Changed(flagCountryName) {
		// warning: country array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveCertificateCertificateSearchDataIssuerPropOrganizationFlags(depth int, m *models.CertificateCertificateSearchDataIssuer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagOrganizationName := fmt.Sprintf("%v.organization", cmdPrefix)
	if cmd.Flags().Changed(flagOrganizationName) {
		// warning: organization array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

// Extra schema cli for CertificateCertificateSearchDataSubject

// register flags to command
func registerModelCertificateCertificateSearchDataSubjectFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCertificateCertificateSearchDataSubjectPropCommonName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCertificateCertificateSearchDataSubjectPropCommonName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: common_name []string array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCertificateCertificateSearchDataSubjectFlags(depth int, m *models.CertificateCertificateSearchDataSubject, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, CommonNameAdded := retrieveCertificateCertificateSearchDataSubjectPropCommonNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || CommonNameAdded

	return nil, retAdded
}

func retrieveCertificateCertificateSearchDataSubjectPropCommonNameFlags(depth int, m *models.CertificateCertificateSearchDataSubject, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagCommonNameName := fmt.Sprintf("%v.common_name", cmdPrefix)
	if cmd.Flags().Changed(flagCommonNameName) {
		// warning: common_name array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

// Extra schema cli for CertificateCertificateSearchDataSubjectAltName

// register flags to command
func registerModelCertificateCertificateSearchDataSubjectAltNameFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCertificateCertificateSearchDataSubjectAltNamePropDNSNames(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCertificateCertificateSearchDataSubjectAltNamePropDNSNames(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: dns_names []string array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCertificateCertificateSearchDataSubjectAltNameFlags(depth int, m *models.CertificateCertificateSearchDataSubjectAltName, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, DNSNamesAdded := retrieveCertificateCertificateSearchDataSubjectAltNamePropDNSNamesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DNSNamesAdded

	return nil, retAdded
}

func retrieveCertificateCertificateSearchDataSubjectAltNamePropDNSNamesFlags(depth int, m *models.CertificateCertificateSearchDataSubjectAltName, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagDNSNamesName := fmt.Sprintf("%v.dns_names", cmdPrefix)
	if cmd.Flags().Changed(flagDNSNamesName) {
		// warning: dns_names array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

// Extra schema cli for CertificateCertificateSearchDataValidity

// register flags to command
func registerModelCertificateCertificateSearchDataValidityFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCertificateCertificateSearchDataValidityPropEnd(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCertificateCertificateSearchDataValidityPropLength(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCertificateCertificateSearchDataValidityPropStart(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCertificateCertificateSearchDataValidityPropEnd(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagEndDescription := ``

	var flagEndName string
	if cmdPrefix == "" {
		flagEndName = "end"
	} else {
		flagEndName = fmt.Sprintf("%v.end", cmdPrefix)
	}

	var flagEndDefault string

	_ = cmd.PersistentFlags().String(flagEndName, flagEndDefault, flagEndDescription)

	return nil
}

func registerCertificateCertificateSearchDataValidityPropLength(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagLengthDescription := ``

	var flagLengthName string
	if cmdPrefix == "" {
		flagLengthName = "length"
	} else {
		flagLengthName = fmt.Sprintf("%v.length", cmdPrefix)
	}

	var flagLengthDefault int64

	_ = cmd.PersistentFlags().Int64(flagLengthName, flagLengthDefault, flagLengthDescription)

	return nil
}

func registerCertificateCertificateSearchDataValidityPropStart(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagStartDescription := ``

	var flagStartName string
	if cmdPrefix == "" {
		flagStartName = "start"
	} else {
		flagStartName = fmt.Sprintf("%v.start", cmdPrefix)
	}

	var flagStartDefault string

	_ = cmd.PersistentFlags().String(flagStartName, flagStartDefault, flagStartDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCertificateCertificateSearchDataValidityFlags(depth int, m *models.CertificateCertificateSearchDataValidity, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, EndAdded := retrieveCertificateCertificateSearchDataValidityPropEndFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || EndAdded

	err, LengthAdded := retrieveCertificateCertificateSearchDataValidityPropLengthFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || LengthAdded

	err, StartAdded := retrieveCertificateCertificateSearchDataValidityPropStartFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || StartAdded

	return nil, retAdded
}

func retrieveCertificateCertificateSearchDataValidityPropEndFlags(depth int, m *models.CertificateCertificateSearchDataValidity, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagEndName := fmt.Sprintf("%v.end", cmdPrefix)
	if cmd.Flags().Changed(flagEndName) {

		var flagEndName string
		if cmdPrefix == "" {
			flagEndName = "end"
		} else {
			flagEndName = fmt.Sprintf("%v.end", cmdPrefix)
		}

		flagEndValue, err := cmd.Flags().GetString(flagEndName)
		if err != nil {
			return err, false
		}
		m.End = flagEndValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCertificateCertificateSearchDataValidityPropLengthFlags(depth int, m *models.CertificateCertificateSearchDataValidity, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagLengthName := fmt.Sprintf("%v.length", cmdPrefix)
	if cmd.Flags().Changed(flagLengthName) {

		var flagLengthName string
		if cmdPrefix == "" {
			flagLengthName = "length"
		} else {
			flagLengthName = fmt.Sprintf("%v.length", cmdPrefix)
		}

		flagLengthValue, err := cmd.Flags().GetInt64(flagLengthName)
		if err != nil {
			return err, false
		}
		m.Length = flagLengthValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCertificateCertificateSearchDataValidityPropStartFlags(depth int, m *models.CertificateCertificateSearchDataValidity, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagStartName := fmt.Sprintf("%v.start", cmdPrefix)
	if cmd.Flags().Changed(flagStartName) {

		var flagStartName string
		if cmdPrefix == "" {
			flagStartName = "start"
		} else {
			flagStartName = fmt.Sprintf("%v.start", cmdPrefix)
		}

		flagStartValue, err := cmd.Flags().GetString(flagStartName)
		if err != nil {
			return err, false
		}
		m.Start = flagStartValue

		retAdded = true
	}

	return nil, retAdded
}
