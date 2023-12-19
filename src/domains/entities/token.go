package entities

import "time"

type Fingerprint struct {
	Fingerprint int64  `db:"fingerprint" json:"fingerprint"`
	Os          string `db:"os" json:"os"`
	UserAgent   string `db:"user_agent" json:"user_agent"`
}

type RefreshToken struct {
	// TODO: avoid duplicated code
	Fingerprint int64     `db:"fingerprint" json:"fingerprint"`
	Os          string    `db:"os" json:"os"`
	UserAgent   string    `db:"user_agent" json:"user_agent"`
	Id          int64     `db:"id" json:"id"`
	AccountId   int64     `db:"account_id" json:"account_id"`
	Expires     time.Time `db:"expires" json:"expires"`
	Token       string    `db:"token" json:"token"`
}

func (r *RefreshToken) GetFieldsMap() map[string]interface{} {
	return map[string]interface{}{
		"account_id":  r.AccountId,
		"fingerprint": r.Fingerprint,
		"os":          r.Os,
		"user_agent":  r.UserAgent,
		"expires":     r.Expires,
		"token":       r.Token,
	}
}

func HaveSameFingerprints(first, second *RefreshToken) bool {
	return first.Fingerprint == second.Fingerprint && first.Os == second.Os &&
		first.UserAgent == second.UserAgent && first.AccountId == second.AccountId
}
