// Package fate provides BaZi (八字) analysis and Five Element (五行) calculations
// for Chinese name selection (起名). It calculates the Four Pillars, determines
// the Day Master's strength, and identifies favorable and unfavorable elements
// using either the Balance (平衡用神) or Pattern (格局用神) method.
//
// # Usage
//
//	input := &fate.FateInput{
//	    BirthDate: time.Now(),
//	    Gender:    1,
//	    Method:    fate.XiYongMethodBalance,
//	}
//	data, err := fate.GetFateData(input)
//	fmt.Println(data.WuxingXiji.UsefulElement)
package fate
