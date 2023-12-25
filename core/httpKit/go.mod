module github.com/xingcxb/goKit/core/httpKit

go 1.21

require (
	github.com/hashicorp/golang-lru v1.0.2
	github.com/xingcxb/goKit/core/strKit v0.0.0-20231111101058-bdee5e13e835
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/xingcxb/goKit/core/arrayKit v0.0.0-20231225070946-93ee4af6654a // indirect
	github.com/xingcxb/goKit/core/numKit v0.0.0-20231225083528-2c30b772efbb // indirect
	github.com/xingcxb/goKit/core/regKit v0.0.0-20231225083528-2c30b772efbb // indirect
)

replace (
	github.com/xingcxb/goKit/core/arrayKit => ../arrayKit
	github.com/xingcxb/goKit/core/httpKit => ./
	github.com/xingcxb/goKit/core/numKit => ../numKit
	github.com/xingcxb/goKit/core/regKit => ../regKit
	github.com/xingcxb/goKit/core/strKit => ../strKit
)
