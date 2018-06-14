package n2xsdk

type Bgp4Ipv4RouteProfile struct {
 Handler string
}

func(np *Bgp4Ipv4RouteProfile) EnableAutoPathAttribute () error {
 //parameters: RouteProfileHandle PathAttributeAuto
 //AgtBgp4Ipv4RouteProfile EnableAutoPathAttribute, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) DisableAutoPathAttribute () error {
 //parameters: RouteProfileHandle PathAttributeAuto
 //AgtBgp4Ipv4RouteProfile DisableAutoPathAttribute, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) IsAutoPathAttributeEnabled () error {
 //parameters: RouteProfileHandle PathAttributeAuto
 //AgtBgp4Ipv4RouteProfile IsAutoPathAttributeEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) SetNextHopIncrementingRange () error {
 //parameters: RouteProfileHandle FirstNextHop PrefixLength NextHopIncrement NextHopRepeat
 //AgtBgp4Ipv4RouteProfile SetNextHopIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) GetNextHopIncrementingRange ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetNextHopIncrementingRange
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) SetNextHopList ()(string, error) {
 //parameters: RouteProfileHandle NextHopList
 //AgtBgp4Ipv4RouteProfile SetNextHopList
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) GetNextHopList ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetNextHopList
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) SetNextHop () error {
 //parameters: RouteProfileHandle NextHop
 //AgtBgp4Ipv4RouteProfile SetNextHop, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) GetNextHop ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetNextHop
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) GetNextHopAuto ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetNextHopAuto
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) SetRoutesIncrementingRange () error {
 //parameters: RouteProfileHandle FirstRoute PrefixLengthIpv4 RouteIncrement PercentOverlap
 //AgtBgp4Ipv4RouteProfile SetRoutesIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) GetRoutesIncrementingRange ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetRoutesIncrementingRange
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) SetRoutesPrefixLength () error {
 //parameters: RouteProfileHandle PrefixLengthIpv4
 //AgtBgp4Ipv4RouteProfile SetRoutesPrefixLength, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) GetRoutesPrefixLength ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetRoutesPrefixLength
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) SetRoutesSubRangeOffsets () error {
 //parameters: RouteProfileHandle FirstRoute MsbOffsetList
 //AgtBgp4Ipv4RouteProfile SetRoutesSubRangeOffsets, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) GetRoutesSubRangeOffsets ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetRoutesSubRangeOffsets
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) SetRoutesSubRange () error {
 //parameters: RouteProfileHandle SubRangeInstance RouteIncrement Count Repeat
 //AgtBgp4Ipv4RouteProfile SetRoutesSubRange, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) GetRoutesSubRange ()(string, error) {
 //parameters: RouteProfileHandle SubRangeInstance
 //AgtBgp4Ipv4RouteProfile GetRoutesSubRange
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) SetRoutesPerPeer () error {
 //parameters: RouteProfileHandle RoutesPerPeer
 //AgtBgp4Ipv4RouteProfile SetRoutesPerPeer, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) GetRoutesPerPeer ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetRoutesPerPeer
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) GetTotalRoutesInProfile ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetTotalRoutesInProfile
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) GetNthRouteInProfile ()(string, error) {
 //parameters: RouteProfileHandle RouteProfileIndex
 //AgtBgp4Ipv4RouteProfile GetNthRouteInProfile
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) GetNumAdvertisedRoutes ()(string, error) {
 //parameters: RouteProfileInstance
 //AgtBgp4Ipv4RouteProfile GetNumAdvertisedRoutes
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) GetNumWithdrawnRoutes ()(string, error) {
 //parameters: RouteProfileInstance
 //AgtBgp4Ipv4RouteProfile GetNumWithdrawnRoutes
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) EnableLabeling () error {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile EnableLabeling, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) DisableLabeling () error {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile DisableLabeling, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) IsLabelingEnabled () error {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile IsLabelingEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) SetLabelingMode () error {
 //parameters: RouteProfileHandle LabelingMode
 //AgtBgp4Ipv4RouteProfile SetLabelingMode, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) GetLabelingMode ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetLabelingMode
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) SetUserLabelType () error {
 //parameters: RouteProfileHandle UserLabelType
 //AgtBgp4Ipv4RouteProfile SetUserLabelType, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) GetUserLabelType ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetUserLabelType
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) SetUserLabelIncrementingRange () error {
 //parameters: RouteProfileHandle StartLabel LabelIncrement LabelRepeat
 //AgtBgp4Ipv4RouteProfile SetUserLabelIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) GetUserLabelIncrementingRange ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetUserLabelIncrementingRange
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) SetUserLabelList ()(string, error) {
 //parameters: RouteProfileHandle LabelList
 //AgtBgp4Ipv4RouteProfile SetUserLabelList
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) GetUserLabelList ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetUserLabelList
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) SetUserLabel () error {
 //parameters: RouteProfileHandle Label
 //AgtBgp4Ipv4RouteProfile SetUserLabel, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) GetUserLabel ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetUserLabel
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) GetLabels ()(string, error) {
 //parameters: RouteProfileInstance
 //AgtBgp4Ipv4RouteProfile GetLabels
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) GetLabelPoolState ()(string, error) {
 //parameters: RouteProfileInstance
 //AgtBgp4Ipv4RouteProfile GetLabelPoolState
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) GetPeerPoolHandle ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetPeerPoolHandle
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) GetState ()(string, error) {
 //parameters: RouteProfileInstance
 //AgtBgp4Ipv4RouteProfile GetState
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) IsEnabled () error {
 //parameters: RouteProfileInstance
 //AgtBgp4Ipv4RouteProfile IsEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) EnableAsPath () error {
 //parameters: RouteProfileHandle AsPathAttributeType
 //AgtBgp4Ipv4RouteProfile EnableAsPath, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) Enable4ByteAsPath () error {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType
 //AgtBgp4Ipv4RouteProfile Enable4ByteAsPath, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) DisableAsPath () error {
 //parameters: RouteProfileHandle AsPathAttributeType
 //AgtBgp4Ipv4RouteProfile DisableAsPath, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) Disable4ByteAsPath () error {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType
 //AgtBgp4Ipv4RouteProfile Disable4ByteAsPath, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) IsAsPathEnabled () error {
 //parameters: RouteProfileHandle AsPathAttributeType
 //AgtBgp4Ipv4RouteProfile IsAsPathEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) Is4ByteAsPathEnabled () error {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType
 //AgtBgp4Ipv4RouteProfile Is4ByteAsPathEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) EnableAutoAsPath () error {
 //parameters: RouteProfileHandle AsPathAttributeType
 //AgtBgp4Ipv4RouteProfile EnableAutoAsPath, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) EnableAuto4ByteAsPath () error {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType
 //AgtBgp4Ipv4RouteProfile EnableAuto4ByteAsPath, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) DisableAutoAsPath () error {
 //parameters: RouteProfileHandle AsPathAttributeType
 //AgtBgp4Ipv4RouteProfile DisableAutoAsPath, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) DisableAuto4ByteAsPath () error {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType
 //AgtBgp4Ipv4RouteProfile DisableAuto4ByteAsPath, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) IsAutoAsPathEnabled () error {
 //parameters: RouteProfileHandle AsPathAttributeType
 //AgtBgp4Ipv4RouteProfile IsAutoAsPathEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) IsAuto4ByteAsPathEnabled () error {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType
 //AgtBgp4Ipv4RouteProfile IsAuto4ByteAsPathEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) SetAsPathIncrementingRange () error {
 //parameters: RouteProfileHandle AsPathAttributeType FirstAsNumber AsNumberIncrement AsNumberRepeat
 //AgtBgp4Ipv4RouteProfile SetAsPathIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) Set4ByteAsPathIncrementingRange () error {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType First4ByteAsNumber AsNumberIncrement AsNumberRepeat
 //AgtBgp4Ipv4RouteProfile Set4ByteAsPathIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) GetAsPathIncrementingRange ()(string, error) {
 //parameters: RouteProfileHandle AsPathAttributeType
 //AgtBgp4Ipv4RouteProfile GetAsPathIncrementingRange
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) Get4ByteAsPathIncrementingRange ()(string, error) {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType
 //AgtBgp4Ipv4RouteProfile Get4ByteAsPathIncrementingRange
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) SetAsPathList ()(string, error) {
 //parameters: RouteProfileHandle AsPathAttributeType AsNumberList
 //AgtBgp4Ipv4RouteProfile SetAsPathList
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) Set4ByteAsPathList ()(string, error) {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType AsNumberList
 //AgtBgp4Ipv4RouteProfile Set4ByteAsPathList
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) GetAsPathList ()(string, error) {
 //parameters: RouteProfileHandle AsPathAttributeType
 //AgtBgp4Ipv4RouteProfile GetAsPathList
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) Get4ByteAsPathList ()(string, error) {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType
 //AgtBgp4Ipv4RouteProfile Get4ByteAsPathList
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) SetAsPath () error {
 //parameters: RouteProfileHandle AsPathAttributeType AsNumber
 //AgtBgp4Ipv4RouteProfile SetAsPath, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) Set4ByteAsPath () error {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType AsNumber
 //AgtBgp4Ipv4RouteProfile Set4ByteAsPath, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) GetAsPath ()(string, error) {
 //parameters: RouteProfileHandle AsPathAttributeType
 //AgtBgp4Ipv4RouteProfile GetAsPath
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) Get4ByteAsPath ()(string, error) {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType
 //AgtBgp4Ipv4RouteProfile Get4ByteAsPath
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) AddAsPathToListOfLists ()(string, error) {
 //parameters: RouteProfileHandle AsPathAttributeType AsNumberList
 //AgtBgp4Ipv4RouteProfile AddAsPathToListOfLists
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) Add4ByteAsPathToListOfLists ()(string, error) {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType AsNumberList
 //AgtBgp4Ipv4RouteProfile Add4ByteAsPathToListOfLists
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) RemoveAsPathFromListOfLists ()(string, error) {
 //parameters: RouteProfileHandle AsPathAttributeType AsPathListHandle
 //AgtBgp4Ipv4RouteProfile RemoveAsPathFromListOfLists
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) Remove4ByteAsPathFromListOfLists ()(string, error) {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType AsPathListHandle
 //AgtBgp4Ipv4RouteProfile Remove4ByteAsPathFromListOfLists
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) ClearAsPathListOfLists ()(string, error) {
 //parameters: RouteProfileHandle AsPathAttributeType
 //AgtBgp4Ipv4RouteProfile ClearAsPathListOfLists
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) Clear4ByteAsPathListOfLists ()(string, error) {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType
 //AgtBgp4Ipv4RouteProfile Clear4ByteAsPathListOfLists
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) ListAsPathListOfLists ()(string, error) {
 //parameters: RouteProfileHandle AsPathAttributeType AsPathListHandle
 //AgtBgp4Ipv4RouteProfile ListAsPathListOfLists
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) List4ByteAsPathListOfLists ()(string, error) {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeType AsPathListHandle
 //AgtBgp4Ipv4RouteProfile List4ByteAsPathListOfLists
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) GetAsPathAuto ()(string, error) {
 //parameters: RouteProfileHandle AsPathAttributeTypeAuto
 //AgtBgp4Ipv4RouteProfile GetAsPathAuto
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) Get4ByteAsPathAuto ()(string, error) {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeTypeAuto
 //AgtBgp4Ipv4RouteProfile Get4ByteAsPathAuto
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) Get4ByteAsPathAutoWithFormat ()(string, error) {
 //parameters: RouteProfileHandle 4ByteAsPathAttributeTypeAuto FourByteFormatType
 //AgtBgp4Ipv4RouteProfile Get4ByteAsPathAutoWithFormat
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) EnablePathAttribute () error {
 //parameters: RouteProfileHandle PathAttribute
 //AgtBgp4Ipv4RouteProfile EnablePathAttribute, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) DisablePathAttribute () error {
 //parameters: RouteProfileHandle PathAttribute
 //AgtBgp4Ipv4RouteProfile DisablePathAttribute, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) IsPathAttributeEnabled () error {
 //parameters: RouteProfileHandle PathAttribute
 //AgtBgp4Ipv4RouteProfile IsPathAttributeEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) SetOrigin () error {
 //parameters: RouteProfileHandle Origin
 //AgtBgp4Ipv4RouteProfile SetOrigin, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) GetOrigin ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetOrigin
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) SetLocalPreference () error {
 //parameters: RouteProfileHandle LocalPreference
 //AgtBgp4Ipv4RouteProfile SetLocalPreference, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) GetLocalPreference ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetLocalPreference
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) SetMultiExitDiscriminator () error {
 //parameters: RouteProfileHandle MultiExitDiscriminator
 //AgtBgp4Ipv4RouteProfile SetMultiExitDiscriminator, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) GetMultiExitDiscriminator ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetMultiExitDiscriminator
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) SetAggregatorIncrementingRange () error {
 //parameters: RouteProfileHandle FirstAsNumber AsNumberIncrement AsNumberRepeat FirstIpAddress PrefixLength AddressIncrement AddressRepeat
 //AgtBgp4Ipv4RouteProfile SetAggregatorIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) Set4ByteAsAggregatorIncrementingRange () error {
 //parameters: RouteProfileHandle First4ByteAsNumber AsNumberIncrement AsNumberRepeat FirstIpAddress PrefixLength AddressIncrement AddressRepeat
 //AgtBgp4Ipv4RouteProfile Set4ByteAsAggregatorIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) Set4ByteAsAggregatorIncrementingRangeWithFormat () error {
 //parameters: RouteProfileHandle FourByteFormatType First4ByteAsNumber AsNumberIncrement AsNumberRepeat FirstIpAddress PrefixLength AddressIncrement AddressRepeat
 //AgtBgp4Ipv4RouteProfile Set4ByteAsAggregatorIncrementingRangeWithFormat, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) GetAggregatorIncrementingRange ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetAggregatorIncrementingRange
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) Get4ByteAsAggregatorIncrementingRange ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile Get4ByteAsAggregatorIncrementingRange
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) Get4ByteAsAggregatorIncrementingRangeWithFormat ()(string, error) {
 //parameters: RouteProfileHandle FourByteFormatType
 //AgtBgp4Ipv4RouteProfile Get4ByteAsAggregatorIncrementingRangeWithFormat
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) SetAggregatorList ()(string, error) {
 //parameters: RouteProfileHandle AsNumberList IpAddressList
 //AgtBgp4Ipv4RouteProfile SetAggregatorList
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) Set4ByteAsAggregatorList ()(string, error) {
 //parameters: RouteProfileHandle AsNumberList IpAddressList
 //AgtBgp4Ipv4RouteProfile Set4ByteAsAggregatorList
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) Set4ByteAsAggregatorListWithFormat ()(string, error) {
 //parameters: RouteProfileHandle FourByteFormatType AsNumberList IpAddressList
 //AgtBgp4Ipv4RouteProfile Set4ByteAsAggregatorListWithFormat
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) GetAggregatorList ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetAggregatorList
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) Get4ByteAsAggregatorList ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile Get4ByteAsAggregatorList
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) Get4ByteAsAggregatorListWithFormat ()(string, error) {
 //parameters: RouteProfileHandle FourByteFormatType
 //AgtBgp4Ipv4RouteProfile Get4ByteAsAggregatorListWithFormat
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) SetAggregator () error {
 //parameters: RouteProfileHandle AsNumber IpAddress
 //AgtBgp4Ipv4RouteProfile SetAggregator, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) Set4ByteAsAggregator () error {
 //parameters: RouteProfileHandle AsNumber IpAddress
 //AgtBgp4Ipv4RouteProfile Set4ByteAsAggregator, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) Set4ByteAsAggregatorWithFormat () error {
 //parameters: RouteProfileHandle FourByteFormatType AsNumber IpAddress
 //AgtBgp4Ipv4RouteProfile Set4ByteAsAggregatorWithFormat, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) GetAggregator ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetAggregator
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) Get4ByteAsAggregator ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile Get4ByteAsAggregator
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) Get4ByteAsAggregatorWithFormat ()(string, error) {
 //parameters: RouteProfileHandle FourByteFormatType
 //AgtBgp4Ipv4RouteProfile Get4ByteAsAggregatorWithFormat
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) GetAggregatorAuto ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetAggregatorAuto
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) Get4ByteAggregatorAuto ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile Get4ByteAggregatorAuto
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) Get4ByteAggregatorAutoWithFormat ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile Get4ByteAggregatorAutoWithFormat
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) SetOriginatorIdIncrementingRange () error {
 //parameters: RouteProfileHandle FirstOriginatorId PrefixLength AddressIncrement AddressRepeat
 //AgtBgp4Ipv4RouteProfile SetOriginatorIdIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) GetOriginatorIdIncrementingRange ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetOriginatorIdIncrementingRange
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) SetOriginatorIdList ()(string, error) {
 //parameters: RouteProfileHandle OriginatorIdList
 //AgtBgp4Ipv4RouteProfile SetOriginatorIdList
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) GetOriginatorIdList ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetOriginatorIdList
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) SetOriginatorId () error {
 //parameters: RouteProfileHandle OriginatorId
 //AgtBgp4Ipv4RouteProfile SetOriginatorId, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) GetOriginatorId ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetOriginatorId
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) GetOriginatorIdAuto ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetOriginatorIdAuto
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) SetClusterList ()(string, error) {
 //parameters: RouteProfileHandle ClusterList
 //AgtBgp4Ipv4RouteProfile SetClusterList
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) GetClusterList ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetClusterList
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) GetClusterListAuto ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetClusterListAuto
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) SetCommunities () error {
 //parameters: RouteProfileHandle CommunitiesList
 //AgtBgp4Ipv4RouteProfile SetCommunities, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) GetCommunities ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetCommunities
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) SetCommunitiesIncrementingRange () error {
 //parameters: RouteProfileHandle CommunitiesList CommunityIncrement CommunityRepeat
 //AgtBgp4Ipv4RouteProfile SetCommunitiesIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) GetCommunitiesIncrementingRange ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetCommunitiesIncrementingRange
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) SetCommunitiesList ()(string, error) {
 //parameters: RouteProfileHandle CommunitiesList
 //AgtBgp4Ipv4RouteProfile SetCommunitiesList
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) GetCommunitiesList ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetCommunitiesList
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) EnableTrafficDestinations () error {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile EnableTrafficDestinations, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) DisableTrafficDestinations () error {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile DisableTrafficDestinations, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) IsTrafficDestinationsEnabled () error {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile IsTrafficDestinationsEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) SetFlapDistribution () error {
 //parameters: RouteProfileHandle FlapPercentage FlapOffset
 //AgtBgp4Ipv4RouteProfile SetFlapDistribution, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) GetFlapDistribution ()(string, error) {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile GetFlapDistribution
 return "", nil
}

func(np *Bgp4Ipv4RouteProfile) Advertise () error {
 //parameters: RouteProfileIdentifiers
 //AgtBgp4Ipv4RouteProfile Advertise, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) Withdraw () error {
 //parameters: RouteProfileIdentifiers
 //AgtBgp4Ipv4RouteProfile Withdraw, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) Enable () error {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile Enable, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv4RouteProfile) Disable () error {
 //parameters: RouteProfileHandle
 //AgtBgp4Ipv4RouteProfile Disable, m.Object, m.Name);
 return nil
}

