package mpsl

// ParseStr is a legacy version of ParseString.
// Deprecated: use ParseString instead.
func (db *DB) ParseStr(hostname string) (tld, etld, etld1 string, icann bool) {
	return db.ParseString(hostname)
}
