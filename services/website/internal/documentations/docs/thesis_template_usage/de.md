# Verwenden der Vorlage für wissenschaftliche Arbeiten

Wenn du die Vorlage als ZIP-Datei heruntergeladen und entpackt hast, wirst du folgende Struktur sehen:

- *styPackages/*
- *abkuerzungen.csv*
- *bib_entries.csv*
- *citedKeys.csv*
- *main.tex*

Die Datei *main.tex" ist die *Hauptdatei*. Hier schreibst du deine Arbeit im groben so wie in gewöhnlichem LaTeX.

Bis Zeile *109* werden verschiedene Packages geladen. Danach geht's los mit dem Inhalt.

## Titel und Autor im Header/Footer

Mit den Befehlen **\renewcommand{\mytitle}{Titel meiner Arbeit\\mit zwei Zeilen}** und **\renewcommand{\myauthor}{Tom
Template}** werden die Werte angepasst, die oben und unten links auf jeder Seite stehen. Falls der Titel umgebrochen
wird, kannst du auch die Höhe des Headers anpassen: **\renewcommand{\headheight}{27pt}**.

Es gibt zwei Möglichkeiten, den Header und Footer zu formattieren.

### FrontMatter

Mit **\frontmatter** aktivierst du den Front-Matter-Stil. Dieser ist z.B. für das Inhaltsverzeichnis gedacht. Der Stil
wird über folgenden Befehl angegeben:

*\setPlainPageStyle{\mytitle}{\nouppercase\plaintitle}{\myauthor}{\thepage}*

Die Werte für **mytitle** und **myauthor** sowie den **plaintitle** kannst du beliebig über *\renewcommand* setzen. *
*thepage** verweist immer auf die aktuelle Seite.

### MainMatter

Mit **\mainmatter** wechselst du in den Stil, bei dem oben rechts der Titel des aktuellen Oberkapitels (**\part**)
angezeigt wird.

*\setMainPageStyle{\mytitle}{\nouppercase\parttitle}{\myauthor}{\thepage}*

Möchtest du den Titel dort nicht anzeigen, z.B. weil er zu lang ist, entferne den Wert einfach aus **\setMainPageStyle**
bzw. ***\setPlainPageStyle*.
