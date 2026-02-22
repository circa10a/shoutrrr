package format

import t "github.com/circa10a/shoutrrr/pkg/types"

type testEnummer struct {
	Choice int `default:"Maybe" key:"choice"`
}

func (testEnummer) Enums() map[string]t.EnumFormatter {
	return map[string]t.EnumFormatter{
		"Choice": CreateEnumFormatter([]string{"Yes", "No", "Maybe"}),
	}
}

func testRenderTree(r TreeRenderer, v any) string {
	return r.RenderTree(getRootNode(v), "mock")
}
