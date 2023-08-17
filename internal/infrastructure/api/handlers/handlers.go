package handlers

import (
	"github.com/LidenbrockGit/url-shortener/internal/usecases/repos/linkrepo"
	"github.com/LidenbrockGit/url-shortener/internal/usecases/repos/userrepo"
	"github.com/LidenbrockGit/url-shortener/internal/usecases/search/fullurlsearch"
)

type Handlers struct {
	Linkrepo      *linkrepo.LinkRepo
	Userrepo      *userrepo.UserRepo
	Fullurlsearch *fullurlsearch.FullUrl
}
