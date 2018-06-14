package n2xsdk

type MosEstimation struct {
 Handler string
}

func(np *MosEstimation) IsMosEstimationSupported () error {
 //parameters: PortHandle
 //AgtMosEstimation IsMosEstimationSupported, m.Object, m.Name);
 return nil
}

func(np *MosEstimation) SetFactors () error {
 //parameters: Count psaStreamGroupHandles EquipmentImpairmentFactor PacketLossRobustnessFactor
 //AgtMosEstimation SetFactors, m.Object, m.Name);
 return nil
}

func(np *MosEstimation) GetFactors ()(string, error) {
 //parameters: Count psaStreamGroupHandles
 //AgtMosEstimation GetFactors
 return "", nil
}

func(np *MosEstimation) ClearFactors () error {
 //parameters: Count psaStreamGroupHandles
 //AgtMosEstimation ClearFactors, m.Object, m.Name);
 return nil
}

func(np *MosEstimation) IsFactorSet () error {
 //parameters: Count psaStreamGroupHandles
 //AgtMosEstimation IsFactorSet, m.Object, m.Name);
 return nil
}

