// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/cybledev/odin-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for ExposedBucket

// register flags to command
func registerModelExposedBucketFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerExposedBucketPropAgeInDays(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExposedBucketPropCategories(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExposedBucketPropDeleted(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExposedBucketPropEfModAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExposedBucketPropFileCatCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExposedBucketPropFiles(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExposedBucketPropInsAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExposedBucketPropLabels(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExposedBucketPropLastFoundAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExposedBucketPropLastOpenAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExposedBucketPropLfModAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExposedBucketPropName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExposedBucketPropNewFiles(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExposedBucketPropOpen(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExposedBucketPropOwnerID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExposedBucketPropPrevScanAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExposedBucketPropProvider(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExposedBucketPropRegion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExposedBucketPropScanAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExposedBucketPropSensitiveFiles(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerExposedBucketPropAgeInDays(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagAgeInDaysDescription := ``

	var flagAgeInDaysName string
	if cmdPrefix == "" {
		flagAgeInDaysName = "age_in_days"
	} else {
		flagAgeInDaysName = fmt.Sprintf("%v.age_in_days", cmdPrefix)
	}

	var flagAgeInDaysDefault int64

	_ = cmd.PersistentFlags().Int64(flagAgeInDaysName, flagAgeInDaysDefault, flagAgeInDaysDescription)

	return nil
}

func registerExposedBucketPropCategories(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: categories []string array type is not supported by go-swagger cli yet

	return nil
}

func registerExposedBucketPropDeleted(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagDeletedDescription := ``

	var flagDeletedName string
	if cmdPrefix == "" {
		flagDeletedName = "deleted"
	} else {
		flagDeletedName = fmt.Sprintf("%v.deleted", cmdPrefix)
	}

	var flagDeletedDefault bool

	_ = cmd.PersistentFlags().Bool(flagDeletedName, flagDeletedDefault, flagDeletedDescription)

	return nil
}

func registerExposedBucketPropEfModAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagEfModAtDescription := ``

	var flagEfModAtName string
	if cmdPrefix == "" {
		flagEfModAtName = "ef_mod_at"
	} else {
		flagEfModAtName = fmt.Sprintf("%v.ef_mod_at", cmdPrefix)
	}

	var flagEfModAtDefault string

	_ = cmd.PersistentFlags().String(flagEfModAtName, flagEfModAtDefault, flagEfModAtDescription)

	return nil
}

func registerExposedBucketPropFileCatCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: file_cat_count map[string]int64 map type is not supported by go-swagger cli yet

	return nil
}

func registerExposedBucketPropFiles(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagFilesDescription := ``

	var flagFilesName string
	if cmdPrefix == "" {
		flagFilesName = "files"
	} else {
		flagFilesName = fmt.Sprintf("%v.files", cmdPrefix)
	}

	var flagFilesDefault int64

	_ = cmd.PersistentFlags().Int64(flagFilesName, flagFilesDefault, flagFilesDescription)

	return nil
}

func registerExposedBucketPropInsAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagInsAtDescription := ``

	var flagInsAtName string
	if cmdPrefix == "" {
		flagInsAtName = "ins_at"
	} else {
		flagInsAtName = fmt.Sprintf("%v.ins_at", cmdPrefix)
	}

	var flagInsAtDefault string

	_ = cmd.PersistentFlags().String(flagInsAtName, flagInsAtDefault, flagInsAtDescription)

	return nil
}

func registerExposedBucketPropLabels(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: labels []string array type is not supported by go-swagger cli yet

	return nil
}

func registerExposedBucketPropLastFoundAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagLastFoundAtDescription := ``

	var flagLastFoundAtName string
	if cmdPrefix == "" {
		flagLastFoundAtName = "last_found_at"
	} else {
		flagLastFoundAtName = fmt.Sprintf("%v.last_found_at", cmdPrefix)
	}

	var flagLastFoundAtDefault string

	_ = cmd.PersistentFlags().String(flagLastFoundAtName, flagLastFoundAtDefault, flagLastFoundAtDescription)

	return nil
}

func registerExposedBucketPropLastOpenAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagLastOpenAtDescription := ``

	var flagLastOpenAtName string
	if cmdPrefix == "" {
		flagLastOpenAtName = "last_open_at"
	} else {
		flagLastOpenAtName = fmt.Sprintf("%v.last_open_at", cmdPrefix)
	}

	var flagLastOpenAtDefault string

	_ = cmd.PersistentFlags().String(flagLastOpenAtName, flagLastOpenAtDefault, flagLastOpenAtDescription)

	return nil
}

func registerExposedBucketPropLfModAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagLfModAtDescription := ``

	var flagLfModAtName string
	if cmdPrefix == "" {
		flagLfModAtName = "lf_mod_at"
	} else {
		flagLfModAtName = fmt.Sprintf("%v.lf_mod_at", cmdPrefix)
	}

	var flagLfModAtDefault string

	_ = cmd.PersistentFlags().String(flagLfModAtName, flagLfModAtDefault, flagLfModAtDescription)

	return nil
}

func registerExposedBucketPropName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagNameDescription := ``

	var flagNameName string
	if cmdPrefix == "" {
		flagNameName = "name"
	} else {
		flagNameName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var flagNameDefault string

	_ = cmd.PersistentFlags().String(flagNameName, flagNameDefault, flagNameDescription)

	return nil
}

func registerExposedBucketPropNewFiles(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagNewFilesDescription := ``

	var flagNewFilesName string
	if cmdPrefix == "" {
		flagNewFilesName = "new_files"
	} else {
		flagNewFilesName = fmt.Sprintf("%v.new_files", cmdPrefix)
	}

	var flagNewFilesDefault int64

	_ = cmd.PersistentFlags().Int64(flagNewFilesName, flagNewFilesDefault, flagNewFilesDescription)

	return nil
}

func registerExposedBucketPropOpen(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagOpenDescription := ``

	var flagOpenName string
	if cmdPrefix == "" {
		flagOpenName = "open"
	} else {
		flagOpenName = fmt.Sprintf("%v.open", cmdPrefix)
	}

	var flagOpenDefault bool

	_ = cmd.PersistentFlags().Bool(flagOpenName, flagOpenDefault, flagOpenDescription)

	return nil
}

func registerExposedBucketPropOwnerID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagOwnerIDDescription := ``

	var flagOwnerIDName string
	if cmdPrefix == "" {
		flagOwnerIDName = "owner_id"
	} else {
		flagOwnerIDName = fmt.Sprintf("%v.owner_id", cmdPrefix)
	}

	var flagOwnerIDDefault string

	_ = cmd.PersistentFlags().String(flagOwnerIDName, flagOwnerIDDefault, flagOwnerIDDescription)

	return nil
}

func registerExposedBucketPropPrevScanAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagPrevScanAtDescription := ``

	var flagPrevScanAtName string
	if cmdPrefix == "" {
		flagPrevScanAtName = "prev_scan_at"
	} else {
		flagPrevScanAtName = fmt.Sprintf("%v.prev_scan_at", cmdPrefix)
	}

	var flagPrevScanAtDefault string

	_ = cmd.PersistentFlags().String(flagPrevScanAtName, flagPrevScanAtDefault, flagPrevScanAtDescription)

	return nil
}

func registerExposedBucketPropProvider(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagProviderDescription := ``

	var flagProviderName string
	if cmdPrefix == "" {
		flagProviderName = "provider"
	} else {
		flagProviderName = fmt.Sprintf("%v.provider", cmdPrefix)
	}

	var flagProviderDefault string

	_ = cmd.PersistentFlags().String(flagProviderName, flagProviderDefault, flagProviderDescription)

	return nil
}

func registerExposedBucketPropRegion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagRegionDescription := ``

	var flagRegionName string
	if cmdPrefix == "" {
		flagRegionName = "region"
	} else {
		flagRegionName = fmt.Sprintf("%v.region", cmdPrefix)
	}

	var flagRegionDefault string

	_ = cmd.PersistentFlags().String(flagRegionName, flagRegionDefault, flagRegionDescription)

	return nil
}

func registerExposedBucketPropScanAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagScanAtDescription := ``

	var flagScanAtName string
	if cmdPrefix == "" {
		flagScanAtName = "scan_at"
	} else {
		flagScanAtName = fmt.Sprintf("%v.scan_at", cmdPrefix)
	}

	var flagScanAtDefault string

	_ = cmd.PersistentFlags().String(flagScanAtName, flagScanAtDefault, flagScanAtDescription)

	return nil
}

func registerExposedBucketPropSensitiveFiles(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	flagSensitiveFilesDescription := ``

	var flagSensitiveFilesName string
	if cmdPrefix == "" {
		flagSensitiveFilesName = "sensitive_files"
	} else {
		flagSensitiveFilesName = fmt.Sprintf("%v.sensitive_files", cmdPrefix)
	}

	var flagSensitiveFilesDefault int64

	_ = cmd.PersistentFlags().Int64(flagSensitiveFilesName, flagSensitiveFilesDefault, flagSensitiveFilesDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelExposedBucketFlags(depth int, m *models.ExposedBucket, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, AgeInDaysAdded := retrieveExposedBucketPropAgeInDaysFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || AgeInDaysAdded

	err, CategoriesAdded := retrieveExposedBucketPropCategoriesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || CategoriesAdded

	err, DeletedAdded := retrieveExposedBucketPropDeletedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || DeletedAdded

	err, EfModAtAdded := retrieveExposedBucketPropEfModAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || EfModAtAdded

	err, FileCatCountAdded := retrieveExposedBucketPropFileCatCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || FileCatCountAdded

	err, FilesAdded := retrieveExposedBucketPropFilesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || FilesAdded

	err, InsAtAdded := retrieveExposedBucketPropInsAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || InsAtAdded

	err, LabelsAdded := retrieveExposedBucketPropLabelsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || LabelsAdded

	err, LastFoundAtAdded := retrieveExposedBucketPropLastFoundAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || LastFoundAtAdded

	err, LastOpenAtAdded := retrieveExposedBucketPropLastOpenAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || LastOpenAtAdded

	err, LfModAtAdded := retrieveExposedBucketPropLfModAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || LfModAtAdded

	err, NameAdded := retrieveExposedBucketPropNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || NameAdded

	err, NewFilesAdded := retrieveExposedBucketPropNewFilesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || NewFilesAdded

	err, OpenAdded := retrieveExposedBucketPropOpenFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || OpenAdded

	err, OwnerIDAdded := retrieveExposedBucketPropOwnerIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || OwnerIDAdded

	err, PrevScanAtAdded := retrieveExposedBucketPropPrevScanAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || PrevScanAtAdded

	err, ProviderAdded := retrieveExposedBucketPropProviderFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ProviderAdded

	err, RegionAdded := retrieveExposedBucketPropRegionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || RegionAdded

	err, ScanAtAdded := retrieveExposedBucketPropScanAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ScanAtAdded

	err, SensitiveFilesAdded := retrieveExposedBucketPropSensitiveFilesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || SensitiveFilesAdded

	return nil, retAdded
}

func retrieveExposedBucketPropAgeInDaysFlags(depth int, m *models.ExposedBucket, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagAgeInDaysName := fmt.Sprintf("%v.age_in_days", cmdPrefix)
	if cmd.Flags().Changed(flagAgeInDaysName) {

		var flagAgeInDaysName string
		if cmdPrefix == "" {
			flagAgeInDaysName = "age_in_days"
		} else {
			flagAgeInDaysName = fmt.Sprintf("%v.age_in_days", cmdPrefix)
		}

		flagAgeInDaysValue, err := cmd.Flags().GetInt64(flagAgeInDaysName)
		if err != nil {
			return err, false
		}
		m.AgeInDays = flagAgeInDaysValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExposedBucketPropCategoriesFlags(depth int, m *models.ExposedBucket, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagCategoriesName := fmt.Sprintf("%v.categories", cmdPrefix)
	if cmd.Flags().Changed(flagCategoriesName) {
		// warning: categories array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveExposedBucketPropDeletedFlags(depth int, m *models.ExposedBucket, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagDeletedName := fmt.Sprintf("%v.deleted", cmdPrefix)
	if cmd.Flags().Changed(flagDeletedName) {

		var flagDeletedName string
		if cmdPrefix == "" {
			flagDeletedName = "deleted"
		} else {
			flagDeletedName = fmt.Sprintf("%v.deleted", cmdPrefix)
		}

		flagDeletedValue, err := cmd.Flags().GetBool(flagDeletedName)
		if err != nil {
			return err, false
		}
		m.Deleted = flagDeletedValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExposedBucketPropEfModAtFlags(depth int, m *models.ExposedBucket, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagEfModAtName := fmt.Sprintf("%v.ef_mod_at", cmdPrefix)
	if cmd.Flags().Changed(flagEfModAtName) {

		var flagEfModAtName string
		if cmdPrefix == "" {
			flagEfModAtName = "ef_mod_at"
		} else {
			flagEfModAtName = fmt.Sprintf("%v.ef_mod_at", cmdPrefix)
		}

		flagEfModAtValue, err := cmd.Flags().GetString(flagEfModAtName)
		if err != nil {
			return err, false
		}
		m.EfModAt = flagEfModAtValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExposedBucketPropFileCatCountFlags(depth int, m *models.ExposedBucket, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagFileCatCountName := fmt.Sprintf("%v.file_cat_count", cmdPrefix)
	if cmd.Flags().Changed(flagFileCatCountName) {
		// warning: file_cat_count map type map[string]int64 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveExposedBucketPropFilesFlags(depth int, m *models.ExposedBucket, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagFilesName := fmt.Sprintf("%v.files", cmdPrefix)
	if cmd.Flags().Changed(flagFilesName) {

		var flagFilesName string
		if cmdPrefix == "" {
			flagFilesName = "files"
		} else {
			flagFilesName = fmt.Sprintf("%v.files", cmdPrefix)
		}

		flagFilesValue, err := cmd.Flags().GetInt64(flagFilesName)
		if err != nil {
			return err, false
		}
		m.Files = flagFilesValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExposedBucketPropInsAtFlags(depth int, m *models.ExposedBucket, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagInsAtName := fmt.Sprintf("%v.ins_at", cmdPrefix)
	if cmd.Flags().Changed(flagInsAtName) {

		var flagInsAtName string
		if cmdPrefix == "" {
			flagInsAtName = "ins_at"
		} else {
			flagInsAtName = fmt.Sprintf("%v.ins_at", cmdPrefix)
		}

		flagInsAtValue, err := cmd.Flags().GetString(flagInsAtName)
		if err != nil {
			return err, false
		}
		m.InsAt = flagInsAtValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExposedBucketPropLabelsFlags(depth int, m *models.ExposedBucket, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagLabelsName := fmt.Sprintf("%v.labels", cmdPrefix)
	if cmd.Flags().Changed(flagLabelsName) {
		// warning: labels array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveExposedBucketPropLastFoundAtFlags(depth int, m *models.ExposedBucket, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagLastFoundAtName := fmt.Sprintf("%v.last_found_at", cmdPrefix)
	if cmd.Flags().Changed(flagLastFoundAtName) {

		var flagLastFoundAtName string
		if cmdPrefix == "" {
			flagLastFoundAtName = "last_found_at"
		} else {
			flagLastFoundAtName = fmt.Sprintf("%v.last_found_at", cmdPrefix)
		}

		flagLastFoundAtValue, err := cmd.Flags().GetString(flagLastFoundAtName)
		if err != nil {
			return err, false
		}
		m.LastFoundAt = flagLastFoundAtValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExposedBucketPropLastOpenAtFlags(depth int, m *models.ExposedBucket, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagLastOpenAtName := fmt.Sprintf("%v.last_open_at", cmdPrefix)
	if cmd.Flags().Changed(flagLastOpenAtName) {

		var flagLastOpenAtName string
		if cmdPrefix == "" {
			flagLastOpenAtName = "last_open_at"
		} else {
			flagLastOpenAtName = fmt.Sprintf("%v.last_open_at", cmdPrefix)
		}

		flagLastOpenAtValue, err := cmd.Flags().GetString(flagLastOpenAtName)
		if err != nil {
			return err, false
		}
		m.LastOpenAt = flagLastOpenAtValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExposedBucketPropLfModAtFlags(depth int, m *models.ExposedBucket, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagLfModAtName := fmt.Sprintf("%v.lf_mod_at", cmdPrefix)
	if cmd.Flags().Changed(flagLfModAtName) {

		var flagLfModAtName string
		if cmdPrefix == "" {
			flagLfModAtName = "lf_mod_at"
		} else {
			flagLfModAtName = fmt.Sprintf("%v.lf_mod_at", cmdPrefix)
		}

		flagLfModAtValue, err := cmd.Flags().GetString(flagLfModAtName)
		if err != nil {
			return err, false
		}
		m.LfModAt = flagLfModAtValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExposedBucketPropNameFlags(depth int, m *models.ExposedBucket, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagNameName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(flagNameName) {

		var flagNameName string
		if cmdPrefix == "" {
			flagNameName = "name"
		} else {
			flagNameName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		flagNameValue, err := cmd.Flags().GetString(flagNameName)
		if err != nil {
			return err, false
		}
		m.Name = flagNameValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExposedBucketPropNewFilesFlags(depth int, m *models.ExposedBucket, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagNewFilesName := fmt.Sprintf("%v.new_files", cmdPrefix)
	if cmd.Flags().Changed(flagNewFilesName) {

		var flagNewFilesName string
		if cmdPrefix == "" {
			flagNewFilesName = "new_files"
		} else {
			flagNewFilesName = fmt.Sprintf("%v.new_files", cmdPrefix)
		}

		flagNewFilesValue, err := cmd.Flags().GetInt64(flagNewFilesName)
		if err != nil {
			return err, false
		}
		m.NewFiles = flagNewFilesValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExposedBucketPropOpenFlags(depth int, m *models.ExposedBucket, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagOpenName := fmt.Sprintf("%v.open", cmdPrefix)
	if cmd.Flags().Changed(flagOpenName) {

		var flagOpenName string
		if cmdPrefix == "" {
			flagOpenName = "open"
		} else {
			flagOpenName = fmt.Sprintf("%v.open", cmdPrefix)
		}

		flagOpenValue, err := cmd.Flags().GetBool(flagOpenName)
		if err != nil {
			return err, false
		}
		m.Open = flagOpenValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExposedBucketPropOwnerIDFlags(depth int, m *models.ExposedBucket, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagOwnerIDName := fmt.Sprintf("%v.owner_id", cmdPrefix)
	if cmd.Flags().Changed(flagOwnerIDName) {

		var flagOwnerIDName string
		if cmdPrefix == "" {
			flagOwnerIDName = "owner_id"
		} else {
			flagOwnerIDName = fmt.Sprintf("%v.owner_id", cmdPrefix)
		}

		flagOwnerIDValue, err := cmd.Flags().GetString(flagOwnerIDName)
		if err != nil {
			return err, false
		}
		m.OwnerID = flagOwnerIDValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExposedBucketPropPrevScanAtFlags(depth int, m *models.ExposedBucket, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagPrevScanAtName := fmt.Sprintf("%v.prev_scan_at", cmdPrefix)
	if cmd.Flags().Changed(flagPrevScanAtName) {

		var flagPrevScanAtName string
		if cmdPrefix == "" {
			flagPrevScanAtName = "prev_scan_at"
		} else {
			flagPrevScanAtName = fmt.Sprintf("%v.prev_scan_at", cmdPrefix)
		}

		flagPrevScanAtValue, err := cmd.Flags().GetString(flagPrevScanAtName)
		if err != nil {
			return err, false
		}
		m.PrevScanAt = flagPrevScanAtValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExposedBucketPropProviderFlags(depth int, m *models.ExposedBucket, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagProviderName := fmt.Sprintf("%v.provider", cmdPrefix)
	if cmd.Flags().Changed(flagProviderName) {

		var flagProviderName string
		if cmdPrefix == "" {
			flagProviderName = "provider"
		} else {
			flagProviderName = fmt.Sprintf("%v.provider", cmdPrefix)
		}

		flagProviderValue, err := cmd.Flags().GetString(flagProviderName)
		if err != nil {
			return err, false
		}
		m.Provider = flagProviderValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExposedBucketPropRegionFlags(depth int, m *models.ExposedBucket, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagRegionName := fmt.Sprintf("%v.region", cmdPrefix)
	if cmd.Flags().Changed(flagRegionName) {

		var flagRegionName string
		if cmdPrefix == "" {
			flagRegionName = "region"
		} else {
			flagRegionName = fmt.Sprintf("%v.region", cmdPrefix)
		}

		flagRegionValue, err := cmd.Flags().GetString(flagRegionName)
		if err != nil {
			return err, false
		}
		m.Region = flagRegionValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExposedBucketPropScanAtFlags(depth int, m *models.ExposedBucket, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagScanAtName := fmt.Sprintf("%v.scan_at", cmdPrefix)
	if cmd.Flags().Changed(flagScanAtName) {

		var flagScanAtName string
		if cmdPrefix == "" {
			flagScanAtName = "scan_at"
		} else {
			flagScanAtName = fmt.Sprintf("%v.scan_at", cmdPrefix)
		}

		flagScanAtValue, err := cmd.Flags().GetString(flagScanAtName)
		if err != nil {
			return err, false
		}
		m.ScanAt = flagScanAtValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExposedBucketPropSensitiveFilesFlags(depth int, m *models.ExposedBucket, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	flagSensitiveFilesName := fmt.Sprintf("%v.sensitive_files", cmdPrefix)
	if cmd.Flags().Changed(flagSensitiveFilesName) {

		var flagSensitiveFilesName string
		if cmdPrefix == "" {
			flagSensitiveFilesName = "sensitive_files"
		} else {
			flagSensitiveFilesName = fmt.Sprintf("%v.sensitive_files", cmdPrefix)
		}

		flagSensitiveFilesValue, err := cmd.Flags().GetInt64(flagSensitiveFilesName)
		if err != nil {
			return err, false
		}
		m.SensitiveFiles = flagSensitiveFilesValue

		retAdded = true
	}

	return nil, retAdded
}
