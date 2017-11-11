package intervalselection

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
)

type interval struct {
	x, y int
}

type byY []*interval

func (a byY) Len() int      { return len(a) }
func (a byY) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byY) Less(i, j int) bool {
	return a[i].y < a[j].y || (a[i].y == a[j].y && a[i].x < a[j].x)
}

type byX []*interval

func (a byX) Len() int      { return len(a) }
func (a byX) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byX) Less(i, j int) bool {
	return a[i].x < a[j].x || (a[i].x == a[j].x && a[i].y < a[j].y)
}

func Solve(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	s, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < s; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		a := make([]interval, n)
		for j := 0; j < n; j++ {
			scanner.Scan()
			fmt.Sscanf(scanner.Text(), "%d %d", &a[j].x, &a[j].y)
		}
		xs := make([]*interval, n)
		ys := make([]*interval, n)
		for j := 0; j < n; j++ {
			xs[j], ys[j] = &a[j], &a[j]
		}
		sort.Sort(byX(xs))
		sort.Sort(byY(ys))
		xi, ans := 0, 0
		l := make([]*interval, 0)
		for _, inv := range ys {
			sort.Sort(byY(l))
			k := 0
			for ; k < len(l) && l[k].y < inv.y; k++ {
				ans++
			}
			l = l[k:]
			for ; xi < len(xs) && xs[xi].x <= inv.y; xi++ {
				l = append(l, xs[xi])
			}
			sort.Sort(byY(l))
			if len(l) > 2 {
				l = l[:2]
			}
		}
		for j := 0; j < len(l); j++ {
			ans++
		}
		fmt.Fprintf(w, "%d\n", ans)
	}
}

// func Solve(r io.Reader, w io.Writer) {
// 	scanner := bufio.NewScanner(r)
// 	scanner.Scan()
// 	s, _ := strconv.Atoi(scanner.Text())
// 	for i := 0; i < s; i++ {
// 		scanner.Scan()
// 		n, _ := strconv.Atoi(scanner.Text())
// 		a := make([]interval, n)
// 		for j := 0; j < n; j++ {
// 			scanner.Scan()
// 			fmt.Sscanf(scanner.Text(), "%d %d", &a[j].x, &a[j].y)
// 		}
// 		sort.Sort(byY(a))
// 		active := list.New()
// 		ans := make([]interval, 0)
// 		active.PushBack(a[0])
// 		for k := 1; k < n; k++ {
// 			var next *list.Element
// 			for e := active.Front(); e != nil; e = next {
// 				next = e.Next()
// 				if !e.Value.(interval).intersects(a[k]) {
// 					active.Remove(e)
// 					ans = append(ans, e.Value.(interval))
// 				}
// 			}
// 			active.PushBack(a[k])
// 			for active.Len() > 2 {
// 				max := active.Front()
// 				for e := max.Next(); e != nil; e = e.Next() {
// 					if e.Value.(interval).len() > max.Value.(interval).len() {
// 						max = e
// 					}
// 				}
// 				active.Remove(max)
// 			}
// 		}
// 		for e := active.Front(); e != nil; e = e.Next() {
// 			ans = append(ans, e.Value.(interval))
// 		}
// 		fmt.Fprintf(w, "%d\n", len(ans))
// 	}
//}
