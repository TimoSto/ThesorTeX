Wenn du die Vorlage als ZIP-Datei heruntergeladen und entpackt hast, wirst du folgende Struktur sehen:

- *styPackages/*
- *abkuerzungen.csv*
- *bib_entries.csv*
- *citedKeys.csv*
- *main.tex*

Die Datei *main.tex* ist die *Hauptdatei*. Hier schreibst du deine Arbeit im groben so wie in gewöhnlichem LaTeX.

# Wie kann ich die Nummerierung der Kapitel ändern?

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
    \usepackage[latour]{styPackages/ueberschriften}

    \usepackage[latour]{styPackages/inhaltsverzeichnis}
```

# Wie kann ich den Header und Footer ändern?

Im Header un Footer können insgesamt vier Informationen angezeigt werden. In dieser Vorlage wird oben links der Titel
der Arbeit, oben rechts der Titel des aktuellen Kapitels, unten links der Name des Autors (ergo dein Name) und unten
rechts die Seitenzahl angezeigt. Dabei wird zwischen dem Textteil deiner Arbeit und den Verzeichnissen unterschieden:

```latex
\setPlainPageStyle{\mytitle}{\nouppercase\plaintitle}{\myauthor}{\thepage}

\setMainPageStyle{\mytitle}{\nouppercase\parttitle}{\myauthor}{\thepage}
```

Die Attribute *mytitle* *plaintitle* und *myauthor* kannst du über *\renewcommand* beliebig setzen. Die Attribute
*parttitle* und *thepage* beziehen sich auf das aktuelle Oberkapitel (*\part*) und die aktuelle Seitenzahl.

Der Wechsel zwischen den beiden Varianten zu wechseln, nutze *\frontmatter* bzw. *\mainmatter*.

# Wie kann ich ein Abkürzungsverzeichnis erstellen?

Die Abkürzungen werden aus der Datei *abkuerzungen.csv* ausgelesen. Dort kannst du sie in der Form * **Abkürzung**;*
*Bedeutung** * eintragen.
