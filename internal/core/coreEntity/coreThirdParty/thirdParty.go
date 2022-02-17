package coreThirdParty

import (
	"finnomena/internal/core/coreEntity"
	"finnomena/internal/gateway/http"
	"finnomena/internal/gateway/httpStub"
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
