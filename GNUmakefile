VSN = 0.8

GOHOSTARCH = $(shell go env GOHOSTARCH)
GOHOSTOS = $(shell go env GOHOSTOS)

.PHONY: all clean

all:
	GOPATH=$(realpath .) go install slf2tcx tcx+gpx

clean:
	for f in slf2tcx tcx+gpx; do \
		rm -f bin/$$f; \
	done
	for f in gpx.a slf.a tcx.a; do \
		rm -f pkg/$(GOHOSTOS)_$(GOHOSTARCH)/$$f; \
	done
	-rmdir bin/
	-rmdir pkg/$(GOHOSTOS)_$(GOHOSTARCH) pkg/
