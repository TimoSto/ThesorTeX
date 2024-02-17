load("@npm//tests/e2e:@playwright/test/package_json.bzl", playwright_bin = "bin")

def test_playwright(name, workingDir, sut_executable_target, sut_executable, sut_port, system_base_url, test_files, config, specific_deps = [] ):
    playwright_bin.playwright_test(
        name = name,
        args = [
            "test",
            "--config " + config
        ],
        chdir = native.package_name(),
        data = [
            config,
        ] + test_files + specific_deps,
         env = {
             "EXECUTABLE": sut_executable,
             "SYSTEM_BASE_URL": system_base_url,
             "E2E_PORT": sut_port,
             "LOG_LEVEL": "ERROR"
          }
    )
