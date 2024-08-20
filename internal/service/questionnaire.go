package service

import (
	"context"
	"questionnaire/internal/biz"

	pb "questionnaire/api/questionnaire/v1"
)

func (s *QuestionnaireService) CreateQuestionnaire(ctx context.Context, req *pb.CreateQuestionnaireRequest) (*pb.CreateQuestionnaireReply, error) {
	result, err := s.qnc.CreateQuestionnaire(ctx, &biz.Questionnaire{
		Name:        req.Name,
		Description: req.Description,
		Object:      req.Object,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateQuestionnaireReply{
		Data: &pb.QuestionnaireInfo{
			Id:          int64(result.Id),
			Name:        result.Name,
			Description: result.Description,
			Object:      result.Object,
		},
	}, nil
}
func (s *QuestionnaireService) GetQuestionnaire(ctx context.Context, req *pb.GetQuestionnaireRequest) (*pb.GetQuestionnaireReply, error) {
	result, err := s.qnc.GetQuestionnaireById(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.GetQuestionnaireReply{
		Data: &pb.QuestionnaireInfo{
			Id:          int64(result.Id),
			Name:        result.Name,
			Description: result.Description,
			Object:      result.Object,
		},
	}, nil
}
func (s *QuestionnaireService) ListQuestionnaire(ctx context.Context, req *pb.ListQuestionnaireRequest) (*pb.ListQuestionnaireReply, error) {
	results, err := s.qnc.ListQuestionnaire(ctx)
	if err != nil {
		return nil, err
	}
	var data []*pb.QuestionnaireInfo
	for _, result := range results {
		data = append(data, &pb.QuestionnaireInfo{
			Id:          int64(result.Id),
			Name:        result.Name,
			Description: result.Description,
			Object:      result.Object,
		})
	}
	return &pb.ListQuestionnaireReply{Data: data}, nil
}
func (s *QuestionnaireService) UpdateQuestionnaire(ctx context.Context, req *pb.UpdateQuestionnaireRequest) (*pb.UpdateQuestionnaireReply, error) {
	result, err := s.qnc.UpdateQuestionnaire(ctx, &biz.Questionnaire{
		Id:          int(req.Id),
		Name:        req.Name,
		Description: req.Description,
		Object:      req.Object,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateQuestionnaireReply{
		Data: &pb.QuestionnaireInfo{
			Id:          int64(result.Id),
			Name:        result.Name,
			Description: result.Description,
			Object:      result.Object,
		},
	}, nil
}
func (s *QuestionnaireService) DeleteQuestionnaire(ctx context.Context, req *pb.DeleteQuestionnaireRequest) (*pb.DeleteQuestionnaireReply, error) {
	result, err := s.qnc.DeleteQuestionnaireById(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.DeleteQuestionnaireReply{
		Data: &pb.QuestionnaireInfo{
			Id:          int64(result.Id),
			Name:        result.Name,
			Description: result.Description,
			Object:      result.Object,
		},
	}, nil
}
