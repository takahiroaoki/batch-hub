package database

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/takahiroaoki/batch-hub/config"
)

func TestDataSourceName(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		setup func(config *config.MockDatabaseConfig)
		want  string
	}{
		{
			name: "Success",
			setup: func(config *config.MockDatabaseConfig) {
				config.EXPECT().Host().Return("host")
				config.EXPECT().Port().Return("port")
				config.EXPECT().Database().Return("database")
				config.EXPECT().User().Return("user")
				config.EXPECT().Password().Return("password")
			},
			want: "user:password@tcp(host:port)/database?charset=utf8mb4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockConfig := config.NewMockDatabaseConfig(ctrl)
			tt.setup(mockConfig)

			assert.Equal(t, tt.want, DataSourceName(mockConfig))
		})
	}
}

func TestDatabaseURL(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		setup func(config *config.MockDatabaseConfig)
		want  string
	}{
		{
			name: "Success",
			setup: func(config *config.MockDatabaseConfig) {
				config.EXPECT().Host().Return("host")
				config.EXPECT().Port().Return("port")
				config.EXPECT().Database().Return("database")
				config.EXPECT().User().Return("user")
				config.EXPECT().Password().Return("password")
			},
			want: "mysql://user:password@tcp(host:port)/database?charset=utf8mb4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockConfig := config.NewMockDatabaseConfig(ctrl)
			tt.setup(mockConfig)

			assert.Equal(t, tt.want, DatabaseURL(mockConfig))
		})
	}
}
