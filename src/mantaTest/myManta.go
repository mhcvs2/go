package mantaTest

import (
	triton "github.com/joyent/triton-go"
	"github.com/joyent/triton-go/authentication"
	"github.com/joyent/triton-go/storage"
	"github.com/astaxie/beego"
	"path"
	"context"
	"io/ioutil"
	"strings"
	"fmt"
)

// genMantaClient get manta client
func genMantaClient() (*storage.StorageClient, error) {
	//get sdc config from consul
	var sdcKeyID, account string
	sdcKeyID = "d8:69:9d:07:60:0a:0f:2c:19:59:e0:5e:aa:c7:84:51"
	account = "dbba"
	privateKey := "-----BEGIN RSA PRIVATE KEY-----#MIIEowIBAAKCAQEArdxFHj3hv4Ii+cQBu5/K4670JapHts7aBaeNDFin8U58yIE+#d/3/zJd2I5+2m3217HE5FXJD3V+JyLJ4T1m/AY8l1II+2zNqmm/q+QmhBw8rFbJ1#O8y/HuC54LGkq/BGYT+NUtG4Gh6/3JAGHD+TAz4KFUp+pUV+t9VoKx2vPXqPD7xg#iRYCmgt/hIxX+f+tQ4PAIr9hNbsJrm6OMB+eoXV7whY3dC9gfsXrUEHO9NzdLD87#Ffjm1RckLiF8a2ulaKgWVmmUVqb7JcWDPhYMwEaLJGVOw5g3v9+xQrUbY5ObTIy7#C/IlGhAmd6P+hmwSz6OQoNa6KWqnvUekv7am7QIDAQABAoIBAFqS7DcrCdZZF5uC#71wtjOc8l9ifcyjbbl2Pwj1yWluuBff2zPJ6Eq8lINjCNcGfpgz9lz2C/7PuN7uk#rmS4XneTeaPSDqjnilvWflUrHQauckWlaMm0isStUmiqYx1n2WKEVz2UIBMLfeyL#44MH47DGuz4IRx4WrENdHB2KI2ck/DmPG0uPP+eQuI5G4VvNp9bckilHwS76b0RX#cHjF55DxKlP4TPjPDL3k9jRtAw3UxSMsks2xrRtIg0oiF8vEcCXrROaRPgk4FcJf#wuKEchgfOoONcGB1R/7Hye/EDYOjWDbW2O3uZVT/YenLhTXMj3vKMNFN29NErRAb#bLEV1gECgYEA2q4VeMDIo6nPD1GSra8oLFJBpNTpt0Hjqeni9aY2+0CMHi7vChXG#Izrj9RxhRJqKS2F+6p22r1DhMu19Xb0iAZkAZW7ryNcuyaUNsmLsMQPs4VUyhVAQ#k6yMu4Pwb/AFWsxHvOrvS/T+5iP+TBXIzv0YPrxuv5QsGj4/Gq1gdi0CgYEAy4gP#7qz5nwLToQTtGeOOMKJJAyGcijcOnkQc0+/hSZrhogn2kUe06dlpc5r7V7B6G3FF#Z9yf5F7QJNl8poCbhQ0g+V35qLA0ywudybsXd+q1Y+gAoAwtvD+246IUooQXDF93#Abjh87zS0LdcdozglvlETvEWfv2sj/MhVkk0K8ECgYB2MNKgEioe8t9bmy4Yu2uO#EOMz0HOFPZJrumKVfEGJKIjSo3FE1SHi1qhwSOd1acVHGqm66oTbWm5s1RkF+fwQ#Ov6Q1BOR2GOMTq4JdRfNIh78ZszIas6a0g66JoRkK6jpOzGmtJ+jQQYnotqFitye#qwJYngWJe+8eO/hlVcGl/QKBgQC+I+GGfzBAPdrJXZnHis+mXaXJ+BePA/pzHnyz#/jDAm6HYyGgBtzSrFsIuDwZqGGMqyfomGnWBWpYnJssNna4scWRxsjpvPhZD7hk9#gbxd+fX1XKNg4Z/Ect1/8UZHwRDrLTA3eqoUEz37YKFP2zJhuIL5IL98aa5RWLi3#LHJBwQKBgG9PGKQISss7LShb/x/IIAHCWS1Z+mB+V3vG2HhgEQvXJIM1ogRBUCox#i36Eor4ERC5KQj298ERO8LXc9NeL5zDWEam+jhdh7iSaNksmvPRqDOwfu6YkmqTH#Jk+5yAXYQFKcYeHteCgQIt9qPNYXaN6Mbysbj6SFBpKYubKCbTQg#-----END RSA PRIVATE KEY-----#"
	privateKey = strings.Replace(privateKey, "#", "\n", -1)
	//generate signature
	sshKeySigner, err4 := authentication.NewPrivateKeySigner(sdcKeyID, []byte(privateKey), account)
	if err4 != nil {
		fmt.Errorf("authentication.NewPrivateKeySigner, err=", err4)
		return nil, err4
	}
	sdcURL := "https://109.105.30.64"
	mantaURL := "https://109.105.30.55"
	config := &triton.ClientConfig{
		TritonURL:   sdcURL,
		MantaURL:    mantaURL,
		AccountName: account,
		Signers:     []authentication.Signer{sshKeySigner},
	}

	s, err7 := storage.NewClient(config)
	if err7 != nil {
		beego.Error("storage.NewClient: err=", err7)
		return nil, err7
	}

	return s, nil
}

// mantaClient the variable of manta client
var mantaClient *Client

// Client defination of manta client
type Client struct {
	*storage.StorageClient
}

// NewMantaClient create manta client
func NewMantaClient() *Client {
	if mantaClient != nil {
		return mantaClient
	}

	mc, err := genMantaClient()
	if err != nil {
		beego.Error("NewMantaClient, err=", NewMantaClient)
		return nil
	}

	mantaClient = &Client{StorageClient:mc}

	return mantaClient
}

// ListDirectory list all file in the directory
func (c *Client) ListDirectory(p string) ([]string, error) {
	var result []string

	// get all file in the directory
	listInput := &storage.ListDirectoryInput{DirectoryName: p}
	out, err := c.Dir().List(context.Background(), listInput)
	if err != nil {
		beego.Error("ListDirectory, err=", err)
		return result, err
	}
	for _, entry := range out.Entries {
		result = append(result, entry.Name)
	}

	return result, nil
}

// DeleteObject delete one object
func (c *Client) DeleteObject(p, object string) {
	op := path.Join(p, object)
	beego.Debug("DeleteObject op=", op)
	input := &storage.DeleteObjectInput{ObjectPath: op}
	if err := c.Objects().Delete(context.Background(), input); err != nil {
		beego.Error("DeleteObject, err=", err)
	}
}

