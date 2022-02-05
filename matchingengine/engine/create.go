package engine

func CreateEngine(asset string) *MatchingEngine {
	return &MatchingEngine{
		asset: asset,
	}
}
