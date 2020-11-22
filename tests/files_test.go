package tests

import (
	"fmt"
	"github.com/MrNinso/MyGoToolBox/objects/file"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"testing"
	"time"
)

var FolderName string
var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func TestMain(m *testing.M) {
	name, err := ioutil.TempDir("", "MyGoToolBox")

	if err != nil {
		panic(err)
	}

	FolderName = name
	i := m.Run()

	_ = os.Remove(name)

	os.Exit(i)
}

func TestFileWriteReader(t *testing.T) {
	err := os.Mkdir(path.Join(FolderName, "TestFileWriteReader"), os.ModePerm)

	assert.Nil(t, err)

	f, err := os.Create(path.Join(FolderName, "TestFileWriteReader", "file"))

	assert.Nil(t, err)

	myFile := file.File{File: f}

	//Write 5 Random string split by | to file:
	randomString := make([]string, 5)
	for i := 0; i < len(randomString); i++ {
		randomString[i] = RandomString()
	}

	err = myFile.WriteLoop(func(number int) (data []byte, MoveOn bool) {
		return []byte(randomString[number] + "|"), (number + 1) < 5
	}, nil)

	assert.Nil(t, err)

	_, err = myFile.ResetSeekToStart()

	assert.Nil(t, err)

	//Try to read the file content:
	chunks := 0
	err = myFile.ReadAllChuck(func(number int, line []byte) (MoveOn bool) {
		chunks++
		assert.Equal(t, randomString[number], string(line))
		return true
	}, []byte("|")[0], nil)

	assert.Nil(t, err)
	assert.Equal(t, len(randomString), chunks)
}

func RandomString() string {
	return fmt.Sprint(seededRand.Int())
}
