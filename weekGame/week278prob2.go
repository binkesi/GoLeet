package weekgame

// https://leetcode-cn.com/contest/weekly-contest-278/problems/all-divisions-with-the-highest-score-of-a-binary-array/

func maxScoreIndices(nums []int) []int {
	lenN := len(nums)
	type scoreData struct {
		score int
		ind   []int
	}
	maxScore := scoreData{ind: []int{0}}
	for i := 0; i < lenN; i++ {
		if nums[i] == 1 {
			maxScore.score += 1
		}
	}
	curScore := maxScore.score
	for j := 0; j < lenN; j++ {
		if nums[j] == 0 {
			curScore += 1
		} else if nums[j] == 1 {
			curScore -= 1
		}
		if curScore > maxScore.score {
			maxScore.score = curScore
			maxScore.ind = []int{j + 1}
		} else if curScore == maxScore.score {
			maxScore.ind = append(maxScore.ind, j+1)
		}
	}
	return maxScore.ind
}
