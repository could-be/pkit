package templatevar

const (
	DockerComposeYamlTemplate = `
# docker-compose.yaml
version: '3.7'

services:

  {{.ProjectName}}:
    image: {{.DockerRegistry}}/{{.ProjectName}}:0.0.1
    container_name: {{.ProjectName}}
    restart: on-failure
    network_mode: nginx-internal-default
    command: ./{{.ProjectName}} -f config/{{.ProjectName}}.toml
    volumes:
    - ./config:/app/config:rw
    - ./data:/app/data:rw
    - ./report:/app/report:rw
    - /etc/localtime:/etc/localtime:ro
    - /etc/hosts:/etc/hosts:ro


`
)
