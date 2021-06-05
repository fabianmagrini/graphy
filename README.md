# Graphy

Graphy is tool that allows you to delcaritively define the architecture of your system and will generate diagrams

## Example generate

```sh
go run graphy.go generate examples/basic.yml --template dot.tmpl | dot -Tpng -o examples/basic.dot.png
```

## Example using convert (old format)

```sh
go run graphy.go convert examples/basic.yml --template dot.tmpl | dot -Tpng -o examples/basic.dot.png
```

## Using the chokidar filewatcher

```sh
chokidar "examples/*.yml" -c "./run.sh {path}"
```

## Consolidated

```sh
go run graphy.go generate examples/filesA*.yml examples/filesB*.yml --template dot.tmpl | dot -Tpng -o examples/consolidated.dot.png
go run graphy.go generate examples/filesA*.yml examples/filesB*.yml --template puml.tmpl | puml generate -o examples/consolidated.puml.png
```

## Run test filter

```sh
go run graphy.go generate examples/filesA*.yml examples/filesB*.yml --template dot.tmpl --filters examples/test-filters.json | dot -Tpng -o examples/filtered.dot.png
go run graphy.go generate examples/filesA*.yml examples/filesB*.yml --template puml.tmpl --filters examples/test-filters.json | puml generate -o examples/filtered.puml.png
```
