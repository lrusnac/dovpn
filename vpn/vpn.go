// Copyright © 2017 Leonid Rusanc <leonidrusnac4@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vpn

import (
	"context"
	"fmt"

	"github.com/digitalocean/godo"
)

// NewVpnInstance creates a new vps with openvpn
func NewVpnInstance(client *godo.Client) error {
	createRequest := &godo.DropletCreateRequest{
		Name:   "vpn",
		Region: "fra1",
		Size:   "512mb",
		Image: godo.DropletCreateImage{
			ID: 28912386,
		},
	}

	ctx := context.TODO()

	_, _, err := client.Droplets.Create(ctx, createRequest)

	if err != nil {
		fmt.Printf("Something bad happened: %s\n", err)
		return err
	}

	return nil
}

// FindVpnInstance returns the id of the vpn instance
func FindVpnInstance(client *godo.Client) (int, error) {
	droplets, err := dropletList(context.TODO(), client)
	if err != nil {
		fmt.Printf("Something happened while retrieving droplet list %s\n", err)
	}

	for _, d := range droplets {
		if d.Name == "vpn" {
			return d.ID, nil
		}
	}

	return -1, nil
}

// DropVpnInstance destroys the vpn instance if exists
func DropVpnInstance(client *godo.Client) error {
	dropletID, err := FindVpnInstance(client)
	if err != nil {
		fmt.Printf("Something happened while retrieving the vpn droplet%s\n", err)
	}

	_, err = client.Droplets.Delete(context.TODO(), dropletID)

	return err
}

func dropletList(ctx context.Context, client *godo.Client) ([]godo.Droplet, error) {
	// create a list to hold our droplets
	list := []godo.Droplet{}

	// create options. initially, these will be blank
	opt := &godo.ListOptions{}
	for {
		droplets, resp, err := client.Droplets.List(ctx, opt)
		if err != nil {
			return nil, err
		}

		// append the current page's droplets to our list
		for _, d := range droplets {
			list = append(list, d)
		}

		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}

		// set the page we want for the next request
		opt.Page = page + 1
	}

	return list, nil
}
