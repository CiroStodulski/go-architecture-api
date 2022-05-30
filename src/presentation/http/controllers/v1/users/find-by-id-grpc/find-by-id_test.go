package v1_user_grpc

import (
	"go-api/src/core/entities/user"
	"go-api/src/main/container"
	ports_http "go-api/src/presentation/http/ports"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserCase struct {
	mock.Mock
}

func (mock *MockUserCase) GetUser(id string) (*user.User, error) {
	arg := mock.Called()
	result := arg.Get(0)
	return result.(*user.User), arg.Error(1)
}

func newMockUser() *user.User {
	user, _ := user.NewUser("test", "test", "test")
	return user
}
func Test_Controller_GetUser_Grpc(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		userMock := newMockUser()
		mockUseCase := new(MockUserCase)
		id := "752ea551-5e6a-4382-859c-cd09fbe50110"

		mockUseCase.On("GetUser").Return(userMock, nil)

		testService := NewController(&container.Container{
			GetUserGrpcUseCase: mockUseCase,
		})

		result, err := testService.LoadRoute().Handle(ports_http.HttpRequest{
			Params: ports_http.Params{
				ports_http.Param{Key: "id", Value: id},
			},
		})

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, &ports_http.HttpResponse{
			Data:   userMock,
			Status: 200,
		}, result)
	})
}
