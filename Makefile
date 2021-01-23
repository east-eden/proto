.PHONY: proto
proto:
	protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=e.coding.net/mmstudio/blade/proto/go_out --micro_opt=paths=source_relative ./game/*.proto
	protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=e.coding.net/mmstudio/blade/proto/go_out --micro_opt=paths=source_relative ./gate/*.proto
	protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=e.coding.net/mmstudio/blade/proto/go_out --micro_opt=paths=source_relative ./combat/*.proto
	protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=e.coding.net/mmstudio/blade/proto/go_out --micro_opt=paths=source_relative ./account/*.proto
	protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=e.coding.net/mmstudio/blade/proto/go_out --micro_opt=paths=source_relative ./pubsub/*.proto

.PHONY: linux_proto
linux_proto:
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go get github.com/micro/protoc-gen-micro/v2
	
	./bin/protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=e.coding.net/mmstudio/blade/proto/go_out --micro_opt=paths=source_relative ./game/*.proto
	./bin/protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=e.coding.net/mmstudio/blade/proto/go_out --micro_opt=paths=source_relative ./gate/*.proto
	./bin/protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=e.coding.net/mmstudio/blade/proto/go_out --micro_opt=paths=source_relative ./combat/*.proto
	./bin/protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=e.coding.net/mmstudio/blade/proto/go_out --micro_opt=paths=source_relative ./account/*.proto
	./bin/protoc -I=./ --go_out=:../server/proto --micro_out=:../server/proto --go_opt=module=e.coding.net/mmstudio/blade/proto/go_out --micro_opt=paths=source_relative ./pubsub/*.proto
