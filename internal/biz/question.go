package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// Question is a Question model.
type Question struct {
	Id              int
	QuestionnaireId int
	Type            int32
	Question        string
}

type QuestionRepo interface {
	Save(context.Context, *Question) (*Question, error)
	Update(context.Context, *Question) (*Question, error)
	FindByQuestionnaireID(context.Context, int) ([]*Question, error)
	ListAll(context.Context) ([]*Question, error)
	DeleteByID(context.Context, int) (*Question, error)
	GetByID(context.Context, int64, int64) (*Question, error)
}

type QuestionUsecase struct {
	repo QuestionRepo
	log  *log.Helper
}

func NewQuestionUsecase(repo QuestionRepo, logger log.Logger) *QuestionUsecase {
	return &QuestionUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *QuestionUsecase) CreateQuestion(ctx context.Context, q *Question) (*Question, error) {
	uc.log.WithContext(ctx).Infof("CreateQuestionnaire: %v", q.Id)
	return uc.repo.Save(ctx, q)
}
func (uc *QuestionUsecase) GetQuestionByQuestionnaireId(ctx context.Context, id int) ([]*Question, error) {
	uc.log.WithContext(ctx).Infof("GetQuestionnaireById: %v", id)
	return uc.repo.FindByQuestionnaireID(ctx, id)
}
func (uc *QuestionUsecase) GetQuestionByQuestionnaireIdAndQuestionId(ctx context.Context, questionnaireId, questionId int64) (*Question, error) {
	//uc.log.WithContext(ctx).Infof("GetQuestionnaireById: %v", id)
	return uc.repo.GetByID(ctx, questionnaireId, questionId)
}
func (uc *QuestionUsecase) ListQuestions(ctx context.Context) ([]*Question, error) {
	uc.log.WithContext(ctx).Infof("ListQuestionnaire: %v")
	return uc.repo.ListAll(ctx)
}
func (uc *QuestionUsecase) UpdateQuestion(ctx context.Context, q *Question) (*Question, error) {
	uc.log.WithContext(ctx).Infof("UpdateQuestionnaire: %v", q)
	return uc.repo.Update(ctx, q)
}
func (uc *QuestionUsecase) DeleteQuestionById(ctx context.Context, id int) (*Question, error) {
	uc.log.WithContext(ctx).Infof("DeleteQuestionnaireById: %v", id)
	return uc.repo.DeleteByID(ctx, id)
}
