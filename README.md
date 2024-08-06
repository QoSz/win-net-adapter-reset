# Windows Network Adapter Reset Tool

This Go script provides a simple utility to reset all network adapters on a Windows system. It disables all network adapters, waits for 5 seconds, and then re-enables them.

## Features

- Automatically detects all network adapters on the system
- Disables all network adapters
- Waits for 5 seconds
- Re-enables all network adapters
- Requires and checks for administrative privileges

## Prerequisites

- Windows operating system
- Go programming language installed (https://golang.org/doc/install)

## Installation

1. Clone this repository or download the script:
git clone https://github.com/yourusername/windows-network-reset.git
cd windows-network-reset

2. No additional dependencies are required beyond the Go standard library.

## Usage

1. Open a command prompt as an administrator.

2. Navigate to the directory containing the script.

3. Run the script:
go build
./network-reset

**Note:** This script requires administrative privileges to function properly. If not run as an administrator, it will display a message and exit.

## How It Works

1. The script first checks if it has administrative privileges.
2. It then uses the `netsh` command to get a list of all network interfaces.
3. For each interface, it disables the adapter using `netsh`.
4. After disabling all adapters, it waits for 5 seconds.
5. Finally, it re-enables all adapters.

## Warning

This script will temporarily disable all network adapters on your system. Use with caution, especially on remote systems or production environments.

## License

[MIT](https://choosealicense.com/licenses/mit/)