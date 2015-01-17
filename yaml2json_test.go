package yaml2json

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEqual(t *testing.T) {
	yamlFile, err := os.Open("_testdata/fig.yml")
	require.NoError(t, err)
	defer yamlFile.Close()
	yamlFileData, err := ioutil.ReadAll(yamlFile)
	require.NoError(t, err)

	jsonFile, err := os.Open("_testdata/fig.json")
	require.NoError(t, err)
	defer jsonFile.Close()
	jsonFileData, err := ioutil.ReadAll(jsonFile)
	require.NoError(t, err)
	jsonFileData, err = unmarshalMarshalJson(jsonFileData)
	require.NoError(t, err)

	output, err := Convert(yamlFileData)
	require.NoError(t, err)
	require.Equal(t, string(output), string(jsonFileData))
}

// Strips whitespace
func unmarshalMarshalJson(jsonFileData []byte) ([]byte, error) {
	var data interface{}
	err := json.Unmarshal(jsonFileData, &data)
	if err != nil {
		return nil, err
	}
	output, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return output, nil
}
