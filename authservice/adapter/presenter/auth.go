package presenter

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/adapter/presenter/converter"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/adapter/presenter/writer"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/common/utils"
	"github.com/tuanngoo192003/gateway-demo-go/authservice/usecase"
)

type AuthPresenter interface {
}

type AuthPresenterImpl struct {
	usecase usecase.AuthUsecase
}

func NewAuthPresenter(usecase usecase.AuthUsecase) *AuthPresenterImpl {
	return &AuthPresenterImpl{
		usecase: usecase,
	}
}

func (ap *AuthPresenterImpl) Login() gin.HandlerFunc {
	return utils.InvokeUseCase(
		converter.ToLoginInputType,
		ap.usecase.Login,
		writer.ToLoginOutputType,
	)
}

func (ap *AuthPresenterImpl) RefreshToken() gin.HandlerFunc {
	return utils.InvokeUseCase(
		converter.ToRefreshTokenInputType,
		ap.usecase.RefreshToken,
		writer.ToRefreshTokenOutputType,
	)
}
