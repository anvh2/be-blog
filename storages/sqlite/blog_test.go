package sqlite

import (
	"context"
	"testing"

	"github.com/anvh2/z-blogs/grpc-gen/blog"
	// include gorm sqlite
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

var testBlogDb *BlogDb

func TestMain(m *testing.M) {
	db, _ := gorm.Open("sqlite3", ":memory:")
	defer db.Close()
	logger, _ := zap.NewProduction()
	testBlogDb = NewBlogDb(db, logger)
	defer testBlogDb.Close()
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
	fillData(item)

	ctx := context.Background()

	err := testBlogDb.Create(ctx, item)
	assert.Nil(t, err)

	g, err := testBlogDb.Get(ctx, item.Id)
	assert.Nil(t, err)
	assert.Equal(t, item, fillData(g))
}
