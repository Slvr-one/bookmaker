#!/usr/bin/env bash
set -eu

if [ $# -ne 1 ]; then
    echo "Usage: $0 <version>"
    exit 1
fi

version=$1

# Get the current version
get_version() {
    current_version=$(cat VERSION)
    return $current_version
}


# Bump the version
bump_version() {
    new_version=$(semver bump $current_version)
}

# Commit the new version
commit_version() {

}

# Get the current version of the code.

# Bump the version of the code.

# Commit the new version of the code.
git add VERSION
git commit -m "Bumped version to $new_version"
