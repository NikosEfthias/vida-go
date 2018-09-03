GO_FILES=$(shell find . -name "*.go")

vida.linux:main.go ${GO_FILES}
	GOOS=linux go build -ldflags="-w -s" -o $@ $<
.Phony:deploy
deploy:vida.linux
	rsync -vaurz --progress --remove-source-files $< vida:/usr/local/bin/
