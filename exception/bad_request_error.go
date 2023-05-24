package exception

type BadRequest struct{
	Error string
}

func NewBadRequestError(err string) BadRequest {
	return BadRequest{Error: err}
}