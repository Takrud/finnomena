package coreEntity

type CoreEntity struct {
	ThirdParty
}

var coreEntity CoreEntity

func InitCoreEntity(en CoreEntity) {
	coreEntity = en
}

func Entity() CoreEntity {
	return coreEntity
}
