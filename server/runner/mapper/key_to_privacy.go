package mapper

import (
	"fmt"
	awesome2v1 "github.com/abitofhelp/awesome2/gen/go/awesome2/v1"
)

func KeyToPrivacy(key string) (awesome2v1.Privacy, error) {

	var value awesome2v1.Privacy
	if v, found := awesome2v1.Privacy_value[key]; found {
		value = awesome2v1.Privacy(v)
	} else {
		return awesome2v1.Privacy_PRIVACY_UNSPECIFIED, fmt.Errorf("\nfailed to find an awesome2v1.Privacy named '%s'", key)
	}

	return value, nil
}
