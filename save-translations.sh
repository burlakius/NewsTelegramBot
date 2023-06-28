#!/bin/bash

for locale_path in ./internal/translations/locales/*; do
    cp -v $locale_path/out.gotext.json $locale_path/messages.gotext.json
done

go generate ./internal/translations/translations.go
echo ""
echo "Done!"
