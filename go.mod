module github.com/xingcxb/goKit

go 1.21

toolchain go1.21.1

require (
	github.com/redis/go-redis/v9 v9.3.1
	github.com/tidwall/gjson v1.17.0
	github.com/xingcxb/goKit/core/arrayKit v0.0.0-20231225070946-93ee4af6654a
	github.com/xingcxb/goKit/core/cryptoKit v0.0.0-00010101000000-000000000000
	github.com/xingcxb/goKit/core/fileKit v0.0.0-00010101000000-000000000000
	github.com/xingcxb/goKit/core/httpKit v0.0.0-00010101000000-000000000000
	github.com/xingcxb/goKit/core/randomKit v0.0.0-00010101000000-000000000000
	github.com/xingcxb/goKit/core/reflectKit v0.0.0-00010101000000-000000000000
	github.com/xingcxb/goKit/core/strKit v0.0.0-20231225083528-2c30b772efbb
	github.com/xuri/excelize/v2 v2.8.0
	github.com/yeqown/go-qrcode/v2 v2.2.2
	github.com/yeqown/go-qrcode/writer/standard v1.2.2
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fogleman/gg v1.3.0 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/richardlehane/mscfb v1.0.4 // indirect
	github.com/richardlehane/msoleps v1.0.3 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/xingcxb/goKit/core/dateKit v0.0.0-20231225083528-2c30b772efbb // indirect
	github.com/xingcxb/goKit/core/numKit v0.0.0-20231225083528-2c30b772efbb // indirect
	github.com/xingcxb/goKit/core/regKit v0.0.0-20231225083528-2c30b772efbb // indirect
	github.com/xuri/efp v0.0.0-20230802181842-ad255f2331ca // indirect
	github.com/xuri/nfp v0.0.0-20230919160717-d98342af3f05 // indirect
	github.com/yeqown/reedsolomon v1.0.0 // indirect
	golang.org/x/crypto v0.13.0 // indirect
	golang.org/x/image v0.12.0 // indirect
	golang.org/x/net v0.15.0 // indirect
	golang.org/x/text v0.13.0 // indirect
)

replace (
	github.com/xingcxb/goKit/core/arrayKit => ./core/arrayKit
	github.com/xingcxb/goKit/core/cryptoKit => ./core/cryptoKit
	github.com/xingcxb/goKit/core/dateKit => ./core/dateKit
	github.com/xingcxb/goKit/core/fileKit => ./core/fileKit
	github.com/xingcxb/goKit/core/httpKit => ./core/httpKit
	github.com/xingcxb/goKit/core/ipKit => ./core/ipKit
	github.com/xingcxb/goKit/core/numKit => ./core/numKit
	github.com/xingcxb/goKit/core/osKit => ./core/osKit
	github.com/xingcxb/goKit/core/otherKit => ./core/otherKit
	github.com/xingcxb/goKit/core/pathKit => ./core/pathKit
	github.com/xingcxb/goKit/core/randomKit => ./core/randomKit
	github.com/xingcxb/goKit/core/reflectKit => ./core/reflectKit
	github.com/xingcxb/goKit/core/regKit => ./core/regKit
	github.com/xingcxb/goKit/core/runTimeKit => ./core/runTimeKit
	github.com/xingcxb/goKit/core/strKit => ./core/strKit
)
