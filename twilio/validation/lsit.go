package list

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ListOfMatchString(valid []string) schema.SchemaValidateFunc {
	return func(val interface{}, key string) (warns []string, errs []error) {
		v, ok := val.([]interface{})
		if !ok { // type error
			errs = append(errs, fmt.Errorf("expected type of %q to be List", key))
			return warns, errs
		}

		for _, e := range v {
			if _, eok := e.(string); !eok {
				errs = append(errs, fmt.Errorf("expected %q to only contain string elements, found :%v", key, e))
				return warns, errs
			}
		}

		for _, sv := range v {
			find := false
			for _, tv := range valid {
				if sv.(string) == tv {
					find = true
					break
				}
				if !find {
					errs = append(errs, fmt.Errorf("expected %q to event names, found %v", key, sv))
					return warns, errs
				}
			}
		}

		return warns, errs
	}
}
