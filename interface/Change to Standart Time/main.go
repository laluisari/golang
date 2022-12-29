package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Hour   int
	Minute int
}

func ChangeToStandartTime(time interface{}) string {
	//xType := fmt.Sprintf("%T", time) //tipedata
	temp := ""
	switch time.(type) {
	case string:
		temp += FormatString(temp, time)
	case []int:
		temp += FormatSlice(temp, time)
	case map[string]int:
		data := time.(map[string]int)

		if len(data) == 2 {
			if _, ok := data["hour"]; ok {
				if _, ok := data["minute"]; ok {
					h := data["hour"]
					m := data["minute"]
					if h > 12 && m <= 60 {
						h = h - 12
						if h < 10 && m < 10 {
							temp += "0" + strconv.Itoa(h) + ":" + "0" + strconv.Itoa(m) + " PM"
						} else if h < 10 && m >= 10 {
							temp += "0" + strconv.Itoa(h) + ":" + strconv.Itoa(m) + " PM"
						} else {
							temp += strconv.Itoa(h) + ":" + strconv.Itoa(m) + " PM"
						}
					} else if h < 12 && m <= 60 {
						if h < 10 && m < 10 {
							temp += "0" + strconv.Itoa(h) + ":" + "0" + strconv.Itoa(m) + " AM"
						} else if h < 10 && m >= 10 {
							temp += "0" + strconv.Itoa(h) + ":" + strconv.Itoa(m) + " AM"
						} else {
							temp += strconv.Itoa(h) + ":" + strconv.Itoa(m) + "0" + " AM"
						}
					} else {
						temp += strconv.Itoa(h) + ":" + strconv.Itoa(m) + "0" + " PM"
					}
				} else {
					temp = "Invalid input"
				}
			} else {
				temp = "Invalid input"
			}

		} else {
			temp = "Invalid input"
		}

	//	temp += FormatMap(temp, time)
	case Time:
		data := time.(Time)
		h := data.Hour
		m := data.Minute
		if h > 12 && m <= 60 {
			h = h - 12
			if h < 10 && m < 10 {
				temp += "0" + strconv.Itoa(h) + ":" + "0" + strconv.Itoa(m) + " PM"
			} else if h < 10 && m >= 10 {
				temp += "0" + strconv.Itoa(h) + ":" + strconv.Itoa(m) + " PM"
			} else {
				temp += strconv.Itoa(h) + ":" + strconv.Itoa(m) + " PM"
			}
		} else if h < 12 && m <= 60 {
			if h < 10 && m < 10 {
				temp += "0" + strconv.Itoa(h) + ":" + "0" + strconv.Itoa(m) + " AM"
			} else if h < 10 && m >= 10 {
				temp += "0" + strconv.Itoa(h) + ":" + strconv.Itoa(m) + " AM"
			} else {
				temp += strconv.Itoa(h) + ":" + strconv.Itoa(m) + "0" + " AM"
			}
		} else {
			temp += strconv.Itoa(h) + ":" + strconv.Itoa(m) + "0" + " PM"
		}

	}
	// type1 := time.([]int)
	// type2 := time.(string)
	// type3 := time.(map[string]int)
	// type4 := time.(Time)

	return temp // TODO: replace this
}

func main() {
	//fmt.Println(ChangeToStandartTime("23:11")) //done
	//fmt.Println(ChangeToStandartTime([]int{23}))
	fmt.Println(ChangeToStandartTime(map[string]int{"minute": 23, "second": 12}))
	//fmt.Println(ChangeToStandartTime(Time{16, 0}))
}

func FormatString(temp string, time any) string {
	split := strings.Split(time.(string), ":")
	if len(split) == 2 {
		if split[0] == "" || split[1] == "" {
			temp = "Invalid input"
		} else {
			h, _ := strconv.Atoi(split[0])
			m, _ := strconv.Atoi(split[1])
			if h > 12 && m <= 60 {
				h = h - 12
				if h < 10 && m < 10 {
					temp += "0" + strconv.Itoa(h) + ":" + "0" + strconv.Itoa(m) + " PM"
				} else if h < 10 && m >= 10 {
					temp += "0" + strconv.Itoa(h) + ":" + strconv.Itoa(m) + " PM"
				} else {
					temp += strconv.Itoa(h) + ":" + strconv.Itoa(m) + " PM"
				}
			} else if h < 12 && m <= 60 {
				if h < 10 && m < 10 {
					temp += "0" + strconv.Itoa(h) + ":" + "0" + strconv.Itoa(m) + " AM"
				} else if h < 10 && m >= 10 {
					temp += "0" + strconv.Itoa(h) + ":" + strconv.Itoa(m) + " AM"
				} else {
					temp += strconv.Itoa(h) + ":" + strconv.Itoa(m) + "0" + " AM"
				}
			} else {
				temp += strconv.Itoa(h) + ":" + strconv.Itoa(m) + "0" + " PM"
			}
		}

	} else {
		temp = "Invalid input"
	}
	return temp
}

func FormatSlice(temp string, time any) string {
	data := time.([]int)
	if len(data) == 2 {
		h := data[0]
		m := data[1]
		if h > 12 && m <= 60 {
			h = h - 12
			if h < 10 && m < 10 {
				temp += "0" + strconv.Itoa(h) + ":" + "0" + strconv.Itoa(m) + " PM"
			} else if h < 10 && m >= 10 {
				temp += "0" + strconv.Itoa(h) + ":" + strconv.Itoa(m) + " PM"
			} else {
				temp += strconv.Itoa(h) + ":" + strconv.Itoa(m) + " PM"
			}
		} else if h < 12 && m <= 60 {
			if h < 10 && m < 10 {
				temp += "0" + strconv.Itoa(h) + ":" + "0" + strconv.Itoa(m) + " AM"
			} else if h < 10 && m >= 10 {
				temp += "0" + strconv.Itoa(h) + ":" + strconv.Itoa(m) + " AM"
			} else {
				temp += strconv.Itoa(h) + ":" + strconv.Itoa(m) + "0" + " AM"
			}
		} else {
			temp += strconv.Itoa(h) + ":" + strconv.Itoa(m) + "0" + " PM"
		}

	} else {
		temp = "Invalid input"
	}

	return temp
}

func FormatMap(temp string, time any) string {
	data := time.(map[string]int)
	if len(data) == 2 {
		h := data["hour"]
		m := data["minute"]
		if h > 12 && m <= 60 {
			h = h - 12
			if h < 10 && m < 10 {
				temp += "0" + strconv.Itoa(h) + ":" + "0" + strconv.Itoa(m) + " PM"
			} else if h < 10 && m >= 10 {
				temp += "0" + strconv.Itoa(h) + ":" + strconv.Itoa(m) + " PM"
			} else {
				temp += strconv.Itoa(h) + ":" + strconv.Itoa(m) + " PM"
			}
		} else if h < 12 && m <= 60 {
			if h < 10 && m < 10 {
				temp += "0" + strconv.Itoa(h) + ":" + "0" + strconv.Itoa(m) + " AM"
			} else if h < 10 && m >= 10 {
				temp += "0" + strconv.Itoa(h) + ":" + strconv.Itoa(m) + " AM"
			} else {
				temp += strconv.Itoa(h) + ":" + strconv.Itoa(m) + "0" + " AM"
			}
		} else {
			temp += strconv.Itoa(h) + ":" + strconv.Itoa(m) + "0" + " PM"
		}

	} else {
		temp = "Invalid input"
	}
	return temp
}
