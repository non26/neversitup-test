package smile_face

import "strings"

type smileFace struct {
}

func NewSmileFace() *smileFace {
	return &smileFace{}
}

func (s *smileFace) CountSmileFaces(input []string) int {
	count := 0
	for _, smileface := range input {
		lenSmileFace := len(smileface)
		eye := string(smileface[0])
		if lenSmileFace == 2 {
			mouth := string(smileface[1])
			if s.ruleForEyes(eye) && s.ruleForMouth(mouth) {
				count++
			}
		} else if lenSmileFace == 3 {
			nose := string(smileface[1])
			mouth := string(smileface[2])
			if s.ruleForEyes(eye) && s.ruleForNose(nose) && s.ruleForMouth(mouth) {
				count++
			}
		}
	}
	return count
}

func (s *smileFace) ruleForEyes(input string) bool {
	result := strings.Contains(input, ":") || strings.Contains(input, ";")
	return result
}

func (s *smileFace) ruleForNose(input string) bool {
	result := strings.Contains(input, "-") || strings.Contains(input, "~")
	return result
}

func (s *smileFace) ruleForMouth(input string) bool {
	result := strings.Contains(input, ")") || strings.Contains(input, "D")
	return result
}
