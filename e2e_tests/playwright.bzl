load("@aspect_rules_ts//ts:defs.bzl", "ts_config", "ts_project")
load("@pnpm//e2e_tests:@cucumber/cucumber/package_json.bzl", "bin")

# Run service specific playwright tests. This rule ues ts_project to generate js files.
# And runs playwright tests via the cucumber.js runner.
def playwright_test(name, service, sut_executable = None, sut_url = "http://localhost:8448", data = [], debug = False, ignore_https_errors = False, visibility = None, **kwargs):
    ts_config(
        name = name + "_tsconfig",
        src = "tsconfig.json",
    )
    ts_project(
        name = name + "_playwright_js",
        srcs = native.glob(["**/*.ts"]),
        allow_js = True,
        declaration = True,
        out_dir = name + "_build",
        tsconfig = ":" + name + "_tsconfig",
        deps = [
            ":node_modules/@cucumber/cucumber",
            ":node_modules/@playwright",
            ":node_modules/@types/node",
            ":node_modules/typescript",
            ":node_modules/ts-node",
        ],
    )

    _env = {
        "BROWSER": "chromium",
        "HEADLESS": "false",
        "SUT_EXECUTABLE": sut_url,
        "IGNORE_HTTPS_ERRORS": "false",
    }

    bin.cucumber_js_test(
        name = name,
        args = [
            "--config cucumber.json",
        ],
        chdir = native.package_name(),
        data = data + [
            "cucumber.json",
            name + "_playwright_js",
            ":node_modules/playwright",
        ] + native.glob([
            # We lookup only service specific feature files.
            "features/" + service + "/*.feature",
        ]),
        env = _env,
        **kwargs
    )
