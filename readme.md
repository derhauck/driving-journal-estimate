# Driving Journal Estimator
Tool to estimate the daily amount of kilometers you drove given a set of days and the amount of kilometers you drove in 
total. Very useful tool if you have to journal and missed some days but still wan to have an estimate of how much you drove
daily in a given time interval.

----
# CLI
## Completely Random
```shell
dje calendar random
# Flags
# --out | will write the output to file 'output.txt'
# --total float32 | total KM to distribute (default:10000)
# --days int | days to distribute the total KM to (default:20)
```
example command: `dje calendar random --total 100.00 --days 3`


## From Config
```shell
dje calendar config
# Flags
# --out | will write the output to file 'output.txt'
# --file string | config yaml to read (default: config.yaml)
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
