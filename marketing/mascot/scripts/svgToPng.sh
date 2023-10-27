if ! command -v rsvg-convert &> /dev/null
then
    echo "rsvg-convert could not be found"
    exit 1
fi

for i in *.svg; do
    [ -f "$i" ] || break
    echo "converting $i..."
    rsvg-convert "$i" > "${i%.*}.png"
done
