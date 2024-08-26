package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"questionnaire/ent"
	"questionnaire/ent/answer"
	"questionnaire/internal/biz"
)

type answerRepo struct {
	data *Data
	log  *log.Helper
}

func NewAnswerRepo(data *Data, logger log.Logger) biz.AnswerRepo {
	return &answerRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (a *answerRepo) Save(ctx context.Context, answer *biz.Answer) (*biz.Answer, error) {
	result, err := a.data.db.Answer.Create().
		SetQuestionnaireID(answer.QuestionnaireId).SetQuestionID(answer.QuestionId).
		SetAnswerChoice(answer.AnswerChoice).SetAnswerText(answer.AnswerText).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Answer{
		Id:              int64(result.ID),
		QuestionnaireId: result.QuestionnaireID,
		QuestionId:      result.QuestionID,
		AnswerText:      result.AnswerText,
		AnswerChoice:    result.AnswerChoice,
	}, nil
}

func (a *answerRepo) SaveBulk(ctx context.Context, answers []*biz.Answer) ([]*biz.Answer, error) {
	results, err := a.data.db.Answer.MapCreateBulk(answers, func(c *ent.AnswerCreate, i int) {
		c.SetAnswerChoice(answers[i].AnswerChoice).SetAnswerText(answers[i].AnswerText).
			SetQuestionID(answers[i].QuestionId).SetQuestionnaireID(answers[i].QuestionnaireId)
	}).Save(ctx)
	if err != nil {
		return nil, err
	}
	var data []*biz.Answer
	for _, result := range results {
		data = append(data, &biz.Answer{
			Id:              int64(result.ID),
			AnswerChoice:    result.AnswerChoice,
			AnswerText:      result.AnswerText,
			QuestionnaireId: result.QuestionnaireID,
			QuestionId:      result.QuestionID,
		})
	}
	return data, nil
}
func (a *answerRepo) Update(ctx context.Context, answer *biz.Answer) (*biz.Answer, error) {
	return nil, nil
}
func (a *answerRepo) FindByQuestionnaireIDAndQuestionID(ctx context.Context, questionnaireID int64, questionID int64) ([]*biz.Answer, error) {
	results, err := a.data.db.Answer.Query().Where(answer.QuestionIDEQ(questionID), answer.QuestionnaireIDEQ(questionnaireID)).All(ctx)
	if err != nil {
		return nil, err
	}
	var data []*biz.Answer
	for _, result := range results {
		data = append(data, &biz.Answer{
			Id:              int64(result.ID),
			AnswerChoice:    result.AnswerChoice,
			AnswerText:      result.AnswerText,
			QuestionnaireId: result.QuestionnaireID,
			QuestionId:      result.QuestionID,
		})
	}
	return data, nil
}
func (a *answerRepo) ListAll(ctx context.Context) ([]*biz.Answer, error) {
	results, err := a.data.db.Answer.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	var data []*biz.Answer
	for _, result := range results {
		data = append(data, &biz.Answer{
			Id:              int64(result.ID),
			AnswerChoice:    result.AnswerChoice,
			AnswerText:      result.AnswerText,
			QuestionnaireId: result.QuestionnaireID,
			QuestionId:      result.QuestionID,
		})
	}
	return data, nil
}
