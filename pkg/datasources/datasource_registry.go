package datasources

import (
	"fmt"
)

type DatasourceRegistry struct {
	Sources []DatasourceInfo
}

var datasourceRegistry *DatasourceRegistry

func newDatasourceRegistry() *DatasourceRegistry {
	return &DatasourceRegistry{}
}

func GetDatasourceRegistry() *DatasourceRegistry {
	if datasourceRegistry == nil {
		datasourceRegistry = newDatasourceRegistry()
	}
	return datasourceRegistry
}

// Register the given datasource with the registry. If the datasource already exists
// an error is thrown
func (d *DatasourceRegistry) Register(dsInfo DatasourceInfo) error {
	for _, d := range d.Sources {
		if d.ID == dsInfo.ID {
			return fmt.Errorf("the data source %s is already registered", dsInfo.ID)
		}
	}
	d.Sources = append(d.Sources, dsInfo)
	return nil
}

// Get a datasource from the registry given its id.
// If the datasource does not exist, nil is returned
func (d *DatasourceRegistry) Get(id string) *DatasourceInfo {
	for _, d := range d.Sources {
		if d.ID == id {
			return &d
		}
	}
	return nil
}
