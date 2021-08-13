package models

type CLI struct {
	Delimiter  string
	IgnoreCase bool
	Prepare    bool
	Help       bool
	Verbose    bool
	Version    bool
}
