package creds

import (
	"container/list"
	"github.com/portworx/kvdb"
	"go.pedge.io/dlog"
	"github.com/pborman/uuid"
	secrets "github.com/libopenstorage/secrets"
	"fmt"
)

const(
	credLockKey     = "/cred/lock"
)

type CredentialManager struct {
	listeners     *list.List
	kv            kvdb.Kvdb
}

func (cm *CredentialManager) List() (map[string]interface{}, error) {
	kvdb := kvdb.Instance()

	kvlock, err := kvdb.Lock(credLockKey)
	if err != nil {
		dlog.Warnln("Unable to obtain cred lock for creating a cloud credential", err)
		return nil, err
	}
	defer kvdb.Unlock(kvlock)

	db, _, err := readClusterInfo()
	if err != nil {
		return nil, err
	}

	return db.Credentials, nil
}

func (cm *CredentialManager) Create(cred CredentialEntry) error{
	kvdb := kvdb.Instance()
	kvlock, err := kvdb.Lock(credLockKey)
	if err != nil {
		dlog.Warnln("Unable to obtain cred lock for creating a cloud credential", err)
		return nil
	}
	defer kvdb.Unlock(kvlock)

	db, _, err := readClusterInfo()
	if err != nil {
		return err
	}
	uuid := uuid.New()

	secret_inst := secrets.Instance()
	err = secret_inst.PutSecret(uuid, cred.SecretsMap, nil)
	if err != nil {
		return fmt.Errorf("Unable to set cluster secret key. Check" +
			" credentials")
	}

	db.Credentials[uuid] = cred.NonSecretsMap

	_, err = writeClusterInfo(&db)

	if err != nil {
		return err
	}

	return nil
}

func (cm *CredentialManager) Update(uuid string, cred CredentialEntry) error {
	kvdb := kvdb.Instance()
	kvlock, err := kvdb.Lock(credLockKey)
	if err != nil {
		dlog.Warnln("Unable to obtain cred lock for creating a cloud credential", err)
		return nil
	}
	defer kvdb.Unlock(kvlock)

	db, _, err := readClusterInfo()
	if err != nil {
		return err
	}

	db.Credentials[uuid] = cred.NonSecretsMap

	_, err = writeClusterInfo(&db)
	if err != nil {
		return err
	}
	return nil
}

func (cm *CredentialManager) Delete(uuid string) error {
	kvdb := kvdb.Instance()
	kvlock, err := kvdb.Lock(credLockKey)
	if err != nil {
		dlog.Warnln("Unable to obtain cred lock for creating a cloud credential", err)
		return nil
	}
	defer kvdb.Unlock(kvlock)

	db, _, err := readClusterInfo()
	if err != nil {
		return err
	}

	delete(db.Credentials, uuid)

	_, err = writeClusterInfo(&db)
	if err != nil {
		return err
	}
	return nil
}