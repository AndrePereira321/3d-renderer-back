package cache

type CacheStruct struct {
	References ReferenceCache
}

var Cache = CacheStruct{
	References: ReferenceCache{
		RefsByName:          map[string]*ReferenceItemDTO{},
		RefsCodeByTableName: map[string]map[string]*ReferenceValue{},
	},
}

func (cache *CacheStruct) Init() error {
	err := cache.References.Init()
	if err != nil {
		return err
	}
	return nil
}

func (cache *CacheStruct) GetReference(tableName string) *ReferenceItemDTO {
	return cache.References.GetReference(tableName)
}
