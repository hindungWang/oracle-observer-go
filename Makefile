gen:
	protoc -I=proto/src --go_out=. proto/src/*.proto && rm -rf proto/gen && mv github.com/hindungWang/oracle-observer-go/proto/gen proto/ 

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o hello main.go
zip:
	rm -f main.zip && zip -r main.zip . -i  hello 