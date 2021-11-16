// Copyright 2020 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gmailfilter

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"google.golang.org/api/gmail/v1"
)

type FilterGenerator struct {
	GmailfilterService
}

func (g FilterGenerator) createResources(filters []*gmail.Filter) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, f := range filters {
		resources = append(resources, terraformutils.NewResource(
			f.Id,
			f.Id,
			"gmailfilter_filter",
			"gmailfilter",
			map[string]string{},
			[]string{},
			map[string]interface{}{}))
	}
	return resources
}

func (g *FilterGenerator) InitResources() error {
	ctx := context.Background()
	gmailService, err := g.gmailService(ctx)
	if err != nil {
		return err
	}

	filters, err := gmailService.Users.Settings.Filters.List(gmailUser).Do()
	if err != nil {
		return err
	}
	g.Resources = append(g.Resources, g.createResources(filters.Filter)...)

	return nil
}
