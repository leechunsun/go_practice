package controllers

import (
	"chunsheng.isir.com/repositories"
	"chunsheng.isir.com/services"
	"fmt"
	"github.com/kataras/iris/v12/context"
)

type MovieController struct {
	context.Context
}

func (m *MovieController) Get(){
	reps := repositories.NewMovieRepositories()
	m.Context.HTML(fmt.Sprintf("<h1>%s</h1>", services.NewMovieService(reps).GetMovieName(m.Context.Params().Get("mid"))))
	m.Context.Next()
}
