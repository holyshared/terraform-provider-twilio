package validation

import "testing"

func TestListOfMatchStringNoWarningAndNoError(t *testing.T) {
	fn := ListOfMatchString([]string{"A", "B"})
	warns, errs := fn([]interface{}{"A"}, "key")

	if len(warns) > 0 {
		t.Errorf("There should be no warning.")
	}
	if len(errs) > 0 {
		t.Errorf("There should be no errors. %v", errs)
	}
}

func TestListOfMatchStringInvalidElement(t *testing.T) {
	fn := ListOfMatchString([]string{"A", "B"})
	warns, errs := fn([]interface{}{"A", "C"}, "key")

	if len(warns) > 0 {
		t.Errorf("There should be no warning.")
	}
	if len(errs) != 1 {
		t.Errorf("There should be one error. %v", errs)
	}
}
