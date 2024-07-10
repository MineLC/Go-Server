package conf

var DefaultServerConfig = ServerConfig{
	Settings{
		Debug: true,
	},
	Network{
		Host: "0.0.0.0",
		Port: 25565,
	},
	Motd{
		Line:    "&bA Golang server",
		Favicon: "",
		Max:     2024,
		Online:  0,
	},
}

type ServerConfig struct {
	Settings Settings
	Network  Network
	Motd     Motd
}

type Network struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

type Settings struct {
	Debug bool `toml:"debug"`
}

type Motd struct {
	Line    string `toml:"motd"`
	Favicon string `toml:"favicon"`

	Max    int `toml:"max"`
	Online int `toml:"online"`
}
