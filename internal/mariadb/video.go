package mariadb

import (
	"context"

	"github.com/chscz/videdit/internal/model"
)

func (r Repository) CreateVideoUpload(ctx context.Context, file *model.VideoUpload) error {
	return r.DB.WithContext(ctx).Create(&file).Error
}

func (r Repository) CreateVideoRequest(ctx context.Context, req *model.VideoRequest) error {
	return r.DB.WithContext(ctx).Create(&req).Error
}
