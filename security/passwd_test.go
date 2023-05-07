package security

import (
	"log"
	"testing"
)

func TestHashPasswd(t *testing.T) {
	hash := HashPasswd([]byte("123456"))
	log.Println(hash)
}
