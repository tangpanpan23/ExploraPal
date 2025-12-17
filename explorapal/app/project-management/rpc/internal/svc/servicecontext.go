package svc

import (
	"explorapal/app/model/hps"
	"explorapal/app/project-management/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	// 数据库模型
	ProjectModel         hps.ProjectsModel
	ProjectActivityModel hps.ProjectActivitiesModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DBConfig.DataSource)

	return &ServiceContext{
		Config: c,

		ProjectModel:         hps.NewProjectsModel(conn, c.Cache),
		ProjectActivityModel: hps.NewProjectActivitiesModel(conn, c.Cache),
	}
}
