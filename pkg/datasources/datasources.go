package datasources

import (
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/compy/finder.sh/config"
	"github.com/compy/finder.sh/ent"
	"github.com/compy/finder.sh/pkg/services"
)

type ConfigFieldType string

const (
	Text     ConfigFieldType = "Text"
	TextArea ConfigFieldType = "TextArea"
	Checkbox ConfigFieldType = "Checkbox"
	Radio    ConfigFieldType = "Radio"
	Password ConfigFieldType = "Password"
)

type FieldOption struct {
	Text  string
	Value string
}

type ConfigField struct {
	Name        string
	PrettyName  string
	Description string
	Required    bool
	Type        ConfigFieldType
	Options     *[]FieldOption
	Checked     bool
	Default     string
	Value       string
	Placeholder string
}

type Datasource interface {
	GetConfigFields() []ConfigField
}

type DatasourceIndexer struct {
	Config *config.Config
	ORM    *ent.Client
	Tasks  *services.TaskClient
}

type DatasourceInfo struct {
	ID         string
	PrettyName string
	New        func() Datasource
}

type DatasourceIndexPayload struct {
	ID     int
	Type   string
	Config string
}

// initConfig initializes configuration
func (c *DatasourceIndexer) initConfig() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
	c.Config = &cfg
}

func (d *DatasourceIndexer) InitORM() {
	if d.Config == nil {
		d.initConfig()
	}
	getAddr := func(dbName string) string {
		return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
			d.Config.Database.User,
			d.Config.Database.Password,
			d.Config.Database.Hostname,
			d.Config.Database.Port,
			dbName,
		)
	}

	db, err := sql.Open("pgx", getAddr(d.Config.Database.Database))
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	d.ORM = ent.NewClient(ent.Driver(drv))
}

func (d *DatasourceIndexer) InitTasks() {
	if d.Config == nil {
		d.initConfig()
	}
	d.Tasks = services.NewTaskClient(d.Config)
}
