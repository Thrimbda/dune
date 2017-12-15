package arrayutils

type FullListError struct{}

func (e FullListError) Error() string {
	return "list is full!"
}

type BadCurrError struct{}

func (e BadCurrError) Error() string {
	return "bad current pointer position."
}

type EmptyListError struct{}

func (e EmptyListError) Error() string {
	return "list is empty!"
}