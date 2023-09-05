package routes

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/compy/finder.sh/ent"
	"github.com/compy/finder.sh/pkg/controller"
	ds "github.com/compy/finder.sh/pkg/datasources"
	"github.com/compy/finder.sh/pkg/msg"

	"github.com/labstack/echo/v4"
)

type (
	datasources struct {
		controller.Controller
	}
	configurePageData struct {
		DataSource   *ds.DatasourceInfo
		ConfigFields []ds.ConfigField
		FormValues   url.Values
	}

	configureDataSourceForm struct {
		Submission controller.FormSubmission
	}

	datasourcePageData struct {
		DataSources []*ent.DataSource
	}
)

func (c *datasources) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "datasources"

	data := datasourcePageData{}

	if dsList, err := c.Container.ORM.DataSource.Query().Order(ent.Asc("name")).All(ctx.Request().Context()); err == nil {
		data.DataSources = dsList
	} else {
		return c.Fail(err, "Could not list data sources")
	}

	page.Data = data

	return c.RenderPage(ctx, page)
}

func (c *datasources) Add(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "datasources-add"
	page.Data = ds.GetDatasourceRegistry().Sources

	return c.RenderPage(ctx, page)
}

func (c *datasources) Configure(ctx echo.Context) error {
	datasourceTypeId := ctx.Param("datasource")
	src := ds.GetDatasourceRegistry().Get(datasourceTypeId)

	if src == nil {
		// Unknown data source
		return ctx.Redirect(302, "/datasources/add")
	}
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "datasources-configure"

	ctx.Logger().Infof("Got datasource add request for %s", datasourceTypeId)
	dsObj := src.New()
	page.Data = &configurePageData{
		ConfigFields: dsObj.GetConfigFields(),
		DataSource:   src,
	}
	return c.RenderPage(ctx, page)
}

func (c *datasources) New(ctx echo.Context) error {
	id := ctx.Request().PostForm.Get("id")
	if id == "" {
		msg.Danger(ctx, "Invalid data source configuration. No id specified. Please try again")
		return ctx.Redirect(301, "/datasources")
	}

	src := ds.GetDatasourceRegistry().Get(id)
	if src == nil {
		msg.Danger(ctx, "Invalid data source configuration. Datasource type not found. Please try again")
		return ctx.Redirect(301, "/datasources")
	}

	dsObj := src.New()

	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "datasources-configure"

	pd := &configurePageData{
		ConfigFields: dsObj.GetConfigFields(),
		DataSource:   src,
		FormValues:   ctx.Request().PostForm,
	}

	page.Data = pd

	dsConfig := make(map[string]string, 0)

	for _, configField := range dsObj.GetConfigFields() {
		if configField.Required == true && ctx.Request().PostForm.Get(configField.Name) == "" {
			msg.Danger(ctx, configField.Name+" is required.")
			return c.RenderPage(ctx, page)
		} else {
			dsConfig[configField.Name] = ctx.Request().PostFormValue(configField.Name)
		}
	}

	jsonConfig, err := json.Marshal(dsConfig)
	if err != nil {
		msg.Danger(ctx, "Could not save datasource configuration")
		return c.RenderPage(ctx, page)
	}

	_, err = c.Container.ORM.DataSource.
		Create().
		SetName(ctx.Request().PostFormValue("name")).
		SetConfig(string(jsonConfig)).
		SetType(id).
		Save(ctx.Request().Context())

	if err != nil {
		msg.Danger(ctx, "Could not save datasource to database")
		return c.RenderPage(ctx, page)
	}

	//return ctx.JSON(200, ctx.Request().PostForm.Get("id"))
	msg.Success(ctx, "Data source created")
	return ctx.Redirect(301, "/datasources")
}

func (c *datasources) Reindex(ctx echo.Context) error {
	datasourceRowId := ctx.Param("datasource")
	id, err := strconv.Atoi(datasourceRowId)
	if err != nil {
		msg.Danger(ctx, "Datasource not found")
		return ctx.Redirect(301, "/datasources")
	}
	dsource, err := c.Container.ORM.DataSource.Get(ctx.Request().Context(), id)
	if err != nil {
		msg.Danger(ctx, fmt.Sprintf("Error fetching datasource: %s", err.Error()))
		return ctx.Redirect(301, "/datasources")
	}

	err = dsource.Update().
		SetStatus("indexing").
		Exec(ctx.Request().Context())

	if err != nil {
		msg.Danger(ctx, fmt.Sprintf("Error reindexing datasource: %s", err.Error()))
		return ctx.Redirect(301, "/datasources")
	}

	// Dispatch job
	pl := ds.DatasourceIndexPayload{
		ID:     dsource.ID,
		Type:   dsource.Type,
		Config: dsource.Config,
	}
	ctx.Logger().Infof("User is requesting a manual reindex of datasource %s (%s)", dsource.Name, dsource.Type)
	c.Container.Tasks.New(fmt.Sprintf("index_%s", dsource.Type)).Payload(pl).Save()

	msg.Success(ctx, "Datasource queued for reindexing. This may take a few moments.")
	return ctx.Redirect(301, "/datasources")
}
