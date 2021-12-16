package controller

import (
	"errors"
	"foodies/gosrc/cache"
	"foodies/gosrc/httputil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CacheRefresh godoc
// @Summary      Refresh cache from data.sfgov.org
// @Description  Refresh cache from data.sfgov.org
// @Tags         search
// @Accept       json
// @Produce      json
// @Success      200  {object}  httputil.HTTPSuccess
// @Failure      500  {object}  httputil.HTTPError
// @Router       /cache/refresh [get]
func (c *Controller) CacheRefresh(ctx *gin.Context) {
	err := cache.LoadCacheFromURL()
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, errors.New("invalid input"))
		return
	}
	httputil.NewSuccess(ctx, http.StatusOK, "Successfully refreshed cache")
}
