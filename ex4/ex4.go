package ex4

import (
	"math"
)

func findMedianSortedArrays(a []int, b []int) float64 {
	if len(a) > len(b) {
		a, b = b, a
	}

	l, r, s, m, i, j := 0, len(a)-1, len(a)+len(b), (len(a)+len(b))/2, 0, 0

	for {
		i = int(math.Floor((float64(l) + float64(r)) / 2))
		j = m - i - 2
		al, bl, ar, br := math.MinInt, math.MinInt, math.MaxInt, math.MaxInt

		if i >= 0 && i < len(a) {
			al = a[i]
		}

		if i+1 < len(a) && i+1 >= 0 {
			ar = a[i+1]
		}

		if j >= 0 && j < len(b) {
			bl = b[j]
		}

		if j+1 < len(b) && j+1 >= 0 {
			br = b[j+1]
		}

		if al <= br && ar >= bl {
			if bl > al {
				al = bl
			}

			if br < ar {
				ar = br
			}

			if (s % 2) == 0 {
				return (float64(al) + float64(ar)) / 2
			} else {
				return float64(ar)
			}
		} else if al > br {
			r = i - 1
		} else {
			l = i + 1
		}
	}
}
