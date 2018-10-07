package main

// concept: let's say commond divider is x and n1 >= n2
// x is GCD of n1 and n2
// and x is GCD of n2 and (n1-n2) too
// because n1 = (x * d1), n2 = (x * d2), so n1-n2 = x * (d1-d2)

func greatestCommonDivider(n1 int, n2 int) int {
	if n2 > n1 {
		return greatestCommonDivider(n2, n1)
	}

	if n2 == 0 { // when it's devisible, that mean n2 of the previous call (n1 this call) is GCD
		return n1
	}
	return greatestCommonDivider(n2, n1%n2)
}

func main() {

}
