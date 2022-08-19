// This file is subject to the terms and conditions defined in
// file 'LICENSE.txt', which is part of this source code package.
package main

import (
	"github.com/trustero/api/go/receptor_sdk"
	"github.com/trustero/api/go/receptor_sdk/cmd"
	"github.com/trustero/api/go/receptor_v1"
)

const (
	serviceName = "CHANGE ME"
)

// Credential object
type Receptor struct {
	//credentials
}

func (r *Receptor) GetReceptorType() string {
	return "trr-CHANGE-ME"
}

func (r *Receptor) GetKnownServices() []string {
	return []string{serviceName}
}

func (r *Receptor) GetCredentialObj() (credentialObj interface{}) {
	return r
}

func (r *Receptor) Verify(credentials interface{}) (ok bool, err error) {

	return
}

func (r *Receptor) Discover(credentials interface{}) (svcs []*receptor_v1.ServiceEntity, err error) {

	return
}

func (r *Receptor) Report(credentials interface{}) (evidences []*receptor_sdk.Evidence, err error) {

	return
}

func main() {
	cmd.Execute(&Receptor{})
}
