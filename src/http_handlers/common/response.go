package common

type StandardResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (s *StandardResponse) SetError(err error, status int) {
	if err != nil {
		s.Message = err.Error()
	}
	s.Status = status
}
