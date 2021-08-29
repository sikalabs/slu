package install_bin_tool

var Tools = []Tool{
	{
		Name:        "terraform",
		Version:     "1.0.5",
		UrlTemplate: hashicorpUrlTemplate("terraform"),
	},
	{
		Name:        "vault",
		Version:     "1.8.2",
		UrlTemplate: hashicorpUrlTemplate("vault"),
	},
}

func hashicorpUrlTemplate(name string) string {
	return "https://releases.hashicorp.com/" + name + "/{{.Version}}/" + name +
		"_{{.Version}}_{{.Os}}_{{.Arch}}.zip"
}
