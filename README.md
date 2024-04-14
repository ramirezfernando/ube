# Ube
<img width="100%" alt="Ube banner" src="./images/ube-banner.png"/>

[![Latest Release](https://img.shields.io/github/v/release/ramirezfernando/ube)](https://github.com/ramirezfernando/ube/releases/tag/v2.0.0)
![Coverage](https://img.shields.io/badge/Coverage-82.0%25-brightgreen)
![GitHub Downloads (all assets, all releases)](https://img.shields.io/github/downloads/ramirezfernando/ube/total)
![GitHub License](https://img.shields.io/github/license/ramirezfernando/ube)

Ube (ooh-bay) is a program that displays statistics about your code. Ube will show the number of files, and total lines within those files grouped by language.

<img width="600" alt="Ube example" src="./images/ube-example.png"/>

## Table of Contents
- [Features](#features)
- [Installation](#installation)
   - [macOS](#macos)
   - [Linux](#linux)
   - [Windows](#windows)
- [Usage](#usage)
- [Common Issues](#common-issues)
- [Contributing](#contributing)
- [License](#license)

## Features <a name="features"></a>
- Ube has a huge range of languages, supporting over **220** language extensions.
- It's **accurate**, and **consistent** as it counts the number of newline characters `/n` present in a specified path. This ensures consistency across different platforms and text editors.
- The code statistics are displayed in a **clean**, **easy-to-read** format leveraging the following libraries from [Charm](https://github.com/charmbracelet):
    - [Bubble Tea](https://github.com/charmbracelet/bubbletea) - A powerful little TUI framework
    - [Bubbles](https://github.com/charmbracelet/bubbles) - TUI components for Bubble Tea
    - [Log](https://github.com/charmbracelet/log) - A minimal, colorful Go logging library

## Installation <a name="installation"></a>

### macOS <a name="macos"></a>
1. Download the appropriate release archive for your platform from the [Latest Release](https://github.com/ramirezfernando/ube/releases/tag/v2.0.0) page:
   - For Intel-based Macs: [ube_2.0.0_darwin_amd64.tar.gz](https://github.com/ramirezfernando/ube/releases/download/v2.0.0/ube_2.0.0_darwin_amd64.tar.gz)
   - For Apple Silicon Macs: [ube_2.0.0_darwin_arm64.tar.gz](https://github.com/ramirezfernando/ube/releases/download/v2.0.0/ube_2.0.0_darwin_arm64.tar.gz)

2. Extract the archive from your downloads folder using the following command in your terminal:
    ```bash
    $ tar -xzf ~/Downloads/ube_2.0.0_darwin_amd64.tar.gz # or ube_2.0.0_darwin_arm64.tar.gz
    ```
3. Move the extracted binary to a directory in your PATH for convenient access from any directory:
    ```bash
    $ mv ube /usr/local/bin/ube
    ```
4. Verify the installation (might have to allow in settings):
    ```
    $ ube
    ```

### Linux <a name="linux"></a>
1. Download the appropriate release archive for your platform from the [Latest Release](https://github.com/ramirezfernando/ube/releases/tag/v2.0.0) page:
    - For 32-bit systems: [ube_2.0.0_linux_386.tar.gz](https://github.com/ramirezfernando/ube/releases/download/v2.0.0/ube_2.0.0_linux_386.tar.gz)
    - For 64-bit systems: [ube_2.0.0_linux_amd64.tar.gz](https://github.com/ramirezfernando/ube/releases/download/v2.0.0/ube_2.0.0_linux_amd64.tar.gz)
    - For ARM 64-bit systems: [ube_2.0.0_linux_arm64.tar.gz](https://github.com/ramirezfernando/ube/releases/download/v2.0.0/ube_2.0.0_linux_arm64.tar.gz)
2. Extract the archive from your downloads folder using the following command in your terminal:
    ```bash
    $ tar -xzf ~/Downloads/ube_2.0.0_linux_386.tar.gz # or ube_2.0.0_linux_amd64.tar.gz or ube_2.0.0_linux_arm64.tar.gz
    ```
3. Move the extracted binary to a directory in your PATH for convenient access from any directory:
    ```bash
    $ sudo mv ube /usr/local/bin/ube
    ```
4. Set the correct permissions for the binary:
    ```bash
    $ sudo chmod +x /usr/local/bin/ube
    ```
5. Verify the installation:
    ```bash
    $ ube
    ```

### Windows <a name="windows"></a>
1. Download the appropriate release archive for your platform from the [Latest Release](https://github.com/ramirezfernando/ube/releases/tag/v2.0.0) page:
    - For 32-bit systems: [ube_2.0.0_windows_386.tar.gz](https://github.com/ramirezfernando/ube/releases/download/v2.0.0/ube_2.0.0_windows_386.tar.gz)
    - For 64-bit systems: [ube_2.0.0_windows_amd64.tar.gz](https://github.com/ramirezfernando/ube/releases/download/v2.0.0/ube_2.0.0_windows_amd64.tar.gz)
    - For ARM 64-bit systems: [ube_2.0.0_windows_arm64.tar.gz](https://github.com/ramirezfernando/ube/releases/download/v2.0.0/ube_2.0.0_windows_arm64.tar.gz)
2. Extract the archive from your downloads folder using a tool like 7-Zip or WinRAR.
3. Move the extracted binary (ube.exe) to a directory included in your system's PATH environment variable for convenient access. Alternatively, you can run the tool from its extracted location.
4. Verify the installation:
    ```
    $ ube
    ```

## Usage <a name="usage"></a>
```bash
$ ube ./your/path
```

## Common Issues <a name="common-issues"></a>
### “ube” can’t be opened because Apple cannot check it for malicious software.
1. Naviagte to `Settings` > `Security & Privacy`.
2. Under `General` tab, click "Allow Anyway" next to the message about the app being blocked.

## Contributing <a name="contributing"></a>
Contributions are welcome! Please check the [Contributing Guidelines](CONTRIBUTING.md) for a few helpful tips!

## License <a name="license"></a>
This project is licensed under the [MIT License](LICENSE).
