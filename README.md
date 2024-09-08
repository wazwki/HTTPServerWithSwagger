# **How to Set Up API Documentation in a Golang HTTP Server**

# 1. **Before You Start, Check if You Have Swaggo and Install It If Not:**

```bash
go get -u github.com/swaggo/swag/cmd/swag
```

# 2. **Set Up API and Endpoints with the Following Comments:**

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

# 3. **Download Swagger Files for Correct Swagger UI Viewing:**

```bash
curl -o swagger/docs/swagger-initializer.js https://raw.githubusercontent.com/swagger-api/swagger-ui/master/dist/swagger-initializer.js
curl -o swagger/docs/index.html https://raw.githubusercontent.com/swagger-api/swagger-ui/master/dist/index.html
curl -o swagger/docs/swagger-ui.css https://raw.githubusercontent.com/swagger-api/swagger-ui/master/dist/swagger-ui.css
curl -o swagger/docs/swagger-ui-bundle.js https://raw.githubusercontent.com/swagger-api/swagger-ui/master/dist/swagger-ui-bundle.js
curl -o swagger/docs/swagger-ui-standalone-preset.js https://raw.githubusercontent.com/swagger-api/swagger-ui/master/dist/swagger-ui-standalone-preset.js
```

***You need to change the URL in the file "../swagger/docs/swagger-initializer.js" from the default value to your URL: "swagger.json"***

# 4. **Generate Documentation for Your API:**

```bash
swag init -g main.go
```

# 5. **Move the Generated Files:**

```bash
mkdir -p swagger/docs
cp -r docs/* swagger/docs/
rm -r docs
```

# 6. **Start the Server:**

```bash
go run main.go
```

***To check the documentation, use the path:*** http://localhost:8080/swagger/