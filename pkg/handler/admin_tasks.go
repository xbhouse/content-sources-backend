package handler

import (
	"net/http"

	"github.com/content-services/content-sources-backend/pkg/api"
	"github.com/content-services/content-sources-backend/pkg/dao"
	ce "github.com/content-services/content-sources-backend/pkg/errors"
	"github.com/content-services/content-sources-backend/pkg/feature_service_client"
	"github.com/content-services/content-sources-backend/pkg/rbac"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type AdminTaskHandler struct {
	DaoRegistry dao.DaoRegistry
	fsClient    feature_service_client.FeatureServiceClient
}

func checkAccessible(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := CheckAdminTaskAccessible(c.Request().Context()); err != nil {
			return err
		}
		return next(c)
	}
}

func RegisterAdminTaskRoutes(engine *echo.Group, daoReg *dao.DaoRegistry, fsClient *feature_service_client.FeatureServiceClient) {
	if engine == nil {
		panic("engine is nil")
	}
	if daoReg == nil {
		panic("taskInfoReg is nil")
	}
	if fsClient == nil {
		panic("adminClient is nil")
	}

	adminTaskHandler := AdminTaskHandler{
		DaoRegistry: *daoReg,
		fsClient:    *fsClient,
	}
	addRepoRoute(engine, http.MethodGet, "/admin/tasks/", adminTaskHandler.listTasks, rbac.RbacVerbRead, checkAccessible)
	addRepoRoute(engine, http.MethodGet, "/admin/tasks/:uuid", adminTaskHandler.fetch, rbac.RbacVerbRead, checkAccessible)
	addRepoRoute(engine, http.MethodGet, "/admin/features/", adminTaskHandler.listFeatures, rbac.RbacVerbRead, checkAccessible)
}

func (adminTaskHandler *AdminTaskHandler) listTasks(c echo.Context) error {
	pageData := ParsePagination(c)
	filterData := ParseAdminTaskFilters(c)

	tasks, totalTasks, err := adminTaskHandler.DaoRegistry.AdminTask.List(c.Request().Context(), pageData, filterData)
	if err != nil {
		return ce.NewErrorResponse(ce.HttpCodeForDaoError(err), "Error listing tasks", err.Error())
	}

	return c.JSON(http.StatusOK, setCollectionResponseMetadata(&tasks, c, totalTasks))
}

func (adminTaskHandler *AdminTaskHandler) fetch(c echo.Context) error {
	id := c.Param("uuid")

	response, err := adminTaskHandler.DaoRegistry.AdminTask.Fetch(c.Request().Context(), id)
	if err != nil {
		return ce.NewErrorResponse(ce.HttpCodeForDaoError(err), "Error fetching task", err.Error())
	}
	return c.JSON(http.StatusOK, response)
}

func (adminTaskHandler *AdminTaskHandler) listFeatures(c echo.Context) error {
	resp, statusCode, err := adminTaskHandler.fsClient.ListFeatures(c.Request().Context())
	if err != nil {
		return ce.NewErrorResponse(statusCode, "Error listing features", err.Error())
	}

	subsAsFeatResp := api.ListFeaturesResponse{}
	for _, content := range resp.Content {
		subsAsFeatResp.Features = append(subsAsFeatResp.Features, content.Name)
	}

	return c.JSON(http.StatusOK, subsAsFeatResp)
}

func ParseAdminTaskFilters(c echo.Context) api.AdminTaskFilterData {
	filterData := api.AdminTaskFilterData{
		AccountId: DefaultAccountId,
		OrgId:     DefaultOrgId,
		Status:    DefaultStatus,
	}
	err := echo.QueryParamsBinder(c).
		String("account_id", &filterData.AccountId).
		String("org_id", &filterData.OrgId).
		String("status", &filterData.Status).
		String("type", &filterData.Typename).
		BindError()

	if err != nil {
		log.Ctx(c.Request().Context()).Info().Err(err).Msg("error parsing filters")
	}

	return filterData
}
