.PHONY: proto
proto:
	# 服务器导出
	protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=github.com/east-eden/server/proto --micro_opt=paths=source_relative ./global/*.proto
	protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=github.com/east-eden/server/proto --micro_opt=paths=source_relative ./server/game/*.proto
	protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=github.com/east-eden/server/proto --micro_opt=paths=source_relative ./server/gate/*.proto
	protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=github.com/east-eden/server/proto --micro_opt=paths=source_relative ./server/combat/*.proto
	protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=github.com/east-eden/server/proto --micro_opt=paths=source_relative ./server/pubsub/*.proto

	# 客户端导出
	protoc -I=./ --csharp_out=:../ee_client/proto --csharp_opt=base_namespace=global ./global/*.proto

.PHONY: linux_proto
linux_proto:
	# 服务器导出
	./bin/protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=github.com/east-eden/server/proto --micro_opt=paths=source_relative ./global/*.proto
	./bin/protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=github.com/east-eden/server/proto --micro_opt=paths=source_relative ./server/game/*.proto
	./bin/protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=github.com/east-eden/server/proto --micro_opt=paths=source_relative ./server/gate/*.proto
	./bin/protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=github.com/east-eden/server/proto --micro_opt=paths=source_relative ./server/combat/*.proto
	./bin/protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=github.com/east-eden/server/proto --micro_opt=paths=source_relative ./server/pubsub/*.proto

	# 客户端导出
	protoc -I=./ --csharp_out=:../ee_client/proto --csharp_opt=base_namespace=global ./global/*.proto
