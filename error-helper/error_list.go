package error_helper

import (
	"fmt"
	"sync"
)

type List struct {
	*sync.Mutex
	Errors []error
}

func NewErrorList() *List {
	return &List{Mutex: &sync.Mutex{}}
}

func (errorList *List) Error() string {
	errorList.Lock()
	defer errorList.Unlock()

	if len(errorList.Errors) == 0 {
		return ""
	}

	return fmt.Sprintf("multiple errors (total %d): %+v", len(errorList.Errors), errorList.Errors)
}

func (errorList *List) HasErrors() bool {
	errorList.Lock()
	defer errorList.Unlock()

	return len(errorList.Errors) > 0
}

func (errorList *List) ErrOrNil() error {
	if errorList.HasErrors() {
		return errorList
	}

	return nil
}

func (errorList *List) Add(errs ...error) {
	errorList.Lock()
	defer errorList.Unlock()

	for _, err := range errs {
		if err != nil {
			errorList.Errors = append(errorList.Errors, err)
		}
	}
}
