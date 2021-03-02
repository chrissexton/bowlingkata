Feature: Bowling scorer
    This feature scores a single player's bowling game.

    Scenario: Player rolls a gutterball game
        Given the score starts at 0
        When the player rolls a 0 20 times
        Then the score should be 0

    Scenario: Player rolls all 1s
        Given the score starts at 0
        When the player rolls a 1 20 times
        Then the score should be 20

    Scenario: The player rolls a spare
        Given the score starts at 0
        When the player rolls a
            | 5 |
            | 5 |
            | 3 |
            | 0 |
        And we print out the game
        Then the score should be 16

    Scenario: the player rolls a strike
        Given the score starts at 0
        When the player rolls a
            | 10 |
            | 3 |
            | 4 |
        And we print out the game
        Then the score should be 24

    Scenario: the player rolls a strike
        Given the score starts at 0
        When the player rolls a
            | 10 |
            | 10 |
            | 3 |
            | 4 |
        And we print out the game
        Then the score should be 47

    Scenario: the player rolls a perfect game
        Given the score starts at 0
        When the player rolls a 10 12 times
        Then the score should be 300