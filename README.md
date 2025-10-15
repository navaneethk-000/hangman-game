
README
md
6 of 6
# Hangman

1. Get secret word (assuming word is "elephant")
   - No capital letters
   - No punctuation
   - Minimum length is 6
    

2. Display

[v] Secret word : --------  (string)
[v] Guessed letters :       (array of characters)
[v] Chances remaining : 7   (uint)
    Correct guesses :       (array of characters)

3. User enters a letter
   1. The letter than the user enters is there in the secret word
      
      secret := "elephant"
      guess := e
      correctGuesses = ["e"]
      guessedLetters = ["e"]
      chancesRemaining = 7
      
      if all letters correctly guessed
         Game over. Win

   2. The letter than the user enters is not there in the secret word
      
      secret := "elephant"
      guess := x
      correctGuesses = ["e"]
      guessedLetters = ["e", "x"]
      chancesRemaining = 6
      
      if turns left becomes 0
         Game over. Loss

   3. The letter that the user enters is already guessed
      
      secret := "elephant"
      guess := e
      correctGuesses = ["e"]
      guessedLetters = ["e", "x"]
      chancesRemaining = 6