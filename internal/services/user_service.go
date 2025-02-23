package services

import (
	"errors"
	"fmt"
	"math"

	"github.com/radenadri/go-boilerplate/internal/delivery/dto/request"
	"github.com/radenadri/go-boilerplate/internal/delivery/dto/response"
	"github.com/radenadri/go-boilerplate/internal/domain/models"
	"github.com/radenadri/go-boilerplate/internal/repositories"
	"github.com/radenadri/go-boilerplate/pkg"
	"github.com/radenadri/go-boilerplate/utils"
)

type UserService struct {
	UserRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (repository *UserService) GetAllUsers(page int, perPage int) (*response.PaginatedResponse[models.User], error) {
	// Count total items
	totalItems, err := repository.UserRepository.Count()
	if err != nil {
		return nil, err
	}

	// Calculate offset
	offset := (page - 1) * perPage

	// Retrieve items for current page
	results, err := repository.UserRepository.FindAll(perPage, offset)
	if err != nil {
		return nil, err
	}

	// Calculate pagination details
	totalPages := int(math.Ceil(float64(totalItems) / float64(perPage)))
	hasNextPage := page < totalPages
	hasPreviousPage := page > 1

	var previousPage *int
	if hasPreviousPage {
		prev := page - 1
		previousPage = &prev
	}

	var nextPage int
	if hasNextPage {
		nextPage = page + 1
	}

	response := &response.PaginatedResponse[models.User]{
		Success: true,
		Data:    results,
		Pagination: response.Pagination{
			TotalItems:      int(totalItems),
			TotalPages:      totalPages,
			CurrentPage:     page,
			ItemsPerPage:    perPage,
			HasNextPage:     hasNextPage,
			HasPreviousPage: hasPreviousPage,
			NextPage:        nextPage,
			PreviousPage:    previousPage,
		},
		Links: response.Links{
			Self:  fmt.Sprintf("/api/items?page=%d&per_page=%d", page, perPage),
			First: fmt.Sprintf("/api/items?page=1&per_page=%d", perPage),
			Last:  fmt.Sprintf("/api/items?page=%d&per_page=%d", totalPages, perPage),
			Next:  fmt.Sprintf("/api/items?page=%d&per_page=%d", nextPage, perPage),
		},
	}

	return response, nil
}

func (s *UserService) Register(userPayload models.User) (*response.UserResponse, error) {
	hashedPassword, err := pkg.HashPassword(userPayload.Password)
	if err != nil {
		return nil, err
	}

	userData := models.User{
		Name:     userPayload.Name,
		Username: userPayload.Username,
		Email:    userPayload.Email,
		Password: hashedPassword,
	}

	if err := s.UserRepository.Create(&userData); err != nil {
		return nil, err
	}

	userResponse := &response.UserResponse{
		ID:        userData.ID,
		Name:      userData.Name,
		Username:  userData.Username,
		Email:     userData.Email,
		CreatedAt: utils.ToTimestamp(userData.CreatedAt),
	}

	return userResponse, nil
}

func (s *UserService) Login(userLoginPayload request.UserLoginRequest) (*response.UserLoginResponse, error) {
	user, err := s.UserRepository.FindByUsername(userLoginPayload.Username)

	if err != nil {
		return nil, errors.New("user not found")
	}

	if !pkg.CheckPasswordHash(userLoginPayload.Password, user.Password) {
		return nil, errors.New("invalid password")
	}

	token, err := pkg.GenerateJWT(*user)

	if err != nil {
		return nil, err
	}

	return &response.UserLoginResponse{
		Token: token,
	}, nil
}
