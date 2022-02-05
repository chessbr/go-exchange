package entities

type MatchingResultTypeEnum int

const (
	Queued   MatchingResultTypeEnum = 1
	Executed MatchingResultTypeEnum = 2
	Deleted  MatchingResultTypeEnum = 3
)

type MatchingResult struct {
	ResultType MatchingResultTypeEnum
	Quantity   int
	Price      int
}

type ExecutionResult struct {
	Quantity int
	Price    int
}
