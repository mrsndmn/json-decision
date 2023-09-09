package jsondecision

import "errors"

type BooleanCodnition struct {
	Or  []Condition
	And []Condition
}

var ErrBooleanConditionOrAnd = errors.New("boolean condtion can have either `or` or `and` field, not both the same time")
var ErrBooleanConditionEmpty = errors.New("boolean condtion must have either `or` or `and` field")

func (cond BooleanCodnition) Check(decisionParams DecisionParams) (bool, error) {
	if len(cond.And) > 0 && len(cond.Or) > 0 {
		return false, ErrBooleanConditionOrAnd
	}

	if len(cond.And) == 0 && len(cond.Or) == 0 {
		return false, ErrBooleanConditionEmpty
	}

	if len(cond.And) > 0 {
		return cond.CheckAnd(decisionParams)
	}

	return cond.CheckOr(decisionParams)
}

func (cond BooleanCodnition) CheckAnd(decisionParams DecisionParams) (bool, error) {
	for _, cond := range cond.And {
		decision, err := cond.Check(decisionParams)
		if err != nil {
			return false, err
		}

		if !decision {
			return false, nil
		}
	}

	return true, nil
}

func (cond BooleanCodnition) CheckOr(decisionParams DecisionParams) (bool, error) {
	for _, cond := range cond.Or {
		decision, err := cond.Check(decisionParams)
		if err != nil {
			return false, err
		}

		if decision {
			return true, nil
		}
	}

	return false, nil
}
