package routergin

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/LidenbrockGit/url-shortener/internal/entities/linkentity"
	"github.com/LidenbrockGit/url-shortener/internal/entities/userentity"
	"github.com/LidenbrockGit/url-shortener/internal/infrastructure/api/handlers"
	"github.com/LidenbrockGit/url-shortener/internal/infrastructure/api/middlewares/auth"
	"github.com/LidenbrockGit/url-shortener/internal/infrastructure/api/tools/jwttool"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

var _ ServerInterface = &Router{}

type Router struct {
	*gin.Engine
	handlers *handlers.Handlers
}

func NewRouter(h *handlers.Handlers) *Router {
	gr := gin.Default()
	r := &Router{
		handlers: h,
	}

	// Use generated openapi code with middlewares
	RegisterHandlersWithOptions(gr, r, GinServerOptions{
		Middlewares: []MiddlewareFunc{
			func(c *gin.Context) {
				if _, ok := c.Get(JwtScopes); !ok {
					return
				}
				auth.GinAuthMW(c, h.UserRead)
			},
		},
	})

	r.Engine = gr
	return r
}

func (r *Router) GetLinks(ctx *gin.Context) {
	currentUser := getCtxUser(ctx)

	chIn, err := r.handlers.GetLinks()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	var linkJson struct {
		Id         string `json:"id"`
		UserId     string `json:"userId"`
		ShortUrl   string `json:"shortUrl"`
		FullUrl    string `json:"fullUrl"`
		UsageCount int    `json:"usageCount"`
		CreatedAt  string `json:"createdAt"`
	}

	encoder := json.NewEncoder(ctx.Writer)

	ctx.Status(http.StatusOK)
	ctx.Header("content-type", "application/json")
	_, _ = ctx.Writer.WriteString("{")
	defer func() {
		_, _ = ctx.Writer.WriteString("}")
	}()
	for link := range chIn {
		if link.UserId != currentUser.Id {
			continue
		}

		linkJson.Id = link.Id.String()
		linkJson.UserId = link.UserId.String()
		linkJson.ShortUrl = link.ShortUrl
		linkJson.FullUrl = link.FullUrl
		linkJson.UsageCount = link.UsageCount
		linkJson.CreatedAt = link.CreatedAt.String()
		_ = encoder.Encode(linkJson)
	}
}

func (r *Router) PostLinks(ctx *gin.Context) {
	linkDTO := LinkCreate{}
	if err := ctx.BindJSON(&linkDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": CantReadRequestDataErr.Error()})
		return
	}

	switch "" {
	case *linkDTO.ShortUrl, *linkDTO.FullUrl:
		ctx.JSON(http.StatusBadRequest, gin.H{"message": MissingRequiredFieldsErr.Error()})
		return
	}

	// DTO
	link := linkentity.Link{
		UserId:   getCtxUser(ctx).Id,
		ShortUrl: *linkDTO.ShortUrl,
		FullUrl:  *linkDTO.FullUrl,
	}

	link, err := r.handlers.CreateLink(link)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	linkId := link.Id.String()
	linkUserId := link.UserId.String()
	linkCreatedAt := link.CreatedAt.String()
	linkResp := Link{
		CreatedAt:  &linkCreatedAt,
		FullUrl:    &link.FullUrl,
		Id:         &linkId,
		ShortUrl:   &link.ShortUrl,
		UsageCount: &link.UsageCount,
		UserId:     &linkUserId,
	}

	ctx.JSON(http.StatusOK, linkResp)
}

func (r *Router) DeleteLinksLinkId(ctx *gin.Context, linkId string) {
	_, err := r.handlers.DeleteLink(linkId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successful removal",
	})
}

