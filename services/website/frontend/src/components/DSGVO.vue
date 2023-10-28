<script setup lang="ts">

</script>

<template>
  <v-dialog width="950">
    <template v-slot:activator="{props}">
      <v-btn variant="text" color="white" style="font-weight: bold;" v-bind="props" text="Datenschutzerklärung"></v-btn>
    </template>
    <v-card theme="light">
      <v-card-text>
        <h2 class="text-h3">Datenschutzerklärung</h2>
        <v-expansion-panels :multiple="true">
          <v-expansion-panel>
            <v-expansion-panel-title>
              Präamble
            </v-expansion-panel-title>
            <v-expansion-panel-text>
              <p class="text-body-1">Mit der folgenden Datenschutzerklärung möchten wir Sie darüber aufklären, welche
                Arten
                Ihrer personenbezogenen Daten (nachfolgend auch kurz als "Daten" bezeichnet) wir zu welchen Zwecken und
                in welchem Umfang verarbeiten. Die Datenschutzerklärung gilt für alle von uns durchgeführten
                Verarbeitungen
                personenbezogener Daten, sowohl im Rahmen der Erbringung unserer Leistungen als auch insbesondere auf
                unseren Webseiten, in mobilen Applikationen sowie innerhalb externer Onlinepräsenzen, wie z. B. unserer
                Social-Media-Profile (nachfolgend zusammenfassend bezeichnet als "Onlineangebot").</p>
              <p class="text-body-1">Die verwendeten Begriffe sind nicht geschlechtsspezifisch.</p>
              <p class="text-body-1">Stand: 28.10.2023</p>
            </v-expansion-panel-text>
          </v-expansion-panel>
          <v-expansion-panel>
            <v-expansion-panel-title>Kontakt</v-expansion-panel-title>
            <v-expansion-panel-text>thesortex.contact@gmail.com</v-expansion-panel-text>
          </v-expansion-panel>
          <v-expansion-panel>
            <v-expansion-panel-title>Übersicht der Verarbeitungen</v-expansion-panel-title>
            <v-expansion-panel-text>
              <p class="text-body-1">Die nachfolgende Übersicht fasst die Arten der verarbeiteten Daten und die Zwecke
                ihrer Verarbeitung zusammen
                und verweist auf die betroffenen Personen.</p>
              <span class="text-h5">Arten der verarbeiteten Daten</span>
              <v-list>
                <v-list-item>
                  <v-list-item-title>Nutzungsdaten</v-list-item-title>
                </v-list-item>
                <v-list-item>
                  <v-list-item-title>Meta-, Kommunikations- und Verfahrensdaten</v-list-item-title>
                </v-list-item>
              </v-list>
              <span class="text-h5">Kategorien betroffener Nutzer</span>
              <v-list>
                <v-list-item>
                  <v-list-item-title>Besucher der Online-Angebote</v-list-item-title>
                </v-list-item>
              </v-list>
              <span class="text-h5">Zwecke der Verarbeitung</span>
              <v-list>
                <v-list-item>
                  <v-list-item-title>Sicherheitsmaßnahmen</v-list-item-title>
                </v-list-item>
                <v-list-item>
                  <v-list-item-title>Bereitstellung unseres Onlineangebotes und Nutzerfreundlichkeit</v-list-item-title>
                </v-list-item>
                <v-list-item>
                  <v-list-item-title>Informationstechnische Infrastruktur</v-list-item-title>
                </v-list-item>
              </v-list>
            </v-expansion-panel-text>
          </v-expansion-panel>
          <v-expansion-panel>
            <v-expansion-panel-title>Rechtsgrundlage</v-expansion-panel-title>
            <v-expansion-panel-text>
              <p class="text-body-1">Im Folgenden erhalten Sie eine Übersicht der Rechtsgrundlagen der DSGVO, auf deren
                Basis wir personenbezogene Daten verarbeiten. Bitte nehmen Sie zur Kenntnis, dass neben den Regelungen
                der DSGVO nationale Datenschutzvorgaben in Ihrem bzw. unserem Wohn- oder Sitzland gelten können. Sollten
                ferner im Einzelfall speziellere Rechtsgrundlagen maßgeblich sein, teilen wir Ihnen diese in der
                Datenschutzerklärung mit.</p>
              <p class="text-body-1"><span class="font-weight-bold">Berechtigte Interessen (Art. 6 Abs. 1 S. 1 lit. f) DSGVO)</span>
                - Die Verarbeitung ist
                zur Wahrung der berechtigten Interessen des Verantwortlichen oder eines Dritten erforderlich, sofern
                nicht die Interessen oder Grundrechte und Grundfreiheiten der betroffenen Person, die den Schutz
                personenbezogener Daten erfordern, überwiegen.</p>
              <p class="text-body-1"><span
                class="font-weight-bold">Nationale Datenschutzregelungen in Deutschland</span>
                - Zusätzlich zu den Datenschutzregelungen der DSGVO gelten nationale Regelungen zum Datenschutz in
                Deutschland. Hierzu gehört insbesondere das Gesetz zum Schutz vor Missbrauch personenbezogener Daten bei
                der Datenverarbeitung (Bundesdatenschutzgesetz – BDSG). Das BDSG enthält insbesondere Spezialregelungen
                zum Recht auf Auskunft, zum Recht auf Löschung, zum Widerspruchsrecht, zur Verarbeitung besonderer
                Kategorien personenbezogener Daten, zur Verarbeitung für andere Zwecke und zur Übermittlung sowie
                automatisierten Entscheidungsfindung im Einzelfall einschließlich Profiling. Ferner können
                Landesdatenschutzgesetze der einzelnen Bundesländer zur Anwendung gelangen.</p>
            </v-expansion-panel-text>
          </v-expansion-panel>
          <v-expansion-panel>
            <v-expansion-panel-title>Sicherheitsmaßnahmen</v-expansion-panel-title>
            <v-expansion-panel-text>
              <p class="text-body-1">Wir treffen nach Maßgabe der gesetzlichen Vorgaben unter Berücksichtigung des
                Stands der Technik, der Implementierungskosten und der Art, des Umfangs, der Umstände und der Zwecke der
                Verarbeitung sowie der unterschiedlichen Eintrittswahrscheinlichkeiten und des Ausmaßes der Bedrohung
                der Rechte und Freiheiten natürlicher Personen geeignete technische und organisatorische Maßnahmen, um
                ein dem Risiko angemessenes Schutzniveau zu gewährleisten.</p>
              <p class="text-body-1">Zu den Maßnahmen gehören insbesondere die Sicherung der Vertraulichkeit, Integrität
                und Verfügbarkeit von Daten durch Kontrolle des physischen und elektronischen Zugangs zu den Daten als
                auch des sie betreffenden Zugriffs, der Eingabe, der Weitergabe, der Sicherung der Verfügbarkeit und
                ihrer Trennung. Des Weiteren haben wir Verfahren eingerichtet, die eine Wahrnehmung von
                Betroffenenrechten, die Löschung von Daten und Reaktionen auf die Gefährdung der Daten gewährleisten.
                Ferner berücksichtigen wir den Schutz personenbezogener Daten bereits bei der Entwicklung bzw. Auswahl
                von Hardware, Software sowie Verfahren entsprechend dem Prinzip des Datenschutzes, durch
                Technikgestaltung und durch datenschutzfreundliche Voreinstellungen.</p>
              <p class="text-body-1">Um die Daten der Benutzer, die über unsere Online-Dienste übertragen werden, zu
                schützen, verwenden wir TLS/SSL-Verschlüsselung. Secure Sockets Layer (SSL) ist die Standardtechnologie
                zur Sicherung von Internetverbindungen durch Verschlüsselung der zwischen einer Website oder App und
                einem Browser (oder zwischen zwei Servern) übertragenen Daten. Transport Layer Security (TLS) ist eine
                aktualisierte und sicherere Version von SSL. Hyper Text Transfer Protocol Secure (HTTPS) wird in der URL
                angezeigt, wenn eine Website durch ein SSL/TLS-Zertifikat gesichert ist.</p>
            </v-expansion-panel-text>
          </v-expansion-panel>
          <v-expansion-panel>
            <v-expansion-panel-title>Nutzung von Content Delivery Networks</v-expansion-panel-title>
            <v-expansion-panel-text>
              <p class="text-body-1">Wir nutzen das Content Delivery Network (CDN) Amazon CloudFront von Amazon Web
                Services EMEA SARL, 38 avenue John F. Kennedy, L-1855 Luxembourg (AWS), um die Sicherheit und die
                Auslieferungsgeschwindigkeit unserer Website zu erhöhen. Dies entspricht unserem berechtigten Interesse
                (Art. 6 Abs. 1 lit. f DSGVO). Ein CDN ist ein Netzwerk aus weltweit verteilten Servern, das in der Lage
                ist, optimiert Inhalte an den Websitenutzer auszuliefern. Für diesen Zweck können personenbezogene Daten
                in Server-Logfiles von AWS verarbeitet werden. Bitte vergleichen Sie die Ausführungen unter „Hosting“.
                Zusätzlich bewahren wir anonymisierte Logfiles auf, um die Stabilität und Sicherheit unserer Website zu
                gewährleisten.</p>
              <p class="text-body-1">AWS ist Empfänger Ihrer personenbezogenen Daten und als Auftragsverarbeiter für uns
                tätig. Die entspricht unserem berechtigten Interesse im Sinne des Art. 6 Abs. 1 S. 1 lit. f DSGVO,
                selbst
                kein Content Delivery Network zu betreiben.</p>
              <p class="text-body-1">Sie haben das Recht der Verarbeitung zu widersprechen. Ob der Widerspruch
                erfolgreich
                ist, ist im Rahmen einer Interessenabwägung zu ermitteln.
                Die Verarbeitung der unter diesem Abschnitt angegebenen Daten ist weder gesetzlich noch vertraglich
                vorgeschrieben. Die Funktionsfähigkeit der Website ist ohne die Verarbeitung nicht gewährleistet.
                Ihre personenbezogenen Daten werden von AWS so lange gespeichert, wie es für die beschriebenen Zwecke
                erforderlich ist.</p>
              <p class="text-body-1">Weitere Informationen zu Widerspruchs- und Beseitigungsmöglichkeiten gegenüber AWS
                finden Sie unter:
                <a>https://d1.awsstatic.com/legal/privacypolicy/AWS_Privacy_Notice__German_Translation.pdf</a>
                AWS hat Compliance-Maßnahmen für internationale Datenübermittlungen umgesetzt. Diese gelten für alle
                weltweiten Aktivitäten, bei denen AWS personenbezogene Daten von natürlichen Personen in der EU
                verarbeitet. Diese Maßnahmen basieren auf den EU-Standardvertragsklauseln (SCCs). Weitere Informationen
                finden Sie unter: <a>https://d1.awsstatic.com/legal/aws-gdpr/AWS_GDPR_DPA.pdf</a>
              </p>
            </v-expansion-panel-text>
          </v-expansion-panel>
          <v-expansion-panel>
            <v-expansion-panel-title>Löschung von Daten</v-expansion-panel-title>
            <v-expansion-panel-text>
              <p class="text-body-1">Die von uns verarbeiteten Daten werden nach Maßgabe der gesetzlichen Vorgaben
                gelöscht, sobald deren zur Verarbeitung erlaubten Einwilligungen widerrufen werden oder sonstige
                Erlaubnisse entfallen (z. B. wenn der Zweck der Verarbeitung dieser Daten entfallen ist oder sie für den
                Zweck nicht erforderlich sind). Sofern die Daten nicht gelöscht werden, weil sie für andere und
                gesetzlich zulässige Zwecke erforderlich sind, wird deren Verarbeitung auf diese Zwecke beschränkt. D.
                h., die Daten werden gesperrt und nicht für andere Zwecke verarbeitet. Das gilt z. B. für Daten, die aus
                handels- oder steuerrechtlichen Gründen aufbewahrt werden müssen oder deren Speicherung zur
                Geltendmachung, Ausübung oder Verteidigung von Rechtsansprüchen oder zum Schutz der Rechte einer anderen
                natürlichen oder juristischen Person erforderlich ist. Unsere Datenschutzhinweise können ferner weitere
                Angaben zu der Aufbewahrung und Löschung von Daten beinhalten, die für die jeweiligen Verarbeitungen
                vorrangig gelten.</p>
            </v-expansion-panel-text>
          </v-expansion-panel>
          <v-expansion-panel>
            <v-expansion-panel-title>Rechte der betroffenen Personen</v-expansion-panel-title>
            <v-expansion-panel-text>
              <p class="text-body-1">Rechte der betroffenen Personen aus der DSGVO: Ihnen stehen als Betroffene nach der
                DSGVO verschiedene Rechte zu, die sich insbesondere aus Art. 15 bis 21 DSGVO ergeben:
              </p>
              <v-expansion-panels multiple>
                <v-expansion-panel>
                  <v-expansion-panel-title>Widerspruchsrecht</v-expansion-panel-title>
                  <v-expansion-panel-text>Sie haben das Recht, aus Gründen, die sich aus Ihrer besonderen Situation
                    ergeben, jederzeit gegen die Verarbeitung der Sie betreffenden personenbezogenen Daten, die aufgrund
                    von Art. 6 Abs. 1 lit. e oder f DSGVO erfolgt, Widerspruch einzulegen; dies gilt auch für ein auf
                    diese Bestimmungen gestütztes Profiling. Werden die Sie betreffenden personenbezogenen Daten
                    verarbeitet, um Direktwerbung zu betreiben, haben Sie das Recht, jederzeit Widerspruch gegen die
                    Verarbeitung der Sie betreffenden personenbezogenen Daten zum Zwecke derartiger Werbung einzulegen;
                    dies gilt auch für das Profiling, soweit es mit solcher Direktwerbung in Verbindung steht.
                  </v-expansion-panel-text>
                </v-expansion-panel>
                <v-expansion-panel>
                  <v-expansion-panel-title>Widerrufsrecht bei Einwilligungen</v-expansion-panel-title>
                  <v-expansion-panel-text>Sie haben das Recht, erteilte Einwilligungen jederzeit zu widerrufen.
                  </v-expansion-panel-text>
                </v-expansion-panel>
                <v-expansion-panel>
                  <v-expansion-panel-title>Auskunftsrecht</v-expansion-panel-title>
                  <v-expansion-panel-text>Sie haben das Recht, eine Bestätigung darüber zu verlangen, ob betreffende
                    Daten verarbeitet werden und auf Auskunft über diese Daten sowie auf weitere Informationen und Kopie
                    der Daten entsprechend den gesetzlichen Vorgaben.
                  </v-expansion-panel-text>
                </v-expansion-panel>
                <v-expansion-panel>
                  <v-expansion-panel-title>Recht auf Berichtigung</v-expansion-panel-title>
                  <v-expansion-panel-text>Sie haben entsprechend den gesetzlichen Vorgaben das Recht, die
                    Vervollständigung der Sie betreffenden Daten oder die Berichtigung der Sie betreffenden unrichtigen
                    Daten zu verlangen.
                  </v-expansion-panel-text>
                </v-expansion-panel>
                <v-expansion-panel>
                  <v-expansion-panel-title>Recht auf Löschung</v-expansion-panel-title>
                  <v-expansion-panel-text>Sie haben nach Maßgabe der gesetzlichen Vorgaben das Recht, zu verlangen, dass
                    Sie betreffende Daten unverzüglich gelöscht werden.
                  </v-expansion-panel-text>
                </v-expansion-panel>
              </v-expansion-panels>
            </v-expansion-panel-text>
          </v-expansion-panel>
          <v-expansion-panel>
            <v-expansion-panel-title>Bereitstellung des Onlineangebotes und Webhosting</v-expansion-panel-title>
            <v-expansion-panel-text>
              <p class="text-body-1">Wir verarbeiten die Daten der Nutzer, um ihnen unsere Online-Dienste zur Verfügung
                stellen zu können. Zu diesem Zweck verarbeiten wir die IP-Adresse des Nutzers, die notwendig ist, um die
                Inhalte und Funktionen unserer Online-Dienste an den Browser oder das Endgerät der Nutzer zu
                übermitteln</p>
              <v-expansion-panels>
                <v-expansion-panel>
                  <v-expansion-panel-title>Verarbeitete Datenarten</v-expansion-panel-title>
                  <v-expansion-panel-text>
                    <v-list>
                      <v-list-item>Nutzungsdaten (z. B. besuchte Webseiten, Interesse an Inhalten, Zugriffszeiten)
                      </v-list-item>
                      <v-list-item>Meta-, Kommunikations- und Verfahrensdaten (z. .B. IP-Adressen, Zeitangaben,
                        Identifikationsnummern, Einwilligungsstatus).
                      </v-list-item>
                    </v-list>
                  </v-expansion-panel-text>
                </v-expansion-panel>
                <v-expansion-panel>
                  <v-expansion-panel-title>Betroffene Personen</v-expansion-panel-title>
                  <v-expansion-panel-text>
                    <v-list>
                      <v-list-item>Nutzer (z. .B. Webseitenbesucher, Nutzer von Onlinediensten)</v-list-item>
                    </v-list>
                  </v-expansion-panel-text>
                </v-expansion-panel>
                <v-expansion-panel>
                  <v-expansion-panel-title>Zwecke der Verarbeitung</v-expansion-panel-title>
                  <v-expansion-panel-text>
                    <p class="text-body-1">Bereitstellung unseres Onlineangebotes und Nutzerfreundlichkeit;
                      Informationstechnische Infrastruktur (Betrieb und Bereitstellung von Informationssystemen und
                      technischen Geräten (Computer, Server etc.).). Sicherheitsmaßnahmen.</p>
                  </v-expansion-panel-text>
                </v-expansion-panel>
              </v-expansion-panels>
              <p class="text-body-1"><span class="font-weight-bold">Rechtsgrundlagen</span>: Berechtigte Interessen
                (Art. 6 Abs. 1 S. 1 lit. f) DSGVO).</p>
            </v-expansion-panel-text>
          </v-expansion-panel>
          <v-expansion-panel>
            <v-expansion-panel-title>Weitere Hinweise zu Verarbeitungsprozessen, Verfahren und Diensten
            </v-expansion-panel-title>
            <v-expansion-panel-text>
              <v-expansion-panels>
                <v-expansion-panel>
                  <v-expansion-panel-title>Bereitstellung Onlineangebot auf gemietetem Speicherplatz
                  </v-expansion-panel-title>
                  <v-expansion-panel-text>
                    <p class="text-body-1">Für die Bereitstellung unseres Onlineangebotes nutzen wir Speicherplatz,
                      Rechenkapazität und Software, die wir von einem entsprechenden Serveranbieter (auch "Webhoster"
                      genannt) mieten oder anderweitig beziehen; Rechtsgrundlagen: Berechtigte Interessen (Art. 6 Abs. 1
                      S. 1 lit. f) DSGVO).</p>
                  </v-expansion-panel-text>
                </v-expansion-panel>
                <v-expansion-panel>
                  <v-expansion-panel-title>Erhebung von Zugriffsdaten und Logfiles
                  </v-expansion-panel-title>
                  <v-expansion-panel-text>
                    <p class="text-body-1">Der Zugriff auf unser Onlineangebot wird in Form von so genannten
                      "Server-Logfiles" protokolliert. Zu den Serverlogfiles können die Adresse und Name der abgerufenen
                      Webseiten und Dateien, Datum und Uhrzeit des Abrufs, übertragene Datenmengen, Meldung über
                      erfolgreichen Abruf, Browsertyp nebst Version, das Betriebssystem des Nutzers, Referrer URL (die
                      zuvor besuchte Seite) und im Regelfall IP-Adressen und der anfragende Provider gehören. Die
                      Serverlogfiles können zum einen zu Zwecken der Sicherheit eingesetzt werden, z. B., um eine
                      Überlastung der Server zu vermeiden (insbesondere im Fall von missbräuchlichen Angriffen,
                      sogenannten DDoS-Attacken) und zum anderen, um die Auslastung der Server und ihre Stabilität
                      sicherzustellen; Rechtsgrundlagen: Berechtigte Interessen (Art. 6 Abs. 1 S. 1 lit. f) DSGVO). </p>
                  </v-expansion-panel-text>
                </v-expansion-panel>
                <v-expansion-panel>
                  <v-expansion-panel-title>Content-Delivery-Network
                  </v-expansion-panel-title>
                  <v-expansion-panel-text>
                    <p class="text-body-1">Wir setzen ein "Content-Delivery-Network" (CDN) ein. Ein CDN ist ein Dienst,
                      mit dessen Hilfe Inhalte eines Onlineangebotes, insbesondere große Mediendateien, wie Grafiken
                      oder Programm-Skripte, mit Hilfe regional verteilter und über das Internet verbundener Server
                      schneller und sicherer ausgeliefert werden können; Rechtsgrundlagen: Berechtigte Interessen (Art.
                      6 Abs. 1 S. 1 lit. f) DSGVO).</p>
                  </v-expansion-panel-text>
                </v-expansion-panel>
              </v-expansion-panels>
            </v-expansion-panel-text>
          </v-expansion-panel>
          <v-expansion-panel>
            <v-expansion-panel-title>Änderung und Aktualisierung der Datenschutzerklärung</v-expansion-panel-title>
            <v-expansion-panel-text>
              <p class="text-body-1">Wir bitten Sie, sich regelmäßig über den Inhalt unserer Datenschutzerklärung zu
                informieren. Wir passen die Datenschutzerklärung an, sobald die Änderungen der von uns durchgeführten
                Datenverarbeitungen dies erforderlich machen. Wir informieren Sie, sobald durch die Änderungen eine
                Mitwirkungshandlung Ihrerseits (z. B. Einwilligung) oder eine sonstige individuelle Benachrichtigung
                erforderlich wird.</p>
            </v-expansion-panel-text>
          </v-expansion-panel>
          <v-expansion-panel>
            <v-expansion-panel-title>Begriffsdefinitionen</v-expansion-panel-title>
            <v-expansion-panel-text>
              <p class="text-body-1">In diesem Abschnitt erhalten Sie eine Übersicht über die in dieser
                Datenschutzerklärung verwendeten Begrifflichkeiten. Soweit die Begrifflichkeiten gesetzlich definiert
                sind, gelten deren gesetzliche Definitionen. Die nachfolgenden Erläuterungen sollen dagegen vor allem
                dem Verständnis dienen.</p>
              <v-expansion-panels>
                <v-expansion-panel>
                  <v-expansion-panel-title>Personenbezogene Daten</v-expansion-panel-title>
                  <v-expansion-panel-text>"Personenbezogene Daten" sind alle Informationen, die sich auf eine
                    identifizierte oder identifizierbare natürliche Person (im Folgenden "betroffene Person") beziehen;
                    als identifizierbar wird eine natürliche Person angesehen, die direkt oder indirekt, insbesondere
                    mittels Zuordnung zu einer Kennung wie einem Namen, zu einer Kennnummer, zu Standortdaten, zu einer
                    Online-Kennung (z. B. Cookie) oder zu einem oder mehreren besonderen Merkmalen identifiziert werden
                    kann, die Ausdruck der physischen, physiologischen, genetischen, psychischen, wirtschaftlichen,
                    kulturellen oder sozialen Identität dieser natürlichen Person sind.
                  </v-expansion-panel-text>
                </v-expansion-panel>
                <v-expansion-panel>
                  <v-expansion-panel-title>Verarbeitung</v-expansion-panel-title>
                  <v-expansion-panel-text>Verarbeitung" ist jeder mit oder ohne Hilfe automatisierter Verfahren
                    ausgeführte Vorgang oder jede solche Vorgangsreihe im Zusammenhang mit personenbezogenen Daten. Der
                    Begriff reicht weit und umfasst praktisch jeden Umgang mit Daten, sei es das Erheben, das Auswerten,
                    das Speichern, das Übermitteln oder das Löschen.
                  </v-expansion-panel-text>
                </v-expansion-panel>
              </v-expansion-panels>
            </v-expansion-panel-text>
          </v-expansion-panel>
        </v-expansion-panels>
      </v-card-text>
    </v-card>
  </v-dialog>
</template>

<style scoped lang="scss">

</style>