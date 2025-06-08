package service

import (
	"context"

	"github.com/adieos/imk-backend/dto"
	"github.com/adieos/imk-backend/entity"
	"github.com/adieos/imk-backend/repository"
	"github.com/google/uuid"
)

type (
	BSService interface {
		CreateBS(ctx context.Context, req dto.BSCreateRequest, userID string) (dto.BSResponse, error)
		GetBSById(ctx context.Context, bsId string) (dto.BSResponse, error)
		GetAllBS(ctx context.Context) ([]dto.BSResponse, error)
		GetAllBSByUserId(ctx context.Context, userId string) ([]dto.BSResponse, error)
		UpdateBS(ctx context.Context, req dto.BSUpdateRequest) (dto.BSResponse, error)
		ChangeStatusBS(ctx context.Context, bsId string, status string) (dto.BSResponse, error)
	}

	bSService struct {
		BSRepo     repository.BSRepository
		jwtService JWTService
	}
)

func NewBSService(BSRepo repository.BSRepository,
	jwtService JWTService) BSService {
	return &bSService{
		BSRepo:     BSRepo,
		jwtService: jwtService,
	}
}

func (s *bSService) CreateBS(ctx context.Context, req dto.BSCreateRequest, userID string) (dto.BSResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	bs := entity.BankSampah{
		Name:    req.Name,
		Address: req.Address,
		City:    req.City,
		Contact: req.Phone,
		UserID:  uuid.MustParse(userID),
	}

	bs, err := s.BSRepo.CreateBS(ctx, nil, bs)
	if err != nil {
		return dto.BSResponse{}, err
	}

	response := dto.BSResponse{
		Id:          bs.ID.String(),
		Name:        bs.Name,
		Address:     bs.Address,
		City:        bs.City,
		Phone:       bs.Contact,
		OpenHour:    bs.OpenHours,
		Description: bs.Description,
		AcceptAll:   bs.AcceptAll,
	}

	return response, nil
}

func (s *bSService) GetBSById(ctx context.Context, bsId string) (dto.BSResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	bs, err := s.BSRepo.GetBSById(ctx, nil, bsId)
	if err != nil {
		return dto.BSResponse{}, err
	}

	acc, err := s.BSRepo.GetBSAccepts(ctx, nil, bsId)
	if err != nil {
		return dto.BSResponse{}, err
	}

	acceptedWasteTypes := make([]dto.BSWasteType, 0, len(acc))
	for _, a := range acc {
		acceptedWasteTypes = append(acceptedWasteTypes, dto.BSWasteType{
			CategoryID: a.CategoryID.String(),
			Price:      a.Price,
			Quota:      a.Quota,
			Filled:     a.Filled,
		})
	}

	response := dto.BSResponse{
		Id:                 bs.ID.String(),
		Name:               bs.Name,
		Address:            bs.Address,
		City:               bs.City,
		Phone:              bs.Contact,
		OpenHour:           bs.OpenHours,
		Description:        bs.Description,
		AcceptAll:          bs.AcceptAll,
		AcceptedWasteTypes: acceptedWasteTypes,
		Status:             bs.Status,
	}

	return response, nil
}

