package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/compy/finder.sh/ent"
	"github.com/compy/finder.sh/pkg/context"
	"github.com/compy/finder.sh/pkg/controller"
	"github.com/compy/finder.sh/pkg/msg"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type (
	groups struct {
		controller.Controller
	}

	newGroupForm struct {
		Name       string `form:"name" validate:"required"`
		Submission controller.FormSubmission
	}

	modifyGroupMemberForm struct {
		UserID     int `form:"user_id" validate:"required"`
		Submission controller.FormSubmission
	}

	groupPageData struct {
		Groups []*ent.Group
	}

	viewGroupPageData struct {
		Group *ent.Group
		Users []*ent.User
	}
)

func (c *groups) GetGroup(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "group"
	page.Form = modifyGroupMemberForm{}

	groupId := ctx.Param("group")
	log.Infof("Got group ID %s", groupId)

	intGroupId, err := strconv.Atoi(groupId)
	if err != nil {
		// Invalid input
		return c.Fail(err, "Could not get group. Invalid input")
	}

	group, err := c.Container.ORM.Group.Get(ctx.Request().Context(), intGroupId)
	if err != nil {
		// Invalid input
		return c.Fail(err, "Could not get group from database.")
	}

	pageData := viewGroupPageData{}
	pageData.Group = group

	if userList, err := c.Container.ORM.User.Query().Order(ent.Asc("name")).All(ctx.Request().Context()); err == nil {
		pageData.Users = userList
	} else {
		return c.Fail(err, "Could not list users")
	}

	page.Data = pageData

	return c.RenderPage(ctx, page)
}

func (c *groups) AddMember(ctx echo.Context) error {
	intGroupId, err := strconv.Atoi(ctx.Param("group"))
	if err != nil {
		return c.Fail(err, "Invalid group id")
	}

	var form modifyGroupMemberForm
	ctx.Set(context.FormKey, &form)

	// Parse the form values
	if err := ctx.Bind(&form); err != nil {
		return c.Fail(err, "unable to parse form")
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(err, "unable to process form submission")
	}

	if form.Submission.HasErrors() {
		log.Error("Form submission had errors")
		return c.GetGroup(ctx)
	}

	user, err := c.Container.ORM.User.Get(ctx.Request().Context(), form.UserID)
	if err != nil {
		log.Errorf("Could not get user record for user ID %d: %v", form.UserID, err)
		msg.Danger(ctx, fmt.Sprintf("Could not add user to group: %v", err))
		ctx.Response().Header().Set("HX-Redirect", fmt.Sprintf("/groups/%d", intGroupId))
		return ctx.String(http.StatusBadRequest, "")
	}

	err = user.Update().AddGroupIDs(intGroupId).Exec(ctx.Request().Context())
	if err != nil {
		log.Errorf("Error while adding user %d to group %d: %v", form.UserID, intGroupId, err)
		msg.Danger(ctx, fmt.Sprintf("Could not add user to group: %v", err))
		ctx.Response().Header().Set("HX-Redirect", fmt.Sprintf("/groups/%d", intGroupId))
		return ctx.String(http.StatusBadRequest, "")
	}

	msg.Success(ctx, fmt.Sprintf("Added <strong>%s</strong> to group successfully", user.Name))
	ctx.Response().Header().Set("HX-Redirect", fmt.Sprintf("/groups/%d", intGroupId))
	return ctx.String(http.StatusOK, "")
}

func (c *groups) RemoveMember(ctx echo.Context) error {
	intGroupId, err := strconv.Atoi(ctx.Param("group"))
	if err != nil {
		return c.Fail(err, "Invalid group id")
	}

	intUserId, err := strconv.Atoi(ctx.Param("userid"))
	if err != nil {
		return c.Fail(err, "Invalid user id")
	}

	user, err := c.Container.ORM.User.Get(ctx.Request().Context(), intUserId)
	if err != nil {
		log.Errorf("Could not get user record for user ID %d: %v", intUserId, err)
		msg.Danger(ctx, fmt.Sprintf("Could not remove user from group: %v", err))
		ctx.Response().Header().Set("HX-Redirect", fmt.Sprintf("/groups/%d", intGroupId))
		return ctx.String(http.StatusBadRequest, "")
	}

	err = user.Update().RemoveGroupIDs(intGroupId).Exec(ctx.Request().Context())
	if err != nil {
		log.Errorf("Error while removing user %d from group %d: %v", intUserId, intGroupId, err)
		msg.Danger(ctx, fmt.Sprintf("Could not remove user from group: %v", err))
		ctx.Response().Header().Set("HX-Redirect", fmt.Sprintf("/groups/%d", intGroupId))
		return ctx.String(http.StatusBadRequest, "")
	}

	msg.Success(ctx, fmt.Sprintf("Removed <strong>%s</strong> from group successfully", user.Name))
	ctx.Response().Header().Set("HX-Redirect", fmt.Sprintf("/groups/%d", intGroupId))
	return ctx.String(http.StatusOK, "")
}

func (c *groups) DeleteGroup(ctx echo.Context) error {
	intGroupId, err := strconv.Atoi(ctx.Param("group"))
	if err != nil || intGroupId == 1 {
		return c.Fail(err, "Invalid group id")
	}
	err = c.Container.ORM.Group.DeleteOneID(intGroupId).Exec(ctx.Request().Context())
	if err != nil {
		msg.Danger(ctx, "Could not delete group")
	} else {
		msg.Success(ctx, "Group deleted successfully")
	}
	ctx.Response().Header().Set("HX-Redirect", "/groups")
	return ctx.String(http.StatusOK, "")
}

func (c *groups) ListGroups(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "groups"
	page.Form = newGroupForm{}
	data := groupPageData{}

	if form := ctx.Get(context.FormKey); form != nil {
		log.Info("Got form")
		page.Form = form.(*newGroupForm)
	}

	if userList, err := c.Container.ORM.Group.Query().Order(ent.Asc("name")).All(ctx.Request().Context()); err == nil {
		data.Groups = userList
	} else {
		return c.Fail(err, "Could not list groups")
	}

	page.Data = data

	return c.RenderPage(ctx, page)
}

func (c *groups) NewGroup(ctx echo.Context) error {
	var form newGroupForm
	ctx.Set(context.FormKey, &form)

	// Parse the form values
	if err := ctx.Bind(&form); err != nil {
		return c.Fail(err, "unable to parse form")
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(err, "unable to process form submission")
	}

	if form.Submission.HasErrors() {
		log.Error("Form submission had errors")
		return c.ListGroups(ctx)
	}

	// Attempt creating the group
	u, err := c.Container.ORM.Group.
		Create().
		SetName(strings.ToLower(form.Name)).
		Save(ctx.Request().Context())

	switch err.(type) {
	case nil:
		ctx.Logger().Infof("group created: %s", u.Name)
	case *ent.ConstraintError:
		msg.Danger(ctx, "A group with that name already exists.")
		ctx.Response().Header().Set("HX-Redirect", "/groups")
		return ctx.String(http.StatusBadRequest, "")
	default:
		return c.Fail(err, "unable to create group")
	}

	//return c.ListUsers(ctx)
	msg.Success(ctx, "Group created successfully")
	ctx.Response().Header().Set("HX-Redirect", "/groups")
	return ctx.String(http.StatusOK, "")
}
