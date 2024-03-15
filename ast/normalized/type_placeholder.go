package normalized

import (
	"github.com/nar-lang/nar-common/ast"
	"github.com/nar-lang/nar-common/ast/typed"
)

type TPlaceholder struct {
	*typeBase
	name ast.FullIdentifier
}

func NewTPlaceholder(name ast.FullIdentifier) Type {
	return &TPlaceholder{
		typeBase: newTypeBase(ast.Location{}),
		name:     name,
	}
}

func (e *TPlaceholder) annotate(ctx *typed.SolvingContext, params typeParamsMap, source bool, placeholders placeholderMap) (typed.Type, error) {
	if p, ok := placeholders[e.name]; ok {
		return p, nil
	} else {
		placeholders[e.name] = nil
		return nil, nil
	}
}
