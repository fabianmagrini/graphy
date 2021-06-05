#!/bin/bash
inputFile=$1

dotOutputFile="${inputFile%.*}.dot.png"
pumlOutputFile="${inputFile%.*}.puml.png"
c4pumlOutputFile="${inputFile%.*}.c4puml.png"

go run graphy.go generate $inputFile --template templates/dot.tmpl | dot -Tpng -o $dotOutputFile
go run graphy.go generate $inputFile --template templates/puml.tmpl | puml generate -o $pumlOutputFile
go run graphy.go generate $inputFile --template templates/c4puml.tmpl | puml generate -o $c4pumlOutputFile