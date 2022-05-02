// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type domainConfigFile struct {
	Domains []domain `json:"domains"`
}

type domain struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	BaseCurrencyFullName string `json:"baseCurrencyFullName"`
	BaseCurrencySymbol   string `json:"baseCurrencySymbol"`
	AddressPrefix        string `json:"addressPrefix"`
}

func newDomain(id int, name, baseCurrencyFullName, baseCurrencySymbol, addressPrefix string) *domain {
	return &domain{
		ID:                   id,
		Name:                 name,
		BaseCurrencyFullName: baseCurrencyFullName,
		BaseCurrencySymbol:   baseCurrencySymbol,
		AddressPrefix:        addressPrefix,
	}
}

// loadDomains registers and load all pre-defined domains
func loadDomains(domainConfigPath string) map[int]domain {
	domains := make(map[int]domain, 0)

	domainData, err := ioutil.ReadFile(filepath.Clean(domainConfigPath))
	if err != nil {
		fmt.Println(err)
		panic(ErrLoadDomainConfig.Wrap(err))
	}

	var content domainConfigFile
	err = json.Unmarshal(domainData, &content)
	if err != nil {
		panic(ErrLoadDomainConfig.Wrap(err))
	}

	for _, domain := range content.Domains {
		domains[domain.ID] = *newDomain(domain.ID, domain.Name, domain.BaseCurrencyFullName, domain.BaseCurrencySymbol, domain.AddressPrefix)
	}

	return domains
}
