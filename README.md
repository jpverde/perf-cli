# perf-cli

A command-line system performance monitoring tool built with Go that provides detailed system information including CPU, memory, disk I/O statistics, and top processes.

## Features

- **CPU Information**: Get physical and logical CPU counts with real-time usage statistics
- **Memory Information**: Display memory usage statistics in MB with detailed breakdown
- **Disk Information**: Show disk I/O counters and statistics in JSON format
- **Top Processes**: Display top 5 processes by CPU and memory usage
- **Verbose Output**: Get detailed system information when needed

## Installation

### Prerequisites

- Go 1.19 or higher

### Build from Source

```bash
git clone <repository-url>
cd perf-cli
go mod tidy
go build -o perf-cli
```

### Install Dependencies

This tool uses the following Go packages:
- `github.com/shirou/gopsutil/v4` - System and process utilities
- `github.com/spf13/cobra` - CLI framework

## Usage

### Basic Command Structure

```bash
./perf-cli sysinfo [flags]
```

### Available Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--cpu` | `-c` | Get CPU information (physical/logical cores and usage) |
| `--mem` | `-m` | Get Memory information (total, free, used in MB) |
| `--disk` | `-d` | Get Disk I/O information and statistics |
| `--top` | `-t` | Get Top 5 processes by CPU and Memory usage |
| `--verbose` | `-v` | Get verbose output with additional details |
| `--help` | `-h` | Display help information |

### Command Aliases

You can use the following aliases for the `sysinfo` command:
- `sys`
- `sysinfo`

## Examples

### Get CPU Information

```bash
./perf-cli sysinfo --cpu
# or
./perf-cli sys -c
```

**Output:**
```
Physical CPUs: 8, Logical CPUs: 16
Waiting for 3 seconds...
------- Total CPU Usage -------
Total CPU 0: 0.437318%
```

**Verbose Output:**
```
Physical CPUs: 8, Logical CPUs: 16
Waiting for 3 seconds...
------- Total CPU Usage -------
Total CPU 0: 0.395668%
----------------------------------------
Waiting for 3 seconds...
------- Per logical CPUs Usage -------
CPU 0: 0.000000%
CPU 1: 0.000000%
CPU 2: 0.000000%
CPU 3: 0.334448%
CPU 4: 0.000000%
CPU 5: 0.000000%
CPU 6: 0.000000%
CPU 7: 0.000000%
CPU 8: 0.000000%
CPU 9: 0.000000%
CPU 10: 0.671141%
CPU 11: 0.332226%
CPU 12: 0.332226%
CPU 13: 0.000000%
CPU 14: 0.000000%
CPU 15: 0.000000%
----------------------------------------
```

### Get Memory Information

```bash
./perf-cli sysinfo --mem
# or
./perf-cli sys -m
```

**Output:**
```
Total: 8192 MB, Free: 2048 MB, Used: 6144 MB, UsedPercent: 75.00%
```

**Verbose Output:**
```
Total: 7292 MB, Free: 5474 MB, Used: 1308 MB, UsedPercent: 17.943282%
Sin: 0 MB
Sout: 0 MB
SwapTotal: 2048 MB
SwapFree: 2048 MB
SwapUsedPercent: 0%
```

### Get Disk Information

```bash
./perf-cli sysinfo --disk
# or
./perf-cli sys -d
```

**Output:**
```
Total: 1031018 MB, Free: 968385 MB, Used: 10187 MB, UsedPercent: 1.041096%
```

### Get Top Processes

```bash
./perf-cli sysinfo --top
# or
./perf-cli sys -t
```

**Output:**
```
Top 5 processes by CPU usage:
PID      NAME                 CPU%       MEM%       MEM(MB)   
----------------------------------------------------------------
1234     chrome               15.50      8.20       512       
5678     firefox              12.30      6.10       384       
9012     node                 8.70       4.50       256       
3456     docker               5.20       12.10      768       
7890     postgres             3.80       9.30       592       

Top 5 processes by Memory usage:
PID      NAME                 CPU%       MEM%       MEM(MB)   
----------------------------------------------------------------
3456     docker               5.20       12.10      768       
7890     postgres             3.80       9.30       592       
1234     chrome               15.50      8.20       512       
5678     firefox              12.30      6.10       384       
9012     node                 8.70       4.50       256       
```

Note: Verbose output shows all processes, not just the top 5.

### Combine Multiple Flags

```bash
# Get CPU and Memory information
./perf-cli sysinfo --cpu --mem 
#or
./perf-cli sys -cm

# Get all information with verbose output
./perf-cli sysinfo -c -m -d -t -v
#or
./perf-cli sys -cmdv

# Get memory and top processes
./perf-cli sys -m -t
#or
./perf-cli sys -mt
```

### Verbose Output

Add the `--verbose` or `-v` flag to any command for additional detailed information:

```bash
./perf-cli sysinfo --mem --verbose
```

## Output Format

- **CPU**: Shows physical and logical CPU counts and usage statistics
- **Memory**: Shows total, free, used memory in MB, and swap usage
- **Disk**: Shows disk I/O counters and statistics
- **Processes**: Tabular format showing PID, Name, CPU%, Memory%, and Memory usage in MB

## Error Handling

The tool gracefully handles errors such as:
- Inaccessible processes (permission issues)
- System information retrieval failures
- Invalid process states

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

[Add your license information here]

## Dependencies

- [gopsutil](https://github.com/shirou/gopsutil) - Cross-platform library for retrieving system and process information
- [Cobra](https://github.com/spf13/cobra) - A Commander for modern Go CLI interactions

## System Compatibility

This tool is compatible with:
- Linux
- macOS
- Windows

*Note: Some features may have platform-specific behavior due to underlying system differences.*
