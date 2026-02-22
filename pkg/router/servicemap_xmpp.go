//go:build xmpp

package router

import t "github.com/circa10a/shoutrrr/pkg/types"

func init() {
	serviceMap["xmpp"] = func() t.Service { return &xmpp.Service{} }
}
