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
