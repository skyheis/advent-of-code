/*
Copyright © 2024 Giulio Giannitrapani <giuliogt7@gmail.com>
*/
package twentyfour

import (
	"advent-of-code/utils"
	"fmt"
	"os"
	"time"
	"unicode"

	"github.com/spf13/cobra"
)

// twentyfourCmd represents the twentyfour command
var twelfthCmd = &cobra.Command{
	Use:   "12th",
	Short: "Day 12 of Advent of Code 2024",
	Long: `Day 12 of Advent of Code 2024, finally my second Advent of Code!!!
Here the problem statement: https://adventofcode.com/2024/day/12`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Executing %s\n\n", cmd.Short)

		inputPath := "twentyfour/inputs/twelve.input"
		filecontent, err := os.ReadFile(inputPath)
		utils.CheckInputFileError(err)

		start := time.Now()
		result := dayTwelveCommonPart(filecontent, true)
		end := time.Since(start)

		if cmd.Flag("subject").Changed {
			twelfthSubject(true)
		}
		utils.PrintResult("one", result, end)

		start = time.Now()
		result = dayTwelveCommonPart(filecontent, false)
		end = time.Since(start)

		if cmd.Flag("subject").Changed {
			twelfthSubject(false)
		}
		utils.PrintResult("two", result, end)

	},
}

func init() {
	TwentyfourCmd.AddCommand(twelfthCmd)
}

func dayTwelveCommonPart(filecontent []byte, fences bool) (result int) {
	garden := utils.MakeRuneMatrixByte(filecontent)
	garden = garden[:len(garden)-1]

	meta := gardenNote{
		garden: garden,
		ai:     -1,
	}

	for y, line := range meta.garden {
		for x, plant := range line {
			if unicode.IsUpper(plant) {
				meta.addArea()
				if fences {
					countFences(meta, utils.Coordinates{X: x, Y: y}, plant)
				} else {
					countCorners(meta, utils.Coordinates{X: x, Y: y}, plant)
				}
				// fmt.Println("plants:", meta.areas[meta.ai].plants)
				// fmt.Println()
			}
		}
	}

	for _, area := range meta.areas {
		result += (area.fences * area.plants)
	}
	return
}

type gardenNote struct {
	garden [][]rune
	ai     int
	areas  []struct {
		plants int
		fences int
	}
}

func (g *gardenNote) addArea() {
	g.ai += 1
	g.areas = append(g.areas, struct {
		plants int
		fences int
	}{plants: 0, fences: 0})
}

func (g *gardenNote) addPlant(cur utils.Coordinates) {
	g.garden[cur.Y][cur.X] += 32
	g.areas[g.ai].plants += 1
}

func (g *gardenNote) addFence() {
	g.areas[g.ai].fences += 1
}

func (g *gardenNote) getPlantType(cur utils.Coordinates) rune {
	if utils.InMatrix(g.garden, cur) {
		return g.garden[cur.Y][cur.X]
	}
	return 0
}

func samePlants(a, b rune) bool {
	return a == b || a == b+32
}

func (g *gardenNote) checkFence(cur utils.Coordinates, plant rune) {

	directions := []utils.Coordinates{cur.Up(), cur.Down(), cur.Left(), cur.Right()}

	for _, direction := range directions {
		nextPlant := g.getPlantType(direction)
		if !samePlants(nextPlant, plant) {
			g.addFence()
		}
	}
}

func countFences(meta gardenNote, cur utils.Coordinates, plant rune) {
	if !utils.InMatrix(meta.garden, cur) {
		return
	}
	if meta.garden[cur.Y][cur.X] != plant {
		return
	}

	meta.addPlant(cur)
	meta.checkFence(cur, plant)

	countFences(meta, cur.Up(), plant)
	countFences(meta, cur.Down(), plant)
	countFences(meta, cur.Left(), plant)
	countFences(meta, cur.Right(), plant)
}

