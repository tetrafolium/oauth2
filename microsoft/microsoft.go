// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package microsoft provides constants for using OAuth2 to access Windows Live ID.
package microsoft

import (
	"github.com/tetrafolium/oauth2"
)

// LiveConnectEndpoint is Windows's Live ID OAuth 2.0 endpoint.
var LiveConnectEndpoint = oauth2.Endpoint{
	AuthURL:  "https://login.live.com/oauth20_authorize.srf",
	TokenURL: "https://login.live.com/oauth20_token.srf",
}
