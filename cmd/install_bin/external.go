package install_bin

func InstallBinForExternalGoUse(name string, version string, os string, arch string, binDir string) {
	for _, tool := range Tools {
		if tool.Name != name {
			continue
		}

		getOsFunc := func(x string) string { return x }
		if tool.GetOsFunc != nil {
			getOsFunc = tool.GetOsFunc
		}

		getArchFunc := func(x string) string { return x }
		if tool.GetArchFunc != nil {
			getArchFunc = tool.GetArchFunc
		}

		run(
			tool.Name,
			tool.Aliases,
			tool.SourcePath,
			tool.UrlTemplate,
			func() string { return version },
			tool.GetVersionFunc,
			getUrl,
			getSourcePath,
			getOsFunc,
			getArchFunc,
			func() string { return os },
			func() string { return arch },
			func() string { return binDir },
			tool.RunBeforeInstall,
		)
	}
}
