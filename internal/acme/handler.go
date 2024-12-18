package acme

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"helloadmin/internal/api"
	"helloadmin/pkg/log"
	"net/http"
)

type Handler struct {
	log *log.Logger
	svc Service
}

func NewHandler(log *log.Logger, svc Service) *Handler {
	return &Handler{
		log: log,
		svc: svc,
	}
}

// CreateAcme godoc
//	@Summary	创建配置
//	@Schemes
//	@Description	创建ACME配置
//	@Tags			ACME模块
//	@Accept			json
//	@Produce		json
//	@Security		Bearer
//	@Param			request	body	CreateRequest	true	"params"
//	@Success		200		{object}Response
//	@Router			/acme/create [post]
func (a *Handler) CreateAcme(ctx *gin.Context) {
	req := new(CreateRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		api.Error(ctx, http.StatusBadRequest, err)
		return
	}
	if err := a.svc.CreateAcme(ctx, req); err != nil {
		a.log.WithContext(ctx).Error("svc.CreateAcme error", zap.Error(err))
		api.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	api.Success(ctx, nil)
}
