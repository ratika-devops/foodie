package controller

import (
	"errors"
	"foodies/gosrc/cache"
	"foodies/gosrc/httputil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SearchByTruck godoc
// @Summary      Search by food truck name
// @Description  Search by food truck name
// @Tags         search
// @Accept       json
// @Produce      json
// @Param        truck   path      string  true  "Food truck name"
// @Success      200  {object}  []model.FoodTruck
// @Failure      400  {object}  httputil.HTTPError
// @Router       /searchby/truck/{truck} [get]
func (c *Controller) SearchByTruck(ctx *gin.Context) {
	truck := ctx.Param("truck")
	// aid, err := strconv.Atoi(id)
	if len(truck) == 0 {
		httputil.NewError(ctx, http.StatusBadRequest, errors.New("invalid input"))
		return
	}

	searchresult := cache.Search("truck", truck)

	ctx.JSON(http.StatusOK, searchresult)
}

// SearchByCousine godoc
// @Summary      Search by cousine
// @Description  Search by cousine
// @Tags         search
// @Accept       json
// @Produce      json
// @Param        cousine   path      string  true  "Cousine, food served by food truck"
// @Success      200  {object}  []model.FoodTruck
// @Failure      400  {object}  httputil.HTTPError
// @Router       /searchby/cousine/{cousine} [get]
func (c *Controller) SearchByCousine(ctx *gin.Context) {
	cousine := ctx.Param("cousine")
	// aid, err := strconv.Atoi(id)
	if len(cousine) == 0 {
		httputil.NewError(ctx, http.StatusBadRequest, errors.New("invalid input"))
		return
	}

	searchresult := cache.Search("cousine", cousine)

	ctx.JSON(http.StatusOK, searchresult)
}

// SearchByAddress godoc
// @Summary      Search by food truck address
// @Description  Search by food truck address
// @Tags         search
// @Accept       json
// @Produce      json
// @Param        address   path      string  true  "Food Truck Address"
// @Success      200  {object}  []model.FoodTruck
// @Failure      400  {object}  httputil.HTTPError
// @Router       /searchby/address/{address} [get]
func (c *Controller) SearchByAddress(ctx *gin.Context) {
	address := ctx.Param("address")
	// aid, err := strconv.Atoi(id)
	if len(address) == 0 {
		httputil.NewError(ctx, http.StatusBadRequest, errors.New("invalid input"))
		return
	}

	searchresult := cache.Search("address", address)

	ctx.JSON(http.StatusOK, searchresult)
}
