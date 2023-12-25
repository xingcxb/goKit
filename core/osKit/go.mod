module github.com/xingcxb/goKit/core/osKit

go 1.21

require (
	github.com/dkorunic/iSMC v0.7.0
	github.com/shirou/gopsutil/v3 v3.23.11
	github.com/xingcxb/goKit/core/dateKit v0.0.0-20231225083528-2c30b772efbb
)

require (
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/panotza/gosmc v1.0.0 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/xingcxb/goKit/core/arrayKit v0.0.0-20231225070946-93ee4af6654a // indirect
	github.com/xingcxb/goKit/core/numKit v0.0.0-20231225083528-2c30b772efbb // indirect
	github.com/xingcxb/goKit/core/regKit v0.0.0-20231225083528-2c30b772efbb // indirect
	github.com/xingcxb/goKit/core/strKit v0.0.0-20231225083528-2c30b772efbb // indirect
	github.com/yusufpapurcu/wmi v1.2.3 // indirect
	golang.org/x/sys v0.15.0 // indirect
)

replace (
	github.com/xingcxb/goKit/core/arrayKit => ../arrayKit
	github.com/xingcxb/goKit/core/dateKit => ../dateKit
	github.com/xingcxb/goKit/core/numKit => ../numKit
	github.com/xingcxb/goKit/core/osKit => ./
	github.com/xingcxb/goKit/core/regKit => ../regKit
	github.com/xingcxb/goKit/core/strKit => ../strKit
)
