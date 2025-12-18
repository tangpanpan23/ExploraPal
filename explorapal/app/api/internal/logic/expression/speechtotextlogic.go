package expression

import (
	"context"

	"explorapal/app/api/internal/svc"
	"explorapal/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SpeechToTextLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 语音转文字
func NewSpeechToTextLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SpeechToTextLogic {
	return &SpeechToTextLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SpeechToTextLogic) SpeechToText(req *types.SpeechToTextReq) (resp *types.SpeechToTextResp, err error) {
	// todo: add your logic here and delete this line

	return
}
