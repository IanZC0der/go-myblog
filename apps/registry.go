package apps

import (
	_ "github.com/IanZC0der/go-myblog/apps/blog/api"
	_ "github.com/IanZC0der/go-myblog/apps/blog/impl"
	_ "github.com/IanZC0der/go-myblog/apps/comment/api"
	_ "github.com/IanZC0der/go-myblog/apps/comment/impl"
	_ "github.com/IanZC0der/go-myblog/apps/mq"
	_ "github.com/IanZC0der/go-myblog/apps/token/api"
	_ "github.com/IanZC0der/go-myblog/apps/token/impl"
	_ "github.com/IanZC0der/go-myblog/apps/user/impl"
)
