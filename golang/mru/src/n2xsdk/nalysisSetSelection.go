package n2xsdk

type nalysisSetSelection struct {
 Handler string
}

func(np *nalysisSetSelection) SelectAnalysisSets () error {
 //parameters: NumSets psaAnalysisSetHandles
 //AgtAnalysisSetSelection SelectAnalysisSets, m.Object, m.Name);
 return nil
}

func(np *nalysisSetSelection) ListSelectedAnalysisSets ()(string, error) {
 //parameters: 
 //AgtAnalysisSetSelection ListSelectedAnalysisSets
 return "", nil
}

func(np *nalysisSetSelection) SetPercentiles () error {
 //parameters: NumPercentiles pPercentiles
 //AgtAnalysisSetSelection SetPercentiles, m.Object, m.Name);
 return nil
}

func(np *nalysisSetSelection) RequestDistribution () error {
 //parameters: NumBuckets
 //AgtAnalysisSetSelection RequestDistribution, m.Object, m.Name);
 return nil
}

func(np *nalysisSetSelection) CancelDistribution () error {
 //parameters: 
 //AgtAnalysisSetSelection CancelDistribution, m.Object, m.Name);
 return nil
}

func(np *nalysisSetSelection) WaitForDistribution () error {
 //parameters: 
 //AgtAnalysisSetSelection WaitForDistribution, m.Object, m.Name);
 return nil
}

func(np *nalysisSetSelection) WaitForExpansion () error {
 //parameters: 
 //AgtAnalysisSetSelection WaitForExpansion, m.Object, m.Name);
 return nil
}

func(np *nalysisSetSelection) WaitForScale () error {
 //parameters: 
 //AgtAnalysisSetSelection WaitForScale, m.Object, m.Name);
 return nil
}

func(np *nalysisSetSelection) GetScaleStatus ()(string, error) {
 //parameters: 
 //AgtAnalysisSetSelection GetScaleStatus
 return "", nil
}

func(np *nalysisSetSelection) GetExpansionStatus ()(string, error) {
 //parameters: 
 //AgtAnalysisSetSelection GetExpansionStatus
 return "", nil
}

func(np *nalysisSetSelection) IncreaseScale () error {
 //parameters: CenterTimeSeconds CenterTimeNanoseconds RelativeScaleFactor Width
 //AgtAnalysisSetSelection IncreaseScale, m.Object, m.Name);
 return nil
}

func(np *nalysisSetSelection) DecreaseScale () error {
 //parameters: NumLevels
 //AgtAnalysisSetSelection DecreaseScale, m.Object, m.Name);
 return nil
}

func(np *nalysisSetSelection) ExpandDistribution () error {
 //parameters: BucketNumber
 //AgtAnalysisSetSelection ExpandDistribution, m.Object, m.Name);
 return nil
}

func(np *nalysisSetSelection) ContractDistribution () error {
 //parameters: NumLevels
 //AgtAnalysisSetSelection ContractDistribution, m.Object, m.Name);
 return nil
}

