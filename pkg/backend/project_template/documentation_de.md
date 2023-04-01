title: Wie ist die Vorlage aufgebaut?
Unter ***styPackages/*** liegen die Dateien, in denen die Formatierungen definiert sind.

Die Datei ***main.tex*** ist die *Haupt-Datei*. Du kannst sie in einem TeX-Editor deiner Wahl öffnen und kompilieren.

In der Datei***abkuerzungen.csv
*** kannst du Abkürzungen aufführen. Diese werden dann automatisch im Abürzungsverzeichnis aufgelistet. Näheres dazu unter
***Wie kann ich Abkürzungen hinzufügen?***.

In der Datei ***bib_entries.csv*** kannst du Literatureinträge im *SCV*-Format auflisten. Näheres dazu unter
***Wie kann ich Literatureinträge hinzufügen?***.

Die Datei
***citedKeys.csv*** kann dafür sorgen, dass nur bestimmte Literatureinträge im Literaturverzeichnis auftauchen.

---

title: Wie kann ich am besten anfangen?

Das hängt natürlich von dir ab. Öffne die ***main.tex*** einfach mal in einem TeX-Editor und kompiliere sie über
***pdflatex***.
Dies sollte in jedem Fall ohne Probleme funktionieren. Falls nicht, lege bitte ein Issue in Github an. Dann wird das Problem untersucht werden.

Wenn LaTeX erfolgreich gebaut hat, kannst du damit anfangen, die Werte in der Kopf- und Fußzeile anzupassen (
***Wie kann ich die Kopf- und Fußzeile anpassen?
***) oder du kannst anfangen, die Beispiel-Kapitel durch deine eigenen zu ersetzen und deinen Text zu schreiben.

---


