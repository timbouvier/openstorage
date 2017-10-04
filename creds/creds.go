package creds

import (
	"errors"
	"github.com/portworx/kvdb"
	"container/list"
)

var (
	inst *CredentialManager

	errCredsInitialized    = errors.New("openstorage.creds: already initialized")
	errCredsNotInitialized = errors.New("openstorage.creds: not initialized")
)

const (
	APIVersion = "v1"
)

type CredentialEntry struct {
	NonSecretsMap map[string]interface{}
	SecretsMap map[string]interface{}
}

type CredentialInfo struct {
	Credentials map[string]interface{}
}

// ClusterData interface provides apis to handle data of the cluster
type CredentialData interface {
	Create(CredentialEntry) error

	List() (map[string]interface{}, error)

	Delete(string) error

	Update(string,CredentialEntry) error
}

type Credential interface {
	CredentialData
}

func Init() error {
	if inst != nil {
		return errCredsInitialized
	}

	kv := kvdb.Instance()

	if kv == nil {
		return errors.New("KVDB is not yet initialized.  " +
			"A valid KVDB instance required for the cluster to start.")
	}

	inst = &CredentialManager{
		listeners:    list.New(),
		kv:           kv,
	}

	return nil
}

func Inst() (Credential, error) {

	if inst == nil {
		return nil, errCredsNotInitialized
	}

	return inst, nil
}

type NullCredentialListener struct {

}
