package entries

import (
	"strings"

	"github.com/TimoSto/ThesorTeX/pkg/backend/tex_escaping"
)

func GenerateCsvForEntries(entries []Entry) string {
	file := "key;type;a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z;\n"
	for _, entry := range entries {
		file += entry.Key + ";" + entry.Category + ";"
		for i := 0; i < 26; i++ {
			if i < len(entry.Fields) {
				//because json converts & into amp;
				strToAdd := strings.Replace(entry.Fields[i], "amp;", "", -1)
				strToAdd = tex_escaping.EscapeField(strToAdd)
				file += strToAdd
			}
			file += ";"
		}
		file += "\n"
	}
	file += "empty;empty;a;b;c;d;e;f;g;h;i;j;k;l;m;n;o;p;q;r;s;t;u;v;w;x;y;z;\n"

	return file
}
