package sqlite

import (
	"testing"

	"github.com/anvh2/z-blogs/grpc-gen/blog"
	"github.com/stretchr/testify/assert"
)

var testBlogDb *BlogDb

func TestMain(m *testing.M) {
	
	testBlogDb = NewBlogDb()
}

func TestConv(t *testing.T) {
	item := &blog.BlogData{
		Tags:   []string{"Tech"},
		Images: []string{"/app/demo.png"},
	}

	fillData(item)
	assert.Equal(t, "[\"Tech\"]", item.TagStr)
	assert.Equal(t, "[\"/app/demo.png\"]", item.ImagesStr)

	item1 := &blog.BlogData{
		TagStr:    "[\"Tech\"]",
		ImagesStr: "[\"/app/demo.png\"]",
	}

	fillData(item1)
	assert.Equal(t, []string{"Tech"}, item1.Tags)
	assert.Equal(t, []string{"/app/demo.png"}, item1.Images)
}

func TestInterface(t *testing.T) {
	item := &blog.BlogData{
		Comments: []*blog.Comment{
			&blog.Comment{
				Author:  "anvh2",
				Content: "Greate",
			},
		},
		Tags:   []string{"Tech"},
		Images: []string{"/app/demo.png"},
	}

}
