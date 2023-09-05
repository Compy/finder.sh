package routes

import (
	"fmt"
	"math/rand"

	"github.com/compy/finder.sh/pkg/controller"

	"github.com/labstack/echo/v4"
)

type (
	search struct {
		controller.Controller
	}

	searchResult struct {
		Title string
		URL   string
	}

	searchData struct {
		Query   string
		Results []searchResult
	}
)

func (c *search) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "search"
	page.Name = "search"
	data := searchData{}

	// Fake search results
	var results []searchResult
	if search := ctx.QueryParam("q"); search != "" {
		data.Query = search
		for i := 0; i < 20; i++ {
			title := "Lorem ipsum example ddolor sit amet"
			index := rand.Intn(len(title))
			title = title[:index] + search + title[index:]
			results = append(results, searchResult{
				Title: title,
				URL:   fmt.Sprintf("https://www.%s.com", search),
			})
		}
	}
	data.Results = results
	page.Data = data

	return c.RenderPage(ctx, page)
}
