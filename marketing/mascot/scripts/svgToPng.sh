for i in *.svg; do
    [ -f "$i" ] || break
    rsvg-convert "$i" > "${i%.*}.png"
done
