.PHONY: proto
proto:
	# 服务器导出
	protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=bitbucket.org/funplus/server/proto --micro_opt=paths=source_relative ./global/common/*.proto
	protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=bitbucket.org/funplus/server/proto --micro_opt=paths=source_relative ./server/game/*.proto
	protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=bitbucket.org/funplus/server/proto --micro_opt=paths=source_relative ./server/gate/*.proto
	protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=bitbucket.org/funplus/server/proto --micro_opt=paths=source_relative ./server/combat/*.proto
	protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=bitbucket.org/funplus/server/proto --micro_opt=paths=source_relative ./server/mail/*.proto
	protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=bitbucket.org/funplus/server/proto --micro_opt=paths=source_relative ./server/pubsub/*.proto

	# 客户端导出
	# protoc -I=./ --csharp_out=:../ee_client/proto --csharp_opt=base_namespace=common ./global/common/*.proto

.PHONY: linux_proto
linux_proto:
	# 服务器导出
	./bin/protoc_server -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=bitbucket.org/funplus/server/proto --micro_opt=paths=source_relative ./global/common/*.proto
	./bin/protoc_server -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=bitbucket.org/funplus/server/proto --micro_opt=paths=source_relative ./server/game/*.proto
	./bin/protoc_server -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=bitbucket.org/funplus/server/proto --micro_opt=paths=source_relative ./server/gate/*.proto
	./bin/protoc_server -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=bitbucket.org/funplus/server/proto --micro_opt=paths=source_relative ./server/combat/*.proto
	./bin/protoc_server -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=bitbucket.org/funplus/server/proto --micro_opt=paths=source_relative ./server/mail/*.proto
	./bin/protoc_server -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=bitbucket.org/funplus/server/proto --micro_opt=paths=source_relative ./server/pubsub/*.proto

	# 客户端导出
	./bin/protoc_client -I=./ --csharp_out=:../ee_client/proto --csharp_opt=base_namespace=common ./global/common/*.proto
