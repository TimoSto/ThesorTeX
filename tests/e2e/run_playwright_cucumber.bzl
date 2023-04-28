load("@npm//tests/e2e:@cucumber/cucumber/package_json.bzl", cucmber_bin = "bin")
load("@aspect_rules_ts//ts:defs.bzl", "ts_config", "ts_project")

def run_playwright_cucumber(name, executable, target, port, sbu, data = [], **kwargs):
    ts_config(
        name = name + "_tsconfig",
        src = "tsconfig.json",
    )

    ts_project(
        name = name + "_playwright_js",
        srcs = native.glob(["src/**/*.ts"]),
        allow_js = True,
        declaration = True,
        source_map = True,
        out_dir = name + "_build",
        tsconfig = ":" + name + "_tsconfig",
        deps = [
            "//:node_modules/typescript",
            "//:node_modules/@types/node",
            "//:node_modules/ts-node",
            ":node_modules/@cucumber/cucumber",
            ":node_modules/@playwright/test",
            ":node_modules/playwright",
            ":node_modules/@axe-core/playwright",
        ],
    )

    _env = {
        "EXECUTABLE": executable,
        "SYSTEM_BASE_URL": sbu,
        "E2E_PORT": port,
    }

    cucmber_bin.cucumber_js_test(
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
                "features/" + target + "/*.feature",
                # The playwright inspectors shows now the original typescript source file
                # by including source maps and the original typescript files.
                "src/**/*.ts",
            ]),
            env = _env
        )