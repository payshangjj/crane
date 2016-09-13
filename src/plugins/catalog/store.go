package catalog

func (catalogApi *CatalogApi) Save(catalog *Catalog) error {
	return catalogApi.DbClient.Save(catalog).Error
}

func (catalogApi *CatalogApi) List() ([]Catalog, error) {
	var catalogs []Catalog
	err := catalogApi.DbClient.Find(&catalogs).Error
	return catalogs, err
}

func (catalogApi *CatalogApi) Get(catalogId uint64) (Catalog, error) {
	var catalog Catalog
	err := catalogApi.DbClient.Where("id = ?", catalogId).First(&catalog).Error
	return catalog, err
}

func (catalogApi *CatalogApi) Delete(catalogId uint64) error {
	return catalogApi.DbClient.Delete(&Catalog{ID: catalogId}).Error
}

func (catalogApi *CatalogApi) Update(catalog *Catalog) error {
	return catalogApi.DbClient.Save(catalog).Error
}
