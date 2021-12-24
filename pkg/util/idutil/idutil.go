package idutil

import (
	"crypto/rand"

	"blog-api/pkg/util/iputil"
	"blog-api/pkg/util/stringutil"
	"github.com/sony/sonyflake"
	"github.com/speps/go-hashids/v2"
)

const (
	Alphabet62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	Alphabet36 = "abcdefghijklmnopqrstuvwxyz1234567890"
)

var sf *sonyflake.Sonyflake

func init() {
	var st sonyflake.Settings

	st.MachineID = func() (uint16, error) {
		ip := iputil.GetLocalIP()

		return uint16([]byte(ip)[2])<<8 + uint16([]byte(ip)[3]), nil
	}
	sf = sonyflake.NewSonyflake(st)
}

func GetIntID() (id uint64, err error) {
	if id, err = sf.NextID(); err != nil {
		return 0, err
	}
	return
}

// GetInstanceID return ID form just like secret-2d3to2
func GetInstanceID(uid uint64, prefix string) string {
	hd := hashids.NewData()
	hd.Alphabet = Alphabet36
	hd.MinLength = 6
	hd.Salt = "x20k5x"

	h, err := hashids.NewWithData(hd)
	if err != nil {
		panic(err)
	}

	i, err := h.Encode([]int{int(uid)})
	if err != nil {
		panic(err)
	}

	return prefix + stringutil.Reverse(i)
}

func GetUUID36(prefix string) (string, error) {
	id, err := GetIntID()
	if err != nil {
		return "", err
	}

	hd := hashids.NewData()
	hd.Alphabet = Alphabet36

	hsID, err := hashids.NewWithData(hd)
	if err != nil {
		return "", err
	}

	i, err := hsID.Encode([]int{int(id)})
	if err != nil {
		panic(err)
	}

	return prefix + stringutil.Reverse(i), nil
}

func randString(letters string, n int) string {
	output := make([]byte, n)
	randomness := make([]byte, n)
	_, err := rand.Read(randomness)
	if err != nil {
		panic(err)
	}

	l := len(letters)

	for pos := range output {
		random := randomness[pos]
		// random % 64
		randomPos := random % uint8(l)
		output[pos] = letters[randomPos]
	}

	return string(output)
}

func NewSecretID() string {
	return randString(Alphabet62, 36)
}

func NewSecretKey() string {
	return randString(Alphabet62, 32)
}
