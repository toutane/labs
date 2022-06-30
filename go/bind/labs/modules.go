package labs

import (
	"berty.tech/labs/go/pkg/blmod"

	"berty.tech/labs/go/mod/datastorebench"
	"berty.tech/labs/go/mod/github_api"
)

var modules = []blmod.ModuleFactory{
	datastorebench.New,
	github_api.New,
}
