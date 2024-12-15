package logic

import (
	"context"

	"tree-cloud/core/helper"
	"tree-cloud/core/internal/svc"
	"tree-cloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendLogic {
	return &MailCodeSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendLogic) MailCodeSend(req *types.MailCodeSendRequest) (resp *types.MailCodeSendReply, err error) {
	err = helper.MailSendCode(req.Email, "123456")
	if err != nil {
		return nil, err
	}
	return
}
