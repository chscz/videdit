package util

func NewErrorToMap(err error) map[string]string {
	return map[string]string{"error": err.Error()}
}

func NewDetailErrorToMap(errMsg, errDetail error) map[string]string {
	return map[string]string{"message": errMsg.Error(), "error": errDetail.Error()}
}