func (s *bSService) GetAllBS(ctx context.Context) ([]dto.BSResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	bss, err := s.BSRepo.GetAllBS(ctx, nil)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.BSResponse, 0, len(bss))
	for _, bs := range bss {
		acc, err := s.BSRepo.GetBSAccepts(ctx, nil, bs.ID.String())
		if err != nil {
			return nil, err
		}

		acceptedWasteTypes := make([]dto.BSWasteType, 0, len(acc))
		for _, a := range acc {
			acceptedWasteTypes = append(acceptedWasteTypes, dto.BSWasteType{
				CategoryID: a.CategoryID.String(),
				Price:      a.Price,
				Quota:      a.Quota,
				Filled:     a.Filled,
			})
		}

		response := dto.BSResponse{
			Id:                 bs.ID.String(),
			Name:               bs.Name,
			Address:            bs.Address,
			City:               bs.City,
			Phone:              bs.Contact,
			OpenHour:           bs.OpenHours,
			Description:        bs.Description,
			AcceptAll:          bs.AcceptAll,
			Status:             bs.Status,
			AcceptedWasteTypes: acceptedWasteTypes,
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (s *bSService) GetAllBSByUserId(ctx context.Context, userId string) ([]dto.BSResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	bss, err := s.BSRepo.GetAllBSByUserId(ctx, nil, userId)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.BSResponse, 0, len(bss))
	for _, bs := range bss {
		acc, err := s.BSRepo.GetBSAccepts(ctx, nil, bs.ID.String())
		if err != nil {
			return nil, err
		}

		acceptedWasteTypes := make([]dto.BSWasteType, 0, len(acc))
		for _, a := range acc {
			acceptedWasteTypes = append(acceptedWasteTypes, dto.BSWasteType{
				CategoryID: a.CategoryID.String(),
				Price:      a.Price,
				Quota:      a.Quota,
				Filled:     a.Filled,
			})
		}

		response := dto.BSResponse{
			Id:                 bs.ID.String(),
			Name:               bs.Name,
			Address:            bs.Address,
			City:               bs.City,
			Phone:              bs.Contact,
			OpenHour:           bs.OpenHours,
			Description:        bs.Description,
			AcceptAll:          bs.AcceptAll,
			Status:             bs.Status,
			AcceptedWasteTypes: acceptedWasteTypes,
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (s *bSService) UpdateBS(ctx context.Context, req dto.BSUpdateRequest) (dto.BSResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	bs, err := s.BSRepo.GetBSById(ctx, nil, req.Id)
	if err != nil {
		return dto.BSResponse{}, err
	}

	bs.Name = req.Name
	bs.OpenHours = req.OpenHour
	bs.Contact = req.Phone
	bs.Description = req.Description
	bs.AcceptAll = req.AcceptAll

	// Update accepted waste types if provided
	if len(req.AcceptedWasteTypes) > 0 {
		for _, wasteType := range req.AcceptedWasteTypes {
			acc := entity.Accept{
				CategoryID:   uuid.MustParse(wasteType.CategoryID),
				Price:        wasteType.Price,
				Quota:        wasteType.Quota,
				Filled:       wasteType.Filled,
				BankSampahID: bs.ID,
			}
			_, err := s.BSRepo.CreateBSAccept(ctx, nil, acc)
			if err != nil {
				return dto.BSResponse{}, err
			}
		}
	} else {
		// If no accepted waste types are provided, set AcceptAll to true
		bs.AcceptAll = true
	}
	if _, err := s.BSRepo.UpdateBS(ctx, nil, bs); err != nil {
		return dto.BSResponse{}, err
	}
	response := dto.BSResponse{
		Id:          bs.ID.String(),
		Name:        bs.Name,
		Address:     bs.Address,
		City:        bs.City,
		Phone:       bs.Contact,
		OpenHour:    bs.OpenHours,
		Description: bs.Description,
		AcceptAll:   bs.AcceptAll,
		Status:      bs.Status,
	}
	return response, nil
}

func (s *bSService) ChangeStatusBS(ctx context.Context, bsId string, status string) (dto.BSResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	bs, err := s.BSRepo.ChangeStatusBS(ctx, nil, bsId, status)
	if err != nil {
		return dto.BSResponse{}, err
	}

	response := dto.BSResponse{
		Id:          bs.ID.String(),
		Name:        bs.Name,
		Address:     bs.Address,
		City:        bs.City,
		Phone:       bs.Contact,
		OpenHour:    bs.OpenHours,
		Description: bs.Description,
		AcceptAll:   bs.AcceptAll,
		Status:      bs.Status,
	}

	return response, nil
}
