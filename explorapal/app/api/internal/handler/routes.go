package handler

import (
	"net/http"

	"explorapal/app/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

// RegisterHandlers 注册所有路由处理器
func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	// 项目管理路由组
	registerProjectHandlers(server, serverCtx)

	// 观察阶段路由组
	registerObservationHandlers(server, serverCtx)

	// 提问引导路由组
	registerQuestioningHandlers(server, serverCtx)

	// 表达阶段路由组
	registerExpressionHandlers(server, serverCtx)

	// 成果生成路由组
	registerAchievementHandlers(server, serverCtx)
}

// registerProjectHandlers 注册项目管理路由
func registerProjectHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtAuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/project/create",
					Handler: CreateProjectHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/project/list",
					Handler: GetProjectListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/project/detail",
					Handler: GetProjectDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/project/status/update",
					Handler: UpdateProjectStatusHandler(serverCtx),
				},
			}...,
		),
	)
}

// registerObservationHandlers 注册观察阶段路由
func registerObservationHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtAuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/observation/image/recognize",
					Handler: RecognizeImageHandler(serverCtx),
				},
			}...,
		),
	)
}

// registerQuestioningHandlers 注册提问引导路由
func registerQuestioningHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtAuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/questioning/questions/generate",
					Handler: GenerateQuestionsHandler(serverCtx),
				},
			}...,
		),
	)
}

// registerExpressionHandlers 注册表达阶段路由
func registerExpressionHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtAuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/expression/note/polish",
					Handler: PolishNoteHandler(serverCtx),
				},
			}...,
		),
	)
}

// registerAchievementHandlers 注册成果生成路由
func registerAchievementHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtAuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/achievement/report/generate",
					Handler: GenerateReportHandler(serverCtx),
				},
			}...,
		),
	)
}
