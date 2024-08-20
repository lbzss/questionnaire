package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// Questionnaire is a Questionnaire model.
type Questionnaire struct {
	Id          int
	Name        string
	Description string
	Object      string
}

// QuestionnaireRepo is a Questionnaire repo.
type QuestionnaireRepo interface {
	Save(context.Context, *Questionnaire) (*Questionnaire, error)
	Update(context.Context, *Questionnaire) (*Questionnaire, error)
	FindByID(context.Context, int) (*Questionnaire, error)
	ListAll(context.Context) ([]*Questionnaire, error)
	DeleteByID(context.Context, int) (*Questionnaire, error)
}

// QuestionnaireUsecase is a Questionnaire usecase.
type QuestionnaireUsecase struct {
	repo QuestionnaireRepo
	log  *log.Helper
}

// NewQuestionnaireUsecase new a Questionnaire usecase.
func NewQuestionnaireUsecase(repo QuestionnaireRepo, logger log.Logger) *QuestionnaireUsecase {
	return &QuestionnaireUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateQuestionnaire creates a Questionnaire, and returns the new Questionnaire.
func (uc *QuestionnaireUsecase) CreateQuestionnaire(ctx context.Context, q *Questionnaire) (*Questionnaire, error) {
	uc.log.WithContext(ctx).Infof("CreateQuestionnaire: %v", q.Id)
	return uc.repo.Save(ctx, q)
}
func (uc *QuestionnaireUsecase) GetQuestionnaireById(ctx context.Context, id int) (*Questionnaire, error) {
	uc.log.WithContext(ctx).Infof("GetQuestionnaireById: %v", id)
	return uc.repo.FindByID(ctx, id)
}
func (uc *QuestionnaireUsecase) ListQuestionnaire(ctx context.Context) ([]*Questionnaire, error) {
	uc.log.WithContext(ctx).Infof("ListQuestionnaire: %v")
	return uc.repo.ListAll(ctx)
}
func (uc *QuestionnaireUsecase) UpdateQuestionnaire(ctx context.Context, q *Questionnaire) (*Questionnaire, error) {
	uc.log.WithContext(ctx).Infof("UpdateQuestionnaire: %v", q)
	return uc.repo.Update(ctx, q)
}
func (uc *QuestionnaireUsecase) DeleteQuestionnaireById(ctx context.Context, id int) (*Questionnaire, error) {
	uc.log.WithContext(ctx).Infof("DeleteQuestionnaireById: %v", id)
	return uc.repo.DeleteByID(ctx, id)
}
