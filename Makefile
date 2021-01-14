.PHONY: proto
proto:
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go get github.com/micro/protoc-gen-micro/v2
	PROTOC_ZIP=protoc-3.14.0.-linux-x86_64.zip
	curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/$PROTOC_ZIP
	sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
	sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
	rm -f $PROTOC_ZIP

	protoc -I=./ --go_out=:./go_out --micro_out=:./go_out ./game/*.proto
	protoc -I=./ --go_out=:./go_out --micro_out=:./go_out ./gate/*.proto
	protoc -I=./ --go_out=:./go_out --micro_out=:./go_out ./combat/*.proto
	protoc -I=./ --go_out=:./go_out --micro_out=:./go_out ./account/*.proto
	protoc -I=./ --go_out=:./go_out --micro_out=:./go_out ./pubsub/*.proto

.PHONY: push_coding
push_coding:
	make -C apps/game push_coding
	make -C apps/gate push_coding
	make -C apps/combat push_coding
	make -C apps/client_bots push_coding

.PHONY: push_github
push_github:
	make -C apps/game push_github
	make -C apps/gate push_github
	make -C apps/combat push_github
	make -C apps/client_bots push_github
