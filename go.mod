module github.com/tendermint/iavl

go 1.12

require (
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.4.0
	github.com/tendermint/go-amino v0.14.1
	github.com/tendermint/tendermint v0.32.6
	github.com/tendermint/tm-db v0.2.0
	golang.org/x/crypto v0.0.0-20190313024323-a1f597ede03a
)

replace github.com/tendermint/go-amino v0.14.1 => ../go-amino-v0150
