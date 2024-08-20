package service

import (
	"context"
	pb "questionnaire/api/questionnaire/v1"
	"questionnaire/internal/biz"
)

func (s *QuestionnaireService) CreateQuestion(ctx context.Context, req *pb.CreateQuestionRequest) (*pb.CreateQuestionReply, error) {
	if _, err := s.qnc.GetQuestionnaireById(ctx, int(req.QuestionnaireId)); err != nil {
		return nil, err
	}
	result, err := s.quc.CreateQuestion(ctx, &biz.Question{
		QuestionnaireId: int(req.QuestionnaireId),
		Type:            int32(req.Type),
		Question:        req.Question,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateQuestionReply{
		Data: &pb.QuestionInfo{
			Id:              int64(result.Id),
			QuestionnaireId: req.QuestionnaireId,
			Type:            req.Type,
			Question:        req.Question,
		},
	}, nil
}

func (s *QuestionnaireService) DeleteQuestion(ctx context.Context, req *pb.DeleteQuestionRequest) (*pb.DeleteQuestionReply, error) {
	result, err := s.quc.DeleteQuestionById(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.DeleteQuestionReply{
		Data: &pb.QuestionInfo{
			Id:              int64(result.Id),
			QuestionnaireId: int64(result.QuestionnaireId),
			Type:            pb.QuestionType(result.Type),
			Question:        result.Question,
		},
	}, nil
}
func (s *QuestionnaireService) GetQuestions(ctx context.Context, req *pb.GetQuestionRequest) (*pb.GetQuestionReply, error) {
	results, err := s.quc.GetQuestionByQuestionnaireId(ctx, int(req.QuestionnaireId))
	if err != nil {
		return nil, err
	}
	var data []*pb.QuestionInfo
	for _, result := range results {
		data = append(data, &pb.QuestionInfo{
			Id:              int64(result.Id),
			Question:        result.Question,
			Type:            pb.QuestionType(result.Type),
			QuestionnaireId: int64(result.QuestionnaireId),
		})
	}
	return &pb.GetQuestionReply{Data: data}, nil
}
func (s *QuestionnaireService) ListQuestion(ctx context.Context, req *pb.ListQuestionRequest) (*pb.ListQuestionReply, error) {
	results, err := s.quc.ListQuestions(ctx)
	if err != nil {
		return nil, err
	}
	var data []*pb.QuestionInfo
	for _, result := range results {
		data = append(data, &pb.QuestionInfo{
			Id:              int64(result.Id),
			Question:        result.Question,
			Type:            pb.QuestionType(result.Type),
			QuestionnaireId: int64(result.QuestionnaireId),
		})
	}
	return &pb.ListQuestionReply{Data: data}, nil
}
func (s *QuestionnaireService) UpdateQuestion(ctx context.Context, req *pb.UpdateQuestionRequest) (*pb.UpdateQuestionReply, error) {
	result, err := s.quc.UpdateQuestion(ctx, &biz.Question{
		Id:              int(req.Id),
		QuestionnaireId: int(req.QuestionnaireId),
		Type:            int32(req.Type),
		Question:        req.Question,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateQuestionReply{
		Data: &pb.QuestionInfo{
			Id:              int64(result.Id),
			QuestionnaireId: req.QuestionnaireId,
			Type:            req.Type,
			Question:        req.Question,
		},
	}, nil
}
