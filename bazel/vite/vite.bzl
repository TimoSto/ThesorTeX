load("@npm//:vite/package_json.bzl", vite_bin = "bin")

def vite_build(name, config_path, root, out, _data = []):

    vite_bin.vite(
        name = name,
        args = [
            "build",
            "--config=" + config_path,
        ],
        outs = out,
        srcs = _data,
        env = {
            "VITEST_ROOT": root
        }
    )
