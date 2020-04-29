package main

type AuthMessage interface {
    XXX_AuthPermissions() []string
    XXX_AuthResourceId() string
}
