# Pascals trekant

Pascals trekant er uendelig mange rader av tall som danner en trekant. Hver rad inneholder ett tall mer enn den forrige. Hver rad har enere på kantene. Innenfor enerne er tallene summen av de to tallene i raden over som står på hver sin side av tallet.

![Pascals trekant GIF](PascalTriangleAnimated2.gif)

Dette er veldig nyttig i kombinatorikk og algebra siden tallene i rad $n$ kan uttrykkes som

$$\binom{n}{0} \quad \binom{n}{1} \quad \cdots \quad \binom{n}{n-1} \quad \binom{n}{n} $$

## Oppgave

Lag en funksjon som returnerer en todimensjonal liste med de $n$ første radene i Pascals trekant.

### Eksempel 1

```python
Input: n = 1
Output: [[1]]
```

### Eksempel 2

```python
Input: n = 5
Output: [[1], [1, 1], [1, 2, 1], [1, 3, 3, 1], [1, 4, 6, 4, 1]]
```