func (r *Router) GetLinksLinkId(ctx *gin.Context, linkId string) {
	link, err := r.handlers.GetLink(linkId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	linkIdResp := link.Id.String()
	linkUserId := link.UserId.String()
	linkCreatedAt := link.CreatedAt.String()
	linkResp := Link{
		CreatedAt:  &linkCreatedAt,
		FullUrl:    &link.FullUrl,
		Id:         &linkIdResp,
		ShortUrl:   &link.ShortUrl,
		UsageCount: &link.UsageCount,
		UserId:     &linkUserId,
	}

	ctx.JSON(http.StatusOK, linkResp)
}

func (r *Router) PutLinksLinkId(ctx *gin.Context, linkId string) {
	linkUUID, err := uuid.Parse(linkId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": IncorrectLinkIdErr.Error()})
		return
	}

	linkDTO := LinkCreate{}
	if err := ctx.BindJSON(&linkDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": CantReadRequestDataErr.Error()})
		return
	}

	switch "" {
	case *linkDTO.ShortUrl, *linkDTO.FullUrl:
		ctx.JSON(http.StatusBadRequest, gin.H{"message": MissingRequiredFieldsErr.Error()})
		return
	}

	// DTO
	link := linkentity.Link{
		Id:       linkUUID,
		ShortUrl: *linkDTO.ShortUrl,
		FullUrl:  *linkDTO.FullUrl,
	}

	err = r.handlers.UpdateLink(link)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successful user update",
	})
}

func (r *Router) PostUseUrl(ctx *gin.Context, params PostUseUrlParams) {
	fullUrl, err := r.handlers.UseUrl(params.ShortUrl)
	if err != nil {
		ctx.JSON(http.StatusNoContent, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"fullUrl": fullUrl})
}

func (r *Router) PostLogin(ctx *gin.Context) {
	userDTO := UserLogin{}
	if err := ctx.BindJSON(&userDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": CantReadRequestDataErr.Error()})
		ctx.Abort()
		return
	}

	switch "" {
	case userDTO.Login, userDTO.Password:
		ctx.JSON(http.StatusBadRequest, gin.H{"message": MissingRequiredFieldsErr.Error()})
		ctx.Abort()
		return
	}

	user, err := r.handlers.Login(userDTO.Login, userDTO.Password)
	if err != nil {
		body := gin.H{"message": err.Error()}
		if errors.Is(err, handlers.WrongPasswordErr) {
			ctx.JSON(http.StatusBadRequest, body)
		} else {
			ctx.JSON(http.StatusInternalServerError, body)
		}
		ctx.Abort()
		return
	}

	jwtKey := []byte(viper.GetString("jwt_key"))
	jwt, err := jwttool.GenerateJWT(gin.H{"userId": user.Id.String()}, jwtKey)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": UnhandledInternalError.Error()})
		//TODO: записать ошибку в логи
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successful user creation",
		"token":   jwt,
	})
}

func (r *Router) PostLogout(ctx *gin.Context) {}

func (r *Router) PostRegist(ctx *gin.Context) {
	regUser := UserRegist{}
	if err := ctx.BindJSON(&regUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": CantReadRequestDataErr.Error()})
		ctx.Abort()
		return
	}

	switch "" {
	case regUser.Name, regUser.Login, regUser.Password:
		ctx.JSON(http.StatusBadRequest, gin.H{"message": MissingRequiredFieldsErr.Error()})
		ctx.Abort()
		return
	}

	// DTO
	user := userentity.User{
		Name:     regUser.Name,
		Login:    regUser.Login,
		Password: regUser.Password,
	}

	userId, err := r.handlers.Regist(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successful user creation",
		"userId":  userId.String(),
	})
}

func (r *Router) DeleteAccount(ctx *gin.Context) {
	_, err := r.handlers.UserDelete(getCtxUser(ctx).Id.String())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successful removal",
	})
}

func (r *Router) GetAccount(ctx *gin.Context) {
	user, err := r.handlers.UserRead(getCtxUser(ctx).Id.String())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// DTO
	var userDTO struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Login string `json:"login"`
	}
	userDTO.Id = user.Id.String()
	userDTO.Name = user.Name
	userDTO.Login = user.Login

	ctx.JSON(http.StatusOK, userDTO)
}

func (r *Router) PutAccount(ctx *gin.Context) {
	userDTO := struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}{}
	if err := ctx.BindJSON(&userDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": CantReadRequestDataErr.Error()})
		return
	}

	user := userentity.User{
		Id:       getCtxUser(ctx).Id,
		Name:     userDTO.Name,
		Password: userDTO.Password,
	}

	err := r.handlers.UserUpdate(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successful update",
	})
}
