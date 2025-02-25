#/bin/sh
REQUIRED_VERSION="v4.5.0"
VERSION="$(sqlboiler --version |  awk '{print $NF}')"
if [ "$VERSION" != "$REQUIRED_VERSION" ]; then
    echo "sqlboiler required $REQUIRED_VERSION , current $VERSION"
    exit 1
fi


sqlboiler sqlite3  --struct-tag-casing camel  --wipe --no-hooks
