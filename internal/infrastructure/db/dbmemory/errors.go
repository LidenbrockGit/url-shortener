package dbmemory

import "errors"

var UserFindErr = errors.New("cannot find user")
var UserLoginDuplicatingErr = errors.New("login is already used")

var LinkFindErr = errors.New("cannot find link")
var LinkShortUrlDuplicatingErr = errors.New("short url is already used")
var LinkUpdateErr = errors.New("cannot update link")
