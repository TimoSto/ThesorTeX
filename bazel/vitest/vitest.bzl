
def vitest_run(name, config_path, root, vitest_bin, data = []):

    vitest_bin.vitest_test(
        name = name,
        args = [
            "run",
            "--config=" + config_path,
        ],
        data = data,
        env = {
            "VITE_ROOT": root
        }
    )