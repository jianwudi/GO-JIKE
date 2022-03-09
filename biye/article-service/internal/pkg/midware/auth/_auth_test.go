package auth

import (
	"fmt"
	"testing"
)

func TestGeneralToken(t *testing.T) {
	str := GenerateToken()
	fmt.Println(str)
}
