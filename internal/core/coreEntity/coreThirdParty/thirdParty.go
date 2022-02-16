package coreThirdParty

import (
	"finno/internal/core/coreEntity"
	"finno/internal/gateway/http"
	"finno/internal/gateway/httpStub"
)

func NewEntity() coreEntity.CoreEntity {
	entities := coreEntity.CoreEntity{}
	NewThirdPartyEntity(&entities)
	return entities
}

func NewStubThirdPartyEntity(entities *coreEntity.CoreEntity) {
	entities.ThirdParty.HTTPFundService = httpStub.NewHTTPGateway()
}

func NewThirdPartyEntity(entities *coreEntity.CoreEntity) {
	entities.ThirdParty.HTTPFundService = http.NewHTTPGateway()
}