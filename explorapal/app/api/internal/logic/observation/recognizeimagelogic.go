package observation

import (
	"context"
	"database/sql"
	"time"

	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"
	"explorapal/app/model/hps"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecognizeImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 识别图片内容
func NewRecognizeImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecognizeImageLogic {
	return &RecognizeImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecognizeImageLogic) RecognizeImage(req *types.RecognizeImageReq) (resp *types.RecognizeImageResp, err error) {
	var observationId int64

	// 如果没有提供observation_id，创建一个新的观察记录
	if req.ObservationId == 0 {
		observationId = time.Now().UnixNano() / 1000000 // 生成观察记录ID

		// 创建观察记录
		observation := &hps.Observations{
			ObservationId: observationId,
			ProjectId:     req.ProjectId,
			UserId:        req.UserId,
			ImageUrl:      req.ImageUrl, // ImageUrl是string类型
			Category:      sql.NullString{String: req.Category, Valid: true}, // Category是sql.NullString类型
			CreateTime:    time.Now(), // 使用CreateTime
			UpdateTime:    time.Now(), // 使用UpdateTime
		}

		// 插入数据库
		_, err = l.svcCtx.ObservationModel.Insert(l.ctx, observation)
		if err != nil {
			l.Errorf("创建观察记录失败: %v", err)
			return nil, err
		}

		l.Infof("自动创建观察记录: ID=%d, ProjectID=%d", observationId, req.ProjectId)
	} else {
		observationId = req.ObservationId
	}

	// 调用AI服务进行图像识别
	prompt := req.Prompt
	if prompt == "" {
		prompt = "请分析这张图片，识别其中的主要物体，描述其特征，并提供相关信息。"
	}

	// 调用AI客户端的图像分析方法
	aiResp, err := l.svcCtx.AIClient.AnalyzeImage(l.ctx, req.ImageUrl, prompt)
	if err != nil {
		l.Errorf("AI图像识别失败: %v", err)
		return nil, err
	}

	// 解析AI响应
	var recognitionResult types.RecognitionResult
	recognitionResult = types.RecognitionResult{
		Description: aiResp.Description,
		Confidence:  aiResp.Confidence,
	}

	// 构建响应
	resp = &types.RecognizeImageResp{
		ObjectName:       aiResp.ObjectName,
		Category:         req.Category,
		Confidence:       aiResp.Confidence,
		Description:      aiResp.Description,
		KeyFeatures:      aiResp.KeyFeatures,
		ScientificName:   aiResp.ScientificName,
		ObservationId:    observationId,
		Recognition:      recognitionResult,
		Suggestions:      []string{"建议观察更多细节", "尝试不同角度拍摄"},
		NextActions:      []string{"记录观察结果", "查找相关资料"},
		InterestingFacts: []string{"这是一个有趣的发现"},
	}

	l.Infof("图像识别完成: ObservationID=%d, Category=%s", observationId, req.Category)

	return resp, nil
}
