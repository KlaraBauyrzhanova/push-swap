package internal

func PushSwap() *logger {

	avgB := 0
	algorithmSortA()
	peekCheck()

	for range b {
		for i := len(b) - 1; i >= 0; i-- {
			if len(b) > 2 {
				avgB = b.middleNumber()
			}
			for j := len(b) - 1; j > 0; j-- {
				if b[j] != b.maxElem() {
					if b[j] < avgB {
						b.ra()
						log.Add("rb")
						j++
					} else {
						a.pa(&b)
						log.Add("pa")
						i--
					}
				} else {
					a.pa(&b)
					log.Add("pa")
					i--
					algorithmSortA()
					break
				}
			}

		}
	}

	peekCheck()
	for i := len(b) - 1; i >= 0; i-- {
		a.pa(&b)
		log.Add("pa")
	}
	log.Add(a)
	return log
}

func peekCheck() {
	peekA, _ := a.peek()
	peekB, _ := b.peek()
	if peekA < peekB {
		a.pa(&b)
		a.sa()
		log.Add("pa")
		log.Add("sa")
	}
}

func CheckSolution(input []string) string {
	for idx := range input {
		switch input[idx] {
		case "pa":
			a.pa(&b)
		case "pb":
			a.pb(&b)
		case "ra":
			a.ra()
		case "rb":
			b.ra()
		case "sa":
			a.sa()
		case "sb":
			b.sa()
		case "ss":
			a.sa()
			b.sa()
		case "rr":
			a.ra()
			b.ra()
		case "rra":
			a.rr()
		case "rrb":
			b.rr()
		case "rrr":
			a.rr()
			b.rr()

		}
	}

	if a.isSort() && len(b) < 1 {
		return "OK"
	}
	return "KO"
}

func algorithmSortA() {
	msg := ""
	for i := len(a) - 1; i > 0; i-- {
		if a.isSort() {
			break
		}
		avgA := a.middleNumber()
		if a[i] > a[i-1] && a[i] != a.maxElem() {
			if b != nil {
				msg = b.algorithmSortB()
			}
			if msg == "sb" {
				a.sa()
				b.sa()
				log.Add("ss")
			} else {
				a.sa()
				log.Add("sa")
			}
		}
		if len(a) == 2 || a.isSort() {
			if !a.isSort() {
				a.sa()
				log.Add("sa")
			}
			break
		}
		if a[0] < avgA {
			if b != nil && b.algorithmSortB() == "rrb" {
				a.rr()
				b.rr()
				log.Add("rrr")
			} else {
				a.rr()
				log.Add("rra")
			}
		}
		if a[i] > avgA {
			if b != nil {
				msg = b.algorithmSortB()
			}
			if msg == "rb" {
				a.ra()
				b.ra()
				log.Add("rr")
			} else {
				a.ra()
				log.Add("ra")
			}
			i++
			if msg != "" {
				if msg == "sb" {
					b.sa()
					log.Add("sb")
				} else if msg == "rb" {
					b.ra()
					log.Add("rb")
				} else if msg == "rrb" {
					b.rr()
					log.Add("rrb")
				}
			}

		} else {
			a.pb(&b)
			log.Add("pb")
		}
	}
}
