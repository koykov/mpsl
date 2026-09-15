package mpsl

// ParseStr is a legacy version of ParseString.
// Deprecated: use ParseString instead.
func (db *DB) ParseStr(hostname string) (tld, etld, etld1 string, icann bool) {
	return db.ParseString(hostname)
}

// GetTLDStr is a legacy version of GetTLDString.
// Deprecated: use GetTLDString instead.
func (db *DB) GetTLDStr(hostname string) (tld string, icann bool) {
	return db.GetTLDString(hostname)
}

// GetEffectiveTLDStr is a legacy version of GetEffectiveTLDString.
// Deprecated: use GetEffectiveTLDString instead.
func (db *DB) GetEffectiveTLDStr(hostname string) (etld string) {
	return db.GetEffectiveTLDString(hostname)
}

// GetEffectiveTLDPlusOneStr is a legacy version of GetEffectiveTLDPlusOneString.
// Deprecated: use GetEffectiveTLDPlusOneString instead.
func (db *DB) GetEffectiveTLDPlusOneStr(hostname string) (etld1 string) {
	return db.GetEffectiveTLDPlusOneString(hostname)
}

// GetETLDStr is a legacy version of GetETLDString.
// Deprecated: use GetETLDString instead.
func (db *DB) GetETLDStr(hostname string) string {
	return db.GetETLDString(hostname)
}

// GetETLD1Str is a legacy version of GetETLD1String.
// Deprecated: use GetETLD1String instead.
func (db *DB) GetETLD1Str(hostname string) string {
	return db.GetETLD1String(hostname)
}
