function semver {
    . .env
    ver=$1
    #patch=$patch
    #minor=$minor
    #majot=$major
    if [ "$ver" == "patch" ]; then
        echo 'patch + 1'
        let $patch++
    elif [ "$ver" == "minor" ]; then
        echo 'minor + 1'
        let $minor++
        patch=0
    elif [ "$ver" == "major" ]; then
        echo 'major + 1'
        let $major++
        patch=0
        minor=0
    else 
        echo '--please supply version upgrade severety--'
    fi
    echo "patch=$patch" >> .env
    echo "minor=$minor" >> .env
    echo "major=$major" >> .env

    this_ver="$major.$minor.$patch"
    echo "this_ver='$major.$minor.$patch'" >> .env
    return this_ver
}
semver $1