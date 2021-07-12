# drone-sonar-dotnet-plugin
The plugin of Drone CI to integrate with SonarQube (previously called Sonar), which is an open source code quality management platform.

Detail tutorials: [DOCS.md](DOCS.md).

### Build process
build go binary file: 
`GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o drone-sonar`

build docker image
`docker build -t yahuiwong/drone-sonar-dotnet-plugin .`


### Testing the docker image:
```commandline
docker run --rm \
  -e DRONE_REPO=test \
  -e PLUGIN_SOURCES=. \
  -e PLUGIN_SONAR_HOST=http://sonar.domain.com \
  -e PLUGIN_SONAR_TOKEN=fa4353111bs53dc531260e9cgd1692d0329cf059 \
  yahuiwong/drone-sonar-dotnet-plugin
```

### Pipeline example
```yaml
steps
- name: code-analysis
  image: yahuiwong/drone-sonar-dotnet-plugin
  settings:
      sonar_host:
        from_secret: sonar_host
      sonar_token:
        from_secret: sonar_token
```
