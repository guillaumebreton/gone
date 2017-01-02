package painter

import "strings"

var colon string = `
..
#.
..
#.
..
`

var zero string = `
######.
#....#.
#....#.
#....#.
######.
`
var one string = `
.....#.
.....#.
.....#.
.....#.
.....#.
`
var two string = `
######.
.....#.
######.
#......
######.
`

var three string = `
######.
.....#.
...###.
.....#.
######.
`

var four string = `
#......
#......
#...#..
######.
....#..
`

var five string = `
######.
#......
######.
.....#.
######.
`

var six string = `
######.
#......
######.
#....#.
######.
`

var seven string = `
######.
.....#.
.....#.
.....#.
.....#.
`

var height string = `
######.
#....#.
######.
#....#.
######.
`

var nine string = `
######.
#....#.
######.
.....#.
######.
`

// smallFont defines the font use to display the timer on termbox
var smallFont = map[rune][][]rune{
	':': asArray(colon),
	'1': asArray(one),
	'2': asArray(two),
	'3': asArray(three),
	'4': asArray(four),
	'5': asArray(five),
	'6': asArray(six),
	'7': asArray(seven),
	'8': asArray(height),
	'9': asArray(nine),
	'0': asArray(zero),
}

// Convert a character as an array of rune
func asArray(chars string) [][]rune {
	result := [][]rune{}
	line := []rune{}
	str := strings.TrimPrefix(chars, "\n")
	for _, c := range str {
		if c == '\n' {
			result = append(result, line)
			line = []rune{}
		} else {
			line = append(line, c)
		}
	}
	return result
}
