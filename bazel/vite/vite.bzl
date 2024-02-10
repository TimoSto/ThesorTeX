
def vite_build(name, config_path, root, out, vite_bin, data = []):

    vite_bin.vite(
        name = name,
        args = [
            "build",
            "--config=" + config_path,
        ],
        outs = out,
        srcs = data,
        env = {
            "VITE_ROOT": root
        }
    )
