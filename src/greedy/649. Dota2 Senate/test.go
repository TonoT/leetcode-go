package main

import "fmt"

//In the world of Dota2, there are two parties: the Radiant and the Dire.
//
//The Dota2 senate consists of senators coming from two parties. Now the senate wants to make a decision about a change in the Dota2 game. The voting for this change is a round-based procedure. In each round, each senator can exercise one of the two rights:
//
//Ban one senator's right:
//A senator can make another senator lose all his rights in this and all the following rounds.
//Announce the victory:
//If this senator found the senators who still have rights to vote are all from the same party, he can announce the victory and make the decision about the change in the game.
//
//
//Given a string representing each senator's party belonging. The character 'R' and 'D' represent the Radiant party and the Dire party respectively. Then if there are n senators, the size of the given string will be n.
//
//The round-based procedure starts from the first senator to the last senator in the given order. This procedure will last until the end of voting. All the senators who have lost their rights will be skipped during the procedure.
//
//Suppose every senator is smart enough and will play the best strategy for his own party, you need to predict which party will finally announce the victory and make the change in the Dota2 game. The output should be Radiant or Dire.
func main() {
	fmt.Println(predictPartyVictory("RRRDDDDD"))
}
func predictPartyVictory(senate string) string {
	var s []byte
	m, n, cm, cn := 0, 0, 0, 0
	for _, v := range senate {
		if v == 'R' {
			if n > 0 {
				n--
			} else {
				s = append(s, byte(v))
				m++
				cm++
			}
		} else {
			if m > 0 {
				m--
			} else {
				s = append(s, byte(v))
				n++
				cn++
			}
		}
	}
	if cm != 0 && cn != 0 {
		if m > n {
			for m > 0 {
				for i, v := range s {
					if v == 'D' {
						s = append(s[:i], s[i+1:]...)
						break
					}
				}

				m--
			}
		} else {
			for n > 0 {
				for i, v := range s {
					if v == 'R' {
						s = append(s[:i], s[i+1:]...)
						break
					}
				}
				n--
			}

		}
		return predictPartyVictory(string(s))
	}
	if cm == 0 {
		return "Dire"
	} else {
		return "Radiant"
	}

}
