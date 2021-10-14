// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package cmp

import (
	"context"
	"errors"
	"log"

	"github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/client"
	"github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/models"
	"github.com/HewlettPackard/hpegl-vmaas-terraform-resources/internal/utils"
	"github.com/tshihad/tftags"
)

type domain struct {
	dClient *client.DomainAPIService
}

func newDomain(dClient *client.DomainAPIService) *domain {
	return &domain{dClient: dClient}
}

func (n *domain) Read(ctx context.Context, d *utils.Data, meta interface{}) error {
	log.Printf("[INFO] Get Domain")

	name := d.GetString("name")

	// Pre check
	if err := d.Error(); err != nil {
		return err
	}
	// Get all domain with filter as name
	resp, err := utils.Retry(ctx, meta, func(ctx context.Context) (interface{}, error) {
		return n.dClient.GetAllDomains(ctx, map[string]string{nameKey: name})
	})
	if err != nil {
		return err
	}
	domains := resp.(models.GetAllDomains)
	if len(domains.NetworkDomains) != 1 {
		return errors.New("error coudn't find exact domain, please check the name")
	}

	tftags.Set(d, domains.NetworkDomains[0])

	// post check
	return d.Error()
}
