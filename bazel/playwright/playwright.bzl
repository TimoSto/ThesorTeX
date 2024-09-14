
def test_playwright(name, deps, test_files, config, playwright_bin, base_url, sut_executable = None, env = {} ):
    _env = {
         "LOG_LEVEL": "ERROR",
         "SUT_EXECUTABLE": "../../$(rootpath {})".format(sut_executable),
         "BASE_URL": base_url
    }

    for (k, v) in env.items():
        _env[k] = v

    playwright_bin.playwright_test(
        name = name,
        args = [
            "test",
        ],
        chdir = native.package_name(),
        data = [
            config,
        ] + deps + test_files,
         env = _env
    )
