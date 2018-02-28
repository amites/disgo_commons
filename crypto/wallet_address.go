package crypto

import (
	"github.com/dispatchlabs/disgo_commons/constants"
	"math/rand"
	"encoding/hex"
	"io/ioutil"
	"fmt"
	"os"
	"log"
	"path/filepath"
	"os/user"
)

// NewWalletAddress
func NewWalletAddress() ([constants.AddressLength]byte, error) {

	// TODO: How do we generate the private key?
	privateKey := make([]byte, constants.HashLength)
	rand.Read(privateKey)

	// Create address.
	hash := Sum256(privateKey)
	address := [constants.AddressLength]byte{}
	for i := 0; i < constants.AddressLength; i++ {
		address[i] = hash[i+12]
	}

	return address, nil
}

func GetWalletAddress() (string, error) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
	}

	disgoDir := usr.HomeDir + "/.disgo"
	disgoAccountFile := disgoDir + "/disgo_account"
	fmt.Println( usr.HomeDir )
	fmt.Println( disgoDir )
	fmt.Println( disgoAccountFile )

	var result string
	addrBytes, err := ioutil.ReadFile(disgoAccountFile)

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	if err != nil || len(addrBytes) == 0 {
		if _, err := os.Stat(disgoDir); os.IsNotExist(err) {
			err = os.MkdirAll(disgoDir, 0755)
			if err != nil {
				panic(err)
			}
		}

		file, err := os.Create(disgoAccountFile)
		if err != nil {
			log.Fatal("Cannot create file", err)
		}
		defer file.Close()
		wa, err := NewWalletAddress()
		if err != nil {
			log.Fatal("Cannot create wallet", err)
		}

		result = ToWalletAddressString(wa)
		fmt.Fprintf(file, result)
	}
	if addrBytes != nil {
		result = string(addrBytes)
	}
	fmt.Printf(result)
	return result, nil
}

// ToWalletAddress
func ToWalletAddress(bytes []byte) [constants.AddressLength]byte {
	address := [constants.AddressLength]byte{}
	copy(address[:], bytes)
	return address
}

// ToWalletAddressString
func ToWalletAddressString(address [constants.AddressLength]byte) string {
	return hex.EncodeToString(address[:])
}

func AddressStringToBytes(address string) ([]byte, error) {
	result, err := hex.DecodeString(address)
	return result, err
}

func ToWalletAddressBytes(address [constants.AddressLength]byte) []byte {
	byteArray := make([]byte, constants.AddressLength)
	copy(byteArray, address[:])
	return byteArray
}