package coreStub

import (
	"finnomena/internal/core/coreEntity"
	"finnomena/internal/core/coreEntity/coreThirdParty"
)

func NewStubEntity() coreEntity.CoreEntity {
	entities := coreEntity.CoreEntity{}
	coreThirdParty.NewStubThirdPartyEntity(&entities)
	return entities
}
