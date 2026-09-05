package api

import "github.com/danielgtaylor/huma/v2"

// NewGroup keeps operation IDs globally unique and maps each resource to an Orval module.
func NewGroup(api huma.API, path, resource string) *huma.Group {
	group := huma.NewGroup(api, path)
	group.UseSimpleModifier(func(op *huma.Operation) {
		op.OperationID += "-" + resource
		op.Tags = []string{resource}
	})
	return group
}
