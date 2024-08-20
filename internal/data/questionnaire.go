package data

import (
	"context"
	"questionnaire/ent/questionnaire"

	"questionnaire/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type questionnaireRepo struct {
	data *Data
	log  *log.Helper
}

// NewQuestionnaireRepo .
func NewQuestionnaireRepo(data *Data, logger log.Logger) biz.QuestionnaireRepo {
	return &questionnaireRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *questionnaireRepo) Save(ctx context.Context, g *biz.Questionnaire) (*biz.Questionnaire, error) {
	result, err := r.data.db.Questionnaire.Create().
		SetDescription(g.Description).SetName(g.Name).SetAnonymous(g.Id).SetObject(g.Object).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Questionnaire{
		Id:          result.ID,
		Name:        result.Name,
		Description: result.Description,
		Object:      result.Object,
	}, nil
}

func (r *questionnaireRepo) Update(ctx context.Context, g *biz.Questionnaire) (*biz.Questionnaire, error) {
	result, err := r.data.db.Questionnaire.UpdateOneID(g.Id).
		SetName(g.Name).SetDescription(g.Description).SetObject(g.Object).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Questionnaire{
		Id:          result.ID,
		Name:        result.Name,
		Description: result.Description,
		Object:      result.Object,
	}, nil
}

func (r *questionnaireRepo) FindByID(ctx context.Context, id int) (*biz.Questionnaire, error) {
	result, err := r.data.db.Questionnaire.Query().Where(questionnaire.IDEQ(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Questionnaire{
		Id:          result.ID,
		Name:        result.Name,
		Description: result.Description,
		Object:      result.Object,
	}, nil
}

func (r *questionnaireRepo) ListAll(ctx context.Context) ([]*biz.Questionnaire, error) {
	results, err := r.data.db.Questionnaire.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	var questionnaires []*biz.Questionnaire
	for _, result := range results {
		questionnaires = append(questionnaires, &biz.Questionnaire{
			Id:          result.ID,
			Name:        result.Name,
			Description: result.Description,
			Object:      result.Object,
		})
	}
	return questionnaires, nil
}

func (r *questionnaireRepo) DeleteByID(ctx context.Context, id int) (*biz.Questionnaire, error) {
	result, err := r.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	count, err := r.data.db.Questionnaire.Delete().Where(questionnaire.IDEQ(id)).Exec(ctx)
	if err != nil || count == 0 {
		return nil, err
	}
	return &biz.Questionnaire{
		Id:          result.Id,
		Name:        result.Name,
		Description: result.Description,
		Object:      result.Object,
	}, nil
}
