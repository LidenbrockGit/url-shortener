package routergin

import "errors"

var CantReadRequestDataErr error = errors.New("can't read request data")
var MissingRequiredFieldsErr error = errors.New("missing required fields in request")
var UnhandledInternalError error = errors.New("unhandled internal error")
var IncorrectLinkIdErr = errors.New("incorrect link id")
