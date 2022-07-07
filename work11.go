package main 

import "fmt"


func countContry() {
  data := []string{"Kazakhstan", "Czech Republic", "Guam", "Bermuda", "Guam", "Czech Republic", "Burkina Faso", "Bermuda", "Guam", "Czech Republic", "Bermuda", "Kazakhstan", "Somalia", "Burkina Faso", "Czech Republic", "Czech Republic", "Czech Republic", "United States of America", "United States of America", "Burkina Faso", "Somalia", "Kazakhstan", "Guam", "Czech Republic", "Singapore", "Saint Pierre and Miquelon", "Saint Pierre and Miquelon", "Guam", "Kazakhstan", "Burkina Faso", "Namibia", "Guam", "Guam", "Somalia", "Singapore", "Burkina Faso", "Czech Republic", "Saint Pierre and Miquelon", "Singapore", "Saint Pierre and Miquelon", "Kazakhstan", "Kazakhstan", "Singapore", "Czech Republic", "Bermuda", "Burkina Faso", "Kazakhstan", "United States of America", "Guam", "Somalia", "Bermuda", "Bermuda", "Saint Pierre and Miquelon", "Bermuda", "Somalia", "Singapore", "Saint Pierre and Miquelon", "Saint Pierre and Miquelon", "Somalia", "Guam", "Burkina Faso", "Saint Pierre and Miquelon", "Bermuda", "Czech Republic", "Kazakhstan", "Singapore", "Guam", "Nepal", "United States of America", "Nepal", "Nepal", "Singapore", "Guam", "Somalia", "Bermuda", "Guam", "Czech Republic", "Nepal", "Bermuda", "Singapore", "Guam", "Burkina Faso", "Czech Republic", "Nepal", "Kazakhstan", "Namibia", "United States of America", "Nepal", "Burkina Faso", "Somalia", "Bermuda", "United States of America", "Somalia", "Namibia", "Nepal", "United States of America", "Somalia", "Somalia", "Singapore", "Bermuda"}

  countMap := map[string]int{}
  for _, s := range data {
    if count , ok := countMap[s]; ok {
      fmt.Println(count)
    }
  }

}

