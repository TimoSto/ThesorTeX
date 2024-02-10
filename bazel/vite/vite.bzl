
def vite_build(name, config_path, root, out, outDir, vite_bin, data = []):

    vite_bin.vite(
        name = name,
        args = [
            "build",
            "--config=" + config_path,
            "--out-dir=" + outDir
        ],
        outs = out,
        srcs = data,
        env = {
            "VITE_ROOT": root
        }
    )
