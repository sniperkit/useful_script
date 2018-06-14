package n2xsdk

type MplsLabelManager struct {
 Handler string
}

func(np *MplsLabelManager) SetLabelRange () error {
 //parameters: hPort Minimum Maximum Increment
 //AgtMplsLabelManager SetLabelRange, m.Object, m.Name);
 return nil
}

func(np *MplsLabelManager) GetLabelRange ()(string, error) {
 //parameters: hPort
 //AgtMplsLabelManager GetLabelRange
 return "", nil
}

func(np *MplsLabelManager) SetSutImplicitNullLabel () error {
 //parameters: hPort Label
 //AgtMplsLabelManager SetSutImplicitNullLabel, m.Object, m.Name);
 return nil
}

func(np *MplsLabelManager) GetSutImplicitNullLabel ()(string, error) {
 //parameters: hPort
 //AgtMplsLabelManager GetSutImplicitNullLabel
 return "", nil
}

func(np *MplsLabelManager) EnableSutImplicitNullLabel () error {
 //parameters: hPort
 //AgtMplsLabelManager EnableSutImplicitNullLabel, m.Object, m.Name);
 return nil
}

func(np *MplsLabelManager) DisableSutImplicitNullLabel () error {
 //parameters: hPort
 //AgtMplsLabelManager DisableSutImplicitNullLabel, m.Object, m.Name);
 return nil
}

func(np *MplsLabelManager) IsSutImplicitNullLabelEnabled () error {
 //parameters: hPort
 //AgtMplsLabelManager IsSutImplicitNullLabelEnabled, m.Object, m.Name);
 return nil
}

