// This file is subject to the terms and conditions defined in
// file 'LICENSE.txt', which is part of this source code package.
package main

/* TODO: Change import name */
import (
	"github.com/trustero/api/go/receptor_sdk"
	"github.com/trustero/api/go/receptor_sdk/cmd"
	"github.com/trustero/api/go/receptor_v1"
	"receptor/trr-receptorName/receptorPackageName"
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
	/* TODO: Change to package name */
	return receptorPackageName.GetReceptorTypeImpl()
}

// Set the names of the services in the const declaration above
// This will let the receptor inform Trustero about itself
// Feel free to add or remove services as needed
func (r *Receptor) GetKnownServices() []string {
	/* TODO: Change to package name */
	return receptorPackageName.GetKnownServicesImpl()
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
func (r *Receptor) Verify(credentials interface{}) (ok bool, err error) {
	c := credentials.(*Receptor)
	/* TODO: Change to package name */
	/* TODO: Change crediential field names */
	return receptorPackageName.VerifyImpl(c.field1, c.field2)
}

// The Discover function returns a list of Service Entities. This function
// makes any relevant API calls to the Service Provider to gather information
// about how many Service Entity Instances are in use. If at any point this
// function runs into an error, log that error and continue
func (r *Receptor) Discover(credentials interface{}) (svcs []*receptor_v1.ServiceEntity, err error) {
	c := credentials.(*Receptor)
	/* TODO: Change to package name */
	/* TODO: Change crediential field names */
	return receptorPackageName.DiscoverImpl(c.field1, c.field2)
}

// Report will often make the same API calls made in the Discover call, but it
// will additionally create evidences with the data returned from the API calls
func (r *Receptor) Report(credentials interface{}) (evidences []*receptor_sdk.Evidence, err error) {
	c := credentials.(*Receptor)
	/* TODO: Change to package name */
	/* TODO: Change crediential field names */
	return receptorPackageName.ReportImpl(c.field1, c.field2)
}

func main() {
	cmd.Execute(&Receptor{})
}
