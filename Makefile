NAMESPACE  := logicmonitor
REPOSITORY := collectorset-controller-proto
VERSION    := 1.0.0


default: build
build:
	docker build --build-arg VERSION=$(VERSION) -t $(NAMESPACE)/$(REPOSITORY):latest .
	docker run --rm -v "$(shell pwd)":/out --entrypoint=cp $(NAMESPACE)/$(REPOSITORY):latest /tmp/api.pb.go /out/api
