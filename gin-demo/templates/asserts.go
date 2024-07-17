package templates

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets90c6e63561da4ea648316ad4669eaa00289ec369 = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>Index Template</title>\n</head>\n<body>\n    <p>Hello, {{.Foo}}</p>\n</body>\n</html>"
var _Assets8daf7115b2ff9fea2c99f960639db5ab04fb10c6 = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n  <meta charset=\"UTF-8\">\n  <title>Bar Template</title>\n</head>\n<body>\n  <p>Can you see this? - {{.Bar}}</p>\n</body>\n</html>"
var _Assets014f674c1d8c8b0d915c0d5545c4b0cf6c73cd60 = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>Index Template</title>\n</head>\n<body>\n    <p>Hello, {{.Foo}}</p>\n</body>\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"bar.htm", "index.htm", "users"}, "/users": []string{"index.htm"}}, map[string]*assets.File{
	"/users/index.htm": &assets.File{
		Path:     "/users/index.htm",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1682730843, 1682730843977984183),
		Data:     []byte(_Assets90c6e63561da4ea648316ad4669eaa00289ec369),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1683949052, 1683949052831903127),
		Data:     nil,
	}, "/bar.htm": &assets.File{
		Path:     "/bar.htm",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1682731752, 1682731752179962126),
		Data:     []byte(_Assets8daf7115b2ff9fea2c99f960639db5ab04fb10c6),
	}, "/index.htm": &assets.File{
		Path:     "/index.htm",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1682730843, 1682730843977984183),
		Data:     []byte(_Assets014f674c1d8c8b0d915c0d5545c4b0cf6c73cd60),
	}, "/users": &assets.File{
		Path:     "/users",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1683948979, 1683948979046993141),
		Data:     nil,
	}}, "")
