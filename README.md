# Övningsuppgifter Python + GO

## Funktionssyntax
En funktion i Python definieras enligt följande:
```
def my_function(argument1, argument2):
    # do something
    return output
```
## Övningar

### Easy
#### E 1.1 Ålderskontroll (If-else)
Skriv en funktion som tar in användarens ålder och sedan skriver ut om de är minderåriga (under 18 år), vuxna (mellan 18 och 65) eller pensionärer (över 65).

#### E 1.2 Jämnt eller udda (If-else)
Skriv en funktion som tar in en enskild siffra och skriver ut om den är udda eller jämn.

#### E 1.3 - Intervall (If-else)
Skapa funktionen `in_range(lower, upper)` som avgör ifall ett tal är inom det angivna intervallet. Om talet finns i ändpunkterna räknas det som inom intervallet.

#### E 1.4 - Favoritfrukter (Listor)
Skriva en funktion som:
1. Skapar en lista som innehåller fem element, som är dina favoritfrukter.
2. Lägger till en ny frukt i listan.
3. Ändrar det tredje elementet till en annan frukt.
4. Tar bort det sista elementet i listan.
5. Skriver ut varje frukt i listan på en ny rad.

#### E 1.5 List slicing (Listor)
Skriva en funktion som:
1. Skapar en lista som innehåller numren 1 till 10.
2. Använder list slicing för att skriva ut de första fem numren.
3. Använder list slicing för att skriva ut de sista fyra numren.
4. Skriver ut varannat nummer i listan, börjande med det första.
5. Skapar en ny lista som är en kopia av den omvända original listan.

#### E 1.6 Minst och störst (Listor)
1. Skapar en lista med 10 slumpmässiga heltal.
2. Letar reda på en inbyggd funktion för att hitta det största talet.
3. Letar reda på en funktion för att hitta det minsta talet.
4. Hittar summan av alla tal i listan.
5. Sorterar listan från lägst till högst.

#### E 1.7 Kombinera listor (Listor)
Skapa en funktion som tar två listor med heltal som argument och returnerar en kombinerad och sorterad lista av dessa.

#### E 1.8 Jämna tal i listor (Listor, Loopar)
Skapa en funktion som tar in en lista med helt som argument, och returnerar hur många jämna tal det finns i listan.

#### E 1.9 Summan av två listor (Listor, Loopar)
Skapa en funktion som tar in två lika långa listor som argument. Returnera en lista där varje element är summan av de två listornas respektive element.

#### E 1.10 Grundläggande nyckel-värde-parning

- Skapa en dictionary som representerar en person; den ska innehålla förnamn, efternamn, ålder och e-postadress.
- Skriv ut varje informationsdel individuellt med hjälp av nycklarna.
- Lägg till ett nytt nyckel-värde-par som representerar personens hemstad.
- Kontrollera om personen har en nyckel "mellannamn". Om inte, lägg till det med ett värde efter eget val.
- Uppdatera personens ålder genom att öka den med ett.

#### E 1.11 Loopa genom Dictionaries (Dictionaries, Loopar)

- Skapa en dictionary programmatiskt (ej manuellt) där nycklarna är nummer mellan 1 och 15 och värdena är nycklarnas kvadrater.
- Skriv ut varje nyckel-värde-par på en ny rad.
- Beräkna och skriv ut summan av alla värden i dictionaryn.

#### E 1.12 Reading a File:

- Write a Python script that reads a file and prints its contents to the console. Create a text file with several lines of text for testing purposes.
- Additional: Modify the script to read the file line by line and print each one with its line number.
 

### Medium
#### M 1.1 - Nästlade listor
En nästlad lista är en lista inuti en annan lista.

Skriva en funktion som:
1. Skapa en nästlad lista där varje element är en lista som innehåller tre tal.
2. Skriver ut det första talet från varje nästlad lista med hjälp av en loop.
3. Lägger till ett fjärde tal i varje nästlad lista.
4. Använder en loop för att skriva ut summan av talen i varje nästlad lista.

