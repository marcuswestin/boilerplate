package db

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"github.com/golang/protobuf/proto"
)

func Marshal(pb proto.Message) []byte {
	bytes, err := proto.Marshal(pb)
	if err != nil {
		panic(err)
	}
	return bytes
}

func CreateAccountInvite(accountId int64) string {
	return createAccountSecret(accountId, 6)
}
func CreateAccountLogin(accountId int64) string {
	return createAccountSecret(accountId, 8)
}

func LookupAccountSecret(secret string) (accountId int64, found bool) {
	found = SelectOneMaybe(&accountId, `
		SELECT account_id FROM account_secret
		WHERE secret=? AND active=true`, secret)
	return
}

func LookupAccountInvite(secret string) (*Account, string) {
	accountId, found := LookupAccountSecret(secret)
	if !found {
		wasUsed := SelectOneMaybe(&accountId, `
			SELECT account_id FROM account_secret
			WHERE secret=? AND active=false`, secret)
		if wasUsed {
			return nil, "You have already enrolled"
		} else {
			return nil, "We can't find this invite"
		}
	}

	var account Account
	SelectOne(&account, `
		SELECT * FROM account
		WHERE account_id=?`, accountId)
	return &account, ""
}

// Internal
///////////

func createAccountSecret(accountId int64, numBytes int) string {
	for i := 0; i < 10; i++ {
		secret := genSecret(numBytes)
		duplicate := InsertIgnoreDuplicate(`
			INSERT INTO account_secret
			(account_id, secret, active) VALUES (?, ?, ?)`,
			accountId, secret, true)
		if !duplicate {
			return secret
		}
		if i == 5 {
			log.Println("WARNING: Could not generate secret after 5 tries")
		}
	}
	panic("Could not generate invite secret!")
}
func genSecret(numBytes int) string {
	b := make([]byte, numBytes)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}
