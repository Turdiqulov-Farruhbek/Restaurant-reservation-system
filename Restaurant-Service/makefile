mig-run:
	migrate create -ext sql -dir migrations -seq reservation 

mig-up:
	migrate -database 'postgres://postgres:root@localhost:5432/reservation_db?sslmode=disable' -path migrations up

mig-down:
	migrate -database 'postgres://postgres:root@localhost:5432/reservation_db?sslmode=disable' -path migrations down

prot-exp:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	export PATH="$PATH:$(go env GOPATH)/bin"

gen-proto-user:
	protoc --go_out=genproto/ \
    --go-grpc_out=genproto/ \
	submodule/user/*.proto

gen-proto-reservation:
	protoc --go_out=genproto/ \
    --go-grpc_out=genproto/ \
	submodule/reservation/*.proto

gen-proto-payment:
	protoc --go_out=genproto/ \
    --go-grpc_out=genproto/ \
	submodule/payment/*.proto