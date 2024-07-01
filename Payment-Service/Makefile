mig-run:
	migrate create -ext sql -dir migrations -seq payment 

mig-up:
	migrate -database 'postgres://postgres:root@localhost:5432/payment_db?sslmode=disable' -path migrations up

mig-down:
	migrate -database 'postgres://postgres:root@localhost:5432/payment_db?sslmode=disable' -path migrations down

mig-fix:
	migrate -database 'postgres://postgres:root@localhost:5432/payment_db?sslmode=disable' -path migrations force 1

prot-exp:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	export PATH="$PATH:$(go env GOPATH)/bin"

gen-proto-payment:
	protoc --go_out=genproto/ \
    --go-grpc_out=genproto/ \
	submodule/payment/*.proto