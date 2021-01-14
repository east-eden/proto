.PHONY: proto
proto:
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go get github.com/micro/protoc-gen-micro/v2
	
	pwd
	ls -al
	./bin/protoc -I=./ --go_out=:./go_out --micro_out=:./go_out ./game/*.proto
	./bin/protoc -I=./ --go_out=:./go_out --micro_out=:./go_out ./gate/*.proto
	./bin/protoc -I=./ --go_out=:./go_out --micro_out=:./go_out ./combat/*.proto
	./bin/protoc -I=./ --go_out=:./go_out --micro_out=:./go_out ./account/*.proto
	./bin/protoc -I=./ --go_out=:./go_out --micro_out=:./go_out ./pubsub/*.proto

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
