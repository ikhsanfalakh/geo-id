package service

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ikhsanfalakh/geo-id/internal/model"
)

type LocationService struct {
	DataDir string
}

func NewLocationService(dataDir string) *LocationService {
	return &LocationService{DataDir: dataDir}
}

func (s *LocationService) readJSON(path string, v interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewDecoder(file).Decode(v)
}

func (s *LocationService) GetStates() ([]model.Region, error) {
	var regions []model.Region
	path := filepath.Join(s.DataDir, "states.json")
	if err := s.readJSON(path, &regions); err != nil {
		return nil, err
	}
	return regions, nil
}

func (s *LocationService) GetState(code string) (*model.Region, error) {
	states, err := s.GetStates()
	if err != nil {
		return nil, err
	}
	for _, state := range states {
		if state.Code == code {
			return &state, nil
		}
	}
	return nil, fmt.Errorf("state not found")
}

func (s *LocationService) GetCities(stateCode string) ([]model.Region, error) {
	var regions []model.Region
	path := filepath.Join(s.DataDir, "cities", stateCode+".json")
	if err := s.readJSON(path, &regions); err != nil {
		return nil, err
	}
	return regions, nil
}

func (s *LocationService) GetCity(code string) (*model.Region, error) {
	// This is inefficient without a mapping or knowing the parent state.
	// The original API might have a better way or just iterates everything?
	// Based on file structure cities/[state_id].json, we can't easily find a city by ID without checking all state files.
	// However, the original API endpoint /cities/:id implies it's possible.
	// For now, let's implement a search across all available city files if we have to, 
	// OR we assume the user provides parent ID in a real app, but the API spec says /cities/:id.
	
	// Optimization: In a real app we would load this into memory or a DB.
	// For this file-based approach, we might need to walk the directory.
	
	files, err := filepath.Glob(filepath.Join(s.DataDir, "cities", "*.json"))
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		var regions []model.Region
		if err := s.readJSON(file, &regions); err == nil {
			for _, r := range regions {
				if r.Code == code {
					return &r, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("city not found")
}

func (s *LocationService) GetDistricts(cityCode string) ([]model.Region, error) {
	var regions []model.Region
	path := filepath.Join(s.DataDir, "districts", cityCode+".json")
	if err := s.readJSON(path, &regions); err != nil {
		return nil, err
	}
	return regions, nil
}

func (s *LocationService) GetDistrict(code string) (*model.Region, error) {
	files, err := filepath.Glob(filepath.Join(s.DataDir, "districts", "*.json"))
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		var regions []model.Region
		if err := s.readJSON(file, &regions); err == nil {
			for _, r := range regions {
				if r.Code == code {
					return &r, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("district not found")
}

func (s *LocationService) GetVillages(districtCode string) ([]model.Region, error) {
	var regions []model.Region
	path := filepath.Join(s.DataDir, "villages", districtCode+".json")
	if err := s.readJSON(path, &regions); err != nil {
		return nil, err
	}
	return regions, nil
}

func (s *LocationService) GetVillage(code string) (*model.Region, error) {
	files, err := filepath.Glob(filepath.Join(s.DataDir, "villages", "*.json"))
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		var regions []model.Region
		if err := s.readJSON(file, &regions); err == nil {
			for _, r := range regions {
				if r.Code == code {
					return &r, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("village not found")
}
