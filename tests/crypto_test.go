package tests

import (
	"testing"

	"github.com/Cyberpull/gokit/crypto"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type CryptoTestSuite struct {
	suite.Suite

	key string
}

func (x *CryptoTestSuite) SetupSuite() {
	x.key = "as8098sda8078asdg0a870asdg893954"
}

func (x *CryptoTestSuite) TestRandom() {
	value, err := crypto.Rand.String(30)

	require.NoError(x.T(), err)
	require.IsType(x.T(), "", value)
	require.Len(x.T(), value, 30)
}

func (x *CryptoTestSuite) TestEncryptDecryptAES() {
	text := "Just testing encription"

	encrypted, err := crypto.Encrypt.AES(text, x.key)
	require.NoError(x.T(), err)

	decrypted, err := crypto.Decrypt.AES(encrypted, x.key)
	require.NoError(x.T(), err)

	require.EqualValues(x.T(), text, decrypted)
}

func (x *CryptoTestSuite) TestHmac() {
	key := []byte(x.key)
	text := []byte("Just testing Hmac")

	md5Value, err := crypto.HMAC.MD5(key, text)
	require.NoError(x.T(), err)
	require.EqualValues(x.T(), "8d63aac19b2e39e1ff923352149522ad", md5Value)

	sha1Value, err := crypto.HMAC.Sha1(key, text)
	require.NoError(x.T(), err)
	require.EqualValues(x.T(), "aa96b938eaca469ae7f1ea473099f0e8db8c945c", sha1Value)

	sha256Value, err := crypto.HMAC.Sha256(key, text)
	require.NoError(x.T(), err)
	require.EqualValues(x.T(), "7831d8b179454b4b9b17ae223dea1cab996f733a09b48fd47ded6fe8349ac390", sha256Value)

	sha512Value, err := crypto.HMAC.Sha512(key, text)
	require.NoError(x.T(), err)
	require.EqualValues(x.T(), "148be65cc75c4e2da549c60438e8505ac7595f3ad0cc88404b02b751538dd8924b1c5bce12170a457a9261c9e459e48685f4a61fdf547dc91d8706c1ca09c441", sha512Value)
}

func (x *CryptoTestSuite) TestHash() {
	text := []byte("Just testing Hash")

	md5Value, err := crypto.Hash.MD5(text)
	require.NoError(x.T(), err)
	require.EqualValues(x.T(), "9e3671dcd4beaed33945945639692c74", md5Value)

	sha1Value, err := crypto.Hash.Sha1(text)
	require.NoError(x.T(), err)
	require.EqualValues(x.T(), "3808978a0a982bd310d525902654ff333d278f42", sha1Value)

	sha256Value, err := crypto.Hash.Sha256(text)
	require.NoError(x.T(), err)
	require.EqualValues(x.T(), "610bc697806d2dba072a76d32f198385cec0a96788926ad984b236b91253069f", sha256Value)

	sha512Value, err := crypto.Hash.Sha512(text)
	require.NoError(x.T(), err)
	require.EqualValues(x.T(), "0139a1a7b2191b58e386dfdd4f209f0be29fead35fed1e87e2907c3f8d800af000ac27c4b9725e3d1ad2105297b4a8c8248ac396d4b280eca2021a0767812c75", sha512Value)
}

// ===============================

func TestCrypto(t *testing.T) {
	suite.Run(t, new(CryptoTestSuite))
}