func (g *gardenNote) checkCorners(cur utils.Coordinates, plant rune) {

	direct := []utils.Coordinates{cur.Up(), cur.Right(), cur.Down(), cur.Left()}
	diagCorn := []func(utils.Coordinates) utils.Coordinates{
		utils.Coordinates.UpRight,
		utils.Coordinates.DownRight,
		utils.Coordinates.DownLeft,
		utils.Coordinates.UpLeft,
	}

	for i := 0; i < 4; i++ {
		if samePlants(g.getPlantType(direct[i]), plant) && samePlants(g.getPlantType(direct[(i+1)%4]), plant) && !samePlants(g.getPlantType(diagCorn[i](cur)), plant) {
			g.addFence()
		}
		if !samePlants(g.getPlantType(direct[i]), plant) && !samePlants(g.getPlantType(direct[(i+1)%4]), plant) {
			g.addFence()
		}
	}
}

func countCorners(meta gardenNote, cur utils.Coordinates, plant rune) {
	if !utils.InMatrix(meta.garden, cur) {
		return
	}
	if meta.garden[cur.Y][cur.X] != plant {
		return
	}

	meta.addPlant(cur)
	meta.checkCorners(cur, plant)

	countCorners(meta, cur.Up(), plant)
	countCorners(meta, cur.Down(), plant)
	countCorners(meta, cur.Left(), plant)
	countCorners(meta, cur.Right(), plant)
}

/*

if up ==, left != ; no
if up !=, left != ; convex
if up ==, left == ; is upleft == ? no, concave
if up !=, left == ; no

*/

