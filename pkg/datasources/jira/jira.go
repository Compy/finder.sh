package jira

import (
	"log"

	"github.com/compy/finder.sh/pkg/datasources"
)

type JiraDataSource struct {
	datasources.DatasourceIndexer
}

func init() {
	log.Printf("Initializing JIRA datasource plugin...")
	datasources.GetDatasourceRegistry().Register(datasources.DatasourceInfo{
		ID:         "jira",
		PrettyName: "Jira",
		New:        func() datasources.Datasource { return NewJiraDatasource() },
	})
}

func NewJiraDatasource() *JiraDataSource {
	newDatasource := &JiraDataSource{}
	// Since we're going to use the database, initialize the ORM
	newDatasource.InitORM()
	// We're going to be feeding documents into the main indexer pipeline, so initialize some task services
	// we can use
	newDatasource.InitTasks()
	log.Printf("New JIRA data source initialized...")
	return newDatasource
}

func (d *JiraDataSource) GetConfigFields() []datasources.ConfigField {
	return []datasources.ConfigField{
		{
			Name:        "url",
			PrettyName:  "Jira URL",
			Description: "Your Jira cloud URL",
			Type:        datasources.Text,
			Placeholder: "https://myorg.atlassian.net",
			Required:    true,
		},
		{
			Name:        "token",
			PrettyName:  "API Access Token",
			Description: "Your Jira API access token",
			Type:        datasources.Password,
		},
		{
			Name:        "username",
			PrettyName:  "Jira Username (Email)",
			Description: "Your Jira username or email address",
			Type:        datasources.Text,
			Placeholder: "me@company.com",
		},
	}
}
