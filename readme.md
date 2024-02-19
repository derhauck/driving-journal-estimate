# Driving Journal Estimator


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

---

## From Config
```shell
dje calendar config
# Flags
# --out | will write the output to file 'output.txt'
# --file string | config yaml to read (default: config.yaml)
```

```Yaml
total: float32
baseline: float32
days:
  - date: string
    count: int
```
---