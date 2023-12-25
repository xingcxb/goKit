module github.com/xingcxb/goKit/core/randomKit

go 1.21

require (
	github.com/xingcxb/goKit/core/dateKit v0.0.0-20231225083528-2c30b772efbb
	github.com/xingcxb/goKit/core/strKit v0.0.0-20231225083528-2c30b772efbb
)

require (
	github.com/xingcxb/goKit/core/arrayKit v0.0.0-20231225070946-93ee4af6654a // indirect
	github.com/xingcxb/goKit/core/numKit v0.0.0-20231225083528-2c30b772efbb // indirect
	github.com/xingcxb/goKit/core/regKit v0.0.0-20231225083528-2c30b772efbb // indirect
)

replace (
	github.com/xingcxb/goKit/core/randomKit => ./
	github.com/xingcxb/goKit/core/arrayKit => ../arrayKit
	github.com/xingcxb/goKit/core/dateKit => ../dateKit
	github.com/xingcxb/goKit/core/numKit => ../numKit
	github.com/xingcxb/goKit/core/regKit => ../regKit
	github.com/xingcxb/goKit/core/strKit => ../strKit
)
