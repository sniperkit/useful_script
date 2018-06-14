package n2xsdk

type MulticastGroupPool struct {
 Handler string
}

func(np *MulticastGroupPool) GetType ()(string, error) {
 //parameters: Handle
 //AgtMulticastGroupPool GetType
 return "", nil
}

func(np *MulticastGroupPool) GetName ()(string, error) {
 //parameters: Handle
 //AgtMulticastGroupPool GetName
 return "", nil
}

func(np *MulticastGroupPool) GetLockCount ()(string, error) {
 //parameters: Handle
 //AgtMulticastGroupPool GetLockCount
 return "", nil
}

func(np *MulticastGroupPool) SetMulticastAddressRange () error {
 //parameters: PoolHandle FirstMulticastAddress NumAddresses Modifier
 //AgtMulticastGroupPool SetMulticastAddressRange, m.Object, m.Name);
 return nil
}

func(np *MulticastGroupPool) GetMulticastAddressRange ()(string, error) {
 //parameters: PoolHandle
 //AgtMulticastGroupPool GetMulticastAddressRange
 return "", nil
}

func(np *MulticastGroupPool) SetMulticastAddressRangeWithPrefix () error {
 //parameters: PoolHandle FirstMulticastAddress PrefixLength NumAddresses Modifier
 //AgtMulticastGroupPool SetMulticastAddressRangeWithPrefix, m.Object, m.Name);
 return nil
}

func(np *MulticastGroupPool) GetMulticastAddressRangeWithPrefix ()(string, error) {
 //parameters: PoolHandle
 //AgtMulticastGroupPool GetMulticastAddressRangeWithPrefix
 return "", nil
}

func(np *MulticastGroupPool) ListReservedGroups ()(string, error) {
 //parameters: 
 //AgtMulticastGroupPool ListReservedGroups
 return "", nil
}

func(np *MulticastGroupPool) AddReservedGroup () error {
 //parameters: MulticastAddress
 //AgtMulticastGroupPool AddReservedGroup, m.Object, m.Name);
 return nil
}

func(np *MulticastGroupPool) RemoveReservedGroup () error {
 //parameters: MulticastAddress
 //AgtMulticastGroupPool RemoveReservedGroup, m.Object, m.Name);
 return nil
}

func(np *MulticastGroupPool) ListIpv6ReservedGroups ()(string, error) {
 //parameters: 
 //AgtMulticastGroupPool ListIpv6ReservedGroups
 return "", nil
}

func(np *MulticastGroupPool) AddIpv6ReservedGroup () error {
 //parameters: MulticastAddress
 //AgtMulticastGroupPool AddIpv6ReservedGroup, m.Object, m.Name);
 return nil
}

func(np *MulticastGroupPool) RemoveIpv6ReservedGroup () error {
 //parameters: MulticastAddress
 //AgtMulticastGroupPool RemoveIpv6ReservedGroup, m.Object, m.Name);
 return nil
}

func(np *MulticastGroupPool) EnableTrafficDestinations () error {
 //parameters: PoolHandle
 //AgtMulticastGroupPool EnableTrafficDestinations, m.Object, m.Name);
 return nil
}

func(np *MulticastGroupPool) DisableTrafficDestinations () error {
 //parameters: PoolHandle
 //AgtMulticastGroupPool DisableTrafficDestinations, m.Object, m.Name);
 return nil
}

func(np *MulticastGroupPool) IsTrafficDestinationEnabled () error {
 //parameters: PoolHandle
 //AgtMulticastGroupPool IsTrafficDestinationEnabled, m.Object, m.Name);
 return nil
}

