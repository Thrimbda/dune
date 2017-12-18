package linkutils

type NullCurrError struct{}

func (e NullCurrError) Error() string {
	return "curr is nil!"
}