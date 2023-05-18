load("@npm//tests/e2e:@playwright/test/package_json.bzl", playwright_bin = "bin")

def run_playwright(name, executable, target, port, sbu, data = [], **kwargs):

    _env = {
      "EXECUTABLE": executable,
      "SYSTEM_BASE_URL": sbu,
      "E2E_PORT": port,
      "LOG_LEVEL": "ERROR"
    }

    playwright_bin.playwright_test(
        name = name,
        args = [
            "test",
            "--config src/visual/playwright-visual.config.ts",
            target,
        ],
        chdir = native.package_name(),
        data = data + [
            ":node_modules/@playwright/test",
            ":node_modules/playwright",
        ] + native.glob([
            "src/**/*.ts",
            "src/**/*.png",
        ]),
        env = _env
    )
