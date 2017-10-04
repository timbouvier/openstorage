package creds

import (
	"github.com/libopenstorage/openstorage/api/client"
	"io/ioutil"
	"bytes"
	"strings"
	"errors"
	"github.com/libopenstorage/openstorage/creds"
)

const(
	credPath  	          = "/creds"
	AddCreds                  = "/addcreds"
	DeleteCreds               = "/deletecreds"
	GetCreds                  = "/getcreds"
	ListCreds                 = "/listcreds"
	ValidateCreds             = "/validatecreds"
)

type credClient struct {
	c *client.Client
}

func newCredClient(c *client.Client) credClient {
	return credClient{c}
}

// String description of this driver.
func (c *credClient) Name() string {
	return "CredentialClient"
}

func (c *credClient) List() ([]creds.CredentialEntry, error) {
	creds := make([]creds.CredentialEntry, 0)

	if err := c.c.Get().Resource(credPath + ListCreds).Do().Unmarshal(&creds); err != nil {
		return nil, err
	}

	return creds, nil
}

func (c *credClient) Create(cred creds.CredentialEntry) error{
	response := c.c.Post().Resource(credPath + AddCreds).Body(cred).Do()

	if response.Error() != nil {
		return response.Error()
	}

	return nil
}

func (c *credClient) Delete(uuid string) error {

	body, err := c.c.Get().Resource(credPath + DeleteCreds +"?uuid=" + uuid).Do().Body();
	if err != nil {
		return err
	}

	response := bytes.NewBuffer(body)

	contents, err := ioutil.ReadAll(response)
	if err != nil {
		return err
	}

	return errors.New(strings.TrimSpace(string(contents)))
}

func (c *credClient) Update(credID string, cred creds.CredentialEntry) error {
	response := c.c.Put().Resource(credPath + AddCreds).Instance(credID).Body(cred).Do()

	if response.Error() != nil {
		return response.Error()
	}

	return nil
}