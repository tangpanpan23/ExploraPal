package svc

import (
	"net/http"

	"explorapal/app/api/internal/config"
	"explorapal/app/model/hps"
	"explorapal/third/openai"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config

	// 数据库模型
	ProjectModel         hps.ProjectsModel
	ProjectActivityModel hps.ProjectActivitiesModel
	ObservationModel     hps.ObservationsModel
	QuestionModel        hps.QuestionsModel
	ExpressionModel      hps.ExpressionsModel
	AchievementModel     hps.AchievementsModel

	// AI服务客户端
	AIClient *openai.Client

	// 中间件
	JwtAuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DBConfig.DataSource)

	// 初始化AI客户端
	aiConfig := &openai.Config{
		TAL_MLOPS_APP_ID:  c.AIService.TAL_MLOPS_APP_ID,
		TAL_MLOPS_APP_KEY: c.AIService.TAL_MLOPS_APP_KEY,
		BaseURL:           c.AIService.BaseURL,
		Timeout:           c.AIService.Timeout,
		MaxTokens:         c.AIService.MaxTokens,
		Temperature:       c.AIService.Temperature,
	}
	aiClient := openai.NewClient(aiConfig)

	return &ServiceContext{
		Config: c,

		// 数据库模型
		ProjectModel:         hps.NewProjectsModel(conn, c.Cache),
		ProjectActivityModel: hps.NewProjectActivitiesModel(conn, c.Cache),
		ObservationModel:     hps.NewObservationsModel(conn, c.Cache),
		QuestionModel:        hps.NewQuestionsModel(conn, c.Cache),
		ExpressionModel:      hps.NewExpressionsModel(conn, c.Cache),
		AchievementModel:     hps.NewAchievementsModel(conn, c.Cache),

		// AI服务
		AIClient: aiClient,

		// 中间件
		JwtAuthMiddleware: func(next http.HandlerFunc) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				// 暂时跳过JWT验证，直接调用下一个处理器
				// TODO: 实现JWT认证逻辑
				next(w, r)
			}
		},
	}
}
