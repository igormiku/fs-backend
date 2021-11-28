APP_NAME = ft_backend
VERSION = 1.0
COMPILE_NAME = ${APP_NAME}-${VERSION}
GOBIN=build/bin

all:
	go build -o ${GOBIN}/${COMPILE_NAME} main.go

clean:
	rm -fr $(GOBIN)/*