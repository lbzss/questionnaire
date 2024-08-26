package service

import (
	"context"
	pb "questionnaire/api/questionnaire/v1"
	"questionnaire/internal/biz"
)

func (s *QuestionnaireService) SubmitAnswer(ctx context.Context, req *pb.SubmitAnswerSingleRequest) (*pb.SubmitAnswerSingleReply, error) {
	result, err := s.au.SubmitAnswer(ctx, &biz.Answer{
		AnswerChoice:    int(req.AnswerChoice),
		AnswerText:      req.AnswerText,
		QuestionnaireId: req.QuestionnaireId,
		QuestionId:      req.QuestionId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.SubmitAnswerSingleReply{
		Data: &pb.AnswerInfo{
			Id:              result.Id,
			QuestionnaireId: req.QuestionnaireId,
			QuestionId:      req.QuestionId,
			AnswerText:      req.AnswerText,
			AnswerChoice:    req.AnswerChoice,
		},
	}, nil
}

func (s *QuestionnaireService) SubmitAnswerBulk(ctx context.Context, req *pb.SubmitAnswerBulkRequest) (*pb.SubmitAnswerBulkReply, error) {
	var answersBulk []*biz.Answer
	for _, answerInfo := range req.AnswerBulkInfo {
		answersBulk = append(answersBulk, &biz.Answer{
			AnswerChoice:    int(answerInfo.AnswerChoice),
			AnswerText:      answerInfo.AnswerText,
			QuestionnaireId: req.QuestionnaireId,
			QuestionId:      answerInfo.QuestionId,
		})
	}
	resultsBulk, err := s.au.SubmitAnswerBulk(ctx, answersBulk)
	if err != nil {
		return nil, err
	}
	var answerInfo []*pb.AnswerInfo
	for _, resultBulk := range resultsBulk {
		answerInfo = append(answerInfo, &pb.AnswerInfo{
			AnswerText:      resultBulk.AnswerText,
			AnswerChoice:    int64(resultBulk.AnswerChoice),
			QuestionnaireId: resultBulk.QuestionnaireId,
			QuestionId:      resultBulk.QuestionId,
			Id:              resultBulk.Id,
		})
	}
	return &pb.SubmitAnswerBulkReply{
		Data: answerInfo,
	}, nil
}

func (s *QuestionnaireService) GetAnswerWithQuestionnaireIdAndQuestionId(ctx context.Context,
	req *pb.GetAnswerWithQuestionnaireIdAndQuestionIdRequest) (*pb.GetAnswerWithQuestionnaireIdAndQuestionIdReply,
	error) {
	answers, err := s.au.GetAnswerWithQuestionnaireIdAndQuestionId(ctx, req.QuestionnaireId, req.QuestionId)
	if err != nil {
		return nil, err
	}
	question, err := s.quc.GetQuestionByQuestionnaireIdAndQuestionId(ctx, req.QuestionnaireId, req.QuestionId)
	if err != nil {
		return nil, err
	}

	var answerData []*pb.AnswerData
	for _, answer := range answers {
		answerData = append(answerData, &pb.AnswerData{
			AnswerText:   answer.AnswerText,
			AnswerChoice: int64(answer.AnswerChoice),
		})
	}

	return &pb.GetAnswerWithQuestionnaireIdAndQuestionIdReply{
		Data: &pb.GetAnswerWithQuestionnaireIdAndQuestionIdInfo{
			QuestionInfo: &pb.QuestionInfo{
				Id:              int64(question.Id),
				Question:        question.Question,
				Type:            pb.QuestionType(question.Type),
				QuestionnaireId: int64(question.QuestionnaireId),
			},
			Data: answerData,
		},
	}, nil
}
