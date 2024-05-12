package file

import (
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/lgp/http-server-example/internal/service/file/internal/mock_file_repository"

	sfssErrors "gitlab.com/lgp/http-server-example/internal/errors"
)

func TestFileServiceGetExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	assert := assert.New(t)

	mockReadRepository := mock_file_repository.NewMockfileReadRepository(ctrl)
	mockWriteRepository := mock_file_repository.NewMockfileWriteRepository(ctrl)

	mockReadRepository.EXPECT().Exists(gomock.Eq("file.txt")).Return(nil)
	mockReadRepository.EXPECT().Get(gomock.Eq("file.txt")).Return(tempData, nil)

	service := NewFileService(mockReadRepository, mockWriteRepository)
	data, err := service.Get("file.txt")

	assert.NoError(err)
	assert.Equal(tempData, data)
}

func TestFileServiceGetNotExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	assert := assert.New(t)

	mockReadRepository := mock_file_repository.NewMockfileReadRepository(ctrl)
	mockWriteRepository := mock_file_repository.NewMockfileWriteRepository(ctrl)

	mockReadRepository.EXPECT().Exists(gomock.Eq("file.txt")).Return(os.ErrNotExist)

	service := NewFileService(mockReadRepository, mockWriteRepository)
	_, err := service.Get("file.txt")

	assert.ErrorIs(err, sfssErrors.ErrFileNotExists)
}

func TestFileServiceUpload(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	assert := assert.New(t)

	mockReadRepository := mock_file_repository.NewMockfileReadRepository(ctrl)
	mockWriteRepository := mock_file_repository.NewMockfileWriteRepository(ctrl)

	mockReadRepository.EXPECT().Exists(gomock.Eq("file.txt")).Return(os.ErrNotExist)
	mockWriteRepository.EXPECT().Upload(gomock.Eq("file.txt"), gomock.Eq(tempData)).Return(nil)

	service := NewFileService(mockReadRepository, mockWriteRepository)
	err := service.Upload("file.txt", tempData)

	assert.NoError(err)
}

func TestFileServiceUploadAlreadyExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	assert := assert.New(t)

	mockReadRepository := mock_file_repository.NewMockfileReadRepository(ctrl)
	mockWriteRepository := mock_file_repository.NewMockfileWriteRepository(ctrl)

	mockReadRepository.EXPECT().Exists(gomock.Eq("file.txt")).Return(nil)

	service := NewFileService(mockReadRepository, mockWriteRepository)
	err := service.Upload("file.txt", tempData)

	assert.ErrorIs(err, sfssErrors.ErrFileAlreadyExists)
}

func TestFileServiceUploadNotPermitted(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	assert := assert.New(t)

	mockReadRepository := mock_file_repository.NewMockfileReadRepository(ctrl)
	mockWriteRepository := mock_file_repository.NewMockfileWriteRepository(ctrl)

	mockReadRepository.EXPECT().Exists(gomock.Eq("file.txt")).Return(os.ErrNotExist)
	mockWriteRepository.EXPECT().Upload(gomock.Eq("file.txt"), gomock.Eq(tempData)).Return(os.ErrPermission)

	service := NewFileService(mockReadRepository, mockWriteRepository)
	err := service.Upload("file.txt", tempData)

	assert.ErrorIs(err, sfssErrors.ErrInternalError)
}

func TestFileServiceDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	assert := assert.New(t)

	mockReadRepository := mock_file_repository.NewMockfileReadRepository(ctrl)
	mockWriteRepository := mock_file_repository.NewMockfileWriteRepository(ctrl)

	mockReadRepository.EXPECT().Exists(gomock.Eq("file.txt")).Return(nil)
	mockWriteRepository.EXPECT().Delete(gomock.Eq("file.txt")).Return(nil)

	service := NewFileService(mockReadRepository, mockWriteRepository)
	err := service.Delete("file.txt")

	assert.NoError(err)
}

func TestFileServiceDeleteNotExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	assert := assert.New(t)

	mockReadRepository := mock_file_repository.NewMockfileReadRepository(ctrl)
	mockWriteRepository := mock_file_repository.NewMockfileWriteRepository(ctrl)

	mockReadRepository.EXPECT().Exists(gomock.Eq("file.txt")).Return(os.ErrNotExist)

	service := NewFileService(mockReadRepository, mockWriteRepository)
	err := service.Delete("file.txt")

	assert.ErrorIs(err, sfssErrors.ErrFileNotExists)
}

func TestFileServiceDeleteNotPermitted(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	assert := assert.New(t)

	mockReadRepository := mock_file_repository.NewMockfileReadRepository(ctrl)
	mockWriteRepository := mock_file_repository.NewMockfileWriteRepository(ctrl)

	mockReadRepository.EXPECT().Exists(gomock.Eq("file.txt")).Return(nil)
	mockWriteRepository.EXPECT().Delete(gomock.Eq("file.txt")).Return(os.ErrPermission)

	service := NewFileService(mockReadRepository, mockWriteRepository)
	err := service.Delete("file.txt")

	assert.ErrorIs(err, sfssErrors.ErrInternalError)
}
