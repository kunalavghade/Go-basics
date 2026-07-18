#!/bin/bash
set -x

cd "/Users/kunal.avghade/Documents/Go-basics"

# Section 1
git mv "A simple Start" "section-1"
mkdir -p "section-1/1-simple-start"
git mv "section-1/main.go" "section-1/1-simple-start/"

# Section 2
git mv "Fundamental" "section-2"
mkdir -p "section-2/1-variables" "section-2/2-functions" "section-2/3-even-odd" "section-2/4-slice-loops" "section-2/5-oo"
git mv "section-2/variables.go" "section-2/1-variables/main.go"
git mv "section-2/function_and_return_type.go" "section-2/2-functions/main.go"
git mv "section-2/even_odd.go" "section-2/3-even-odd/main.go"
git mv "section-2/slice_and_loops.go" "section-2/4-slice-loops/main.go"
git mv "section-2/oo" "section-2/5-oo/oo"

# Section 3
git mv "Data Struct" "section-3"
mkdir -p "section-3/1-data-struct"
git mv "section-3/main.go" "section-3/1-data-struct/"

# Section 4
git mv "Strings" "section-4"
mkdir -p "section-4/1-strings"
git mv "section-4/main.go" "section-4/1-strings/"

# Section 5
git mv "Map" "section-5"
mkdir -p "section-5/1-map"
git mv "section-5/main.go" "section-5/1-map/"

# Section 6
git mv "Interfaces" "section-6"
mkdir -p "section-6/1-interfaces" "section-6/2-http-interface" "section-6/3-filereader" "section-6/4-shape"
git mv "section-6/main.go" "section-6/1-interfaces/"
git mv "section-6/http_inteface.go" "section-6/2-http-interface/main.go"
git mv "section-6/filereader.go" "section-6/3-filereader/main.go"
git mv "section-6/shape.go" "section-6/4-shape/main.go"

# Section 7
git mv "Channel and Routine" "section-7"
mkdir -p "section-7/1-channel-routine"
git mv "section-7/main.go" "section-7/1-channel-routine/"

# Rename section 10 and 11
git mv "section-10" "section-8"
git mv "section-11" "section-9"

