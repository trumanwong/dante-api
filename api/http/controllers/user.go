package controllers

import (
	"bufio"
	"dante-api/configs"
	"github.com/gin-gonic/gin"
	"github.com/trumanwong/go-internal/util"
	"net/http"
	"os/exec"
	"strings"
)

type UserController struct {
	Controller
}

type User struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Store
// Create/Update User
func (this *UserController) Store(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		util.Response(ctx, nil, http.StatusBadRequest, err.Error())
		return
	}

	err := exec.Command(
		configs.Config.Command.Sockd,
		"add",
		user.Name,
		user.Password,
	).Run()
	if err != nil {
		util.Response(ctx, nil, http.StatusInternalServerError, err.Error())
		return
	}
	util.Response(ctx, nil, http.StatusOK, "Success")
}

// Show
// query Users
func (this *UserController) Show(ctx *gin.Context) {
	cmd := exec.Command(
		configs.Config.Command.Sockd,
		"show",
	)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		util.Response(ctx, nil, http.StatusInternalServerError, err.Error())
		return
	}
	if err := cmd.Start(); err != nil {
		util.Response(ctx, nil, http.StatusInternalServerError, err.Error())
		return
	}

	buf := bufio.NewReader(stdout)
	i := 0
	res := make([]string, 0)
	for {
		output, _, err := buf.ReadLine()
		if err != nil {
			if err.Error() != "EOF" {
				util.Response(ctx, nil, http.StatusInternalServerError, err.Error())
				return
			}
			break
		}
		if i > 0 {
			res = append(res, string(output))
		}
		i++
	}

	if err := cmd.Wait(); err != nil {
		util.Response(ctx, nil, http.StatusInternalServerError, err.Error())
		return
	}
	util.Response(ctx, res, http.StatusOK, "Success")
}

// Delete
// delete user
func (this *UserController) Delete(ctx *gin.Context) {
	name := strings.Trim(ctx.PostForm("name"), " ")
	if len(name) == 0 {
		util.Response(ctx, nil, http.StatusBadRequest, "Please input username")
		return
	}
	err := exec.Command(
		configs.Config.Command.Sockd,
		"del",
		name,
	).Run()
	if err != nil {
		util.Response(ctx, nil, http.StatusInternalServerError, err.Error())
		return
	}
	util.Response(ctx, nil, http.StatusOK, "Success")
}
