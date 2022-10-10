# Mange øyer
<p>
    <figure align="center">
        <img src="dalle2islands.png" alt="Mockup" width="700"/>
        <figcaption><i>Hentet fra DALL-E 2.</i></figcaption>
    </figure>
</p>

I den vedlagte [tekstfilen](kart.txt) får du et `m x n` binært rutenett som representerer et kart over øyer, hvor `1` er land og `0` er vann. Finn ut hvor mange forskjellige øyer kartet inneholder.

Ruter med land som er koblet sammen hotisontalt eller vertikalt utgjør en øy. Hvis to ruter med land står diagonalt ovenfor hverandre så teller det som 2 øyer. Du kan anta at alt utafor kartet er vann. 

*Eksempel:* 

<img src="eksempel.png" alt="Mockup" width="300"/>

Dette kartet består av 3 øyer (en rød rute representerer 3 øyer).

---


<details>
    <summary>Svar</summary>
    Kartet i tekstfilen består av 13 øyer!
</details>