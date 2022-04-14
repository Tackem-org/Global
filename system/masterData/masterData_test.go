package masterData_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/Tackem-org/Global/system/masterData"
	"github.com/stretchr/testify/assert"
)

func TestMasterDataSetup(t *testing.T) {
	goodFile := "test.json"
	missingFile := "some.json"
	t2 := masterData.Infostruct{
		URL:             "TestKey",
		Port:            50001,
		RegistrationKey: "localhost",
	}

	os.Remove(goodFile)
	os.Unsetenv("REGKEY")
	os.Unsetenv("URL")
	os.Unsetenv("PORT")
	//First Run
	assert.False(t, masterData.Setup(""), "First Setup Run With No Data Should Fail")
	assert.False(t, masterData.UP.Check(), "Check The Master is marked is down")
	assert.Equal(t, masterData.DefaultPort, masterData.Info.Port, "Checking Port is set to the Default Option")
	assert.Equal(t, masterData.DefaultURL, masterData.Info.URL, "Checking the URL is set to the Default Option")
	assert.Equal(t, "", masterData.Info.RegistrationKey, "Checking the Registration Key is blank")

	//Second Run
	assert.Nil(t, os.Setenv("URL", t2.URL))
	assert.Nil(t, os.Setenv("PORT", fmt.Sprint(t2.Port)))
	_, keyPresent1 := os.LookupEnv("REGKEY")
	assert.False(t, keyPresent1)

	assert.False(t, masterData.Setup(""), "Second Setup Run With Env Var missing Reg Key Data Should Fail from lack of Reg Key")
	_, urlPresent1 := os.LookupEnv("URL")
	assert.True(t, urlPresent1)
	_, portPresent1 := os.LookupEnv("PORT")
	assert.True(t, portPresent1)

	//Third Run
	assert.Nil(t, os.Setenv("REGKEY", t2.RegistrationKey))
	_, keyPresent2a := os.LookupEnv("REGKEY")
	assert.True(t, keyPresent2a)

	assert.False(t, masterData.Setup(""), "Third Setup Run With Env Var Data Should Fail from not being able to save it")

	_, urlPresent2 := os.LookupEnv("URL")
	assert.False(t, urlPresent2)
	_, portPresent2 := os.LookupEnv("PORT")
	assert.False(t, portPresent2)
	_, keyPresent2 := os.LookupEnv("REGKEY")
	assert.False(t, keyPresent2)

	//Forth Run
	os.Setenv("REGKEY", t2.RegistrationKey)
	os.Setenv("URL", t2.URL)
	os.Setenv("PORT", fmt.Sprint(t2.Port))
	_, urlPresent3a := os.LookupEnv("URL")
	assert.True(t, urlPresent3a)
	_, portPresent3a := os.LookupEnv("PORT")
	assert.True(t, portPresent3a)
	_, keyPresent3a := os.LookupEnv("REGKEY")
	assert.True(t, keyPresent3a)

	assert.True(t, masterData.Setup(goodFile), "Forth Setup Run With Env Var Data Should Pass")

	//Fifth Run
	assert.True(t, masterData.Setup(goodFile), "Fifth Setup Run With Json Data should pass")

	//Sixth Run
	os.Unsetenv("REGKEY")
	os.Unsetenv("URL")
	os.Unsetenv("PORT")
	os.Remove(missingFile)
	assert.False(t, masterData.Setup(missingFile), "Sixth Setup Run With no Json Data should Fail")

	assert.Nil(t, os.Remove(goodFile))
}

func TestResetFuncs(t *testing.T) {
	assert.NotPanics(t, func() { masterData.ResetFuncs() })
}
