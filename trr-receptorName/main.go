// This file is subject to the terms and conditions defined in
// file 'LICENSE.txt', which is part of this source code package.
package main

import (
	"github.com/trustero/api/go/receptor_sdk"
	"github.com/trustero/api/go/receptor_sdk/cmd"
	"github.com/trustero/api/go/receptor_v1"
)

const (
	receptorName = "CHANGE ME"
	serviceName1 = "CHANGE ME"
	serviceName2 = "CHANGE ME"
)

// This struct holds the credentials the receptor needs to authenticate with the
// service provider. A display name and placeholder tag should be provided
// which will be used in the UI when activating the receptor.
// This is what will be returned in the GetCredentialObj call
type Receptor struct {
	//credentials
}

// Set the name of the receptor in the const declaration above
// This will let the receptor inform Trustero about itself
func (r *Receptor) GetReceptorType() string {
	return "trr-custom"
}

// Set the names of the services in the const declaration above
// This will let the receptor inform Trustero about itself
// Feel free to add or remove services as needed
func (r *Receptor) GetKnownServices() []string {
	return []string{serviceName1, serviceName2}
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

	return
}

// The Discover function returns a list of Service Entities. This function 
// makes any relevant API calls to the Service Provider to gather information
// about how many Service Entity Instances are in use. If at any point this 
// function runs into an error, log that error and continue
func (r *Receptor) Discover(credentials interface{}) (svcs []*receptor_v1.ServiceEntity, err error) {

	return
}

// Report will often make the same API calls made in the Discover call, but it
// will additionally create evidences with the data returned from the API calls
func (r *Receptor) Report(credentials interface{}) (evidences []*receptor_sdk.Evidence, err error) {

	return
}

func main() {
	cmd.Execute(&Receptor{})
}
