load("@npm//tests/e2e:@cucumber/cucumber/package_json.bzl", cucmber_bin = "bin")

def run_cucumber(name, workingDir, sut_executable_target, sut_executable, sut_port, system_base_url, test_files):
    cucmber_bin.cucumber_js_test(
        name = name,
        args = [
            "--config cucumber.json",
        ],
        chdir = workingDir,
        data = [
            sut_executable_target,
            "cucumber.json",
            "//:node_modules/typescript",
            "//:node_modules/@types/node",
            "//:node_modules/ts-node",
            "//tests/e2e:node_modules/@cucumber/cucumber",
            "//tests/e2e:playwright_base_deps",
            "//tests/e2e:playwright_app_steps",
        ] + test_files,
        env = {
            "EXECUTABLE": sut_executable,
            "SYSTEM_BASE_URL": system_base_url,
            "E2E_PORT": sut_port,
            "LOG_LEVEL": "ERROR"
         }
    )
