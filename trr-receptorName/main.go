// This file is subject to the terms and conditions defined in
// file 'LICENSE.txt', which is part of this source code package.
package main

/* TODO: Change import name */
import (
	receptorPackage "receptor/trr-receptorName/receptorPackage"

	"github.com/trustero/api/go/receptor_sdk"
	"github.com/trustero/api/go/receptor_sdk/cmd"
	"github.com/trustero/api/go/receptor_v1"
)

// This struct holds the credentials the receptor needs to authenticate with the
// service provider. A display name and placeholder tag should be provided
// which will be used in the UI when activating the receptor.
// This is what will be returned in the GetCredentialObj call
type Receptor struct {
	//credentials
	/* TODO: Change this to the needed credential fields */
	field1 int
	field2 string
}

// Set the name of the receptor in the const declaration above
// This will let the receptor inform Trustero about itself
func (r *Receptor) GetReceptorType() string {
	return receptorPackage.GetReceptorTypeImpl()
}

// Set the names of the services in the const declaration above
// This will let the receptor inform Trustero about itself
// Feel free to add or remove services as needed
func (r *Receptor) GetKnownServices() []string {
	return receptorPackage.GetKnownServicesImpl()
}

// This will return Receptor struct defined above when the receptor is asked to
// identify itself
func (r *Receptor) GetCredentialObj() (credentialObj interface{}) {
	return r
}

// This function will call into the service provider API with the provided
// credentials and confirm that the credentials are valid. Usually a simple
// API call like GET org name. If the credentials are not valid,
// return a relevant error message
func (r *Receptor) Verify(credentials interface{}, config interface{}) (ok bool, err error) {
	c := credentials.(*Receptor)
	/* TODO: Change crediential field names */
	return receptorPackage.VerifyImpl(c.field1, c.field2)
}

// The Discover function returns a list of Service Entities. This function
// makes any relevant API calls to the Service Provider to gather information
// about how many Service Entity Instances are in use. If at any point this
// function runs into an error, log that error and continue
func (r *Receptor) Discover(credentials interface{}, config interface{}) (svcs []*receptor_v1.ServiceEntity, err error) {
	c := credentials.(*Receptor)
	/* TODO: Change crediential field names */
	return receptorPackage.DiscoverImpl(c.field1, c.field2)
}

// Report will often make the same API calls made in the Discover call, but it
// will additionally create evidences with the data returned from the API calls
func (r *Receptor) Report(credentials interface{}, config interface{}) (evidences []*receptor_sdk.Evidence, err error) {
	c := credentials.(*Receptor)
	/* TODO: Change crediential field names */
	return receptorPackage.ReportImpl(c.field1, c.field2)
}

// Configure returns a ReceptorConfiguration object that represents the configuration
// of the receptor.
func (r *Receptor) Configure(credentials interface{}) (config *receptor_v1.ReceptorConfiguration, err error) {
	return nil, nil
}

// GetAuthMethods returns the authentication methods supported by the receptor
// This is usually implemented for services that support multiple authentication
// so as to support previously configured authentication methods that need not be
// the preferred method anymore. If only one authentication method is available the
// function should return nil
func (r *Receptor) GetAuthMethods() interface{} {
	return nil
}

// GetInstructions returns the instructions for the receptor in markdown format
// The instructions should be in the 'resources' directory so that they can be
// embedded in the receptor executable. Helper functions has been provided in
// the receptorPackage to embed the instructions from file content.
func (r *Receptor) GetInstructions() (string, error) {
	return receptorPackage.GetInstructionsImpl()
}

// GetLogo returns the logo for the receptor in svg format. The logo should be in
// the 'resources' directory so that it can be embedded in the receptor executable.
// Helper function has been provided in the receptorPackage to embed the logo from
// the file content. The logo file should be in svg format.
func (r *Receptor) GetLogo() (string, error) {
	return receptorPackage.GetLogoImpl()
}

// / GetConfigObj will return the configuration required by the receptor
// to be able to collect data from the service provider API. This is usually
// implemented for services that require additional configuration like Jira and Slack
// which require the user to provide the filters and channel name respectively for
// the data to be collected. If no additional configuration is required, return nil.
func (r *Receptor) GetConfigObj() interface{} {
	return nil
}

// GetConfigObjDesc will return an object that describes the configuration required
// by the receptor to be able to collect data from the service provider API.
// The sdk supports a Config struct that can be used to describe the configuration
// which can be used to generate a form in the UI for the user to map where the evidence
// collected will flow. If no additional configuration is required, return nil.

func (r *Receptor) GetConfigObjDesc() interface{} {
	return nil
}

// GetEvidenceInfo returns all the evidence information that the receptor will collect.
func (r *Receptor) GetEvidenceInfo() []*receptor_sdk.Evidence {
	evidences := []*receptor_sdk.Evidence{}
	return evidences
}

func main() {
	cmd.Execute(&Receptor{})
}