#### M 1.2 Mest förekommande talet
Skapa en funktion som tar in lista som argument, och returnerar det värdet som förekommer flest gånger i listan.

#### M 1.3 Största siffran
Skapa en funktion som tar in en lista med heltal som argument och returnerar det största talet. Du får inte använda `max()`

#### M 1.4 Lista baklänges
Skapa en funktion som tar in två listor. Funktionen returnar sant/falskt om den ena listan är den andra baklänges.

#### M 1.5 FizzBuzz
FizzBuzz är ett klassiskt intervjuproblem.

Skriv en funktion som skriver ut alla tal från 1 till ett givet nummer. Om numret är delbart med 3 skriv ut "Fizz". Om numret är delbart med 5 skriv ut "Buzz". Om numret är delbart med 3 och 5 skriv ut "FizzBuzz".

#### M 1.6 Varannan bokstav
Skriv en funktion som tar in en sträng som argument och returnerar varannan bokstav.

#### M 1.7 Rövarspråket
Gör ett program som översätter en sträng till rövarspråket. (https://sv.wikipedia.org/wiki/R%C3%B6varspr%C3%A5ket)

#### M 1.8 Räkning med Dictionaries
Givet en lång sträng av ord, skapa en dictionary som visar antalet av varje ord i strängen.

#### M 1.9 Invertera en Dictionary
Skapa en dictionary med objekt där nycklarna är unika, men värdena är inte det. Skriv en funktion för att invertera dictionaryn, genom att göra värdena från den ursprungliga dictionaryn till nycklar och nycklarna från den ursprungliga dictionaryn till värden i den nya. Eftersom nycklar är unika kan vissa värden komma att bli listor.

#### M 1.10 Copying a File

- Write a program that copies the contents of one file to another file.
- Additional: Modify the script to only copy lines that contain a specific word.

### Hard
Många av dessa övningar är lånade, varför engelska förekommer som instruktionsspråk ibland.

#### H 1.1 Two Sum (Leetcode #1)
Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:

Input: `nums = [2,7,11,15]`, `target = 9`

Output: `[0,1]`

Output: Because `nums[0] + nums[1] == 9`, we return `[0, 1]`.

Example 2: 

Input: `nums = [3,2,4]`, `target = 6`

Output: `[1,2]`

Example 3:

Input: `nums = [3,3]`, `target = 6`

Output: `[0,1]`

Constraints:
- `2 <= nums.length <= 10^4`
- `-10^9 <= nums[i<= 10^9]`
- `-10^9 <= target <= 10^9`
- `Only one valid answer exists.`

#### H 1.2 Roman to Integer
https://leetcode.com/problems/roman-to-integer/

#### H 1.3 Happy numbers
https://leetcode.com/problems/happy-number/


### Extreme
#### Ex 1.1 Longest Unique Substring (Leetcode#3)
Given a string `s`, find the length of the **longest** substring without repeating characters.

Example 1:

Input: `s = "abcabcbb"`

Output: `3`

Explanation: The answer is "abc", with the length of 3.

Example 2:

Input: `s = "bbbbb"`

Output: `1`

Explanation: The answer is "b", with the length of 1.

Example 3:

Input: `s = "pwwkew"`

Output: `3`

Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Example 4:

Input: `s` = `""`

Output: `0`

Constraints:

- `0 <= s.length <= 5 * 10^4`
- s consists of English letters, digits, symbols and spaces.

#### Ex 1.2 Integer to Roman
https://leetcode.com/problems/integer-to-roman/


### Nightmare
#### N 1.1 Median of Two Sorted Arrays
https://leetcode.com/problems/median-of-two-sorted-arrays/

#### N 1.2 Longest valid parenthesis
https://leetcode.com/problems/longest-valid-parentheses/