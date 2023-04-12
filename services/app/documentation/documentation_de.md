MainTitle: Verwendung des Tools für Literaturmanagement
title: Wofür kann ich dieses Tool benutzen?
Dieses Tool erleichtert das Erstellen und Formattieren von Literatureinträgen. Es baut auf der ***Vorlage für Haus- und Abschlussarbeiten*** auf.

Die Motivation hierfür war, dass Anpassungen an *BibLaTeX* sehr umständlich und kompliziert sein können. Der Funktionsumfang grob zusammengefasst ist:
- Erstelle Projekte auf Basis der Vorlage
- Erstelle und bearbeite Literatureinträge über die Weboberfläche
- Diese Literatureinträge kannst du dann ganz einfach in deinem Dokument zitieren
- Erstelle und bearbeite Literatur-Kategorien über die Weboberfläche
- Importiere deine Bibliothek oder einzelne Einträge aus Citavi
---
title: Wie fang ich am besten an?
Gehe auf die ***Downloads*** Seite und lade das Paket für dein Betriebssystem herunter. Wenn du die heruntergeladene ZIP-Datei entpackst, solltest du eine ausführbare Datei sehen. Kopiere diese Datei in den Ordner, in dem die LaTeX-Projekte erstellt werden sollen und starte sie über einen Doppelklick.

Jetzt taucht ein Terminal-Fenster bei dir auf, in welchem die Anwendung läuft. Sobald du dieses Fenster schließt, wird die Anwendung beendet. Außerdem wurde ein Ordner ***projects*** mit einem Beispielprojekt erstellt.

Wenn du die Adresse ***http://localhost:8448*** in einem Browser aufrufst, solltest du die Anwendung sehen.

![Liste der Projekte](./app_images/startpage.png)

Jetzt kannst du anfangen, im ***example***-Projekt die im folgenden beschriebenen Schritte durchzuführen. Oder du erstellt dein eigenes Projekt.

---

title: Wie kann ich einen Literatureintrag anlegen?

Navigiere zunächst in das Projekt, in dem du einen Eintrag hinzufügen möchtest. Klicke dazu auf den entsprechenden Listeneintrag auf der Startseite.

![Projektüberischt: example](./app_images/project_overview.png)

Unter *Literatureinträge* findest du bereits einen *testEntry*. Um einen neuen Eintrag zu erstellen, klicke auf das ***+***-Icon.

Unter *Allgemein* musst du nun einen im Projekt eindeutigen Schlüssel für den neuen Eintrag eingeben. Diesen wirst du verwenden, wenn du diesen Eintrag zitieren willst.

![Neuer Literatureintrag: Allgemein](./app_images/entry_editor_general.png)

Sobald du eine Kategorie ausgewählt hast, werden die Felder dieser Kategorie angezeigt und du kannst Werte eingeben.

![Neuer Literatureintrag: Felder](./app_images/entry_editor_fields.png)

Oben kannst du sehen, wie der Eintrag im Literaturverzeichnis und in einem Zitat aussehen würde.

Wenn du auf das ***Speichern***-Icon in der Toolbar klickst, wird der Eintrag gespeichert. Wenn du nun zurück navigierst, siehst du deinen Eintrag in der Liste.

![Projektüberischt: Mit neuem Eintrag](./app_images/entry_added.png)
