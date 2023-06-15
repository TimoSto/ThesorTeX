load("@npm//tests/e2e:@cucumber/cucumber/package_json.bzl", cucmber_bin = "bin")

def run_playwright_cucumber(name, executable, target, port, sbu, data = [], **kwargs):

    _env = {
        "EXECUTABLE": executable,
        "SYSTEM_BASE_URL": sbu,
        "E2E_PORT": port,
        "LOG_LEVEL": "ERROR"
    }

    cucmber_bin.cucumber_js_test(
            name = name,
            args = [
                "--config cucumber.json",
            ],
            chdir = native.package_name(), # change the wd to /tests/e2e
            data = data + [
                "tsconfig.json",
                "cucumber.json",
                "//:node_modules/typescript",
                "//:node_modules/@types/node",
                "//:node_modules/ts-node",
                ":node_modules/@cucumber/cucumber",
                ":node_modules/@playwright/test",
                ":node_modules/playwright",
                ":node_modules/@axe-core/playwright",
            ] + native.glob([
                "features/" + target + "/*.feature",
                "src/**/*.ts",
            ]),
            env = _env
        )