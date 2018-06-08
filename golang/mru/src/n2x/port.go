package n2x

import (
	"fmt"
	"strings"
)

type Port struct {
	Name    string
	Handler string
	*NSession
}

func (p *Port) LegacyGetSutIPAddresses() ([]string, error) {
	cmd := fmt.Sprintf("AgtEthernetAddresses ListSutIpAddresses %s", p.Handler)
	res, err := p.Invoke(cmd)
	if err != nil {
		return nil, fmt.Errorf("Cannot get sup ip address : %s", err.Error())
	}

	ips := make([]string, 0, 10)
	fields := strings.Split(res, " ")
	for _, field := range fields {
		ips = append(ips, strings.TrimSpace(field))
	}

	return ips, nil
}
