package stubs

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CreateRepository(fileName string, modPath string) string {
	// Split the string by underscores
	parts := strings.Split(fileName, "_")

	// Create a caser for title casing
	caser := cases.Title(language.English)

	// Capitalize the first letter of each part
	for i, part := range parts {
		parts[i] = caser.String(part)
	}

	// Join the parts back together
	fileNameUpper := strings.Join(parts, "")
	packageName := strings.ReplaceAll(fileName, "_", "")
	return fmt.Sprintf(`
package %vrepo

import (
	"context"
	"%v/app/models"
)

type %vRepository interface {
	Save(ctx context.Context, data *models.%v) error
	FindAll(ctx context.Context) ([]models.%v, error)
	FindByID(ctx context.Context, id uint) (*models.%v, error)
	Update(ctx context.Context, data *models.%v) error
	Delete(ctx context.Context, id uint) error
}
	`, packageName, modPath, fileNameUpper, fileNameUpper, fileNameUpper, fileNameUpper, fileNameUpper)
}

func CreateRepositoryImpl(fileName string, modPath string) string {
	// Split the string by underscores
	parts := strings.Split(fileName, "_")

	// Capitalize the first letter of each part
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}

	// Join the parts back together
	fileNameUpper := strings.Join(parts, "")
	packageName := strings.ReplaceAll(fileName, "_", "")
	alias := string(fileName[0])
	return fmt.Sprintf(`
package %vrepo

import (
	"context"
	"%v/app/facades"
	"%v/app/models"

	"gorm.io/gorm"
)

type %vRepositoryImpl struct {
	orm *gorm.DB
}

func New%vRepository() %vRepository {
	return &%vRepositoryImpl{orm: facades.Orm()}
}

// Delete implements %vRepository.
func (%v *%vRepositoryImpl) Delete(ctx context.Context, id uint) error {
		panic("implement me")
}

// FindAll implements %vRepository.
func (%v *%vRepositoryImpl) FindAll(ctx context.Context) ([]models.%v, error) {
		panic("implement me")
}

// FindByID implements %vRepository.
func (%v *%vRepositoryImpl) FindByID(ctx context.Context, id uint) (*models.%v, error) {
		panic("implement me")
}

// Save implements %vRepository.
func (%v *%vRepositoryImpl) Save(ctx context.Context, data *models.%v) error {
		panic("implement me")
}

// Update implements %vRepository.
func (%v *%vRepositoryImpl) Update(ctx context.Context, data *models.%v) error {
	panic("implement me")
}
`, packageName, modPath, modPath, fileNameUpper, fileNameUpper, fileNameUpper, fileNameUpper, fileNameUpper, alias, fileNameUpper, fileNameUpper, alias, fileNameUpper, fileNameUpper, fileNameUpper, alias, fileNameUpper, fileNameUpper, fileNameUpper, alias, fileNameUpper, fileNameUpper, fileNameUpper, alias, fileNameUpper, fileNameUpper)
}

func CreateService(fileName string, modPath string) string {
	// Split the string by underscores
	parts := strings.Split(fileName, "_")

	// Capitalize the first letter of each part
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}

	// Join the parts back together
	fileNameUpper := strings.Join(parts, "")
	packageName := strings.ReplaceAll(fileName, "_", "")
	return fmt.Sprintf(`
package %vservice

import (
	"context"
	"%v/app/dto"
	"%v/app/http/requests"
)

type %vService interface {
	Create(ctx context.Context, req requests.%vCreateRequest) error
	Delete(ctx context.Context, id uint) error
	FindAll(ctx context.Context) ([]dto.%vData, error)
	FindByID(ctx context.Context, id uint) (*dto.%vData, error)
	Show(ctx context.Context, id uint) (*dto.%vData, error)
	Update(ctx context.Context, req requests.%vUpdateRequest) error
}
`, packageName, modPath, modPath, fileNameUpper, fileNameUpper, fileNameUpper, fileNameUpper, fileNameUpper, fileNameUpper)
}

func CreateRequest(fileName string, operations []string) string {
	// Split the string by underscores
	parts := strings.Split(fileName, "_")

	// Capitalize the first letter of each part
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}

	// Join the parts back together
	fileNameUpper := strings.Join(parts, "")

	var result string
	result = `	
package requests

`
	for _, operation := range operations {
		operationUpper := strings.ToUpper(string(operation[0])) + operation[1:]
		result += fmt.Sprintf(`
type %v%vRequest struct {
}

	`, fileNameUpper, operationUpper)
	}

	return result
}

func CreateDto(fileName string) string {
	// Split the string by underscores
	parts := strings.Split(fileName, "_")

	// Capitalize the first letter of each part
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}

	// Join the parts back together
	fileNameUpper := strings.Join(parts, "")
	return fmt.Sprintf(`
package dto

type %vData struct {
}
	
`, fileNameUpper)
}

func CreateServiceImpl(fileName string, modPath string) string {
	// Split the string by underscores
	parts := strings.Split(fileName, "_")

	// Capitalize the first letter of each part
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}

	// Join the parts back together
	fileNameUpper := strings.Join(parts, "")
	fileNameRepo := strings.ToLower(string(fileNameUpper[0])) + fileNameUpper[1:]
	aliasService := string(fileName[0])
	aliasParam := aliasService + "r"
	packageName := strings.ReplaceAll(fileName, "_", "")
	return fmt.Sprintf(`
package %vservice

import (
	"context"
	"%v/app/dto"
	"%v/app/http/requests"
	%vrepo "%v/app/repositories/%v_repo"
)

type %vServiceImpl struct {
	%vRepository %vrepo.%vRepository
}

func New%vService(%v %vrepo.%vRepository) %vService {
	return &%vServiceImpl{%vRepository: %v}
}

// Create implements %vService.
func (%v *%vServiceImpl) Create(ctx context.Context, req requests.%vCreateRequest) error {
	panic("unimplemented")
}

// Delete implements %vService.
func (%v *%vServiceImpl) Delete(ctx context.Context, id uint) error {
	panic("unimplemented")
}

// FindAll implements %vService.
func (%v *%vServiceImpl) FindAll(ctx context.Context) ([]dto.%vData, error) {
	panic("unimplemented")
}

// FindByID implements %vService.
func (%v *%vServiceImpl) FindByID(ctx context.Context, id uint) (*dto.%vData, error) {
	panic("unimplemented")
}

// Show implements %vService.
func (%v *%vServiceImpl) Show(ctx context.Context, id uint) (*dto.%vData, error) {
	panic("unimplemented")
}

// Update implements %vService.
func (%v *%vServiceImpl) Update(ctx context.Context, req requests.%vUpdateRequest) error {
	panic("unimplemented")
}
`, packageName, modPath, modPath, packageName, modPath, fileName, fileNameUpper, fileNameRepo, packageName, fileNameUpper, fileNameUpper, aliasParam, packageName, fileNameUpper, fileNameUpper, fileNameUpper, fileNameRepo, aliasParam, fileNameUpper, aliasService, fileNameUpper, fileNameUpper, fileNameUpper, aliasService, fileNameUpper, fileNameUpper, aliasService, fileNameUpper, fileNameUpper, fileNameUpper, aliasService, fileNameUpper, fileNameUpper, fileNameUpper, aliasService, fileNameUpper, fileNameUpper, fileNameUpper, aliasService, fileNameUpper, fileNameUpper)
}
