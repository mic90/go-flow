#!/usr/bin/env bash

GOPATH=$(go env GOPATH)
GOBIN="$GOPATH/bin"
GENNY_BIN="$GOBIN/genny"

INFILE="port/port_gen.go"
ARRAY_INFILE="port/port_array_gen.go"

OUTFILE="port/port"

# simple types
cat "$INFILE"       | "$GENNY_BIN" gen "ValueType=byte,int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,bool,float32,float64"         > "$OUTFILE"_typed.go
cat "$ARRAY_INFILE" | "$GENNY_BIN" gen "ArrayValueType=byte,int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,bool,float32,float64"    > "$OUTFILE"_array_typed.go
