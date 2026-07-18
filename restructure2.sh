#!/bin/bash
set -x

cd "/Users/kunal.avghade/Documents/Go-basics"

# Move Fundamental files
mv "Fundamental/variables.go" "section-2/1-variables/main.go"
mv "Fundamental/function_and_return_type.go" "section-2/2-functions/main.go"
mv "Fundamental/even_odd.go" "section-2/3-even-odd/main.go"
mv "Fundamental/slice_and_loops.go" "section-2/4-slice-loops/main.go"
mv "Fundamental/oo"/* "section-2/5-oo/"
mv "Fundamental/README.md" "section-2/README.md"
rm -rf "Fundamental"

# Move Map files
mv "Map/main.go" "section-5/1-map/main.go"
mv "Map/README.md" "section-5/README.md"
rm -rf "Map"

# Move Interfaces files
mv "Interfaces/main.go" "section-6/1-interfaces/main.go"
mv "Interfaces/http_inteface.go" "section-6/2-http-interface/main.go"
mv "Interfaces/filereader.go" "section-6/3-filereader/main.go"
mv "Interfaces/shape.go" "section-6/4-shape/main.go"
mv "Interfaces/README.md" "section-6/README.md"
rm -rf "Interfaces"

# Add everything to git
git add -A

