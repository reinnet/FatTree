package builder

import "github.com/reinnet/topology/model"

// Builder type can build a toplogy.
type Builder interface {
	Build() model.Config
}
