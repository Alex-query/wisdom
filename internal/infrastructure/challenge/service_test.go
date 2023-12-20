package challenge

import (
	"errors"
	"testing"
	"wisdom/mocks"
)

func TestService_GenerateTaskToResolve(t *testing.T) {
	type fields struct {
		Hasher mocks.Hasher
	}
	type args struct {
		clientID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				Hasher: mocks.Hasher{},
			},
			args: args{
				clientID: "----",
			},
			want:    "123",
			wantErr: false,
		},
		{
			name: "test2",
			fields: fields{
				Hasher: mocks.Hasher{},
			},
			args: args{
				clientID: "----",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		c := &Service{
			Hasher: &tt.fields.Hasher,
		}
		var errOut error
		if tt.wantErr {
			errOut = errors.New("error")
		}
		tt.fields.Hasher.On("GenerateRandomPrefixToken").Return(tt.want, errOut)
		got, err := c.GenerateTaskToResolve(tt.args.clientID)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Service.GenerateTaskToResolve() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			return
		}
		if got != tt.want {
			t.Errorf("%q. Service.GenerateTaskToResolve() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestService_VerifySolution(t *testing.T) {
	type fields struct {
		Hasher mocks.Hasher
	}
	type args struct {
		clientID string
		solution string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "test1",
			fields: fields{
				Hasher: mocks.Hasher{},
			},
			args: args{
				clientID: "----",
				solution: "123",
			},
			want: true,
		},
		{
			name: "test2",
			fields: fields{
				Hasher: mocks.Hasher{},
			},
			args: args{
				clientID: "----",
				solution: "123",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		c := &Service{
			Hasher: &tt.fields.Hasher,
		}

		tt.fields.Hasher.On("GenerateRandomPrefixToken").Return("123", nil)
		_, err := c.GenerateTaskToResolve(tt.args.clientID)
		if err != nil {
			t.Errorf("%q. Service.GenerateTaskToResolve() error = %v", tt.name, err)
			return
		}
		tt.fields.Hasher.On("Check", tt.args.solution, "123").Return(tt.want)
		got, err := c.VerifySolution(tt.args.clientID, tt.args.solution)
		if got != tt.want {
			t.Errorf("%q. Service.VerifySolution() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
