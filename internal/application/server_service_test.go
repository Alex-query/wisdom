package application

import (
	"errors"
	"testing"
	"wisdom/internal/domain/entity"
	"wisdom/mocks"
)

func TestApplicationServiceServer_GetSecurityCode(t *testing.T) {
	type fields struct {
		serverMock mocks.ServerService
		challenge  mocks.ChallengeService
		wisdom     mocks.WisdomRepository
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
				serverMock := mocks.ServerService{}
				serverMock.On("SendMessage", entity.ServerMessage{
					ClientID: "clientID111",
					Content:  []byte(`{"command":"get_security_code","meta":{"code":200,"task_to_resolve":"123","request_id":""}}`),
				}).Return(nil)
				challenge := mocks.ChallengeService{}
				challenge.On("GenerateTaskToResolve", "clientID111").Return("123", nil)
				wisdom := mocks.WisdomRepository{}
				return fields{
					serverMock: serverMock,
					challenge:  challenge,
					wisdom:     wisdom,
				}
			},
			args: args{
				clientID: "clientID111",
			},
			want: nil,
		},
		{
			name: "",
			fields: func() fields {
				serverMock := mocks.ServerService{}
				challenge := mocks.ChallengeService{}
				challenge.On("GenerateTaskToResolve", "clientID111").Return("", errors.New("error"))
				wisdom := mocks.WisdomRepository{}
				return fields{
					serverMock: serverMock,
					challenge:  challenge,
					wisdom:     wisdom,
				}
			},
			args: args{
				clientID: "clientID111",
			},
			want: errors.New("error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.fields()
			s := NewApplicationServiceServer(&f.serverMock, nil, &f.challenge, &f.wisdom)
			got := s.GetSecurityCode(tt.args.clientID, RequestMessageDTO{})
			if got != nil && tt.want != nil && got.Error() != tt.want.Error() {
				t.Errorf("ApplicationServiceServer.GetSecurityCode() = %v, want %v", got, tt.want)
			}
			if (got == nil || tt.want == nil) && got != tt.want {
				t.Errorf("ApplicationServiceServer.GetSecurityCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApplicationServiceServer_GetWisdom(t *testing.T) {
	type fields struct {
		serverMock mocks.ServerService
		challenge  mocks.ChallengeService
		wisdom     mocks.WisdomRepository
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
				serverMock := mocks.ServerService{}
				serverMock.On("SendMessage", entity.ServerMessage{
					ClientID: "clientID111",
					Content:  []byte(`{"command":"get_wisdom","meta":{"code":200,"task_to_resolve":"123","request_id":""},"data":{"wisdom":"wisdom"}}`),
				}).Return(nil)
				challenge := mocks.ChallengeService{}
				challenge.On("GenerateTaskToResolve", "clientID111").Return("123", nil)
				wisdom := mocks.WisdomRepository{}
				wisdom.On("GetRandomWisdom").Return("wisdom", nil)
				return fields{
					serverMock: serverMock,
					challenge:  challenge,
					wisdom:     wisdom,
				}
			},
			args: args{
				clientID: "clientID111",
			},
			want: nil,
		},
		{
			name: "",
			fields: func() fields {
				serverMock := mocks.ServerService{}
				challenge := mocks.ChallengeService{}
				challenge.On("GenerateTaskToResolve", "clientID111").Return("", errors.New("error"))
				wisdom := mocks.WisdomRepository{}
				wisdom.On("GetRandomWisdom").Return("wisdom", nil)
				return fields{
					serverMock: serverMock,
					challenge:  challenge,
					wisdom:     wisdom,
				}
			},
			args: args{
				clientID: "clientID111",
			},
			want: errors.New("error"),
		},
		{
			name: "",
			fields: func() fields {
				serverMock := mocks.ServerService{}
				challenge := mocks.ChallengeService{}
				challenge.On("GenerateTaskToResolve", "clientID111").Return("", nil)
				wisdom := mocks.WisdomRepository{}
				wisdom.On("GetRandomWisdom").Return("wisdom", errors.New("error_wisdom"))
				return fields{
					serverMock: serverMock,
					challenge:  challenge,
					wisdom:     wisdom,
				}
			},
			args: args{
				clientID: "clientID111",
			},
			want: errors.New("error_wisdom"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.fields()
			s := NewApplicationServiceServer(&f.serverMock, nil, &f.challenge, &f.wisdom)
			got := s.GetWisdom(tt.args.clientID, RequestMessageDTO{})
			if got != nil && tt.want != nil && got.Error() != tt.want.Error() {
				t.Errorf("ApplicationServiceServer.GetSecurityCode() = %v, want %v", got, tt.want)
			}
			if (got == nil || tt.want == nil) && got != tt.want {
				t.Errorf("ApplicationServiceServer.GetSecurityCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
