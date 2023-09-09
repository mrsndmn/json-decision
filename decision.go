package jsondecision

type Condition struct {
	BooleanCodnition
	OperationCodnition
}

type DecisionParams map[string]string

func (cond Condition) Check(decisionParams DecisionParams) (bool, error) {
	if len(cond.And) > 0 || len(cond.Or) > 0 {
		booleanCondition := BooleanCodnition{
			And: cond.And,
			Or:  cond.Or,
		}
		return booleanCondition.Check(decisionParams)
	}

	operationCodnition := OperationCodnition{
		Operation: cond.Operation,
		Field:     cond.Field,
		Value:     cond.Value,
	}

	return operationCodnition.Check(decisionParams)
}
