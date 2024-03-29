package mpsl

import "github.com/koykov/byteconv"

// GetTLDStr returns TLD part of string hostname and ICANN flag.
func (db *DB) GetTLDStr(hostname string) (tld string, icann bool) {
	var btld []byte
	_, _, btld, icann = db.Parse(byteconv.S2B(hostname))
	tld = byteconv.B2S(btld)
	return
}

// GetEffectiveTLDStr returns only eTLD part of string hostname.
func (db *DB) GetEffectiveTLDStr(hostname string) (etld string) {
	_, betld, _, _ := db.Parse(byteconv.S2B(hostname))
	etld = byteconv.B2S(betld)
	return
}

// GetEffectiveTLDPlusOneStr return only eTLD1 part of string hostname.
func (db *DB) GetEffectiveTLDPlusOneStr(hostname string) (etld1 string) {
	_, _, betld1, _ := db.Parse(byteconv.S2B(hostname))
	etld1 = byteconv.B2S(betld1)
	return
}

// GetETLDStr is a shorthand alias of GetEffectiveTLDStr.
func (db *DB) GetETLDStr(hostname string) string {
	return db.GetEffectiveTLDStr(hostname)
}

// GetETLD1Str is a shorthand alias of GetEffectiveTLDPlusOneStr.
func (db *DB) GetETLD1Str(hostname string) string {
	return db.GetEffectiveTLDPlusOneStr(hostname)
}
