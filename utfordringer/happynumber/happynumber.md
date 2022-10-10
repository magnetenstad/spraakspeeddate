# Happy number

Skriv en funksjon som tar inn et positivt heltall n, og avgjør om tallet er `happy`.

Et tall er `happy` dersom det oppfyller disse kriteriene:

- Start med et positivt heltall og legg sammen kvadratet til hvert siffer.

1234 = 1<sup>2</sup> + 2<sup>2</sup> + 3<sup>2</sup> + 4<sup>2</sup> = 30

- Repeter prosessen enten til tallet når 1, eller om det fortsetter i en uendelig sykel.

- Tall som når 1 er `happy`

Funksjonen skal returnere `true` dersom inputen var `happy`, og `false` ellers.

## Tester

- 19 -> true

- 2 -> false

- 338 -> true

- 99 -> false

- 100 -> true
