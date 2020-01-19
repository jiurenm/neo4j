package grpc

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	neo4j "neo4j/api"
	service2 "neo4j/internal/service"
)

type Server struct {
	service service2.Service
}

func New(service *service2.Service) *grpc.Server {
	s := Server{service: *service}
	server := grpc.NewServer()
	neo4j.RegisterGraphServiceServer(server, &s)
	return server
}

func (s *Server) Find(ctx context.Context, req *neo4j.Req) (*neo4j.Output, error) {
	if req.Relationship == "INFO" && req.Name == "" {
		names := s.service.FindAllName()
		var results = make([]*neo4j.Result, 0)
		for _, name := range names {
			results = append(results, &neo4j.Result{
				Name: name,
			})
		}
		return &neo4j.Output{Result: results}, nil
	} else {
		var results = make([]*neo4j.Result, 0)
		results = append(results, &neo4j.Result{Info: s.service.Find(req.Name, req.Relationship)})
		return &neo4j.Output{Result: results}, nil
	}
}

func (s *Server) FindByName(context.Context, *neo4j.Input) (*neo4j.Result, error) {
	panic("implement me")
}

func (s *Server) FindCauseByName(context.Context, *neo4j.Input) (*neo4j.Output, error) {
	panic("implement me")
}

func (s *Server) FindDiagnosisByName(context.Context, *neo4j.Input) (*neo4j.Output, error) {
	panic("implement me")
}

func (s *Server) FindHarmByName(context.Context, *neo4j.Input) (*neo4j.Output, error) {
	panic("implement me")
}

func (s *Server) FindNoteByName(context.Context, *neo4j.Input) (*neo4j.Output, error) {
	panic("implement me")
}

func (s *Server) FindPreventionByName(context.Context, *neo4j.Input) (*neo4j.Output, error) {
	panic("implement me")
}

func (s *Server) FindSymptomByName(context.Context, *neo4j.Input) (*neo4j.Output, error) {
	panic("implement me")
}

func (s *Server) FindTreatmentByName(context.Context, *neo4j.Input) (*neo4j.Output, error) {
	panic("implement me")
}
