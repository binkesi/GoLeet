package weekgame

// https://leetcode.cn/contest/weekly-contest-307/problems/minimum-hours-of-training-to-win-a-competition/

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) (ans int) {
	preSum := make([]int, len(experience))
	preSum[0] = 2*experience[0] + 1
	for i := 1; i < len(experience); i++ {
		preSum[i] = preSum[i-1] + experience[i]
	}
	maxE := 0
	for i := 1; i < len(experience); i++ {
		if experience[i]-preSum[i-1] >= maxE {
			maxE = experience[i] - preSum[i-1] + 1
		}
	}
	if initialExperience < maxE+experience[0]+1 {
		ans = ans + maxE + (experience[0] + 1 - initialExperience)
	}
	energySum := 0
	for i := 0; i < len(energy); i++ {
		energySum += energy[i]
	}
	if initialEnergy <= energySum {
		ans = ans + energySum + 1 - initialEnergy
	}
	return
}
