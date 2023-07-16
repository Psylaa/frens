package router

import (
	"github.com/bwoff11/frens/service/user"
)

type UsersRepo struct {
	Service *user.Service
}
