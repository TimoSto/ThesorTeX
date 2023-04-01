title: Wie ist die Vorlage aufgebaut?
Unter ***styPackages/*** liegen die Dateien, in denen die Formatierungen definiert sind.

Die Datei ***main.tex*** ist die *Haupt-Datei*. Du kannst sie in einem TeX-Editor deiner Wahl öffnen und kompilieren.

In der Datei ***abkuerzungen.csv*** kannst du Abkürzungen aufführen. Diese werden dann automatisch im Abürzungsverzeichnis aufgelistet. Näheres dazu unter ***Wie kann ich Abkürzungen hinzufügen?***.

In der Datei ***bib_entries.csv*** kannst du Literatureinträge im *CSV*-Format auflisten. Näheres dazu unter ***Wie kann ich Literatureinträge hinzufügen?***.

Die Datei ***citedKeys.csv*** kann dafür sorgen, dass nur bestimmte Literatureinträge im Literaturverzeichnis auftauchen.

---

title: Wie kann ich am besten anfangen?

Das hängt natürlich von dir ab. Öffne die ***main.tex*** einfach mal in einem TeX-Editor und kompiliere sie über ***pdflatex***.
Dies sollte in jedem Fall ohne Probleme funktionieren. Falls nicht, lege bitte ein Issue in Github an. Dann wird das Problem untersucht werden.

Wenn LaTeX erfolgreich gebaut hat, kannst du damit anfangen, die Werte in der Kopf- und Fußzeile anzupassen (***Wie kann ich die Kopf- und Fußzeile anpassen?***) oder du kannst anfangen, die Beispiel-Kapitel durch deine eigenen zu ersetzen und deinen Text zu schreiben.

---

title: Wie kann ich die Kopf- und Fußzeile anpassen?
Im Header und Footer können insgesamt vier Informationen angezeigt werden. In dieser Vorlage wird oben links der Titel
der Arbeit, oben rechts der Titel des aktuellen Kapitels, unten links der Name des Autors (ergo dein Name) und unten
rechts die Seitenzahl angezeigt. Dabei wird zwischen dem Textteil deiner Arbeit und den Verzeichnissen unterschieden:

```latex
% Textteil
\setMainPageStyle{\mytitle}{\nouppercase\parttitle}{\myauthor}{\thepage}
% Verzeichnisse und Sonstiges
\setPlainPageStyle{\mytitle}{\nouppercase\plaintitle}{\myauthor}{\thepage}
```

Die Attribute ***mytitle***,  ***plaintitle*** und ***myauthor*** kannst du über ***\renewcommand*** beliebig setzen.
Die Attribute ***parttitle*** und ***thepage*** beziehen sich auf das aktuelle Oberkapitel (***\part***) und die aktuelle Seitenzahl.

Der Wechsel zwischen den beiden Varianten zu wechseln, nutze ***\frontmatter*** (zu *PlainPageStyle*) bzw. ***\mainmatter*** (zu *MainPageStyle*).

---

title: Wie kann ich die Nummerierung der Kapitel ändern?
Standardmäßig werden die Kapitel numerisch gezählt:

- 1 Kapitel 1
- 1.1 Unterkapitel 1.1
- 1.1.1 Unterkapitel 1.1.1
- ...

```latex
    \usepackage[numeric]{styPackages/ueberschriften}

    \usepackage[numeric]{styPackages/inhaltsverzeichnis}
```

Du kannst aber auch auf eine alphanumerische Zählung umstellen:

- A Kapitel 1
- I Unterkapitel 1.1
- 1 Unterkapitel 1.1.1
- a) Unterkapitel 1.1.1.1
- aa) Unterkapitel 1.1.1.1.1
- (1) Unterkapitel 1.1.1.1.1.1

```latex
    \usepackage[alphaNumeric]{styPackages/ueberschriften}

    \usepackage[alphaNumeric]{styPackages/inhaltsverzeichnis}
```

---

title: Wie kann ich Abkürzungen hinzufügen?
Die Abkürzungen werden aus der Datei ***abkuerzungen.csv*** ausgelesen. Dort kannst du sie in der Form *Abkürzung*;*Bedeutung*; eintragen.

Wenn du den Befehl ***\printabbreviations*** in deinem Dokument aufrufst, wird das Abkürzungsverzeichnis ausegeben.

---

title: Wie kann ich Anhänge hinzufügen?
Wenn deine Arbeit Anhänge hat, kannst du sie strukturiert ausgeben. Die Befehle ***\anhangI***, ***\anhangII*** und ***\anhangIII*** sind im Grunde nichts anderes als die gewöhnlichen Kapitel.
Sie werden allerdings separat nummeriert und über den Befehl ***\listofanhang*** ausgegeben.
