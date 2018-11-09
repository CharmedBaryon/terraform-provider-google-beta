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
	"regexp"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceBinaryAuthorizationAttestor() *schema.Resource {
	return &schema.Resource{
		Create: resourceBinaryAuthorizationAttestorCreate,
		Read:   resourceBinaryAuthorizationAttestorRead,
		Update: resourceBinaryAuthorizationAttestorUpdate,
		Delete: resourceBinaryAuthorizationAttestorDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBinaryAuthorizationAttestorImport,
		},

		Schema: map[string]*schema.Schema{
			"attestation_authority_note": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"note_reference": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: compareSelfLinkOrResourceName,
						},
						"public_keys": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ascii_armored_pgp_public_key": {
										Type:     schema.TypeString,
										Required: true,
									},
									"comment": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"delegation_service_account_email": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceBinaryAuthorizationAttestorCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandBinaryAuthorizationAttestorName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandBinaryAuthorizationAttestorDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	userOwnedDrydockNoteProp, err := expandBinaryAuthorizationAttestorAttestationAuthorityNote(d.Get("attestation_authority_note"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("attestation_authority_note"); !isEmptyValue(reflect.ValueOf(userOwnedDrydockNoteProp)) && (ok || !reflect.DeepEqual(v, userOwnedDrydockNoteProp)) {
		obj["userOwnedDrydockNote"] = userOwnedDrydockNoteProp
	}

	url, err := replaceVars(d, config, "https://binaryauthorization.googleapis.com/v1beta1/projects/{{project}}/attestors?attestorId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Attestor: %#v", obj)
	res, err := sendRequest(config, "POST", url, obj)
	if err != nil {
		return fmt.Errorf("Error creating Attestor: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Attestor %q: %#v", d.Id(), res)

	return resourceBinaryAuthorizationAttestorRead(d, meta)
}

func resourceBinaryAuthorizationAttestorRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://binaryauthorization.googleapis.com/v1beta1/projects/{{project}}/attestors/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("BinaryAuthorizationAttestor %q", d.Id()))
	}

	if err := d.Set("name", flattenBinaryAuthorizationAttestorName(res["name"])); err != nil {
		return fmt.Errorf("Error reading Attestor: %s", err)
	}
	if err := d.Set("description", flattenBinaryAuthorizationAttestorDescription(res["description"])); err != nil {
		return fmt.Errorf("Error reading Attestor: %s", err)
	}
	if err := d.Set("attestation_authority_note", flattenBinaryAuthorizationAttestorAttestationAuthorityNote(res["userOwnedDrydockNote"])); err != nil {
		return fmt.Errorf("Error reading Attestor: %s", err)
	}
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Attestor: %s", err)
	}

	return nil
}

func resourceBinaryAuthorizationAttestorUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandBinaryAuthorizationAttestorName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandBinaryAuthorizationAttestorDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	userOwnedDrydockNoteProp, err := expandBinaryAuthorizationAttestorAttestationAuthorityNote(d.Get("attestation_authority_note"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("attestation_authority_note"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, userOwnedDrydockNoteProp)) {
		obj["userOwnedDrydockNote"] = userOwnedDrydockNoteProp
	}

	url, err := replaceVars(d, config, "https://binaryauthorization.googleapis.com/v1beta1/projects/{{project}}/attestors/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Attestor %q: %#v", d.Id(), obj)
	_, err = sendRequest(config, "PUT", url, obj)

	if err != nil {
		return fmt.Errorf("Error updating Attestor %q: %s", d.Id(), err)
	}

	return resourceBinaryAuthorizationAttestorRead(d, meta)
}

func resourceBinaryAuthorizationAttestorDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://binaryauthorization.googleapis.com/v1beta1/projects/{{project}}/attestors/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Attestor %q", d.Id())
	res, err := sendRequest(config, "DELETE", url, obj)
	if err != nil {
		return handleNotFoundError(err, d, "Attestor")
	}

	log.Printf("[DEBUG] Finished deleting Attestor %q: %#v", d.Id(), res)
	return nil
}

func resourceBinaryAuthorizationAttestorImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	parseImportId([]string{"projects/(?P<project>[^/]+)/attestors/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config)

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenBinaryAuthorizationAttestorName(v interface{}) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenBinaryAuthorizationAttestorDescription(v interface{}) interface{} {
	return v
}

func flattenBinaryAuthorizationAttestorAttestationAuthorityNote(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})
	transformed["note_reference"] =
		flattenBinaryAuthorizationAttestorAttestationAuthorityNoteNoteReference(original["noteReference"])
	transformed["public_keys"] =
		flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeys(original["publicKeys"])
	transformed["delegation_service_account_email"] =
		flattenBinaryAuthorizationAttestorAttestationAuthorityNoteDelegationServiceAccountEmail(original["delegationServiceAccountEmail"])
	return []interface{}{transformed}
}
func flattenBinaryAuthorizationAttestorAttestationAuthorityNoteNoteReference(v interface{}) interface{} {
	return v
}

func flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeys(v interface{}) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		transformed = append(transformed, map[string]interface{}{
			"comment":                      flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysComment(original["comment"]),
			"id":                           flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysId(original["id"]),
			"ascii_armored_pgp_public_key": flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysAsciiArmoredPgpPublicKey(original["asciiArmoredPgpPublicKey"]),
		})
	}
	return transformed
}
func flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysComment(v interface{}) interface{} {
	return v
}

func flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysId(v interface{}) interface{} {
	return v
}

func flattenBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysAsciiArmoredPgpPublicKey(v interface{}) interface{} {
	return v
}

func flattenBinaryAuthorizationAttestorAttestationAuthorityNoteDelegationServiceAccountEmail(v interface{}) interface{} {
	return v
}

func expandBinaryAuthorizationAttestorName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationAttestorDescription(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNote(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNoteReference, err := expandBinaryAuthorizationAttestorAttestationAuthorityNoteNoteReference(original["note_reference"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNoteReference); val.IsValid() && !isEmptyValue(val) {
		transformed["noteReference"] = transformedNoteReference
	}

	transformedPublicKeys, err := expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeys(original["public_keys"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPublicKeys); val.IsValid() && !isEmptyValue(val) {
		transformed["publicKeys"] = transformedPublicKeys
	}

	transformedDelegationServiceAccountEmail, err := expandBinaryAuthorizationAttestorAttestationAuthorityNoteDelegationServiceAccountEmail(original["delegation_service_account_email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDelegationServiceAccountEmail); val.IsValid() && !isEmptyValue(val) {
		transformed["delegationServiceAccountEmail"] = transformedDelegationServiceAccountEmail
	}

	return transformed, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNoteNoteReference(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	r := regexp.MustCompile("projects/(.+)/notes/(.+)")
	if r.MatchString(v.(string)) {
		return v.(string), nil
	}

	project, err := getProject(d, config)
	if err != nil {
		return nil, err
	}

	return fmt.Sprintf("projects/%s/notes/%s", project, v.(string)), nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeys(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedComment, err := expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysComment(original["comment"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedComment); val.IsValid() && !isEmptyValue(val) {
			transformed["comment"] = transformedComment
		}

		transformedId, err := expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysId(original["id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedId); val.IsValid() && !isEmptyValue(val) {
			transformed["id"] = transformedId
		}

		transformedAsciiArmoredPgpPublicKey, err := expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysAsciiArmoredPgpPublicKey(original["ascii_armored_pgp_public_key"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAsciiArmoredPgpPublicKey); val.IsValid() && !isEmptyValue(val) {
			transformed["asciiArmoredPgpPublicKey"] = transformedAsciiArmoredPgpPublicKey
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysComment(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysId(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysAsciiArmoredPgpPublicKey(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBinaryAuthorizationAttestorAttestationAuthorityNoteDelegationServiceAccountEmail(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}
