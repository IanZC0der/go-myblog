package impl

import (
	"github.com/IanZC0der/go-myblog/apps/blog"
	"github.com/IanZC0der/go-myblog/ioc"
)

func init() {
	ioc.DefaultControllerContainer().Register(&blogServiceImpl{})
}

type blogServiceImpl struct {
}

func (b *blogServiceImpl) Init() error {
	return nil
}

func (b *blogServiceImpl) Name() string {
	return blog.AppName
}
