package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/LidenbrockGit/url-shortener/internal/infrastructure/api/handlers"
	"github.com/LidenbrockGit/url-shortener/internal/infrastructure/api/routers/routergin"
	_ "github.com/LidenbrockGit/url-shortener/internal/infrastructure/config"
	"github.com/LidenbrockGit/url-shortener/internal/infrastructure/db/dbmemory"
	"github.com/LidenbrockGit/url-shortener/internal/infrastructure/server/defserver"
	"github.com/LidenbrockGit/url-shortener/internal/usecases/repos/linkrepo"
	"github.com/LidenbrockGit/url-shortener/internal/usecases/repos/userrepo"
	"github.com/LidenbrockGit/url-shortener/internal/usecases/search/fullurlsearch"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()

	userDB := dbmemory.NewUserRepo()
	user := userrepo.NewUserRepo(userDB)

	linkDB := dbmemory.NewLinkRepo()
	link := linkrepo.NewLinkRepo(linkDB)
	fullUrlSearch := fullurlsearch.NewFullUrl(linkDB)

	hds := &handlers.Handlers{
		Userrepo:      user,
		Linkrepo:      link,
		Fullurlsearch: fullUrlSearch,
	}
	router := routergin.NewRouter(hds)
	server := defserver.NewServer(":"+os.Getenv("PORT"), router)
	server.Start()

	<-ctx.Done()
	server.Stop()
}
