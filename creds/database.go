package creds

import (
	"bytes"
	"encoding/json"
	"strings"

	"go.pedge.io/dlog"
	"github.com/portworx/kvdb"
)

const (
	// CredentialsDBKey is the key at which cloud credential is store in kvdb
	CredentialsDBKey = "credentials"
)

func readClusterInfo() (CredentialInfo, uint64, error) {
	kvdb := kvdb.Instance()

	db := CredentialInfo{
		Credentials: make(map[string]interface{},0),
	}

	kv, err := kvdb.Get(CredentialsDBKey)

	if err != nil && !strings.Contains(err.Error(), "Key not found") {
		dlog.Warnln("Warning, could not read credential database")
		return db, 0, err
	}

	if kv == nil || bytes.Compare(kv.Value, []byte("{}")) == 0 {
		dlog.Infoln("Credentials are uninitialized...")
		return db, 0, nil
	}
	if err := json.Unmarshal(kv.Value, &db); err != nil {
		dlog.Warnln("Fatal, Could not parse credential database ", kv)
		return db, 0, err
	}

	return db, kv.KVDBIndex, nil
}

func writeClusterInfo(db *CredentialInfo) (*kvdb.KVPair, error) {
	kvdb := kvdb.Instance()
	b, err := json.Marshal(db)
	if err != nil {
		dlog.Warnf("Fatal, Could not marshal cluster database to JSON: %v", err)
		return nil, err
	}

	kvp, err := kvdb.Put(CredentialsDBKey, b, 0)
	if err != nil {
		dlog.Warnf("Fatal, Could not marshal cluster database to JSON: %v", err)
		return nil, err
	}
	return kvp, nil
}
