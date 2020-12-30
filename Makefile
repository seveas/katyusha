# Let's not rebuild the parser if we don't have antlr available
ifeq ("", "$(strip $(shell which antlr))")
	antlr_sources :=
else
	antlr_sources := scripting/parser/katyusha_base_listener.go scripting/parser/katyusha_lexer.go scripting/parser/katyusha_listener.go scripting/parser/katyusha_parser.go
endif

# Let's not rebuild the protobuf code if we don't have protobuf available
ifeq ("", "$(strip $(shell which protoc))")
	protobuf_sources :=
else ifeq ("", "$(strip $(shell which protoc-gen-go))")
	protobuf_sources :=
else ifeq ("", "$(strip $(shell which protoc-gen-go-crpc))")
	protobuf_sources :=
else
	protobuf_sources = provider/plugin/common/plugin.pb.go provider/plugin/common/plugin_grpc.pb.go
endif

katyusha: go.mod *.go cmd/katyusha/*.go scripting/*.go provider/*/*.go provider/plugin/common/*.go $(protobuf_sources) $(antlr_sources)
	go build -o "$@" github.com/seveas/katyusha/cmd/katyusha

%_grpc.pb.go: %.proto
	protoc --go-grpc_out=. $^

%.pb.go: %.proto
	protoc --go_out=. $^

katyusha-provider-%: cmd/katyusha-provider-%/*.go provider/%/*.go provider/plugin/common/* provider/plugin/server/* $(protobuf_sources)
	go build -o "$@" github.com/seveas/katyusha/cmd/$@

ssh-agent-proxy: go.mod cmd/ssh-agent-proxy/*.go
	go build -o "$@" github.com/seveas/katyusha/cmd/ssh-agent-proxy

$(antlr_sources): scripting/Katyusha.g4
	(cd scripting; antlr -Dlanguage=Go -o parser Katyusha.g4)

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

provider/plugin/testdata/bin/katyusha-provider-ci: provider/plugin/testdata/provider/ci/*.go provider/plugin/testdata/cmd/katyusha-provider-ci/*.go provider/plugin/common/* provider/plugin/server/* $(protobuf_sources)
	go build -o "$@" github.com/seveas/katyusha/provider/plugin/testdata/cmd/katyusha-provider-ci

test: fmt vet tidy provider/plugin/testdata/bin/katyusha-provider-ci
	go test ./...

test-integration:
	go mod vendor
	make -C integration/pki
	docker-compose down || true
	docker-compose build
	docker-compose up --exit-code-from katyusha --abort-on-container-exit
	docker-compose down

dist_oses := darwin dragonfly freebsd linux netbsd openbsd windows
ssh_agent_oses := darwin dragonfly freebsd linux netbsd openbsd
build_all:
	@echo Building katyusha
	@$(foreach os,$(dist_oses),echo " - for $(os)" && mkdir -p dist/$(os)-amd64 && GOOS=$(os) GOARCH=amd64 go build -tags no_extra -ldflags '-s -w' -o dist/$(os)-amd64/ github.com/seveas/katyusha/cmd/katyusha;)
	@echo Building ssh-agent-proxy
	@$(foreach os,$(ssh_agent_oses),echo " - for $(os)" && GOOS=$(os) GOARCH=amd64 go build -ldflags '-s -w' -o dist/$(os)-amd64/ github.com/seveas/katyusha/cmd/ssh-agent-proxy;)

clean:
	rm -f katyusha
	rm -f katyusha-provider-example
	rm -f ssh-agent-proxy
	go mod tidy

fullclean: clean
	rm -rf dist/
	rm -f $(antlr_sources)
