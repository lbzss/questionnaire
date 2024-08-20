package data

import (
	"context"
	"questionnaire/ent/question"
	"questionnaire/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type questionRepo struct {
	data *Data
	log  *log.Helper
}

// NewQuestionnaireRepo .
func NewQuestionRepo(data *Data, logger log.Logger) biz.QuestionRepo {
	return &questionRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *questionRepo) Save(ctx context.Context, g *biz.Question) (*biz.Question, error) {
	result, err := r.data.db.Question.Create().
		SetType(int(g.Type)).SetQuestion(g.Question).SetQuestionnaireID(int64(g.QuestionnaireId)).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Question{
		Id:              result.ID,
		Type:            int32(result.Type),
		Question:        result.Question,
		QuestionnaireId: int(result.QuestionnaireID),
	}, nil
}

func (r *questionRepo) Update(ctx context.Context, g *biz.Question) (*biz.Question, error) {
	result, err := r.data.db.Question.UpdateOneID(g.Id).
		SetType(int(g.Type)).SetQuestion(g.Question).SetQuestionnaireID(int64(g.QuestionnaireId)).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Question{
		Id:              result.ID,
		Type:            int32(result.Type),
		Question:        result.Question,
		QuestionnaireId: int(result.QuestionnaireID),
	}, nil
}

func (r *questionRepo) FindByQuestionnaireID(ctx context.Context, id int) ([]*biz.Question, error) {
	results, err := r.data.db.Question.Query().Where(question.QuestionnaireIDEQ(int64(id))).All(ctx)
	if err != nil {
		return nil, err
	}
	var questions []*biz.Question
	for _, question := range results {
		questions = append(questions, &biz.Question{
			Id:              question.ID,
			QuestionnaireId: int(question.QuestionnaireID),
			Type:            int32(question.Type),
			Question:        question.Question,
		})
	}
	return questions, nil
}

func (r *questionRepo) ListAll(ctx context.Context) ([]*biz.Question, error) {
	results, err := r.data.db.Question.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	var questions []*biz.Question
	for _, question := range results {
		questions = append(questions, &biz.Question{
			Id:              question.ID,
			QuestionnaireId: int(question.QuestionnaireID),
			Type:            int32(question.Type),
			Question:        question.Question,
		})
	}
	return questions, nil
}

func (r *questionRepo) DeleteByID(ctx context.Context, id int) (*biz.Question, error) {
	result, err := r.data.db.Question.Query().Where(question.IDEQ(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	count, err := r.data.db.Question.Delete().Where(question.IDEQ(id)).Exec(ctx)
	if err != nil || count == 0 {
		return nil, err
	}
	return &biz.Question{
		Id:              result.ID,
		Type:            int32(result.Type),
		Question:        result.Question,
		QuestionnaireId: int(result.QuestionnaireID),
	}, nil
}
