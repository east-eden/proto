.PHONY: proto
proto:
	# 服务器导出
	protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=bitbucket.org/east-eden/server/proto --micro_opt=paths=source_relative ./game/*.proto
	protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=bitbucket.org/east-eden/server/proto --micro_opt=paths=source_relative ./gate/*.proto
	protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=bitbucket.org/east-eden/server/proto --micro_opt=paths=source_relative ./combat/*.proto
	protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=bitbucket.org/east-eden/server/proto --micro_opt=paths=source_relative ./account/*.proto
	protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=bitbucket.org/east-eden/server/proto --micro_opt=paths=source_relative ./pubsub/*.proto

	# 客户端导出
	protoc -I=./ --csharp_out=:../client/proto --csharp_opt=base_namespace=game ./game/*.proto

.PHONY: linux_proto
linux_proto:
	
	# 服务器导出
	./bin/protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=bitbucket.org/east-eden/server/proto --micro_opt=paths=source_relative ./game/*.proto
	./bin/protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=bitbucket.org/east-eden/server/proto --micro_opt=paths=source_relative ./gate/*.proto
	./bin/protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=bitbucket.org/east-eden/server/proto --micro_opt=paths=source_relative ./combat/*.proto
	./bin/protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=bitbucket.org/east-eden/server/proto --micro_opt=paths=source_relative ./account/*.proto
	./bin/protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=bitbucket.org/east-eden/server/proto --micro_opt=paths=source_relative ./pubsub/*.proto

	# 客户端导出
	protoc -I=./ --csharp_out=:../client/proto --csharp_opt=base_namespace=game ./game/*.proto
