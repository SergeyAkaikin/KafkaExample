package test

import (
	"task/internal/config"
	"testing"
)

type configStruct struct {
	TestField1 string `json:"testField1"`
	TestField4 string `json:"testField4"`
	TestField5 string `json:"testField5"`
	TestField2 int    `json:"testField2"`
	TestField3 int    `json:"testField3"`
}

var pathConf = `./config_test.json`

func TestConfig(t *testing.T) {

	var loader config.Loader = config.JsonLoader{}

	conf1 := configStruct{}
	conf2 := configStruct{
		TestField1: `test1`,
		TestField2: 5,
		TestField3: 10,
		TestField4: `test2`,
		TestField5: `test3`,
	}

	if err := loader.Load(pathConf, &conf1); err != nil {
		t.Errorf("Error in loading config: %v", pathConf)
	}

	if conf2 != conf1 {
		t.Errorf("conf1 != conf2, conf1=%v, conf2=%v\n", conf1, conf2)
	}

}
