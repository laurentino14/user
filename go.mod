module github.com/laurentino14/user

go 1.19

replace github.com/laurentino14/user/prisma => ./prisma

require (
	github.com/99designs/gqlgen v0.17.20
	github.com/disintegration/imaging v1.6.2
	github.com/gobuffalo/envy v1.10.2
	github.com/golang-jwt/jwt/v4 v4.4.2
	github.com/iancoleman/strcase v0.0.0-20190422225806-e506e3ef7365
	github.com/joho/godotenv v1.4.0
	github.com/prisma/prisma-client-go v0.16.2
	github.com/rs/cors v1.8.2
	github.com/shopspring/decimal v1.3.1
	github.com/takuoki/gocase v1.0.0
	github.com/vektah/gqlparser/v2 v2.5.1
)

require (
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	golang.org/x/image v0.0.0-20191009234506-e7c1f5e7dbb8 // indirect
)
