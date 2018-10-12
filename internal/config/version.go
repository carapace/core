package config

type Version struct {
	Name             string
	Allowed          map[string]struct{}
	DeprecatedFields map[string]struct{}
	Deprecated       bool

	// map of config types to corresponding go structs
	objects map[string]interface{}
}

var V1 = Version{
	Name:             "v1",
	Allowed:          map[string]struct{}{},
	DeprecatedFields: map[string]struct{}{},
	Deprecated:       false,
}
