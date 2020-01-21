package templatevar

const (
	DockerfileTemplate = `
FROM alpine:3

RUN apk update \
&& apk add ca-certificates tzdata\
&& rm -rf /var/cache/apk/*

RUN ln -fs /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY {{.ProjectName}} config.toml /opt/
WORKDIR /opt

ENTRYPOINT ["/opt/{{.ProjectName}}"]

`
)
