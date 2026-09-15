package mpsl

import "github.com/koykov/byteconv"

// GetTLDString returns TLD part of string hostname and ICANN flag.
func (db *DB) GetTLDString(hostname string) (tld string, icann bool) {
	var btld []byte
	_, _, btld, icann = db.Parse(byteconv.S2B(hostname))
	tld = byteconv.B2S(btld)
	return
}

// GetEffectiveTLDString returns only eTLD part of string hostname.
func (db *DB) GetEffectiveTLDString(hostname string) (etld string) {
	_, betld, _, _ := db.Parse(byteconv.S2B(hostname))
	etld = byteconv.B2S(betld)
	return
}

// GetEffectiveTLDPlusOneString return only eTLD1 part of string hostname.
func (db *DB) GetEffectiveTLDPlusOneString(hostname string) (etld1 string) {
	_, _, betld1, _ := db.Parse(byteconv.S2B(hostname))
	etld1 = byteconv.B2S(betld1)
	return
}

// GetETLDString is a shorthand alias of GetEffectiveTLDStr.
func (db *DB) GetETLDString(hostname string) string {
	return db.GetEffectiveTLDString(hostname)
}

// GetETLD1String is a shorthand alias of GetEffectiveTLDPlusOneStr.
func (db *DB) GetETLD1String(hostname string) string {
	return db.GetEffectiveTLDPlusOneString(hostname)
}
