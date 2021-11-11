package v1

import (
	"github.com/labstack/echo/v4"
	_ "github.com/rewebcan/ys-memoli/docs"
	"github.com/rewebcan/ys-memoli/internal/entity"
	"github.com/rewebcan/ys-memoli/internal/usecase/repo"
	"net/http"
)

const _internalServerErrorMsg = "something went wrong, error is reported"

type settingsRoutes struct {
	settingsRepository repo.SettingsRepository
}

func newV1settingsRoutes(handler *echo.Echo, settingsRepository repo.SettingsRepository) {
	sv1 := &settingsRoutes{settingsRepository}
	handler.GET("api/v1/settings/:key", sv1.getSetting)
	handler.PUT("api/v1/settings", sv1.setSetting)
}

// @Summary     Set a setting
// @Description An example of how to use memoliDB, to set a setting with a key and value
// @ID          setSetting
// @Tags  	    setSetting
// @Accept 		json
// @Product 	json
// @Param 		setting body entity.StoreSettingRequest true "Setting"
// @Success     204
// @Failure     400,422,500 {object} errResponse{message=string}
// @Router      /api/v1/settings [put].
func (sv1 *settingsRoutes) setSetting(c echo.Context) error {
	req := new(entity.StoreSettingRequest)
	err := c.Bind(&req)
	if err != nil {
		return responseError(c, http.StatusUnprocessableEntity, err, err.Error())
	}
	err = req.Validate()
	if err != nil {
		return responseError(c, http.StatusBadRequest, err, err.Error())
	}

	if err = sv1.settingsRepository.Set(req.Key, req.Value); err != nil {
		return responseError(c, http.StatusInternalServerError, err, _internalServerErrorMsg)
	}

	return responseSuccess(c, http.StatusNoContent, "")
}

// @Summary     Get a setting value from key
// @Description An example of how to use memoliDB, to get a setting value with a key
// @ID          getSetting
// @Tags  	    getSetting
// @Param  	    key path string true "Key"
// @Produce     json
// @Success     200 {object} successResponse{data=entity.SettingsResponse}
// @Failure     500 {object} errResponse{message=string}
// @Router      /api/v1/settings/{key} [get].
func (sv1 *settingsRoutes) getSetting(c echo.Context) error {
	key := c.Param("key")

	value, err := sv1.settingsRepository.Get(key)
	if err != nil {
        return responseError(c, http.StatusNotFound, err, "key is not exist")
    }

	return responseSuccess(c, http.StatusOK, entity.SettingsResponse{
		Key:   key,
		Value: value,
	})
}
