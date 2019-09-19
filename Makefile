default:
	go mod vendor \
        && go build -v -mod=vendor -o ./go-link-check \
        && ./go-link-check
