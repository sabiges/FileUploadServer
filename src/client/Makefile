include ../Makefile.include


TARGET=store
COVERAGE_TARGET=store_test


install:
	@echo ""
	@echo ${GOROOT}
	@echo ${GOPATH}
	@echo "... Generate version"
	@echo ".... Build and install " ${TARGET}
	${GOROOT}/bin/go build -ldflags="-s -w" -o ${TARGET} *.go
	cp ${TOPROOT}/src/client/${TARGET} ${TOPROOT}/bin/
clean:
	@echo ""
	@echo ".... Cleaning " ${TARGET}
	@echo ".."
	#${GOROOT}/bin/go clean -x client/${TARGET}
	rm -f ${TARGET} bin/${TARGET}
	rm -f bin/${COVERAGE_TARGET}
	@echo ".... Clean Done"
test:
	@echo ""
	@echo ".... Creating binary for Coverage"
	rm -f bin/${COVERAGE_TARGET}
	${GOROOT}/bin/go test -c -v -race -coverpkg="" -o bin/${COVERAGE_TARGET} 
	@echo ".... Done"
