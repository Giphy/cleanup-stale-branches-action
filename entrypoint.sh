output=$(/usr/bin/cleanup-stale-branches-action)
echo "OUTPUT=$output" >> "$GITHUB_OUTPUT"