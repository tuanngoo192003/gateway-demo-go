package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/adapter/presenter"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/domain/ent"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/domain/repository"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/infra/config"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/infra/middleware"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/usecase/interactor"
)

func NewAuthController(r *gin.Engine, cfg *config.Config, client *ent.Client, jwtSecret string) {
	authRepo := repository.NewAccountRepository(client)
	tokenRepo := repository.NewRefreshTokenRepository(client)

	// 2. Initialize usecase (interactor)
	authUsecase := interactor.NewAuthInteractor(authRepo, tokenRepo, jwtSecret)

	// 3. Initialize presenter (controller adapter)
	authPresenter := presenter.NewAuthPresenter(authUsecase)

	// 4. Register routes
	protected := r.Group("/")
	{
		protected.POST("refresh", middleware.ContextMiddleware([]byte(cfg.JWT.Secret)), authPresenter.RefreshToken())
		protected.POST("login", authPresenter.Login())
	}
}
