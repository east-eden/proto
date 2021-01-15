.PHONY: proto
proto:
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go get github.com/micro/protoc-gen-micro/v2
	
	./bin/protoc -I=./ --go_out=:./go_out --micro_out=:./go_out --go_opt=module=e.coding.net/mmstudio/blade/proto/go_out --micro_opt=paths=source_relative ./game/*.proto
	./bin/protoc -I=./ --go_out=:./go_out --micro_out=:./go_out --go_opt=module=e.coding.net/mmstudio/blade/proto/go_out --micro_opt=paths=source_relative ./gate/*.proto
	./bin/protoc -I=./ --go_out=:./go_out --micro_out=:./go_out --go_opt=module=e.coding.net/mmstudio/blade/proto/go_out --micro_opt=paths=source_relative ./combat/*.proto
	./bin/protoc -I=./ --go_out=:./go_out --micro_out=:./go_out --go_opt=module=e.coding.net/mmstudio/blade/proto/go_out --micro_opt=paths=source_relative ./account/*.proto
	./bin/protoc -I=./ --go_out=:./go_out --micro_out=:./go_out --go_opt=module=e.coding.net/mmstudio/blade/proto/go_out --micro_opt=paths=source_relative ./pubsub/*.proto
