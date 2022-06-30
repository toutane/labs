package github_api

import (
	"context"

	"berty.tech/labs/go/pkg/blmod"
)

type module struct {
	blmod.UnimplementedModule
}

var _ blmod.Module = (*module)(nil)

func New() (blmod.Module, error) {
	return &module{}, nil
}

var _ blmod.ModuleFactory = New

func (m *module) Info() (*blmod.ModuleInfo, error) {
	return &blmod.ModuleInfo{
		Name:             "github-api",
		DisplayName:      "Github api",
		IconKind:         blmod.ModuleInfo_ICON_KIND_UTF,
		IconData:         ([]byte)("🐹"),
		ShortDescription: "Github api",
	}, nil
}

func (m *module) Run(ctx context.Context, args []byte, mc blmod.ModuleContext) error {
	return mc.Send("Hello module!")
}
