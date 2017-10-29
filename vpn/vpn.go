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

	newDroplet, _, err := client.Droplets.Create(ctx, createRequest)

	if err != nil {
		fmt.Printf("Something bad happened: %s\n", err)
		return err
	}

	fmt.Printf("the new vpn server is at: %v", newDroplet)
	return nil
}

// ExistsVpnInstance returns true or false if there is a vpn instance
func ExistsVpnInstance(client *godo.Client) (bool, error) {
	return false, nil
}

// DropVpnInstance destroys the vpn instance if exists
func DropVpnInstance(client *godo.Client) error {
	return nil
}
