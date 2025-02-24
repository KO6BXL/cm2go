package main

import (
	"log"

	"github.com/ko6bxl/cm2go/block"
	"github.com/ko6bxl/cm2go/build"
)

//this func is waaay to long
//I tried putting everything into a map, but golang hates me :(

// TODO: Make this function cleaner
// TODO: Test connections
func main() {
	//test blocks
	log.Println("starting test")
	log.Println("Starting with blocks")
	//test nor
	var norcol block.Collection
	nor := norcol.Append(block.NOR())
	nor.Offset.X = 1
	nor.Offset.Y = 2
	nor.Offset.Z = -3
	log.Println("Compiling NOR...")
	norstr, err := build.Compile([]block.Collection{norcol})

	if err != nil {
		log.Fatal(err)
	}

	if strVal(norstr) {
		log.Println("NOR SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}
	//test and
	var andcol block.Collection
	and := andcol.Append(block.AND())
	and.Offset.X = 1
	and.Offset.Y = 2
	and.Offset.Z = -3
	log.Println("Compiling AND...")
	andstr, err := build.Compile([]block.Collection{andcol})

	if err != nil {
		log.Fatal(err)
	}
	if strVal(andstr) {
		log.Println("AND SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}
	//test or
	var orcol block.Collection
	or := orcol.Append(block.OR())
	or.Offset.X = 1
	or.Offset.Y = 2
	or.Offset.Z = -3
	log.Println("Compiling OR...")
	orstr, err := build.Compile([]block.Collection{orcol})

	if err != nil {
		log.Fatal(err)
	}
	if strVal(orstr) {
		log.Println("OR SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}
	//test xor
	var xorcol block.Collection
	xor := xorcol.Append(block.XOR())
	xor.Offset.X = 1
	xor.Offset.Y = 2
	xor.Offset.Z = -3
	log.Println("Compiling XOR...")
	xorstr, err := build.Compile([]block.Collection{xorcol})
	if err != nil {
		log.Fatal(err)
	}
	if strVal(xorstr) {
		log.Println("XOR SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}
	//test button
	var buttoncol block.Collection
	button := buttoncol.Append(block.BUTTON())
	button.Offset.X = 1
	button.Offset.Y = 2
	button.Offset.Z = -3
	log.Println("Compiling button...")
	buttonstr, err := build.Compile([]block.Collection{buttoncol})
	if err != nil {
		log.Fatal(err)
	}
	if strVal(buttonstr) {
		log.Println("Button SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}
	//test flipflop
	var flipflopcol block.Collection
	flipflop := flipflopcol.Append(block.FLIPFLOP())
	flipflop.Offset.X = 1
	flipflop.Offset.Y = 2
	flipflop.Offset.Z = -3
	log.Println("Compiling FlipFlop...")
	flipflopstr, err := build.Compile([]block.Collection{flipflopcol})
	if err != nil {
		log.Fatal(err)
	}
	if strVal(flipflopstr) {
		log.Println("FlipFlop SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}

	//test flipflop state
	flipflop.State = true
	flipflop.Offset.X = 2
	flipflop.Offset.Y = 3
	flipflop.Offset.Z = 4
	log.Println("Compiling FlipFlop w/ State = true...")
	flipflopstr, err = build.Compile([]block.Collection{flipflopcol})
	if err != nil {
		log.Fatal(err)
	}
	if strVal(flipflopstr) {
		log.Println("FlipFlop SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}
	//test flipflop props
	flipflop.Offset.X = 3
	flipflop.Offset.Y = 4
	flipflop.Offset.Z = 5
	log.Println("Compiling FlipFlop w/ Properties = true...")
	flipflopstr, err = build.Compile([]block.Collection{flipflopcol})
	if err != nil {
		log.Fatal(err)
	}
	if strVal(flipflopstr) {
		log.Println("FlipFlop SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}
	//test led
	var ledcol block.Collection
	ledcolor := block.Color{
		R: 100,
		G: 150,
		B: 200,
	}
	ledparam := block.LEDParams{
		Color:      ledcolor,
		OpacityOn:  100,
		OpacityOff: 25,
		Analog:     true,
	}
	led := ledcol.Append(block.LED(&ledparam))
	led.Offset.X = 1
	led.Offset.Y = 2
	led.Offset.Z = -3
	log.Println("Compiling LED...")
	ledstr, err := build.Compile([]block.Collection{ledcol})
	if err != nil {
		log.Fatal(err)
	}
	if strVal(ledstr) {
		log.Println("LED SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}
	//test led props
	led.Offset.X = 3
	led.Offset.Y = 4
	led.Offset.Z = 5
	log.Println("Compiling LED w/ props...")
	ledstr, err = build.Compile([]block.Collection{ledcol})
	if err != nil {
		log.Fatal(err)
	}
	if strVal(ledstr) {
		log.Println("LED SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}
	//test sound
	var soundcol block.Collection
	sound := soundcol.Append(block.SOUND(555, 1))
	sound.Offset.X = 1
	sound.Offset.Y = 2
	sound.Offset.Z = -3
	log.Println("Compiling sound...")
	soundstr, err := build.Compile([]block.Collection{soundcol})
	if err != nil {
		log.Fatal(err)
	}
	if strVal(soundstr) {
		log.Println("Sound SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}
	//test sound props
	sound.Offset.X = 3
	sound.Offset.Y = 4
	sound.Offset.Z = 5
	log.Println("Compiling sound w/ props ...")
	soundstr, err = build.Compile([]block.Collection{soundcol})
	if err != nil {
		log.Fatal(err)
	}
	if strVal(soundstr) {
		log.Println("Sound SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}
	//test conductor
	var condcol block.Collection
	cond := condcol.Append(block.CONDUCTOR())
	cond.Offset.X = 1
	cond.Offset.Y = 2
	cond.Offset.Z = -3
	log.Println("Compiling conductor...")
	condstr, err := build.Compile([]block.Collection{condcol})
	if err != nil {
		log.Fatal(err)
	}
	if strVal(condstr) {
		log.Println("Conductor SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}
	//test custom
	var custcol block.Collection
	cust := custcol.Append(block.CUSTOM())
	cust.Offset.X = 1
	cust.Offset.Y = 2
	cust.Offset.Z = -3
	log.Println("Compiling Custom...")
	custstr, err := build.Compile([]block.Collection{custcol})
	if err != nil {
		log.Fatal(err)
	}
	if strVal(custstr) {
		log.Println("Custom SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}
	//test nand
	var nandcol block.Collection
	nand := nandcol.Append(block.NAND())
	nand.Offset.X = 1
	nand.Offset.Y = 2
	nand.Offset.Z = -3
	log.Println("Compiling NAND...")
	nandstr, err := build.Compile([]block.Collection{nandcol})
	if err != nil {
		log.Fatal(err)
	}
	if strVal(nandstr) {
		log.Println("NAND SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}
	//test xnor
	var xnorcol block.Collection
	xnor := xnorcol.Append(block.XNOR())
	xnor.Offset.X = 1
	xnor.Offset.Y = 2
	xnor.Offset.Z = -3
	log.Println("Compiling XNOR...")
	xnorstr, err := build.Compile([]block.Collection{xnorcol})
	if err != nil {
		log.Fatal(err)
	}
	if strVal(xnorstr) {
		log.Println("XNOR SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}
	//test rand
	var randcol block.Collection
	rand := randcol.Append(block.RANDOM())
	rand.Offset.X = 1
	rand.Offset.Y = 2
	rand.Offset.Z = -3
	log.Println("Compiling random...")
	randstr, err := build.Compile([]block.Collection{randcol})
	if err != nil {
		log.Fatal(err)
	}
	if strVal(randstr) {
		log.Println("Random SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}
	//test text (lol)
	var textcol block.Collection
	text := textcol.Append(block.TEXT(80))
	text.Offset.X = 1
	text.Offset.Y = 2
	text.Offset.Z = -3
	log.Println("Compiling text...")
	textstr, err := build.Compile([]block.Collection{textcol})
	if err != nil {
		log.Fatal(err)
	}
	if strVal(textstr) {
		log.Println("Text SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}
	//test text props
	text.Offset.X = 3
	text.Offset.Y = 4
	text.Offset.Z = 5
	log.Println("Compiling text w/ props...")
	textstr, err = build.Compile([]block.Collection{textcol})
	if err != nil {
		log.Fatal(err)
	}
	if strVal(textstr) {
		log.Println("Text SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}
	//test tile
	var tilecol block.Collection
	tile := tilecol.Append(block.TILE(block.Color{R: 5, G: 5, B: 5}, 5))
	tile.Offset.X = 1
	tile.Offset.Y = 2
	tile.Offset.Z = -3
	log.Println("Compiling tile...")
	tilestr, err := build.Compile([]block.Collection{tilecol})
	if err != nil {
		log.Fatal(err)
	}
	if strVal(tilestr) {
		log.Println("Tile SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}
	//test tile props
	tile.Offset.X = 3
	tile.Offset.Y = 4
	tile.Offset.Z = 5
	log.Println("Compiling tile w/ props...")
	tilestr, err = build.Compile([]block.Collection{tilecol})
	if err != nil {
		log.Fatal(err)
	}
	if strVal(tilestr) {
		log.Println("Tile SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}
	//test node
	var nodecol block.Collection
	node := nodecol.Append(block.NODE())
	node.Offset.X = 1
	node.Offset.Y = 2
	node.Offset.Z = -3
	log.Println("Compiling node...")
	nodestr, err := build.Compile([]block.Collection{nodecol})
	if err != nil {
		log.Fatal(err)
	}
	if strVal(nodestr) {
		log.Println("Node SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}
	//test delay
	var delaycol block.Collection
	delay := delaycol.Append(block.DELAY(5))
	delay.Offset.X = 1
	delay.Offset.Y = 2
	delay.Offset.Z = -3
	log.Println("Compiling delay...")
	delaystr, err := build.Compile([]block.Collection{delaycol})
	if err != nil {
		log.Fatal(err)
	}
	if strVal(delaystr) {
		log.Println("Delay SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}
	//test delay props
	delay.Offset.X = 3
	delay.Offset.Y = 4
	delay.Offset.Z = 5
	log.Println("Compiling delay w/ props...")
	delaystr, err = build.Compile([]block.Collection{delaycol})
	if err != nil {
		log.Fatal(err)
	}
	if strVal(delaystr) {
		log.Println("Delay SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}
	//test antenna
	var antcol block.Collection
	ant := antcol.Append(block.ANTENNA(69))
	ant.Offset.X = 1
	ant.Offset.Y = 2
	ant.Offset.Z = -3
	log.Println("Compiling antenna...")
	antstr, err := build.Compile([]block.Collection{antcol})
	if err != nil {
		log.Fatal(err)
	}
	if strVal(antstr) {
		log.Println("Antenna SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}
	//test antenna props
	ant.Offset.X = 3
	ant.Offset.Y = 4
	ant.Offset.Z = 5
	log.Println("Compiling antenna...")
	antstr, err = build.Compile([]block.Collection{antcol})
	if err != nil {
		log.Fatal(err)
	}
	if strVal(antstr) {
		log.Println("Antenna SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}
	//test conductorV2
	var cond2col block.Collection
	cond2 := cond2col.Append(block.CONDUCTORV2())
	cond2.Offset.X = 1
	cond2.Offset.Y = 2
	cond2.Offset.Z = -3
	log.Println("Compiling ConductorV2...")
	cond2str, err := build.Compile([]block.Collection{cond2col})
	if err != nil {
		log.Fatal(err)
	}
	if strVal(cond2str) {
		log.Println("ConductorV2 SUCCESS!")
	} else {
		log.Fatal("FAILED!")
	}
	for _ = range 5 {
		log.Println("TEST SUCCESS!")
	}
}

// validate strings func
func strVal(str string) bool {
	//setup parsing functions
	var word string
	var words []string

	//parse basic params
	for _, x := range str {

		if x == ',' || x == '?' {
			words = append(words, word)
			word = ""
		} else {
			word += string(x)
		}
	}
	//store params in easy to identify vars
	id := words[0]
	state := words[1]
	X := words[2]
	Y := words[3]
	Z := words[4]
	//raw properties, not separated
	rprops := words[5]

	//setup parse vars for properties
	var prop string
	var props []string

	//parse through raw properties
	for _, x := range rprops {

		if x == '+' {
			props = append(props, prop)
			prop = ""
		} else {
			prop += string(x)
		}
	}
	//extracts the final prop or the only prop
	props = append(props, prop)

	//debug
	log.Println(str)
	log.Println(words)
	log.Println(props)
	switch X {
	// 1,2,-3 means test translation
	// 2,3,4 means test state
	// 3,4,5 means test properties

	//test tanslations
	case "1":
		//tests if compiler correctly added translations
		if Y == "2" && Z == "-3" {
			return true
		} else {
			log.Println(words)
			return false
		}
	//test if state is valid
	case "2":
		//tests if compiler set state correctly
		if state == "1" {
			return true
		} else {
			log.Println(words)
			return false
		}
	//tests if compiler set the properties
	case "3":
		switch id {
		//flipflop
		case "5":
			if props[0] == "2" && props[1] == "0" {
				return true
			} else {
				log.Println(words)
				return false
			}
		//led
		case "6":
			testSuccess := true
			ledTestProps := []string{"100", "150", "200", "100", "25", "1"}
			for i, ledp := range props {
				log.Println(ledp)
				if ledp != ledTestProps[i] {
					testSuccess = false
				}
			}
			if testSuccess {
				return true
			} else {
				log.Println(words)
				return false
			}
		//sound
		case "7":
			if props[0] == "555" && props[1] == "1" {
				return true
			} else {
				log.Println(words)
				return false
			}
		//text
		case "13":
			if props[0] == "80" {
				return true
			} else {
				log.Println(words)
				return false
			}
		//tile
		case "14":
			testsuccess := true
			tiletestprops := []string{"5", "5", "5", "5"}

			for i, prop := range props {
				if prop != tiletestprops[i] {
					testsuccess = false
				}
			}

			if testsuccess {
				return true
			} else {
				log.Println(words)
				return false
			}
		//delay
		case "16":
			if props[0] == "5" {
				return true
			} else {
				log.Println(words)
				return false
			}
		//antenna
		case "17":
			if props[0] == "69" {
				return true
			} else {
				log.Println(words)
				return false
			}
		//these shouldn't be run unless some bs happens
		default:
			log.Println("NO ID FOUND! Couldn't pick a property test.")
			return false
		}
	default:
		log.Println("NO X FOUND! Couldn't pick a test")
		return false
	}

}
