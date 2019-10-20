package main

import "fmt"

func main() {
	var sportskitSet []sportsKit

	redBall := &ball{2, "red"}
	sportskitSet = append(sportskitSet, redBall)

	blackBall := &ball{2, "black"}
	sportskitSet = append(sportskitSet, blackBall)

	whiteBat := &bat{2, "white"}
	sportskitSet = append(sportskitSet, whiteBat)

	duplicateSportsKitSet := duplicate(sportskitSet)

	for _, sportsKit := range duplicateSportsKitSet {
		fmt.Printf("%+v\n", sportsKit)
	}
}

func duplicate(sportsKitArray []sportsKit) []sportsKit {
	var duplicateSportsKitSet []sportsKit
	for _, sportsKit := range sportsKitArray {
		duplicateSportsKitSet = append(duplicateSportsKitSet, sportsKit.clone())
	}
	return duplicateSportsKitSet
}