func twelfthSubject(partOne bool) {
	if partOne {
		fmt.Println(`--- Day 12: Garden Groups ---
Why not search for the Chief Historian near the gardener and his massive farm? There's plenty of food, so The Historians grab something to eat while they search.

You're about to settle near a complex arrangement of garden plots when some Elves ask if you can lend a hand. They'd like to set up fences around each region of garden plots, but they can't figure out how much fence they need to order or how much it will cost. They hand you a map (your puzzle input) of the garden plots.

Each garden plot grows only a single type of plant and is indicated by a single letter on your map. When multiple garden plots are growing the same type of plant and are touching (horizontally or vertically), they form a region. For example:

	AAAA
	BBCD
	BBCC
	EEEC

This 4x4 arrangement includes garden plots growing five different types of plants (labeled A, B, C, D, and E), each grouped into their own region.

In order to accurately calculate the cost of the fence around a single region, you need to know that region's area and perimeter.

The area of a region is simply the number of garden plots the region contains. The above map's type A, B, and C plants are each in a region of area 4. The type E plants are in a region of area 3; the type D plants are in a region of area 1.

Each garden plot is a square and so has four sides. The perimeter of a region is the number of sides of garden plots in the region that do not touch another garden plot in the same region. The type A and C plants are each in a region with perimeter 10. The type B and E plants are each in a region with perimeter 8. The lone D plot forms its own region with perimeter 4.

Visually indicating the sides of plots in each region that contribute to the perimeter using - and |, the above map's regions' perimeters are measured as follows:

	+-+-+-+-+
	|A A A A|
	+-+-+-+-+     +-+
				|D|
	+-+-+   +-+   +-+
	|B B|   |C|
	+   +   + +-+
	|B B|   |C C|
	+-+-+   +-+ +
			|C|
	+-+-+-+   +-+
	|E E E|
	+-+-+-+

Plants of the same type can appear in multiple separate regions, and regions can even appear within other regions. For example:

	OOOOO
	OXOXO
	OOOOO
	OXOXO
	OOOOO

The above map contains five regions, one containing all of the O garden plots, and the other four each containing a single X plot.

The four X regions each have area 1 and perimeter 4. The region containing 21 type O plants is more complicated; in addition to its outer edge contributing a perimeter of 20, its boundary with each X region contributes an additional 4 to its perimeter, for a total perimeter of 36.

Due to "modern" business practices, the price of fence required for a region is found by multiplying that region's area by its perimeter. The total price of fencing all regions on a map is found by adding together the price of fence for every region on the map.

In the first example, region A has price 4 * 10 = 40, region B has price 4 * 8 = 32, region C has price 4 * 10 = 40, region D has price 1 * 4 = 4, and region E has price 3 * 8 = 24. So, the total price for the first example is 140.

In the second example, the region with all of the O plants has price 21 * 36 = 756, and each of the four smaller X regions has price 1 * 4 = 4, for a total price of 772 (756 + 4 + 4 + 4 + 4).

Here's a larger example:

	RRRRIICCFF
	RRRRIICCCF
	VVRRRCCFFF
	VVRCCCJFFF
	VVVVCJJCFE
	VVIVCCJJEE
	VVIIICJJEE
	MIIIIIJJEE
	MIIISIJEEE
	MMMISSJEEE

It contains:

 - A region of R plants with price 12 * 18 = 216.
 - A region of I plants with price 4 * 8 = 32.
 - A region of C plants with price 14 * 28 = 392.
 - A region of F plants with price 10 * 18 = 180.
 - A region of V plants with price 13 * 20 = 260.
 - A region of J plants with price 11 * 20 = 220.
 - A region of C plants with price 1 * 4 = 4.
 - A region of E plants with price 13 * 18 = 234.
 - A region of I plants with price 14 * 22 = 308.
 - A region of M plants with price 5 * 12 = 60.
 - A region of S plants with price 3 * 8 = 24.

So, it has a total price of 1930.

What is the total price of fencing all regions on your map?`)
	} else {
		fmt.Println(`--- Part Two ---
Fortunately, the Elves are trying to order so much fence that they qualify for a bulk discount!

Under the bulk discount, instead of using the perimeter to calculate the price, you need to use the number of sides each region has. Each straight section of fence counts as a side, regardless of how long it is.

Consider this example again:

	AAAA
	BBCD
	BBCC
	EEEC

The region containing type A plants has 4 sides, as does each of the regions containing plants of type B, D, and E. However, the more complex region containing the plants of type C has 8 sides!

Using the new method of calculating the per-region price by multiplying the region's area by its number of sides, regions A through E have prices 16, 16, 32, 4, and 12, respectively, for a total price of 80.

The second example above (full of type X and O plants) would have a total price of 436.

Here's a map that includes an E-shaped region full of type E plants:

	EEEEE
	EXXXX
	EEEEE
	EXXXX
	EEEEE

The E-shaped region has an area of 17 and 12 sides for a price of 204. Including the two regions full of type X plants, this map has a total price of 236.

This map has a total price of 368:

	AAAAAA
	AAABBA
	AAABBA
	ABBAAA
	ABBAAA
	AAAAAA

It includes two regions full of type B plants (each with 4 sides) and a single region full of type A plants (with 4 sides on the outside and 8 more sides on the inside, a total of 12 sides). Be especially careful when counting the fence around regions like the one full of type A plants; in particular, each section of fence has an in-side and an out-side, so the fence does not connect across the middle of the region (where the two B regions touch diagonally). (The Elves would have used the Möbius Fencing Company instead, but their contract terms were too one-sided.)

The larger example from before now has the following updated prices:

 - A region of R plants with price 12 * 10 = 120.
 - A region of I plants with price 4 * 4 = 16.
 - A region of C plants with price 14 * 22 = 308.
 - A region of F plants with price 10 * 12 = 120.
 - A region of V plants with price 13 * 10 = 130.
 - A region of J plants with price 11 * 12 = 132.
 - A region of C plants with price 1 * 4 = 4.
 - A region of E plants with price 13 * 8 = 104.
 - A region of I plants with price 14 * 16 = 224.
 - A region of M plants with price 5 * 6 = 30.
 - A region of S plants with price 3 * 6 = 18.

Adding these together produces its new total price of 1206.

What is the new total price of fencing all regions on your map?`)
	}
	fmt.Println()
}
