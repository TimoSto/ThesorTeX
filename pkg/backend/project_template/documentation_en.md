title: How is the template structured?
The files in which the formatting is defined are located under ***styPackages/***.

The file ***main.tex*** is the *main file*. You can open it in a TeX editor of your choice and compile it.

In the file***abbreviations.csv
*** you can list abbreviations. These will then be listed automatically in the abbreviations directory. For more information see
***How can I add abbreviations?***.

In the file ***bib_entries.csv*** you can list literature entries in *SCV* format. For more information see
***How can I add literature entries?***.

The file ***citedKeys.csv*** can ensure that only certain literature entries appear in the bibliography.

---

title: What is the best way to start?

That depends on you, of course. Just open the ***main.tex*** in a TeX editor and compile it via ***pdflatex***.
This should work without problems in any case. If not, please create an issue in Github. Then the problem will be investigated.

If LaTeX has built successfully, you can start adjusting the values in the header and footer (
***How can I adjust the header and footer?
***) or you can start replacing the example chapters with your own and write your text.

---
