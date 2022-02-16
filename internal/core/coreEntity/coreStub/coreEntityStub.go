package coreStub

import (
	"finno/internal/core/coreEntity"
	"finno/internal/core/coreEntity/coreThirdParty"
)

func NewStubEntity() coreEntity.CoreEntity {
	entities := coreEntity.CoreEntity{}
	coreThirdParty.NewStubThirdPartyEntity(&entities)
	return entities
}