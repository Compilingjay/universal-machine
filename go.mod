module github.com/Compilingjay/universal-machine

go 1.23.0

require um v0.0.0

replace (
	um v0.0.0 => ./internal/um
)