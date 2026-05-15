package core

type ArtInstituteOfChicagoError struct {
	IsArtInstituteOfChicagoError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewArtInstituteOfChicagoError(code string, msg string, ctx *Context) *ArtInstituteOfChicagoError {
	return &ArtInstituteOfChicagoError{
		IsArtInstituteOfChicagoError: true,
		Sdk:              "ArtInstituteOfChicago",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *ArtInstituteOfChicagoError) Error() string {
	return e.Msg
}
