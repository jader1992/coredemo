package framework

import (
	"testing"
)

func Test_filterChildNodes(t *testing.T)  {
	root := &node{
		isLast: false,
		segment: "",
		handler: func(c *Context) error {
			return nil
		},
		childes: []*node{
			{
				isLast: true,
				segment: "FOO",
				handler: func(c *Context) error {
					return nil
				},
				childes: nil,
			},
			{
				isLast: false,
				segment: ":id",
				handler: nil,
				childes: nil,
			},
		},
	}

	nodes := root.filterChildNodes("FOO")
	if len(nodes) != 2 {
		t.Error("foo error")
	}
	t.Log(nodes)

	nodes = root.filterChildNodes(":foo")
	if len(nodes) != 2 {
		t.Error(":foo error")
	}
	t.Log(nodes)
}

func Test_matchNode(t *testing.T) {
	root := &node{
		isLast:  false,
		segment: "",
		handler: func(*Context) error { return nil },
		childes: []*node{
			{
				isLast:  true,
				segment: "FOO",
				handler: nil,
				childes: []*node{
					{
						isLast:   true,
						segment:  "BAR",
						handler:  func(*Context) error { panic("not implemented") },
						childes:  []*node{},
					},
				},
			},
			{
				isLast:  true,
				segment: ":id",
				handler: nil,
				childes:  nil,
			},
		},
	}

	{
		node := root.matchNode("foo/bar")
		if node == nil {
			t.Error("match normal node error")
		}
	}

	{
		node := root.matchNode("test")
		if node == nil {
			t.Error("match test")
		}
	}

}
