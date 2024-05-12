package handler

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"gitlab.com/lgp/http-server-example/internal/handler/http/internal/mock_file_service"

	sfssError "gitlab.com/lgp/http-server-example/internal/errors"
)

func TestHandlerGet(t *testing.T) {
	tempData := []byte("some temp data")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	assert := assert.New(t)

	mockService := mock_file_service.NewMockfileService(ctrl)
	mockService.EXPECT().Get(gomock.Eq("test.txt")).Return(tempData, nil)

	h := NewHandler(mux.NewRouter(), mockService)
	server := httptest.NewServer(h)
	defer server.Close()

	resp, err := http.Get(fmt.Sprintf("%s/%s", server.URL, path.Join("get", "test.txt")))

	assert.NoError(err)
	assert.Equal(resp.StatusCode, http.StatusOK)

	data, _ := io.ReadAll(resp.Body)
	assert.Equal(tempData, data)
}

func TestHandlerGetNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	assert := assert.New(t)

	mockService := mock_file_service.NewMockfileService(ctrl)
	mockService.EXPECT().Get(gomock.Eq("test.txt")).Return(
		nil,
		fmt.Errorf("%w: %s: %w", sfssError.ErrFileNotExists, "test.txt", os.ErrNotExist),
	)

	h := NewHandler(mux.NewRouter(), mockService)
	server := httptest.NewServer(h)
	defer server.Close()

	resp, err := http.Get(fmt.Sprintf("%s/%s", server.URL, path.Join("get", "test.txt")))

	assert.NoError(err)
	assert.Equal(http.StatusNotFound, resp.StatusCode)
}

func createFormFile(name string) (io.Reader, string) {
	buf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(buf)
	defer bodyWriter.Close()
	fileWriter, _ := bodyWriter.CreateFormFile("file", name)
	file, _ := os.Open(name)
	defer file.Close()

	io.Copy(fileWriter, file)

	return buf, bodyWriter.FormDataContentType()
}

func TestHandlerUpload(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	assert := assert.New(t)

	tempFile, _ := os.Open("testdata/upload-file.txt")
	tempData, _ := io.ReadAll(tempFile)

	mockService := mock_file_service.NewMockfileService(ctrl)
	mockService.EXPECT().Upload(gomock.Eq("upload-file.txt"), gomock.Eq(tempData)).Return(nil)

	h := NewHandler(mux.NewRouter(), mockService)
	server := httptest.NewServer(h)
	defer server.Close()

	body, content := createFormFile("testdata/upload-file.txt")
	resp, err := http.Post(fmt.Sprintf("%s/%s", server.URL, "upload"), content, body)

	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
}

func TestHandlerAlreadyExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	assert := assert.New(t)

	tempFile, _ := os.Open("testdata/upload-file.txt")
	tempData, _ := io.ReadAll(tempFile)

	mockService := mock_file_service.NewMockfileService(ctrl)
	mockService.EXPECT().Upload(gomock.Eq("upload-file.txt"), gomock.Eq(tempData)).Return(
		fmt.Errorf("%w: %s: %w", sfssError.ErrFileAlreadyExists, "upload-file.txt", os.ErrExist),
	)

	h := NewHandler(mux.NewRouter(), mockService)
	server := httptest.NewServer(h)
	defer server.Close()

	body, content := createFormFile("testdata/upload-file.txt")
	resp, err := http.Post(fmt.Sprintf("%s/%s", server.URL, "upload"), content, body)

	assert.NoError(err)
	assert.Equal(http.StatusConflict, resp.StatusCode)
}

func TestHandlerUploadBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	assert := assert.New(t)

	mockService := mock_file_service.NewMockfileService(ctrl)

	h := NewHandler(mux.NewRouter(), mockService)
	server := httptest.NewServer(h)
	defer server.Close()

	_, content := createFormFile("testdata/upload-file.txt")
	resp, err := http.Post(fmt.Sprintf("%s/%s", server.URL, "upload"), content, nil)

	assert.NoError(err)
	assert.Equal(http.StatusBadRequest, resp.StatusCode)
}

func TestHandlerDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	assert := assert.New(t)

	mockService := mock_file_service.NewMockfileService(ctrl)
	mockService.EXPECT().Delete(gomock.Eq("test.txt")).Return(nil)

	h := NewHandler(mux.NewRouter(), mockService)
	server := httptest.NewServer(h)
	defer server.Close()

	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", server.URL, path.Join("delete?name=test.txt")), nil)
	resp, err := http.DefaultClient.Do(req)

	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
}

func TestHandlerDeleteNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	assert := assert.New(t)

	mockService := mock_file_service.NewMockfileService(ctrl)
	mockService.EXPECT().Delete(gomock.Eq("test.txt")).Return(
		fmt.Errorf("%w: %s: %w", sfssError.ErrFileNotExists, "test.txt", os.ErrNotExist),
	)

	h := NewHandler(mux.NewRouter(), mockService)
	server := httptest.NewServer(h)
	defer server.Close()

	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", server.URL, path.Join("delete?name=test.txt")), nil)
	resp, err := http.DefaultClient.Do(req)

	assert.NoError(err)
	assert.Equal(http.StatusNotFound, resp.StatusCode)
}
