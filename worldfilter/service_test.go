package worldfilter

import (
	_ "embed"
	"testing"
)

//go:embed key
var sk string

func TestFilter(t *testing.T) {
	var s = Service{
		API:       "http://muse.seasungame.com:10104/filter",
		AppId:     "xsjweb",
		ChannelId: "jx3box",
		SecretKey: sk,
		fnSwitch:  func() bool { return true },
	}
	r, _ := s.Filter("剑侠3牛逼")
	t.Log(r)

}
