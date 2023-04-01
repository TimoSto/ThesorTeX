title: How is the template structured?
The files in which the formatting is defined are located under ***styPackages/***.

The file ***main.tex*** is the *main file*. You can open it in a TeX editor of your choice and compile it.

In the file ***abbreviations.csv*** you can list abbreviations. These will then be listed automatically in the abbreviations directory. For more information see ***How can I add abbreviations?***.

In the file ***bib_entries.csv*** you can list literature entries in *CSV* format. For more information see ***How can I add literature entries?***.

The file ***citedKeys.csv*** can ensure that only certain literature entries appear in the bibliography.

---

title: What is the best way to start?

That depends on you, of course. Just open the ***main.tex*** in a TeX editor and compile it via ***pdflatex***.
This should work without problems in any case. If not, please create an issue in Github. Then the problem will be investigated.

If LaTeX has built successfully, you can start adjusting the values in the header and footer (***How can I adjust the header and footer?***) or you can start replacing the example chapters with your own and write your text.

---

title: How can I customise the header and footer?
A total of four pieces of information can be displayed in the header and footer. In this template, the title of the work is displayed at the top left
the title of the current chapter, the name of the author (ergo your name) at the bottom left and the page number at the bottom right.
the page number at the bottom right. A distinction is made between the text part of your work and the indexes:

```latex
% text part
\setMainPageStyle{\mytitle}{\nouppercase\parttitle}{\myauthor}{\thepage}
% Directories and miscellaneous
\setPlainPageStyle{\mytitle}{\nouppercase\plaintitle}{\myauthor}{\thepage}
```

You can set the attributes ***mytitle***, ***plaintitle*** and ***myauthor*** as you wish via ***\renewcommand***.
The attributes ***parttitle*** and ***thepage*** refer to the current upper chapter (***\part***) and the current page number.

To switch between the two variants, use ***\frontmatter*** (to *PlainPageStyle*) or ***\mainmatter*** (to *MainPageStyle*).

---

title: How can I change the numbering of the chapters?
By default, the chapters are numbered numerically:

- 1 Chapter 1
- 1.1 Subchapter 1.1
- 1.1.1 Subchapter 1.1.1
- ...

```latex
\usepackage[numeric]{styPackages/headings}

\usepackage[numeric]{styPackages/table-of-contents}
```

You can also switch to an alphanumeric count:

- A Chapter 1
- I Subchapter 1.1
- 1 Subchapter 1.1.1
- a) Subchapter 1.1.1.1
- aa) Subchapter 1.1.1.1
- (1) Subchapter 1.1.1.1.1

```latex
\usepackage[alphaNumeric]{styPackages/headings}

\usepackage[alphaNumeric]{styPackages/table of contents}
```

---

title: How can I add abbreviations?
The abbreviations are read from the file ***abkuerzungen.csv***. There you can enter them in the form *abbreviation*;*meaning*;.

If you call up the command ***\printabbreviations*** in your document, the list of abbreviations is output.

---

title: How can I add attachments?
If your work has appendices, you can output them in a structured way. The commands ***\annexI***, ***\annexII*** and ***\annexIII*** are basically nothing other than the usual chapters.
However, they are numbered separately and output via the command ***\listofanhang***.
