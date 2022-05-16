package queries

import "github.com/checkmarxDev/ast-sast-export/internal/integration/soap"

type Repo struct {
	soapClient soap.Adapter
}

func NewRepo(soapClient soap.Adapter) *Repo {
	return &Repo{soapClient: soapClient}
}

func (e *Repo) GetQueriesList() (*soap.GetQueryCollectionResponse, error) {
	return e.soapClient.GetQueryCollection()
}
