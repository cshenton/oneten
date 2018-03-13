package oneten

import "errors"

// Cell is a representation of a rule 110 state with fixed boundaries left
// and right.
type Cell struct {
	Values []bool
	Left   bool
	Right  bool
}

// New creates a new cell with the provided dimension and boundaries.
func New(values []bool, left, right bool) (c *Cell, err error) {
	if len(values) < 1 {
		err = errors.New("values must have length 1 or more")
		return nil, err
	}

	c = &Cell{
		Values: values,
		Left:   left,
		Right:  right,
	}
	return c, nil
}

// Rule is the rule 110 transition rule. Ugly but correct.
func Rule(l, m, r bool) bool {
	if l {
		if m {
			if r {
				return false
			}
			return true
		}
		if r {
			return true
		}
		return false
	}
	if m {
		if r {
			return true
		}
		return true
	}
	if r {
		return true
	}
	return false

}

// Next transitions the cell to its next state, according to the transition rules.
func (c *Cell) Next() {
	var l bool
	var r bool
	next := make([]bool, len(c.Values))

	for i := range c.Values {
		if i == 0 {
			l = c.Left
		} else {
			l = c.Values[i-1]
		}
		if i == len(c.Values)-1 {
			r = c.Right
		} else {
			r = c.Values[i]
		}

		next[i] = false
	}
	c.Values = next
}
