<script setup lang="ts">

</script>

<template>
  <v-dialog width="950">
    <template v-slot:activator="{props}">
      <v-btn icon color="background" style="font-weight: bold; margin-top:-5px" v-bind="props"
             title="Datenschutzerklärung">
        <svg xmlns="http://www.w3.org/2000/svg" version="1.0" width="30" height="30"
             viewBox="0 0 30 30" preserveAspectRatio="xMidYMid meet">
          <g transform="translate(0,30) scale(0.006000000,-0.006000000)" fill="#FFFFFF" stroke="none">
            <path
              d="M2517 5109 c-9 -6 -40 -56 -68 -112 -29 -56 -56 -105 -62 -108 -5 -4 -52 -12 -104 -19 -115 -15 -147 -36 -147 -94 0 -30 9 -46 50 -89 27 -28 63 -68 81 -87 l31 -35 -18 -108 c-20 -114 -19 -129 13 -164 32 -36 73 -29 176 28 l92 50 91 -50 c102 -57 143 -64 175 -28 32 35 33 50 13 164 l-18 108 26 30 c15 17 51 56 81 87 46 48 55 64 55 94 0 58 -32 79 -148 94 -52 8 -100 16 -104 19 -5 3 -33 52 -61 108 -28 56 -60 107 -70 113 -25 13 -61 12 -84 -1z m63 -320 c16 -38 45 -57 99 -67 l42 -7 -36 -43 c-34 -39 -37 -46 -33 -97 1 -30 1 -55 -2 -55 -3 1 -16 7 -30 15 -33 19 -91 19 -116 0 -33 -25 -39 -18 -36 40 4 52 2 58 -34 98 l-38 42 43 7 c56 10 85 29 101 67 7 17 16 31 20 31 4 0 13 -14 20 -31z" />
            <path
              d="M3580 4928 c-17 -9 -41 -46 -76 -119 l-51 -105 -109 -17 c-94 -14 -112 -20 -131 -42 -42 -49 -32 -78 59 -171 l81 -84 -18 -121 c-17 -119 -17 -121 2 -147 41 -55 69 -52 209 26 l71 40 89 -51 c105 -59 142 -64 181 -24 l27 26 -13 98 c-6 54 -14 109 -17 123 -5 21 8 40 75 112 69 73 81 91 81 120 0 58 -32 78 -153 96 l-105 15 -47 98 c-26 54 -54 104 -63 112 -25 23 -64 29 -92 15z m61 -327 c23 -41 54 -61 94 -61 38 0 38 -15 -1 -61 -33 -40 -35 -45 -28 -96 l6 -54 -33 20 c-43 26 -69 26 -119 1 l-40 -20 6 58 c6 58 6 60 -30 100 -40 44 -41 42 31 57 31 6 62 37 76 78 9 24 13 22 38 -22z" />
            <path
              d="M1477 4819 c-10 -6 -41 -57 -68 -112 -27 -56 -57 -105 -67 -108 -9 -4 -57 -12 -105 -19 -109 -15 -142 -37 -140 -95 1 -33 11 -47 82 -119 l80 -81 -15 -90 c-23 -138 -21 -165 10 -192 39 -33 68 -29 172 28 l94 51 95 -51 c104 -56 141 -63 172 -28 32 35 33 50 14 167 l-18 111 79 83 c94 101 104 131 62 174 -24 23 -43 31 -109 41 -143 22 -128 12 -176 118 -24 52 -54 104 -68 114 -26 21 -66 24 -94 8z m100 -366 c8 -7 35 -15 59 -19 l44 -6 -36 -42 c-34 -39 -36 -44 -32 -99 3 -38 1 -56 -6 -53 -6 2 -27 11 -46 19 -35 15 -51 13 -121 -17 -14 -6 -15 -1 -9 37 8 54 -2 82 -41 123 l-30 31 42 7 c59 9 76 21 100 68 l21 42 19 -40 c11 -21 27 -44 36 -51z" />
            <path
              d="M4429 4091 c-14 -10 -44 -62 -68 -114 -48 -106 -33 -97 -175 -118 -96 -15 -130 -39 -130 -92 0 -31 11 -47 81 -120 l80 -83 -18 -113 c-20 -121 -16 -148 23 -175 36 -26 67 -19 163 35 l90 51 97 -51 c106 -56 133 -61 172 -30 31 25 33 59 11 194 l-14 89 79 82 c44 45 81 90 81 100 5 77 -20 97 -151 117 -57 9 -106 20 -110 24 -3 4 -23 46 -44 93 -21 47 -49 95 -63 108 -29 27 -72 28 -104 3z m71 -311 c23 -45 40 -57 99 -66 l42 -7 -30 -31 c-40 -42 -51 -73 -41 -125 7 -37 5 -41 -9 -35 -71 30 -86 32 -121 17 -19 -8 -40 -17 -46 -19 -7 -3 -9 13 -6 50 4 52 2 57 -32 98 -32 37 -35 43 -19 46 79 14 94 23 115 67 12 25 23 45 25 45 1 0 12 -18 23 -40z" />
            <path
              d="M697 4079 c-9 -5 -39 -57 -66 -113 l-49 -104 -43 -7 c-155 -24 -173 -29 -196 -51 -43 -43 -33 -76 57 -169 44 -46 80 -91 80 -100 0 -9 -7 -64 -16 -121 -16 -103 -16 -106 3 -132 40 -53 64 -53 173 6 54 28 103 52 108 52 5 0 49 -22 98 -50 97 -54 111 -57 154 -35 41 21 46 53 26 178 l-17 108 80 85 c68 72 81 91 81 119 0 62 -30 82 -152 101 l-108 16 -45 96 c-25 53 -53 104 -62 114 -18 20 -77 24 -106 7z m73 -319 c23 -45 38 -55 98 -65 l43 -7 -41 -44 c-41 -44 -41 -46 -34 -100 l7 -56 -41 21 c-51 26 -63 26 -113 1 l-40 -21 7 49 c7 54 -2 80 -43 124 l-25 27 38 6 c53 10 82 30 101 70 9 19 18 35 20 35 1 0 12 -18 23 -40z" />
            <path
              d="M2415 3823 c-250 -53 -451 -248 -516 -501 -15 -60 -19 -110 -19 -247 l0 -172 -36 -6 c-46 -8 -97 -47 -118 -90 -14 -31 -16 -97 -16 -606 0 -546 1 -574 20 -628 43 -121 130 -214 250 -265 55 -23 57 -23 549 -26 319 -2 512 1 544 8 126 27 240 117 299 237 l33 68 3 575 c2 316 0 590 -3 608 -9 48 -74 110 -125 118 l-40 7 0 172 c0 137 -4 187 -19 247 -92 358 -446 579 -806 501z m241 -164 c138 -25 285 -132 352 -256 51 -96 62 -157 62 -340 l0 -163 -510 0 -510 0 0 163 c0 112 5 181 15 220 68 263 323 425 591 376z m584 -1014 l0 -85 -100 0 c-94 0 -103 -2 -125 -25 -17 -16 -25 -35 -25 -60 0 -25 8 -44 25 -60 22 -23 31 -25 125 -25 l100 0 0 -85 0 -85 -90 0 c-95 0 -142 -14 -154 -45 -13 -35 -6 -76 19 -100 22 -23 31 -25 125 -25 l100 0 0 -85 0 -85 -91 0 c-120 0 -158 -21 -159 -88 0 -22 8 -41 25 -57 22 -23 31 -25 125 -25 l100 0 0 -27 c-2 -76 -55 -156 -132 -199 l-53 -29 -495 0 -495 0 -53 29 c-77 43 -130 123 -132 199 l0 27 100 0 c94 0 103 2 125 25 17 16 25 35 25 57 -1 67 -39 88 -159 88 l-91 0 0 85 0 85 100 0 c94 0 103 2 125 25 25 24 32 65 19 100 -12 31 -59 45 -154 45 l-90 0 0 85 0 85 100 0 c94 0 103 2 125 25 34 33 34 87 0 120 -22 23 -31 25 -125 25 l-100 0 0 85 0 85 680 0 680 0 0 -85z" />
            <path
              d="M2455 2451 c-177 -80 -200 -332 -40 -439 l50 -34 5 -108 c5 -97 8 -112 30 -135 32 -35 88 -35 120 0 22 23 25 38 30 135 l5 108 50 34 c160 107 137 359 -40 439 -54 25 -156 25 -210 0z m162 -173 c25 -23 31 -74 13 -108 -14 -27 -64 -41 -100 -28 -36 12 -50 35 -50 79 0 70 85 106 137 57z" />
            <path
              d="M390 2978 c-17 -9 -41 -46 -76 -119 l-51 -105 -109 -17 c-94 -14 -112 -20 -132 -42 -41 -49 -31 -78 60 -171 l82 -84 -18 -121 c-18 -119 -18 -122 1 -147 41 -56 75 -51 219 31 l61 35 89 -50 c105 -60 142 -65 181 -25 l27 26 -13 98 c-6 54 -14 109 -17 123 -4 21 9 40 75 112 69 73 81 92 81 120 0 58 -32 78 -153 96 l-105 15 -47 98 c-26 54 -54 104 -63 112 -25 23 -64 29 -92 15z m59 -326 c20 -40 55 -62 96 -62 38 0 38 -15 -1 -61 -33 -40 -35 -45 -28 -96 l6 -54 -33 20 c-43 26 -69 26 -119 1 l-40 -20 6 58 c6 58 6 60 -30 100 -40 44 -41 43 30 57 36 7 58 27 73 68 6 15 13 27 16 27 2 0 13 -17 24 -38z" />
            <path
              d="M4670 2983 c-31 -12 -50 -37 -95 -132 l-47 -98 -105 -16 c-81 -12 -111 -21 -129 -38 -32 -30 -33 -88 -3 -115 12 -10 48 -49 81 -86 l59 -67 -16 -88 c-22 -126 -20 -155 10 -183 39 -36 77 -30 179 28 l90 51 70 -40 c138 -79 168 -82 209 -27 19 25 19 28 1 147 l-18 121 82 84 c91 93 101 122 59 171 -19 23 -37 29 -131 43 l-109 16 -47 101 c-26 55 -54 106 -61 112 -18 15 -60 23 -79 16z m40 -317 c0 -25 40 -64 73 -71 72 -15 71 -13 31 -57 -36 -40 -36 -42 -30 -100 l6 -58 -40 20 c-50 25 -76 25 -119 -1 l-33 -20 6 54 c7 51 5 56 -28 96 -39 46 -39 61 -1 61 41 0 76 22 96 62 19 36 39 43 39 14z" />
            <path
              d="M743 1968 c-13 -6 -43 -52 -75 -117 l-53 -106 -105 -17 c-125 -20 -105 -14 -130 -40 -12 -13 -20 -35 -20 -56 0 -30 12 -46 82 -118 l81 -84 -16 -99 c-22 -128 -21 -143 4 -175 32 -41 77 -35 183 24 l89 50 92 -50 c106 -57 143 -62 180 -25 30 29 30 31 8 167 l-17 108 82 84 c94 96 104 126 54 175 -23 24 -41 30 -123 41 -53 7 -101 15 -106 19 -5 3 -30 49 -55 102 -57 121 -92 148 -155 117z m108 -374 c13 -7 39 -13 57 -14 17 0 32 -3 32 -6 0 -4 -16 -25 -36 -47 -35 -39 -36 -41 -30 -98 l7 -59 -33 20 c-38 23 -83 25 -118 5 -38 -22 -43 -18 -36 29 8 49 -6 87 -43 120 -30 27 -27 33 20 39 46 6 81 34 100 79 l12 30 22 -42 c12 -23 33 -49 46 -56z" />
            <path
              d="M4397 1949 c-10 -6 -41 -57 -68 -112 -27 -56 -57 -105 -67 -108 -9 -4 -57 -12 -105 -19 -116 -16 -143 -38 -138 -114 0 -10 37 -56 82 -101 l80 -83 -15 -88 c-23 -132 -21 -168 10 -193 40 -31 68 -26 171 30 l94 50 94 -50 c111 -60 142 -64 179 -20 31 36 30 27 8 160 l-19 107 84 91 c93 102 100 125 56 170 -23 22 -43 30 -108 40 -143 22 -128 12 -176 118 -24 52 -54 104 -68 114 -26 21 -66 24 -94 8z m92 -362 c5 -5 33 -15 62 -21 l52 -12 -38 -42 c-36 -40 -37 -45 -33 -98 3 -37 1 -53 -6 -50 -6 2 -27 11 -46 19 -35 15 -50 13 -121 -17 -14 -6 -16 -2 -9 35 10 52 -1 83 -43 127 l-31 32 27 0 c52 0 95 26 118 71 l21 43 19 -39 c11 -22 23 -43 28 -48z" />
            <path
              d="M1436 1229 c-10 -6 -40 -57 -66 -114 -54 -114 -43 -106 -179 -126 -101 -14 -131 -36 -131 -95 0 -32 9 -46 81 -120 l81 -83 -16 -94 c-22 -129 -21 -153 9 -182 35 -36 66 -31 174 26 l96 52 92 -52 c117 -66 158 -66 192 -2 11 19 10 44 -5 136 l-17 113 77 83 c83 90 95 114 76 158 -15 38 -38 48 -155 66 -54 9 -99 17 -100 18 -9 31 -95 196 -109 209 -22 20 -71 23 -100 7z m97 -356 c12 -13 37 -25 63 -29 24 -3 44 -9 44 -11 0 -3 -16 -24 -36 -46 -35 -38 -36 -42 -30 -93 7 -57 3 -62 -29 -40 -37 25 -70 26 -110 6 -22 -11 -40 -20 -42 -20 -1 0 -1 27 0 59 2 57 0 62 -36 99 l-38 39 42 7 c59 9 76 21 100 68 l21 42 15 -30 c9 -16 24 -39 36 -51z" />
            <path
              d="M3620 1128 c-17 -9 -43 -48 -78 -118 l-53 -105 -106 -17 c-122 -20 -153 -39 -153 -95 0 -27 14 -48 80 -120 44 -48 80 -92 80 -97 0 -6 -7 -49 -15 -96 -21 -119 -20 -142 10 -170 38 -36 77 -30 180 28 l90 50 70 -40 c138 -78 167 -81 209 -25 19 26 19 28 1 147 l-18 120 81 84 c91 93 101 122 59 171 -19 23 -37 29 -131 43 l-109 16 -47 101 c-26 55 -53 105 -61 112 -24 19 -63 24 -89 11z m71 -346 c12 -21 28 -32 50 -37 74 -15 74 -14 46 -43 -41 -44 -49 -69 -42 -123 l7 -50 -38 20 c-49 26 -70 26 -118 0 l-38 -20 6 54 c7 51 5 57 -28 100 -37 48 -37 57 0 57 41 0 75 23 96 64 22 43 23 43 59 -22z" />
            <path
              d="M2503 828 c-12 -13 -40 -63 -62 -112 -23 -50 -43 -91 -44 -92 -2 -2 -51 -11 -110 -20 -98 -15 -110 -19 -133 -47 -40 -47 -28 -77 66 -172 l79 -80 -19 -114 c-20 -122 -17 -148 22 -175 36 -25 64 -19 159 34 49 28 93 50 99 50 6 0 50 -22 99 -50 95 -53 123 -59 159 -34 39 27 43 54 22 176 l-20 113 74 75 c84 85 99 109 90 144 -11 46 -43 63 -150 80 -56 9 -104 16 -106 16 -3 0 -24 42 -48 93 -53 115 -72 137 -120 137 -25 0 -43 -7 -57 -22z m75 -297 c15 -42 48 -68 98 -76 l45 -6 -37 -41 c-35 -39 -36 -43 -32 -99 2 -33 3 -59 1 -59 -1 0 -23 10 -48 22 l-45 22 -45 -22 c-25 -12 -47 -22 -48 -22 -2 0 -1 26 1 59 4 56 3 60 -32 99 l-37 41 45 6 c50 8 83 34 98 76 5 16 14 29 18 29 4 0 13 -13 18 -29z" />
          </g>
        </svg>
      </v-btn>
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