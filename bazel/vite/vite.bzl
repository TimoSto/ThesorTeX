
def vite_build(name, config_path, root, out, vite_bin, data = [], outDir = None):

    args = [
        "build",
        "--config=" + config_path
    ]

    if outDir:
        args = [
            "build",
            "--config=" + config_path,
            "--out-dir=" + outDir
        ]

    vite_bin.vite(
        name = name,
        args = args,
        outs = out,
        srcs = data,
        env = {
            "VITE_ROOT": root
        }
    )
