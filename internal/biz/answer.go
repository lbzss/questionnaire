package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Answer struct {
	Id              int64
	AnswerChoice    int
	AnswerText      string
	QuestionnaireId int64
	QuestionId      int64
}

type AnswerRepo interface {
	Save(context.Context, *Answer) (*Answer, error)
	SaveBulk(context.Context, []*Answer) ([]*Answer, error)
	Update(context.Context, *Answer) (*Answer, error)
	FindByQuestionnaireIDAndQuestionID(context.Context, int64, int64) ([]*Answer, error)
	ListAll(context.Context) ([]*Answer, error)
}

type AnswerUsecase struct {
	repo AnswerRepo
	log  *log.Helper
}

func NewAnswerUsecase(repo AnswerRepo, logger log.Logger) *AnswerUsecase {
	return &AnswerUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AnswerUsecase) SubmitAnswer(ctx context.Context, q *Answer) (*Answer, error) {
	uc.log.WithContext(ctx).Infof("CreateQuestionnaire: %v", q.Id)
	return uc.repo.Save(ctx, q)
}

func (uc *AnswerUsecase) SubmitAnswerBulk(ctx context.Context, qs []*Answer) ([]*Answer, error) {
	return uc.repo.SaveBulk(ctx, qs)
}

func (uc *AnswerUsecase) GetAnswerWithQuestionnaireIdAndQuestionId(ctx context.Context, questionnaireId, questionId int64) ([]*Answer, error) {
	return uc.repo.FindByQuestionnaireIDAndQuestionID(ctx, questionnaireId, questionId)
}
