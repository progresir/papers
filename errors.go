package papers

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrSnils                        = Error("Incorrect SNILS")
	ErrSnilsCotainInvalidCharacters = Error("SNIPS must consist only of numbers")
	ErrINN                          = Error("Incorrect INN")
)
