gen:
	protoc -I=proto/src --go_out=. proto/src/*.proto && rm -rf proto/gen && mv github.com/hindungWang/oracle-observer-go/proto/gen proto/ 