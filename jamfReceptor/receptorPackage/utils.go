/* TODO: Name package */
package receptorPackage

import (
	"net/http"
	"sync"

	"github.com/rs/zerolog/log"
	computers "github.com/trustero/jamf-api-client-go/classic/computers"
)

/* TODO: Write any helper functions here */
type Discovered struct {
	Computers []*computers.Computer `json:"computers"`
}

type computerResponse struct {
	computer *computers.Computer
	apiCalls []string
}

const serviceId = "trs86"

func findComputers(computerService *computers.Service) (computerInfo []*computers.Computer, apiCalls []string, err error) {
	computerList, resp, err := computerService.List()
	if err != nil {
		return
	}
	apiCalls = append(apiCalls, resp.Request.URL.String())

	out := make(chan *computerResponse, 100)

	go getComputerInfo(computerService, computerList, out)
	for computer := range out {
		computerInfo = append(computerInfo, computer.computer)
		apiCalls = append(apiCalls, computer.apiCalls...)
	}
	return
}

func getComputerInfo(computerService *computers.Service, computerList []computers.ComputerNameId, out chan *computerResponse) {
	defer close(out)
	var pg sync.WaitGroup

	for _, p := range computerList {
		pg.Add(1)
		go func(id int) {
			defer pg.Done()
			result := &computerResponse{}
			var resp *http.Response
			var err error
			if result.computer, resp, err = computerService.GetById(id); err != nil {
				log.Error().Msgf(err.Error())
				return
			}
			result.apiCalls = append(result.apiCalls, resp.Request.URL.String())
			out <- result
		}(p.Id)
	}
	pg.Wait()
}
