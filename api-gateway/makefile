mig-run:
	migrate create -ext sql -dir migrations -seq create_table

mig-up:
	migrate -database 'postgres://postgres:root@localhost:5432/db_name?sslmode=disable' -path migrations up

mig-down:
	migrate -database 'postgres://postgres:root@localhost:5432/db_name?sslmode=disable' -path migrations down

gen-proto:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	export PATH="$PATH:$(go env GOPATH)/bin"
	protoc --go_out=. --go-grpc_out=. *.proto
	protoc --go_out=. --go-grpc_out=. *.proto
	
gen-protoAll:
#proto fileni hammasini bittada generatsiya qilish 
	protoc --go_out=./ \
	--go-grpc_out=./ \
	submodule/user/*.proto

	protoc --go_out=./ \
	--go-grpc_out=./ \
	submodule/reservation/*.proto

	protoc --go_out=./ \
	--go-grpc_out=./ \
	submodule/payment/*.proto


	