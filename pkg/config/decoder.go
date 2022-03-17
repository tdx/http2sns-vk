package config

import (
	"fmt"
	"strings"
)

type MapEndpointArn map[string]string

// Decode string as map items.
// Ex: "/endpoint:arn:aws:sns:eu-central-1:000000000000:sns-test" ->
//  map["/endpoint"]:"arn:aws:sns:eu-central-1:000000000000:sns-test"

func (mea *MapEndpointArn) Decode(value string) error {
	m := map[string]string{}
	pairs := strings.Split(value, ",")
	for _, pair := range pairs {
		kvpair := strings.SplitN(pair, ":", 2)
		if len(kvpair) != 2 {
			return fmt.Errorf("invalid map item: %q", pair)
		}
		// skip empty values
		if kvpair[0] == "" || kvpair[1] == "" {
			continue
		}
		m[kvpair[0]] = kvpair[1]
	}
	*mea = MapEndpointArn(m)
	return nil
}
