package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ctx *appContext) AdminPage(gc *gin.Context) {
	bs5 := ctx.config.Section("ui").Key("bs5").MustBool(false)
	emailEnabled, _ := ctx.config.Section("invite_emails").Key("enabled").Bool()
	notificationsEnabled, _ := ctx.config.Section("notifications").Key("enabled").Bool()
	gc.HTML(http.StatusOK, "admin.html", gin.H{
		"bs5":            bs5,
		"cssFile":        ctx.cssFile,
		"contactMessage": "",
		"email_enabled":  emailEnabled,
		"notifications":  notificationsEnabled,
	})
}

func (ctx *appContext) InviteProxy(gc *gin.Context) {
	code := gc.Param("invCode")
	/* Don't actually check if the invite is valid, just if it exists, just so the page loads quicker. Invite is actually checked on submit anyway. */
	// if ctx.checkInvite(code, false, "") {
	if _, ok := ctx.storage.invites[code]; ok {
		email := ctx.storage.invites[code].Email
		gc.HTML(http.StatusOK, "form.html", gin.H{
			"bs5":            ctx.config.Section("ui").Key("bs5").MustBool(false),
			"cssFile":        ctx.cssFile,
			"contactMessage": ctx.config.Section("ui").Key("contac_message").String(),
			"helpMessage":    ctx.config.Section("ui").Key("help_message").String(),
			"successMessage": ctx.config.Section("ui").Key("success_message").String(),
			"jfLink":         ctx.config.Section("jellyfin").Key("public_server").String(),
			"validate":       ctx.config.Section("password_validation").Key("enabled").MustBool(false),
			"requirements":   ctx.validator.getCriteria(),
			"email":          email,
			"username":       !ctx.config.Section("email").Key("no_username").MustBool(false),
		})
	} else {
		gc.HTML(404, "invalidCode.html", gin.H{
			"bs5":            ctx.config.Section("ui").Key("bs5").MustBool(false),
			"cssFile":        ctx.cssFile,
			"contactMessage": ctx.config.Section("ui").Key("contac_message").String(),
		})
	}
}

func (ctx *appContext) NoRouteHandler(gc *gin.Context) {
	gc.HTML(404, "404.html", gin.H{
		"bs5":            ctx.config.Section("ui").Key("bs5").MustBool(false),
		"cssFile":        ctx.cssFile,
		"contactMessage": ctx.config.Section("ui").Key("contact_message").String(),
	})
}