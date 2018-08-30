#!/usr/bin/env bash

GOPATH=$(go env GOPATH)
GOBIN="$GOPATH/bin"
GENNY_BIN="$GOBIN/genny"

INFILE="property/property_gen.go"
OUTFILE="property/property_typed.go"

# simple types
cat "$INFILE" | "$GENNY_BIN" gen "PropValueType=byte,int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,bool,float32,float64"   > "$OUTFILE"
