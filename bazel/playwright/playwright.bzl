
def test_playwright(name, deps, test_files, config, playwright_bin ):
    playwright_bin.playwright_test(
        name = name,
        args = [
            "test",
        ],
        chdir = native.package_name(),
        data = [
            config,
        ] + deps + test_files,
         env = {
             "LOG_LEVEL": "ERROR"
          }
    )
