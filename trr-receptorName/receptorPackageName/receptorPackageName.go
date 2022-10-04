/* TODO: Name package */
package receptorPackageName

import (
	"github.com/trustero/api/go/receptor_sdk"
	"github.com/trustero/api/go/receptor_v1"
	"receptor/trr-receptorName/logging"
)

const (
	receptorName = "trr-bitbucket"
	serviceName  = "Custom Service"
)

func GetReceptorTypeImpl() string {
	return receptorName
}

func GetKnownServicesImpl() []string {
	return []string{serviceName}
}

func VerifyImpl( /* TODO: Put needed receptor creds here */ field1 int, field2 string) (ok bool, err error) {
	receptorLog.Debug("Entering ReportImpl")
	/* TODO: Implement Verify logic here */
	receptorLog.Debug("Leaving ReportImpl")
	return
}

func DiscoverImpl( /* TODO: Put needed receptor creds here */ field1 int, field2 string) (svcs []*receptor_v1.ServiceEntity, err error) {
	/* TODO: Implement Discover logic here */
	receptorLog.Debug("Leaving ReportImpl")
	return
}

func ReportImpl( /* TODO: Put needed receptor creds here */ field1 int, field2 string) (evidences []*receptor_sdk.Evidence, err error) {
	receptorLog.Debug("Entering ReportImpl")
	/* TODO: Implement Report logic here */
	receptorLog.Debug("Leaving ReportImpl")
	return
}
