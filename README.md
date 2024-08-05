#How to config API documentation for go HTTP-server 

##Step 1:
Before starting need to check have you swaggo and install if not:
```bash
go get -u github.com/swaggo/swag/cmd/swag
```

##Step 2:
Configure API and endpoints with next comments:
```go
main.go:
// @title Example API
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /
```
```go
anyHandler:
// anyHandler godoc
// @Summary What does this endpoint
// @Description Here description endpoint
// @Tags Here tag endpoint
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Example answer for 200 code"
// @Router /endpoint [get]
```

##Step 3:
Download swagger files for correctly view swagger-panel:
```bash
curl -o swagger/docs/swagger-initializer.js https://raw.githubusercontent.com/swagger-api/swagger-ui/master/dist/swagger-initializer.js
curl -o swagger/docs/index.html https://raw.githubusercontent.com/swagger-api/swagger-ui/master/dist/index.html
curl -o swagger/docs/swagger-ui.css https://raw.githubusercontent.com/swagger-api/swagger-ui/master/dist/swagger-ui.css
curl -o swagger/docs/swagger-ui-bundle.js https://raw.githubusercontent.com/swagger-api/swagger-ui/master/dist/swagger-ui-bundle.js
curl -o swagger/docs/swagger-ui-standalone-preset.js https://raw.githubusercontent.com/swagger-api/swagger-ui/master/dist/swagger-ui-standalone-preset.js
```
Need to change url in file "../swagger/docs/swagger-initializer.js" from default to your url: "swagger.json"

##Step 4:
Generate docs for your API:
```bash
swag init -g main.go
```

##Step 5:
Relocate generated files:
```bash
mkdir -p swagger/docs
cp -r docs/* swagger/docs/
rm docs
```

##Step 6:
Start server:
```bash
go run main.go
```
For check documentation use path:
http://localhost:8080/swagger/
