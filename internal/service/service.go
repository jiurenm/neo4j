package service

import (
	"fmt"
	"log"
	"neo4j/internal/dao"
)

type Service struct {
	db *dao.DB
}

func New(db *dao.DB) *Service {
	return &Service{db: db}
}

func (s *Service) FindAllName() []string {
	result, err := s.db.FindName()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return result
}

func (s *Service) Find(name string, relationship string) string {
	name = fmt.Sprintf(".*%s", name)
	switch relationship {
	case "CAUSED":
		result, err := s.db.FindCauseByName(name)
		if err != nil {
			log.Fatal(err)
			return ""
		}
		return result
	case "DIAGNOSED":
		result, err := s.db.FindDiagnosisByName(name)
		if err != nil {
			log.Fatal(err)
			return ""
		}
		return result
	case "HAS_HARM":
		result, err := s.db.FindHarmByName(name)
		if err != nil {
			log.Fatal(err)
			return ""
		}
		return result
	case "HAS_NOTE":
		result, err := s.db.FindNoteByName(name)
		if err != nil {
			log.Fatal(err)
			return ""
		}
		return result
	case "PREVENTED":
		result, err := s.db.FindPreventionByName(name)
		if err != nil {
			log.Fatal(err)
			return ""
		}
		return result
	case "HAS_SYMPTOM":
		result, err := s.db.FindSymptomByName(name)
		if err != nil {
			log.Fatal(err)
			return ""
		}
		return result
	case "TREATED":
		result, err := s.db.FindTreatmentByName(name)
		if err != nil {
			log.Fatal(err)
			return ""
		}
		return result
	case "INFO":
		result, err := s.db.FindByName(name)
		if err != nil {
			log.Fatal(err)
			return ""
		}
		return result
	default:
		return ""
	}
}
