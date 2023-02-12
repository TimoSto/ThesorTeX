package versions

func GetToolVersions(dev bool) []VersionInfo {
	var versions []VersionInfo

	if dev {
		versions = []VersionInfo{
			{
				Name: "1.0.0",
				Date: "10-01-2023",
			},
			{
				Name: "0.0.2",
				Date: "02-01-2023",
			},
			{
				Name: "0.0.1",
				Date: "20-12-2022",
			},
		}
	}

	return versions
}
