package genx

import (
	"fmt"

	"github.com/rogeecn/genx/field"
	"gorm.io/gorm/clause"
)

// Cond convert expression array to condition array
func Cond(exprs ...clause.Expression) []Condition {
	return exprToCondition(exprs...)
}

var _ Condition = &CondContainer{}

type CondContainer struct {
	value interface{}
	err   error
}

func (c *CondContainer) BeCond() interface{} { return c.value }
func (c *CondContainer) CondError() error    { return c.err }

func exprToCondition(exprs ...clause.Expression) []Condition {
	conds := make([]Condition, 0, len(exprs))
	for _, e := range exprs {
		conds = append(conds, &CondContainer{value: e})
	}
	return conds
}

func condToExpression(conds []Condition) ([]clause.Expression, error) {
	if len(conds) == 0 {
		return nil, nil
	}
	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		if cond == nil {
			continue
		}
		if err := cond.CondError(); err != nil {
			return nil, err
		}

		switch cond.(type) {
		case *CondContainer, field.Expr, SubQuery:
		default:
			return nil, fmt.Errorf("unsupported condition: %+v", cond)
		}

		switch e := cond.BeCond().(type) {
		case []clause.Expression:
			exprs = append(exprs, e...)
		case clause.Expression:
			exprs = append(exprs, e)
		}
	}
	return exprs, nil
}
