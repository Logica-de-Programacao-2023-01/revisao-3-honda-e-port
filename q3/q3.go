package q3

//Você está subindo uma escada. Leva n passos para chegar ao topo.
//
//A cada vez, você pode subir 1 ou 2 degraus. De quantas maneiras distintas você pode subir até o topo?

func ClimbStairs(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	maneiras := 0
	if n > 2 {
		maneiras = ClimbStairs(n-1) + ClimbStairs(n-2)
		return maneiras
	}
	return 0
}
