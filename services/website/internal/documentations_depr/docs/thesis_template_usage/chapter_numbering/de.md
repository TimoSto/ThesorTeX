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
