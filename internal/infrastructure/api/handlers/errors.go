package handlers

import "errors"

var WrongPasswordErr = errors.New("wrong password")
var IncorrectUserIdErr = errors.New("incorrect user id")
var IncorrectLinkIdErr = errors.New("incorrect link id")
