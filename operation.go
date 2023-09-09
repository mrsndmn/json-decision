package jsondecision

import (
	"errors"
	"fmt"
)

type OperationCodnition struct {
	Operation string
	Field     string
	Value     string
}

type Operation string

const OperationEqual Operation = "=="
const OperationGreater Operation = ">"
const OperationGreaterOrEqual Operation = ">="
const OperationLess Operation = "<"
const OperationLessOrEqual Operation = "<="

var ErrOperationConditionInvalidOperation = errors.New("invalid operation")

func (cond OperationCodnition) Check(decisionParams DecisionParams) (bool, error) {

	switch Operation(cond.Operation) {
	case OperationEqual:
		return cond.CheckEqual(decisionParams)
	case OperationGreater:
		return cond.CheckGreater(decisionParams)
	case OperationGreaterOrEqual:
		return cond.CheckGreaterOrEqual(decisionParams)
	case OperationLess:
		return cond.CheckLess(decisionParams)
	case OperationLessOrEqual:
		return cond.CheckLessOrEqual(decisionParams)
	}

	return false, fmt.Errorf("%w: %s", ErrOperationConditionInvalidOperation, cond.Operation)
}

func (cond OperationCodnition) CheckEqual(decisionParams DecisionParams) (bool, error) {
	return decisionParams[cond.Field] == cond.Value, nil
}

func (cond OperationCodnition) CheckGreater(decisionParams DecisionParams) (bool, error) {
	return decisionParams[cond.Field] > cond.Value, nil
}

func (cond OperationCodnition) CheckGreaterOrEqual(decisionParams DecisionParams) (bool, error) {
	return decisionParams[cond.Field] >= cond.Value, nil
}

func (cond OperationCodnition) CheckLess(decisionParams DecisionParams) (bool, error) {
	return decisionParams[cond.Field] < cond.Value, nil
}

func (cond OperationCodnition) CheckLessOrEqual(decisionParams DecisionParams) (bool, error) {
	return decisionParams[cond.Field] <= cond.Value, nil
}
