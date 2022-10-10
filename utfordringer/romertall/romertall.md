# Romertall

Når man skriver romertall så bruker man syv symboler: `I`, `V`, `X`, `L`, `C`, `D` og `M`. Disse har følgende verdier:

````
I	= 1
V	= 5
X	= 10
L	= 50
C	= 100
D	= 500
M	= 1000
````

### Regler
1. Hvis symbolene er sotert fra høyest til lavest verdi så er tallet bare summen av symbolene:

    `MCLIII = 1000 + 100 + 50 + 1 + 1 + 1 = 1153`

2. Hvis symbolet til venstre har lavere verdi enn symbolet til høyre, så substraherer man symbolet til venstre sin verdi fra symbolet til høyre sin verdi. Det er 6 slike tilfeller hvor substraksjon kan oppstå:

   * `I` plassert forran `V` (5) og `X` (10) gir henholdsvis 4 og 9. 
   * `X` plassert forran `L` (50) og `C` (100) gir henholdsvis 40 og 90. 
   * `C` plassert forran `D` (500) og `M` (1000) gir henholdsvis 400 og 900. 

   </br>

   Eksempel:

   *  `XLI = (-10 + 50) + 1 = 41`
   *  `XCIX = (-10 + 100) + (-1 + 10) = 99`

### Oppgave
Lag en funksjon `toRoman(int number)` som tar inn et heltall og retunerer tallet som romertall.