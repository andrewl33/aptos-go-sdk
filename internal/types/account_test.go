package types

import (
	"github.com/aptos-labs/aptos-go-sdk/bcs"
	"github.com/aptos-labs/aptos-go-sdk/crypto"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	defaultMetadata = "0x2ebb2ccac5e027a87fa0e2e5f656a3a4238d6a48d93ec9b610d570fc0aa0df12"
	defaultStore    = "0x8a9d57692a9d4deb1680eaf107b83c152436e10f7bb521143fa403fa95ef76a"
	defaultOwner    = "0xc67545d6f3d36ed01efc9b28cbfd0c1ae326d5d262dd077a29539bcee0edce9e"
)

func TestAccountSpecialString(t *testing.T) {
	var aa AccountAddress
	aa[31] = 3
	aas := aa.String()
	if aas != "0x3" {
		t.Errorf("wanted 0x3 got %s", aas)
	}

	var aa2 AccountAddress
	err := aa2.ParseStringRelaxed("0x3")
	if err != nil {
		t.Errorf("unexpected err %s", err)
	}
	if aa2 != aa {
		t.Errorf("aa2 != aa")
	}
}

func TestSpecialAddresses(t *testing.T) {
	var addr AccountAddress
	err := addr.ParseStringRelaxed("0x0")
	assert.NoError(t, err)
	assert.Equal(t, AccountZero, addr)
	err = addr.ParseStringRelaxed("0x1")
	assert.NoError(t, err)
	assert.Equal(t, AccountOne, addr)
	err = addr.ParseStringRelaxed("0x2")
	assert.NoError(t, err)
	assert.Equal(t, AccountTwo, addr)
	err = addr.ParseStringRelaxed("0x3")
	assert.NoError(t, err)
	assert.Equal(t, AccountThree, addr)
	err = addr.ParseStringRelaxed("0x4")
	assert.NoError(t, err)
	assert.Equal(t, AccountFour, addr)
}

func TestSerialize(t *testing.T) {
	inputs := [][]byte{
		{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
		{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0F},
		{0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
		{0x02, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
		{0x00, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
		{0x00, 0x04, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
		{0x00, 0x00, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
	}

	for i := 0; i < len(inputs); i++ {
		addr := AccountAddress(inputs[i])
		bytes, err := bcs.Serialize(&addr)
		assert.NoError(t, err)
		assert.Equal(t, bytes, inputs[i])

		newAddr := AccountAddress{}
		err = bcs.Deserialize(&newAddr, bytes)
		assert.NoError(t, err)
		assert.Equal(t, addr, newAddr)
	}
}

func TestStringOutput(t *testing.T) {
	inputs := [][]byte{
		{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
		{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0F},
		{0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
		{0x02, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
		{0x00, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
		{0x00, 0x04, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
		{0x00, 0x00, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x12, 0x34, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
	}
	expected := []string{
		"0x0",
		"0x1",
		"0xf",
		"0x1234123412341234123412341234123412341234123412340123456789abcdef",
		"0x0234123412341234123412341234123412341234123412340123456789abcdef",
		"0x0034123412341234123412341234123412341234123412340123456789abcdef",
		"0x0004123412341234123412341234123412341234123412340123456789abcdef",
		"0x0000123412341234123412341234123412341234123412340123456789abcdef",
	}
	expectedLong := []string{
		"0x0000000000000000000000000000000000000000000000000000000000000000",
		"0x0000000000000000000000000000000000000000000000000000000000000001",
		"0x000000000000000000000000000000000000000000000000000000000000000f",
		"0x1234123412341234123412341234123412341234123412340123456789abcdef",
		"0x0234123412341234123412341234123412341234123412340123456789abcdef",
		"0x0034123412341234123412341234123412341234123412340123456789abcdef",
		"0x0004123412341234123412341234123412341234123412340123456789abcdef",
		"0x0000123412341234123412341234123412341234123412340123456789abcdef",
	}
	for i := 0; i < len(inputs); i++ {
		addr := AccountAddress(inputs[i])
		assert.Equal(t, expected[i], addr.String())
		assert.Equal(t, expectedLong[i], addr.StringLong())
	}
}

func TestAccountAddress_ParseStringRelaxed_Error(t *testing.T) {
	var owner AccountAddress
	err := owner.ParseStringRelaxed("0x")
	assert.Error(t, err)
	err = owner.ParseStringRelaxed("0xF1234567812345678123456781234567812345678123456781234567812345678")
	assert.Error(t, err)
	err = owner.ParseStringRelaxed("NotHex")
	assert.Error(t, err)
}
func TestAccountAddress_ObjectAddressFromObject(t *testing.T) {
	var owner AccountAddress
	err := owner.ParseStringRelaxed(defaultOwner)
	assert.NoError(t, err)

	var objectAddress AccountAddress
	err = objectAddress.ParseStringRelaxed(defaultMetadata)
	assert.NoError(t, err)

	var expectedDerivedAddress AccountAddress
	err = expectedDerivedAddress.ParseStringRelaxed(defaultStore)
	assert.NoError(t, err)

	derivedAddress := owner.ObjectAddressFromObject(&objectAddress)
	assert.NoError(t, err)

	assert.Equal(t, expectedDerivedAddress, derivedAddress)
}

func TestGenerateEd25519Account(t *testing.T) {
	message := []byte{0x12, 0x34}
	account, err := NewEd25519Account()
	assert.NoError(t, err)
	output, err := account.Sign(message)
	assert.NoError(t, err)
	assert.Equal(t, crypto.AccountAuthenticatorEd25519, output.Variant)
	assert.True(t, output.Auth.Verify(message))
}

func TestNewAccountFromSigner(t *testing.T) {
	message := []byte{0x12, 0x34}
	key, err := crypto.GenerateEd25519PrivateKey()
	assert.NoError(t, err)

	account, err := NewAccountFromSigner(key)
	assert.NoError(t, err)
	output, err := account.Sign(message)
	assert.NoError(t, err)
	assert.Equal(t, crypto.AccountAuthenticatorEd25519, output.Variant)
	assert.True(t, output.Auth.Verify(message))

	authKey := key.AuthKey()
	assert.Equal(t, authKey[:], account.Address[:])
}

func TestNewAccountFromSignerWithAddress(t *testing.T) {
	message := []byte{0x12, 0x34}
	key, err := crypto.GenerateEd25519PrivateKey()
	assert.NoError(t, err)

	authenticationKey := crypto.AuthenticationKey{}

	account, err := NewAccountFromSigner(key, authenticationKey)
	assert.NoError(t, err)
	output, err := account.Sign(message)
	assert.NoError(t, err)
	assert.Equal(t, crypto.AccountAuthenticatorEd25519, output.Variant)
	assert.True(t, output.Auth.Verify(message))

	assert.Equal(t, AccountZero, account.Address)
}

func TestNewAccountFromSignerWithAddressMulti(t *testing.T) {
	key, err := crypto.GenerateEd25519PrivateKey()
	assert.NoError(t, err)

	authenticationKey := crypto.AuthenticationKey{}

	_, err = NewAccountFromSigner(key, authenticationKey, authenticationKey)
	assert.Error(t, err)
}
