package passman

import "testing"

/*****************************************************************************
Rough Specification:

1. passman should be able to take a passstring generated by passgen
2. encrypt the password string using AES (or bcrypt) and return a valid hash
3. store hash in a serializable datastructure in a way that is storage agnostic
so it can be stored locally, or in hosted in a cloud environment
4. allow use of oauth2 for user authentication

*****************************************************************************/
var tmppass = "fdsFDF43$"

func TestBuildRecord(t *testing.T) {

}
