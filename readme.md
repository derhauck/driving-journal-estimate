# Driving Journal Estimator

The Driving Journal Estimator is a tool designed to estimate your daily kilometers driven, ideal for backfilling missed entries in a driving journal. 

## Features

- **Random Distribution**: Distributes total kilometers across days randomly.
- **Configurable Estimation**: Allows more precise control based on your preferences.

## Installation
The Driving Journal Estimator is available for Linux and macOS. Install it using our convenient install script:

```shell
curl -sSL https://raw.githubusercontent.com/DERHauck/driving-journal-estimate/main/install.sh | sudo bash
```
This script automatically detects your OS, downloads the latest dje binary, and sets it up for immediate use.


## CLI Usage

### Random Distribution
Generates estimates randomly. Useful for a quick approximation without specific requirements.

```shell
dje calendar random
--total float32: Total kilometers to distribute (default: 10000).
--days int: Number of days to spread the total kilometers over (default: 20).
--out: Optional. Writes output to output.txt (default: prints to console).
```
### Configuration-based Distribution
Generates estimates based on a detailed YAML configuration for more accurate distribution reflecting actual driving patterns.

```shell
dje calendar config
--file string: The YAML configuration file specifying the distribution logic (default: "config.yaml").
--out: Optional. Writes output to output.txt (default: prints to console).
```
#### Type - config.yaml
```Yaml
total: float32 | total KM to distribute
baseline: float32 | smaller value means more differences, higher value means less differences between days with same `count`
days:
  - date: string | display name (irrelevant for calculation)
    count: int | weight of the day, the higher the more of the kilometers will be distributed relative to the other days
```
example config: [config.yaml](cli/config.yaml)

example command: `dje calendar config --file config.yaml`

---
