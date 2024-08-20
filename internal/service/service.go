package service

import (
	"github.com/google/wire"
	pb "questionnaire/api/questionnaire/v1"
	"questionnaire/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewQuestionnaireService)

type QuestionnaireService struct {
	pb.UnimplementedQuestionnaireServer

	qnc *biz.QuestionnaireUsecase
	quc *biz.QuestionUsecase
}

func NewQuestionnaireService(qnc *biz.QuestionnaireUsecase, quc *biz.QuestionUsecase) *QuestionnaireService {
	return &QuestionnaireService{qnc: qnc, quc: quc}
}
