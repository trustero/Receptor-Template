/* TODO: Name package */
package receptorPackage

import (
	receptorLog "github.com/trustero/api/go/examples/gitlab_receptor/logging"

	"github.com/trustero/api/go/receptor_sdk"
	"github.com/trustero/api/go/receptor_v1"
	"github.com/trustero/jamf-api-client-go/classic/computers"
)

const (
	receptorName = "trr-jamf"
	serviceName  = "Jamf Service"
)

func GetReceptorTypeImpl() string {
	return receptorName
}

func GetKnownServicesImpl() []string {
	return []string{serviceName}
}

func VerifyImpl(userName string, password string, baseUrl string) (ok bool, err error) {
	receptorLog.Info("Entering VerifyImpl")
	ok = true

	var j *computers.Service
	j, err = computers.NewService(userName, password, baseUrl, nil)
	if err != nil {
		receptorLog.Err(err, "Could not verify, error in Jamf Computers NewService for %s ", userName)
		return false, nil
	}

	// Example: Get All Computers
	if _, _, err = j.List(); err != nil {
		return false, err
	}

	receptorLog.Info("Leaving VerifyImpl")
	return
}

func DiscoverImpl(userName string, password string, baseUrl string) (svcs []*receptor_v1.ServiceEntity, err error) {
	receptorLog.Info("Entering DiscoverImpl")
	//services := receptor_sdk.NewServiceEntities()

	receptorLog.Info("Leaving DiscoverImpl")
	return
}

func ReportImpl(userName string, password string, baseUrl string) (evidences []*receptor_sdk.Evidence, err error) {
	receptorLog.Info("Entering ReportImpl")
	/* TODO: Implement Report logic here */
	receptorLog.Info("Leaving ReportImpl")
	return
}
