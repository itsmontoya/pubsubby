package string

//go:generate genny -pkg=string -in=$GOPATH/src/github.com/itsmontoya/pubsubby/pubsubby.go -out=pubsubby.go gen "Key=string Value=string"

// New will return a new Pubsubby
func New() *Pubsubby {
	return Pubsubby(newPubsubby())
}

// Pubsubby is an exported pubsubby
type Pubsubby pubsubby
