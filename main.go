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

package main

import (
	"context"
	"fmt"

	"github.com/lrusnac/dovpn/vpn"

	"github.com/digitalocean/godo"
	"github.com/lrusnac/dovpn/cmd"
	"golang.org/x/oauth2"
)

type tokenSource struct {
	AccessToken string
}

func (t *tokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func main() {
	cmd.Execute()

	pat := "aaaa"
	tokenSource := &tokenSource{
		AccessToken: pat,
	}
	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)

	err := vpn.NewVpnInstance(client)
	if err != nil {
		fmt.Printf("error creating the vpn instance: %v\n", err)
	}

	dropletID, err := vpn.FindVpnInstance(client)
	if err != nil {
		fmt.Printf("error finding the vpn instance: %v\n", err)
	}
	fmt.Printf("found droplet: %d\n", dropletID)

	err = vpn.DropVpnInstance(client)
	if err != nil {
		fmt.Printf("error destroying the vpn instance: %v\n", err)
	}
}
