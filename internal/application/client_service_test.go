package application

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"testing"
	"wisdom/internal/domain/entity"
	"wisdom/mocks"
)

func TestNewApplicationServiceClient_Init(t *testing.T) {
	type fields struct {
		clientService mocks.ClientService
		challenge     mocks.ChallengeService
	}
	type args struct {
		clientID string
	}
	tests := []struct {
		name   string
		fields func() fields
		args   args
		want   error
	}{
		{
			name: "",
			fields: func() fields {
				clientService := mocks.ClientService{}
				clientService.On("SubscribeMessages", mock.Anything, mock.Anything).Return(nil)
				return fields{
					clientService: clientService,
				}
			},
			args: args{},
			want: nil,
		},
		{
			name: "",
			fields: func() fields {
				clientService := mocks.ClientService{}
				clientService.On("SubscribeMessages", mock.Anything, mock.Anything).Return(errors.New("error"))
				return fields{
					clientService: clientService,
				}
			},
			args: args{},
			want: errors.New("error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.fields()
			s := NewApplicationServiceClient(&f.clientService, nil, &f.challenge)
			got := s.Init()
			if got != nil && tt.want != nil && got.Error() != tt.want.Error() {
				t.Errorf("ApplicationServiceClient.Init() = %v, want %v", got, tt.want)
			}
			if (got == nil || tt.want == nil) && got != tt.want {
				t.Errorf("ApplicationServiceClient.Init() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewApplicationServiceClient_GetWisdom(t *testing.T) {
	type fields struct {
		clientService mocks.ClientService
		challenge     mocks.ChallengeService
	}
	type args struct {
		clientID string
	}
	tests := []struct {
		name   string
		fields func() fields
		args   args
		want   error
	}{
		{
			name: "",
			fields: func() fields {
				clientService := mocks.ClientService{}
				clientService.On("SendMessage", entity.ServerMessage{
					Content: []byte(`{"data":null,"command":"get_security_code","meta":{"security_token":""}}`),
				}).Return(nil)
				return fields{
					clientService: clientService,
				}
			},
			args: args{},
			want: errors.New("security token is invalid"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.fields()
			s := NewApplicationServiceClient(&f.clientService, nil, &f.challenge)
			got := s.GetWisdom()
			if got != nil && tt.want != nil && got.Error() != tt.want.Error() {
				t.Errorf("ApplicationServiceClient.GetWisdom() = %v, want %v", got, tt.want)
			}
			if (got == nil || tt.want == nil) && got != tt.want {
				t.Errorf("ApplicationServiceClient.GetWisdom() = %v, want %v", got, tt.want)
			}
		})
	}
}
