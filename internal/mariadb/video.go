package mariadb

import (
	"context"

	"github.com/chscz/videdit/internal/model"
)

func (r Repository) CreateVideoUpload(ctx context.Context, file *model.VideoUpload) error {
	return r.DB.WithContext(ctx).Create(&file).Error
}

func (r Repository) CreateVideoRequest(ctx context.Context, req *model.VideoCreate) error {
	return r.DB.WithContext(ctx).Create(&req).Error
}

func (r Repository) GetUploadVideoList(ctx context.Context) ([]*model.VideoUpload, error) {
	var uploadVideos []*model.VideoUpload
	if err := r.DB.WithContext(ctx).Order("created_at DESC").Find(&uploadVideos).Error; err != nil {
		return nil, err
	}
	return uploadVideos, nil
}

func (r Repository) GetCreateVideoList(ctx context.Context) ([]*model.VideoCreate, error) {
	var createVideos []*model.VideoCreate
	if err := r.DB.WithContext(ctx).Order("created_at DESC").Find(&createVideos).Error; err != nil {
		return nil, err
	}
	return createVideos, nil
}
