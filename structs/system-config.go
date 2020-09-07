package structs

type SystemConfig struct {
	ZAuth struct {
		API struct {
			Port       string `yaml:"port"`
			Production bool   `yaml:"production"`
		} `yaml:"api"`
		Radius struct {
			HostURL string `yaml:"hostURL"`
			Secret  string `yaml:"secret"`
		} `yaml:"radius"`
		MySQL struct {
			HostIP   string `yaml:"hostIP"`
			Port     int    `yaml:"port"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			DBName   string `yaml:"dbName"`
		}
		Operator struct {
			ShareKey string `yaml:"shareKey"`
			HostURL  string `yaml:"hostURL"`
		} `yaml:"Operator"`
	} `yaml:"zauth"`
}
