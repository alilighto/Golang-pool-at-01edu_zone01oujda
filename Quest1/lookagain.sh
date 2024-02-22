find . -type f -iname '*.sh' -printf '%f\n' | cut -d. -f1 | sort -r
