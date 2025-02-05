// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceHealthcareFhirStore() *schema.Resource {
	return &schema.Resource{
		Create: resourceHealthcareFhirStoreCreate,
		Read:   resourceHealthcareFhirStoreRead,
		Update: resourceHealthcareFhirStoreUpdate,
		Delete: resourceHealthcareFhirStoreDelete,

		Importer: &schema.ResourceImporter{
			State: resourceHealthcareFhirStoreImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(240 * time.Second),
			Update: schema.DefaultTimeout(240 * time.Second),
			Delete: schema.DefaultTimeout(240 * time.Second),
		},

		Schema: map[string]*schema.Schema{
			"dataset": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"disable_referential_integrity": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"disable_resource_versioning": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"enable_history_import": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"enable_update_create": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"notification_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pubsub_topic": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceHealthcareFhirStoreCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandHealthcareFhirStoreName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	enableUpdateCreateProp, err := expandHealthcareFhirStoreEnableUpdateCreate(d.Get("enable_update_create"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_update_create"); !isEmptyValue(reflect.ValueOf(enableUpdateCreateProp)) && (ok || !reflect.DeepEqual(v, enableUpdateCreateProp)) {
		obj["enableUpdateCreate"] = enableUpdateCreateProp
	}
	disableReferentialIntegrityProp, err := expandHealthcareFhirStoreDisableReferentialIntegrity(d.Get("disable_referential_integrity"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disable_referential_integrity"); !isEmptyValue(reflect.ValueOf(disableReferentialIntegrityProp)) && (ok || !reflect.DeepEqual(v, disableReferentialIntegrityProp)) {
		obj["disableReferentialIntegrity"] = disableReferentialIntegrityProp
	}
	disableResourceVersioningProp, err := expandHealthcareFhirStoreDisableResourceVersioning(d.Get("disable_resource_versioning"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disable_resource_versioning"); !isEmptyValue(reflect.ValueOf(disableResourceVersioningProp)) && (ok || !reflect.DeepEqual(v, disableResourceVersioningProp)) {
		obj["disableResourceVersioning"] = disableResourceVersioningProp
	}
	enableHistoryImportProp, err := expandHealthcareFhirStoreEnableHistoryImport(d.Get("enable_history_import"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_history_import"); !isEmptyValue(reflect.ValueOf(enableHistoryImportProp)) && (ok || !reflect.DeepEqual(v, enableHistoryImportProp)) {
		obj["enableHistoryImport"] = enableHistoryImportProp
	}
	labelsProp, err := expandHealthcareFhirStoreLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	notificationConfigProp, err := expandHealthcareFhirStoreNotificationConfig(d.Get("notification_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("notification_config"); !isEmptyValue(reflect.ValueOf(notificationConfigProp)) && (ok || !reflect.DeepEqual(v, notificationConfigProp)) {
		obj["notificationConfig"] = notificationConfigProp
	}

	url, err := replaceVars(d, config, "{{HealthcareBasePath}}{{dataset}}/fhirStores?fhirStoreId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new FhirStore: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating FhirStore: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{dataset}}/fhirStores/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating FhirStore %q: %#v", d.Id(), res)

	return resourceHealthcareFhirStoreRead(d, meta)
}

func resourceHealthcareFhirStoreRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{HealthcareBasePath}}{{dataset}}/fhirStores/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("HealthcareFhirStore %q", d.Id()))
	}

	res, err = resourceHealthcareFhirStoreDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if err := d.Set("name", flattenHealthcareFhirStoreName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading FhirStore: %s", err)
	}
	if err := d.Set("enable_update_create", flattenHealthcareFhirStoreEnableUpdateCreate(res["enableUpdateCreate"], d)); err != nil {
		return fmt.Errorf("Error reading FhirStore: %s", err)
	}
	if err := d.Set("disable_referential_integrity", flattenHealthcareFhirStoreDisableReferentialIntegrity(res["disableReferentialIntegrity"], d)); err != nil {
		return fmt.Errorf("Error reading FhirStore: %s", err)
	}
	if err := d.Set("disable_resource_versioning", flattenHealthcareFhirStoreDisableResourceVersioning(res["disableResourceVersioning"], d)); err != nil {
		return fmt.Errorf("Error reading FhirStore: %s", err)
	}
	if err := d.Set("enable_history_import", flattenHealthcareFhirStoreEnableHistoryImport(res["enableHistoryImport"], d)); err != nil {
		return fmt.Errorf("Error reading FhirStore: %s", err)
	}
	if err := d.Set("labels", flattenHealthcareFhirStoreLabels(res["labels"], d)); err != nil {
		return fmt.Errorf("Error reading FhirStore: %s", err)
	}
	if err := d.Set("notification_config", flattenHealthcareFhirStoreNotificationConfig(res["notificationConfig"], d)); err != nil {
		return fmt.Errorf("Error reading FhirStore: %s", err)
	}

	return nil
}

func resourceHealthcareFhirStoreUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	enableUpdateCreateProp, err := expandHealthcareFhirStoreEnableUpdateCreate(d.Get("enable_update_create"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_update_create"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enableUpdateCreateProp)) {
		obj["enableUpdateCreate"] = enableUpdateCreateProp
	}
	labelsProp, err := expandHealthcareFhirStoreLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	notificationConfigProp, err := expandHealthcareFhirStoreNotificationConfig(d.Get("notification_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("notification_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, notificationConfigProp)) {
		obj["notificationConfig"] = notificationConfigProp
	}

	url, err := replaceVars(d, config, "{{HealthcareBasePath}}{{dataset}}/fhirStores/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating FhirStore %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("enable_update_create") {
		updateMask = append(updateMask, "enableUpdateCreate")
	}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("notification_config") {
		updateMask = append(updateMask, "notificationConfig")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	_, err = sendRequestWithTimeout(config, "PATCH", url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating FhirStore %q: %s", d.Id(), err)
	}

	return resourceHealthcareFhirStoreRead(d, meta)
}

func resourceHealthcareFhirStoreDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{HealthcareBasePath}}{{dataset}}/fhirStores/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting FhirStore %q", d.Id())
	res, err := sendRequestWithTimeout(config, "DELETE", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "FhirStore")
	}

	log.Printf("[DEBUG] Finished deleting FhirStore %q: %#v", d.Id(), res)
	return nil
}

func resourceHealthcareFhirStoreImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*Config)

	fhirStoreId, err := parseHealthcareFhirStoreId(d.Id(), config)
	if err != nil {
		return nil, err
	}

	d.Set("dataset", fhirStoreId.DatasetId.datasetId())
	d.Set("name", fhirStoreId.Name)

	return []*schema.ResourceData{d}, nil
}

func flattenHealthcareFhirStoreName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenHealthcareFhirStoreEnableUpdateCreate(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenHealthcareFhirStoreDisableReferentialIntegrity(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenHealthcareFhirStoreDisableResourceVersioning(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenHealthcareFhirStoreEnableHistoryImport(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenHealthcareFhirStoreLabels(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenHealthcareFhirStoreNotificationConfig(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["pubsub_topic"] =
		flattenHealthcareFhirStoreNotificationConfigPubsubTopic(original["pubsubTopic"], d)
	return []interface{}{transformed}
}
func flattenHealthcareFhirStoreNotificationConfigPubsubTopic(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandHealthcareFhirStoreName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareFhirStoreEnableUpdateCreate(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareFhirStoreDisableReferentialIntegrity(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareFhirStoreDisableResourceVersioning(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareFhirStoreEnableHistoryImport(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareFhirStoreLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandHealthcareFhirStoreNotificationConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPubsubTopic, err := expandHealthcareFhirStoreNotificationConfigPubsubTopic(original["pubsub_topic"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPubsubTopic); val.IsValid() && !isEmptyValue(val) {
		transformed["pubsubTopic"] = transformedPubsubTopic
	}

	return transformed, nil
}

func expandHealthcareFhirStoreNotificationConfigPubsubTopic(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceHealthcareFhirStoreDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	// Take the returned long form of the name and use it as `self_link`.
	// Then modify the name to be the user specified form.
	// We can't just ignore_read on `name` as the linter will
	// complain that the returned `res` is never used afterwards.
	// Some field needs to be actually set, and we chose `name`.
	d.Set("self_link", res["name"].(string))
	res["name"] = d.Get("name").(string)
	return res, nil
}
