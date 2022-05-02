// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

const (
	NativeCurrencyAddr = "0x0000000000000000000000000000000000000000"
)

type resourceConfigFile struct {
	Resources []resource `json:"resources"`
}

type resource struct {
	ID           string `json:"id"`
	Symbol       string `json:"symbol"`
	DomainId     int    `json:"domainId"`
	Domain       domain `json:"domain"`
	Decimal      int    `json:"decimal"`
	TokenAddress string `json:"tokenAddress"`
}

func newResource(tokenAddress string, decimal int, symbol string, domain domain) *resource {
	return &resource{
		ID:           ResourceIDBuilder(tokenAddress, domain.ID),
		Symbol:       symbol,
		Domain:       domain,
		DomainId:     domain.ID,
		Decimal:      decimal,
		TokenAddress: tokenAddress,
	}
}

// loadResources registers and load all pre-defined resources
func loadResources(resourceConfigPath string, domains map[int]domain) map[string]resource {
	resources := make(map[string]resource, 0)

	resourceData, err := ioutil.ReadFile(filepath.Clean(resourceConfigPath))
	if err != nil {
		panic(ErrLoadResourceConfig.Wrap(err))
	}

	var content resourceConfigFile
	err = json.Unmarshal(resourceData, &content)
	if err != nil {
		panic(ErrLoadResourceConfig.Wrap(err))
	}

	for _, resource := range content.Resources {
		resources[ResourceIDBuilder(resource.TokenAddress, resource.DomainId)] =
			*newResource(strings.ToLower(resource.TokenAddress), resource.Decimal, resource.Symbol, domains[resource.DomainId])
	}

	return resources
}

// ResourceIDBuilder builds the resourceID according to fee handler contract
func ResourceIDBuilder(tokenAddress string, domainId int) string {
	return fmt.Sprintf("%s%d", strings.ToLower(tokenAddress), domainId)
}
