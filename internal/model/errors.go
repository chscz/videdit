package model

import (
	"fmt"
)

type VideoEditorError struct {
	Message string `json:"message"`
	Err     string `json:"error"`
}

func (ve *VideoEditorError) Error() string {
	if ve.Err != "" {
		return fmt.Sprintf("%s:%v", ve.Message, ve.Err)
	}
	return ve.Message
}

func NewVideoEditorError(index int, err error) *VideoEditorError {
	return &VideoEditorError{
		Message: fmt.Sprintf("%d번 째 영상 오류", index),
		Err:     err.Error(),
	}
}

func NewErrorToMap(err error) map[string]string {
	return map[string]string{"error": err.Error()}
}

func NewDetailErrorToMap(errMsg, errDetail error) map[string]string {
	return map[string]string{"message": errMsg.Error(), "error": errDetail.Error()}
}
