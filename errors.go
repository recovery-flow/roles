package roles

import "github.com/pkg/errors"

var ErrorRole = errors.New("Invalid role type")
var ErrorNoPermission = errors.New("User is have not permission")
var ErrorRolePriority = errors.New("You can't change/delete user with the same or higher role")
