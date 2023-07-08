package application_test

import (
	gomock "github.com/golang/mock/gomock"
	application "github.com/phelipperibeiro/golang-hexagonal-architecture/application"
	mock_application "github.com/phelipperibeiro/golang-hexagonal-architecture/application/mocks"
	require "github.com/stretchr/testify/require"
	testing "testing"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mock_application.NewMockProductInterface(ctrl)
	repository := mock_application.NewMockProductRepositoryInterface(ctrl)
	repository.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()
	service := application.ProductService{Repository: repository}

	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)

}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mock_application.NewMockProductInterface(ctrl)
	repository := mock_application.NewMockProductRepositoryInterface(ctrl)
	repository.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
	service := application.ProductService{Repository: repository}

	result, err := service.Create("Product 1", 10)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_EnableDisable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mock_application.NewMockProductInterface(ctrl)
	product.EXPECT().Enable().Return(nil)
	product.EXPECT().Disable().Return(nil)

	repository := mock_application.NewMockProductRepositoryInterface(ctrl)
	repository.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
	service := application.ProductService{Repository: repository}

	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

	result, err = service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}
