module github.com/xingcxb/goKit/core/ipKit

go 1.21

require (
	github.com/PuerkitoBio/goquery v1.8.1
	github.com/tidwall/gjson v1.17.0
	github.com/xingcxb/goKit/core/httpKit v0.0.0-20231225070946-93ee4af6654a
	github.com/xingcxb/goKit/core/randomKit v0.0.0-20231225083528-2c30b772efbb
	github.com/xingcxb/goKit/core/strKit v0.0.0-20231225083528-2c30b772efbb
)

require (
	github.com/andybalholm/cascadia v1.3.1 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/xingcxb/goKit/core/arrayKit v0.0.0-20231225070946-93ee4af6654a // indirect
	github.com/xingcxb/goKit/core/dateKit v0.0.0-20231225083528-2c30b772efbb // indirect
	github.com/xingcxb/goKit/core/numKit v0.0.0-20231225083528-2c30b772efbb // indirect
	github.com/xingcxb/goKit/core/regKit v0.0.0-20231225083528-2c30b772efbb // indirect
	golang.org/x/net v0.7.0 // indirect
)

replace (
	github.com/xingcxb/goKit/core/ipKit => ./
	github.com/xingcxb/goKit/core/arrayKit => ../arrayKit
	github.com/xingcxb/goKit/core/httpKit => ../httpKit
	github.com/xingcxb/goKit/core/numKit => ../numKit
	github.com/xingcxb/goKit/core/randomKit => ../randomKit
	github.com/xingcxb/goKit/core/regKit => ../regKit
	github.com/xingcxb/goKit/core/strKit => ../strKit
)
