
def vitest_run(name, root, vitest_bin, data = [], config_path = "" ):

    args = ["run"]

    if config_path != "":
        args += ["--config=" + config_path]

    vitest_bin.vitest_test(
        name = name,
        args = args,
        data = data,
        env = {
            "VITE_ROOT": root
        }
    )