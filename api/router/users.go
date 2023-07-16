package router

import (
	"github.com/bwoff11/frens/service"
)

type UsersRepo struct {
	Service *service.UserService
}
